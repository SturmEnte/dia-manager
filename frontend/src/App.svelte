<script lang="ts">
   import { onMount } from "svelte";

   import Router from "svelte-spa-router";

   import Main from "./pages/Main.svelte";
   import Login from "./pages/Login.svelte";
   import Singup from "./pages/Singup.svelte";

   import DictionaryManager from "./util/DictionaryManager";

   const dictionaryManager: DictionaryManager = new DictionaryManager("/dictionaries", "en", undefined);

   onMount(async () => {
      await dictionaryManager.loadDictionaries();

      console.log(dictionaryManager.getEntry("login", "login"));
   });
</script>

<Router
   routes={{
      "/": Main,
      "/login": Login,
      "/signup": Singup,
   }}
/>
