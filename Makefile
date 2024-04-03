migrateup:
	migrate -verbose -path migrations -database "postgresql://postgres:postgres@localhost:5432/advertising?sslmode=disable" up 

migratedown:
	migrate -verbose -path migrations -database "postgresql://postgres:postgres@localhost:5432/advertising?sslmode=disable" down 

create:
	migrate create -ext sql -dir migrations create_ads_table

.PHONY: migrateup migratedown create