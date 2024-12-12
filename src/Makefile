
run : build
	./.bin/auth  --config_path=$(CONFIGPATH) --config_name=$(CONFIGNAME)

build :
	go build -o ./.bin/auth ./cmd/main.go

clear : 
	rm -rf .bin