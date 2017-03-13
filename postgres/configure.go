package postgres


var maxIdleConns = 0
var maxOpenConns = 0

func GetMaxIdleConns() int {
	return maxIdleConns
}

func SetMaxIdleConns(conns int) {
	maxIdleConns = conns
	Database.SetMaxIdleConns(maxIdleConns)
}

func GetMaxOpenConns() int {
	return maxOpenConns
}

func SetMaxOpenConns(conns int) {
	maxOpenConns = conns
	Database.SetMaxOpenConns(maxOpenConns)
}
