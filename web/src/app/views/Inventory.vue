<script setup>
import { onMounted, ref } from "vue";
import api from "../../services/api";

import Structure from "../components/Structure.vue";

// Make inventory reactive so Vue updates the template when data arrives
const inventory = ref([]);

// Editor
const currentTab = ref("create-structure");

// Use a unique ID generator to prevent key collisions
const generateId = () => Date.now() + Math.random();

// Create Structure
const structureName = ref("");

const generallInformation = ref([]);
const attributes = ref([]);

function addGeneralInformation() {
	generallInformation.value.push({
		id: "gi" + generateId(),
		name: "",
		value: "",
	});
}

function addAttribute() {
	attributes.value.push({
		id: "att" + generateId(),
		name: "",
		required: false,
	});
}

addGeneralInformation();
addAttribute();

async function createStructure() {
	let generallInformationFormated = [];
	generallInformation.value.forEach((elem) => {
		generallInformationFormated.push({ name: elem.name, value: elem.value });
	});

	let attributesFormated = [];
	attributes.value.forEach((elem) => {
		attributesFormated.push({ name: elem.name, required: elem.required });
	});

	const id = await api.createItemStructures(structureName.value, generallInformationFormated, attributesFormated);

	if (!id) {
		alert("Failed");
		return;
	}

	// Insert new structure
	inventory.value.push({
		id: id,
		name: structureName.value,
		attributes: attributesFormated,
		generalInformation: generallInformationFormated,
		items: [],
	});

	// Clear
	structureName.value = "";
	generallInformation.value = [];
	attributes.value = [];

	addGeneralInformation();
	addAttribute();

	alert("Created");
}

// Create Item
const selectedStructure = ref("");
const attributesCreateItem = ref({});
const amount = ref(0);

function newStructureSelected() {
	attributesCreateItem.value = {};

	const structure = inventory.value.find((s) => s.id === selectedStructure.value);

	if (!structure) return;

	structure.attributes.forEach((attribute) => {
		attributesCreateItem.value[attribute.name] = "";
	});
}

async function createItems() {
	if (!selectedStructure.value) {
		alert("Bitte wählen Sie eine Struktur aus.");
		return;
	}

	if (amount.value <= 0) {
		alert("Anzahl muss größer als 0 sein.");
		return;
	}

	const structure = inventory.value.find((s) => s.id === selectedStructure.value);
	if (!structure) return;

	// Check required fields
	for (const attr of structure.attributes) {
		if (attr.required && !attributesCreateItem.value[attr.name]) {
			alert(`Das Attribute "${attr.name}" muss ausgefüllt werden.`);
			return;
		}
	}

	const itemsToCreate = [];
	for (let i = 0; i < amount.value; i++) {
		itemsToCreate.push({ ...attributesCreateItem.value });
	}

	const ids = await api.createItems(selectedStructure.value, itemsToCreate);

	if (!ids) {
		alert("Fehler beim Erstellen der Artikel.");
		return;
	}

	// Add to local inventory
	ids.forEach((id) => {
		structure.items.push({
			id: id,
			structure_id: selectedStructure.value,
			data: { ...attributesCreateItem.value },
			created_at: new Date().toISOString(),
		});
	});

	// Reset form
	attributesCreateItem.value = {};
	structure.attributes.forEach((attribute) => {
		attributesCreateItem.value[attribute.name] = "";
	});
	amount.value = 0;

	alert(`${ids.length} Artikel erstellt.`);
}

onMounted(async () => {
	// Map inventory
	const structures = await api.getItemStructures();

	for (let i = 0; i < structures.length; i++) {
		const structure = structures[i];

		// push objects into the reactive array and include the id on the value
		inventory.value.push({
			id: structure.id,
			name: structure.name,
			attributes: structure.attributes,
			generalInformation: structure.general_information,
			items: [],
		});
	}

	const items = await api.getItems();

	for (let i = 0; i < items.length; i++) {
		const item = items[i];

		const bucket = inventory.value.find((s) => s.id === item.structure_id);
		if (!bucket) {
			console.log("Couldn't find a structure for item with id", item.id);
			continue;
		}

		bucket.items.push(item);
	}
});
</script>

