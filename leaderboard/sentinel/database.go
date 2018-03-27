package main

import (
	"encoding/json"
	"fmt"

	documentdb "github.com/TsuyoshiUshio/documentdb-go"
)

// DB interface
type DB interface {
	Get(id string) *Team
	GetAll() []*Team
	Add(u *Team) *Team
	Update(u *Team) *Team
	Remove(id string) error
}

// UsersDB implement DB interface
type TeamDB struct {
	Database   string
	Collection string
	db         *documentdb.Database
	coll       *documentdb.Collection
	client     *documentdb.DocumentDB
}

// DocumentDB config
type Config struct {
	Url       string `json:"url"`
	MasterKey string `json:"masterKey"`
}

// Return new UserDB
// Test if database and collection exist. if not, create them.
func NewDB(db, coll string, config *Config) (teamdb TeamDB) {
	teamdb.Database = db
	teamdb.Collection = coll
	teamdb.client = documentdb.New(config.Url, documentdb.Config{config.MasterKey})
	// Find or create `test` db and `users` collection
	if err := teamdb.findOrDatabase(db); err != nil {
		panic(err)
	}
	if err := teamdb.findOrCreateCollection(coll); err != nil {
		panic(err)
	}
	return
}

// Get All teams
func (t *TeamDB) GetAll() (teams *[]Team, err error) {
	err = t.client.ReadDocuments(t.coll.Self, &teams)
	return
}
func (t *TeamDB) RemoveDB() error {
	return t.client.DeleteDatabase(t.db.Self)
}

// Create team
func (t *TeamDB) Add(team *Team) (err error) {
	jsonbytes, err := json.Marshal(*team)
	if err != nil {
		panic(err)
	}
	fmt.Println("--------")
	fmt.Println(string(jsonbytes))
	err = t.client.UpsertDocument(t.coll.Self, team)
	if err != nil {
		// Currently, If you use Upsert, we've got and error with it's body as "
		// It is because the Status Code is not match 200 vs 201. However it's OK
		// Upsert case. See the detail https://github.com/a8m/documentdb-go/pull/7
		if ", " != err.Error() {
			return err
		}
	}
	return nil
}

// Find or create collection by id
func (t *TeamDB) findOrCreateCollection(name string) (err error) {
	if colls, err := t.client.QueryCollections(t.db.Self, fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(colls) == 0 {
		if coll, err := t.client.CreateCollection(t.db.Self, fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {
			t.coll = coll
		}
	} else {
		t.coll = &colls[0]
	}
	return
}

// Find or create database by id
func (t *TeamDB) findOrDatabase(name string) (err error) {
	if dbs, err := t.client.QueryDatabases(fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(dbs) == 0 {
		if db, err := t.client.CreateDatabase(fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {
			t.db = db
		}
	} else {
		t.db = &dbs[0]
	}
	return
}
