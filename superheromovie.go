package main

type superheroMovie struct{
	Movie
}

func (superheroMovie) string() (string) {
	return "Super Hero"
}

func (superheroMovie) CollectRatings() (float32, error) {
	return 8, nil
}