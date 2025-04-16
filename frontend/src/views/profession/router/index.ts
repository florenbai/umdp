import router from '@/router';

export function gotoAddProfession() {
  router.push(`/profession/add`);
}

export function gotoProfessionList() {
  router.push(`/profession/list`);
}

export function gotoEditProfession(id: number) {
  router.push(`/profession/edit/${id}`);
}
