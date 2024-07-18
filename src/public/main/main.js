(async () => {
	const ENTRY_ATTRIBUTE_IDS = ["name", "ref", "lot", "date-of-manufacture", "expiry-date", "start-of-use", "end-of-use"];

	const dictionaryManager = new DictionaryManager(getCookieValue("language"));
	await dictionaryManager.init();

	const languageSelector = new LanguageSelector(document.getElementById("language-selector"));
	languageSelector.addLanguageChangeListener(async (language) => {
		// Preload the language to the local storage
		await DictionaryManager.prototype.loadDictionary(language);

		document.cookie = `language=${language};path=/`;
		location.reload();
	});

	function generateEntry(data) {
		let element = document.createElement("div");
		element.classList.add("entry");

		for (attributeId of ENTRY_ATTRIBUTE_IDS) {
			if (data[attributeId] == undefined) continue;

			let attributeElement = document.createElement("div");
			attributeElement.classList.add("attribute");
			attributeElement.innerText = dictionaryManager.get(attributeId) + data[attributeId];
			element.appendChild(attributeElement);
		}

		return element;
	}
})();
