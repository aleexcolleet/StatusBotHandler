package stores

import (
	"MST_FV/config"
	"MST_FV/internal/domain/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresStore struct {
	db      *sql.DB //database connection
	dbCfg   databaseConfig
	connStr string
}

type databaseConfig struct {
	Port     string
	Host     string
	Database string
	User     string
	Password string
	DBName   string
}

/*
Url connector key value pair str

The final part is a query parameter that tells how to handle [SSL/TLS] encryption connection.
Disable meaning SSL is completely disabled. For a local server security is normally guaranteed so
it's unnecessary for now.
*/

func NewPostgresStore(cfg config.Config) (*PostgresStore, error) {

	dbCfgLoad := databaseConfig{
		Port:     cfg.Database.Port,
		Host:     cfg.Database.Host,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
	}

	//Create the urlConnectionString
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbCfgLoad.User, dbCfgLoad.Password, dbCfgLoad.Host, dbCfgLoad.Port, dbCfgLoad.DBName)

	//create the database instance
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("connection string not valid: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	} //connection test

	return &PostgresStore{
		db:      db,
		dbCfg:   dbCfgLoad,
		connStr: connStr,
	}, nil
}

// LoadUrls adaptation saves a list of URLs into the urls table in database
func (s *PostgresStore) LoadUrls(ctx context.Context, urls models.URLs) error {

	for _, url := range urls.Urls {
		_, err := s.db.ExecContext(ctx, `INSERT INTO urls(url) VALUES($1) ON CONFLICT(url) DO NOTHING`, url)
		if err != nil {
			return fmt.Errorf("error inserting url: %v", err)
		}
	}
	return nil
}

// GetUrls recovers all urls store in db and returns it in the models.URLs struct
func (s *PostgresStore) GetUrls(ctx context.Context) (models.URLs, error) {

	var urls models.URLs
	var url string

	rows, err := s.db.QueryContext(ctx, `SELECT url FROM urls`)
	if err != nil {
		return models.URLs{}, fmt.Errorf("error querying urls on GetUrls: %v", err)
	}
	defer rows.Close()

	for rows.Next() { // goes across rows one by one
		err = rows.Scan(&url)
		if err != nil {
			return models.URLs{}, fmt.Errorf("error scanning rows in GetUrls: %v", err)
		}
		urls.Urls = append(urls.Urls, url)
	}
	if err = rows.Err(); err != nil { //Check whether rows had any error at the end
		return models.URLs{}, fmt.Errorf("error iterating rows: %v", err)
	}
	return urls, nil
}

// LoadStatusResponse stores and actualizes into the db the Urls and UrlsData response from http requests
func (s *PostgresStore) LoadStatusResponse(ctx context.Context, urls models.URLs) error {

	if len(urls.UrlsData) == 0 { //small check
		return nil
	}

	//if url already exists, update it before failing
	for _, urlData := range urls.UrlsData {
		_, err := s.db.ExecContext(ctx, `
			INSERT INTO url_status (url, status, comment, status_code)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (url) DO UPDATE
			SET status = $2, COMMENT = $3, status_code = $4
		`, urlData.Url, urlData.Status, urlData.Comment, urlData.StatusCode)
		if err != nil {
			return fmt.Errorf("error inserting url status: %v", err)
		}
	}
	return nil
}

// GetStatusResponse retrieves all URL status responses from url_status table
func (s *PostgresStore) GetStatusResponse(ctx context.Context) ([]models.URLData, error) {

	var urlsData []models.URLData

	rows, err := s.db.QueryContext(ctx, `SELECT url, status, comment, status_code FROM url_status`)
	if err != nil {
		return nil, fmt.Errorf("error querying urls on GetStatusResponse: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var urlData models.URLData
		err = rows.Scan(&urlData.Url, &urlData.Status, &urlData.Comment, &urlData.StatusCode)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows in GetStatusResponse: %v", err)
		}
		urlsData = append(urlsData, urlData)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows in GetStatusResponse: %v", err)
	}

	return urlsData, nil
}
