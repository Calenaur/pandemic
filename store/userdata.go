package store

import (
	"database/sql"

	"github.com/Calenaur/pandemic/model"
	"github.com/calenaur/pandemic/config"
)

type UserdataStore struct {
	db  *sql.DB
	cfg *config.Config
}

type UserResearcher struct {
	Tier           int    `json:"tier"`
	Research_speed int    `json:"research_speed"`
	Salary         int    `json:"salary"`
	Maximum_traits int    `json:"maximum_traits"`
	Rarity         int    `json:"rarity"`
	Name           string `json:"name"`
}

func NewUserdataStore(db *sql.DB, cfg *config.Config) *UserdataStore {
	return &UserdataStore{
		db:  db,
		cfg: cfg,
	}
}

func (ud *UserdataStore) SetUserDisease(userid string, diseaseid int) error {

	q := `
	INSERT INTO user_disease (user, disease)
	VALUES (?, ?)
	`
	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userid, diseaseid)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) GetUserDisease(userid string) ([]*model.Disease, error) {
	var (
		id          int
		tier        int
		name        string
		description string
		rarity      int
	)
	q := `
	SELECT disease.tier, disease.name, disease.description, disease.rarity
	FROM user_disease
	JOIN disease ON user_disease.disease = disease.id
	WHERE user_disease.user = ?
	`

	rows, err := ud.db.Query(q, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.Disease, 0, 10)
	for rows.Next() {
		err = rows.Scan(&id, &tier, &name, &description, &rarity)
		if err != nil {
			return nil, err
		}
		results = append(results, &model.Disease{id, tier, name, description, rarity})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, err

}

//FIXME ID to user_disease
func (ud *UserdataStore) UpdateUserDisease(userid string, diseaseid int, id int) error {
	q := `
	UPDATE user_disease
	SET user = ?, disease = ?
	WHERE id = ?
	`
	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(diseaseid, userid)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) SetUserTier(user string, tier int) error {
	q := `
	INSERT INTO user_tier (user, tier)
	VALUES (?, ?)
	`

	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user, tier)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) GetUserTier(userid string) ([]*model.Tier, error) {
	var (
		id    int
		name  string
		color string
	)
	q := `
	SELECT tier.id, tier.name, tier.color
	FROM user_tier
	JOIN tier ON user_tier.tier = tier.id
	WHERE user = ?
	`
	rows, err := ud.db.Query(q, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.Tier, 0, 10)
	for rows.Next() {
		err = rows.Scan(&id, &name, &color)
		if err != nil {
			return nil, err
		}
		results = append(results, &model.Tier{id, name, color})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, err
}

func (ud *UserdataStore) UpdateUserTier(userid string, tier int, id int) error {
	q := `
	UPDATE user_tier
	SET userid = ?, tier = ?
	WHERE id = ?
	`
	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tier, userid, id)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) SetUserEvent(userid string, eventid int) error {
	q := `
	INSERT INTO user_event (user, event)
	VALUES (?, ?)
	`
	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userid, eventid)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) GetUserEvent(userid string) ([]*model.Event, error) {
	var (
		id          int
		name        string
		description string
		rarity      int
		tier        int
	)
	q := `
	SELECT event.name, event.description, event.rarity, event.tier
	FROM user_event
	JOIN event ON user_event.event = event.id
	WHERE user = ?;
	`
	rows, err := ud.db.Query(q, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.Event, 0, 10)
	for rows.Next() {
		err = rows.Scan(&id, &name, &description, &rarity, &tier)
		if err != nil {
			return nil, err
		}
		results = append(results, &model.Event{id, name, description, rarity, tier})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, err
}

//FIXME userid or id
func (ud *UserdataStore) UpdateUserEvent(userid string, event int) error {
	q := `
	UPDATE user_event
	SET event = ?
	WHERE user = ?
	`
	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event, userid)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) SetUserResearcher(userid string, researcherid int, name string) error {
	q := `
	INSERT INTO user_researcher (user, researcher, name)
	VALUES (?, ?, ?)
	`
	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userid, researcherid, name)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) GetUserResearcher(userid string) ([]*UserResearcher, error) {
	var (
		tier           int
		research_speed int
		salary         int
		maximum_traits int
		rarity         int
		name           string
	)
	q := `
	SELECT researcher.tier, researcher.research_speed, researcher.salary, researcher.maximum_traits, researcher.rarity, user_researcher.name
	FROM user_researcher
	JOIN researcher ON user_researcher.researcher = researcher.id
	WHERE user = ?;
	`

	rows, err := ud.db.Query(q, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*UserResearcher, 0, 10)
	for rows.Next() {
		err = rows.Scan(&tier, &research_speed, &salary, &maximum_traits, &rarity, &name)
		if err != nil {
			return nil, err
		}
		results = append(results, &UserResearcher{tier, research_speed, salary, maximum_traits, rarity, name})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, err
}

func (ud *UserdataStore) UpdateUserResearcher(userid string, oldName string, newName string) error {
	q := `
	UPDATE user_researcher
	SET name = ?
	WHERE user = ? AND researcher = ?
	`
	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newName, userid, oldName)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) SetUserResearcherTrait(userid string, userResearchername string, researcherTraitname string) error {
	q := `
	INSERT INTO user_researcher_trait (user_researcher, researcher_trait)
	VALUES ((SELECT  id FROM user_researcher WHERE user = ? AND name = ?), (SELECT id FROM researcher_trait WHERE name = ?))
	`
	stmt, err := ud.db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userid, userResearchername, researcherTraitname)
	if err != nil {
		return err
	}

	return err
}

