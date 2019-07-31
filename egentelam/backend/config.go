package main

type Config struct {
	//dane do bazy Mysql
	mysqlLogin    string `json:"id"` // do pobierania z GET
	mysqlPassword string `json:"login"`
	mysqlHost     string `json:"password"`
	mysqlDatabase string `json:"userDescription"`
	mysqlPort     int    `json:points`

	//dane do bazy couchdb
	couchdbLogin    string `json:"id"` // do pobierania z GET
	couchdbPassword string `json:"login"`
	couchdbHost     string `json:"password"`
	couchdbDatabase string `json:"userDescription"`
	couchdbPort     int    `json:points`
}
