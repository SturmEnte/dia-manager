<script>
	import LogoutIconBlack from "../assets/icons/material/logout-black.svg";

	import LanguageSelector from "../lib/LanguageSelector.svelte";
</script>

<main>
	<div id="header">
		<div id="main-title">Dia Manager</div>
		<div id="menu">
			<select id="product-selector">
				<option value="catheters">Catheters</option>
				<option value="sensors">Sensors</option>
			</select>
			<div>
				<LanguageSelector />
			</div>
			<a href="/logout" id="logout-link">
				<img src={LogoutIconBlack} alt="Logout" />
			</a>
		</div>
	</div>

	<div id="products">
		<div id="catheters" class="product">Catheters</div>
		<div id="sensors" class="product">Sensors</div>
	</div>

	<div id="author">Made by SturmEnte with ❤️</div>
</main>

<style>
	:root {
		--margin: 3vw;
		--title-size: 3vw;
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
	#menu a,
	#menu div {
		margin-left: 0.8vw;
	}

	#logout-link,
	#product-selector {
		background: var(--secondary);
		height: var(--title-size);
	}

	#logout-link {
		width: var(--title-size);

		border-radius: 100%;
		text-align: center;
		cursor: pointer;
	}

	#product-selector {
		border-radius: 0.8vw;
		border: none;
		font-size: calc(var(--title-size) * 0.5);
	}

	#logout-link {
		display: flex;
		align-items: center;
		justify-content: center;
	}

	#logout-link img {
		width: calc(var(--title-size) * 0.8);
		height: calc(var(--title-size) * 0.8);
	}

	#products {
		background: blue;
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
