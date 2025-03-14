package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/muhammedkucukaslan/library-management-service/app/auth"
	"github.com/muhammedkucukaslan/library-management-service/app/author"
	"github.com/muhammedkucukaslan/library-management-service/app/book"
	bookCategory "github.com/muhammedkucukaslan/library-management-service/app/book_category"
	"github.com/muhammedkucukaslan/library-management-service/app/healthcheck"
	"github.com/muhammedkucukaslan/library-management-service/app/loan"
	"github.com/muhammedkucukaslan/library-management-service/app/punishment"
	"github.com/muhammedkucukaslan/library-management-service/app/user"
	"github.com/muhammedkucukaslan/library-management-service/domain"
	fiberInfra "github.com/muhammedkucukaslan/library-management-service/infra/fiber"
	"github.com/muhammedkucukaslan/library-management-service/infra/jwt"
	"github.com/muhammedkucukaslan/library-management-service/infra/postgres"
	resendInfra "github.com/muhammedkucukaslan/library-management-service/infra/resend"
	"github.com/robfig/cron/v3"
)

type Request any
type Response any

type HandlerInterface[R Request, Res Response] interface {
	Handle(ctx context.Context, req *R) (*Res, error)
}

func handle[R Request, Res Response](handler HandlerInterface[R, Res]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req R
		if err := c.BodyParser(&req); err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if err := c.ParamsParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if err := c.QueryParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if err := c.ReqHeaderParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if c.Locals("requireAuth") == true {
			role, ok := c.Locals("role").(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "invalid role in context",
				})
			}

			userID, ok := c.Locals("userID").(int)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "invalid user_id in context",
				})
			}

			ctx := context.WithValue(c.UserContext(), domain.RoleKey, role)
			ctx = context.WithValue(ctx, domain.UserIDKey, userID)
			c.SetUserContext(ctx)
		}

		res, err := handler.Handle(c.UserContext(), &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(res)
	}
}

func authMiddleware(c *fiber.Ctx) error {

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing authorization header",
		})
	}

	tokenString := token[7:]
	tokenService := jwt.NewTokenService(os.Getenv("JWT_SECRET_KEY"), time.Hour*24)

	payload, err := tokenService.ValidateToken(tokenString)

	if err != nil {
		fmt.Println("ERROR", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}
	c.Locals("userID", payload.UserID)
	c.Locals("role", payload.Role)

	return c.Next()
}

func main() {

	godotenv.Load()

	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Concurrency:  256 * 1024,
	})

	repo := postgres.NewRepository()
	defer repo.Close()

	tokenService := jwt.NewTokenService(os.Getenv("JWT_SECRET_KEY"), time.Hour*24)
	cookieService := fiberInfra.NewFiberCookieService()
	resendEmailService := resendInfra.NewResendEmailService(os.Getenv("RESEND_API_KEY"))

	c := cron.New()

	punishmentCronJob := punishment.NewPunishLoanCronJob(repo, resendEmailService)
	c.AddFunc("@every 5m", punishmentCronJob.Execute)
	c.Start()

	healthcheckHandler := healthcheck.NewHealthcheckHandler()

	signupHandler := auth.NewSignupHandler(repo, tokenService, cookieService, resendEmailService)
	loginHandler := auth.NewLoginHandler(repo, tokenService, cookieService)

	getUserHandler := user.NewGetUserHandler(repo)
	deleteUserHandler := user.NewDeleteUserHandler(repo)
	updateUserHandler := user.NewUpdateUserHandler(repo)

	createAuthorHandler := author.NewCreateAuthorHandler(repo)
	deleteAuthorHandler := author.NewDeleteAuthorHandler(repo)

	getBookHandler := book.NewGetBookHandler(repo)
	createBookHandler := book.NewCreateBookHandler(repo)
	deleteBookHandler := book.NewDeleteBookHandler(repo)

	getBookCategoriesHandler := bookCategory.NewGetBookCategoriesHandler(repo)
	createBookCategoryHandler := bookCategory.NewCreateCategoryHandler(repo)
	deleteBookCategoryHandler := bookCategory.NewDeleteCategoryHandler(repo)

	borrowBookHandler := loan.NewBorrowBookHandler(repo)
	returnBookHandler := loan.NewReturnBookHandler(repo)
	getLoansHandler := loan.NewGetLoanHandler(repo)

	punishUserHandler := punishment.NewPunishUserHandler(repo)

	app.Get("/healthcheck", handle[healthcheck.HealthcheckRequest, healthcheck.HealthcheckResponse](healthcheckHandler))
	app.Use(fiberInfra.ContextMiddleware)

	api := app.Group("/api")

	api.Post("/signup", handle[auth.SignupRequest, auth.SignupResponse](signupHandler))
	api.Post("/login", handle[auth.LoginRequest, auth.LoginResponse](loginHandler))

	api.Use(func(c *fiber.Ctx) error {
		c.Locals("requireAuth", true)
		return c.Next()
	}, authMiddleware)

	usersApp := api.Group("/users")
	usersApp.Get("/:id", handle[user.GetUserRequest, user.GetUserResponse](getUserHandler))
	usersApp.Delete("/:id", handle[user.DeleteUserRequest, user.DeleteUserResponse](deleteUserHandler))
	usersApp.Put("/:id", handle[user.UpdateUserRequest, user.UpdateUserResponse](updateUserHandler))
	usersApp.Post("/:id/punish", handle[punishment.PunishUserRequest, punishment.PunishUserResponse](punishUserHandler))

	authorsApp := api.Group("/authors")
	authorsApp.Post("/", handle[author.CreateAuthorRequest, author.CreateAuthorResponse](createAuthorHandler))
	authorsApp.Delete("/:id", handle[author.DeleteAuthorRequest, author.DeleteAuthorResponse](deleteAuthorHandler))

	booksApp := api.Group("/books")
	booksApp.Post("/", handle[book.CreateBookRequest, book.CreateBookResponse](createBookHandler))
	booksApp.Get("/:id", handle[book.GetBookRequest, book.GetBookResponse](getBookHandler))
	booksApp.Delete("/:id", handle[book.DeleteBookRequest, book.DeleteBookResponse](deleteBookHandler))

	bookCategoriesApp := api.Group("/categories")
	bookCategoriesApp.Post("/", handle[bookCategory.CreateCategoryRequest, bookCategory.CreateCategoryResponse](createBookCategoryHandler))
	bookCategoriesApp.Delete("/:id", handle[bookCategory.DeleteCategoryRequest, bookCategory.DeleteCategoryResponse](deleteBookCategoryHandler))
	bookCategoriesApp.Get("/", handle[bookCategory.GetBookCategoriesRequest, bookCategory.GetBookCategoriesResponse](getBookCategoriesHandler))

	loanApp := api.Group("/loans")
	loanApp.Post("/", handle[loan.BorrowBookRequest, loan.BorrowBookResponse](borrowBookHandler))
	loanApp.Get("/", handle[loan.GetLoansRequest, loan.GetLoansResponse](getLoansHandler))
	loanApp.Post("/:id/return", handle[loan.ReturnBookRequest, loan.ReturnBookResponse](returnBookHandler))

	go func() {
		if err := app.Listen(":3000"); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}()
	gracefulShutdown(app)
}

func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	app.Shutdown()
}
