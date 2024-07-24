<script lang="ts">
   import Router from "svelte-spa-router";
   import Main from "./pages/Main.svelte";
   import Login from "./pages/Login.svelte";
   import Singup from "./pages/Singup.svelte";
</script>

<Router
   routes={{
      "/": Main,
      "/login": Login,
      "/signup": Singup,
   }}
/>
