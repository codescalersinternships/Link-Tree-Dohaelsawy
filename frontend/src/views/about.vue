<script setup lang="ts">

import Header from '@/components/header.vue'





</script>

<template>

    <div class="wrapper">
        <div class="content">
            <Header></Header>
            <div class="container">
                <div class="flex justify-center items-center gap-4 about">
                    <div>
                        <img src="/src/assets/me.jpg" alt="me" class="rounded-lg">
                    </div>
                    <div class="data">
                        <p>Hi, welcome! It's me, Doha, a recent graduate from Helwan University,<br>class of 2024.<br> This
                            project was developed during my internship at CodeScalers, <br> a company focused on helping
                            businesses, especially startups, <br> to build and operate their data centers and cloud
                            architectures.<br> I implemented the Linktree project using the following technologies:<br>
                            Backend: Go (Gin framework). <br> Frontend: Vue, CSS, and HTML. <br> Testing: Go's default
                            testing package for the backend <br>and Cypress for the frontend.<br> Storage: PostgreSQL, AWS
                            S3. <br>
                            Deployment: Docker, Docker Compose, Kubernetes, and Helm.<br> I really enjoyed building this
                            project and learned a lot along the way!<br> If you have any comments, suggestions, advice,
                            or
                            ideas on how to improve it,<br> please feel free to reach out to me on <br> dohaelsawy18@gmail.com.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>