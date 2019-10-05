Vue.component('app-header', {
    template: //html
        `
        <div class="md-layout  md-gutter">
            <div class="md-layout-item  md-size-80 md-small-size-60">
                
            </div>

            <div class="md-layout-item  md-size-20 md-small-size-40 ">
                <md-field>
                    <md-select 
                    v-model="channelID" 
                    name="channelID" 
                    id="channelID" 
                    placeholder="Sala" 
                    @md-selected="selectchannelID"
                    >
                        <md-option value="global">Global</md-option>
                        <md-option value="channelID 1">channelID1</md-option>
                    </md-select>
                </md-field>
            </div>
        </div>
    `,
    props: ["status", "user"],
    data: function () {
        return {
            channelID: "global"
        }
    },
    mounted() {
        this.$emit("select-channelID", this.channelID)
    },
    methods: {
        selectchannelID: function () {
            this.$emit("select-channelID", this.channelID)
        }
    },
})