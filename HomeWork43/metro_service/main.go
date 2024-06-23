package main

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
}
