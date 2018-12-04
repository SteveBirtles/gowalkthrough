package models

import "fmt"

type Manufacturer struct {
	ManufacturerId int    `json:"manufacturerId"`
	Name           string `json:"name"`
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
	_, err := database.Exec("insert into Manufacturers (ManufacturerId, Name) values (?, '?')",
		m.ManufacturerId, m.Name)
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateManufacturer(m Manufacturer) {
	_, err := database.Exec("update Manufacturers set Name = '?' where ManufacturerId = ?",
		m.Name, m.ManufacturerId)
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteManufacturer(manufacturerId int) {
	_, err := database.Exec("delete from Messages where ManufacturerId = ?", manufacturerId)
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
