package main

import "flag"

const(

	// Database connection string
	// Will be initialize after parsing application run flags
	DB_CONNECTION_STRING string
	
	// Database connection settings
	DB_HOST *string = flag.String("dbhost", "localhost", "Database host")
	DB_PORT *string = flag.String("dbport", "5432", "Database port")
	DB_USER *string = flag.String("dbuser", "ezarechkin", "Database user")
	DB_NAME *string = flag.String("dbname", "point", "Database name")
	DB_PASS *string = flag.String("dbpass", "", "Database password")
	
	// Tcp ))) listener bind address
	BIND_ADDRESS *string = flag.String("bind_address", "127.0.0.1:3000", "Tcp listener address")
)