# Go Meetup 2022

Slides for 2022 Go Meetup

## Running the slides

```sh
make play
```

Since `present` tool does not work in module mode, certain depenencies need to be fetched in GOPATH mode:

```sh
cd $GOPATH # just navigate outside any Go module
GO111MODULE=off go get golang.org/x/exp/constraints
GO111MODULE=off go get golang.org/x/exp/slices
GO111MODULE=off go get golang.org/x/exp/maps
```
