package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"shuTeacher/infrastructure"
	"shuTeacher/service/token"
)

func main() {
	password := os.Args[1]
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(password), -1)
	_, err := infrastructure.DB.Exec(`
	INSERT INTO token(tokenhash) VALUES ($1);
	`, string(encrypted))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(token.GenerateJWT(password))
}
