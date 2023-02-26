
export default interface Cart {
  visitor_id: number
  status: number
  subtotal: number
  total: number
  discount_class: number
  discount_rate: number
  discount_price: number
  paid_price: number
  change: number
  list: CartDetail[]
}

type CartDetail = {
  tax: number
  tax_rate: number
  product_id: number
  category_id: number
  data: []
  options: []
  quantity: number
  price: number
}