package handler

import (
	dbpostgres "apiAluno/data"
	modelos "apiAluno/model"
	"apiAluno/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsuariosRouter(r *gin.RouterGroup) {
	r.POST("", logar)
	r.POST("/novo", adicionarUsuario)
}

// adiciona um novo usuário
func adicionarUsuario(c *gin.Context) {
	var req modelos.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Senha = services.SHA256Encoder(req.Senha)

	err := dbpostgres.AdicionarUsuario(req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(201, gin.H{
		"Sucesso": "Usuário adicionado",
	})
}

// faz login e retorna token
func logar(c *gin.Context) {
	var req modelos.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := dbpostgres.BuscarUsuario(req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	if user.Senha != services.SHA256Encoder(req.Senha) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "senha incorreta"})
		return
	}

	token, err := services.NewJWTService().GenerateToken(uint(user.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "creddencial inválida"})
		return
	}
	fmt.Println(token)
	c.JSON(200, gin.H{"token": token})
}
