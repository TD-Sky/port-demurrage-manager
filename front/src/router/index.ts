import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router"
const Layout = () => import("@/layout/index.vue")

/** 常驻路由 */
export const constantRoutes: Array<RouteRecordRaw> = [
    {
        path: "/redirect",
        component: Layout,
        meta: {
            hidden: true
        },
        children: [
            {
                path: "/redirect/:path(.*)",
                component: () => import("@/views/redirect/index.vue")
            }
        ]
    },

    {
        path: "/login",
        component: () => import("@/views/login/index.vue"),
        meta: {
            hidden: true
        }
    },

    {
        path: "/",
        redirect: "/warehouse",
    },

    {
        path: "/warehouse",
        component: Layout,
        children: [
            {
                path: "",
                component: () => import("@/views/warehouse/index.vue"),
                name: "Warehouse",
                meta: {
                    title: "场地管理",
                    icon: "warehouse",
                    affix: true,
                }
            }
        ]
    },

    {
        path: "/company",
        component: Layout,
        children: [
            {
                path: "",
                component: () => import("@/views/company/index.vue"),
                name: "Company",
                meta: {
                    title: "货代公司",
                    icon: "company",
                    affix: true,
                }
            }
        ]
    },

    {
        path: "/storage",
        component: Layout,
        children: [
            {
                path: "",
                component: () => import("@/views/storage/index.vue"),
                name: "Store",
                meta: {
                    title: "库存信息",
                    icon: "store",
                    affix: true,
                }
            }
        ]
    },

    {
        path: "/load",
        component: Layout,
        children: [
            {
                path: "",
                component: () => import("@/views/load/index.vue"),
                name: "Load",
                meta: {
                    title: "出库信息",
                    icon: "load",
                    affix: true
                }
            },

            {
                path: "post",
                component: () => import("@/views/load/PostForm.vue"),
                meta: {
                    hidden: true,
                },
            },
        ]
    },

    {
        path: "/statistics",
        component: Layout,
        children: [
            {
                path: "",
                component: () => import("@/views/statistics/index.vue"),
                name: "Statistics",
                meta: {
                    title: "统计图",
                    icon: "statistics",
                    affix: true,
                }
            }
        ]
    },
]

/**
 * 动态路由
 * 用来放置有权限（roles 属性）的路由
 * 必须带有 name 属性
 */
export const asyncRoutes: Array<RouteRecordRaw> = [
    {
        path: "/:pathMatch(.*)*", // 必须将 'ErrorPage' 路由放在最后, Must put the 'ErrorPage' route at the end
        component: Layout,
        redirect: "/404",
        name: "ErrorPage",
        meta: {
            title: "错误页面",
            icon: "404",
            hidden: true
        },
        children: [
            {
                path: "401",
                component: () => import("@/views/error-page/401.vue"),
                name: "401",
                meta: {
                    title: "401"
                }
            },
            {
                path: "404",
                component: () => import("@/views/error-page/404.vue"),
                name: "404",
                meta: {
                    title: "404"
                }
            }
        ]
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: constantRoutes
})

/** 重置路由 */
export function resetRouter() {
    // 注意：所有动态路由路由必须带有 name 属性，否则可能会不能完全重置干净
    try {
        router.getRoutes().forEach((route) => {
            const { name, meta } = route
            if (name && meta.roles?.length) {
                router.hasRoute(name) && router.removeRoute(name)
            }
        })
    } catch (error) {
        // 强制刷新浏览器，不过体验不是很好
        window.location.reload()
    }
}

export default router
