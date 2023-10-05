package controllers

import (
	"net/http"

	"github.com/Jonathan0012/backend-test-ottodigital/app"
	"github.com/Jonathan0012/backend-test-ottodigital/models"
	"github.com/gin-gonic/gin"
)

var db = app.NewDB()

func IndexUsers(c *gin.Context) {
	users := []models.Users{}
	var u models.Users

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	for results.Next() {
		err = results.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Created_at, &u.Updated_at)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, u)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Users": users})
}

func CreateUser(c *gin.Context) {
	var u models.Users
	if err := c.ShouldBindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	insert, err := db.Query("INSERT INTO users (name,email,password,created_at,updated_at) VALUES (?,?,?, now(),now())", u.Name, u.Email, u.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	defer insert.Close()
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "User created"})
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")
	var u models.Users

	result := db.QueryRow("SELECT * FROM users WHERE id=?", id)

	err := result.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Created_at, &u.Updated_at)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Users": u})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u models.Users
	if err := c.ShouldBindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	update, err := db.Query("UPDATE users SET name=?,email=?,password=?,updated_at=now() WHERE id=?", u.Name, u.Email, u.Password, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if update.Next() {
		err = update.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Created_at, &u.Updated_at)
		if err != nil {
			panic(err.Error())
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User Succesfully Updated"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var u models.Users

	result := db.QueryRow("SELECT * FROM users WHERE id=?", id)
	err := result.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Created_at, &u.Updated_at)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	delete, err := db.Query("DELETE FROM users WHERE id=?", id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	defer delete.Close()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User Succesfully Deleted"})
}
