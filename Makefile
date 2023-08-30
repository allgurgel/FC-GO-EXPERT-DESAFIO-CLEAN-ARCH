createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq create_order_table

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" --verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" --verbose down

.PHONY: migrate migratedown createmigration
	