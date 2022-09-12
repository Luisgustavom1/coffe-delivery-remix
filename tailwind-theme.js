module.exports = {
   extend: {
      colors: {
        'yellow-light': '#F1E9C9',
        'yellow': '#DBAC2C',
        'yellow-dark': '#C47F17',
        'purple-light': '#EBE5F9',
        'purple': '#8047f8',
        'purple-dark': '#4B2995',
        'base-title': '#272221',
        'base-subtitle': '#403937',
        'base-text': '#574F4D',
        'base-label': '#8D8686',
        'base-hover': '#D7D5D5',
        'base-button': '#E6E5E5',
        'base-input': '#EDEDED',
        'base-card': '#F3F2F2',
        'background': '#FAFAFA',
      },
      screens: {
        xl: '1120px'
      }
    },
    fontFamily: {
      header: ["'Baloo 2'", 'cursive'],
      text: ["'Roboto'", 'sans-serif'],
    },
    typography: {
      regular: {
        DEFAULT: {
          fontFamily: 'Roboto',
          fontWeight: '400',
        },
        s: {
          fontSize: '14px',
          lineHeight: '18.2px'
        },
        m: {
          fontSize: '16px',
          lineHeight: '20.8px'
        },
        l: {
          fontSize: '20px',
          lineHeight: '26px'
        },
      },
      title: {
        DEFAULT: {
          fontFamily: 'Baloo 2',
          fontWeight: 800
        },
        l: {
          fontSize: '32px',
          lineHeight: '41.6px'
        },
        xl: {
          fontSize: '48px',
          lineHeight: '62.4px'
        }
      }
    },
}