# golang+docker example: not building

docker build failing:

```bash
$ sudo docker build . -t local/pruebago
Sending build context to Docker daemon  375.3kB
Step 1/7 : FROM golang:1.17.7 as build
1.17.7: Pulling from library/golang
0c6b8ff8c37e: Already exists 
412caad352a3: Already exists 
e6d3e61f7a50: Already exists 
461bb1d8c517: Already exists 
9297634c9537: Already exists 
c9cefb987250: Pull complete 
8560fc463426: Pull complete 
Digest: sha256:1a35cc2c5338409227c7293add327ebe42e1ee5465049f6c57c829588e3f8a39
Status: Downloaded newer image for golang:1.17.7
 ---> 260de46bdc85
Step 2/7 : WORKDIR /app
 ---> Running in 6715ab07bf52
Removing intermediate container 6715ab07bf52
 ---> 72f8e8b11ee6
Step 3/7 : COPY ./src/go.mod .
 ---> 242d9ba170f9
Step 4/7 : COPY ./src/go.sum .
 ---> 20d2d420847a
Step 5/7 : RUN go mod download
 ---> Running in c9616cb3c669
Removing intermediate container c9616cb3c669
 ---> 4b0a7f4b27af
Step 6/7 : COPY ./src/* ./
 ---> 1d3f245124b4
Step 7/7 : RUN go build -o /tmp/built
 ---> Running in 72358fb165c4
main.go:8:2: no required module provides package github.com/David-Lor/go-example/internal/foo; to add it:
        go get github.com/David-Lor/go-example/internal/foo
main.go:9:2: no required module provides package github.com/David-Lor/go-example/internal/foo/bar; to add it:
        go get github.com/David-Lor/go-example/internal/foo/bar
The command '/bin/sh -c go build -o /tmp/built' returned a non-zero code: 1
```

build and run from docker run success:

```bash
$ sudo docker run -it --rm -v $(pwd):/data golang:1.17.7
root@e468a186536f:/go# cd /data/src
root@e468a186536f:/data/src# go build -o /tmp/built
go: downloading github.com/gammazero/workerpool v1.1.2
go: downloading github.com/gammazero/deque v0.1.0
root@e468a186536f:/data/src# /tmp/built
internal/foo
internal/foo/bar
root@e468a186536f:/data/src# go run main.go
internal/foo
internal/foo/bar
```
