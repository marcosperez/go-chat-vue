Vue.component('app-header', {
    template: //html
        `
        <div class="md-layout  md-gutter">
            <div class="md-layout-item  md-size-80 md-small-size-60">
                
            </div>

            <div class="md-layout-item  md-size-20 md-small-size-40 ">
                <md-field>
                    <md-select 
                    v-model="chatID" 
                    name="chatID" 
                    id="chatID" 
                    placeholder="Sala" 
                    @md-selected="selectchatID"
                    >
                        <md-option value="global">Global</md-option>
                        <md-option value="chatID 1">chatID1</md-option>
                    </md-select>
                </md-field>
            </div>
        </div>
    `,
    props: ["status", "user"],
    data: function () {
        return {
            chatID: "global"
        }
    },
    mounted() {
        this.$emit("select-chatID", this.chatID)
    },
    methods: {
        selectchatID: function () {
            this.$emit("select-chatID", this.chatID)
        }
    },
})