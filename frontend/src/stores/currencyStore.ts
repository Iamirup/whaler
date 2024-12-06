import { defineStore } from 'pinia';

export const useCurrencyStore = defineStore('main', {
  state: () => ({
    currency: ''
  }),
  actions: {
    setCurency(value: string) {
      this.currency = value;
    }
  }
});