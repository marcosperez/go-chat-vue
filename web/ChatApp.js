Vue.use(VueMaterial.default)
var app = new Vue({
    el: '#app',
    data: {
        status: "desconectado",
        user: null,
    },
    mounted() {
        socketClient.onConnect = this.connected;
        socketClient.onDisconect = this.connected;
    },
    methods: {
        connected: (evt) => {
            app.status = "conectado";
        },
        disconnect: (evt) => {
            app.status = "desconectado";
        },
        login: function (name) {
            this.user = name;
            socketClient.connect();
        }
    },
})