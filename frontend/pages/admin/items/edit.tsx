import React, { useState, useEffect, useRef, useCallback } from "react"
import axios from 'axios'
import Link from 'next/link'
import Image from 'next/image'
import Item from '../../../models/item'
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

  const getAllCategories = useCallback(async () => {
    await getCategoriesForAdmin()
      .then((res) => {
        console.log(res)
      })
  }, [])
  
  useEffect(()  => {
    getAllCategories();
    setItem([
      {
        name: 'test',
        category_id: 1,
        is_display: 1,
      }
    ])
  }, [])
  
  const save = async (item: Item[]) => {
    await postItem(item)
      .then((res) => {
        console.log(res)
      })
  }

  return (
    <>
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>
      <div className="grid gap-6 mb-6 md:grid-cols-2">
        <div>
          <label htmlFor="name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">名称</label>
          <input type="text" id="name" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Flowbite" required />
        </div>  
        <div>
          <label htmlFor="price" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">値段</label>
          <input type="price" id="price" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="0000" pattern="[0-9]{3}-[0-9]{2}-[0-9]{3}" required />
        </div>
      </div>
      <select id="countries" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
        <option defaultValue="JP">Choose a country</option>
        <option value="US">United States</option>
        <option value="CA">Canada</option>
        <option value="FR">France</option>
        <option value="DE">Germany</option>
      </select>
      <div className="flex items-start mb-6">
        <div className="flex items-center h-5">
          <input id="category_id" type="checkbox" value="" className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-blue-600 dark:ring-offset-gray-800" required />
        </div>
        <label htmlFor="category_id" className="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">表示</label>
      </div>
      <button 
        onClick={() => save(item)}
        type="submit" className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Submit</button>
    </>
  )
}

export default ItemEdit