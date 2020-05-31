package store

import (
	"database/sql"

	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/model"
	_ "github.com/go-sql-driver/mysql"
)

type UserStore struct {
	db  *sql.DB
	cfg *config.Config
}

func NewUserStore(db *sql.DB, cfg *config.Config) *UserStore {
	return &UserStore{
		db:  db,
		cfg: cfg,
	}
}

func (us *UserStore) GetByID(id int64) (*model.User, error) {
	stmt, err := us.db.Prepare(`
		SELECT 
			*
		FROM user
		WHERE id = ?
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
		//		&user.Password,
		&user.Session,
		//		&user.SessionDate,
		&user.Manufacture,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserStore) UserLogin(username string, password string) (*model.User, error) {
	q := `
	SELECT id, username, session, manufacture
	FROM user 
	WHERE username = ? AND password = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(username, password)
	user, err := us.CreateUserFromRow(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// func (us *UserStore) DeleteByID(id int) (bool, error) {
// 	stmt, err := us.db.Prepare(`
// 		DELETE
// 		FROM user
// 		WHERE id = ?
// 	`)
// 	if err != nil {
// 		return false, err
// 	}

// 	defer stmt.Close()
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// create a new User
// func (us *UserStore) CreateUser(username string) (string, error) {
// 	stmt, err := us.db.Prepare(`
//  		INSERT INTO
//  			user (
//  				username
//  			)
//  			VALUES (
//  				"?"
//  			)

//  	`)
// 	if err != nil {
// 		return "Error:", err
// 	}

// 	defer stmt.Close()
// 	if err != nil {
// 		return "Error:", err
// 	}
// }

// 	return "User Created", nil
// }
