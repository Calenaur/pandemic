package store

import "github.com/Calenaur/pandemic/model"

func (us *UserStore) ShowFriends(id string) ([]*model.Friend, error) {

	var (
		name    string
		balance int64
		tier    string
	)
	q := `
	SELECT f.username, f.balance, t.name
	FROM user u, user_friend uf, user f , tier t
	WHERE u.id = uf.user AND uf.friend = f.id
	AND f.tier = t.id
	AND u.id = ? AND uf.status = 1`

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
		err = rows.Scan(&name, &balance, &tier)
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
	SET status = 1
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
	_, err = stmt.Exec(id, friendName, friendName, id)
	if err != nil {
		return err
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
