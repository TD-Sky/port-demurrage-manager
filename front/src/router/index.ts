import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            redirect: "/store",
        },
        {
            path: "/store",
            component: () => import("../components/store.vue")
        },
        {
            path: "/load",
            component: { template: '<div>Load</div>' },
        },
        {
            path: "/fee",
            component: { template: '<div>Fee</div>' },
        },
    ],
})
