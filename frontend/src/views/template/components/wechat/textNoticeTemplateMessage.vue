<template>
  <a-card class="general-card">
    <a-row :gutter="80">
      <a-col :span="8">
        <a-form-item label="消息来源">
          <a-input
            v-model="data.source"
            placeholder="请输入消息来源,可为空"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item
          label="主标题"
          :field="`config[${props.configKey}].textNoticeTemplate.mainTitle.title`"
          :rules="[
            {
              required: true,
              message: '请输入主标题',
            },
          ]"
        >
          <a-input
            v-model="data.mainTitle.title"
            placeholder="请输入主标题"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item label="标题辅助信息">
          <a-input
            v-model="data.mainTitle.desc"
            placeholder="请输入标题辅助信息,可为空"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item label="二级普通文本">
          <a-input
            v-model="data.subTitleText"
            placeholder="请输入二级普通文本,可为空"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item label="关键数据样式内容">
          <a-input
            v-model="data.emphasisContent.title"
            placeholder="请输入关键数据样式内容,可为空"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item label="关键数据样式描述">
          <a-input
            v-model="data.emphasisContent.desc"
            placeholder="请输入关键数据样式描述,可为空"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
  <a-card class="general-card">
    <template #title> 整体卡片点击跳转事件 </template>
    <a-row :gutter="80">
      <a-col :span="8">
        <a-form-item
          label="跳转事件类型"
          :field="`config[${props.configKey}].textNoticeTemplate.cardAction.type`"
          :rules="[
            {
              required: true,
              message: '请选择跳转事件类型',
            },
          ]"
        >
          <CardActionSelect v-model="data.cardAction.type"></CardActionSelect>
        </a-form-item>
      </a-col>
      <a-col v-if="data.cardAction.type === 1" :span="8">
        <a-form-item label="跳转事件的url">
          <a-input
            v-model="data.cardAction.url"
            placeholder="请输入跳转事件的url"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col v-if="data.cardAction.type === 2" :span="8">
        <a-form-item label="小程序appid">
          <a-input
            v-model="data.cardAction.appid"
            placeholder="请输入小程序appid"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col v-if="data.cardAction.type === 2" :span="8">
        <a-form-item label="小程序pagepath">
          <a-input
            v-model="data.cardAction.pagepath"
            placeholder="请输入小程序pagepath"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
  <a-card class="general-card">
    <template #title> 二级标题和文本列表设置 </template>
    <template #extra>
      <a-button type="primary" @click="handleAdd">
        <template #icon>
          <icon-plus />
        </template>
        新增二级标题设置
      </a-button>
    </template>
    <a-row
      v-for="(horizontalContent, index) of data.horizontalContentList"
      :key="index"
      :gutter="80"
    >
      <a-col :span="8">
        <a-form-item label="二级标题">
          <a-input
            v-model="horizontalContent.keyname"
            placeholder="请输入二级标题，不超过5个字"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item label="二级文本">
          <a-input
            v-model="horizontalContent.value"
            placeholder="请输入二级文本，不超过30个字"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-button
        type="primary"
        :style="{ marginLeft: '10px', marginTop: '30px' }"
        @click="handleDelete(index)"
      >
        <template #icon> <icon-delete /> </template>
        <template #default>删除</template>
      </a-button>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { TextNoticeTemplateData } from '@/api/template';
  import { onMounted, watch, ref } from 'vue';
  import CardActionSelect from './cardActionSelect.vue';

  export interface Props {
    textNoticeTemplate?: TextNoticeTemplateData;
    configKey: any;
  }
  const emit = defineEmits(['messageChanged']);
  const props = withDefaults(defineProps<Props>(), {
    textNoticeTemplate: () => {
      return {
        source: '',
        mainTitle: {
          title: '',
          desc: '',
        },
        subTitleText: '',
        emphasisContent: {
          title: '',
          desc: '',
        },
        horizontalContentList: [],
        cardAction: {
          type: 1,
        },
      };
    },
  });

  const data = ref<TextNoticeTemplateData>({
    source: '',
    mainTitle: {
      title: '',
      desc: '',
    },
    subTitleText: '',
    emphasisContent: {
      title: '',
      desc: '',
    },
    horizontalContentList: [],
    cardAction: {
      type: 1,
    },
  });

  onMounted(() => {
    data.value = props.textNoticeTemplate;
  });

  watch(data, () => {
    if (data.value === null) {
      data.value = {
        source: '',
        mainTitle: {
          title: '',
          desc: '',
        },
        subTitleText: '',
        emphasisContent: {
          title: '',
          desc: '',
        },
        horizontalContentList: [],
        cardAction: {
          type: 1,
        },
      };
    }
    emit('messageChanged', data.value);
  });

  const handleAdd = () => {
    if (data.value.horizontalContentList !== undefined) {
      data.value.horizontalContentList.push({
        keyname: '',
        value: '',
      });
    }
  };

  const handleDelete = (index: number) => {
    if (data.value.horizontalContentList !== undefined) {
      data.value.horizontalContentList.splice(index, 1);
    }
  };
</script>

<style scoped lang="less">
  .arco-card-header {
    padding-left: 0px;
  }

  .general-card {
    border: 0;
  }
</style>
