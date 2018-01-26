package health

import (
	"database/sql"
	"log"
)

type DbSqlChecker struct {
	dns        string
	driverName string
}

func NewDbSqlChecker(driverName string, dns string) DbSqlChecker {
	return DbSqlChecker{dns: dns, driverName: driverName}
}

func (dbSqlChecker DbSqlChecker) CheckHealth() HealthCheckStatus {
	db, err := sql.Open(dbSqlChecker.driverName, dbSqlChecker.dns)
	if err != nil {
		log.Printf("Cannot acquire connection : %s \n", err)
		return Down
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Printf("Cannot acquire connection : %s \n", err)
		return Down
	}
	return Up
}
