/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./html/**/*.templ",
    "./cmd/**/*.go",
    "./static/**/*.html",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
