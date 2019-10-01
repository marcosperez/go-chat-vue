Vue.component('login', {
    template: //html 
        `
        <div  class="login-form" >
                <md-field class=" animated heartBeat delay-2s">
                    <label>Nombre de usuarios</label>
                    <md-input v-model="userName"  v-on:keyup.enter="login"></md-input>
                </md-field>
                <md-button class="md-raised md-primary" @click="login">Login</md-button>
        </div>
    `,
    props: ["user"],
    data: function () {
        return {
            userName: this.user,
        }
    },
    methods: {
        login: function () {
            var self = this;
            fetch(`${config.baseUrl}/users`, {
                method: 'POST', // or 'PUT'
                body: JSON.stringify({ name: this.userName }), // data can be `string` or {object}!
                headers: {
                    "Content-Type": "application/json"
                }
            })
                .then(function (response) {
                    return response.json();
                })
                .then(function (data) {
                    console.log(data);
                    self.$emit('login', data.name)
                });
        }
    },
})