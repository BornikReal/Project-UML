package film

type Film struct {
	ID          int64
	Name        string
	Description string
	Cover       string
	Genres      []string
	Actors      []string
	Images      []string
	Trailers    []string
	Ratings     []int64
}

func (f *Film) fromModel(model filmModel) {
	if f == nil {
		return
	}
	*f = Film{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Cover:       model.Cover,
		Genres:      model.Genres,
		Actors:      model.Actors,
		Images:      model.Images,
		Trailers:    model.Trailers,
		Ratings:     model.Ratings,
	}
}
