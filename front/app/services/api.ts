import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:3004/api'
});

export default api