/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../../**/*.{css,templ,html}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["night"],
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
};
