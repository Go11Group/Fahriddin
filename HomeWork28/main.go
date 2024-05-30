package main

import (
	"fmt"
	"github.com/Go11Group/at_lesson/lesson28/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	st := postgres.NewStudentRepo(db)

	// cr := postgres.NewCourseRepo(db)

	// res,err :=  cr.GetAllCourse()
	// if err != nil{
	// 	panic(err)
	// }
	// for _,r := range res{
	// 	fmt.Println(r)
	// }

	// res,err := cr.GetByID("3bc47cc4-a7c1-4f9e-a245-979b62a69a22")
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(res)

	// err = cr.Create(model.Course{})
	// if err != nil {
	// 	panic(err)
	// }

	// err = cr.Update("13dd50f6-b6e4-42f1-a081-f56a87b08fa1",model.Course{})
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println("Update Success!!!")

	// err = cr.Delete("13dd50f6-b6e4-42f1-a081-f56a87b08fa1")
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println("Delete Success!!!")

	// res,err := st.GetAllStudents()
	// if err != nil{
	// 	panic(err)
	// }
	// for _,r := range res{
	// 	fmt.Println(r)
	// }

	// res,err := st.GetByID("f521af2b-31a4-4416-b639-a31db11decda")
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(res)

	// err = st.Create(model.User{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Information entered")


	// err = st.Update("4bd77c5e-b8c8-4fb8-b1d2-c8e895580b42",model.User{})
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println("Update Success!!!")

	// err = st.Delete("4bd77c5e-b8c8-4fb8-b1d2-c8e895580b42")
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println("Delate succes!!!")
}
