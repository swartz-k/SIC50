export interface ILocaleItem {
  cn: string
  en: string
}

const locales0 = {
  add_step: {
    cn: '增加步骤',
    en: 'Add Step',
  },
  clear: {
    cn: '清楚',
    en: 'Clear',
  },
  close: {
    cn: '关闭',
    en: 'Close',
  },
  click_drag_upload_text: {
    cn: '点击上传文件或拖拽文件到这里',
    en: 'Click or Drag to upload',
  },
  collapse: {
    cn: '折叠',
    en: 'Collapse',
  },
  create_at: {
    cn: '创建时间',
    en: 'Create Time',
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
  result: {
    cn: '任务结果',
    en: 'Task Result',
  },
  refresh: {
    cn: '刷新',
    en: 'Refresh',
  },
  search: {
    cn: '搜索',
    en: 'Search',
  },
  status: {
    cn: '状态',
    en: 'Status',
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
  task_submit_success_tooltip: {
    cn: '任务提交成功，请稍后用以下任务 ID 查询结果',
    en: 'Task Submit Success. Later you can use task id query result',
  },
  total_images: {
    cn: '总共处理图片数',
    en: 'Total Images',
  },
  total_tasks: {
    cn: '总共处理任务数',
    en: 'Total Tasks',
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
