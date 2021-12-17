# Go + Docker -> DigitalOcean

### Create a go project
Ideally with some http capabilities so we can test it in the browser

`go mod init daily`

```
package main

func main() {
	fmt.Println("Go web app started on port: 3000")

	http.HandleFunc("/", homePage)

	http.ListenAndServe(":3000", nil)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Homepage")
}
```


### Create Dockerfile

```
FROM golang:1.17.4-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
```

### Build the docker image
`docker build -t daily .`

Show all images on our Docker installation

`docker images`

Run it locally, exposing port #1 to the world to the #2 TCP port that you used in your app

`docker run -d -p 3000:3000 daily`  <-- detached

To kill a process

`docker ps`

`docker kill [Container ID]`

### On DO console:
`git clone https://github.com/user/project.git app`       <-- same folder as above in Dockerfile

`cd app`

`docker build -t daily .`

`docker run -d -p 3000:3000 daily`