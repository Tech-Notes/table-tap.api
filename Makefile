
.PHONY: migration_up migration_down

migration_up:
	migrate -path ./migrations/ -database $(TABLE_TAP_DB) -verbose up

migration_down:
	migrate -path ./migrations/ -database $(TABLE_TAP_DB) -verbose down

migration_fix:
	migrate -path ./migrations/ -database $(TABLE_TAP_DB) force $(VERSION)

create_migration:
	migrate create -ext sql -dir ./migrations/ -seq $(SEQ_NAME)