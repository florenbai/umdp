import router from '@/router';

export function gotoAddTemplate() {
  router.push(`/template/add`);
}

export function gotoTemplateList() {
  router.push(`/template/list`);
}

export function gotoEditTemplate(id: number) {
  router.push(`/template/edit/${id}`);
}
