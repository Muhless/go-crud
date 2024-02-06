package categoriesmodel

import (
	"github.com/Muhless/go-crud/config"
	"github.com/Muhless/go-crud/entities"
)

func GetALL() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	//looping rows di baris ke-9
	for rows.Next() {
		var category entities.Category
		if rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
	INSERT INTO categories (name, created_at, updated_at)
	VALUE (?, ?, ?)`,
		category.Name, category.CreatedAt, category.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertID > 0
}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id,name FROM categories WHERE id=?`, id)

	var categories entities.Category
	if err := row.Scan(&categories.Id, &categories.Name); err != nil {
		panic(err.Error())
	}

	return categories
}

func Update(id int, categories entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name=?, updated_at=? WHERE id=?`, categories.Name, categories.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}
