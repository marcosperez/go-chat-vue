var app = new Vue({
    el: '#app',
    data: {
        status: "offline"
    },
    mounted() {
        socketClient.connect();
    },
})