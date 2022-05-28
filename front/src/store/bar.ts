import { defineStore } from "pinia";

export const bar_store = defineStore("bar", {
    state: () => {
        return {
            collapsed: true,
        }
    }
});
