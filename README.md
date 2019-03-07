# webp2jpg
convert the *.webp to *.jpg


# Dependence
```bash
  golang.org/x/image/webp
  gopkg.in/urfave/cli.v1
``` 

```bash
  go get gopkg.in/urfave/cli.v1
  git clone https://github.com/golang/image $GOPATH/src/golang.org/x/image
  cd $GOPATH/src/golang.org/x/image/webp && go build
```
# Download

```bash
  go get github.com/sndnvaps/webp2jpg
```

# Build && Run

```bash
  cd $GOPATH/src/github.com/sndnvaps/webp2jpg
  go build
```

# Convert the webp file to jpeg

```bash
  put the test.webp file into webp2jpg folder
  ./webp2jpg -t jpg -s test.webp
  ./webp2jpg --type jpg --source test.webp
  ./webp2jpg --type jpeg --source test.webp
  ./webp2jpg --type png --source test.webp
  ./webp2jpg --type gif --source test.webp
```
