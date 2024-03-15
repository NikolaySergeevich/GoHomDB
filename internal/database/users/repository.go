package users

import (
	"context"
	"fmt"

	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
)

/*
Создаёт новый репозиторий()
Принимает созданное в env/setup.go подключение и время ожидания
*/
func New(userDB *pgx.Conn, timeout time.Duration) *Repository {
	return &Repository{userDB: userDB, timeout: timeout}
}

/*
База данных конкретного пользователя какого-то
*/
type Repository struct {
	userDB  *pgx.Conn
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateUserReq) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	r.userDB.QueryRow(context.Background(), "INSERT INTO users (id, username, password) VALUES($1, $2, $3) returning username, password, created_at, updated_at",
		req.ID, req.Username, req.Password).Scan(&u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	u.ID = req.ID

	return u, nil
}

func (r *Repository) FindByID(ctx context.Context, userID uuid.UUID) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	row := r.userDB.QueryRow(context.Background(), "select * from users where id=$1", userID)
	
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, fmt.Errorf("row scan ByID: %w", err)
	}

	return u, nil
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	
	row := r.userDB.QueryRow(context.Background(), "select * from users where username = $1", username)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, fmt.Errorf("row scan ByUsername: %w", err)
	}
	return u, nil
}
