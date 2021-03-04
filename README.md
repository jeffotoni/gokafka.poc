# gokafka.poc

A repository where you have an excellent example of a game made with a Game made in Go using Webassembly, communicating with kafka and everything running in GKE.

We have the game compiled for webassembly, we have the Apis in Go exposing to be consumed by the game and communicating with pkg kafka and some examples of kafka using docker examples using confluent, kafka bitnami and kafka fast where I did a summer clean to test the location using docker.

<h2 align="center">
  <br/>
  <img src="https://github.com/jeffotoni/gokafka.poc/blob/beta/img/game.png" alt="logo" width="900" />
  <br />
  <br />
  <br />

### ping server
```bash
$ curl -i -XPOST localhost:8181/ping
```

### Create Topic
```bash
$ curl -i -XPOST -H "Content-type:application/json" \
localhost:8181/topic \
-d '{"name":"mytopic8", "partition":3, "replication_factor":1}'
```

### List all topic
```bash
$ curl -i -XGET -H "Content-type:application/json" \
localhost:8181/topic
```

### Delete topic
```bash
$ curl -i -XDELETE -H "Content-type:application/json" \
localhost:8181/topic/mytopic8
```

### Producer
```bash
$ curl -i -XPOST -H "X-Key-User:12121212121" \
-H "Content-type:application/json" localhost:8181/producer \
-d '{"key":"x100001", "name":"jefferson otoni","score":3000}'
```
