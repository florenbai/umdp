import axios from 'axios';
import qs from 'query-string';
import { HttpResponse } from './interceptor';

export interface TemplateInterface {
  id: number;
  type: string;
}
export interface TemplateRecord {
  id: number;
  templateName: string;
  professionId: number;
  professionName: string;
  createdAt: string;
  updatedAt: string;
}

export interface TemplateTestParameters {
  receiver: string[];
  cc: string[];
  variable: any;
}
export interface TemplateTestParam {
  templateId: number;
  channel: string;
  parameters: TemplateTestParameters;
}

export interface EmailTemplateConfig {
  id: number;
  type: string;
  title: string;
  content: string;
}

export interface MainTitle {
  title: string;
  desc: string;
}

export interface HorizontalContentList {
  keyname: string;
  value: string;
}

export interface ButtonList {
  text: string;
  style: number;
  key: string;
}

export interface EmphasisContent {
  title: string;
  desc: string;
}

export interface CardAction {
  type?: number;
  url?: string;
  appid?: string;
  pagepath?: string;
}
export interface ButtonTemplateData {
  source: string;
  mainTitle: MainTitle;
  horizontalContentList: HorizontalContentList[];
  buttonList: ButtonList[];
  cardAction: CardAction;
  callback: string;
}

export interface TextNoticeTemplateData {
  source: string;
  mainTitle: MainTitle;
  emphasisContent: EmphasisContent;
  horizontalContentList?: HorizontalContentList[];
  subTitleText?: string;
  cardAction: CardAction;
}

export interface WechatTemplateConfig {
  id: number;
  type: string;
  messageType: number;
  content?: string;
  buttonTemplate?: ButtonTemplateData;
  textNoticeTemplate?: TextNoticeTemplateData;
}

export interface WechatBotTemplateConfig {
  id: number;
  type: string;
  messageType: number;
  content: string;
}

export interface DingTalkTemplateConfig {
  id: number;
  type: string;
  messageType: number;
  title?: string;
  content?: string;
}

export interface PhoneTemplateConfig {
  id: number;
  type: string;
  callCount: number;
  content: string;
}

export interface FeiShuTemplateConfig {
  id: number;
  type: string;
  messageType: number;
  title?: string;
  content?: string;
}

export interface AliyunTemplateParam {
  name: string;
  value: string;
}

export interface AliyunSmsTemplateConfig {
  id: number;
  type: string;
  signName: string;
  templateCode: string;
  templateParam?: AliyunTemplateParam[];
  smsUpExtendCode?: string;
  outId: string;
}

export interface TencentSmsTemplateConfig {
  id: number;
  type: string;
  signName?: string;
  appId: string;
  templateCode: string;
  templateParam?: string[];
  smsUpExtendCode?: string;
}

export interface TemplateForm {
  templateName: string;
  professionId: number | null;
  config: any[];
}

export interface TemplateParams extends Partial<TemplateRecord> {
  current: number;
  pageSize: number;
}

export interface TemplateListRes {
  list: TemplateRecord[];
  total: number;
}

export function queryTemplateList(params: TemplateParams) {
  return axios.get<TemplateListRes>('/api/v1/template/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function submitTemplateForm(data: TemplateForm) {
  return axios.post<HttpResponse>('/api/v1/template', data);
}

export function deleteTemplate(id: number) {
  return axios.delete(`/api/v1/template/${id}`);
}

export function getTemplateById(id: string) {
  return axios.get<TemplateRecord>(`/api/v1/template/${id}`);
}

export function submitEditTemplateForm(id: string, data: TemplateForm) {
  return axios.put<HttpResponse>(`/api/v1/template/${id}`, data);
}

export function submitTestTemplate(data: any) {
  return axios.post<HttpResponse>('/api/v1/template/test', data);
}
