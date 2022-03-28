package dbpostgres

import (
	modelos "apiAluno/model"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-54-173-77-184.compute-1.amazonaws.com"
	port     = 5432
	user     = "aaiiammvzpolvz"
	password = "afea01c8f848df0591ddef549a0f8b0410acac6655ced6e17780c3441fcab550"
	dbname   = "de50i5dsk573pe"
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
func ListarAlunos() (alunos []modelos.EstruturaAluno, err error) {
	db := conectar()
	defer db.Close()
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
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&aluno.ID, &aluno.Nome, &aluno.Idade, &aluno.Matricula, &aluno.Curso)
		if err != nil {
			return nil, err
		}
		alunos = append(alunos, aluno)
	}

	return
}

// busca alunos no banco de dados postgres
func BuscarAluno(alunoID int64) (aluno *modelos.EstruturaAluno, err error) {
	db := conectar()
	defer db.Close()
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
		return nil, err
	}

	return
}

// adiciona alunos no bando de dados postgres
func AdicionarAluno(aluno modelos.EstruturaAluno) (err error) {
	db := conectar()
	defer db.Close()
	sqlStatement := `INSERT INTO public.t_aluno
	(nome, matricula, idade, curso)
	VALUES($1::TEXT, $2::TEXT, $3::BIGINT, $4::TEXT)`
	_, err = db.Exec(sqlStatement, aluno.Nome, aluno.Matricula, aluno.Idade, aluno.Curso)
	if err != nil {
		fmt.Println("erro aqui: ", err)
		return err
	}

	return
}

// deleta alunos no banco de dados postgres
func DeletarAluno(alunoID int64) (err error) {
	db := conectar()
	defer db.Close()
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

// adiciona usuários no bando de dados postgres
func AdicionarUsuario(user modelos.User) (err error) {
	db := conectar()
	defer db.Close()
	sqlStatement := `INSERT INTO public.t_usuarios
	(usuario, senha)
	VALUES($1, $2)`
	_, err = db.Exec(sqlStatement, user.Username, user.Senha)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

// obtém senha criptografada do usuário no banco
func BuscarUsuario(userBusca modelos.User) (user *modelos.User, err error) {
	db := conectar()
	defer db.Close()
	user = new(modelos.User)
	if err = db.
		QueryRow(`
			SELECT
				tu.id,
				tu.usuario,
				tu.senha
			FROM t_usuarios tu
			WHERE tu.usuario =$1`, userBusca.Username).
		Scan(&user.ID, &user.Username, &user.Senha); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	fmt.Println(user)

	return
}
