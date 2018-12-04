package models

import "fmt"

type Accessory struct {
	AccessoryId int    `json:"id"`
	CategoryId  int    `json:"categoryId"`
	ConsoleId   int    `json:"consoleId"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	ThirdParty  bool   `json:"thirdParty"`
	ImageURL    string `json:"imageURL"`
}

func SelectAllAccessories() []Accessory {
	accessorys := make([]Accessory, 0)
	rows, err := database.Query("select AccessoryId, CategoryId, ConsoleId, Description, Quantity, ThirdParty, ImageURL from Accessories")
	if err != nil {
		fmt.Println("Database select all error:", err)
		return accessorys
	}
	defer rows.Close()
	for rows.Next() {
		var a Accessory
		err = rows.Scan(&a.AccessoryId, &a.CategoryId, &a.ConsoleId, &a.Description, &a.Quantity, &a.ThirdParty, &a.ImageURL)
		if err != nil {
			fmt.Println("Database select all error:", err)
			break
		}
		accessorys = append(accessorys, a)
	}
	return accessorys
}

func SelectAccessory(accessoryId int) Accessory {
	var a Accessory
	rows, err := database.Query("select AccessoryId, CategoryId, ConsoleId, Description, Quantity, ThirdParty, ImageURL from Accessories where AccessoryId = ?", accessoryId)
	if err != nil {
		fmt.Println("Database select error:", err)
		return a
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&a.AccessoryId, &a.CategoryId, &a.ConsoleId, &a.Description, &a.Quantity, &a.ThirdParty, &a.ImageURL)
	if err != nil {
		fmt.Println("Database select error:", err)
	}
	return a
}

func InsertAccessory(a Accessory) {
	_, err := database.Exec("insert into Accessories (AccessoryId, CategoryId, ConsoleId, Description, Quantity, ThirdParty, ImageURL) values (?, ?, ?, '?', ?, ?, '?')",
		a.AccessoryId, a.CategoryId, a.ConsoleId, a.Description, a.Quantity, a.ThirdParty, a.ImageURL)
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateAccessory(a Accessory) {
	_, err := database.Exec("update Accessories set CategoryId = ?, ConsoleId = ?, Description = '?', Quantity = ?, ThirdParty = ?, ImageURL = '?' where AccessoryId = ?",
		a.CategoryId, a.ConsoleId, a.Description, a.Quantity, a.ThirdParty, a.ImageURL, a.AccessoryId)
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteAccessory(accessoryId int) {
	_, err := database.Exec("delete from Messages where AccessoryId = ?", accessoryId)
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
