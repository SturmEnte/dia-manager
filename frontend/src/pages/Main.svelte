<script>
   import LanguageSelector from "../lib/LanguageSelector.svelte";
   import Product from "../lib/Product.svelte";

   import LogoutButton from "../lib/LogoutButton.svelte";

   export let dictionaryManager;

   // The available products
   // This hard coded list can later be replaced with a custom system
   const PRODUCTS = [
      {
         name: "Catheters",
         databaseId: "catheter",
      },
      {
         name: "Sensors",
         databaseId: "sensor",
      },
   ];

   let prodcuts;

   function changeProduct(event) {
      let selectedProduct = event.srcElement.value;

      for (let element of products.children) {
         console.log(element.attributes.databaseId.value);
         console.log(selectedProduct);

         if (element.attributes.databaseId.value == selectedProduct) {
            element.hidden = false;
            continue;
         }
         element.hidden = true;
      }
   }
</script>

<main>
   <div id="header">
      <div id="main-title">Dia Manager</div>
      <div id="menu">
         <select id="product-selector" on:input={changeProduct}>
            {#each PRODUCTS as { name, databaseId }, index}
               <option value={databaseId}>{name}</option>
            {/each}
         </select>
         <div>
            <LanguageSelector />
         </div>
         <div>
            <LogoutButton />
         </div>
      </div>
   </div>

   <div id="products" bind:this={prodcuts}>
      {#each PRODUCTS as { name, databaseId }, index (databaseId)}
         <Product {name} {databaseId} {index} />
      {/each}
   </div>

   <div id="author">Made by SturmEnte with ❤️</div>
</main>

<style>
   :root {
      --margin: 3vw;
      --title-size: 3vw;
      --header-size: 4vw;
   }

   main {
      background: var(--background);
      padding: var(--margin);
      box-sizing: border-box;
      width: 100vw;
      height: 100vh;
   }

   #header {
      width: 100%;
      height: var(--header-size);
      margin-bottom: calc(var(--margin) / 2);
      display: flex;
      align-items: center;
      justify-content: space-between;
   }

   #main-title {
      font-size: var(--title-size);
      color: var(--primary);
   }

   #main-title:after {
      content: "";
      display: block;
      width: 100%;
      height: 1px;
      background-color: black;
      position: relative;
      bottom: 0;
      left: 0;
   }

   #menu {
      display: flex;
      align-items: center;
      justify-content: flex-end;
      height: 100%;
   }

   #menu select,
   #menu div {
      margin-left: 0.8vw;
   }

   #product-selector {
      background: var(--secondary);
      height: var(--title-size);
   }

   #product-selector {
      border-radius: 0.8vw;
      border: none;
      font-size: calc(var(--title-size) * 0.5);
   }

   #products {
      width: 100%;
      height: calc(100vh - 2 * var(--margin) - var(--header-size) - calc(var(--margin) / 2));
      overflow: hidden;
   }

   #author {
      position: absolute;
      width: 100vw;
      left: 0;
      bottom: 1vw;
      display: flex;
      align-items: center;
      justify-content: center;
      color: rgb(126, 126, 126);
      font-size: 0.95vw;
   }
</style>
