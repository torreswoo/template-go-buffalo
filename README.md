# Torres Template Buffalo 

[Powered by Buffalo](http://gobuffalo.io)

Torres template project using Buffalo. A Go web development eco-system, designed to make the life of a Go web developer easier.
- use SQLite3

## requirements
```bash
# requirements
$ gvm use go1.10  # go1.8.1+
$ nvm use node

# check
$ echo $GOPATH
$ go version
$ gcc -v
$ node -v
$ npm -v
```

## install & start
```bash
# install
$ go get -u -v github.com/gobuffalo/buffalo/buffalo
 
# make project
$ buffalo new <nombre> --db-type sqlite3
$ buffalo db create -a
$ buffalo db migrate up
$ buffalo dev           # localhost:3000
```
- [http://127.0.0.1:3000](http://127.0.0.1:3000)

## Docker support
```bash
$ docker build -f Dockerfile -t torreswoo/go-blog .
$ docker run --name go-blog -p 3000:3000 torreswoo/go-blog
```