// This class contains the dictionary manager
// The dictionary manager is responsible for loading the right dictionary and serving the dictionary
// with a fallback dictionary if a dictionary is not complete
class DictionaryManager {
	DEFAULT_LANGUAGE = "en";

	defaultDictionary;
	dictionary;

	language;

	constructor(language) {
		this.language = language;
	}

	async init() {
		// Load the fallback dictionary
		this.defaultDictionary = await this.loadDictionary(this.DEFAULT_LANGUAGE);

		// Load the normal dictionary
		if (this.language) {
			this.dictionary = await this.loadDictionary(this.language);
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

	updateDictionary(language) {
		// TBD
	}

	get(key) {
		if (this.dictionary[key] !== undefined) return this.dictionary[key];
		return this.defaultDictionary[key];
	}
}
