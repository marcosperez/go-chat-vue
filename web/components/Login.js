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
            this.$emit('login', this.userName)
        }
    },
})