package package_initialization

var connection string

func init() {
	connection = "mysql"
}

func GetDatabase() string {
	return connection
}
