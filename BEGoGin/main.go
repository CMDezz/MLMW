package main

import (
	"MLMW/BEGoGin/infras/apis"
	"MLMW/BEGoGin/infras/connections"
	"MLMW/BEGoGin/utils"
	"log"
)

func main() {
	//Init ENV Config
	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatal("INIT ERR: Cannot load config \n", err)
	}

	//Init DB Connection
	sqlxDB, err := connections.InitDBConnection(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("INIT ERR: Cannot connect to Database \n", err)
	}

	// migrate db
	err = utils.MigrateDatabase(config.MigrationDir, config.DBSource)
	if err != nil {
		log.Fatal("INIT ERR: Cannot excute the migration \n", err)
	}

	// init server connection
	server, err := apis.InitServerConnection(config, sqlxDB)
	if err != nil {
		log.Fatal("INIT ERR: Cannot connect to Server \n", err)
	}

	// => START THE APP
	if err = server.Start(config.SVAddress); err != nil {
		log.Fatal("INIT ERR: Cannot start the Server \n", err)
	}
}
