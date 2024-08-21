<script>
   import LogoutIconBlack from "../assets/icons/material/logout-black.svg";

   function logout() {
      const token = localStorage.getItem("token");

      localStorage.removeItem("token");

      // Try to delete the token from the database
      // After that, load login page
      fetch("/api/logout", {
         method: "DELETE",
         headers: {
            Authorization: token,
         },
      }).finally((err) => {
         window.location.href = "/#/login";
      });
   }
</script>

<div>
   <div id="logout-link" on:click={logout}>
      <img src={LogoutIconBlack} alt="Logout" />
   </div>
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
