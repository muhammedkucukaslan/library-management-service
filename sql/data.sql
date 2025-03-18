-- Users
INSERT INTO users (
        first_name,
        second_name,
        role,
        password,
        email,
        address,
        birth_date,
        created_at,
        updated_at
    )
VALUES (
        'Emma',
        'Johnson',
        'user',
        '$2a$10$aYFLlRsr9/fXttif9ozSr.k2wG92ZyGGU38uuZ1ms8lzzYtZOba/6',
        'emma.johnson@example.com',
        '456 Oak St, Springfield, ST 54321',
        '1992-03-15',
        NOW (),
        NOW ()
    ),
    (
        'Michael',
        'Williams',
        'user',
        '$2a$10$QhxozZSydAi8FU7pgCyEveDC/4Pq0GoB.K0BeTYCq/tzePO1YsveW',
        'michael.williams@example.com',
        '789 Pine Ave, Riverdale, RD 67890',
        '1985-07-22',
        NOW (),
        NOW ()
    ),
    (
        'Sophia',
        'Brown',
        'user',
        '$2a$10$mNG7/MI5jIpuLHQVAHPcAOlsPvdHWZ3ynVMdXW9TbMWnI/HxizWcu',
        'sophia.brown@example.com',
        '101 Cedar Ln, Lakeside, LS 13579',
        '1993-11-08',
        NOW (),
        NOW ()
    ),
    (
        'Daniel',
        'Jones',
        'user',
        '$2a$10$VnACnraH/VyswnxdBd10GOvuPjiWTsHIUz/VRrvwsuPIcKhZAE98K',
        'daniel.jones@example.com',
        '202 Maple Dr, Hillcrest, HC 24680',
        '1988-05-19',
        NOW (),
        NOW ()
    ),
    (
        'Olivia',
        'Garcia',
        'user',
        '$2a$10$aYFLlRsr9/fXttif9ozSr.k2wG92ZyGGU38uuZ1ms8lzzYtZOba/6',
        'olivia.garcia@example.com',
        '303 Elm St, Westfield, WF 97531',
        '1991-09-27',
        NOW (),
        NOW ()
    ),
    (
        'William',
        'Miller',
        'admin',
        '$2a$10$QhxozZSydAi8FU7pgCyEveDC/4Pq0GoB.K0BeTYCq/tzePO1YsveW',
        'william.miller@example.com',
        '404 Birch Rd, Eastview, EV 86420',
        '1983-12-04',
        NOW (),
        NOW ()
    ),
    (
        'Ava',
        'Davis',
        'user',
        '$2a$10$mNG7/MI5jIpuLHQVAHPcAOlsPvdHWZ3ynVMdXW9TbMWnI/HxizWcu',
        'ava.davis@example.com',
        '505 Walnut Blvd, Northtown, NT 75319',
        '1995-02-11',
        NOW (),
        NOW ()
    ),
    (
        'James',
        'Rodriguez',
        'user',
        '$2a$10$VnACnraH/VyswnxdBd10GOvuPjiWTsHIUz/VRrvwsuPIcKhZAE98K',
        'james.rodriguez@example.com',
        '606 Spruce Ct, Southville, SV 42086',
        '1987-06-30',
        NOW (),
        NOW ()
    ),
    (
        'Isabella',
        'Martinez',
        'user',
        '$2a$10$aYFLlRsr9/fXttif9ozSr.k2wG92ZyGGU38uuZ1ms8lzzYtZOba/6',
        'isabella.martinez@example.com',
        '707 Pineapple Way, Beachside, BS 31795',
        '1994-10-17',
        NOW (),
        NOW ()
    ),
    (
        'Alexander',
        'Hernandez',
        'user',
        '$2a$10$QhxozZSydAi8FU7pgCyEveDC/4Pq0GoB.K0BeTYCq/tzePO1YsveW',
        'alexander.hernandez@example.com',
        '808 Orange Ave, Mountainview, MV 86421',
        '1989-04-23',
        NOW (),
        NOW ()
    ),
    (
        'Charlotte',
        'Lopez',
        'user',
        '$2a$10$mNG7/MI5jIpuLHQVAHPcAOlsPvdHWZ3ynVMdXW9TbMWnI/HxizWcu',
        'charlotte.lopez@example.com',
        '909 Grape St, Valleytown, VT 97532',
        '1996-08-09',
        NOW (),
        NOW ()
    ),
    (
        'Benjamin',
        'Gonzalez',
        'admin',
        '$2a$10$VnACnraH/VyswnxdBd10GOvuPjiWTsHIUz/VRrvwsuPIcKhZAE98K',
        'benjamin.gonzalez@example.com',
        '121 Cherry Ln, Meadowland, ML 64208',
        '1984-01-26',
        NOW (),
        NOW ()
    ),
    (
        'Mia',
        'Wilson',
        'user',
        '$2a$10$aYFLlRsr9/fXttif9ozSr.k2wG92ZyGGU38uuZ1ms8lzzYtZOba/6',
        'mia.wilson@example.com',
        '232 Peach Blvd, Riverside, RS 25874',
        '1993-05-03',
        NOW (),
        NOW ()
    ),
    (
        'Ethan',
        'Anderson',
        'user',
        '$2a$10$QhxozZSydAi8FU7pgCyEveDC/4Pq0GoB.K0BeTYCq/tzePO1YsveW',
        'ethan.anderson@example.com',
        '343 Plum Dr, Oceanview, OV 36985',
        '1986-11-29',
        NOW (),
        NOW ()
    ),
    (
        'Amelia',
        'Thomas',
        'user',
        '$2a$10$mNG7/MI5jIpuLHQVAHPcAOlsPvdHWZ3ynVMdXW9TbMWnI/HxizWcu',
        'amelia.thomas@example.com',
        '454 Lemon Ct, Foresthill, FH 14785',
        '1997-07-16',
        NOW (),
        NOW ()
    ),
    (
        'Henry',
        'Taylor',
        'user',
        '$2a$10$VnACnraH/VyswnxdBd10GOvuPjiWTsHIUz/VRrvwsuPIcKhZAE98K',
        'henry.taylor@example.com',
        '565 Apple Way, Downtown, DT 69874',
        '1990-03-21',
        NOW (),
        NOW ()
    ),
    (
        'Luna',
        'Moore',
        'user',
        '$2a$10$aYFLlRsr9/fXttif9ozSr.k2wG92ZyGGU38uuZ1ms8lzzYtZOba/6',
        'luna.moore@example.com',
        '676 Banana Ave, Uptown, UT 25896',
        '1994-12-07',
        NOW (),
        NOW ()
    ),
    (
        'Sebastian',
        'Jackson',
        'admin',
        '$2a$10$QhxozZSydAi8FU7pgCyEveDC/4Pq0GoB.K0BeTYCq/tzePO1YsveW',
        'sebastian.jackson@example.com',
        '787 Coconut St, Midtown, MT 36987',
        '1988-06-14',
        NOW (),
        NOW ()
    ) -- Authors
