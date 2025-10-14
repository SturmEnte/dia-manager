<script setup>
const props = defineProps({
	id: {
		type: String,
		required: true,
	},
	startedAt: {
		type: [String, Date],
		required: true,
	},
	endedAt: {
		type: [String, Date],
		required: false,
	},
});

function formatDate(value) {
	if (!value) return "-";

	const date = value instanceof Date ? value : new Date(value);

	if (isNaN(date.getTime())) return "-";

	const dd = String(date.getDate()).padStart(2, "0");
	const mm = String(date.getMonth() + 1).padStart(2, "0");
	const yyyy = date.getFullYear();
	const hh = String(date.getHours()).padStart(2, "0");
	const min = String(date.getMinutes()).padStart(2, "0");

	return `${dd}.${mm}.${yyyy} ${hh}:${min}`;
}
</script>

<template>
	<div class="catheter">
		<div class="side">
			<div><span class="attr-title">Gestartet:</span> {{ formatDate(props.startedAt) }}</div>
			<div><span class="attr-title">Beendet:</span> {{ formatDate(props.endedAt) }}</div>
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
	flex: 1;
}

.attr-title {
	font-weight: 600;
}
</style>
