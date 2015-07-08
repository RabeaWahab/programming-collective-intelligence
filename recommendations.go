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
					"Superman Returns" : 4.0},
				"Rabee" : map[string]float64{
					"Snakes on a Plane" : 4.49,
					"You, Me and Dupree" : 0.8,
					"Superman Returns" : 4.0}}

func main() {
	Euclidean("Toby", "Rabee", critics)
	Pearson("Toby", "Rabee", critics)
	Similarity("Rabee", "pearson", 2, critics)
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
	var numenator float64 = 0
	for _, movieName := range commonMovies {
		numenator += math.Pow(user1Ratings[movieName] - user2Ratings[movieName], 2)
	}

	// square root the numenator and divide one by it to get a higher ranking for similarity between 0-1
	result := 1 / (1+(math.Sqrt(numenator)))

	// print result
	fmt.Println(result)

	return result
}

// Pearson Correlation Score
func Pearson(user1 string, user2 string, critics map[string]map[string]float64) float64 {
	var commonMovies = make(map[string]float64)

	user1Ratings := critics[user1]
	user2Ratings := critics[user2]

	// return false if no movies in any of the two users
	if len(user1Ratings) == 0 || len(user2Ratings) == 0 {
		return 0
	}

	// check for similar movies and add them to a seperate array
	var usersSumOfProducts float64 = 0
	var sumUser1Ratings float64 = 0
	var sumUser1PowRatings float64 = 0
	var sumUser2Ratings float64 = 0
	var sumUser2PowRatings float64 = 0


	for movieName, rating := range user1Ratings {
		if user2Ratings[movieName] > 0 {
			commonMovies[movieName] = rating
			usersSumOfProducts += user1Ratings[movieName] * user2Ratings[movieName]

			sumUser1Ratings 	+= user1Ratings[movieName]
			sumUser1PowRatings 	+= math.Pow(user1Ratings[movieName], 2)

			sumUser2Ratings 	+= user2Ratings[movieName]
			sumUser2PowRatings 	+= math.Pow(user2Ratings[movieName], 2)
		}
	}

	// number of common movies , then cast to float64 for the formula
	numberOfElementsInCommon := float64(len(commonMovies))

	// Formula
	numenator 	:= usersSumOfProducts - ((sumUser1Ratings*sumUser2Ratings)/numberOfElementsInCommon)
	denominator := math.Sqrt((sumUser1PowRatings - (math.Pow(sumUser1Ratings, 2)/numberOfElementsInCommon)) *
					(sumUser2PowRatings - (math.Pow(sumUser2Ratings, 2)/numberOfElementsInCommon)))

	if denominator == 0 {
		return 0
	}

	result := numenator/denominator

	// print result
	fmt.Println(result)

	return result
}

func Similarity(user string, method string, heighestSimilaritiesLimit int, critics map[string]map[string]float64) map[string]float64 {
	var results = make(map[string]float64)
	var similarity float64 = 0

	for critic, _ := range critics {
		if critic != user {
			if method == "pearson" {
				similarity = Pearson(critic, user, critics)
			}

			if method == "euclidean" {
				similarity = Euclidean(critic, user, critics)
			}

			results[critic] = similarity
		}
	}

	fmt.Println(results);
	return results
}
