import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'
import LanguageDectector from 'i18next-browser-languagedetector'
import { locales } from 'i18n/locales'

i18n
  .use(LanguageDectector)
  .use(initReactI18next)
  .init({
    resources: {
      en: {
        translations: Object.entries(locales).reduce((p, [k, v]) => {
          return {
            ...p,
            [k]: v.en,
          }
        }, {}),
      },
      zh: {
        translations: Object.entries(locales).reduce((p, [k, v]) => {
          return {
            ...p,
            [k]: v.cn,
          }
        }, {}),
      },
    },
    fallbackLng: 'zh',
    debug: false,
    ns: ['translations'],
    defaultNS: 'translations',

    keySeparator: false,
    interpolation: {
      escapeValue: false,
    },
  })

export default i18n
