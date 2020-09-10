# gokafka.poc

Uma poc para presentar Kafka e Go

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
