## Description

Exercise to make use of read-through pattern by building cache on top of db in GO. This is just mimicking of the real world application having persistance db and cache on top of it.

100 data items are populated in the db.
500 concurrent http requests are made.
Only 100 db hit are made with rest served from cache.

## Pre Requirements

* GO installed in the system and any supported IDE

## Test

```bash
# run test
$ go test

# test run data
db query count 100
cache hit count 400
PASS
ok  	db-cache-wrapper	0.986
```