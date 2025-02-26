package stores

import (
	"database/sql"
	"fmt"
)

type PostgresStore struct {
	db      *sql.DB
	dbCfg   databaseConfig
	connStr string
}

type databaseConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	DBName   string
}

/*
	Url connector key value pair str

	The final part is a query parameter that tells how to handle [SSL/TLS] encryption connection.
	Disable meaning SSL is completely disabled. For a local server security is normally guaranteed so
	it's unnecessary for now.

*/

func NewPostgresStore(db *sql.DB, dbCfg databaseConfig) (*PostgresStore, error) {

	dbCfgLoad := databaseConfig{
		Host:     dbCfg.Host,
		Port:     dbCfg.Port,
		Database: dbCfg.Database,
		Username: dbCfg.Username,
		Password: dbCfg.Password,
		DBName:   dbCfg.DBName,
	}

	//Create the urlConnectionString
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DBName)

	//create the database instance
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return &PostgresStore{
		db:      db,
		dbCfg:   dbCfgLoad,
		connStr: connStr,
	}, nil
}

/* TODO implementations for Postgres

func (s *PostgresStore) LoadUrls(ctx context.Context, urls models.URLs) error {
	// Implement the logic to load URLs into the database
	return nil
}

func (s *PostgresStore) GetUrls(ctx context.Context) (models.URLs, error) {
	// Implement the logic to retrieve URLs from the database
	return models.URLs{}, nil
}

func (s *PostgresStore) LoadStatusResponse(ctx context.Context, urls models.URLs) error {
	// Implement the logic to load status responses into the database
	return nil
}

func (s *PostgresStore) GetStatusResponse(ctx context.Context) ([]models.URLData, error) {
	// Implement the logic to retrieve status responses from the database
	return []models.URLData{}, nil
}
*/
