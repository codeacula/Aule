// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/codeacula/Aule/discord-bot/models"
	"github.com/codeacula/Aule/discord-bot/restapi/operations"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate swagger generate server --target ..\..\discord-bot --name DiscordBot --spec ..\swagger.yml --principal interface{}

type ApiKeys struct {
}

type ContextKey string

const mongoClientKey ContextKey = "mongoDatabase"

type MiningNode struct {
	NodeType    string    `bson:"nodeType"`
	NodeSize    string    `bson:"nodeSize"`
	RoomID      int64     `bson:"roomId"`
	ClaimedBy   string    `bson:"claimedBy"`
	LastChecked time.Time `bson:"lastChecked"`
}

var localContext = context.Background()

func convertToMiningNode(incomingNode *models.Node) MiningNode {
	return MiningNode{
		NodeType:    *incomingNode.NodeType,
		NodeSize:    *incomingNode.NodeSize,
		RoomID:      *incomingNode.RoomID,
		ClaimedBy:   incomingNode.ClaimedBy,
		LastChecked: time.Now(),
	}
}

func getNodeSummary(ctx *context.Context) {

}

func addNodeToUpdates(ctx *context.Context, node MiningNode) {
	client := (*ctx).Value(ContextKey("mongoDatabase")).(*mongo.Client)
	database := client.Database("discord-mining-bot")
	coll := database.Collection("node-updates")
	_, err := coll.InsertOne(localContext, node)

	if err != nil {
		log.Fatal(err)
	}
}

func updateNodeInDb(ctx *context.Context, node MiningNode) {
	client := (*ctx).Value(ContextKey("mongoDatabase")).(*mongo.Client)
	database := client.Database("discord-mining-bot")
	coll := database.Collection("nodes")
	opts := options.Update().SetUpsert(true)

	filter := bson.D{{Key: "roomId", Value: node.RoomID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "nodeType", Value: node.NodeType},
		{Key: "nodeSize", Value: node.NodeSize},
		{Key: "claimedBy", Value: node.ClaimedBy},
		{Key: "lastChecked", Value: node.LastChecked},
	}}}

	_, err := coll.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		panic(err)
	}
}

func configureFlags(api *operations.DiscordBotAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DiscordBotAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(localContext, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to MongoDB")
	}

	localContext = context.WithValue(localContext, mongoClientKey, client)

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	if api.GetNodeHandler == nil {
		api.GetNodeHandler = operations.GetNodeHandlerFunc(func(params operations.GetNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetNode has not yet been implemented")
		})
	}

	api.GetSummaryHandler = operations.GetSummaryHandlerFunc(func(params operations.GetSummaryParams) middleware.Responder {
		getNodeSummary(&localContext)
		return middleware.NotImplemented("operation operations.GetSummary has not yet been implemented")
	})

	api.PostMessageHandler = operations.PostMessageHandlerFunc(func(params operations.PostMessageParams) middleware.Responder {
		miningNode := convertToMiningNode(params.Body)
		addNodeToUpdates(&localContext, miningNode)
		updateNodeInDb(&localContext, miningNode)
		return operations.NewPostMessageOK().WithPayload("Ok")
	})

	api.ReturnPingHandler = operations.ReturnPingHandlerFunc(func(params operations.ReturnPingParams) middleware.Responder {
		return operations.NewReturnPingOK().WithPayload("pong")
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
