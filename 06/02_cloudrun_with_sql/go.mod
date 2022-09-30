module 0delta/cloudrun_with_sql

go 1.12

require (
	github.com/0Delta/CloudRunSample/handler v0.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.9.0
)

replace github.com/0Delta/CloudRunSample/handler => ./handler
