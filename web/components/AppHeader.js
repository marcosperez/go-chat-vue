Vue.component('app-header', {
    template: //html
        `
        <div class="md-layout  md-gutter">
            <div class="md-layout-item  md-size-80 md-small-size-60">
                
            </div>

            <div class="md-layout-item  md-size-20 md-small-size-40 ">
                <md-field>
                    <md-select 
                    v-model="roomName" 
                    name="roomName" 
                    id="roomName" 
                    placeholder="Sala" 
                    @md-selected="selectRoom"
                    >
                        <md-option value="global">Global</md-option>
                        <md-option value="room 1">ROOM1</md-option>
                    </md-select>
                </md-field>
            </div>
        </div>
    `,
    props: ["status", "user"],
    data: function () {
        return {
            roomName: "global"
        }
    },
    mounted() {
        this.$emit("select-room", this.roomName)
    },
    methods: {
        selectRoom: function () {
            this.$emit("select-room", this.roomName)
        }
    },
})