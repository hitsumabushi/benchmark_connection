# benchmark_connection
measure limit of the number of concurrent connections

## Limitation
- Support only IPv4 address

## How to use
Firstly, you have to launch a server.
Then, you run clients.
These clients try to connect specified server tcp4 socket.

### Server

```sh
./benchmark_connection --beginport 60000 --endport 60010 server
```

| parameter | explanation  |
|-----------|--------------|
| beginport | lower limit of number |
| endport   | higher limit of number |

### Client

```sh
./benchmark_connection --beginport 60000 --endport 60010 client --host 127.0.0.1
```

| parameter | explanation  |
|-----------|--------------|
| beginport | lower limit of port |
| endport   | higher limit of port |
| host      | target host IP/FQDN |
