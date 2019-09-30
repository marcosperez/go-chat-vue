function SocketClient() {

    this.url = "ws://127.0.0.1:8357/ws";

    this.onConnect = () => console.log("conectado")
    this.onDisconect = () => console.log("desconectado")

    this.connect = function () {
        const client = this;
        this.socket = new WebSocket(this.url);
        const socket = this.socket;

        this.socket.onopen = function () {
            client.onConnect();
            console.log("connected to " + this.url);
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
}

var socketClient = new SocketClient();