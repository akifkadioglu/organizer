package run

import (
	"encoding/json"
	"os"

	"github.com/akifkadioglu/organizer/database"
	"github.com/akifkadioglu/organizer/database/entities"
	"github.com/akifkadioglu/organizer/olog"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GetPassword is a function to get apps and names password
func GetPassword(app, username string) {
	var passwordEntity []entities.Password
	tx := database.Get()
	if app != "" {
		tx = tx.Where("app = ?", app)

	}
	if username != "" {
		tx = tx.Where("username = ?", username)
	}
	err := tx.Find(&passwordEntity).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			olog.Fatal("There is no password with this username")
		} else {
			olog.Fatal("Error while reaching password")
		}
	}
	data := [][]string{}
	for _, v := range passwordEntity {
		data = append(data, []string{v.App, v.Username, v.Password})
	}
	header := []string{"App", "Username", "Password"}
	logPasswordTable(header, data)
}

// PasswordDownload is a function to download passwords.json
func PasswordDownload() {
	var passwordEntity []entities.Password
	var err error
	err = database.Get().Find(&passwordEntity).Error
	if err != nil {
		olog.Fatal("Error while reaching passwords")
	}
	file, err := os.Create("passwords.json")

	if err != nil {
		olog.Fatal("Error while creating passwords.json")
	}

	defer file.Close()
	a, err := json.Marshal(passwordEntity)
	if err != nil {
		olog.Fatal("Error while marshalling passwords.")
	}
	n := len(a)
	s := string(a[:n])

	file.WriteString(s)

	olog.Info("Passwords downloaded as passwords.json")
}

// PasswordGet is a function to get passwords
func GeneratePassword(app, username string) {
	var passwordEntity entities.Password
	passwordEntity.App = app
	passwordEntity.Username = username
	hash, err := bcrypt.GenerateFromPassword([]byte(uuid.New().String()), bcrypt.DefaultCost)
	if err != nil {
		olog.Fatal("Error while generating password")
	}
	passwordEntity.Password = string(hash)

	err = database.Get().
		Where("app = ?", app).
		Where("username = ?", username).
		Save(&passwordEntity).Error

	if err != nil {
		olog.Fatal("Error while creating password")
	}
	data := [][]string{{passwordEntity.App, passwordEntity.Username, passwordEntity.Password}}
	header := []string{"App", "Username", "Password"}
	logPasswordTable(header, data)
}

// RemovePassword is a function to remove password
func RemovePassword(app, username string) {
	var passwordEntity entities.Password
	database.Get().
		Where("app = ?", app).
		Where("username = ?", username).
		Delete(&passwordEntity)

	olog.Info(username + ": Password removed successfully")

}

// ListPassword is a function to list passwords
func ListPassword(app string) {
	var passwordEntity []entities.Password
	database.Get().Find(&passwordEntity)
	data := [][]string{}
	for _, v := range passwordEntity {
		data = append(data, []string{v.App, v.Username, v.Password})
	}

	header := []string{"App", "Username", "Password"}
	logPasswordTable(header, data)
}
