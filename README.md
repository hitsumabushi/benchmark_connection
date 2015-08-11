# benchmark_connection
benchmark connection limit

## How to use

### Server

```sh
./benchmark_connection --beginport 60000 --endport 60010 server
```

| parameter | explanation  |
|-----------|--------------|
| beginport | lower limit of use port |
| endport   | higher limit of use port |



### Client

```sh
./benchmark_connection --beginport 60000 --endport 60010 client --host 127.0.0.1
```

| parameter | explanation  |
|-----------|--------------|
| beginport | lower limit of use port |
| endport   | higher limit of use port |
| host      | target host IP/FWDN |
