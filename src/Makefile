
run : build
	./.bin/auth  --config_path=$(CONFIGPATH) --config_name=$(CONFIGNAME)

build : update
	go build -o ./.bin/auth ./cmd/main.go

update : 
	go mod tidy
	
clear : 
	rm -rf .bin