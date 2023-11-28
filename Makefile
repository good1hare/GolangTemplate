migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

migrate-up: ### migration up sqlite data/user.db
	migrate -path migrations -database 'mysql://root:password@tcp(0.0.0.0:3303)/db_name?query' up
.PHONY: migrate-up

migrate-down: ### migration down
	migrate -path migrations -database 'mysql://root:password@tcp(0.0.0.0:3303)/db_name?query' down
.PHONY: migrate-down