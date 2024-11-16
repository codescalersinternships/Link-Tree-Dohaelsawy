<script setup lang="ts">
import { ref } from 'vue';
import { RouterLink, useRouter } from 'vue-router';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useAuthStore } from '@/stores/authStore'
type PAYLOAD = {
  password: string;
  email: string;
};
const form = ref<PAYLOAD>({
  password: '',
  email: '',
});
const router = useRouter();
const store = useAuthStore();
const onSubmit = async () => {
  try {
    await store.loginUser(form.value);
    router.push('/');
  } catch (error) {
    alert(error);
    console.error('Logout failed', error);
  } finally {
  }
};
</script>

<template>
  <div class="wrapper">

    <div class="content">
      <div class="flex flex-col justify-center items-center min-h-screen">
        <div class="mx-auto w-full max-w-md">
          <form @submit.prevent="onSubmit">
            <Card class="overflow-y-auto">
              <CardHeader class="space-y-1">
                <CardTitle class="text-2xl"> Welcome Back </CardTitle>
                <CardDescription> Enter your details below to login </CardDescription>
              </CardHeader>
              <CardContent class="grid gap-4">
                <div class="grid gap-2">
                  <Input id="email" type="text" placeholder="email" v-model="form.email" cy="login-email" />
                </div>
                <div class="grid gap-2">
                  <Input id="password" type="password" placeholder="password" v-model="form.password" cy="login-password"/>
                </div>
              </CardContent>
              <CardFooter class="flex-col space-y-2">
                <Button class="w-full" type="submit" cy="login-btn"> Login </Button>
                <p>
                  Don't have an account?
                  <RouterLink to="/auth/register"
                    class="border-b border-gray-500 text-muted-foreground hover:text-primary">
                    Register
                  </RouterLink>
                </p>
              </CardFooter>
            </Card>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>