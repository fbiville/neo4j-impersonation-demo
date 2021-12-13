package main

import (
	"fmt"
	"github.com/fbiville/impersonation-demo/pkg/io"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	authenticatedUser := "neo4j"

	driver, err := neo4j.NewDriver("neo4j://localhost", neo4j.BasicAuth(authenticatedUser, "s3cr3t", ""))
	io.MaybePanic(err)
	defer io.MaybePanicFn(driver.Close)
	session := driver.NewSession(neo4j.SessionConfig{
		BoltLogger: neo4j.ConsoleBoltLogger(),
	})
	defer io.MaybePanic(session.Close())
	favouriteMovieTitle, err := readFavouriteMovie(session)
	io.MaybePanic(err)
	fmt.Printf("Favourite movie title is: %s\n", favouriteMovieTitle)
}

func readFavouriteMovie(session neo4j.Session) (interface{}, error) {
	result, err := session.Run("MATCH (m:FavouriteMovie) RETURN m.title AS title LIMIT 1", map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	record, err := result.Single()
	if err != nil {
		return nil, err
	}
	value, _ := record.Get("title")
	return value, nil
}
