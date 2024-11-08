/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  content: ["./ui/html/**/*.{html,tmpl}"],
  theme: {
    extend: {
            maxWidth: {
                '95vw' : '95vw',
            },
             height: {
                '1/4-screen': '25vh',
                '1/2-screen': '50vh',
              },
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
