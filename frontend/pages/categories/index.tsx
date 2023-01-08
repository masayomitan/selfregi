import React, { useState, useEffect, useRef, useCallback } from "react";
import axios from 'axios';
import Link from 'next/link';
import Image from 'next/image'
import Category from '../../models/category'
import { getCategories } from "../../apiClient/category"


const CategoryIndex = () => {
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }

  const [categories, setCategories] = useState<Category[]>([])

  const getAllCategories = useCallback(async () => {
    await getCategories()
      .then((res: any) => {
        console.log(res.data)
        if (res.data !== undefined) {
          setCategories(res.data)
        }
      })
  }, [])
  
  useEffect(()  => {
    getAllCategories();

  }, [])

  return (
    <>
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>
      <ul>
        {categories.map((v: any, i: any) =>
          <li key={i}>
            <Link href={{
              pathname: "/categories/[category_id]",
              query: {category_id: v.id}
            }}>
              <p>{v.name}</p>
            </Link>
          </li>
        )}
      </ul>
    </>
  )
}

export default CategoryIndex