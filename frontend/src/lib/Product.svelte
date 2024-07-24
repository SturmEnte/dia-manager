<script>
   // Props
   export let name = "Catheter";
   export let databaseId = "catheter";
   export let index;

   let hidden = index === 0 ? false : true;
</script>

<main {databaseId} {hidden}>{name}</main>

<style>
   main {
      width: 100%;
      height: 100%;
      box-sizing: border-box;
   }
</style>
