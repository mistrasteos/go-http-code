# go-http-code
Simple http server written in go to play with HTTP response status codes

## Listening port
Listening port can be configured using `PORT`environment variable


**Dockerfile**
```
docker run --rm -it --name go-http-code -e PORT=5555 go-http-code:latest
2024/03/25 08:41:18 INFO Listening on port=5555
```

**shell**
```
PORT=6666 ./go-http-code
2024/03/25 09:42:02 INFO Listening on port=6666
```

## Examples

```
curl -i http://localhost:4444/504
HTTP/1.1 504 Gateway Timeout
Content-Type: text/plain;

504, Gateway Timeout
```
```
curl -i http://localhost:4444/205
HTTP/1.1 205 Reset Content
Content-Type: text/plain;

205, Reset Content
```

```
curl -i http://localhost:4444/thisistheway
HTTP/1.1 404 Not Found
Content-Type: text/plain;

Not found, /thisistheway
```
