package store

type Users struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessLevel int    `json:"accesslevel"`
	Balance     int    `json:"balance"`
	Manufacture int    `json:"manufacture"`
}

func (us *UserStore) ListAll(offset int64) ([]*Users, error) {
	// Query
	var (
		// users       []Users
		id          string
		username    string
		accesslevel int
		balance     int
		manufacture int
	)
	q := `
	SELECT id, username, accesslevel, balance, manufacture
	FROM user
	LIMIT 10 OFFSET ?;`

	rows, err := us.db.Query(q, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*Users, 0, 10)
	for rows.Next() {
		err = rows.Scan(&id, &username, &accesslevel, &balance, &manufacture)
		if err != nil {
			return nil, err
		}
		results = append(results, &Users{id, username, accesslevel, balance, manufacture})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, err
}
