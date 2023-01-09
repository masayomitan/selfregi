import { immerable, produce } from 'immer'

export interface Item {
  id?: number
  name: string
  category_id: number
  is_display: number
  price: number
  tax: number
  tax_rate: number
  temporary_stock: number
}

export default Item