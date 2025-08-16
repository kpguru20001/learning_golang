package handlers

import (
	"shadow-docs/graph"
	"shadow-docs/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

func GraphQLHandler() gin.HandlerFunc {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func GQLPlaygroundHandler() gin.HandlerFunc {
	opts := playground.ApolloSandboxOption(playground.WithApolloSandboxInitialStateIncludeCookies(true))
	h := playground.ApolloSandboxHandler("GraphQL", "/graphql", opts)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
