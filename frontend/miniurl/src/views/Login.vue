<script setup>
    function loginHandler(event) {
        login()
    }
</script>
<script>
    async function login() {
        let username = document.querySelector("#username").value
        let password = document.querySelector("#password").value

        let body = btoa(`${username}:${password}`)
        console.log("body: " , body)

        const response = await fetch("/api/v1/login", {
            method: "POST", // *GET, POST, PUT, DELETE, etc.
            headers: {
            // "Content-Type": "application/json",
            'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: body, // body data type must match "Content-Type" header
        });
        let responseLogin = await response.json();
        console.log('responseLogin: ' , responseLogin)
    }
</script>
<template>
    <div>
        <label for="username">Username</label>
        <input id="username" type="email" />
    </div>
    <div>
        <label for="password">Password</label>
        <input id="password" type="password" />
    </div>
    <div>
        <button @click="loginHandler">Login</button>
    </div>
</template>