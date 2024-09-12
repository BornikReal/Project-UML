package rating

type ratingModel struct {
	ID        int64   `db:"id"`
	Score     float32 `db:"score"`
	Review    string  `db:"review"`
	FilmID    int64   `db:"film_id"`
	UserID    int64   `db:"user_id"`
	IsSpecial bool    `db:"is_special"`
}
