#bd
DB_DRIVE = postgres
DB_USER = postgres
DB_PASS = qwerty
DB_IP = localhost
DB_PORT = 5436
DB_NAME = postgres
DB_OPT = sslmode=disable

#DATABASE
MIGRATIONSDIR = repository/schema


run : build
	./.bin/auth  --config_path=$(CONFIGPATH) --config_name=$(CONFIGNAME)

build : update
	go build -o ./.bin/auth ./cmd/main.go

update : 
	go mod tidy

docker-up :
	chmod +x docker/run.sh
	sh -C docker/run.sh

migr_init : docker-up
	migrate create -ext sql -dir $(MIGRATIONSDIR) -seq init
	
migr_up : docker-up
	migrate -path $(MIGRATIONSDIR) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up

migr_down : migr_up
	migrate -path $(MIGRATIONSDIR) -database "$(DB_DRIVE)://$(DB_USER):$(DB_PASS)@$(DB_IP):$(DB_PORT)/$(DB_NAME)?$(DB_OPT)" -verbose up

clear : 
	rm -rf .bin