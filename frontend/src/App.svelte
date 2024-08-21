<script lang="ts">
   import { onMount } from "svelte";

   import Router from "svelte-spa-router";
   import { wrap } from "svelte-spa-router/wrap";

   import Main from "./pages/Main.svelte";
   import Login from "./pages/Login.svelte";
   import Singup from "./pages/Singup.svelte";

   import DictionaryManager from "./util/DictionaryManager";

   const dictionaryManager: DictionaryManager = new DictionaryManager("/dictionaries", "en", undefined);
   let loaded = false;

   let mainPage;
   let loginPage;
   let signupPage;

   onMount(async () => {
      // Load dictionary manager
      await dictionaryManager.loadDictionaries();

      // Create components for the router
      mainPage = wrap({ component: Main, props: { dictionaryManager } });
      loginPage = wrap({ component: Login, props: { dictionaryManager } });
      signupPage = wrap({ component: Singup, props: { dictionaryManager } });

      loaded = true;
   });
</script>

{#if loaded}
   <Router
      routes={{
         "/": mainPage,
         "/login": loginPage,
         "/signup": signupPage,
      }}
   />
{/if}
