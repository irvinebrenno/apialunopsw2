package modelos

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
