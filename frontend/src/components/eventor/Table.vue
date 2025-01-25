<template>
    <main class="mt-6 mx-3 xs:mx-5 sm:mx-14"> 
        <!-- Live Transactions Section -->
        <section class="container mx-0 mt-96 xl:mt-12 px-0 py-6 text-end">
            <div class="flex items-center justify-between mb-4">
                <button @click="showModal = true" class="bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 text-white py-2 px-6 rounded-lg shadow-md transition duration-200 ease-in-out transform hover:scale-105">Filter</button>
                <h2 class="text-lg font-semibold">Live <span class="text-green-500">●</span></h2>
            </div>
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
                        <tr v-for="transaction in transactions" :key="transaction.hash">
                            <td class="px-4 py-3 text-sm text-gray-700">● {{ transaction.hash }}</td>
                            <td class="px-4 py-3 text-sm text-gray-700">{{ transaction.block_id }}</td>
                            <td class="px-4 py-3 text-sm text-gray-700">{{ transaction.time }}</td>
                            <td class="px-4 py-3 text-sm text-gray-700">{{ transaction.type }}</td>
                            <td class="px-4 py-3 text-sm text-gray-700">{{ transaction.output_total_usd }}</td>
                            <td class="px-4 py-3 text-sm text-gray-700">{{ transaction.token }}</td>
                            <td class="px-4 py-3 text-sm text-green-600">{{ transaction.res }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </section> 

        <!-- Filter Modal -->
        <div v-if="showModal" class="fixed inset-0 flex items-center justify-center bg-gray-500 bg-opacity-75">
            <div class="bg-white p-6 rounded-lg text-left shadow-lg w-96">
                <h2 class="text-lg font-semibold mb-4">Filter Transactions</h2>
                <div class="mb-4">
                    <label class="block text-sm font-medium">Cryptocurrency</label>
                    <select v-model="filters.cryptocurrency" class="border-gray-300 rounded-md w-full p-2 mt-2">
                        <option value="">Select cryptocurrency</option>
                        <option value="bitcoin">bitcoin</option>
                        <option value="ethereum">ethereum</option>
                        <option value="dogecoin">dogecoin</option> 
                    </select>
                </div>
                <div class="mb-4">
                    <label class="block text-sm font-medium">Minimum Amount</label>
                    <input v-model="filters.minAmount" class="border-gray-300 rounded-md w-full p-2 mt-2" type="number" placeholder="Enter minimum amount"/>
                </div>
                <div class="mb-4">
                    <label class="block text-sm font-medium">Age</label>
                    <div class="flex items-center mt-2">
                        <span class="text-sm">Last</span>
                        <input v-model="filters.age.number" class="border-gray-300 rounded-md w-16 p-2 mx-2" type="number" placeholder="Number"/>
                        <select v-model="filters.age.unit" class="border-gray-300 rounded-md p-2">
                            <option value="seconds">Seconds</option>
                            <option value="minutes">Minutes</option>
                            <option value="hours">Hours</option>
                            <option value="days">Days</option>
                        </select>
                    </div>
                </div>
                <div class="flex justify-end">
                    <button @click="applyFilters" class="bg-gradient-to-r from-blue-400 to-blue-600 hover:from-blue-500 hover:to-blue-700 text-white py-2 px-4 rounded-md mr-2">Apply Filters</button>
                    <button @click="showModal = false" class="bg-gray-300 text-gray-700 py-2 px-4 rounded-md">Close</button>
                </div>
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import axios, { AxiosError } from 'axios';

interface FilterParams {
  cryptocurrency: string;
  minAmount: number;
  age: string;
}

interface AgeFilter {
  number: number;
  unit: 'seconds' | 'minutes' | 'hours' | 'days';
}

const transactions = ref<Transaction[]>([]);
const filters = ref({
    cryptocurrency: '',
    minAmount: 0,
    age: {
        number: 0,
        unit: 'seconds' as const
    }
});

const showModal = ref(false);

const convertToSeconds = (number: number, unit: string): number => {
    const conversions = {
        'seconds': 1,
        'minutes': 60,
        'hours': 3600,
        'days': 86400
    };
    return number * conversions[unit as keyof typeof conversions];
};

interface ApiResponse {
    table: Transaction[];
    new_cursor: string;
}

interface Transaction {
    block_id: number;
    output_total_usd: number;
    hash: string;
    time: string;
    res: string;
    token: string;
    type: string;
}

const fetchTransactions = async () => {
    try {
        const params: FilterParams = {
            cryptocurrency: filters.value.cryptocurrency,
            minAmount: filters.value.minAmount,
            age: convertToSeconds(filters.value.age.number, filters.value.age.unit).toString()
        };
        const response = await axios.get<ApiResponse>('https://api.example.com/transactions', { params });
        transactions.value = response.data.table;
    } catch (error) {
        if (error instanceof AxiosError) {
            console.error('Error fetching transactions:', error.message);
            if (error.response) {
                console.error('Server Error:', error.response.status);
            }
        }
    }
};

const applyFilters = async () => {
    await fetchTransactions();
    showModal.value = false;
};

let intervalId: number;

onMounted(() => {
    fetchTransactions();
    intervalId = window.setInterval(fetchTransactions, 5000); 
});

onUnmounted(() => {
    window.clearInterval(intervalId); 
});

</script>

<style scoped>
</style>