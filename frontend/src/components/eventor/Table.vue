<template>
    <main class="mt-6 mx-3 xs:mx-5 sm:mx-14"> 
        <!-- Live Transactions Section -->
        <section class="container mx-0 mt-96 xl:mt-12 px-0 py-6 text-end">
        <h2 class="text-lg font-semibold mb-4">Live <span class="text-green-500">●</span></h2>
        <div class="bg-white shadow-md rounded-lg overflow-x-auto">

            <table class="min-w-full">
                <thead class="bg-red-400">
                    <tr>
                        <th class="px-4 py-2 text-left text-xs font-medium text-white drop-shadow-[0_1.4px_1.4px_rgba(0,0,0,0.8)]">Hash</th>
                        <th class="px-4 py-2 text-left text-xs font-medium text-white drop-shadow-[0_1.4px_1.4px_rgba(0,0,0,0.8)]">Block</th>
                        <th class="px-4 py-2 text-left text-xs font-medium text-white drop-shadow-[0_1.4px_1.4px_rgba(0,0,0,0.8)]">AGE</th>
                        <th class="px-4 py-2 text-left text-xs font-medium text-white drop-shadow-[0_1.4px_1.4px_rgba(0,0,0,0.8)]">Type</th>
                        <th class="px-4 py-2 text-left text-xs font-medium text-white drop-shadow-[0_1.4px_1.4px_rgba(0,0,0,0.8)]">From</th>
                        <th class="px-4 py-2 text-left text-xs font-medium text-white drop-shadow-[0_1.4px_1.4px_rgba(0,0,0,0.8)]">To</th>
                        <th class="px-4 py-2 text-left text-xs font-medium text-white drop-shadow-[0_1.4px_1.4px_rgba(0,0,0,0.8)]">Token</th>
                        <th class="px-4 py-2 text-left text-xs font-medium text-white drop-shadow-[0_1.4px_1.4px_rgba(0,0,0,0.8)]">RES</th>
                    </tr>
                </thead>
                <tbody class="bg-white divide-y divide-slate-50">
                    <tr>
                        <td class="px-4 py-3 text-sm text-gray-700">● e9c0ff5++28d6</td>
                        <td class="px-4 py-3 text-sm text-gray-700">12345678</td>
                        <td class="px-4 py-3 text-sm text-gray-700">12 secs ago</td>
                        <td class="px-4 py-3 text-sm text-gray-700">Transfer</td>
                        <td class="px-4 py-3 text-sm text-gray-700">$ 1,250.615</td>
                        <td class="px-4 py-3 text-sm text-gray-700">TKNcBq...7nGwtVV</td>
                        <td class="px-4 py-3 text-sm text-gray-700">USD Token</td>
                        <td class="px-4 py-3 text-sm text-green-600">OK</td>
                    </tr>
                    <tr class="bg-red-100">
                        <td class="px-4 py-3 text-sm text-gray-700">● e9c0ff5++28d6</td>
                        <td class="px-4 py-3 text-sm text-gray-700">12345678</td>
                        <td class="px-4 py-3 text-sm text-gray-700">12 secs ago</td>
                        <td class="px-4 py-3 text-sm text-gray-700">Transfer</td>
                        <td class="px-4 py-3 text-sm text-red-600">$ 15,290.615</td>
                        <td class="px-4 py-3 text-sm text-gray-700">TKNcBq...7nGwtVV</td>
                        <td class="px-4 py-3 text-sm text-gray-700">USD Token</td>
                        <td class="px-4 py-3 text-sm text-green-600">OK</td>
                    </tr>
                    <tr>
                    <td class="px-4 py-3 text-sm text-gray-700">● e9c0ff5++28d6</td>
                    <td class="px-4 py-3 text-sm text-gray-700">12345678</td>
                    <td class="px-4 py-3 text-sm text-gray-700">12 secs ago</td>
                    <td class="px-4 py-3 text-sm text-gray-700">Transfer</td>
                    <td class="px-4 py-3 text-sm text-gray-700">$ 1,250.615</td>
                    <td class="px-4 py-3 text-sm text-gray-700">TKNcBq...7nGwtVV</td>
                    <td class="px-4 py-3 text-sm text-gray-700">USD Token</td>
                    <td class="px-4 py-3 text-sm text-green-600">OK</td>
                </tr>
                <tr>
                    <td class="px-4 py-3 text-sm text-gray-700">● e9c0ff5++28d6</td>
                    <td class="px-4 py-3 text-sm text-gray-700">12345678</td>
                    <td class="px-4 py-3 text-sm text-gray-700">12 secs ago</td>
                    <td class="px-4 py-3 text-sm text-gray-700">Transfer</td>
                    <td class="px-4 py-3 text-sm text-gray-700">$ 1,250.615</td>
                    <td class="px-4 py-3 text-sm text-gray-700">TKNcBq...7nGwtVV</td>
                    <td class="px-4 py-3 text-sm text-gray-700">USD Token</td>
                    <td class="px-4 py-3 text-sm text-green-600">OK</td>
                </tr>
                <tr>
                <td class="px-4 py-3 text-sm text-gray-700">● e9c0ff5++28d6</td>
                <td class="px-4 py-3 text-sm text-gray-700">12345678</td>
                <td class="px-4 py-3 text-sm text-gray-700">12 secs ago</td>
                <td class="px-4 py-3 text-sm text-gray-700">Transfer</td>
                <td class="px-4 py-3 text-sm text-gray-700">$ 1,250.615</td>
                <td class="px-4 py-3 text-sm text-gray-700">TKNcBq...7nGwtVV</td>
                <td class="px-4 py-3 text-sm text-gray-700">USD Token</td>
                <td class="px-4 py-3 text-sm text-green-600">OK</td>
                </tr>
                <tr class="bg-red-100">
                <td class="px-4 py-3 text-sm text-gray-700">● e9c0ff5++28d6</td>
                <td class="px-4 py-3 text-sm text-gray-700">12345678</td>
                <td class="px-4 py-3 text-sm text-gray-700">12 secs ago</td>
                <td class="px-4 py-3 text-sm text-gray-700">Transfer</td>
                <td class="px-4 py-3 text-sm text-red-600">$ 15,290.615</td>
                <td class="px-4 py-3 text-sm text-gray-700">TKNcBq...7nGwtVV</td>
                <td class="px-4 py-3 text-sm text-gray-700">USD Token</td>
                <td class="px-4 py-3 text-sm text-green-600">OK</td>
            </tr>
            <tr>
            <td class="px-4 py-3 text-sm text-gray-700">● e9c0ff5++28d6</td>
            <td class="px-4 py-3 text-sm text-gray-700">12345678</td>
            <td class="px-4 py-3 text-sm text-gray-700">12 secs ago</td>
            <td class="px-4 py-3 text-sm text-gray-700">Transfer</td>
            <td class="px-4 py-3 text-sm text-gray-700">$ 1,250.615</td>
            <td class="px-4 py-3 text-sm text-gray-700">TKNcBq...7nGwtVV</td>
            <td class="px-4 py-3 text-sm text-gray-700">USD Token</td>
            <td class="px-4 py-3 text-sm text-green-600">OK</td>
            </tr>
                </tbody>
            </table>
        </div>
        </section> 
    </main>
</template>

<script setup lang="ts">

</script>

<style>

</style>