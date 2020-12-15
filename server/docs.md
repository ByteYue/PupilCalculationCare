```
.
├── db
│   ├── db
│   │   └── sqlc        #database control
│   │       ├── account.sql.go
│   │       ├── account_test.go
│   │       ├── db.go
│   │       ├── go.mod
│   │       ├── go.sum
│   │       ├── main_test.go
│   │       ├── models.go
│   │       ├── querier.go
│   │       └── util
│   │           └── random.go
│   └── Makefile
├── Dockerfile          #develop environment deployment
├── docs.md             #development document
├── expression          #expression generate
│   ├── check.go        #check if satisfies
│   ├── filehandler.go  #transfer2json
│   ├── expression_test.go  #expression generate test
│   └── randomOp.go     #generatre expression
├── go.mod              #server dependencyies
├── go.sum
├── model               #data struct
│   ├── request
│   │   └── user.go
│   ├── response
│   └── technology.go
├── my-app              #binary
├── server.go           #the starting point of the server
├── router              #router function
│   ├── initialize.go   #initialize the router
│   ├── router_test.go  #router test
│   ├── sys_user.go     #group initialize
│   └── user_api.go     #APIs
├── stack               #simple stack container
│   └── stack.go
└── web                 #REST APIs,web server
    ├── app.go
    └── app_test.go

```
