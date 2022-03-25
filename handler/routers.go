package aluno

import (
	dbpostgres "apiAluno/data"
	modelos "apiAluno/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// rotas para aluno
func Router(r *gin.RouterGroup) {
	r.GET("", listarAlunos)
	r.POST("", adicionarAluno)
	r.POST("/logar", logar)
	r.GET("/:aluno_id", buscarAluno)
	r.PUT("/:aluno_id", editarAluno)
	r.DELETE("/:aluno_id", deletarAluno)
}

// handler para listagem de alunos
func listarAlunos(c *gin.Context) {
	alunos := dbpostgres.ListarAlunos()
	c.JSON(200, alunos)
}

// handler para buscar aluno pelo id
func buscarAluno(c *gin.Context) {
	statusHttp := http.StatusOK

	alunoID, err := strconv.ParseInt(c.Param("aluno_id"), 10, 64)
	if err != nil {
		return
	}

	aluno, err := dbpostgres.BuscarAluno(alunoID)
	if err != nil || aluno == nil {
		statusHttp = http.StatusBadRequest
	}
	c.JSON(statusHttp, aluno)
}

// handler para login
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

// handler para adicionar um aluno
func adicionarAluno(c *gin.Context) {
	statusHttp := http.StatusNoContent
	var req modelos.EstruturaAluno

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := dbpostgres.AdicionarAluno(req)
	if err != nil {
		statusHttp = http.StatusBadRequest
	}

	c.JSON(statusHttp, nil)
}

// handler para editar um aluno
func editarAluno(c *gin.Context) {
	statusHttp := http.StatusNoContent
	var req modelos.EstruturaAluno

	alunoID, err := strconv.ParseInt(c.Param("aluno_id"), 10, 64)
	if err != nil {
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = alunoID

	err = dbpostgres.EditarAluno(req)
	if err != nil {
		statusHttp = http.StatusBadRequest
	}

	c.JSON(statusHttp, nil)
}

// handler para deletar um aluno
func deletarAluno(c *gin.Context) {
	statusHttp := http.StatusNoContent
	alunoID, err := strconv.ParseInt(c.Param("aluno_id"), 10, 64)
	if err != nil {
		return
	}

	err = dbpostgres.DeletarAluno(alunoID)
	if err != nil {
		statusHttp = http.StatusBadRequest
	}

	c.JSON(statusHttp, nil)
}
