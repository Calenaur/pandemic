package store

import (
	"database/sql"

	"github.com/Calenaur/pandemic/model"
	"github.com/calenaur/pandemic/config"
)

type TierStore struct {
	db  *sql.DB
	cfg *config.Config
}

func NewTierStore(db *sql.DB, cfg *config.Config) *TierStore {
	return &TierStore{
		db:  db,
		cfg: cfg,
	}
}

func (ts *TierStore) CreateTierFromRow(row *sql.Row) (*model.Tier, error) {
	tier := &model.Tier{}
	err := row.Scan(
		&tier.ID,
		&tier.Name,
		&tier.Color,
	)
	if err != nil {
		return nil, err
	}

	return tier, err
}

func (ts *TierStore) CreateTiersFromRows(rows *sql.Rows) ([]*model.Tier, error) {
	var tiers []*model.Tier
	for rows.Next() {
		tier := &model.Tier{}
		err := rows.Scan(
			&tier.ID,
			&tier.Name,
			&tier.Color,
		)
		if err != nil {
			return nil, err
		}

		tiers = append(tiers, tier)
	}

	return tiers, nil
}

func (ts *TierStore) GetTierByID(id int) (*model.Tier, error) {
	stmt, err := ts.db.Prepare(`
	SELECT id, name, color
	FROM tier
	WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	tier, err := ts.CreateTierFromRow(row)
	if err != nil {
		return nil, err
	}

	return tier, err
}

func (ts *TierStore) GetTierList() ([]*model.Tier, error) {
	q := `
	SELECT id, name, color
	FROM tier
	`
	rows, err := ts.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ts.CreateTiersFromRows(rows)
}

func (ts *TierStore) SetTier(name string, color string) error {
	q := `
	INSERT INTO tier (name, color)
	VALUES (?, ?)
	`
	stmt, err := ts.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, color)
	if err != nil {
		return err
	}

	return err
}

func (ts *TierStore) UpdateTier(id int, name string, color string) error {
	q := `
	UPDATE tier
	SET name = ?, color = ?
	WHERE id = ?
	`
	stmt, err := ts.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, color, id)
	if err != nil {
		return err
	}

	return err
}

func (ts *TierStore) DeleteTier(id int) error {
	q := `
	DELETE FROM tier
	WHERE id = ?
	`
	stmt, err := ts.db.Prepare(q)
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
