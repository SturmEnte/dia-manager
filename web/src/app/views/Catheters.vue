<script setup>
import { onMounted, ref } from "vue";
import api from "../../services/api";

import Catheter from "../components/Catheter.vue";

const catheters = ref([]);

onMounted(async () => {
	catheters.value = await api.getCatheters();
});
</script>

<template>
	<div id="main">
		<div id="history" class="window">
			<div class="title">Historie</div>
			<div id="catheters"><Catheter v-for="catheter in catheters" :key="catheter.id" :id="catheter.id" :started-at="catheter.startedAt" :ended-at="catheter.endedAt" /></div>
		</div>
		<div id="create" class="window"></div>
	</div>
</template>

<style scoped>
#main {
	display: flex;
	flex-direction: row;
	height: 100%;
}

#main #history {
	flex: 2;
	margin-right: calc(var(--padding) / 2);
	height: 100%;
	display: flex;
	flex-direction: column;
}

#main #create {
	flex: 1;
	margin-left: calc(var(--padding) / 2);
}
</style>
