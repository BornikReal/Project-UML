package film

type filmModel struct {
	ID          int64    `db:"id"`
	Name        string   `db:"name"`
	Description string   `db:"description"`
	Cover       string   `db:"cover"`
	Genres      []string `db:"genres"`
	Actors      []string `db:"actors"`
	Images      []string `db:"images"`
	Trailers    []string `db:"trailers"`
	Ratings     []int64  `db:"ratings"`
}
