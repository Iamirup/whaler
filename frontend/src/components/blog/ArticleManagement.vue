<template>
<div class="p-6 bg-gray-100 min-h-screen">
    <div class="flex justify-between items-center mb-6">
        <h2 class="text-3xl font-semibold text-gray-700">Manage Articles</h2>
        <button @click="newArticle" class="px-4 py-2 bg-blue-600 text-white rounded-md shadow-md hover:bg-blue-700">Add Article</button>
    </div>
    <div v-if="editingArticle" class="mb-6">
        <input v-model="editingArticle.title" placeholder="Title" class="border p-2 mb-3 w-full rounded-md focus:outline-none focus:border-blue-500" />
        <textarea v-model="editingArticle.content" placeholder="Content" class="border p-2 w-full rounded-md focus:outline-none focus:border-blue-500"></textarea>
        <button @click="saveArticle" class="px-4 py-2 bg-green-600 text-white rounded-md shadow-md hover:bg-green-700 mt-2">Save</button>
    </div>
    <div v-for="article in articles" :key="article.article_id" class="mb-4 p-4 bg-white rounded-lg shadow-lg">
        <h3 class="text-xl font-semibold text-gray-800">{{ article.title }}</h3>
        <p class="text-gray-600">{{ article.content }}</p>
    <div class="mt-4 flex space-x-2">
        <button @click="editArticle(article)" class="px-4 py-2 bg-yellow-500 text-white rounded-md shadow-md hover:bg-yellow-600">Edit</button>
        <button @click="deleteArticle(article.article_id)" class="px-4 py-2 bg-red-600 text-white rounded-md shadow-md hover:bg-red-700">Delete</button>
    </div>
    </div>
</div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios, { type AxiosRequestConfig } from 'axios';
import { refreshJWT } from '../refreshJWT';

export default defineComponent({
    setup() {
        const articles = ref<Array<{ 
            article_id: string; 
            title: string; 
            content: string; 
            author_id: string; 
            author_username: string;
            likes: number;
            date: string;
        }>>([]);

        const editingArticle = ref<{ 
            article_id?: string; 
            title: string; 
            content: string; 
        } | null>(null);

        const fetchMyArticles = async () => {
            return await axios.get('/api/blog/v1/my-articles?limit=20')
                .then(response => {
                    articles.value = response.data.articles;
                })
                .catch(async error => {
                    if (error.response.data.need_refresh){
                        const isRefreshed = await refreshJWT(); 
                        if (!isRefreshed) { return; }
                        fetchMyArticles();
                    } else {
                        console.error('Failed to get articles', error);
                    }
                });
        };

        const newArticle = () => {
            editingArticle.value = { 
                article_id: '', 
                title: '', 
                content: '', 
            };
        };

        const saveArticle = async () => {
            if (editingArticle.value) {
                const articleData = { article_id: editingArticle.value.article_id, title: editingArticle.value.title, content: editingArticle.value.content };
                if (editingArticle.value.article_id) {
                // Update existing article
                return await axios.patch(`/api/blog/v1/article`, articleData)
                    .then(() => {
                        fetchMyArticles();
                        editingArticle.value = null;
                    })
                    .catch(async error => {
                        if (error.response.data.need_refresh){
                            const isRefreshed = await refreshJWT(); 
                            if (!isRefreshed) { return; }
                            saveArticle();
                        } else {
                            console.error('Failed to update article', error);
                        }
                    });
                } else {
                // Add new article
                return await axios.post('/api/blog/v1/article', articleData)
                    .then(() => {
                        fetchMyArticles();
                        editingArticle.value = null;
                    })
                    .catch(async error => {
                        if (error.response.data.need_refresh){
                            const isRefreshed = await refreshJWT(); 
                            if (!isRefreshed) { return; }
                            saveArticle();
                        } else {
                            console.error('Failed to add article', error);
                        }
                    });
                }
            }
        };

        const editArticle = (article: { 
            article_id?: string; 
            title: string; 
            content: string; 
        }) => {
            editingArticle.value = { ...article };
        }; 

        const deleteArticle = async (id: string) => {
            const config: AxiosRequestConfig = { 
                headers: { 'Content-Type': 'application/json' }, 
                data: { article_id: id } 
            };
            return await axios.delete(`/api/blog/v1/article`, config)
                .then(() => {
                    fetchMyArticles();
                })
                .catch(async error => {
                    if (error.response.data.need_refresh){
                        const isRefreshed = await refreshJWT(); 
                        if (!isRefreshed) { return; }
                        deleteArticle(id);
                    } else {
                        console.error('Failed to delete article', error);
                    }
                });
        };

        onMounted(fetchMyArticles);

        return {
            articles,
            editingArticle,
            newArticle,
            saveArticle,
            editArticle,
            deleteArticle,
        };
    },
});
</script>

<style scoped>
</style>