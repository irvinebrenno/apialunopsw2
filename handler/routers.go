package aluno

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("", listarAlunos)
	r.POST("", adicionarAluno)
	r.POST("/logar", logar)
	r.GET("/:aluno_id", buscarAluno)
	r.PUT("/:aluno_id", editarAluno)
	r.DELETE("/:aluno_id", deletarAluno)
}

func listarAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "aluno 1",
	})
}

func buscarAluno(c *gin.Context) {
	alunoID, err := strconv.ParseInt(c.Param("aluno_id"), 10, 64)
	if err != nil {
		return
	}
	var res EstruturaAluno
	res.ID = alunoID
	c.JSON(200, res)
}

func logar(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "manu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	fmt.Println(json)

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func adicionarAluno(c *gin.Context) {
	var req EstruturaAluno
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, req)
}

func editarAluno(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "aluno 1",
	})
}

func deletarAluno(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "aluno 1",
	})
}

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EstruturaAluno struct {
	ID        int64  `json:"id" binding:"required"`
	Nome      string `json:"nome" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
	Idade     int64  `json:"idade" binding:"required"`
	Curso     string `json:"curso" binding:"required"`
}
