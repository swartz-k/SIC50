export interface ILocaleItem {
  cn: string
  en: string
}

const locales0 = {
  add_step: {
    cn: '增加步骤',
    en: 'Add Step',
  },
  click_drag_upload_text: {
    cn: '点击上传文件或拖拽文件到这里',
    en: 'Click or Drag to upload',
  },
  collapse: {
    cn: '折叠',
    en: 'Collapse',
  },
  docs: {
    cn: '文档',
    en: 'Docs',
  },
  overview: {
    cn: '概览',
    en: 'Overview',
  },
  quer_result: {
    cn: '查询结果',
    en: 'Query Result',
  },
  search: {
    cn: '搜索',
    en: 'Search',
  },
  step_images: {
    cn: '实验步骤结果',
    en: 'Step Images',
  },
  step_num: {
    cn: '步骤数',
    en: 'Step Num',
  },
  step_concentration: {
    cn: '浓度',
    en: 'concentration',
  },
  submit: {
    cn: '提交',
    en: 'Submit',
  },
  task: {
    cn: '任务',
    en: 'Task',
  },
  upload_file: {
    cn: '实验文件',
    en: 'Lib Image',
  },
  welcome: {
    cn: '欢迎',
    en: 'welcome',
  },
}

export const locales: { [key in keyof typeof locales0]: ILocaleItem } = locales0
