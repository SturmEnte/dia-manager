<script setup>
const props = defineProps({
	id: {
		type: String,
		required: true,
	},
	startedAt: {
		type: Date,
		required: true,
	},
	endedAt: {
		type: Date,
		required: false,
	},
});
</script>

<template>
	<div class="catheter">
		<div>Gestartet {{ props.startedAt }}</div>
		<div v-if="endedAt">Beendet {{ props.endedAt }}</div>
	</div>
</template>

<style>
.catheter {
	background: blue;
	border-radius: var(--radius);
}
</style>
