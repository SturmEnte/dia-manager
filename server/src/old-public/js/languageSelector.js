const LANGUAGES = [
	{
		name: "en",
		emoji: "🇬🇧",
	},
	{
		name: "de",
		emoji: "🇩🇪",
	},
];

const DEFAULT_LANGUAGE = "en";

// languageSelector.oninput = async () => {
// 	// Preload the language to the local storage
// 	await DictionaryManager.prototype.loadDictionary(languageSelector.value);

// 	document.cookie = `language=${languageSelector.value};path=/`;
// 	location.reload();
// };

class LanguageSelector {
	element;
	languageChangeListeners = [];

	constructor(element) {
		if (element) {
			this.element = element;
		} else {
			this.element = document.createElement("select");
			document.body.appendChild(languageSelector);
		}

		this.element.classList.add("language-selector");

		for (let language of LANGUAGES) {
			const languageOption = document.createElement("option");
			languageOption.value = language.name;
			languageOption.innerText = language.emoji;
			this.element.appendChild(languageOption);
		}

		// Set selected language to the language saved in the cookie
		// If there is not saved language, select the default language
		if (getCookieValue("language")) {
			this.element.value = getCookieValue("language");
		} else {
			this.element.value = DEFAULT_LANGUAGE;
		}

		this.element.oninput = () => {
			this.selectorInputEventHandler();
		};
	}

	addLanguageChangeListener(listener) {
		console.log(this.languageChangeListeners);
		this.languageChangeListeners.push(listener);
		console.log(this.languageChangeListeners);
	}

	selectorInputEventHandler() {
		console.log(this);
		console.log(this.languageChangeListeners);
		this.languageChangeListeners.forEach((listener) => {
			listener(this.element.value);
		});
	}
}
