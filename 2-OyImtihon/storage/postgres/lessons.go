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
type UpdateLesson struct {
	LessonId  	*string     `json:"lesson_id"`
	CourseId   *string    `json:"course_id"`
	Title    	*string    	`json:"title"`
	Content 	*string 		`json:"content"`
}
//funcsiyalardi bazaga method qib olish uchun
type LesseonsRepo struct {
	DB *sql.DB
}

func NewLessonsRepo(db *sql.DB) *LesseonsRepo {
	return &LesseonsRepo{DB:db}
}

//Lesson yaratish
func (l *LesseonsRepo) Create(lesson modul.Lesson) error{
	tr, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = l.DB.Exec(`INSERT INTO Lessons(course_id,title,content,created_at,updated_at) 
	VALUES($1,$2,$3,$4,$5)`, lesson.CourseId, lesson.Title, lesson.Content, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

//Lesson Id bo'yicha o'qish
func (l *LesseonsRepo) Read(Id string) (modul.Lesson, error) {
	lesson := modul.Lesson{}
	row := l.DB.QueryRow("Select lesson_id,course_id,title,content FROM Lessons WHERE lesson_id = $1 AND deleted_at = 0", Id)
	err := row.Scan(&lesson.LessonId,&lesson.CourseId, &lesson.Title, &lesson.Content)
	if err != nil {
		return modul.Lesson{}, err
	}
	return lesson, nil
}

//Filter bo'yicha Lesson o'zgartirish 
func (l *LesseonsRepo) Update(updateFilter UpdateLesson) (error) {
	tr, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT lesson_id
	  FROM Lessons
	  WHERE deleted_at = 0 AND lesson_id = $1
	`
	err = l.DB.QueryRow(query,*updateFilter.LessonId).Err()

	if  err != nil {
		return err
	}

	query = `
	  UPDATE Lessons SET 
	`

	if updateFilter.CourseId != nil {
		params = append(params, fmt.Sprintf("course_id = $%d", len(args)+1))
		args = append(args, *updateFilter.CourseId)
	}

	if updateFilter.Title != nil {
		params = append(params, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, *updateFilter.Title)
	}

	if updateFilter.Content != nil {
		params = append(params, fmt.Sprintf("content = $%d", len(args)+1))
		args = append(args, *updateFilter.Content)
	}

	params = append(params, fmt.Sprintf("updated_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return err
	}

	args = append(args, *updateFilter.LessonId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE lesson_id = $%d AND deleted_at = 0", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err = l.DB.Exec(query, args...)

	if err != nil {
		return err
	}
	return nil
}

//Id boyicha Lesson di ochirish
func (l *LesseonsRepo) DELETE(Id string) (error){
	tr, err := l.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	res,err := l.DB.Exec(`UPDATE Lessons SET
		deleted_at = date_part('epoch', current_timestamp)::INT
	   where lesson_id = $1 and deleted_at = 0`,Id)
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


//Filter boyicha qidirib ularni chiqarish
func (l *LesseonsRepo) GetAll(fLesson modul.FilterLesson) ([]modul.Lesson, error) {
	tr, err := l.DB.Begin()
	if err != nil {
		return []modul.Lesson{},err
	}
	defer tr.Commit()

	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT lesson_id,course_id,title,content FROM Lessons WHERE deleted_at = 0 "

	if  fLesson.CourseID != "" {
		params["course_id"] = fLesson.CourseID
		filter += "AND course_id = :course_id "
	}
	if fLesson.Title != "" {
		params["title"] = fLesson.Title
		filter += "AND title = :title "
	}
	if fLesson.Content != "" {
		params["content"] = fLesson.Content
		filter += "AND content = :content"
	}

	if fLesson.Limit > 0 {
		params["limit"] = fLesson.Limit
		filter += " limit :limit "
	}

	if fLesson.Offset > 0 {
		params["offset"] = fLesson.Offset
		filter += " offset :offset "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query,params)

	rows, err := l.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var lessons []modul.Lesson
	for rows.Next() {
		var lesson modul.Lesson

		err = rows.Scan(&lesson.LessonId,&lesson.CourseId,&lesson.Title,&lesson.Content)

		if err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lessons, err
}
