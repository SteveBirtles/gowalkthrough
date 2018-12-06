package models

import "fmt"

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	rows, err := database.Query("select Username, Password, SessionToken from Admins where Username = ?", username)
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
	statement, err := database.Prepare("insert into Admins (Username, Password, SessionToken) values (?, ?, ?)")
	if err == nil { _, err = statement.Exec(a.Username, a.Password, a.SessionToken) }
	if err != nil {
		fmt.Println("Database insert error:", err)
	}
}

func UpdateAdmin(a Admin) {
	statement, err := database.Prepare("update Admins set Password = ?, SessionToken = ? where Username = ?")
	if err == nil { _, err = statement.Exec(a.Password, a.SessionToken, a.Username) }
	if err != nil {
		fmt.Println("Database update error:", err)
	}
}

func DeleteAdmin(username string) {
	statement, err := database.Prepare("delete from Admins where Username = ?")
	if err == nil { _, err = statement.Exec(username) }
	if err != nil {
		fmt.Println("Database delete error:", err)
	}
}
