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
	<div>
		<div id="historie" class="window">
			<div class="title">Historie</div>
			<div id="catheters"><Catheter v-for="catheter in catheters" :key="catheter.id" :id="catheter.id" :started-at="catheter.startedAt" :ended-at="catheter.endedAt" /></div>
		</div>
	</div>
</template>

<style scoped></style>
