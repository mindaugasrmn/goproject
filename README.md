<h1>Golang starter project</h1>
Golang starter project for projects based on REST and Angular
<br />
<h2>Features:</h2>
REST<br />
JWT<br />
Mongodb Middleware for storing data<br />
Redis for stroing used tokens<br />
Config.json for setting db settings<br />
Serving static files for angular<br />
Angular with Material starter project included<br />

<h2>How to run:</h2>
create file in GOPATH path, lets say server.go with content:
<pre>
package main

import (
    "github.com/gorilla/context"
	"github.com/mindaugasrmn/goproject/routers"
	"github.com/mindaugasrmn/goproject/settings"
	"github.com/codegangsta/negroni"
	"github.com/mindaugasrmn/goproject/core/mongo"
	"net/http"
	"log"
)

func main() {
	log.Print("Server started and listening on port 5000")
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.Use(mongo.MongoMiddleware())
	n.UseHandler(context.ClearHandler(router))
	http.ListenAndServe(":5000", n)
}

</pre>

then:
<pre>
$ cd $GOPATH
$ go get
$ mv src/github.com/mindaugasrmn/goproject/static .
</pre>

then start mongo and redis, and then run application:
<pre>
$ go run server.go
</pre>