func (ud *UserdataStore) GetUserResearcherTrait(userid string) ([]*model.ResearcherTrait, error) {
	var (
		id          int
		tier        int
		name        string
		description string
		rarity      int
	)

	q := `
	SELECT researcher_trait.tier, researcher_trait.name, researcher_trait.description, researcher_trait.rarity
	FROM user_researcher
	JOIN researcher ON user_researcher.researcher = researcher.id
	JOIN user_researcher_trait ON researcher.id = user_researcher
	JOIN researcher_trait ON user_researcher.id = researcher_trait
	WHERE user = ?
	`

	rows, err := ud.db.Query(q, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	results := make([]*model.ResearcherTrait, 0, 10)
	for rows.Next() {
		err = rows.Scan(&id, &tier, &name, &description, &rarity)
		if err != nil {
			return nil, err
		}
		results = append(results, &model.ResearcherTrait{id, tier, name, description, rarity})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, err
}

//FIXME userid or id
// func (ud *UserdataStore) UpdateResearcherTrait(tier int, name string, description string, rarity int, researcherid int) error {
// 	q := `
// 	UPDATE TABLE researcer_trait
// 	SET tier = ?, name = ?, description = ?, rarity = ?
// 	WHERE id = ?
// 	`
// 	stmt, err := ud.db.Prepare(q)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(tier, name, description, rarity, researcherid)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

// TODO add researcher
// func (ud *UserdataStore) SetResearcher(tier int, researchSpeed int, salary int, maximumTraits int, rarity int) error {
// 	q := `
// 	INSERT INTO researcher (tier, research_speed, salary, maximum_traits, rarity)
// 	VALUES (?, ?, ?, ?, ?)
// 	`
// 	stmt, err := ud.db.Prepare(q)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(tier, researchSpeed, salary, maximumTraits, rarity)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

// func (ud *UserdataStore) GetResearcher(researcherId int) ([]*Researcher, error) {
// 	var (
// 		tier          int
// 		researchSpeed int
// 		salary        int
// 		maximumTraits int
// 		rarity        int
// 	)

// 	q := `
// 	SELECT researcher_trait.tier, researcher_trait.name, researcher_trait.description, researcher_trait.rarity
// 	FROM user_researcher
// 	JOIN researcher ON user_researcher.researcher = researcher.id
// 	JOIN user_researcher_trait ON researcher.id = user_researcher
// 	JOIN researcher_trait ON user_researcher.id = researcher_trait
// 	WHERE user = ?
// 	`

// 	rows, err := ud.db.Query(q, userid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	if err != nil {
// 		return nil, err
// 	}
// 	results := make([]*ResearcherTrait, 0, 10)
// 	for rows.Next() {
// 		err = rows.Scan(&tier, &name, &description, &rarity)
// 		if err != nil {
// 			return nil, err
// 		}
// 		results = append(results, &ResearcherTrait{tier, name, description, rarity})
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return results, err
// }
