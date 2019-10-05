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
        channelID: ""
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
                apiClient.getChannelMessage().then((data) => {
                    app.users = data.users || [];
                });
            }
        },
        selectchannelID: function (channelID) {
            this.channelID = channelID;

        }
    },
})