package main

type Movie struct {
	ID        string    `json:"id"`
	Isbn      string    `json:"isbn"`
	Title     string    `json:"title"`
	Directors *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Movies []Movie
