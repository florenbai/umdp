<template>
  <a-select placeholder="请选择渠道" :options="statusOptions" multiple>
  </a-select>
</template>

<script lang="ts" setup>
  import { onMounted, ref, watch } from 'vue';
  import { getAllChannel, ChannelRecord, ChannelListRes } from '@/api/channel';
  import { SelectOptionGroup } from '@arco-design/web-vue';

  const props = defineProps<{
    selected?: number[];
  }>();
  const selectList = ref<ChannelListRes>();
  const labelsSelect: string[] = [];
  const numChannel: Map<number, string> = new Map();
  const statusOptions = ref<SelectOptionGroup[]>();

  const initSelect = () => {
    labelsSelect.length = 0;
    if (props.selected !== undefined) {
      props.selected.forEach((item: number) => {
        const channel = numChannel.get(item);
        if (channel !== undefined) {
          labelsSelect.push(channel);
        }
      });
    }
    const opts: SelectOptionGroup[] = [];
    const labels: string[] = [];
    if (selectList.value === undefined) {
      return;
    }
    selectList.value.list.forEach((item: ChannelRecord) => {
      if (labels.includes(item.channelName)) {
        opts.forEach((optItem, num) => {
          if (optItem.label === item.channelName) {
            if (
              labelsSelect.includes(`${item.channelName}`) &&
              props.selected !== undefined &&
              !props.selected.includes(Number(`${item.id}`))
            ) {
              opts[num].options.push({
                label: `${item.channelName} - ${item.channelTag}`,
                value: item.id,
                disabled: true,
              });
            } else {
              opts[num].options.push({
                label: `${item.channelName} - ${item.channelTag}`,
                value: item.id,
                disabled: false,
              });
            }
          }
        });
      } else {
        if (
          labelsSelect.includes(`${item.channelName}`) &&
          props.selected !== undefined &&
          !props.selected.includes(Number(`${item.id}`))
        ) {
          opts.push({
            isGroup: true,
            label: `${item.channelName}`,
            options: [
              {
                label: `${item.channelName} - ${item.channelTag}`,
                value: item.id,
                disabled: true,
              },
            ],
          });
        } else {
          opts.push({
            isGroup: true,
            label: `${item.channelName}`,
            options: [
              {
                label: `${item.channelName} - ${item.channelTag}`,
                value: item.id,
                disabled: false,
              },
            ],
          });
        }

        labels.push(`${item.channelName}`);
      }
    });
    statusOptions.value = opts;
  };

  onMounted(async () => {
    const { data } = await getAllChannel();
    if (data.list) {
      selectList.value = data;
      data.list.forEach((item: ChannelRecord) => {
        numChannel.set(item.id, `${item.channelName}`);
      });
      initSelect();
    }
  });

  watch(props, () => {
    initSelect();
  });

  const handleChange = (selVal: any) => {
    labelsSelect.length = 0;
    selVal.forEach((item: number) => {
      const channel = numChannel.get(item);
      if (channel !== undefined) {
        labelsSelect.push(channel);
      }
    });
    const opts: SelectOptionGroup[] = [];
    const labels: string[] = [];
    if (selectList.value === undefined) {
      return;
    }
    selectList.value.list.forEach((item: ChannelRecord) => {
      if (labels.includes(item.channelName)) {
        opts.forEach((optItem, num) => {
          if (optItem.label === item.channelName) {
            if (
              labelsSelect.includes(`${item.channelName}`) &&
              !selVal.includes(Number(`${item.id}`))
            ) {
              opts[num].options.push({
                label: `${item.channelName} - ${item.channelTag}`,
                value: item.id,
                disabled: true,
              });
            } else {
              opts[num].options.push({
                label: `${item.channelName} - ${item.channelTag}`,
                value: item.id,
                disabled: false,
              });
            }
          }
        });
      } else {
        if (
          labelsSelect.includes(`${item.channelName}`) &&
          !selVal.includes(Number(`${item.id}`))
        ) {
          opts.push({
            isGroup: true,
            label: `${item.channelName}`,
            options: [
              {
                label: `${item.channelName} - ${item.channelTag}`,
                value: item.id,
                disabled: true,
              },
            ],
          });
        } else {
          opts.push({
            isGroup: true,
            label: `${item.channelName}`,
            options: [
              {
                label: `${item.channelName} - ${item.channelTag}`,
                value: item.id,
                disabled: false,
              },
            ],
          });
        }

        labels.push(`${item.channelName}`);
      }
    });
    statusOptions.value = opts;
  };
</script>
