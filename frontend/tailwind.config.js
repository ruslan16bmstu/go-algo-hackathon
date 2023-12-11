/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    fontFamily: {
      sans: ['Noto Sans', 'Noto Sans SC', 'sans-serif'],
      serif: ['Noto Serif', 'serif']
    },
    extend: {},
  },
  plugins: [],
}

