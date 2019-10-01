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
                client.send('ping', JSON.stringify("ping...."), function (data) {
                    console.log('ACK from server wtih data: ', data);
                })
            }, 15000);
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
        var data = `${topic}:${msg}`
        this.socket.send(data);
    };

    this.subscribe = function (topic) {
        this.socket.send(`subscribe:${topic}`);
    }
}

var socketClient = new SocketClient();