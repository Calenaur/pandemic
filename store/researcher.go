package store

import (
	"database/sql"

	"github.com/Calenaur/pandemic/model"
	"github.com/calenaur/pandemic/config"
)

type ResearcherStore struct {
	db  *sql.DB
	cfg *config.Config
}

func NewResearcherStore(db *sql.DB, cfg *config.Config) *ResearcherStore {
	return &ResearcherStore{
		db:  db,
		cfg: cfg,
	}
}

func (rs *ResearcherStore) CreateResearcherFromRow(row *sql.Row) (*model.Researcher, error) {
	researcher := &model.Researcher{}
	err := row.Scan(
		&researcher.ID,
		&researcher.Tier,
		&researcher.ResearcherSpeed,
		&researcher.Salary,
		&researcher.MaximumTraits,
		&researcher.Rarity,
	)
	if err != nil {
		return nil, err
	}

	return researcher, err
}

func (rs *ResearcherStore) CreateResearchersFromRows(rows *sql.Rows) ([]*model.Researcher, error) {
	researchers := []*model.Researcher{}
	for rows.Next() {
		researcher := &model.Researcher{}
		err := rows.Scan(
			&researcher.ID,
			&researcher.Tier,
			&researcher.ResearcherSpeed,
			&researcher.Salary,
			&researcher.MaximumTraits,
			&researcher.Rarity,
		)
		if err != nil {
			return nil, err
		}

		researchers = append(researchers, researcher)
	}

	return researchers, nil
}

func (rs *ResearcherStore) CreateResearcherTraitFromRow(row *sql.Row) (*model.ResearcherTrait, error) {
	researcherTrait := &model.ResearcherTrait{}
	err := row.Scan(
		&researcherTrait.ID,
		&researcherTrait.Tier,
		&researcherTrait.Name,
		&researcherTrait.Description,
		&researcherTrait.Rarity,
	)
	if err != nil {
		return nil, err
	}

	return researcherTrait, err
}

func (rs *ResearcherStore) CreateResearcherTraitsFromRows(rows *sql.Rows) ([]*model.ResearcherTrait, error) {
	researcherTraits := []*model.ResearcherTrait{}
	for rows.Next() {
		researcherTrait := &model.ResearcherTrait{}
		err := rows.Scan(
			&researcherTrait.ID,
			&researcherTrait.Tier,
			&researcherTrait.Name,
			&researcherTrait.Description,
			&researcherTrait.Rarity,
		)
		if err != nil {
			return nil, err
		}

		researcherTraits = append(researcherTraits, researcherTrait)
	}

	return researcherTraits, nil
}

func (rs *ResearcherStore) GetByID(id int) (*model.Researcher, error) {
	stmt, err := rs.db.Prepare(`
		SELECT m.id, m.name, m.description, m.research_cost, m.maximum_traits, m.base_value, m.tier
		FROM researcher m
		WHERE m.id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	researcher, err := rs.CreateResearcherFromRow(row)
	if err != nil {
		return nil, err
	}

	return researcher, err
}

func (rs *ResearcherStore) Getresearchers() ([]*model.Researcher, error) {

	q := `
	SELECT m.id, m.name, m.description, m.research_cost, m.maximum_traits, m.base_value, m.tier
	FROM researcher m`
	rows, err := rs.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return rs.CreateResearchersFromRows(rows)
}

func (rs *ResearcherStore) GetTraitByID(id int) (*model.ResearcherTrait, error) {
	stmt, err := rs.db.Prepare(`
		SELECT mt.id, mt.name, mt.description, mt.tier
		FROM researcher_trait mt
		WHERE mt.id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	researcher, err := rs.CreateResearcherTraitFromRow(row)
	if err != nil {
		return nil, err
	}

	return researcher, err
}

func (rs *ResearcherStore) GetTraits() ([]*model.ResearcherTrait, error) {

	q := `
		SELECT mt.id, mt.name, mt.description, mt.tier
		FROM researcher_trait mt
		`
	rows, err := rs.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return rs.CreateResearcherTraitsFromRows(rows)
}
