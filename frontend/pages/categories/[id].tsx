import React, { useState, useEffect, useRef, useCallback } from "react";
import { useRouter } from 'next/router'
import axios from 'axios';
import Link from 'next/link';
import Image from 'next/image'
import Category from '../../models/category'
import Item from '../../models/item'
import { getCategories } from "../../apiClient/category"
import { getItems } from "../../apiClient/item"



const Category = () => {
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  const [items, setItems] = useState<Item[]>([])
  const [categories, setCategories] = useState<Category[]>([])

  const router = useRouter()
  const pathId = router.query.id
  const getMenu = useCallback(async (pathId: any) => {
    if (pathId === undefined) return 
    await getItems(pathId)
      .then((res: any) => {
          setItems(res.data)
      })
  }, [])

  const getAllCategories = useCallback(async () => {
    await getCategories()
      .then((res: any) => {
          setCategories(res.data)
      })
  }, [])
  
  useEffect(()  => {
    getMenu(pathId)
    getAllCategories()

  }, [pathId])

  return (
    <>
      <div className="fixed top-0 left-0 z-20 w-full border-b border-gray-200 bg-white py-2.5 px-6 sm:px-4">
        <div className="container mx-auto flex max-w-6xl flex-wrap items-center justify-between">
          <a href="#" className="flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" className="mr-3 h-6 text-blue-500 sm:h-9">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21 7.5l-9-5.25L3 7.5m18 0l-9 5.25m9-5.25v9l-9 5.25M3 7.5l9 5.25M3 7.5v9l9 5.25m0-9v9" />
            </svg>
            <span className="self-center whitespace-nowrap text-xl font-semibold">selfregi</span>
          </a>
          <div className="mt-2 sm:mt-0 sm:flex md:order-2">
          </div>
          <div className="hidden w-full items-center justify-between md:order-1 md:flex md:w-auto" id="navbar-sticky">
            <ul className="mt-4 flex flex-col rounded-lg border border-gray-100 bg-gray-50 p-4 md:mt-0 md:flex-row md:space-x-8 md:border-0 md:bg-white md:text-sm md:font-medium">
              {categories.map((c, i) =>
              <Link 
              href={{
                pathname: '/categories/' + c.id,
              }}
              >
                <p className="block rounded bg-blue-700 py-2 pl-3 pr-4 text-white md:bg-transparent md:p-0 md:text-blue-700" aria-current="page">{c.name}</p>
              </Link>
              )}

            </ul>
          </div>
        </div>
      </div>
      <section className="py-10 bg-gray-100">
        <div className="mx-auto grid max-w-6xl  grid-cols-1 gap-6 p-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {items.map((item, i) =>
            <article className="rounded-xl bg-white p-3 shadow-lg hover:shadow-xl">
              <Link 
                href={{
                  pathname: '/items/' + item.id,
                }}
                >
                  <div className="relative flex items-end overflow-hidden rounded-xl">
                    <img src="https://i.imgur.com/GIeyjWd.jpg" alt="Hotel Photo" />
                    <div className="flex items-center space-x-1.5 rounded-lg bg-blue-500 px-4 py-1.5 text-white duration-100 hover:bg-blue-600">
                      <button className="text-sm">Add to cart</button>
                    </div>
                  </div>

                  <div className="mt-1 p-2">
                    <h2 className="text-slate-700">{item.name}</h2>
                    <p className="mt-1 text-sm text-slate-400">{item.temporary_stock}</p>
                    <div className="mt-3 flex items-end justify-between">
                      <p>
                        <span className="text-lg font-bold text-blue-500">$850</span>
                      </p>
                    </div>
                  </div>
              </Link>
            </article>
          )}
        </div>
      </section>
    </>
  )
}

export default Category