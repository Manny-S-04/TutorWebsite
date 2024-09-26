/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  content: ["./ui/html/**/*.{html,tmpl}"],
  theme: {
    extend: {
            maxWidth: {
                '95vw' : '95vw',
            }
        },
  },
  plugins: [],
}

/**
 *
    safelist: [
        {
            pattern: /.*\
        }
    ]

 */ 
