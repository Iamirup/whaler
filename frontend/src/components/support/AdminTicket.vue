<template>
    <div class="mt-30 container mx-auto p-6 bg-gray-100 min-h-screen">
      <div class="bg-white rounded-lg shadow-md p-4">
        <h1 class="text-3xl font-bold mb-4">All Support Tickets</h1>
  
        <!-- Ticket List -->
        <div class="space-y-4">
          <div v-for="ticket in tickets" :key="ticket.ticket_id" class="ticket-card bg-gray-50 p-6 rounded-lg shadow-md">
            <div class="flex justify-between items-center">
              <h2 class="text-2xl font-bold">{{ ticket.title }}</h2>
              <span :class="ticket.is_done ? 'text-green-600' : 'text-red-600'">{{ ticket.is_done ? 'Done' : 'Open' }}</span>
            </div>
            <p class="mt-2">{{ ticket.content }}</p>
            <p class="mt-2 text-sm text-gray-600">Created by {{ ticket.username }} on {{ formatDate(ticket.date) }}</p>
  
            <!-- Reply to Ticket -->
            <div v-if="!ticket.is_done" class="mt-4">
              <label for="reply" class="block text-sm font-medium text-gray-700">Reply</label>
              <textarea v-model="ticket.reply_text" id="reply" placeholder="Reply" class="textarea mb-2"></textarea>
              <button @click="replyToTicket(ticket.ticket_id, ticket.reply_text ?? '')" class="btn-secondary">Reply</button>
            </div>
  
            <!-- Display Replies -->
            <div v-if="ticket.reply_text" class="mt-4 bg-gray-200 p-4 rounded-lg">
              <p><strong>Reply:</strong> {{ ticket.reply_text }}</p>
              <p class="text-sm text-gray-600">Replied on {{ formatDate(ticket.reply_date) }}</p>
            </div>
          </div>
        </div>

        <!-- Load More Button -->
        <div class="flex justify-center mt-6">
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
name: 'AdminTickets',
setup() {
    const router = useRouter()
  
    const tickets = ref<Ticket[]>([]);
    const cursor = ref<string | null>(null);
    const limit = ref<number>(20); 
  
    const fetchAllTickets = async (eCursor: string | null, limit: number): Promise<void> => {
    try {
        const response = await axios
          .get('/api/support/v1/tickets/all', {
            params: { cursor: eCursor, limit: limit },
          });
        tickets.value = [...tickets.value, ...response.data.tickets];
        cursor.value = response.data.new_cursor;
      } catch (error: any) {
        if (error.response.data.need_refresh){
          const isRefreshed = await refreshService.refreshJWT(); 
          if (!isRefreshed) { router.push('/login'); return; }
          fetchAllTickets(eCursor, limit);
        } else {
          alertService.showAlert(error.response.data.errors[0].message, "error");
          console.error(error);
        }
      }
    };

    const replyToTicket = async (ticketId: string, reply_text: string): Promise<void> => {
    try {
        await axios.post('/api/support/v1/ticket/reply', { ticket_id: ticketId, reply_text: reply_text ?? '' });
        tickets.value = [];
        cursor.value = null;
        return await fetchAllTickets(null, limit.value);
      } catch (error: any) {
        if (error.response.data.need_refresh){
          const isRefreshed = await refreshService.refreshJWT(); 
          if (!isRefreshed) { router.push('/login'); return; }
          replyToTicket(ticketId, reply_text);
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
      fetchAllTickets(cursor.value, limit.value);
    };

    onMounted(() => {
    fetchAllTickets(null, limit.value);
    });

    return { tickets, replyToTicket, formatDate, loadMoreTickets };
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
