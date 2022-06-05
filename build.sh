go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_proxy/lepus_proxy.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_alarm/lepus_alarm.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_task/lepus_task.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_collector/mysql/lepus_mysql_mon.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_collector/redis/lepus_redis_mon.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_collector/mongo/lepus_mongo_mon.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_collector/oracle/lepus_oracle_mon.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_collector/postgres/lepus_postgres_mon.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_collector/greatsql/lepus_greatsql_mon.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_collector/sqlserver/lepus_sqlserver_mon.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_collector/web/lepus_web_mon.go
go build -o ./bin/ -ldflags "-s -w" -a ./src/lepus_agent/lepus_server_agent.go

