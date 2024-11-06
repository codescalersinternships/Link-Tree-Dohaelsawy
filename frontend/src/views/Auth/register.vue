<script setup lang="ts">
import { ref } from "vue";
import { RouterLink, useRouter } from "vue-router";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useAuthStore } from "@/stores/authStore";
type PAYLOAD = {
  first_name: string;
  last_name: string;
  username: string;
  email: string;
  password: string;
};
const confirm_password = "";
const form = ref<PAYLOAD>({
  password: "",
  username: "",
  first_name: "",
  last_name: "",
  email: "",
});
const router = useRouter();
const store = useAuthStore();
const onSubmit = async () => {
  try {
    await store.registerUser(form.value);
    router.push("/auth/login");
  } catch (error) {
    alert(error);
    console.error("register failed", error);
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
                <CardTitle class="text-2xl"> Create an account </CardTitle>
              </CardHeader>
              <CardContent class="grid gap-4">
                <div class="grid gap-2">
                  <Input
                    id="first_name"
                    type="text"
                    placeholder="first name"
                    v-model="form.first_name"
                  />
                </div>
                <div class="grid gap-2">
                  <Input
                    id="last_name"
                    type="text"
                    placeholder="last name"
                    v-model="form.last_name"
                  />
                </div>
                <div class="grid gap-2">
                  <Input
                    id="email"
                    type="text"
                    placeholder="email"
                    v-model="form.email"
                  />
                </div>
                <div class="grid gap-2">
                  <Input
                    id="username"
                    type="text"
                    placeholder="username"
                    v-model="form.username"
                  />
                </div>
                <div class="grid gap-2">
                  <Input
                    id="password"
                    type="password"
                    placeholder="password"
                    v-model="form.password"
                    cy="password"
                  />
                </div>
              </CardContent>
              <CardFooter class="flex-col space-y-2">
                <Button class="w-full" type="submit" cy="register-btn"> Register </Button>
                <p>
                  Already have an account?
                  <RouterLink
                    to="/auth/login"
                    class="border-b border-gray-500 text-muted-foreground hover:text-primary"
                    cy="login-btn"
                  >
                    Login
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
