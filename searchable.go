package main

type Searcheble interface {
	CollectRatings() (float32, error)
}