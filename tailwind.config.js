const plugin = require('tailwindcss/plugin')
const theme = require('./tailwind-theme')

module.exports = {
  content: ["./app/**/*.{ts,tsx,jsx,js}"],
  theme: theme,
  plugins: [
    plugin(({ addComponents }) => 
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
    )
  ]
};