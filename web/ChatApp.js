Vue.use(VueMaterial.default)
var app = new Vue({
    el: '#app',
    data: {
        status: "desconectado",
        user: null,
        users: [],
        roomName: ""
    },
    mounted() {
        socketClient.onConnect = this.connected;
        socketClient.onDisconect = this.disconnect;
    },
    methods: {
        connected: (evt) => {
            app.status = "conectado";
            socketClient.subscribe('global');
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
        },
        selectRoom: function (roomName) {
            this.roomName = roomName;

        }
    },
})