INSERT INTO authors (first_name, second_name, bio, birth_date)
VALUES (
        'James',
        'Patterson',
        'James Patterson is an American author known for his thriller and mystery novels.',
        '1947-03-22'
    ),
    (
        'Margaret',
        'Atwood',
        'Margaret Atwood is a Canadian author known for her dystopian fiction and poetry.',
        '1939-11-18'
    ),
    (
        'Haruki',
        'Murakami',
        'Haruki Murakami is a Japanese author known for his surrealist fiction.',
        '1949-01-12'
    ),
    (
        'Chimamanda',
        'Adichie',
        'Chimamanda Adichie is a Nigerian author known for her novels and feminist writings.',
        '1977-09-15'
    ),
    (
        'Stephen',
        'King',
        'Stephen King is an American author known for his horror and supernatural fiction.',
        '1947-09-21'
    ),
    (
        'Zadie',
        'Smith',
        'Zadie Smith is a British novelist and essayist known for her multicultural narratives.',
        '1975-10-25'
    ),
    (
        'Gabriel',
        'García Márquez',
        'Gabriel García Márquez was a Colombian author known for magical realism.',
        '1927-03-06'
    ),
    (
        'Toni',
        'Morrison',
        'Toni Morrison was an American novelist and professor known for her explorations of Black identity.',
        '1931-02-18'
    ),
    (
        'Salman',
        'Rushdie',
        'Salman Rushdie is a British-Indian novelist known for his blend of magical realism and historical fiction.',
        '1947-06-19'
    ),
    (
        'Elena',
        'Ferrante',
        'Elena Ferrante is an Italian novelist known for her Neapolitan Novels series.',
        '1943-04-05'
    ),
    (
        'Yuval Noah',
        'Harari',
        'Yuval Noah Harari is an Israeli historian and philosopher known for his books on human evolution and history.',
        '1976-02-24'
    ),
    (
        'Khaled',
        'Hosseini',
        'Khaled Hosseini is an Afghan-American novelist known for his emotional storytelling.',
        '1965-03-04'
    ),
    (
        'Neil',
        'Gaiman',
        'Neil Gaiman is a British author known for his fantasy and comic book writing.',
        '1960-11-10'
    ),
    (
        'Isabel',
        'Allende',
        'Isabel Allende is a Chilean writer known for her magical realist novels.',
        '1942-08-02'
    ),
    (
        'Tana',
        'French',
        'Tana French is an Irish-American author known for her psychological crime novels.',
        '1973-05-10'
    ),
    (
        'Orhan',
        'Pamuk',
        'Orhan Pamuk is a Turkish novelist and Nobel Prize winner known for his works about Turkish identity.',
        '1952-06-07'
    ),
    (
        'Min Jin',
        'Lee',
        'Min Jin Lee is a Korean-American author known for her historical fiction.',
        '1968-11-11'
    ),
    (
        'Colson',
        'Whitehead',
        'Colson Whitehead is an American novelist known for his explorations of American history.',
        '1969-11-06'
    ),
    (
        'Bernardine',
        'Evaristo',
        'Bernardine Evaristo is a British author known for her experimental fiction and poetry.',
        '1959-05-28'
    ),
    (
        'Richard',
        'Powers',
        'Richard Powers is an American novelist known for his works exploring science and technology.',
        '1957-06-18'
    ) -- Categories
