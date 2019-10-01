Vue.component('app-header', {
    template: //html
        `
    <div class="md-layout">
        <div class="md-layout-item" md-medium-size-75>Usuario: {{user}}</div>
        <div class="md-layout-item" md-medium-size-25>
        {{status}}
        </div>
    </div>
    `,
    props: ["status", "user"]
})