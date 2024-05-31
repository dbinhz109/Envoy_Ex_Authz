
# Example repo to demonstrate Envoy External Authorization with Golang GRPC service

|  Request | header  | status  | http_server_response  |
|:-:|:-:|:-:|:-:|
| / private | x-allow-private = "admin"  | 200  | "message": "This is a private endpoint"  |
| / private |  -  | 200  | "message": "This is a private endpoint"  |
| / public  |  -  | 200  | "message": "This is a public endpoint"  |

# To run this example

1. Start the envoy server

```sh
envoy -c envoy.yml
```

2. start the go_grpc_filter & go_simple_http servers by navigating to the cluster root.
```sh
go run main.go
```

3. Go to http://localhost:10000