INSERT INTO categories (title)
values ('Mahtematical Analysis'),
    ('Physics'),
    ('Physics'),
    ('Chemistry'),
    ('Biology'),
    ('History'),
    ('Geography'),
    ('Literature'),
    ('Philosophy'),
    ('Economics'),
    ('Sociology'),
    ('Psychology'),
    ('Political Science'),
    ('Law'),
    ('Computer Science'),
    ('Engineering'),
    ('Medicine'),
    ('Agriculture'),
    ('Business'),
    ('Management'),
    ('Marketing'),
    ('Finance'),
    ('Accounting') -- Books
INSERT INTO books (
        id,
        title,
        isbn,
        author_id,
        publisher,
        edition,
        description,
        created_at,
        updated_at
    )
VALUES (
        48,
        'Beyond Limits',
        '978-0-553-34567-8',
        5,
        'Simon & Schuster',
        1,
        'Michael Jordan shares his philosophy on perseverance and excellence.',
        NOW (),
        NOW ()
    ),
    (
        49,
        'Whispers in the Wind',
        '978-0-062-45678-9',
        6,
        'Graywolf Press',
        3,
        'A collection of evocative poems exploring themes of identity and belonging.',
        NOW (),
        NOW ()
    ),
    (
        50,
        'The Forgotten Path',
        '978-0-307-56789-0',
        7,
        'Oxford University Press',
        1,
        'A coming-of-age novel set against the backdrop of political upheaval.',
        NOW (),
        NOW ()
    ),
    (
        51,
        'Shadows of the Past',
        '978-0-679-67890-1',
        8,
        'Vintage Books',
        2,
        'A psychological thriller about a woman confronting her troubled childhood.',
        NOW (),
        NOW ()
    ),
    (
        52,
        'The Quantum Paradox',
        '978-0-812-78901-2',
        9,
        'MIT Press',
        1,
        'An exploration of quantum physics and its philosophical implications.',
        NOW (),
        NOW ()
    ),
    (
        53,
        'Lost in Translation',
        '978-0-375-89012-3',
        10,
        'Columbia University Press',
        3,
        'A novel about cultural identity and the immigrant experience in America.',
        NOW (),
        NOW ()
    ),
    (
        54,
        'The Art of Silence',
        '978-0-226-90123-4',
        11,
        'Yale University Press',
        1,
        'A meditation on the power of silence in an increasingly noisy world.',
        NOW (),
        NOW ()
    ),
    (
        55,
        'Desert Bloom',
        '978-0-394-01234-5',
        12,
        'Beacon Press',
        2,
        'A novel about resilience and renewal set in the American Southwest.',
        NOW (),
        NOW ()
    ),
    (
        56,
        'The Last Emperor',
        '978-0-452-12345-6',
        13,
        'Knopf',
        1,
        'A historical account of the fall of the last imperial dynasty.',
        NOW (),
        NOW ()
    ),
    (
        57,
        'Oceans Depth',
        '978-0-571-23456-7',
        14,
        'Faber & Faber',
        2,
        'A maritime adventure novel set in the early 19th century.',
        NOW (),
        NOW ()
    ),
    (
        58,
        'The Infinite Mind',
        '978-0-684-34567-8',
        15,
        'Princeton University Press',
        1,
        'A groundbreaking work on consciousness and cognitive science.',
        NOW (),
        NOW ()
    ),
    (
        59,
        'City of Dreams',
        '978-0-805-45678-9',
        16,
        'Metropolitan Books',
        3,
        'A sprawling urban novel exploring the diverse lives of city dwellers.',
        NOW (),
        NOW ()
    ),
    (
        60,
        'The Historians Daughter',
        '978-0-316-56789-0',
        17,
        'Little, Brown and Company',
        1,
        'A family saga spanning generations of women who preserve cultural histories.',
        NOW (),
        NOW ()
    ),
    (
        61,
        'Frost & Fire',
        '978-0-399-67890-1',
        18,
        'Riverhead Books',
        2,
        'A fantasy novel set in a world divided between eternal winter and summer.',
        NOW (),
        NOW ()
    ),
    (
        62,
        'The Economics of Everything',
        '978-0-670-78901-2',
        19,
        'Harvard University Press',
        1,
        'A provocative analysis of how economic principles shape everyday decisions.',
        NOW (),
        NOW ()
    ),
    (
        63,
        'Whispers in the Dark',
        '978-0-374-89012-3',
        20,
        'Farrar, Straus and Giroux',
        4,
        'A collection of gothic short stories exploring the supernatural.',
        NOW (),
        NOW ()
    ),
    (
        64,
        'The Language of Birds',
        '978-0-811-90123-4',
        21,
        'New Directions',
        1,
        'A poetic novel about a woman who discovers she can communicate with birds.',
        NOW (),
        NOW ()
    ),
    (
        65,
        'Tomorrows Promise',
        '978-0-312-01234-5',
        22,
        'St. Martins Press',
        2,
        'A science fiction novel about a utopian future society facing unexpected challenges.',
        NOW (),
        NOW ()
    ) -- Book Categories
