import React, { useState, useEffect, useRef, useCallback } from "react";
import { useRouter } from 'next/router'
import axios from 'axios'
import { useRecoilState, RecoilEnv } from "recoil"
import Link from 'next/link';
import Image from 'next/image'
import Item from '../../models/item'
import Cart from '../../models/cart'
import { getItem } from "../../apiClient/item"
import { atom } from 'recoil'

const Category = () => {
  RecoilEnv.RECOIL_DUPLICATE_ATOM_KEY_CHECKING_ENABLED = false
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  const cartState = atom<Cart[]>({
    key: "carts",
    default: [],
  })

  const [item, setItem] = useState<Item>()
  const [cart, setCart] = useRecoilState<Cart[]>(cartState)

  const router = useRouter()
  const pathId = router.query.id
  const getMenuDetail = useCallback(async (pathId: any) => {
    if (pathId === undefined) return 
    await getItem(pathId)
      .then((res: any) => {
        setItem(res.data)
      })
      .catch((err) => {
        console.log(err)
        router.push('/')
      })
  }, [])
  console.log(cart)
  
  useEffect(()  => {
    getMenuDetail(pathId);
  }, [pathId])

  const itemInCart = (cart: Cart[], item: Item) => {
    cart.find((detail) => {
      detail.list.find((product) => {
        if (product.product_id === item.id) {
          return true
        } else {
          return false
        }
      })
    })
  }

  const addCart = (item: Item) => {
    let test: any = {}
    if (Object.keys(cart).length === 0) {
      console.log(item)
      test = {
        visitor_id: 1,
        status: item.id,
        subtotal: 1,
        total: 1,
        discount_class: 1,
        discount_rate: 1,
        discount_price: 1,
        paid_price: 1,
        change: 1,
        list: [
  
        ]
      }
    }
    setCart(test)
    console.log(cart)
    router.back()
  }

  const changeCart = (item: Item) => {

  }

  return (
    <>
      <section className="py-10 font-poppins dark:bg-gray-800">
        {item ?
        <div className="max-w-6xl px-4 mx-auto">
          <div className="flex flex-wrap mb-24 -mx-4">
            <div className="w-full px-4 mb-8 md:w-1/2 md:mb-0">
              <div className="sticky top-0 overflow-hiden ">
                <div className="relative mb-6 lg:mb-10 lg:h-96">
                  <a className="absolute left-0 transform lg:ml-2 top-1/2 translate-1/2" href="#">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="w-5 h-5 text-blue-500 bi bi-chevron-left dark:text-blue-200" viewBox="0 0 16 16">
                      <path fill-rule="evenodd" d="M11.354 1.646a.5.5 0 0 1 0 .708L5.707 8l5.647 5.646a.5.5 0 0 1-.708.708l-6-6a.5.5 0 0 1 0-.708l6-6a.5.5 0 0 1 .708 0z"></path>
                    </svg>
                  </a>
                  <img className="object-contain w-full lg:h-full" src="https://i.postimg.cc/0jwyVgqz/Microprocessor1-removebg-preview.png" alt=""/>
                  <a className="absolute right-0 transform lg:mr-2 top-1/2 translate-1/2" href="#">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="w-5 h-5 text-blue-500 bi bi-chevron-right dark:text-blue-200" viewBox="0 0 16 16">
                      <path fill-rule="evenodd" d="M4.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L10.293 8 4.646 2.354a.5.5 0 0 1 0-.708z"></path>
                    </svg>
                  </a>
                </div>
                <div className="flex-wrap hidden -mx-2 md:flex">
                  <div className="w-1/2 p-2 sm:w-1/4">
                    <a className="block border border-gray-200 hover:border-blue-400 dark:border-gray-700 dark:hover:border-blue-300" href="#">
                      <img className="object-contain w-full lg:h-28" src="https://i.postimg.cc/Z5KhRkD6/download-1-removebg-preview.png" alt="" />
                    </a>
                  </div>
                </div>
              </div>
            </div>
            <div className="w-full px-4 md:w-1/2">
              <div className="lg:pl-20">
                <div className="mb-6 ">
                  <h2 className="max-w-xl mt-6 mb-6 text-xl font-semibold leading-loose tracking-wide text-gray-700 md:text-2xl dark:text-gray-300">
                    {item.name}
                  </h2>
                  <div className="flex flex-wrap items-center mb-6">
                    <ul className="flex mb-4 mr-2 lg:mb-0">
                      <li>
                        <a href="#">
                          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className="w-4 mr-1 text-red-500 dark:text-gray-400 bi bi-star " viewBox="0 0 16 16">
                            <path d="M2.866 14.85c-.078.444.36.791.746.593l4.39-2.256 4.389 2.256c.386.198.824-.149.746-.592l-.83-4.73 3.522-3.356c.33-.314.16-.888-.282-.95l-4.898-.696L8.465.792a.513.513 0 0 0-.927 0L5.354 5.12l-4.898.696c-.441.062-.612.636-.283.95l3.523 3.356-.83 4.73zm4.905-2.767-3.686 1.894.694-3.957a.565.565 0 0 0-.163-.505L1.71 6.745l4.052-.576a.525.525 0 0 0 .393-.288L8 2.223l1.847 3.658a.525.525 0 0 0 .393.288l4.052.575-2.906 2.77a.565.565 0 0 0-.163.506l.694 3.957-3.686-1.894a.503.503 0 0 0-.461 0z"></path>
                          </svg>
                        </a>
                      </li>
                    </ul>
                  </div>
                  <p className="inline-block text-2xl font-semibold text-gray-700 dark:text-gray-400 ">
                    <span>Rs.7,000.00</span>
                  </p>
                </div>
                <div className="mb-6">
                  <h2 className="mb-2 text-lg font-bold text-gray-700 dark:text-gray-400">System Specs :</h2>
                  <div className="bg-gray-100 dark:bg-gray-700 rounded-xl">
                    <div className="p-3 lg:p-5 ">
                      <div className="p-2 rounded-xl lg:p-6 dark:bg-gray-800 bg-gray-50">
                        <div className="flex flex-wrap justify-center gap-x-10 gap-y-4">
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="flex flex-wrap items-center mb-6">
                  <div className="mb-4 mr-4 lg:mb-0">
                    <div className="w-28">
                      <div className="relative flex flex-row w-full h-10 bg-transparent rounded-lg">
                        <button className="w-20 h-full text-gray-600 bg-gray-100 border-r rounded-l outline-none cursor-pointer dark:border-gray-700 dark:hover:bg-gray-700 dark:text-gray-400 hover:text-gray-700 dark:bg-gray-900 hover:bg-gray-300">
                          <span className="m-auto text-2xl font-thin">-</span>
                        </button>
                        <button className="w-20 h-full text-gray-600 bg-gray-100 border-l rounded-r outline-none cursor-pointer dark:border-gray-700 dark:hover:bg-gray-700 dark:text-gray-400 dark:bg-gray-900 hover:text-gray-700 hover:bg-gray-300">
                          <span className="m-auto text-2xl font-thin">+</span>
                        </button>
                      </div>
                    </div>
                  </div>
                  <div className="mb-4 lg:mb-0">
                    <button className="flex items-center justify-center w-full h-10 p-2 mr-4 text-gray-700 border border-gray-300 lg:w-11 hover:text-gray-50 dark:text-gray-200 dark:border-blue-600 hover:bg-blue-600 hover:border-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500 dark:hover:border-blue-500 dark:hover:text-gray-100">
                      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" className=" bi bi-heart" viewBox="0 0 16 16">
                        <path d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01L8 2.748zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15z"></path>
                      </svg>
                    </button>
                  </div>
                  {(itemInCart) ?
                    <button onClick={() => addCart(item)} className="w-full px-4 py-3 text-center text-blue-600 bg-blue-100 border border-blue-600 dark:hover:bg-gray-900 dark:text-gray-400 dark:border-gray-700 dark:bg-gray-700 hover:bg-blue-600 hover:text-gray-100 lg:w-1/2 rounded-xl">
                      Add to cart
                    </button>
                    :
                    <button onClick={() => addCart(item)} className="w-full px-4 py-3 text-center text-blue-600 bg-blue-100 border border-blue-600 dark:hover:bg-gray-900 dark:text-gray-400 dark:border-gray-700 dark:bg-gray-700 hover:bg-blue-600 hover:text-gray-100 lg:w-1/2 rounded-xl">
                      edit to cart
                    </button>
                  }
                </div>
              </div>
            </div>
          </div>
        </div>
        : 'no item' 
       }
      </section>
    </>
  )
}

export default Category