package modelos

type User struct {
	ID       int64
	Senha    string `json:"password" validate:"required"`
	Username string `json:"username"`
}

type EstruturaAluno struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
	Idade     int64  `json:"idade" binding:"required"`
	Curso     string `json:"curso" binding:"required"`
}
