export interface IPageReq {
  page: number
  size: number
}

export interface IPageResp<T> {
  page: number
  size: number
  total: number
  data: T
}