INSERT INTO book_categories (book_id, category_id)
VALUES (48, 9),
    (48, 22),
    (49, 4),
    (49, 18),
    (50, 7),
    (50, 11),
    (51, 5),
    (51, 21),
    (52, 10),
    (52, 25),
    (53, 3),
    (53, 17),
    (54, 8),
    (54, 14),
    (55, 6),
    (55, 19),
    (56, 2),
    (56, 23),
    (57, 1),
    (57, 16),
    (58, 10),
    (58, 24),
    (59, 3),
    (59, 13),
    (60, 2),
    (60, 20),
    (61, 1),
    (61, 26),
    (62, 9),
    (62, 27),
    (63, 5),
    (63, 15),
    (64, 4),
    (64, 22),
    (65, 1),
    (65, 10) -- Book Stocks
INSERT INTO book_stocks (
        book_id,
        total_quantity,
        available_quantity,
        created_at,
        updated_at
    )
VALUES (48, 150, 147, NOW (), NOW ()),
    (49, 100, 98, NOW (), NOW ()),
    (50, 130, 125, NOW (), NOW ()),
    (51, 110, 110, NOW (), NOW ()),
    (52, 90, 88, NOW (), NOW ()),
    (53, 120, 120, NOW (), NOW ()),
    (54, 80, 80, NOW (), NOW ()),
    (55, 200, 194, NOW (), NOW ()),
    (56, 75, 75, NOW (), NOW ()),
    (57, 120, 120, NOW (), NOW ()),
    (58, 100, 100, NOW (), NOW ()),
    (59, 160, 158, NOW (), NOW ()),
    (60, 120, 116, NOW (), NOW ()),
    (61, 150, 150, NOW (), NOW ()),
    (62, 90, 90, NOW (), NOW ()),
    (63, 110, 110, NOW (), NOW ()),
    (64, 100, 97, NOW (), NOW ()),
    (65, 120, 120, NOW (), NOW ()) -- Loans
INSERT INTO loans (
        id,
        user_id,
        book_id,
        started_at,
        due_at,
        returned_at,
        status
    )
