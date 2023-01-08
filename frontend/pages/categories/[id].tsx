import React, { useState, useEffect, useRef, useCallback } from "react";
import { useRouter } from 'next/router'
import axios from 'axios';
import Link from 'next/link';
import Image from 'next/image'
// import Category from '../../models/category'
import { getItemsDisplayOn } from "../../apiClient/item"



const Category = () => {
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  // const [categories, setCategories] = useState<Category[]>([])
  const router = useRouter()
  const pathId = router.query.id
  const getItems = useCallback(async (pathId: any) => {
    await getItemsDisplayOn(pathId)
      .then((res) => {
        console.log(res)
      })
  }, [])
  
  useEffect(()  => {
    getItems(pathId);

  }, [pathId])

  return (
    <>
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>
      <ul>
        {/* {categories.map((v, i) =>
          <li key={i}>
            <Link href={{
              pathname: "/categories/[category_id]",
              query: {category_id: v.id}
            }}>
              <a>{v.name}</a>
            </Link>
          </li>
        )} */}
      </ul>
    </>
  )
}

export default Category