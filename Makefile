main:
	go run cmd/main.go -c configs/config.toml
MysqlStart:
	mysql.server start

MysqlStop:
	mysql.server stop

.PHONY: main MysqlStart MysqlStop