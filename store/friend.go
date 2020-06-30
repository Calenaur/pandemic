package store

import (
	"errors"

	"github.com/Calenaur/pandemic/model"
)

func (us *UserStore) ShowFriends(id string, user string) ([]*model.Friend, error, error) {

	var (
		name    string
		balance int
		tier    int
	)
	q1 := `
	SELECT f.username, f.balance, f.tier
	FROM user_friend uf
	JOIN user u ON u.id = uf.user
	JOIN user f ON f.id = uf.friend
	WHERE (u.id = ? OR f.id = ?)
	AND uf.status = 1
	GROUP BY f.username`

	q2 := `
	SELECT u.username, u.balance, u.tier
	FROM user_friend uf
	JOIN user u ON u.id = uf.user
	JOIN user f ON f.id = uf.friend
	WHERE (u.id = ? OR f.id = ?)
 	AND uf.status = 1
 	GROUP BY u.username`

	rows1, err1 := us.db.Query(q1, id, id)
	if err1 != nil {
		return nil, err1, nil
	}
	rows2, err2 := us.db.Query(q2, id, id)
	if err2 != nil {
		return nil, err2, nil
	}
	defer rows1.Close()
	if err1 != nil {
		return nil, err1, nil
	}
	defer rows2.Close()
	if err2 != nil {
		return nil, err2, nil
	}

	// make slice
	results := make([]*model.Friend, 0, 10)
	names := make([]string, 0, 10)
	for rows1.Next() {
		err1 = rows1.Scan(&name, &balance, &tier)
		if err1 != nil {
			return nil, err1, err2
		}
		//fmt.Println(tier,name, description, rarity)
		if name == user || Contains(names, name) {
		} else {
			results = append(results, &model.Friend{name, balance, tier})
			names = append(names, name)
		}

	}
	err1 = rows1.Err()
	if err1 != nil {
		return nil, err1, nil
	}

	for rows2.Next() {
		err2 = rows2.Scan(&name, &balance, &tier)
		if err2 != nil {
			return nil, err2, nil
		}
		//fmt.Println(tier,name, description, rarity)
		if name == user || Contains(names, name) {
		} else {
			results = append(results, &model.Friend{name, balance, tier})
			names = append(names, name)
		}

	}
	err2 = rows2.Err()
	if err2 != nil {
		return nil, nil, err2
	}

	//fmt.Println(results)

	return results, err1, err2
}

func (us *UserStore) ShowPendingFriends(id string, user string) ([]*model.Friend, error, error) {

	var (
		name    string
		balance int
		tier    int
	)
	q1 := `
	SELECT f.username, f.balance, f.tier
	FROM user_friend uf
	JOIN user u ON u.id = uf.user
	JOIN user f ON f.id = uf.friend
	WHERE (u.id = ? OR f.id = ?)
	AND uf.status = 0
	GROUP BY f.username`

	q2 := `
	SELECT u.username, u.balance, u.tier
	FROM user_friend uf
	JOIN user u ON u.id = uf.user
	JOIN user f ON f.id = uf.friend
	WHERE (u.id = ? OR f.id = ?)
 	AND uf.status = 0
 	GROUP BY u.username`

	rows1, err1 := us.db.Query(q1, id, id)
	if err1 != nil {
		return nil, err1, nil
	}
	rows2, err2 := us.db.Query(q2, id, id)
	if err2 != nil {
		return nil, err2, nil
	}
	defer rows1.Close()
	if err1 != nil {
		return nil, err1, nil
	}
	defer rows2.Close()
	if err2 != nil {
		return nil, err2, nil
	}

	// make slice
	results := make([]*model.Friend, 0, 10)
	names := make([]string, 0, 10)
	for rows1.Next() {
		err1 = rows1.Scan(&name, &balance, &tier)
		if err1 != nil {
			return nil, err1, err2
		}
		//fmt.Println(tier,name, description, rarity)
		if name == user || Contains(names, name) {
		} else {
			results = append(results, &model.Friend{name, balance, tier})
			names = append(names, name)
		}

	}
	err1 = rows1.Err()
	if err1 != nil {
		return nil, err1, nil
	}

	for rows2.Next() {
		err2 = rows2.Scan(&name, &balance, &tier)
		if err2 != nil {
			return nil, err2, nil
		}
		//fmt.Println(tier,name, description, rarity)
		if name == user || Contains(names, name) {
		} else {
			results = append(results, &model.Friend{name, balance, tier})
			names = append(names, name)
		}

	}
	err2 = rows2.Err()
	if err2 != nil {
		return nil, nil, err2
	}

	//fmt.Println(results)

	return results, err1, err2
}

func (us *UserStore) SendFriendRequest(id string, friendName string) error {
	// Query
	q1 := `INSERT INTO 
	user_friend (user, friend, status) 
	VALUES (?,(
	SELECT id FROM user WHERE username=?), 0)`

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
	var q = `UPDATE 
	user_friend
	SET status = ?
	WHERE (user = ? AND friend = (SELECT id FROM user WHERE username = ?)) OR (user = (SELECT id FROM user WHERE username = ?) AND friend = ?) `
	stmt1, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt1.Close()
	_, err = stmt1.Exec(response, id, friendName, friendName, id)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStore) DeleteFriend(id string, friendName string) error {
	// Query
	var q = `DELETE FROM 
	user_friend
	WHERE ( user = ? AND friend = (
	SELECT id FROM user WHERE username= ? ))
	OR ( user = (SELECT id FROM user WHERE username = ? ) AND friend = ? )`
	stmt, err := us.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id, friendName, friendName, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		err = errors.New("Method failed")
	}
	return err
}

func (us *UserStore) SendFriendBalance(id string, friendName string, balance string) error {
	var status string
	// Query
	q := `
	SELECT status
	FROM user_friend
	WHERE (user = ? AND friend = (SELECT id FROM user WHERE username = ?)) OR (user = (SELECT id FROM user WHERE username = ?) AND friend = ?)
`

	row, err := us.db.Query(q, id, friendName, friendName, id)
	if err != nil {
		return err
	}
	defer row.Close()
	if err != nil {
		return err
	}
	for row.Next() {
		err = row.Scan(&status)
		if err != nil {
			return err
		}
	}
	//println(status)
	if status == "1" {
		//Also Query
		q0 := `
		UPDATE
		user SET balance = (balance + ?)
		WHERE username = ?`

		stmt0, err := us.db.Prepare(q0)
		if err != nil {
			return err
		}
		defer stmt0.Close()
		_, err = stmt0.Exec(balance, friendName)
		if err != nil {
			return err
		}

		q1 := `
		UPDATE
		user SET balance = (balance - ?)
		WHERE id = ?`

		stmt1, err := us.db.Prepare(q1)
		if err != nil {
			return err
		}
		defer stmt1.Close()
		_, err = stmt1.Exec(balance, id)
		if err != nil {
			return err
		}

	}
	return err
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
