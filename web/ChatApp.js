Vue.use(VueMaterial.default)
var app = new Vue({
    el: '#app',
    data: {
        status: "desconectado",
        user: {
            name: null,
            id: null,
        },
        users: [],
        chatID: ""
    },
    mounted() {
        socketClient.onConnect = this.connected;
        socketClient.onDisconect = this.disconnect;
    },
    methods: {
        connected: (evt) => {
            app.status = "conectado";
        },
        disconnect: (evt) => {
            app.status = "desconectado";
        },
        login: function (user) {
            app.user = user;
            if (user.id) {
                socketClient.connect(user.id);
                apiClient.getChatMessage().then((data) => {
                    app.users = data.users || [];
                });
            }
        },
        selectchatID: function (chatID) {
            this.chatID = chatID;

        }
    },
})