hello:
	echo "gotodo"

run:
	go run main/gotodo/main.go embedded run --manifest manifest.yaml

make update:
	go get -v github.com/mkawserm/abesh
	go get -v github.com/mkawserm/httpserver2
