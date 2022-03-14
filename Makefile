# dbml 文件 转 sql 文件
dbml2sql:
	dbml2sql database.dbml -o database.sql --postgres
# 生成在线数据库文档
dbdocs:
	dbdocs build database.dbml

# docker 运行 postgresql 数据库
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# 创建数据库
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root soybean_admin

# 删除数据库
dropdb:
	docker exec -it postgres12 dropdb soybean_admin

# 执行向上迁移数据库
migrateup:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5432/soybean_admin?sslmode=disable" -verbose up

# 执行向下迁移数据库
migratedown:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5432/soybean_admin?sslmode=disable" -verbose down

# 根据sqlc文件 自动生成 curd golang 代码
sqlc:
	sqlc generate
# 执行测试
test:
	go test -cover ./...
# 启动服务
server:
	go run main.go

# 生成测试接口
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/honghuangdc/soybean-admin-go/db/sqlc Store

.PHONY: dbml2sql dbdocs postgres createdb dropdb migrateup migratedown sqlc test server mock
# 查看帮助
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help