import {ApiClient} from '.'

const apiBaseUrl = process.env.NEXT_PUBLIC_BASE_URL
const apiClient = new ApiClient(apiBaseUrl)

export const getItem = (itemId: string) => {
  return apiClient.get(
    '/api/items/get/' + itemId,
  )
}

export const getItems = (categoryId: string) => {
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