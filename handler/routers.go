package aluno

import (
	dbpostgres "apiAluno/data"
	modelos "apiAluno/model"
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
	alunos := dbpostgres.ListarAlunos()
	c.JSON(200, alunos)
}

func buscarAluno(c *gin.Context) {
	alunoID, err := strconv.ParseInt(c.Param("aluno_id"), 10, 64)
	if err != nil {
		return
	}
	var res modelos.EstruturaAluno
	res.ID = alunoID
	c.JSON(200, res)
}

func logar(c *gin.Context) {
	var json modelos.Login
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
	var req modelos.EstruturaAluno
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
