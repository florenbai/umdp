<template>
  <a-button type="text" size="small" @click="handleClick"> 查看 </a-button>
  <a-modal
    v-model:visible="visible"
    :hide-cancel="true"
    :closable="false"
    draggable
    :hide-title="true"
    @ok="handleOk"
  >
    <a-space direction="vertical" size="large" fill>
      <a-descriptions title="请求参数">
        <a-descriptions-item>
          <code>{{ props.parameters }}</code>
        </a-descriptions-item>
      </a-descriptions>
      <a-descriptions v-if="props.message !== ''" title="异常信息">
        <a-descriptions-item>
          {{ props.message }}
        </a-descriptions-item>
      </a-descriptions>
    </a-space>
  </a-modal>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';

  const props = defineProps({
    message: {
      type: String,
    },
    parameters: {
      type: String,
    },
  });

  const visible = ref(false);
  const handleOk = () => {
    visible.value = false;
  };
  const handleClick = () => {
    visible.value = true;
  };
</script>
