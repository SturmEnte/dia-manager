(async () => {
	const ENTRY_ATTRIBUTE_IDS = ["name", "ref", "lot", "date-of-manufacture", "expiry-date", "start-of-use", "end-of-use"];

	// Dictionary manager
	const dictionaryManager = new DictionaryManager(getCookieValue("language"));
	await dictionaryManager.init();

	// Language selector
	const languageSelector = new LanguageSelector(document.getElementById("language-selector"));
	languageSelector.addLanguageChangeListener(async (language) => {
		// Preload the language to the local storage
		await DictionaryManager.prototype.loadDictionary(language);

		document.cookie = `language=${language};path=/`;
		location.reload();
	});

	// Product selector
	const productSelector = document.getElementById("product-selector");
	const productsContainerElement = document.getElementById("products");

	let products = [];

	for (productElement of productsContainerElement.children) {
		if (productElement.id != productSelector.value) {
			productElement.hidden = true;
		}
		products.push({ id: productElement.id, element: productElement });
	}

	productSelector.oninput = () => {
		for (product of products) {
			console.log(product);

			if (product.id != productSelector.value) {
				product.element.hidden = true;
				console.log("hide");
				continue;
			}
			product.element.hidden = false;
		}
	};

	//Functions
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
