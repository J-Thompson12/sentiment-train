formatted = "Ohh look at me i'm testing emojis i'm freaking AWESOME!!! 😀 😁 😂 🤣 😃 😄";

formatted = formatted.replace(/([\uD83C-\uDBFF\uDC00-\uDFFF]{2})/gm, ' $1');

console.log(formatted);
