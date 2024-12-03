import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

// 定义路由
const routes: Array<RouteRecordRaw> = [
    {
        path: "/index",
        name: "index",
        component: () => import("../views/index.vue"), // 懒加载组件
    }
];

// 创建路由实例
const index = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL), // 使用 HTML5 历史模式
    routes,
});

export default index;
