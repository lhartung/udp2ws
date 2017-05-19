# UDP-to-WS Proxy

This proxy receives messages on a UDP port and broadcasts them to one or more
WebSocket clients.

It is based on the [gorilla/websocket](https://github.com/gorilla/websocket)
[chat example](https://github.com/gorilla/websocket/tree/master/examples/chat).

## Compiling

    go get github.com/gorilla/websocket
    go build .

## Running the Server

    ./udp2ws

By default the HTTP server listens on port 8080.  Open http://localhost:8080/
in your web browser to view the WebSocket side of the proxy.

By default the proxy listens on UDP port 4000.  You can use `netcat` or
a similar utility to send UDP messages to the proxy.  For example:

    echo "Hello, World" >/dev/udp/localhost/4000

or

    echo "Hello, World" | netcat -u localhost 4000 -w1

The message should appear in your web browser.

## Configuration

Usage of ./udp2ws:

  -http string
    	HTTP service address (default ":8080")
  -udp string
    	UDP listening address (default ":4000")

