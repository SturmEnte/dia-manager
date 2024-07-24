<script>
   // @ts-nocheck

   let username;
   let password;
   let error;

   function login() {
      fetch("/api/auth/login", {
         method: "post",
         headers: {
            "Content-Type": "application/json",
         },
         body: JSON.stringify({
            username: username.value,
            password: password.value,
         }),
      }).then(async (res) => {
         if (!res.headers.has("Content-Type") || !res.headers.get("Content-Type").includes("application/json")) {
            error.innerText = "Error in server response";
            return;
         }

         const body = await res.json();

         if (body.error) {
            error.innerText = "Error: " + body.error;
            return;
         }

         localStorage.setItem("token", body.access_token);
         document.cookie = "loggedIn=true;path=/";

         error.innerText = "Success. Redirecting shortly...";
         error.style.color = "green";

         setTimeout(() => {
            window.location.href = "/#/";
         }, 1000);
      });
   }
</script>

<main>
   <label for="username">Username</label>
   <input bind:this={username} type="text" id="username" />
   <label for="password">Password</label>
   <input bind:this={password} type="password" id="password" />
   <button on:click={login}>Login</button>
   <a href="/#/signup">Sign up</a>
   <div bind:this={error} id="error"></div>
</main>

<style>
   #error {
      color: red;
   }
</style>
