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

// cria e retorna uma conexão com o bando de dados postgres
func conectar() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
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

// lista alunos contidos no bando de dados postgres
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

// busca alunos no banco de dados postgres
func BuscarAluno(alunoID int64) (aluno *modelos.EstruturaAluno, err error) {
	db := conectar()
	aluno = new(modelos.EstruturaAluno)
	if err = db.
		QueryRow(`
			SELECT
				ta.id,
				ta.nome,
				ta.idade,
				ta.matricula,
				ta.curso
			FROM t_aluno ta
			WHERE ta.id =$1`, alunoID).
		Scan(&aluno.ID, &aluno.Nome,
			&aluno.Idade, &aluno.Matricula,
			&aluno.Curso); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return
}

// adiciona alunos no bando de dados postgres
func AdicionarAluno(aluno modelos.EstruturaAluno) (err error) {
	db := conectar()
	sqlStatement := `INSERT INTO public.t_aluno
	(nome, matricula, idade, curso)
	VALUES($1::TEXT, $2::TEXT, $3::BIGINT, $4::TEXT)`
	_, err = db.Exec(sqlStatement, aluno.Nome, aluno.Matricula, aluno.Idade, aluno.Curso)
	if err != nil {
		return err
	}

	return
}

// deleta alunos no banco de dados postgres
func DeletarAluno(alunoID int64) (err error) {
	db := conectar()
	sqlStatement := `DELETE FROM public.t_aluno
	WHERE id=$1::BIGINT`
	_, err = db.Exec(sqlStatement, alunoID)
	if err != nil {
		return err
	}

	return
}

// edita alunos no banco de dados postgres
func EditarAluno(aluno modelos.EstruturaAluno) (err error) {
	db := conectar()
	sqlStatement := `UPDATE public.t_aluno
	SET nome=$2::TEXT, matricula=$3::TEXT, idade=$4::BIGINT, curso=$5::TEXT
	WHERE id=$1::BIGINT`
	_, err = db.Exec(sqlStatement, aluno.ID, aluno.Nome, aluno.Matricula, aluno.Idade, aluno.Curso)
	if err != nil {
		return err
	}

	return
}
