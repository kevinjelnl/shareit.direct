# Makefile to make building the image / containers easier
#
# builds the app in the local dir and names it "main"
build_app:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# builds the multistage docker image with Dockerfile.scratch and pulls in the binary
build_image:
	docker build -t shareit.direct:latest -f Dockerfile .

# runs the container
run_container:
	docker run -it -p 8080:8080 shareit.direct:latest



