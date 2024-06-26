package server_setup

import "database/sql"

func ConnectToDatabase(dbURL string) (*sql.DB, error) {
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Verify the connection is valid
	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, err
	}

	return dbConn, nil
}
