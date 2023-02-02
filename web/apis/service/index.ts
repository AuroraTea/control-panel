import axios from 'axios'

const service = axios.create({
  baseURL: 'http://localhost:5222/api/', // both development and production
  timeout: 2200,
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
    'Content-Type': 'application/json',
  },
})

service.interceptors.response.use(
  (res) => res.data,
  (err) => {
    let isPreErr = true
    const stopPreErr = () => { isPreErr = false }

    setTimeout(() => {
      if (isPreErr) {
        console.log(err)
        window.nMessage.error(err?.response?.data || 'Backend Error')
      }
    })
    return Promise.reject({ ...err, stopPreErr })
  },
)

export default service
