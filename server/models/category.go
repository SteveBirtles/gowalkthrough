package models

import "fmt"

type Category struct {
	CategoryId int `json:"categoryId"`
	Name string `json:"name"`
}

func SelectAllCategorys() []Category {
	categorys := make([]Category, 0)
	rows, err := database.Query("select CategoryId, Name from Categories")
	if err != nil {
		fmt.Println("Database select all error:", err)
		return categorys
	}
	defer rows.Close()
	for rows.Next() {
		var c Category
		err = rows.Scan(&c.CategoryId, &c.Name)
		if err != nil {
			fmt.Println("Database select all error:", err)
			break
		}
		categorys = append(categorys, c)
	}
	return categorys
}

func SelectCategory(categoryId int) Category {
	var c Category
	rows, err := database.Query("select CategoryId, Name from Categories where CategoryId = ?", categoryId)
	if err != nil {
		fmt.Println("Database select error:", err)
		return c
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&c.CategoryId, &c.Name)
	if err != nil {
		fmt.Println("Database select error:", err)
	}
	return c
}

func InsertCategory(c Category) {
	statement, err := database.Prepare("insert into Categories (CategoryId, Name) values (?, ?)")
	if err == nil { _, err = statement.Exec(c.CategoryId, c.Name) }
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateCategory(c Category) {
	statement, err := database.Prepare("update Categories set Name = ? where CategoryId = ?")
	if err == nil { _, err = statement.Exec(c.Name, c.CategoryId) }
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteCategory(categoryId int) {
	statement, err := database.Prepare("delete from Categories where CategoryId = ?")
	if err == nil { _, err = statement.Exec(categoryId) }
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
