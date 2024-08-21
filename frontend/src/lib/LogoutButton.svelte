<script>
   import LogoutIconBlack from "../assets/icons/material/logout-black.svg";
</script>

<div>
   <a href="/logout" id="logout-link">
      <img src={LogoutIconBlack} alt="Logout" />
   </a>
</div>

<style>
   #logout-link {
      background: var(--secondary);
      width: var(--title-size);
      height: var(--title-size);
      border-radius: 100%;
      text-align: center;
      cursor: pointer;
      display: flex;
      align-items: center;
      justify-content: center;
   }

   img {
      width: calc(var(--title-size) * 0.8);
      height: calc(var(--title-size) * 0.8);
   }
</style>
