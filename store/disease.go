package store

import (
	"database/sql"

	"github.com/Calenaur/pandemic/model"
	"github.com/calenaur/pandemic/config"
)

type DiseaseStore struct {
	db  *sql.DB
	cfg *config.Config
}

func NewDiseaseStore(db *sql.DB, cfg *config.Config) *DiseaseStore {
	return &DiseaseStore{
		db:  db,
		cfg: cfg,
	}
}

func (ds *DiseaseStore) CreateDiseaseFromRow(row *sql.Row) (*model.Disease, error) {
	disease := &model.Disease{}
	err := row.Scan(
		&disease.ID,
		&disease.Tier,
		&disease.Name,
		&disease.Description,
		&disease.Rarity,
	)
	if err != nil {
		return nil, err
	}

	return disease, err
}

func (ds *DiseaseStore) CreateDiseasesFromRows(rows *sql.Rows) ([]*model.Disease, error) {
	diseases := []*model.Disease{}
	for rows.Next() {
		disease := &model.Disease{}
		err := rows.Scan(
			&disease.ID,
			&disease.Tier,
			&disease.Name,
			&disease.Description,
			&disease.Rarity,
		)
		if err != nil {
			return nil, err
		}

		diseases = append(diseases, disease)
	}

	return diseases, nil
}

func (ds *DiseaseStore) GetByID(id int) (*model.Disease, error) {
	stmt, err := ds.db.Prepare(`
		SELECT d.id, d.tier, d.name, d.description, d.rarity
		FROM disease d
		WHERE d.id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	medication, err := ds.CreateDiseaseFromRow(row)
	if err != nil {
		return nil, err
	}

	return medication, err
}

func (ds *DiseaseStore) GetDiseases() ([]*model.Disease, error) {

	q := `
	SELECT d.id, d.tier, d.name, d.description, d.rarity
	FROM disease d`
	rows, err := ds.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ds.CreateDiseasesFromRows(rows)
}

func (ds *DiseaseStore) GetDiseasesForUser(id string) ([]*model.Disease, error) {

	q := `
	SELECT d.id , d.tier , d.name , d.description, d.rarity 
	FROM disease d, user_disease ud, user u 
	WHERE d.id = ud.disease AND ud.user = u.id AND u.id = ?`

	rows, err := ds.db.Query(q, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ds.CreateDiseasesFromRows(rows)
}

func (ds *DiseaseStore) GetDiseasesList(id string) ([]*model.Disease, error) {

	q := `
	SELECT d.id , d.tier , d.name , d.description, d.rarity 
	FROM disease d, user_disease ud, user u 
	WHERE d.id = ud.disease AND ud.user = u.id AND u.id != ?`

	rows, err := ds.db.Query(q, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ds.CreateDiseasesFromRows(rows)
}

func (ds *DiseaseStore) SelectDisease(id string, disease string) error {
	q := `
	INSERT
	INTO
	user_disease
	VAlUES (? , ?)
	`
	stmt1, err := ds.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt1.Close()
	_, err = stmt1.Exec(id, disease)
	if err != nil {
		return err
	}

	return err
}

func (ds *DiseaseStore) UnSelectDisease(id string, disease string) error {
	q := `
	DELETE
	FROM
	user_disease
	WHERE user = ? AND disease = ?
	`
	stmt1, err := ds.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt1.Close()
	_, err = stmt1.Exec(id, disease)
	if err != nil {
		return err
	}

	return err
}
