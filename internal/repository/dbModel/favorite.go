package dbModel

type Favorite struct {
	Login  string `db:"login"`
	Name   string `db:"name"`
	Author string `db:"author"`
	Genre  string `db:"genre"`
}
