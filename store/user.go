package store

import (
	"database/sql"

	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/model"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	db  *sql.DB
	cfg *config.Config
}

type Message struct {
	Error *Error `json:"error"`
}

type Error struct {
	code    string `json:"code"`
	message string `json:"message"`
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
		id, username, balance, manufacture
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

	return user, err
}

func (us *UserStore) CreateUserFromRow(row *sql.Row) (*model.User, error) {
	user := &model.User{}
	err := row.Scan(
		&user.ID,
		&user.Username,
		// &user.Password,
		&user.Balance,
		// &user.SessionDate,
		&user.Manufacture,
	)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (us *UserStore) UserLogin(username string, password string) (*model.User, error) {
	// decipher password
	err := us.decipher(username, password)
	if err != nil {
		return nil, err
	}
	q := `
	SELECT id, username, balance, manufacture
	FROM user 
	WHERE username = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(username)
	user, err := us.CreateUserFromRow(row)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (us *UserStore) UserSignup(username string, passwordString string) error {

	// Hash password
	password := []byte(passwordString)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Query
	q := `
	INSERT INTO user (username, password) 
	VALUES (? , ?)
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(username, hashedPassword)
	if err != nil {
		return err
	}

	return err
}

// decipher hashed password
func (us *UserStore) decipher(username string, passwordString string) error {
	var hashedPasswordString string
	row := us.db.QueryRow("SELECT password FROM user WHERE username = ?", username)

	err := row.Scan(&hashedPasswordString)

	if err != nil {
		return err
	}

	hashedPassword := []byte(hashedPasswordString)
	password := []byte(passwordString)

	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}
	return err
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
