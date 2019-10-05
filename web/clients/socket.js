function SocketClient() {

    this.onConnect = () => console.log("conectado")
    this.onDisconect = () => console.log("desconectado")
    this.userID = null;

    this.connect = function (userID) {
        const self = this;
        self.userID = userID;
        self.socket = new WebSocket(config.wsURL);
        // const socket = self.socket;

        self.socket.onopen = function () {
            self.onConnect();
            self.send("connection")
            setTimeout(() => {
                self.send("chat", { chatID: "global", message: "Hola mundo" })
            })
        }

        self.socket.onclose = function (e) {
            console.log("connection closed (" + e.code + ")");
            self.onDisconect();
        }

        self.socket.onmessage = function (e) {
            console.log("[onmessage] " + e.data);
            msg = JSON.parse(e.data)
            if (!msg || !msg.type) {
                return;
            }
            switch (msg.type) {
                case "ping":
                    console.log("[onmessage][ping]" + e.data);
                    self.send("pong")
                case "chat":
                    console.log("[onmessage][ping]" + e.data);
                    self.send("pong")
                default:
                    console.log("[onmessage] tipo de mensaje no controlado " + e.data);
            }
        }
    }

    this.send = function (type, msg) {
        var data = {
            type: type,
            userID: this.userID,
            data: msg
        }

        this.socket.send(JSON.stringify(data));
    };

    this.subscribe = function (topic) {
        this.send("suscription", { chatID: topic });
    }
}

var socketClient = new SocketClient();