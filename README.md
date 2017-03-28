# go-grpc-prometheus-demo

This repository contains the demo of using [go-grpc-prometheus](https://github.com/grpc-ecosystem/go-grpc-prometheus).  
The Client of gRPC and also the server outputs metrics, and monitoring metrics by using the [Grahana](https://grafana.com/).  

## How to start demo

The demo are built up by [Docker Compose](https://docs.docker.com/compose/), so that you have to install the Docker Compose.

### Step 1

Containers is begin, after using the following make command.
```
$ make start
```

Begin to run containers are:

- Prometheus `localhost:9090`
- Server of gRPC `localhost:8080`
- Client of gRPC `localhost:8081`
- Grafana `localhost:3000`

### Step 2

After starting containers, you are able to access following url and to send request to server from client of gRPC.
```
$ curl -XGET 'http://localhost:8081/echo?m=message'
Response: echo message%
```

The gRPC connection of demo is simple unary RPCs.

### Step 3

After sending any requests, you are able to monitoring of metrics at dashboard of Prometheus.

Navigate to `http://localhost:9090/graph` and From `insert metric at cursor` menu, choose `grpc_client_msg_sent_total` and `grpc_server_msg_sent_total` and so on after will see metrics of gRPC.

### Step 4

Step 4 is setup Grafana.

Navigate to `http://localhost:3000/login` and login with admin user.

- User: admin
- Password: admin

From the Grafana menu, choose `Data Sources` and click `Add data source`, then input by using the following values.

- Name: prometheus
- Type: prometheus
- Url: http://localhost:9090
- Access: direct

At last you can import the dashboard templates from the grafana directory from this repository. From the Grafana menu, choose `Dashboards` and click on `Import`.

### Other

If you want to stop containers, you can use following make command.
```
$ make stop
```