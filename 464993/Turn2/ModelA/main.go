package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Movie struct {
	Title    string
	Year     int
	Genre    string
	Duration int
}

type MovieSlice []Movie

func (ms MovieSlice) SortByTitle() {
	fmt.Println("Sorting by title...")
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Title < ms[j].Title
	})
}

func (ms MovieSlice) SortByYear() {
	fmt.Println("Sorting by year...")
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Year < ms[j].Year
	})
}

func (ms MovieSlice) FilterByGenre(genre string) MovieSlice {
	fmt.Println("Filtering by genre:", genre)
	filtered := []Movie{}
	for _, movie := range ms {
		if movie.Genre == genre {
			filtered = append(filtered, movie)
		}
	}
	return filtered
}

func (ms MovieSlice) Display() {
	for _, movie := range ms {
		fmt.Printf("Title: %s, Year: %d, Genre: %s, Duration: %d minutes\n",
			movie.Title, movie.Year, movie.Genre, movie.Duration)
	}
}

func (ms MovieSlice) Search(query string) MovieSlice {
	fmt.Println("Searching for:", query)
	filtered := []Movie{}
	for _, movie := range ms {
		if strings.Contains(strings.ToLower(movie.Title), strings.ToLower(query)) {
			filtered = append(filtered, movie)
		}
	}
	return filtered
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	movies := MovieSlice{
		{Title: "Inception", Year: 2010, Genre: "Sci-Fi", Duration: 148},
		{Title: "The Godfather", Year: 1972, Genre: "Crime", Duration: 175},
		{Title: "Pulp Fiction", Year: 1994, Genre: "Crime", Duration: 154},
		{Title: "The Shawshank Redemption", Year: 1994, Genre: "Drama", Duration: 142},
		{Title: "The Dark Knight", Year: 2008, Genre: "Crime", Duration: 152},
	}

	for {
		fmt.Println("\nMovie Manager")
		fmt.Println("1. Display all movies")
		fmt.Println("2. Sort by title")
		fmt.Println("3. Sort by year")
		fmt.Println("4. Filter by genre")
		fmt.Println("5. Search by title or keyword")
		fmt.Println("6. Exit")

		fmt.Print("Choose an option: ")
		if !scanner.Scan() {
			fmt.Println("Exiting...")
			return
		}

		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			movies.Display()
		case "2":
			movies.SortByTitle()
			movies.Display()
		case "3":
			movies.SortByYear()
			movies.Display()
		case "4":
			fmt.Print("Enter genre to filter: ")
			if !scanner.Scan() {
				continue
			}
			genre := strings.TrimSpace(scanner.Text())
			filteredMovies := movies.FilterByGenre(genre)
			fmt.Println("Filtered Movies:")
			filteredMovies.Display()
		case "5":
			fmt.Print("Enter movie title or keyword to search: ")
			if !scanner.Scan() {
				continue
			}
			query := strings.TrimSpace(scanner.Text())
			searchResults := movies.Search(query)
			fmt.Println("Search Results:")
			searchResults.Display()
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
