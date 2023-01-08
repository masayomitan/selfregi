import {ApiClient} from '.'

const apiBaseUrl = process.env.NEXT_PUBLIC_BASE_URL
const apiClient = new ApiClient(apiBaseUrl)


export const getCategories = () => {
  return apiClient.get(
    "/api/categories/get"
  )
}

export const getCategoriesForAdmin = () => {
    return apiClient.get(
      "/api/admin/categories/get"
    )
}