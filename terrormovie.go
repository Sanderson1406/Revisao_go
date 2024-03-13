package main

type TerrorMovie struct{
	Movie
}

func (TerrorMovie) string() (string) {
	return "Terror"
}

func (TerrorMovie) CollectRatings() (float32, error) {
	return 10, nil
}