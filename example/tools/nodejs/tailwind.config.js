/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../../**/*.{css,templ,html}", "./node_modules/flowbite/**/*.js"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["night"],
  },
  plugins: [require("flowbite/plugin"), require("@tailwindcss/typography")],
};
