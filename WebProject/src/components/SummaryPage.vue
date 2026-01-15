<template>
  <div class="container">
    <h2>📄 文件总结助手</h2>

    <input type="file" @change="onFileChange" />
    <button @click="upload">上传并总结</button>

    <div v-if="loading">正在总结中...</div>

    <div v-if="summary">
      <h3>总结结果</h3>
      <pre>{{ summary }}</pre>
    </div>

    <div v-if="error" class="error">
      {{ error }}
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { uploadFile } from '../api/summary'

const file = ref(null)
const summary = ref('')
const loading = ref(false)
const error = ref('')

function onFileChange(e) {
  file.value = e.target.files[0]
}

async function upload() {
  if (!file.value) {
    error.value = '请先选择文件'
    return
  }

  loading.value = true
  error.value = ''
  summary.value = ''

  try {
    const res = await uploadFile(file.value)
    summary.value = res.summary
  } catch (err) {
    error.value = '上传或总结失败'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 40px auto;
}

button {
  margin-left: 10px;
}

.error {
  color: red;
  margin-top: 10px;
}
</style>
