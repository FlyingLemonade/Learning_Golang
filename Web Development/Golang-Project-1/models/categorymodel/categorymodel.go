package categorymodel

import (
	"Golang-Project-1/config"
	"Golang-Project-1/entities"
)


func GetAll() []entities.Category{
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil{
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next(){
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil{
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories

}

func Create(category entities.Category) bool{
	result, err := config.DB.Exec(`
	INSERT INTO categories (name,created_at,updated_at)
	VALUES (?,?,?)
	`,
	category.Name, category.CreatedAt, category.UpdatedAt)

	if err != nil{
		panic(err)
	}

	lastInsertId,err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	return lastInsertId > 0
}

func Edit(category entities.Category) bool{
	_, err := config.DB.Exec(`
	UPDATE categories 
	SET name = ?
	WHERE id = ?
	`,category.Name, category.Id)

	if err != nil {
		panic(err)
	}

	return true
}

func Delete(category entities.Category) bool{
	_, err := config.DB.Exec(`
	DELETE FROM categories 
	WHERE id = ?
	`,category.Id)

	if err != nil {
		panic(err)
	}

	return true
}