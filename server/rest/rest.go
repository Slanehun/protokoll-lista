package rest

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterEndpoints(r *gin.Engine, db *sql.DB) {
	r.GET("/list", list(db))
	r.POST("/add", add(db))
	r.DELETE("/event/:eventID/delete", delete(db))
}

const insertValues = "INSERT INTO magyar (government, event, name, position, address, email, phone, comment) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

func list(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM magyar")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		defer rows.Close()

		var events []Event
		for rows.Next() {
			var event Event
			err := rows.Scan(
				&event.ID,
				&event.Government,
				&event.Event,
				&event.Name,
				&event.Position,
				&event.Address,
				&event.Email,
				&event.Phone,
				&event.Comment)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			events = append(events, event)
		}

		c.JSON(http.StatusOK, events)
	}
}

func add(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var event Event

		if err := c.ShouldBindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		result, err := db.Exec(insertValues, event.Government, event.Event, event.Name, event.Position, event.Address, event.Email, event.Phone, event.Comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		event.ID = int(id)

		c.JSON(http.StatusCreated, event)
	}
}

func delete(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID, err := strconv.Atoi(c.Param("eventID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		fmt.Printf("Deleting event with the ID of %v\n", eventID)

		result, err := db.Exec("DELETE FROM magyar WHERE id=?", eventID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}

		// Return a JSON response indicating success
		c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
	}
}