VALUES (
        2,
        4,
        48,
        '2025-02-20 10:15:22.123456+00',
        '2025-03-06 10:15:22.123456+00',
        '2025-03-01 14:22:37.456789+00',
        'RETURNED'
    ),
    (
        3,
        10,
        49,
        '2025-02-25 14:30:45.234567+00',
        '2025-03-11 14:30:45.234567+00',
        '2025-03-10 09:45:12.567890+00',
        'RETURNED'
    ),
    (
        4,
        22,
        50,
        '2025-03-01 09:45:18.345678+00',
        '2025-03-15 09:45:18.345678+00',
        '2025-03-12 16:18:42.678901+00',
        'RETURNED'
    ),
    (
        5,
        23,
        52,
        '2025-03-02 16:20:33.456789+00',
        '2025-03-16 16:20:33.456789+00',
        NULL,
        'BORROWED'
    ),
    (
        6,
        4,
        55,
        '2025-03-03 11:10:56.567890+00',
        '2025-03-17 11:10:56.567890+00',
        NULL,
        'BORROWED'
    ),
    (
        7,
        10,
        59,
        '2025-03-04 13:25:22.678901+00',
        '2025-03-18 13:25:22.678901+00',
        NULL,
        'BORROWED'
    ),
    (
        8,
        22,
        60,
        '2025-03-05 15:40:38.789012+00',
        '2025-03-19 15:40:38.789012+00',
        '2025-03-13 10:33:27.890123+00',
        'RETURNED'
    ),
    (
        9,
        23,
        64,
        '2025-03-06 10:15:44.890123+00',
        '2025-03-20 10:15:44.890123+00',
        NULL,
        'BORROWED'
    ),
    (
        10,
        4,
        49,
        '2025-03-07 14:30:19.901234+00',
        '2025-03-21 14:30:19.901234+00',
        NULL,
        'BORROWED'
    ),
    (
        11,
        10,
        50,
        '2025-03-08 09:45:25.012345+00',
        '2025-03-22 09:45:25.012345+00',
        '2025-03-12 11:27:36.123456+00',
        'RETURNED'
    ),
    (
        12,
        22,
        55,
        '2025-03-08 16:20:41.123456+00',
        '2025-03-22 16:20:41.123456+00',
        NULL,
        'BORROWED'
    ),
    (
        13,
        23,
        59,
        '2025-03-09 11:10:53.234567+00',
        '2025-03-23 11:10:53.234567+00',
        NULL,
        'BORROWED'
    ),
    (
        14,
        4,
        60,
        '2025-03-10 13:25:14.345678+00',
        '2025-03-24 13:25:14.345678+00',
        '2025-03-12 15:42:18.456789+00',
        'RETURNED'
    ),
    (
        15,
        10,
        64,
        '2025-03-10 15:40:33.456789+00',
        '2025-03-24 15:40:33.456789+00',
        NULL,
        'BORROWED'
    ),
    (
        16,
        22,
        48,
        '2025-03-11 10:15:52.567890+00',
        '2025-03-25 10:15:52.567890+00',
        NULL,
        'BORROWED'
    ),
    (
        17,
        23,
        52,
        '2025-03-11 14:30:28.678901+00',
        '2025-03-25 14:30:28.678901+00',
        '2025-03-13 09:18:42.789012+00',
        'RETURNED'
    ),
    (
        18,
        4,
        55,
        '2025-03-12 09:45:37.789012+00',
        '2025-03-26 09:45:37.789012+00',
        NULL,
        'BORROWED'
    ),
    (
        19,
        10,
        59,
        '2025-03-12 16:20:48.890123+00',
        '2025-03-26 16:20:48.890123+00',
        NULL,
        'BORROWED'
    ),
    (
        20,
        22,
        64,
        '2025-03-13 11:10:59.901234+00',
        '2025-03-27 11:10:59.901234+00',
        NULL,
        'BORROWED'
    ),
    (
        21,
        23,
        48,
        '2025-03-13 13:25:11.012345+00',
        '2025-03-27 13:25:11.012345+00',
        NULL,
        'BORROWED'
    ) -- Punishments
INSERT INTO punishments (
        id,
        user_id,
        punisher_id,
        reason,
        end_at,
        created_at
    )
