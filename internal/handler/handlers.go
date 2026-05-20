package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/filipbabicdev/finance-tracker-api/internal/config"
	"github.com/filipbabicdev/finance-tracker-api/internal/model"
	"github.com/gin-gonic/gin"
)

func ReadTransactions(c *gin.Context) {
	rows, err := config.DB.Query("SELECT * FROM transactions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction
		var dateStr string
		err = rows.Scan(&t.ID, &t.Amount, &t.Category, &dateStr, &t.Type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		t.Date, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Date parsing error from DB"})
		}
		transactions = append(transactions, t)
	}

	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	var t model.Transaction

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO transactions (amount, category, date, type) VALUES (?, ?, ?, ?)", t.Amount, t.Category, t.Date.Format("2006-01-02 15:04:05"), t.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created"})
}

func UpdateTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var t model.Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("UPDATE transactions SET amount=?, category=?, date=?, type=? WHERE id=?", t.Amount, t.Category, t.Date.Format("2006-01-02 15:04:05"), t.Type, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated"})
}

func DeleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := config.DB.Exec("DELETE FROM transactions WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tranasction deleted"})
}
