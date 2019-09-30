var app = new Vue({
    el: '#app',
    data: {
        status: "desconectado"
    },
    mounted() {
        socketClient.onConnect = this.connected;
        socketClient.onDisconect = this.connected;
        socketClient.connect();
    },
    methods: {
        connected: (evt) => {
            app.status = "conectado";
        },
        disconnect: (evt) => {
            app.status = "desconectado";
        }
    },
})