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

	// Check if the date was defined successfully
	if (isNaN(date.getTime())) return "-";

	// Format date
	const dd = String(date.getDate()).padStart(2, "0");
	const mm = String(date.getMonth() + 1).padStart(2, "0");
	const yyyy = date.getFullYear();
	const hh = String(date.getHours()).padStart(2, "0");
	const min = String(date.getMinutes()).padStart(2, "0");

	return `${dd}.${mm}.${yyyy} ${hh}:${min}`;
}

function formatDuration(start, end) {
	if (!start) return "-";
	const startDate = start instanceof Date ? start : new Date(start);
	const endDate = end ? (end instanceof Date ? end : new Date(end)) : new Date();

	// Check if the date was defined successfully
	if (isNaN(startDate.getTime()) || isNaN(endDate.getTime())) return "-";

	// Time difference between start and end in ms
	let diffMs = endDate.getTime() - startDate.getTime();
	if (diffMs < 0) return "-";

	// Format the difference from ms to more human readble format (days, hours, minutes)
	const totalMinutes = Math.floor(diffMs / 60000);
	const days = Math.floor(totalMinutes / (60 * 24));
	const hours = Math.floor((totalMinutes % (60 * 24)) / 60);
	const minutes = totalMinutes % 60;

	const parts = [];
	if (days) parts.push(`${days}d`);
	if (hours) parts.push(`${hours}h`);
	parts.push(`${minutes}m`);

	return parts.join(" ");
}
</script>

<template>
	<div class="catheter">
		<div class="side">
			<div><span class="attr-title">Gestartet:</span> {{ formatDate(props.startedAt) }}</div>
			<div><span class="attr-title">Beendet:</span> {{ formatDate(props.endedAt) }}</div>
		</div>
		<div class="side">
			<div><span class="attr-title">Tragedauer:</span> {{ formatDuration(props.startedAt, props.endedAt) }}</div>
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
