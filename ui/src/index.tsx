import ReactDOM from 'react-dom/client'
import Router from './route'
import reportWebVitals from './reportWebVitals'
import 'i18n'
import './index.css'
import axios from 'axios'
import { Toast } from '@douyinfe/semi-ui'
import Cookies from 'js-cookie'
import { useNavigate } from 'react-router-dom'

axios.interceptors.request.use(
  async (config) => {
    const token = Cookies.get('auth')
    config.headers = {
      Authorization: `Bearer ${token}`,
    }
    return config
  },
  (error) => {
    Promise.reject(error)
  },
)

axios.interceptors.response.use(
  (res) => {
    if (res.config.method !== 'get') {
      Toast.success(`[${res.config.method}]${res.config.url}, status: ${res.status}`)
    }
    return res
  },
  (err) => {
    if (err.response) {
      Toast.error(
        `[${err.response.config.method}]${err.response.config.url}, status: ${err.response.status}, data: ${err.response.data} `,
      )
      if (err.response.status === 401) {
        Cookies.remove('username')
        Cookies.remove('auth')
      }
      if (Cookies.get('username') === undefined || Cookies.get('auth') === undefined) {
        const nav = useNavigate()
        nav('/login')
      }
      return Promise.reject(err.response.data)
    }

    if (err.request) {
      Toast.error(`err.request = ${err.request}`)
      return Promise.reject(err.request)
    }

    return Promise.reject(err.message)
  },
)

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
root.render(<Router />)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
