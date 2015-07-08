package main

import (
	"fmt"
	"math"
)

var critics =  map[string]map[string]float64{
				"Lisa Rose": map[string]float64{
					"Lady in the Water" : 2.5,
					"Snakes on a Plane" : 3.5,
					"Just My Luck" : 3.0,
					"Superman Returns" : 3.5,
					"You, Me and Dupree" : 2.5,
					"The Night Listener" :  3.0},
				"Gene Seymour" : map[string]float64{
					"Lady in the Water" : 3.0,
					"Snakes on a Plane" : 3.5,
					"Just My Luck" : 1.5,
					"Superman Returns" : 5.0,
					"You, Me and Dupree" : 3.0,
					"The Night Listener" :  3.5},
				"Michael Phillips" : map[string]float64{
					"Lady in the Water" : 2.5,
					"Snakes on a Plane" : 3.0,
					"Superman Returns" : 3.5,
					"The Night Listener" :  4.0},
				"Claudia Puig" : map[string]float64{
					"Snakes on a Plane" : 3.5,
					"Just My Luck" : 3.0,
					"The Night Listener" :  4.5,
					"Superman Returns" : 4.0,
					"You, Me and Dupree" : 2.5},
				"Mick LaSalle" : map[string]float64{
					"Lady in the Water" : 3.0,
					"Snakes on a Plane" : 4.0,
					"Just My Luck" : 2.0,
					"Superman Returns" : 3.0,
					"The Night Listener" :  3.0,
					"You, Me and Dupree" : 2.0},
				"Jack Matthews" : map[string]float64{
					"Lady in the Water" : 3.0,
					"Snakes on a Plane" : 4.0,
					"The Night Listener" : 3.0,
					"Superman Returns" : 5.0,
					"You, Me and Dupree" : 3.5},
				"Toby" : map[string]float64{
					"Snakes on a Plane" : 4.5,
					"You, Me and Dupree" : 1.0,
					"Superman Returns" : 4.0}}

func main() {
	Euclidean("Toby", "Jack Matthews", critics)
}

// Euclidean Distance Score, to calculate the similarity between two users
func Euclidean(user1 string, user2 string, critics map[string]map[string]float64) float64 {
	var commonMovies = []string{}

	user1Ratings := critics[user1]
	user2Ratings := critics[user2]

	// return false if no movies in any of the two users
	if len(user1Ratings) == 0 || len(user2Ratings) == 0 {
		return 0
	}

	// check for similar movies and add them to a seperate array
	for key, _ := range user1Ratings {
		if user2Ratings[key] > 0 {
			commonMovies = append(commonMovies, key)
		}
	}

	// calculate the first portion of the formula
	// square root of the sum of the difference between ratings of common movies
	var numerator float64 = 0
	for _, movieName := range commonMovies {
		numerator = numerator + math.Pow(user1Ratings[movieName] - user2Ratings[movieName], 2)
	}

	// square root the numenator and divide one by it to get a higher ranking for similarity
	result := 1 / 1+(math.Sqrt(numerator))

	fmt.Println(result)

	return result
}
