# shapes

Set `GOPATH` and append to `PATH`

```
export GOPATH=$MY_GOPATH
export PATH=$PATH:$GOPATH/bin
```

Fetch and install code

```
go get github.com/paulcadman/shapes
go install github.com/paulcadman/shapes/cmd/draw
```

Run the program

`$GOPATH/bin/draw`

Visit <http://localhost:2003/draw>
