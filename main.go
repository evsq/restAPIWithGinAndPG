package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type movie struct {
	ID     int    `json:"id"`
	Rating string `json:"rating"`
	Name   string `json:"name"`
}

func main() {

	connect := "user=admin dbname=test password=test sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/movie", func(c *gin.Context) {
		var movie movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
			return
		}

		sql := fmt.Sprintf(
			"insert into movie (rating, name) values ('%s', '%s') returning id",
			movie.Rating,
			movie.Name)

		var id int
		err := db.QueryRow(sql).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
			return
		}
		movie.ID = id
		c.JSON(http.StatusOK, movie)
	})

	r.GET("/movie", func(c *gin.Context) {
		rows, err := db.Query("select id, rating, name from movie")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
			return
		}
		defer rows.Close()

		movies := []movie{}
		var id int
		var rating, name string
		for rows.Next() {
			if err := rows.Scan(&id, &rating, &name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
				return
			}
			movies = append(movies, movie{id, rating, name})
		}
		c.JSON(http.StatusOK, movies)
	})

	r.PUT("/movie/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"id must be int:": err.Error()})
			return
		}
		var movie movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
			return
		}

		sql := fmt.Sprintf(
			"update movie set rating='%s', name='%s' where id=%d",
			movie.Rating,
			movie.Name,
			id)
		if _, err := db.Query(sql); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	r.DELETE("/movie/:id", func(c *gin.Context) {
		id, e := strconv.Atoi(c.Param("id"))
		if e != nil {
			c.JSON(http.StatusBadRequest, gin.H{"id must be int:": err.Error()})
			return
		}
		sql := fmt.Sprintf("delete from movie where id=%d", id)
		if _, err := db.Query(sql); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	r.Run(":8080")
}
