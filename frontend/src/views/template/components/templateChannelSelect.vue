<template>
  <a-select placeholder="请选择渠道" :options="options"> </a-select>
</template>

<script lang="ts" setup>
  import {
    ProfessionChannelRecord,
    getProfessionChannels,
  } from '@/api/profession';
  import { SelectOptionData } from '@arco-design/web-vue/es/select';
  import { onMounted, ref } from 'vue';

  const props = defineProps<{
    profession: number;
  }>();

  const options = ref<SelectOptionData[]>([]);
  onMounted(async () => {
    const { data } = await getProfessionChannels(props.profession);
    data.forEach((item: ProfessionChannelRecord) => {
      options.value.push({
        label: item.channelName,
        value: item.channelTag,
      });
    });
  });
</script>
