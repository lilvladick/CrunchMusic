package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func GetCategoriesByID(CategoryID int) (*Category, error) {
	query := "SELECT * FROM categories WHERE id = $1"

	row := db.QueryRow(query, CategoryID)
	var Category Category

	err := row.Scan(&Category.ID, &Category.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Category not found with ID %v", CategoryID)
		}
		return nil, fmt.Errorf("error scanning Category row: %v", err)
	}

	return &Category, nil
}

func GetCategoryByName(name string) (*Category, error) {
	query := "SELECT * FROM categories WHERE name = $1"

	row := db.QueryRow(query, name)
	var Category Category

	err := row.Scan(&Category.ID, &Category.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Category not found with name %v", name)
		}
		return nil, fmt.Errorf("error scanning Category row: %v", err)
	}

	return &Category, nil
}

func GetCategories() ([]Category, error) {
	sqlStatement := `
        SELECT * FROM categories;
    `
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error querying Categories: %v", err)
	}
	defer rows.Close()

	var Categories []Category
	for rows.Next() {
		var Category Category
		err := rows.Scan(&Category.ID, &Category.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning Category row: %v", err)
		}
		Categories = append(Categories, Category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over Categories rows: %v", err)
	}

	return Categories, nil
}

func InsertCategory(category Category) error {
	sqlStatement := `
        INSERT INTO categories (name) VALUES ($1)   `

	_, err := db.Exec(sqlStatement, category.Name)
	return err
}
