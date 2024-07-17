// This class contains the dictionary manager
// The dictionary manager is responsible for loading the right dictionary and serving the dictionary
// with a fallback dictionary if a dictionary is not complete
class DictionaryManager {
	DEFAULT_LANGUAGE = "en";

	defaultDictionary;
	dictionary;

	async init() {
		// Load the fallback dictionary
		this.defaultDictionary = await this.loadDictionary(this.DEFAULT_LANGUAGE);

		const language = getCookieValue("language");

		// Load the normal dictionary
		if (language) {
			this.dictionary = await this.loadDictionary(language);
		}
	}

	async loadDictionary(language) {
		if (localStorage.getItem(language)) {
			return JSON.parse(localStorage.getItem(language));
		}
		const response = await fetch(`/${language}.dict.json`);
		const dictionary = await response.json();
		localStorage.setItem(language, JSON.stringify(dictionary));
		return this.dictionary;
	}

	get(key) {}
}
