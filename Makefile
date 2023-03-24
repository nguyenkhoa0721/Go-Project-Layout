include config.yml

echo:
	echo ${APP.grpcServerAddress}

migrateup:
	migrate -path dev/sql -database "postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE_NAME}?sslmode=disable" -verbose up
