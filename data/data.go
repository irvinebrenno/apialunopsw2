package dbpostgres

import (
	modelos "apiAluno/model"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-44-194-92-192.compute-1.amazonaws.com"
	port     = 5432
	user     = "cbwddrgknyovbl"
	password = "d6e2c7033b945c6e0717a66337a9ce6d5229d57327edc5751b2e367a37161e1c"
	dbname   = "d6n9egsujidc12"
)

func conectar() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func ListarAlunos() (alunos []modelos.EstruturaAluno) {
	db := conectar()
	var aluno modelos.EstruturaAluno

	rows, err := db.
		Query(`
			SELECT
				ta.id,
				ta.nome,
				ta.idade,
				ta.matricula,
				ta.curso
			FROM t_aluno ta`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&aluno.ID, &aluno.Nome, &aluno.Idade, &aluno.Matricula, &aluno.Curso)
		alunos = append(alunos, aluno)
	}

	return
}

type Tolete struct {
	ab int64
}
