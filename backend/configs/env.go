package configs

/*
	import (
	   "log"
	   "os"

)
*/
func EnvMongoURI() string {
	const uri = "mongodb://minijo:miata123@127.0.0.1:27017/?maxPoolSize=20&w=majority"
	return uri
	//return os.Getenv("MONGOURI")
}
