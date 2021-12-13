package main

import (
	"fmt"
	"github.com/fbiville/impersonation-demo/pkg/io"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	authenticatedUser := "neo4j"
	executingUser := "jane"

	driver, err := neo4j.NewDriver("neo4j://localhost", neo4j.BasicAuth(authenticatedUser, "s3cr3t", ""))
	io.MaybePanic(err)
	defer io.MaybePanicFn(driver.Close)
	session := driver.NewSession(neo4j.SessionConfig{
		ImpersonatedUser: executingUser,
		BoltLogger:       neo4j.ConsoleBoltLogger(),
	})
	defer io.MaybePanic(session.Close())
	favouriteMovieTitle, err := session.ReadTransaction(readFavouriteMovie)
	io.MaybePanic(err)
	fmt.Printf("Favourite movie title is: %s\n", favouriteMovieTitle)
}

func readFavouriteMovie(tx neo4j.Transaction) (interface{}, error) {
	result, err := tx.Run("MATCH (m:FavouriteMovie) RETURN m.title AS title LIMIT 1", map[string]interface{}{})
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
