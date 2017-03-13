package postgres

import (
	"time"
)

const timeSlice = 500

type PostgresqlStatsStruct struct {
	MaxOpenConnections int
}

var postgresStats PostgresqlStatsStruct

func init() {
	go takePeriodicStats()
}

func takePeriodicStats() {
	for {

		oc := GetOpenConnections()
		if oc > postgresStats.MaxOpenConnections {
			postgresStats.MaxOpenConnections = oc
		}

		time.Sleep(timeSlice * time.Millisecond)
	}
}

func GetOpenConnections() int {
	if Database != nil {
		return Database.Stats().OpenConnections
	}
	return 0
}

func GetMaxOpenConnections() int {
	return postgresStats.MaxOpenConnections
}

func SetMaxOpenConnections(val int) {
	postgresStats.MaxOpenConnections = val
}
