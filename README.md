A shim for watching how quickly things increment

Watch an ES index fill up
```
while sleep 1; do curl -s http://127.0.0.1:9200/index_name/_count; done | go run gorate.go
```