VALUES (
        6,
        10,
        23,
        'Returned book with damaged pages',
        '2025-04-15 00:00:00+00',
        '2025-03-13 14:15:22.123456+00'
    ),
    (
        7,
        4,
        23,
        'Multiple late returns',
        '2025-04-01 00:00:00+00',
        '2025-03-13 15:20:33.234567+00'
    ),
    (
        8,
        22,
        23,
        'Disruptive behavior in reading area',
        '2025-03-20 00:00:00+00',
        '2025-03-13 16:30:44.345678+00'
    ),
    (
        9,
        10,
        22,
        'Eating in the library',
        '2025-03-25 00:00:00+00',
        '2025-03-13 17:40:55.456789+00'
    ),
    (
        10,
        4,
        23,
        'Using phone in silent zone',
        '2025-03-18 00:00:00+00',
        '2025-03-14 09:10:11.567890+00'
    ),
    (
        11,
        22,
        23,
        'Arguing with library staff',
        '2025-04-05 00:00:00+00',
        '2025-03-14 10:20:22.678901+00'
    ),
    (
        12,
        10,
        22,
        'Repeatedly missing appointments',
        '2025-03-22 00:00:00+00',
        '2025-03-14 11:30:33.789012+00'
    ),
    (
        13,
        4,
        23,
        'Talking loudly in study area',
        '2025-03-28 00:00:00+00',
        '2025-03-14 12:40:44.890123+00'
    ),
    (
        14,
        22,
        23,
        'Unauthorized access to restricted section',
        '2025-04-10 00:00:00+00',
        '2025-03-14 13:50:55.901234+00'
    ),
    (
        15,
        10,
        22,
        'Failed to pay overdue fines',
        '2025-03-31 00:00:00+00',
        '2025-03-14 14:00:00.012345+00'
    ),
    (
        16,
        4,
        23,
        'Bringing unauthorized guests',
        '2025-03-21 00:00:00+00',
        '2025-03-14 15:10:11.123456+00'
    ),
    (
        17,
        22,
        23,
        'Misuse of library computers',
        '2025-04-03 00:00:00+00',
        '2025-03-14 16:20:22.234567+00'
    ),
    (
        18,
        10,
        22,
        'Disturbing other patrons',
        '2025-03-27 00:00:00+00',
        '2025-03-14 17:30:33.345678+00'
    ),
    (
        19,
        4,
        23,
        'Falling asleep repeatedly in reading room',
        '2025-03-19 00:00:00+00',
        '2025-03-15 09:40:44.456789+00'
    ),
    (
        20,
        22,
        23,
        'Misplacing reference materials',
        '2025-04-07 00:00:00+00',
        '2025-03-15 10:50:55.567890+00'
    ),
    (
        21,
        10,
        22,
        'Using highlighter on library book',
        '2025-03-24 00:00:00+00',
        '2025-03-15 11:00:00.678901+00'
    ),
    (
        22,
        4,
        23,
        'Loud conversations on phone',
        '2025-03-29 00:00:00+00',
        '2025-03-15 12:10:11.789012+00'
    ),
    (
        23,
        22,
        23,
        'Exceeding maximum borrowing limit',
        '2025-04-12 00:00:00+00',
        '2025-03-15 13:20:22.890123+00'
    ),
    (
        24,
        10,
        22,
        'Violation of library code of conduct',
        '2025-04-02 00:00:00+00',
        '2025-03-15 14:30:33.901234+00'
    ),
    (
        25,
        4,
        23,
        'Inappropriate use of library resources',
        '2025-03-23 00:00:00+00',
        '2025-03-15 15:40:44.012345+00'
    )
SELECT setval (
        pg_get_serial_sequence ('users', 'id'),
        (
            SELECT MAX(id)
            FROM users
        )
    ),
    setval (
        pg_get_serial_sequence ('books', 'id'),
        (
            SELECT MAX(id)
            FROM books
        )
    ),
    setval (
        pg_get_serial_sequence ('book_stocks', 'id'),
        (
            SELECT MAX(id)
            FROM book_stocks
        )
    ),
    setval (
        pg_get_serial_sequence ('categories', 'id'),
        (
            SELECT MAX(id)
            FROM categories
        )
    ),
    setval (
        pg_get_serial_sequence ('book_categories', 'id'),
        (
            SELECT MAX(id)
            FROM book_categories
        )
    ),
    setval (
        pg_get_serial_sequence ('loans', 'id'),
        (
            SELECT MAX(id)
            FROM loans
        )
    ),
    setval (
        pg_get_serial_sequence ('authors', 'id'),
        (
            SELECT MAX(id)
            FROM authors
        )
    ),
    setval (
        pg_get_serial_sequence ('punishments', 'id'),
        (
            SELECT MAX(id)
            FROM punishments
        )
    );