export default class DictionaryManager {
   private PREFIX: string = "[Dictionary Manager] ";

   private dictionaries: Map<string, any>;

   private dictionariesPath: string;
   private defaultDictionary: string; // The default / fallback dictionary name
   private selectedLanguage: string;

   constructor(dictionariesPath: string, defaultDictionary: string, selectedLanguage: string | undefined) {
      this.dictionariesPath = dictionariesPath;
      this.defaultDictionary = defaultDictionary;

      if (selectedLanguage) {
         this.selectedLanguage = selectedLanguage;
      } else {
         this.selectedLanguage = defaultDictionary;
      }

      // Initialize private variables
      this.dictionaries = new Map<string, any>();
   }

   public async loadDictionaries(): Promise<void> {
      try {
         // Fetch config
         const configResponse = await fetch(this.dictionariesPath + "/config.json");
         const config = await configResponse.json();

         // Fetch languages
         for (const language of config.languages) {
            console.log(this.PREFIX + "Trying to load:", language);

            const dictionaryResponse = await fetch(this.dictionariesPath + `/${language}.json`);
            const dictionary = await dictionaryResponse.json();

            this.dictionaries.set(language, dictionary);

            console.log(this.PREFIX + "Successfully loaded:", language);
         }
      } catch (error) {
         console.error("Error while loading dictionaries:", error);
      }
   }

   public selectLanguage(language: string): void {
      this.selectedLanguage = language;
   }

   public getEntry(page: string, key: string): string {
      // Return key from selected language's dictionary if it exists
      const selectedDict = this.dictionaries.get(this.selectedLanguage);

      if (selectedDict && selectedDict[page] && selectedDict[page][key]) {
         return selectedDict[page][key];
      }

      // Otherwise try the same for the default dictionary
      const defaultDict = this.dictionaries.get(this.defaultDictionary);

      if (defaultDict && defaultDict[page] && defaultDict[page][key]) {
         return defaultDict[page][key];
      }

      // If both dictionaries do not have the page / key, then an error message is returned
      console.error(`${this.PREFIX}Could not find key "${key}" of page "${page}"`);
      return `Could not find key "${key}" of page "${page}"`;
   }
}
