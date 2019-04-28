package model

import "shuTeacher/infrastructure"

type Teacher struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Get(id string) (Teacher, error) {
	result := Teacher{
		Id: id,
	}
	row := infrastructure.DB.QueryRow(`
	SELECT name 
	FROM teacher
	WHERE id=$1;
	`, id)
	err := row.Scan(&result.Name)
	return result, err
}

func Save(teacher Teacher) {
	_, _ = infrastructure.DB.Exec(`
	INSERT INTO teacher(id, name)
	VALUES ($1,$2);
	`, teacher.Id, teacher.Name)
}
