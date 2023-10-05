package controllers

import (
	"net/http"

	"github.com/Jonathan0012/backend-test-ottodigital/models"
	"github.com/gin-gonic/gin"
)

func IndexTasks(c *gin.Context) {
	tasks := []models.Tasks{}
	var t models.Tasks

	results, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	for results.Next() {
		err = results.Scan(&t.Id, &t.User_Id, &t.Title, &t.Description, &t.Status, &t.Created_at, &t.Updated_at)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

			return
		}
		tasks = append(tasks, t)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Tasks": tasks})
}

func CreateTask(c *gin.Context) {
	var t models.Tasks
	if err := c.ShouldBindJSON(&t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := db.Query("SELECT * FROM users WHERE id=?", t.User_Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	insert, err := db.Query("INSERT INTO tasks (user_id,title,description,status,created_at,updated_at) VALUES (?,?,?,?, now(),now())", t.User_Id, t.Title, t.Description, t.Status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	defer insert.Close()
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Task created"})
}

func ShowTask(c *gin.Context) {
	id := c.Param("id")
	var t models.Tasks

	result := db.QueryRow("SELECT * FROM tasks WHERE id=?", id)

	err := result.Scan(&t.Id, &t.User_Id, &t.Title, &t.Description, &t.Status, &t.Created_at, &t.Updated_at)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Tasks": t})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var t models.Tasks
	if err := c.ShouldBindJSON(&t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := db.Query("SELECT * FROM users WHERE id=?", t.User_Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	update, err := db.Query("UPDATE tasks SET user_id=?,title=?,description=?,status=?,updated_at=now() WHERE id=?", t.User_Id, t.Title, t.Description, t.Status, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if update.Next() {
		err = update.Scan(&t.Id, &t.User_Id, &t.Title, &t.Description, &t.Status, &t.Created_at, &t.Updated_at)
		if err != nil {
			panic(err.Error())
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Succesfully Updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var t models.Tasks

	result := db.QueryRow("SELECT * FROM tasks WHERE id=?", id)
	err := result.Scan(&t.Id, &t.User_Id, &t.Title, &t.Description, &t.Status, &t.Created_at, &t.Updated_at)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	delete, err := db.Query("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	defer delete.Close()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Succesfully Deleted"})
}
