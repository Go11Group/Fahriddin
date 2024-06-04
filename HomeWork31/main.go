package main

import (
  "database/sql"
  "fmt"
  "math/rand"
  "time"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = "5432"
  user     = "postgres"
  dbname   = "nt"
  password = "0412"
)

func main() {
  conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password)
  db, err := sql.Open("postgres", conn)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  // 1300000 ta malumot insert qilish

  // for i := 0; i < 1300000; i++ {
  //   _, err = db.Exec("insert into persons(id,last_name,first_name,age,courdse,phone) values($1,$2,$3,$4,$5,$6)",
  //     faker.UUIDDigit(), faker.LastName(), faker.FirstName(), rand.Intn(100), rand.Intn(5), faker.Phonenumber())
  //   if err != nil {
  //     panic(err)
  //   }
  //   if i%20000 == 0 {
  //     fmt.Print(i)
  //   }
  // }
  var querys string
  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(10)
  for i := 0; i < 300; i++ {
    go func() {
      curs := rand.Intn(5)
      age := rand.Intn(100)
      t := time.Now()
      err := db.QueryRow("select count(1) from persons where courdse=$1 and age=$2", curs, age).Scan(&querys)
      if err != nil {
        panic(err)
      }
      fmt.Println(i, querys, time.Now().Sub(t))

    }()
  }
  time.Sleep(5 * time.Second)

}
