<template>
  <div class="container mx-auto p-6 bg-gray-100 min-h-screen">
    <div class="bg-white rounded-lg shadow-md p-4">
      <h1 class="text-3xl font-bold mb-4">Support Tickets</h1>

      <!-- Create New Ticket Form -->
      <form @submit.prevent="createTicket" class="mb-6">
        <div class="mb-4">
          <label for="title" class="block text-sm font-medium text-gray-700">Title</label>
          <input v-model="newTicket.title" id="title" type="text" placeholder="Title" class="input mb-2" required />
        </div>
        <div class="mb-4">
          <label for="content" class="block text-sm font-medium text-gray-700">Content</label>
          <textarea v-model="newTicket.content" id="content" placeholder="Content" class="textarea mb-2" required></textarea>
        </div>
        <button type="submit" class="btn-primary w-full">Create Ticket</button>
      </form>

      <!-- Ticket List -->
      <div class="space-y-4">
        <div v-for="ticket in tickets" :key="ticket.ticket_id" class="ticket-card bg-gray-50 p-4 rounded-lg shadow-md">
          <h2 class="text-xl font-bold">{{ ticket.title }}</h2>
          <p class="mt-1">{{ ticket.content }}</p>
          <p class="mt-1 text-sm text-gray-600">Status: <span :class="ticket.is_done ? 'text-green-600' : 'text-red-600'">{{ ticket.is_done ? 'Done' : 'Open' }}</span></p>
          <p class="mt-1 text-sm text-gray-600">Created by {{ ticket.username }} on {{ formatDate(ticket.date) }}</p>

          <!-- Reply to Ticket -->
          <div v-if="!ticket.is_done" class="mt-4">
            <textarea v-model="ticket.replyText" placeholder="Reply" class="textarea mb-2"></textarea>
            <button @click="replyToTicket(ticket.ticket_id, ticket.replyText ?? '')" class="btn-secondary">Reply</button>
          </div>

          <!-- Display Replies -->
          <div v-if="ticket.replyText" class="mt-4 bg-gray-200 p-2 rounded-lg">
            <p><strong>Reply:</strong> {{ ticket.replyText }}</p>
            <p class="text-sm text-gray-600">Replied on {{ formatDate(ticket.replyDate) }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';

export interface Ticket {
  ticket_id: string;
  user_id: string;
  username: string;
  title: string;
  content: string;
  date: string; // ISO 8601 string
  is_done: boolean;
  replyText?: string;
  replyDate?: string; // ISO 8601 string
}

export default defineComponent({
  name: 'SupportTicket',
  setup() {
    const tickets = ref<Ticket[]>([]);
    const newTicket = ref<Partial<Ticket>>({ title: '', content: '' });

    const fetchMyTickets = async () => {
      const response = await axios.get('/api/support/v1/tickets/me');
      tickets.value = response.data;
    };

    onMounted(fetchMyTickets);

    const createTicket = async () => {
      await axios.post('/api/support/v1/ticket/new', newTicket.value);
      newTicket.value = { title: '', content: '' };
      await fetchMyTickets();
    };

    const replyToTicket = async (ticketId: string, replyText: string) => {
      await axios.post('/api/support/v1/ticket/reply', { ticketId, replyText: replyText ?? '' });
      await fetchMyTickets();
    };

    const formatDate = (date?: string) => {
      return date ? new Date(date).toLocaleString() : '';
    };

    return { tickets, newTicket, createTicket, replyToTicket, formatDate };
  },
});
</script>

<style scoped>
.container {
  max-width: 800px;
}
.input, .textarea {
  display: block;
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  transition: border-color 0.3s ease;
}
.input:focus, .textarea:focus {
  border-color: #1f2937;
}
.btn-primary {
  display: inline-block;
  padding: 0.5rem 1rem;
  background-color: #1f2937;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}
.btn-primary:hover {
  background-color: #4b5563;
}
.btn-secondary {
  display: inline-block;
  padding: 0.5rem 1rem;
  background-color: #4b5563;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}
.btn-secondary:hover {
  background-color: #6b7280;
}
.ticket-card {
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 1rem;
}
</style>
