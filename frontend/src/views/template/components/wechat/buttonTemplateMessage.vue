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
          :field="`config[${props.configKey}].buttonTemplate.mainTitle.title`"
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
  <a-card class="general-card">
    <template #title> 整体卡片点击跳转事件 </template>
    <a-row :gutter="80">
      <a-col :span="8">
        <a-form-item label="跳转事件类型">
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
    <template #title> 按钮设置 </template>
    <template #extra>
      <a-button type="primary" @click="handleButtonAdd">
        <template #icon>
          <icon-plus />
        </template>
        新增按钮设置
      </a-button>
    </template>
    <a-row v-for="(bt, index) of data.buttonList" :key="index" :gutter="80">
      <a-col :span="6">
        <a-form-item label="按钮文案">
          <a-input
            v-model="bt.text"
            placeholder="建议不超过10个字"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-col :span="4">
        <a-form-item label="按钮样式">
          <button-style-select v-model="bt.style"></button-style-select>
        </a-form-item>
      </a-col>
      <a-col :span="6">
        <a-form-item label="回调事件">
          <a-input
            v-model="bt.key"
            placeholder="请输入回调事件"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
      <a-button
        type="primary"
        :style="{ marginLeft: '10px', marginTop: '30px' }"
        @click="handleButtonDelete(index)"
      >
        <template #icon> <icon-delete /> </template>
        <template #default>删除</template>
      </a-button>
    </a-row>
    <a-row :gutter="80">
      <a-col :span="10">
        <a-form-item
          label="回调地址"
          :field="`config[${props.configKey}].buttonTemplate.callback`"
          :rules="[
            {
              required: true,
              message: '请输入回调地址',
            },
          ]"
        >
          <a-input
            v-model="data.callback"
            placeholder="请输入回调地址"
            allow-clear
            show-word-limit
          />
        </a-form-item>
      </a-col>
    </a-row>
  </a-card>
</template>

<script lang="ts" setup>
  import { ButtonTemplateData } from '@/api/template';
  import { onMounted, watch, ref } from 'vue';
  import buttonStyleSelect from './buttonStyleSelect.vue';
  import CardActionSelect from './cardActionSelect.vue';

  export interface Props {
    buttonTemplate?: ButtonTemplateData;
    configKey: any;
  }
  const emit = defineEmits(['messageChanged']);
  const props = withDefaults(defineProps<Props>(), {
    buttonTemplate: () => {
      return {
        source: '',
        mainTitle: {
          title: '',
          desc: '',
        },
        horizontalContentList: [],
        buttonList: [],
        cardAction: {
          type: 0,
        },
        callback: '',
      };
    },
  });

  const data = ref<ButtonTemplateData>({
    source: '',
    mainTitle: {
      title: '',
      desc: '',
    },
    horizontalContentList: [],
    buttonList: [],
    cardAction: {
      type: 0,
    },
    callback: '',
  });

  onMounted(() => {
    data.value = props.buttonTemplate;
  });

  watch(data, () => {
    if (data.value === null) {
      data.value = {
        source: '',
        mainTitle: {
          title: '',
          desc: '',
        },
        horizontalContentList: [],
        buttonList: [],
        cardAction: {
          type: 0,
        },
        callback: '',
      };
    }
    emit('messageChanged', data.value);
  });

  const handleAdd = () => {
    data.value.horizontalContentList.push({
      keyname: '',
      value: '',
    });
  };

  const handleButtonAdd = () => {
    data.value.buttonList.push({
      key: '',
      text: '',
      style: 1,
    });
  };
  const handleDelete = (index: number) => {
    data.value.horizontalContentList.splice(index, 1);
  };

  const handleButtonDelete = (index: number) => {
    data.value.buttonList.splice(index, 1);
  };
</script>

<style scoped lang="less">
  .arco-card-header {
    padding-left: 0px;
  }
</style>
