Vue.component('usersList', {
    template: //html
        `
        <md-list  class="users-list">
            <div class="md-layout md-alignment-top-center" v-if="chatID">
                <div class="md-layout-item md-size-20">
                    <span class="md-list-item-text">{{chatID}}</span>
                </div>
            </div>
            <template v-for="u in users">
                <md-list-item   v-if="user.name != u.name">
                    <md-avatar>
                        <md-icon>home</md-icon>
                    </md-avatar>

                    <span class="md-list-item-text">{{u.name}}</span>

                    <md-button class="md-icon-button md-list-action">
                        <md-icon class="md-primary">chat_bubble</md-icon>
                    </md-button>
                </md-list-item>
            </template>
        </md-list>
    `,
    props: ["users", "user", "chatID"]
});
// <user-item></user-item>