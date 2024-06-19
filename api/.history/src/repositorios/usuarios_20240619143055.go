package repositorios

type usuarios struct {  // usuarios com "u" minusculo porque não será exportada
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql)