const plugin = require('tailwindcss/plugin')
const theme = require('./tailwind-theme')

module.exports = {
  content: ["./app/**/*.{ts,tsx,jsx,js}"],
  theme: theme,
  plugins: [
    plugin(function({ addComponents, addUtilities }) {
      addComponents(Object.entries(theme.typography).reduce((accTypography, [key, values]) => {
        const defaultStyle = values.DEFAULT;
        delete values.DEFAULT;

        return Object.entries(values).reduce((accVariants, [variants, styles]) => ({
            ...accVariants, 
            [`.typography-${key}-${variants}`]: {
                ...defaultStyle,
                ...styles,
            }
        }), {
            ...accTypography
        })
      }, {}))

      addUtilities({
        '.button-active': {
          'box-shadow': 'inset 0px -1px 20px 2px rgba(128, 71, 248, 30%)'
        }
      })
      }
    )
  ]
};