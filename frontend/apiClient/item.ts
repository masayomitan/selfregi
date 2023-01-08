import {ApiClient} from '.'

const apiBaseUrl = process.env.NEXT_PUBLIC_BASE_URL
const apiClient = new ApiClient(apiBaseUrl)

export const getItemsDisplayOn = (categoryId: string) => {
  return apiClient.get(
    '/api/categories/' + categoryId + '/items/get',
  )
}

export const postItem = (payload: any) => {
    return apiClient.post(
      '/api/admin/items/post',
      payload
    )
}