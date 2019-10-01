Vue.component('login', {
    template: //html 
        `
        <div  class="login-form" >
                <md-field class=" animated heartBeat delay-2s">
                    <label>Nombre de usuarios</label>
                    <md-input v-model="userName"  v-on:keyup.enter="login"></md-input>
                </md-field>
                <div class="md-layout-item">
                <md-field>
                <md-select v-model="room" name="room" id="room" placeholder="Sala">
                    <md-option value="global">Global</md-option>
                </md-select>
                </md-field>
            </div>
                <md-button class="md-raised md-primary" @click="login">Login</md-button>
        </div>
    `,
    props: ["user"],
    data: function () {
        return {
            userName: this.user,
            room: "global"
        }
    },
    mounted() {
        this.userName = "marcos";
        this.login();
    },
    methods: {
        login: function () {
            var self = this;
            apiClient.login(this.userName, "").then(function (data) {
                self.$emit('login', data.name, this.room)
            });
        }
    },
})