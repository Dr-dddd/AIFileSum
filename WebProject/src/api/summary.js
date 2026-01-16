import axios from 'axios'

const BASE_URL = 'http://localhost:8080'

export async function uploadFile(file) {
  const formData = new FormData()
  formData.append('type', 'file')
  formData.append('file', file)

  const res = await axios.post(
    `${BASE_URL}/chat`,
    formData
  )

  return res.data
}
