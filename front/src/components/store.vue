<script setup lang="ts">
import { onMounted } from 'vue';
import axios from 'axios';
import Layout from '../layout/index.vue';

type GetStore = {
    id: number,
    store_date: string,
    license_plate_number: string,
    stocks: number,
    store_ton: number,
    duration: number
    fee: number,
}

let posts: GetStore[];

onMounted(() => {
    axios.get("/store").then((response) => {
        posts = response.data
        console.log(posts)
    })
})
</script>

<template>
    <Layout>
        <div v-for="post in posts" :key="post.id">
            <h2>{{ post.id }}</h2>
            <p>{{ post.store_date }}</p>
        </div>
    </Layout>
</template>
