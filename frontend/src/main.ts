import { createApp } from 'vue'
import {createPinia} from 'pinia'
import router from "./router"
import App from './App.vue'

// Import base css
import "normalize.css"
import "./css/variables.css"
import "./css/core.css"

// Import header css
import "./css/header/header.css"

// Import footer css
import "./css/footer/footer.css"

// Import editor css
import "./css/editor/editor.css"
import "./css/editor/edfields.css"

// Import vue-multiselect styles
import "vue-multiselect/dist/vue-multiselect.css"
import "./css/multiselect/multiselect.css"

// Import AG-Grid styles
import "ag-grid-community/dist/styles/ag-grid.css"
import "ag-grid-community/dist/styles/ag-theme-alpine.css"
import "./css/ag-grid-styles.css"

createApp(App)
.use(createPinia())
.use(router)
.mount('#app')
