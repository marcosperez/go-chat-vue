Vue.use(VueMaterial.default)
var app = new Vue({
    el: '#app',
    data: {
        status: "desconectado",
        user: null,
        users: [],
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
            app.user = name;
            socketClient.connect();
            apiClient.getChatData().then((data) => {
                app.users = data.users || [];
            });
        }
    },
})