Vue.component('layout', {

    template: //html
        `
    <div class="md-layout-item">
        <div class="md-layout">
            <div class="md-layout-item" style="background-color:red; height: 40px">
                <slot name="header"></slot>
            </div>
        </div>
        <div class="md-layout" style="height: calc(100vh - 65px)">
            <div class="md-layout-item md-size-80 md-small-size-60" style="background-color:green;">
                <slot></slot>
            </div>
            
            <div class="md-layout-item md-size-20 md-small-size-40" style="background-color:blue;">
                <slot name="right-panel"></slot>
            </div>
        </div>
        <div class="md-layout">
            <div class="md-layout-item" style="background-color:red; height: 25px" slot="footer">
                <slot name="footer"></slot>
            </div>
        </div>
    </div>
    `
})