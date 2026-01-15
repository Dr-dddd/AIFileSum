<template>
 <div class="chat">
  <div
    v-for="(msg, index) in messages"
    :key="index"
    :class="['bubble', msg.role]"
  >
    <span v-if="msg.type === 'file'">📎 {{ msg.content }}</span>
    <span v-else>{{ msg.content }}</span>
  </div>
</div>
<div class="input-area">
  <input type="file" @change="onFileChange" />
  <button @click="upload">发送</button>
</div>

</template>

<script setup>
import { ref } from 'vue'
import { uploadFile } from '../api/summary'

const file = ref(null)
const messages = ref([])
const loading = ref(false)

function onFileChange(e) {
  file.value = e.target.files[0]
}

async function upload() {
  if (!file.value) {
    return
  }

  loading.value = true

  messages.value.push({
    role: 'user',
    type: 'file',
    content: file.value.name
  })

   // AI 占位消息
  messages.value.push({
    role: 'assistant',
    type: 'text',
    content: '正在总结中...'
  })

  const aiIndex = messages.value.length - 1

  try {
    const res = await uploadFile(file.value)

    messages.value[aiIndex].content = res.content
  } catch (e) {
    messages.value[aiIndex].content = '总结失败，请重试'
  } finally {
    loading.value = false
    file.value = null
  }

}
</script>

<style scoped>
.chat {
  display: flex;
  flex-direction: column;
  margin-bottom: 20px;
}

.bubble {
  padding: 10px 14px;
  margin: 6px;
  border-radius: 8px;
  max-width: 70%;
  word-break: break-word;
}

.user {
  background: #d1e7ff;
  align-self: flex-end;
}

.assistant {
  background: #f1f1f1;
  align-self: flex-start;
}

.input-area {
  display: flex;
  gap: 10px;
}
</style>