<template>
	<div id="main">
		<div id="structures" class="scrollbar">
			<Structure
				class="structure"
				v-for="structure in inventory"
				:key="structure.id"
				:id="structure.id"
				:name="structure.name"
				:items="structure.items"
				:attributes="structure.attributes"
				:generalInformation="structure.generalInformation"
				:selected="selectedStructure"
			/>
		</div>
		<div id="editor">
			<form id="selection" @input="test">
				<input type="radio" id="create-structure" value="create-structure" v-model="currentTab" />
				<label for="create-structure">Struktur erstellen</label>

				<input type="radio" id="create-item" value="create-item" v-model="currentTab" />
				<label for="create-item">Artikel erstellen</label>
			</form>

			<div class="editor-input scrollbar" v-if="currentTab === 'create-structure'">
				<label for="name">Name:</label>
				<input type="text" v-model="structureName" />

				<br /><br />

				<div>Allgemeine Informationen:</div>
				<div class="general-information-pair" v-for="(elem, index) in generallInformation" :key="elem.id">
					<input type="text" placeholder="Name/Titel" v-model="elem.name" />
					<input type="text" placeholder="Wert" v-model="elem.value" />
					<button v-if="index != 0" @click="generallInformation.splice(index, 1)">-</button>
				</div>
				<button id="add-general-information" @click="addGeneralInformation">+</button>

				<br /><br />

				<div>Attribute:</div>
				<div class="attribute" v-for="(elem, index) in attributes" :key="elem.id">
					<input type="text" placeholder="Name/Titel" v-model="elem.name" />
					<label for="required">Muss ausgefüllt werden?</label>
					<input type="checkbox" id="required" v-model="elem.required" />
					<button v-if="index != 0" @click="attributes.splice(index, 1)">-</button>
				</div>
				<button id="add-attribute" @click="addAttribute">+</button>

				<br /><br />

				<button id="create-structure" @click="createStructure">Erstellen</button>
			</div>

			<div class="editor-input" v-if="currentTab === 'create-item'">
				<label for="structur">Wähle eine Struktur aus:</label>
				<br />
				<select id="structur" v-model="selectedStructure" @change="newStructureSelected">
					<option v-for="structure in inventory" :key="structure.id" :value="structure.id">
						{{ structure.name }}
					</option>
				</select>

				<br /><br />

				<div class="" v-for="attribute in inventory.find((s) => s.id === selectedStructure)?.attributes || []" :key="attribute.name">
					<label :for="attribute.name">{{ attribute.name }}{{ attribute.required ? "*" : "" }}</label>
					<input type="text" :id="attribute.name" v-model="attributesCreateItem[attribute.name]" />
				</div>
				<div>*muss ausgefüllt werden</div>

				<br /><br />

				<input type="number" v-model.number="amount" />

				<br /><br />

				<button @click="createItems">Artikel erstellen</button>
			</div>
		</div>
	</div>
</template>

<style scoped>
#main {
	display: flex;
	flex-direction: row;
	height: 100%;
	overflow: hidden;
}

#editor {
	background: var(--col-2);
	flex: 1;
	margin-left: var(--padding);
	border-radius: var(--radius);
	padding: var(--padding);

	display: flex;
	flex-direction: column;
}

#selection {
	widows: 100%;
	display: flex;
	flex-direction: row;
	align-items: center;
}

#selection input {
	display: none;
}

#selection label {
	background: var(--col-3);
	flex: 1;
	text-align: center;
	border-radius: var(--radius);
	margin-right: var(--padding);
	padding: var(--padding);
	user-select: none;
}

#selection label:last-of-type {
	margin: 0;
}

#selection input:checked + label {
	color: var(--col-accent);
	font-weight: bold;
}

.editor-input {
	padding: var(--padding);
	overflow-y: auto;
}

#structures {
	flex: 2;
	overflow-y: auto;
}

.structure {
	margin-bottom: var(--padding);
}

.structure:last-child {
	margin: 0;
}
</style>
