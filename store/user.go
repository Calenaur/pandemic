package store

import (
	"database/sql"
	"strings"

	"github.com/Calenaur/pandemic/model"
	"github.com/calenaur/pandemic/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

func (us *UserStore) GetByID(id string) (*model.User, error) {
	stmt, err := us.db.Prepare(`
		SELECT 
		id, username, accesslevel, tier, balance
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
		&user.AccessLevel,
		&user.Tier,
		&user.Balance,

		// &user.SessionDate,
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
	SELECT id, username, accesslevel,tier, balance
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
	INSERT INTO user (id, username, password) 
	VALUES (?, ?, ?)
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()
	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	// fmt.Println(id)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, username, hashedPassword)
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

func (us *UserStore) GetUserDetails(id string) (*model.User, error) {

	// Query
	q := `
	SELECT username, accesslevel, tier, balance
	FROM user
	WHERE id = ?`

	stmt, err := us.db.Prepare(q)
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

func (us *UserStore) UpdateBalance(id string, balance string) error {
	// Query
	q := `
	UPDATE
	user 
	SET balance = ?
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance, id)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) UpdateDevice(id string, device string) error {
	// Query
	q := `
	UPDATE
	user 
	SET device = ?
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(device, id)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) GetDevice(id string) (string, error) {
	// Query
	q := `
	SELECT device
	FROM user
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return "", err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)

	device := ""

	err = row.Scan(
		&device,
	)
	if err != nil {
		return "", err
	}

	return device, err
}

func (us *UserStore) GetTraitsForUserMedication(userMedication int) ([]int, error) {
	stmt, err := us.db.Prepare(`
		SELECT umt.medication_trait FROM user_medication_trait umt 
		WHERE umt.user_medication = ?;
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(userMedication)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	traits := []int{}
	for rows.Next() {
		var trait int
		err := rows.Scan(
			&trait,
		)
		if err != nil {
			return nil, err
		}

		traits = append(traits, trait)
	}

	return traits, err
}

func (us *UserStore) GetUserMedications(userID string) ([]*model.UserMedication, error) {
	stmt, err := us.db.Prepare(`
		SELECT um.id, um.medication FROM user_medication um 
		WHERE um.user = ?;
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(userID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	userMedications := []*model.UserMedication{}
	for rows.Next() {
		userMedication := &model.UserMedication{}
		err := rows.Scan(
			&userMedication.ID,
			&userMedication.Medication,
		)
		if err != nil {
			return nil, err
		}

		userMedication.Traits, err = us.GetTraitsForUserMedication(userMedication.ID)
		if err != nil {
			return nil, err
		}

		userMedications = append(userMedications, userMedication)
	}

	return userMedications, nil
}

func (us *UserStore) GetUserMedicationByID(userID string, userMedicationID int) (*model.UserMedication, error) {
	stmt, err := us.db.Prepare(`
		SELECT um.medication FROM user_medication um 
		WHERE um.user = ? AND um.id = ?;
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(userID, userMedicationID)
	userMedication := &model.UserMedication{}
	err = row.Scan(
		&userMedication.Medication,
	)
	if err != nil {
		return nil, err
	}

	userMedication.ID = userMedicationID
	userMedication.Traits, err = us.GetTraitsForUserMedication(userMedicationID)
	if err != nil {
		return nil, err
	}

	return userMedication, nil
}

func (us *UserStore) ResearchMedication(id string, medication string) error {
	q := `
	INSERT INTO user_medication(user, medication) 
	VALUES ( ?, ? );
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, medication)
	if err != nil {
		return err
	}

	return err
}
