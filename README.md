# PSQL LisTify

PostgreSQL Listen Notify by Example

## Usage

LisTify expects postgres to be running on localhost.

Create a listify database:

```
$ createdb -O USERNAME listify
```

Replace `USERNAME` with your postgresql db username.


```
$ go get github.com/Tonkpils/psql-listify
$ cd $GOPATH/src/github.com/Tonkpils/psql-listifify
$ go run listify.go -db.user USERNAME
```

