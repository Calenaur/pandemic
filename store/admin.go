package store

import (
	"github.com/Calenaur/pandemic/model"
)

func (us *UserStore) ListAll(offset int64, limit int64) ([]*model.User, error) {
	// Query
	var (
		// users       []Users
		id          string
		username    string
		accesslevel int
		tier        int
		balance     int
	)
	q := `
	SELECT id, username, accesslevel, tier, balance
	FROM user
	LIMIT ? OFFSET ?;`

	rows, err := us.db.Query(q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.User, 0, 10)
	for rows.Next() {
		err = rows.Scan(&id, &username, &accesslevel, &tier, &balance)
		if err != nil {
			return nil, err
		}
		results = append(results, &model.User{id, username, accesslevel, tier, balance})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, err
}

func (us *UserStore) UserCount() (int, error) {
	var count int
	q := `
	SELECT COUNT(*)
	FROM user;`

	err := us.db.QueryRow(q).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (us *UserStore) MakeUserAdmin(userId string) error {

	q := `
	UPDATE
	user
	SET accesslevel = 100
	WHERE id = ?
	`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId)
	if err != nil {
		return err
	}

	return err

}

func (us *UserStore) DeleteUser(userId string) error {
	//TODO make sure its save delete for foreign keys
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
	_, err = stmt.Exec(userId)
	if err != nil {
		return err
	}

	return err

}
