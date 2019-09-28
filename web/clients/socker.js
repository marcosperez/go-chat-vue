function SocketClient() {

    this.url = "ws://localhost:8000/socket.io/";

    this.connect = function () {
        this.socket =
            this.socket = io(this.url);
        // Abre la conexión
        this.socket.addEventListener('open', function (event) {
            console.log("Conectado")
            this.socket.send('Hello Server!');
        });
    }
}

var socketClient = new SocketClient();