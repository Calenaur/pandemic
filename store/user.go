package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/model"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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
		// &user.Password,
		&user.Balance,
		// &user.SessionDate,
		&user.Manufacture,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
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
		log.Fatal(err)
		fmt.Println(err)
		return nil, err
	}

	return user, nil
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

func (us *UserStore) decipher(username string, passwordString string) error {
	// decipher hashed password
	var hashedPasswordString string
	rows, err := us.db.Query("SELECT password FROM user WHERE username = ?", username)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&hashedPasswordString)
		if err != nil {
			return err
		}
	}
	err = rows.Err()
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

func (us *UserStore) ChangeUserName(id string, newName string) error {

	// Query
	q := `
	UPDATE user 
	SET  username = ? 
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newName, id)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) ChangeUserPassword(id string, password string) error {

	// Hash password
	passwordByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Query
	q := `
	UPDATE user 
	SET  password = ? 
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(hashedPassword, id)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) DeleteAccount(id string) error {

	// Query
	q := `
	DELETE
	FROM user 
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return err
}
