package patterns

type IDbConnection interface{
	Connect()
}

type DbConnection struct {
	Db IDbConnection
}

func (con DbConnection) DbConnect() {
	con.Db.Connect()
}

type MySqlConnection struct {
	ConnectionString string
}
func (con MySqlConnection) Connect() {
	println("MySQL", con.ConnectionString)
}

type PostressConnection struct {
	ConnectionString string
}
func (con PostressConnection) Connect() {
	println("Postgress", con.ConnectionString)
}

type MongoConnection struct {
	ConnectionString string
}
func (con MongoConnection) Connect() {
	println("Moongo", con.ConnectionString)
}

func Strategy() {
	MySqlConnection := MongoConnection{ConnectionString: "MySqlConString"}
	con := DbConnection{Db: MySqlConnection}
	con.DbConnect()

	PostressConnection := PostressConnection{ConnectionString: "PostgressConString"}
	con2 := DbConnection{Db: PostressConnection}
	con2.DbConnect()
}
