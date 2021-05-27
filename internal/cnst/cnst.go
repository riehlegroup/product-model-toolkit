package cnst

// General
const (
	Debug = true
	Empty = ""
	DefaultGrpcPort = "56985"
	SPDX = "spdx"
	Composer = "composer"
	FileHasher = "hasher"
)

// Database constants
const (
	MongoDBDefaultHost = "datastore:27017"
	MongoDBDevelopmentHost = "mongodb://127.0.0.1:60731" // depends
	MongoDBDefaultRetryNumber = 0
	MongoDBDefaultDatabaseName = "pmt"
	MongoDBDefaultCollectionName = "products"
)