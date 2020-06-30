package store

import (
	"database/sql"
	"github.com/Calenaur/pandemic/model"
	"github.com/calenaur/pandemic/config"
	_ "github.com/go-sql-driver/mysql"
)

type MedicationStore struct {
	db  *sql.DB
	cfg *config.Config
}

func NewMedicationStore(db *sql.DB, cfg *config.Config) *MedicationStore {
	return &MedicationStore{
		db:  db,
		cfg: cfg,
	}
}

func (ms *MedicationStore) CreateMedicationFromRow(row *sql.Row) (*model.Medication, error) {
	medication := &model.Medication{}
	err := row.Scan(
		&medication.ID,
		&medication.Name,
		&medication.Description,
		&medication.ResearchCost,
		&medication.MaximumTraits,
		&medication.BaseValue,
		&medication.Tier,
	)
	if err != nil {
		return nil, err
	}

	return medication, err
}

func (ms *MedicationStore) CreateMedicationsFromRows(rows *sql.Rows) ([]*model.Medication, error) {
	var medications []*model.Medication
	for rows.Next() {
		medication := &model.Medication{}
		err := rows.Scan(
			&medication.ID,
			&medication.Name,
			&medication.Description,
			&medication.ResearchCost,
			&medication.MaximumTraits,
			&medication.BaseValue,
			&medication.Tier,
		)
		if err != nil {
			return nil, err
		}

		medications = append(medications, medication)
	}

	return medications, nil
}

func (ms *MedicationStore) CreateMedicationTraitFromRow(row *sql.Row) (*model.MedicationTrait, error) {
	medicationTrait := &model.MedicationTrait{}
	err := row.Scan(
		&medicationTrait.ID,
		&medicationTrait.Name,
		&medicationTrait.Description,
		&medicationTrait.Tier,
	)
	if err != nil {
		return nil, err
	}

	return medicationTrait, err
}

func (ms *MedicationStore) CreateMedicationTraitsFromRows(rows *sql.Rows) ([]*model.MedicationTrait, error) {
	medicationTraits := []*model.MedicationTrait{}
	for rows.Next() {
		medicationTrait := &model.MedicationTrait{}
		err := rows.Scan(
			&medicationTrait.ID,
			&medicationTrait.Name,
			&medicationTrait.Description,
			&medicationTrait.Tier,
		)
		if err != nil {
			return nil, err
		}

		medicationTraits = append(medicationTraits, medicationTrait)
	}

	return medicationTraits, nil
}

func (ms *MedicationStore) GetByID(id int) (*model.Medication, error) {
	stmt, err := ms.db.Prepare(`
		SELECT m.id, m.name, m.description, m.research_cost, m.maximum_traits, m.base_value, m.tier
		FROM medication m
		WHERE m.id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	medication, err := ms.CreateMedicationFromRow(row)
	if err != nil {
		return nil, err
	}

	return medication, err
}

func (ms *MedicationStore) GetMedications() ([]*model.Medication, error) {

	q := `
	SELECT m.id, m.name, m.description, m.research_cost, m.maximum_traits, m.base_value, m.tier
	FROM medication m`
	rows, err := ms.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ms.CreateMedicationsFromRows(rows)
}

func (ms *MedicationStore) GetTraitByID(id int) (*model.MedicationTrait, error) {
	stmt, err := ms.db.Prepare(`
		SELECT mt.id, mt.name, mt.description, mt.tier
		FROM medication_trait mt
		WHERE mt.id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	medication, err := ms.CreateMedicationTraitFromRow(row)
	if err != nil {
		return nil, err
	}

	return medication, err
}

func (ms *MedicationStore) GetTraits() ([]*model.MedicationTrait, error) {

	q := `
		SELECT mt.id, mt.name, mt.description, mt.tier
		FROM medication_trait mt
		`
	rows, err := ms.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ms.CreateMedicationTraitsFromRows(rows)
}

func (ms *MedicationStore) AddMedicationAndTraits(medication string, traits []string) error {
	for _, s := range traits {
		q := `
		INSERT INTO
		user_medication_trait
		VALUES (? , ?)
		`
		stmt1, err := ms.db.Prepare(q)
		if err != nil {
			return err
		}
		defer stmt1.Close()
		_, err = stmt1.Exec(medication, s)
		if err != nil {
			return err
		}
	}
	return nil
}
