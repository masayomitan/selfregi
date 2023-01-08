import { immerable, produce } from 'immer'

export interface Item {
  id?: number
  name: string
  category_id: number
  is_display: number
}

export default Item