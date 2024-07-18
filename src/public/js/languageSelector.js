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

const languageSelector = document.createElement("select");
languageSelector.classList.add("language-selector");

for (language of LANGUAGES) {
	const languageOption = document.createElement("option");
	languageOption.value = language.name;
	languageOption.innerText = language.emoji;
	languageSelector.appendChild(languageOption);
}

// Set selected language to the language saved in the cookie
// If there is not saved language, select the default language
if (getCookieValue("language")) {
	languageSelector.value = getCookieValue("language");
} else {
	languageSelector.value = DEFAULT_LANGUAGE;
}

languageSelector.oninput = async () => {
	// Preload the language to the local storage
	await DictionaryManager.prototype.loadDictionary(languageSelector.value);

	document.cookie = `language=${languageSelector.value};path=/`;
	location.reload();
};

document.body.appendChild(languageSelector);
