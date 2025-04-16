import { HttpResponse } from '@/api/interceptor';
import type { AxiosResponse } from 'axios';
import { getCurrentInstance } from 'vue';

export default function onMessage(response: AxiosResponse<HttpResponse>) {
  const c = getCurrentInstance();
  if (c != null) {
    console.log(response);
    c.appContext.config.globalProperties.$message.success({
      content: response.data.message,
      duration: 5 * 1000,
    });
  }
}
