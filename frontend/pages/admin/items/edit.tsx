import React, { useState, useEffect, useRef, useCallback } from "react"
import { SubmitHandler, useForm } from 'react-hook-form'
import axios from 'axios'
import Link from 'next/link'
import Image from 'next/image'
import Item from '../../../models/item'
import Category from '../../../models/category'
import { getCategoriesForAdmin } from "../../../apiClient/category"
import { postItem } from "../../../apiClient/item"

export const ItemEdit = () => {
  const apiBaseUrl = process.env.NEXT_PUBLIC_BASE_URL
  axios.defaults.baseURL = apiBaseUrl
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  const [item, setItem] = useState<Item[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const { register, handleSubmit } = useForm<Item>({
    mode: "onSubmit",
    reValidateMode: "onChange",
  })

  const handleOnSubmit: SubmitHandler<Item> = async (item) => {
    console.log(item)

    await postItem(item)
      .then((res: any) => {
        console.log(res)
        if (res.data) {

        }
      })
  }

  const getAllCategories = useCallback(async () => {
    await getCategoriesForAdmin()
      .then((res: any) => {
        if (res.data) {
          setCategories(res.data)
        }
      })
  }, [])
  
  useEffect(()  => {
    getAllCategories();
  }, [])

  return (
    <>
      <form onSubmit={handleSubmit(handleOnSubmit)}>
        <div className="gap-6 mb-6 md:grid-cols-2">
          <div>
            <label htmlFor="name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">名称</label>
            <input 
              type="text"
              id="name" 
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700" 
              {...register('name', {
                required: true,
              })}
            />
          </div>  
          <div>
            <label htmlFor="price" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">値段</label>
            <input 
              type="prnumberice" 
              id="price"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 "
              placeholder="0000" 
              {...register("price", {
                valueAsNumber: true,
                required: true,
              })}
            />
          </div>
        </div>
        <label>カテゴリー</label>
        <select 
          id="category_id" 
          className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700"
          {...register("category_id", {
            required: true,
          })}
        >
          <option value="" hidden>選択してください</option>
            {categories.map((value, index) =>
                <option
                    value={value.id} 
                    key={index}>{value.name}
                </option>
              )
            }
        </select>
        <div className="flex items-start mb-6">
          <div className="flex items-center h-5">
            <input 
              id="is_display" 
              type="checkbox" 
              value="" className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300"
              {...register("is_display", {
                required: false,
              })}
            />
          </div>
          <label htmlFor="is_display" className="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">表示</label>
        </div>
        <button 
          type="submit" 
          className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center"
        >
          登録
        </button>
      </form>
    </>
  )
}

export default ItemEdit