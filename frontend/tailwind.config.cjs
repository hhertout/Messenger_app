/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      "./index.html",
      "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
      fontFamily: {
        sans: ['Inter', 'Avenir', 'Helvetica', 'Arial', 'sans-serif']
      },
      colors: {
        'dark': '#333',
        'blue': '#1fb6ff',
        'dark-blue': '#317399',
        'purple': '#7e5bef',
        'pink': '#ff49db',
        'orange': '#ff7849',
        'green': '#13ce66',
        'yellow': '#ffc82c',
        'gray-dark': '#273444',
        'gray': '#8492a6',
        'gray-light': '#d3dce6',
        'white': '#F1F1F1',
        'red' : '#F00'
      },
      container: {
        center: true,
      },
      extend: {},
    },
    plugins: [],
  }
  