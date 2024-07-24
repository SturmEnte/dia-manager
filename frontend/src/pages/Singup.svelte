<script>
   // @ts-nocheck

   let username;
   let password;
   let error;

   function signup() {
      fetch("/api/auth/signup", {
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

         error.innerText = body.message;
         error.style.color = "green";
         window.location.href = "/#/login";
      });
   }
</script>

<main>
   <label for="username">Username</label>
   <input bind:this={username} type="text" id="username" />
   <label for="password">Password</label>
   <input bind:this={password} type="password" id="password" />
   <button on:click={signup}>Sign up</button>
   <a href="/#/login">Login</a>
   <div bind:this={error} id="error"></div>
</main>

<style>
   #error {
      color: red;
   }
</style>
