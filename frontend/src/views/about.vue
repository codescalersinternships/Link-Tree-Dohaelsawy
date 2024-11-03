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
                        <img src="/src/assets/me.jpg" alt="me" class="rounded-lg" cy="me-img">
                    </div>
                    <div class="data text-left">
                        <p>Hi, welcome! It's me, Doha, a recent graduate from Helwan University,<br>class of 2024.<br> This
                            project was developed during my internship at <strong>CodeScalers</strong>, <br> a company focused on helping
                            businesses, especially startups, <br> to build and operate their data centers and cloud
                            architectures.<br> I implemented the Linktree project using the following technologies:<br>
                            <b>Backend</b>: Go (Gin framework). <br> <b>Frontend</b>: Vue, CSS, and HTML. <br> <b>Testing</b>: Go's default
                            testing package for the backend <br>and Cypress for the frontend.<br> <b>Storage</b>: PostgreSQL, AWS
                            S3. <br>
                            <b>Deployment</b>: Docker, Docker Compose, Kubernetes, and Helm.<br> I really enjoyed building this
                            project and learned a lot along the way!<br> If you have any comments, suggestions, advice,
                            or ideas on how to improve it, please feel free to reach out to me on <br> dohaelsawy18@gmail.com.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>