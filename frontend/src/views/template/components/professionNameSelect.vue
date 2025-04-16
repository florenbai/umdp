<template>
  <a-select placeholder="请选择所属业务" :options="statusOptions"> </a-select>
</template>

<script lang="ts" setup>
  import { SelectOptionData } from '@arco-design/web-vue/es/select/interface';
  import useRequest from '@/hooks/request';
  import { computed } from 'vue';
  import { getAllProfession, ProfessionRecord } from '@/api/profession';

  const { response: data } = useRequest<ProfessionRecord[]>(getAllProfession);

  const statusOptions = computed<SelectOptionData[]>(() => {
    if (!data.value) {
      return [];
    }
    const opts = data.value.map((item: ProfessionRecord) => ({
      label: `${item.professionName}`,
      value: item.id,
    }));
    return opts;
  });
</script>
