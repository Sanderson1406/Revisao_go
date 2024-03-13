package main

type cultMovie struct {
	Movie
}

func (cultMovie) string() (string) {
	return "Cult"
}

func (cultMovie) CollectRatings() (float32, error) {
	return 5 , nil
}
