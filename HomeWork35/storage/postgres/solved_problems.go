package postgres

import (
	"database/sql"
	"my_module/modul"
)

type Solved_problemsRepo struct {
	DB *sql.DB
}

func NewSolved_problemsRepo(db *sql.DB) *Solved_problemsRepo {
	return &Solved_problemsRepo{DB: db}
}

func (s *Solved_problemsRepo) Create(sp modul.Solved_problem) error {
	problem := modul.Problem{}
	row := s.DB.QueryRow("Select Answer from Problems where id = $1", sp.Problem_id)
	err := row.Scan(&problem.Answer)
	answer := problem.Answer
	if err != nil{
		return err
	}
	if answer == sp.Answer_problem{
		sp.Result = true
	}
	_, err = s.DB.Exec(`INSER INTO Solved_problems(User_id,Problem_id,Answer_problem,Result) 
	VALUES($1,$2,$3,$4)`, sp.User_id, sp.Problem_id, sp.Answer_problem, sp.Result)
	if err != nil {
		return err
	}
	return err
}

func (s *Solved_problemsRepo) GetById(Id string) (modul.Solved_problem,error){
	row :=s.DB.QueryRow("Select * From Solved_problems Where id = $1",Id)
	sp := modul.Solved_problem{}
	err := row.Scan(&sp.Id,&sp.User_id,&sp.Problem_id,sp.Answer_problem,sp.Result)
	if err != nil{
		return modul.Solved_problem{},err
	}
	return sp,nil
}	


func (s *Solved_problemsRepo) Update(sp modul.Solved_problem)error{
	_,err := s.DB.Exec(`Update Solved_problems 
	Set User_id = $1,Problem_id = $2,Answer_problem = $3,Result = $4 Where id = $5`,
	sp.User_id,sp.Problem_id,sp.Answer_problem,sp.Result,sp.Id)
	if err != nil{
		return err
	}
	return nil
}

func (s *Solved_problemsRepo)Delete(Id string) error{
	_,err := s.DB.Exec("Delete From Solved_problems where id = $1",Id)
	if err != nil{
		return err
	}
	return nil
}