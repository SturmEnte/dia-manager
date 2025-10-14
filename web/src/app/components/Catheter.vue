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
		<div class="side">
			<div><span class="attr-title">Gestartet:</span> {{ props.startedAt }}</div>
			<div><span class="attr-title">Beendet:</span> {{ props.endedAt }}</div>
		</div>
		<div class="side">
			<div><span class="attr-title">Tragedauer:</span> {{ null }}</div>
			<div><span class="attr-title">Wechselgrund:</span> {{ null }}</div>
		</div>
	</div>
</template>

<style>
.catheter {
	background: var(--col-3);
	border-radius: var(--radius);
	padding: var(--padding);
	display: flex;
}

.catheter .side {
	width: 50%;
}

.attr-title {
	font-weight: 600;
}
</style>
