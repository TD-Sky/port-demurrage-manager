import { defineComponent } from "vue";
import { bar_store } from "./bar";

export default defineComponent({
    setup() {
        const barStore = bar_store()

        return {
            barStore,
        }
    }
})
