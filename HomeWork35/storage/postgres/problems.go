package postgres

import (
	"database/sql"
	"my_module/modul"
)

type ProblemRepo struct {
	DB *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{DB: db}
}

func (p *ProblemRepo) Create(problem modul.Problem) error {
	_, err := p.DB.Exec(`INSERT INTO Problems(Problem_num,Title,Status,Description)
	($1,$2,$3,$4)`,
		problem.Problem_num, problem.Title, problem.Status, problem.Description)

	if err != nil {
		return err
	}

	return nil
}
func (p *ProblemRepo) GetById(Id string) (modul.Problem, error) {

	row := p.DB.QueryRow("Select * From Problems Where id = $1", Id)

	problem := modul.Problem{}

	err := row.Scan(&problem.Id, &problem.Problem_num, &problem.Title, &problem.Status, &problem.Description)
	if err != nil {
		return modul.Problem{}, err
	}

	return problem, nil
}

func (p *ProblemRepo) Update(problem modul.Problem) error {
	_, err := p.DB.Exec(`Update Problems Set 
	Id = $1,
	Problem_num = $2,
	Title = $3,
	Status = $4,
	Description = $5
	Where id = $6
	`, problem.Problem_num, problem.Title, problem.Status, problem.Description, problem.Id)

	return err
}

func (p *ProblemRepo) Delete(Id string) error {
	_, err := p.DB.Exec("Delete From Problems Where Id = $1", Id)
	return err
}
