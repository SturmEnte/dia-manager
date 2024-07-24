<script>
   // Code from old language selector
   // To be implemented

   /*
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
*/
</script>

<main>
   <select>
      <option value="test">🏳️‍🌈</option>
   </select>
</main>

<style>
   select {
      background: var(--secondary);
      height: var(--title-size);
      width: var(--title-size);

      border-radius: 100%;
      text-align: center;
      cursor: pointer;

      -webkit-appearance: none;
      -moz-appearance: none;
      appearance: none;
      border: none;
      font-size: calc(var(--title-size) * 0.6);
   }

   select,
   option {
      font-family: "Noto Color Emoji", sans-serif;
   }
</style>
