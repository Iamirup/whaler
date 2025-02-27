<template>
  <div class="pt-36">
    <div class="mx-24 mb-12 bg-white rounded-lg shadow-md p-4 min-h-screen">
      <h1 class="text-3xl font-bold mb-4 text-center">My Support Tickets</h1>

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
        <div v-for="ticket in tickets" :key="ticket.ticket_id" class="ticket-card bg-gray-50 p-4 break-words rounded-lg shadow-md">
          <h2 class="text-xl font-bold">{{ ticket.title }}</h2>
          <p class="mt-1">{{ ticket.content }}</p>
          <p class="mt-1 text-sm text-gray-600">Status: <span :class="ticket.is_done ? 'text-green-600' : 'text-red-600'">{{ ticket.is_done ? 'Done' : 'Open' }}</span></p>
          <p class="mt-1 text-sm text-gray-600">Created by {{ ticket.username }} on {{ formatDate(ticket.date) }}</p>

          <!-- Display Replies -->
          <div v-if="ticket.reply_text" class="mt-4 bg-gray-200 p-2 rounded-lg">
            <p><strong>Reply:</strong> {{ ticket.reply_text }}</p>
            <p class="text-sm text-gray-600">Replied on {{ formatDate(ticket.reply_date) }}</p>
          </div>
        </div>
      </div>

      <!-- Load More Button -->
      <div class="flex justify-center mt-4">
        <button @click="loadMoreTickets" class="btn-secondary">More</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';
import { alertService } from '../alertor';
import { useRouter } from 'vue-router';
import { refreshService } from '../refreshJWT';

export interface Ticket {
  ticket_id: string;
  user_id: string;
  username: string;
  title: string;
  content: string;
  date: string;
  is_done: boolean;
  reply_text?: string;
  reply_date?: string;
}

export default defineComponent({
  name: 'UserTickets',
  setup() {
    const router = useRouter()

    const tickets = ref<Ticket[]>([]);
    const newTicket = ref<Partial<Ticket>>({ title: '', content: '' });
    const cursor = ref<string | null>(null);
    const limit = ref<number>(5);

    const fetchMyTickets = async (eCursor: string | null, limit: number): Promise<void> => {
      try {
        const response = await axios
          .get('/api/support/v1/tickets/me', {
            params: { cursor: eCursor, limit: limit },
          });
        tickets.value = [...tickets.value, ...response.data.tickets];
        cursor.value = response.data.new_cursor;
      } catch (error: any) {
        if (error.response.data.need_refresh){
          const isRefreshed = await refreshService.refreshJWT(); 
          if (!isRefreshed) { router.push('/login'); return; }
          await fetchMyTickets(eCursor, limit);
        } else {
          alertService.showAlert(error.response.data.errors[0].message, "error");
          console.error(error);
        }
      }
    };

    const createTicket = async (): Promise<void> => {
      try {
        await axios.post('/api/support/v1/ticket/new', newTicket.value);
        newTicket.value = { title: '', content: '' };
        tickets.value = [];
        cursor.value = null;
        return await fetchMyTickets(null, limit.value);
      } catch (error: any) {
        if (error.response.data.need_refresh){
          const isRefreshed = await refreshService.refreshJWT(); 
          if (!isRefreshed) { router.push('/login'); return; }
          await createTicket();
        } else {
          alertService.showAlert(error.response.data.errors[0].message, "error");
          console.error(error);
        }
      }
    };

    const formatDate = (date?: string): string => {
      return date ? new Date(date).toLocaleString() : '';
    };

    const loadMoreTickets = (): void => {
      fetchMyTickets(cursor.value, limit.value);
    };

    onMounted(() => {
      fetchMyTickets(null, limit.value);
    });

    return { tickets, newTicket, createTicket, formatDate, loadMoreTickets };
  },
});
</script>

<style scoped>
.container {
  max-width: 1000px;
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
