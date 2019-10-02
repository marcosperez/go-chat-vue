function SocketClient() {

    this.onConnect = () => console.log("conectado")
    this.onDisconect = () => console.log("desconectado")

    this.connect = function () {
        const client = this;
        this.socket = new WebSocket(config.wsURL);
        // const socket = this.socket;

        this.socket.onopen = function () {
            client.onConnect();
            console.log("connected to " + config.wsURL);
            client.send("ping", "ping....")
            setInterval(() => {
                client.send('ping', 'ping', function (data) {
                    console.log('ACK from server wtih data: ', data);
                })
            }, 2000);
        }

        this.socket.onclose = function (e) {
            console.log("connection closed (" + e.code + ")");
            client.onDisconect();
        }

        this.socket.onmessage = function (e) {
            console.log("message received: " + e.data);
        }
    }

    this.send = function (topic, msg) {
        var data = {
            type: topic,
            data: msg
        }
        this.socket.send(JSON.stringify(data));
    };

    this.subscribe = function (topic) {
        this.send("subscripcion", { room: topic });
    }
}

var socketClient = new SocketClient();