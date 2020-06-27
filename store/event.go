package store

import (
	"database/sql"

	"github.com/Calenaur/pandemic/model"
	"github.com/calenaur/pandemic/config"
)

type EventStore struct {
	db  *sql.DB
	cfg *config.Config
}

func NewEventStore(db *sql.DB, cfg *config.Config) *EventStore {
	return &EventStore{
		db:  db,
		cfg: cfg,
	}
}

func (es *EventStore) CreateEventFromRow(row *sql.Row) (*model.Event, error) {
	event := &model.Event{}
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Rarity,
		&event.Tier,
	)
	if err != nil {
		return nil, err
	}

	return event, err
}

func (es *EventStore) CreateEventsFromRows(rows *sql.Rows) ([]*model.Event, error) {
	var events []*model.Event
	for rows.Next() {
		event := &model.Event{}
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Rarity,
			&event.Tier,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (es *EventStore) GetByID(id int) (*model.Event, error) {
	stmt, err := es.db.Prepare(`
		SELECT e.id, e.name, e.description, e.rarity, e.tier
		FROM event e
		WHERE e.id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(id)
	medication, err := es.CreateEventFromRow(row)
	if err != nil {
		return nil, err
	}

	return medication, err
}

func (es *EventStore) GetEvents() ([]*model.Event, error) {

	q := `
	SELECT e.id, e.name, e.description, e.rarity, e.tier
	FROM event e`
	rows, err := es.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return es.CreateEventsFromRows(rows)
}

func (es *EventStore) UpdateTier(id string, tier string) error {
	// Query
	q := `
	UPDATE
	user 
	SET tier = ?
	WHERE id = ?
	`
	stmt, err := es.db.Prepare(q)
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

func (es *EventStore) SubscribeToEvent(id string, event string) error {
	// Query
	q := `
	INSERT
	INTO
	user_event( user , event )
	VALUES( ? , ? )
	`
	stmt, err := es.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, event)
	if err != nil {
		return err
	}

	return err
}

func (es *EventStore) GetMyEvents(id string) ([]*model.Event, error) {
	q := `
	SELECT e.id, e.name, e.description, e.rarity, e.tier
	FROM event e, user u, user_event ue
	WHERE u.id = ue.user AND ue.event = e.id AND u.id = ?`
	rows, err := es.db.Query(q, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return es.CreateEventsFromRows(rows)
}

func (es *EventStore) UnSubscribeToEvent(id string, event string) error {
	// Query
	q := `
	DELETE
	FROM
	user_event
	WHERE user =  ? AND event = ? 
	`
	stmt, err := es.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, event)
	if err != nil {
		return err
	}

	return err
}
