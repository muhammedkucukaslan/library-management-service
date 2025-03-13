package postgres

import (
	"context"

	"github.com/muhammedkucukaslan/library-management-service/domain"
)

func (r *Repository) PunishUser(ctx context.Context, punishment *domain.Punishment) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO punishments (user_id, punisher_id, reason, end_date) VALUES ($1, $2, $3, $4)", punishment.UserID, punishment.PunisherID, punishment.Reason, punishment.EndDate)
	if err != nil {
		return err
	}
	return nil
}
