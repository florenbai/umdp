import router from '@/router';

export function gotoAddChannel() {
  router.push(`/channel/add`);
}

export function gotoChannelList() {
  router.push(`/channel/list`);
}

export function gotoEditChannel(id: number) {
  router.push(`/channel/edit/${id}`);
}
