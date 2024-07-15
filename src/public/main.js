const entryTemplate = document.getElementById("entry-template");

let entry = document.createElement("div");
entry.className = "entry";
entry.appendChild(entryTemplate.content.cloneNode(true));

document.body.appendChild(entry);
