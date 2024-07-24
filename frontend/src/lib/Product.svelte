<script>
   // Props
   export let name = "Catheter";
   export let databaseId = "catheter";
   export let index;

   let hidden = index === 0 ? false : true;
</script>

<main {databaseId} {hidden}>{name}</main>

<style></style>
