<script setup lang="ts">
import { defineModel } from 'vue';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';

interface Props {
  name?: string;
  url?: string;
}
defineProps<Props>();

const isOpen = defineModel('isOpen', {
  required: true,
  type: Boolean,
});
const emit = defineEmits(['onClose']);
</script>

<template>
  <Dialog :open="isOpen" @update:open="emit('onClose')">
    <DialogContent>
      <DialogHeader>
        <DialogTitle v-if="name"> {{ name }} </DialogTitle>
        <DialogContent v-if="url">
          {{ url }}
        </DialogContent>
      </DialogHeader>

      <slot />
    </DialogContent>
  </Dialog>
</template>