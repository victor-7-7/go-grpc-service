package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"go-course/grpc-service/internal/rocket"
	"log"
	"os"
)

// https://github.com/jmoiron/sqlx

type Schema struct {
	create string
	drop   string
}

// https://jmoiron.github.io/sqlx/

var defaultSchema = Schema{
	create: `
CREATE TABLE IF NOT EXISTS rockets (
    id text,
    name text,
    type text,
    flights integer DEFAULT 0
);
`,
	drop: `DROP TABLE IF EXISTS rockets;`,
}

type Store struct {
	db *sqlx.DB
}

// New - returns a new Store object.
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbName, dbPassword, dbSSLMode)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}

	db.MustExec(defaultSchema.create)
	return Store{db: db}, nil
}

// GetRocketById - retrieves a rocket from the database by id.
func (s Store) GetRocketById(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	row := s.db.QueryRow(`SELECT id, name, type FROM rockets WHERE id=$1;`, id)
	err := row.Scan(&rkt.Id, &rkt.Name, &rkt.Type)
	//err := s.db.Get(&rkt, `SELECT * FROM rockets WHERE id=$1`, id)
	if err != nil {
		log.Println(err.Error())
		return rocket.Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - inserts a rocket into the rockets table.
func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	log.Println("Store | InsertRocket:", fmt.Sprintf("%#v", rkt))
	//_, err := s.db.Exec(`INSERT INTO rockets (id, name, type) VALUES ($1, $2, $3)`, rkt.Id, rkt.Name, rkt.Type)
	_, err := s.db.NamedExec(`INSERT INTO rockets (id, name, type) VALUES (:id, :name, :type)`, rkt)

	if err != nil {
		return rocket.Rocket{}, errors.New("failed to insert rocket into db")
	}
	return rocket.Rocket{
		Id:   rkt.Id,
		Name: rkt.Name,
		Type: rkt.Type,
	}, nil
}

func (s Store) DeleteRocket(id string) error {
	uid, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("DELETE FROM rockets WHERE id=$1", uid)
	if err != nil {
		return err
	}
	return nil
}
