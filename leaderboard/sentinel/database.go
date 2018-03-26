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

// Create team
func (u *TeamDB) Add(team *Team) (err error) {
	jsonbytes, err := json.Marshal(*team)
	if err != nil {
		panic(err)
	}
	fmt.Println("--------")
	fmt.Println(string(jsonbytes))
	err = u.client.UpsertDocument(u.coll.Self, team)
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
func (u *TeamDB) findOrCreateCollection(name string) (err error) {
	if colls, err := u.client.QueryCollections(u.db.Self, fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(colls) == 0 {
		if coll, err := u.client.CreateCollection(u.db.Self, fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {
			u.coll = coll
		}
	} else {
		u.coll = &colls[0]
	}
	return
}

// Find or create database by id
func (u *TeamDB) findOrDatabase(name string) (err error) {
	if dbs, err := u.client.QueryDatabases(fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(dbs) == 0 {
		if db, err := u.client.CreateDatabase(fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {
			u.db = db
		}
	} else {
		u.db = &dbs[0]
	}
	return
}
