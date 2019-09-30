module 0delta/cloudrun_with_sql

go 1.12

require (
	github.com/0Delta/CloudRunSample/handler v0.0.0
	github.com/GoogleCloudPlatform/cloudsql-proxy v0.0.0-20190828224159-d93c53a4824c
	github.com/go-sql-driver/mysql v1.4.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.10
)

replace github.com/0Delta/CloudRunSample/handler => ./handler
