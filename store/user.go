package store

import (
	"database/sql"
	"strings"

	"github.com/calenaur/pandemic/config"
	"github.com/calenaur/pandemic/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

func (us *UserStore) GetByID(id string) (*model.User, error) {
	stmt, err := us.db.Prepare(`
		SELECT 
		id, username, accesslevel, tier, balance, manufacture
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
	SELECT id, username, accesslevel,tier, balance, manufacture
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

func (us *UserStore) GetUserDetails(id string) (string, string, string, error) {

	// Query
	q := `
	SELECT username, balance, manufacture
	FROM user
	WHERE id = ?`

	stmt, err := us.db.Prepare(q)
	if err != nil {
		return "", "", "", err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)

	username := ""
	balance := ""
	manufacture := ""

	err = row.Scan(
		&username,
		&balance,
		&manufacture,
	)
	if err != nil {
		return "", "", "", err
	}

	return string(username), string(balance), string(manufacture), err
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

func (us *UserStore) UpdateManufacture(id string, manufacture string) error {
	// Query
	q := `
	UPDATE
	user 
	SET manufacture = ?
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(manufacture, id)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) GetDiseases(id string) ([]*model.Disease, error) {

	var (
		tier        int
		name        string
		description string
		rarity      int
	)
	q := `
	SELECT d.tier , d.name , d.description, d.rarity 
	FROM disease d, user_disease ud, user u 
	WHERE d.id = ud.disease AND ud.user = u.id AND u.id = ?`

	rows, err := us.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.Disease, 0, 5)
	for rows.Next() {
		err = rows.Scan(&tier, &name, &description, &rarity)
		if err != nil {
			return nil, err
		}
		//fmt.Println(tier,name, description, rarity)
		results = append(results, &model.Disease{tier, name, description, rarity})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	//fmt.Println(results)

	return results, err
}

func (us *UserStore) GetDiseasesList(id string) ([]*model.Disease, error) {

	var (
		tier        int
		name        string
		description string
		rarity      int
	)
	q := `
	SELECT d.tier , d.name , d.description, d.rarity 
	FROM disease d, user_disease ud, user u 
	WHERE d.id = ud.disease AND ud.user = u.id AND u.id != ?`

	rows, err := us.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.Disease, 0, 10)
	for rows.Next() {
		err = rows.Scan(&tier, &name, &description, &rarity)
		if err != nil {
			return nil, err
		}
		//fmt.Println(tier,name, description, rarity)
		results = append(results, &model.Disease{tier, name, description, rarity})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	//fmt.Println(results)

	return results, err
}

func (us *UserStore) GetMedications(id string) ([]*model.Medication, error) {

	var (
		name           string
		description    string
		research_cost  int
		maximum_traits int
		rarity         int
		teir           int
	)
	q := `
	SELECT m.name, m.description, m.research_cost, m.maximum_traits, m.rarity, m.tier
	FROM medication m, user_medication um, user u 
	WHERE m.id = um.medication AND um.user = u.id AND u.id = ?`

	rows, err := us.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.Medication, 0, 5)
	for rows.Next() {
		err = rows.Scan(&name, &description, &research_cost, &maximum_traits, &rarity, &teir)
		if err != nil {
			return nil, err
		}
		//fmt.Println(tier,name, description, rarity)
		results = append(results, &model.Medication{name, description, research_cost, maximum_traits, rarity, teir})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	//fmt.Println(results)

	return results, err
}

func (us *UserStore) GetMedicationsList(id string) ([]*model.Medication, error) {

	var (
		name           string
		description    string
		research_cost  int
		maximum_traits int
		rarity         int
		teir           int
	)
	q := `
	SELECT m.name, m.description, m.research_cost, m.maximum_traits, m.rarity, m.tier
	FROM medication m, user_medication um, user u 
	WHERE m.id = um.medication AND um.user = u.id AND u.id != ?`

	rows, err := us.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.Medication, 0, 10)
	for rows.Next() {
		err = rows.Scan(&name, &research_cost, &description, &maximum_traits, &rarity, &teir)
		if err != nil {
			return nil, err
		}
		//fmt.Println(tier,name, description, rarity)
		results = append(results, &model.Medication{name, description, research_cost, maximum_traits, rarity, teir})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	//fmt.Println(results)

	return results, err
}

func (us *UserStore) ResearchMedication(id string, medication string) error {
	q := `
	INSERT INTO user_medication(user, medication) 
	VALUES ( ?,(SELECT m.id FROM medication m WHERE m.name = "?"  ));
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

func (us *UserStore) ShowFriends(id string) ([]*model.Friend, error) {

	var (
		name    string
		balance int64
		tier    int64
	)
	q := `
	SELECT f.username, f.balance, f.tier
	FROM user u, user_friend uf,user f 
	WHERE u.id = uf.user AND uf.friend = f.id AND u.id = ? AND uf.status = 1 `

	rows, err := us.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.Friend, 0, 10)
	for rows.Next() {
		err = rows.Scan(&name, &balance)
		if err != nil {
			return nil, err
		}
		//fmt.Println(tier,name, description, rarity)
		results = append(results, &model.Friend{name, balance, tier})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	//fmt.Println(results)

	return results, err
}

func (us *UserStore) UpdateTier(id string, tier string) error {
	// Query
	q := `
	UPDATE
	user 
	SET tier = ?
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tier, id)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) SendFriendRequest(id string, friendName string) error {
	// Query
	q1 := `INSERT INTO 
	user_friend (user, friend) 
	VALUES (?,(
	SELECT id FROM user WHERE username=?))`

	stmt1, err := us.db.Prepare(q1)
	if err != nil {
		return err
	}
	defer stmt1.Close()
	_, err = stmt1.Exec(id, friendName)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) RespondFriendRequest(id string, friendName string, response int64) error {
	// Query
	q := `UPDATE 
	user_friend
	SET status = ?
	WHERE user = ? AND friend = (SELECT id FROM user WHERE username = ?)`

	stmt1, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt1.Close()
	_, err = stmt1.Exec(id, friendName)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) DeleteFriend(id string, friendName string) error {
	// Query
	q := `DELETE FROM 
	user_friend
	WHERE ( user = ? AND friend = (
	SELECT id FROM user WHERE username= ? ))
	OR ( user = (SELECT id FROM user WHERE username = ? ) AND friend = ? )`

	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, friendName, friendName, id)
	if err != nil {
		return err
	}

	return err
}
