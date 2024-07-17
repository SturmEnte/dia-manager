const ENTRY_ATTRIBUTE_IDS = ["name", "ref", "lot", "produced-date", "expiration-date", "usage-start", "usage-end"];

const DICTIONARY = {
	name: "Name: ",
	ref: "REF-Nummer: ",
	lot: "LOT-Nummer: ",
	"date-of-manufacture": "Date of Manufacture: ",
	"expiry-date": "Expiry Date: ",
	"start-of-use": "Start of use: ",
	"end-of-use": "End of use: ",
};

//document.body.appendChild(generateEntry({ name: "Foo", ref: "Bar" }));

function generateEntry(data) {
	let element = document.createElement("div");
	element.classList.add("entry");

	for (attributeId of ENTRY_ATTRIBUTE_IDS) {
		if (data[attributeId] == undefined) continue;

		let attributeElement = document.createElement("div");
		attributeElement.classList.add("attribute");
		attributeElement.innerText = DICTIONARY[attributeId] + data[attributeId];
		element.appendChild(attributeElement);
	}

	return element;
}

//#region Template
// const entryTemplate = document.getElementById("entry-template");

// let entry = document.createElement("div");
// entry.className = "entry";
// let clone = entryTemplate.content.cloneNode(true);

// console.log(clone);

// entry.appendChild(clone);

// document.body.appendChild(entry);
//#endregion
