const integerOnly = (text) => {
     return text.replace(/[^0-9]/g, '');
}

const resort = (data, field, asc) => {
     return data.sort((a, b) => {
          const aValue = getValueByPath(a, field);
          const bValue = getValueByPath(b, field);
          if (aValue === bValue) {
               return 0;
          }

          if (asc === true) {
               return aValue > bValue ? 1 : -1;     
          }

          return aValue > bValue ? -1 : 1;
     });
}

function getValueByPath(obj, path) {
     return path.split('.').reduce((acc, key) => (acc && acc[key] !== 'undefined' ? acc[key] : undefined), obj);
   }

export { integerOnly, resort }
