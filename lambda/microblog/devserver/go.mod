module microblog-devserver

go 1.16

require (
	github.com/aws/aws-sdk-go v1.38.51
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	microblog v0.0.0
)

replace microblog v0.0.0 => ../
