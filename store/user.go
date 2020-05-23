package store

import (
	"database/sql"
	"github.com/calenaur/pandemic/model"
	"github.com/calenaur/pandemic/config"
	_ "github.com/go-sql-driver/mysql"
)

type UserStore struct {
	db *sql.DB
	cfg *config.Config
}

func NewUserStore(db *sql.DB, cfg *config.Config) *UserStore {
	return &UserStore{
		db: db,
		cfg: cfg,
	}
}

func (us *UserStore) GetByID(id int64) (*model.User, error) {
	stmt, err := us.db.Prepare(`
		SELECT 
			u.id, u.username
		FROM user u
		WHERE u.id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	user, err := us.CreateUserFromRow(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserStore) CreateUserFromRow(row *sql.Row) (*model.User, error) {
	user := &model.User{}
	err := row.Scan(
		&user.ID, 
		&user.Username, 
	)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}