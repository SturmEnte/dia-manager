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
		type: Array,
		required: true,
	},
	generalInformation: {
		type: Array,
		required: true,
	},
	items: {
		type: Array,
		required: true,
	},
	selected: {
		type: String,
		required: true,
	},
});
</script>

<template>
	<div id="container">
		<div id="structure" :style="{ background: props.selected == props.id ? 'rgb(255,0,0)' : '' }">
			<div class="info">
				<div class="attr-name">Name:</div>
				<div>{{ props.name }}</div>
			</div>
			<div class="info">
				<div class="attr-name">Menge:</div>
				<div>{{ props.items.length }}</div>
			</div>

			<div class="info" v-for="information in props.generalInformation" :key="information.name">
				<div class="attr-name">{{ information.name }}:</div>
				<div>{{ information.value }}</div>
			</div>
		</div>
		<div id="items">
			<div class="item" v-for="item in props.items" :key="item.id">
				<div class="info">
					<div class="attr-name">ID:</div>
					<div>{{ item.id }}</div>
				</div>
				<div class="info">
					<div class="attr-name">Structure ID:</div>
					<div>{{ item.structure_id }}</div>
				</div>
				<div class="info">
					<div class="attr-name">Created At:</div>
					<div>{{ item.created_at }}</div>
				</div>
				<div class="info" v-for="(value, key) in item.data" :key="key">
					<div class="attr-name">{{ key }}:</div>
					<div>{{ value }}</div>
				</div>
			</div>
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

.info {
	display: flex;
	flex-direction: row;
	align-items: center;
	justify-content: space-between;
}

#items {
	padding: var(--padding);
	box-sizing: border-box;
	border-radius: var(--radius);
	background: var(--col-3);
	display: flex;
	flex-direction: column;
	gap: var(--padding);
}

.item {
	background: var(--col-2);
	border-radius: var(--radius);
	padding: var(--padding);
	box-sizing: border-box;
}
</style>
