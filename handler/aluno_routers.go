package handler

import (
	dbpostgres "apiAluno/data"
	modelos "apiAluno/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// rotas para aluno
func Router(r *gin.RouterGroup) {
	r.GET("", listarAlunos)
	r.POST("", adicionarAluno)
	r.GET("/:aluno_id", buscarAluno)
	r.PUT("/:aluno_id", editarAluno)
	r.DELETE("/:aluno_id", deletarAluno)
}

// handler para listagem de alunos
func listarAlunos(c *gin.Context) {
	alunos, err := dbpostgres.ListarAlunos()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, alunos)
}

// handler para buscar aluno pelo id
func buscarAluno(c *gin.Context) {
	alunoID, err := strconv.ParseInt(c.Param("aluno_id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	aluno, err := dbpostgres.BuscarAluno(alunoID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, aluno)
}

// handler para adicionar um aluno
func adicionarAluno(c *gin.Context) {
	var req modelos.EstruturaAluno

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := dbpostgres.AdicionarAluno(req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(201, nil)
}

// handler para editar um aluno
func editarAluno(c *gin.Context) {
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
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(201, nil)
}

// handler para deletar um aluno
func deletarAluno(c *gin.Context) {
	alunoID, err := strconv.ParseInt(c.Param("aluno_id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	err = dbpostgres.DeletarAluno(alunoID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(201, nil)
}
