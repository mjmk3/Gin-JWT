package main

func main() {
	database.Connect("root:root@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()
}
