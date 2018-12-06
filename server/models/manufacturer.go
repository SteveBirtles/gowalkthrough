package models

import "fmt"

type Manufacturer struct {
	ManufacturerId int `json:"manufacturerId"`
	Name string `json:"name"`
}

func SelectAllManufacturers() []Manufacturer {
	manufacturers := make([]Manufacturer, 0)
	rows, err := database.Query("select ManufacturerId, Name from Manufacturers")
	if err != nil {
		fmt.Println("Database select all error:", err)
		return manufacturers
	}
	defer rows.Close()
	for rows.Next() {
		var m Manufacturer
		err = rows.Scan(&m.ManufacturerId, &m.Name)
		if err != nil {
			fmt.Println("Database select all error:", err)
			break
		}
		manufacturers = append(manufacturers, m)
	}
	return manufacturers
}

func SelectManufacturer(manufacturerId int) Manufacturer {
	var m Manufacturer
	rows, err := database.Query("select ManufacturerId, Name from Manufacturers where ManufacturerId = ?", manufacturerId)
	if err != nil {
		fmt.Println("Database select error:", err)
		return m
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&m.ManufacturerId, &m.Name)
	if err != nil {
		fmt.Println("Database select error:", err)
	}
	return m
}

func InsertManufacturer(m Manufacturer) {
	statement, err := database.Prepare("insert into Manufacturers (ManufacturerId, Name) values (?, ?)")
	if err == nil { _, err = statement.Exec(m.ManufacturerId, m.Name) }
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateManufacturer(m Manufacturer) {
	statement, err := database.Prepare("update Manufacturers set Name = ? where ManufacturerId = ?")
	if err == nil { _, err = statement.Exec(m.Name, m.ManufacturerId) }
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteManufacturer(manufacturerId int) {
	statement, err := database.Prepare("delete from Manufacturers where ManufacturerId = ?")
	if err == nil { _, err = statement.Exec(manufacturerId) }
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
