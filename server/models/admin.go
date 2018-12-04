package models

import "fmt"

type Admin struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	SessionToken string `json:"sessionToken"`
}

func SelectAllAdmins() []Admin {
	admins := make([]Admin, 0)
	rows, err := database.Query("select Username, Password, SessionToken from Admins")
	if err != nil {
		fmt.Println("Database select all error:", err)
		return admins
	}
	defer rows.Close()
	for rows.Next() {
		var a Admin
		err = rows.Scan(&a.Username, &a.Password, &a.SessionToken)
		if err != nil {
			fmt.Println("Database select all error:", err)
			break
		}
		admins = append(admins, a)
	}
	return admins
}

func SelectAdmin(username string) Admin {
	var a Admin
	rows, err := database.Query("select Username, Password, SessionToken from Admins where Username = '?'", username)
	if err != nil {
		fmt.Println("Database select error:", err)
		return a
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&a.Username, &a.Password, &a.SessionToken)
	if err != nil {
		fmt.Println("Database select error:", err)
	}
	return a
}

func InsertAdmin(a Admin) {
	_, err := database.Exec("insert into Admins (Username, Password, SessionToken) values ('?', '?', '?')",
		a.Username, a.Password, a.SessionToken)
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateAdmin(a Admin) {
	_, err := database.Exec("update Admins set Password = '?', SessionToken = '?' where Username = '?'",
		a.Password, a.SessionToken, a.Username)
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteAdmin(username string) {
	_, err := database.Exec("delete from Messages where Username = '?'", username)
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
