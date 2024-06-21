package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/modul"
	"strings"
	"time"
)
//Update qilishda filter qib ozgartiradi
type UpdateEnrollment struct {
	EnrollmentId   *string `json:"enrollment_id"`
	UserId         *string `json:"user_id"`
	CourseId       *string `json:"course_id"`
	EnrollmentDate *string `json:"enrollment_date"`
}
//funcsiyalardi bazaga method qib olish uchun
type EnrollmentsRepo struct {
	DB *sql.DB
}

func NewEnrollmentsRepo(db *sql.DB) *EnrollmentsRepo {
	return &EnrollmentsRepo{DB: db}
}

func (e *EnrollmentsRepo) Create(enrollment modul.Enrollment) error {
	tr, err := e.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	_, err = e.DB.Exec(`INSERT INTO Enrollments(user_id,course_id,enrollment_date,created_at,updated_at)
	VALUES($1,$2,$3,$4,$5)`, enrollment.UserId, enrollment.CourseId, time.Now(), time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (e *EnrollmentsRepo) Read(Id string) (modul.Enrollment, error) {
	tr, err := e.DB.Begin()
	if err != nil {
		return modul.Enrollment{},err
	}
	defer tr.Commit()

	enrollment := modul.Enrollment{}
	row := e.DB.QueryRow("Select enrollment_id,user_id,course_id,enrollment_date FROM Enrollments WHERE enrollment_id = $1 AND deleted_at = 0", Id)
	err = row.Scan(&enrollment.EnrollmentId,&enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate)
	if err != nil {
		return modul.Enrollment{}, err
	}
	return enrollment, nil
}

func (e *EnrollmentsRepo) Update(updateFilter UpdateEnrollment) (error) {
	tr, err := e.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT enrollment_id
	  FROM Enrollments
	  WHERE deleted_at = 0 AND enrollment_id = $1
	`
	err = e.DB.QueryRow(query, *updateFilter.EnrollmentId).Err()

	if err != nil {
		return err
	}

	query = `
	  UPDATE Enrollments SET 
	`

	
	if updateFilter.UserId != nil {
		params = append(params, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, *updateFilter.UserId)
	}

	if updateFilter.CourseId != nil {
		params = append(params, fmt.Sprintf("course_id = $%d", len(args)+1))
		args = append(args, *updateFilter.CourseId)
	}

	if updateFilter.EnrollmentDate != nil {
		params = append(params, fmt.Sprintf("enrollment_date = $%d", len(args)+1))
		args = append(args, *updateFilter.EnrollmentDate)
	}

	params = append(params, fmt.Sprintf("updated_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return err
	}

	args = append(args, *updateFilter.EnrollmentId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE enrollment_id = $%d AND deleted_at = 0", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err = e.DB.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (e *EnrollmentsRepo) DELETE(Id string) (error){
	tr, err := e.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	res,err := e.DB.Exec(`UPDATE Enrollments SET
		deleted_at = date_part('epoch', current_timestamp)::INT
	   where enrollment_id = $1 and deleted_at = 0`,Id)
	if err != nil{
	return err
	}
	x,err := res.RowsAffected()
	
	if err != nil{
		return err
	}

	if x == 0{
		return fmt.Errorf("no such information is available or %s not found already deleted", Id)
	}
	return nil
}


func (e *EnrollmentsRepo) GetAll(fEnrollment modul.FilterEnrollment) ([]modul.Enrollment, error) {
	tr, err := e.DB.Begin()
	if err != nil {
		return []modul.Enrollment{}, err
	}
	defer tr.Commit()
	
	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT enrollment_id,user_id,course_id,enrollment_date FROM Enrollments WHERE deleted_at = 0 "

	if fEnrollment.UserID != "" {
		params["user_id"] = fEnrollment.UserID
		filter += "AND user_id = :user_id "
	}

	if fEnrollment.CourseID != "" {
		params["course_id"] = fEnrollment.CourseID
		filter += "AND course_id = :course_id "
	}

	if fEnrollment.EnrollmentDate != nil {
		params["enrollment_date"] = fEnrollment.EnrollmentDate.Format("2006-01-02")
	}

	if fEnrollment.Limit > 0 {
		params["limit"] = fEnrollment.Limit
		filter += " limit :limit "
	}

	if fEnrollment.Offset > 0 {
		params["offset"] = fEnrollment.Offset
		filter += " offset :offset "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query,params)

	rows, err := e.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	fmt.Println(query,params)

	var enrollments []modul.Enrollment
	for rows.Next() {
		var enrollment modul.Enrollment

		err = rows.Scan(&enrollment.EnrollmentId,&enrollment.UserId,&enrollment.CourseId,&enrollment.EnrollmentDate)

		if err != nil {
			return nil, err
		}

		enrollments = append(enrollments, enrollment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return enrollments, err
}
