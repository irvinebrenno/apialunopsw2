package main

import (
	handler "apiAluno/handler"
	"apiAluno/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	apiAlunos := r.Group("alunos", middlewares.Auth())
	contas := r.Group("")
	handler.Router(apiAlunos)
	handler.UsuariosRouter(contas)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
