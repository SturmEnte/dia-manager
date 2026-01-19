<script setup>
const props = defineProps({
	id: {
		type: String,
		required: true,
	},
	name: {
		type: String,
		required: true,
	},
	attributes: {
		type: Object,
		required: true,
	},
	items: {
		type: Array,
		required: true,
	},
});
</script>

<template>
	<div id="structure">
		<div class="row">
			<div class="attr-name">Name:</div>
			<div>{{ props.name }}</div>
		</div>
		<div class="row">
			<div class="attr-name">Menge:</div>
			<div>{{ props.items.length }}</div>
		</div>
		<div class="row">
			<div class="attr-name">Name:</div>
			<div>{{ props.name }}</div>
		</div>
	</div>
</template>

<style scoped>
#structure {
	background: var(--col-2);
	border-radius: var(--radius);
	display: flex;
	flex-direction: column;
	width: 100%;
	padding: var(--padding);
	box-sizing: border-box;
}

.row {
	display: flex;
	flex-direction: row;
	align-items: center;
	justify-content: space-between;
	width: 100%;
}
</style>
