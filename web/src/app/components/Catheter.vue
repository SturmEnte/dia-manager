<script setup>
import { ref } from "vue";
import api from "../../services/api";

const CHANGE_REASONS = ["Unbekannt", "Andere", "Regulärer Wechsel", "Entzündung", "Langsame Insulinreaktion"];
let CHANGE_REASONS_REF = ref([]);
CHANGE_REASONS_REF.value = CHANGE_REASONS;

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
	changeReason: {
		type: Number,
		required: true,
	},
});

const emit = defineEmits(["deleted"]);

let editMode = ref();
editMode.value = false;

let start = ref();
let end = ref();

start.value = formatDateTimeLocal(props.startedAt);
end.value = formatDateTimeLocal(props.endedAt);

// Data before edit mode is entered
let oldStart;
let oldEnd;

// Toggle edit mode and discard unsaved changes
function toggleEditmode() {
	if (editMode.value) {
		start.value = oldStart;
		end.value = oldEnd;
	} else {
		oldStart = start.value;
		oldEnd = end.value;
	}

	editMode.value = !editMode.value;
}

async function saveChanges() {
	await api.updateCatheter(props.id, start.value, end.value);
	editMode.value = false;
}

async function deleteCatheter() {
	const res = await api.deleteCatheter(props.id);
	if (!res.id) return;
	emit("deleted", props.id);
}

function checkStartInput() {
	// Prevents a empty value to be tried to saved with would cause an error
	// Preventing the event is not possible for this type of event
	if (start.value == "") {
		start.value = oldStart;
	}
}
</script>

<template>
	<div class="catheter-container">
		<div class="catheter">
			<div class="side" v-if="!editMode">
				<div><span class="attr-title">Gestartet:</span> {{ formatDate(start) }}</div>
				<div><span class="attr-title">Beendet:</span> {{ formatDate(end) }}</div>
			</div>
			<div class="side" v-if="!editMode">
				<div><span class="attr-title">Tragedauer:</span> {{ formatDuration(props.startedAt, props.endedAt) }}</div>
				<div><span class="attr-title">Wechselgrund:</span> {{ CHANGE_REASONS_REF[props.changeReason < CHANGE_REASONS.length && props.changeReason >= 0 ? props.changeReason : 0] }}</div>
			</div>

			<!-- Edit catheter -->
			<div class="side" v-if="editMode">
				<div><span class="attr-title">Gestartet:</span> <input class="form-input" type="datetime-local" v-model="start" @input="checkStartInput" /></div>
				<div><span class="attr-title">Beendet:</span> <input class="form-input" type="datetime-local" v-model="end" :min="start || undefined" /></div>
			</div>
			<div class="side" v-if="editMode">
				<div><span class="attr-title">Tragedauer:</span> {{ formatDuration(start, end) }}</div>
				<div><span class="attr-title">Wechselgrund:</span> {{}}</div>
			</div>
		</div>
		<div class="buttons" v-if="editMode">
			<button @click="saveChanges">Save</button>
			<button @click="deleteCatheter">Delete</button>
		</div>
		<!-- <div class="catheter" v-if="editMode"></div> -->
		<div class="editButton" @click="toggleEditmode()">Edit</div>
	</div>
</template>

<style scoped>
.catheter-container {
	position: relative;
	display: flex;
	flex-direction: column;
}

.catheter {
	background: var(--col-3);
	border-radius: var(--radius);
	padding: var(--padding);
	display: flex;
	width: 100%;
}

.catheter .side {
	flex: 1;
}

.attr-title {
	font-weight: 600;
}

.form-input {
	background: var(--col-2);
}

.buttons {
	display: flex;
}

.buttons button {
	border: none;
	background: var(--col-2);
	color: var(--col-font);
	border-radius: var(--radius);
	font-size: 1rem;
	flex: 1;
	cursor: pointer;
}

.buttons button:first-of-type {
	margin-right: var(--padding);
}

.editButton {
	position: absolute;
	right: var(--padding);
	top: var(--padding);
	cursor: pointer;
	background: var(--col-2);
	padding: calc(var(--padding) / 2);
	border-radius: var(--radius);
	font-size: 0.9rem;
}
</style>

<script>
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
	if (!(start && end)) return "-";
	const startDate = start instanceof Date ? start : new Date(start);
	const endDate = end instanceof Date ? end : new Date(end);

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

function formatDateTimeLocal(value) {
	if (!value) return "";

	const date = value instanceof Date ? value : new Date(value);
	if (isNaN(date.getTime())) return "";

	// Format: YYYY-MM-DDTHH:mm
	const yyyy = date.getFullYear();
	const mm = String(date.getMonth() + 1).padStart(2, "0");
	const dd = String(date.getDate()).padStart(2, "0");
	const hh = String(date.getHours()).padStart(2, "0");
	const min = String(date.getMinutes()).padStart(2, "0");

	return `${yyyy}-${mm}-${dd}T${hh}:${min}`;
}
</script>
