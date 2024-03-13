package main

type defaultMovie struct{
	Movie
}

func (defaultMovie) string() (string) {
	return "Default"
}

func (defaultMovie) CollectRatings() (float32, error) {
	return 3, nil
}