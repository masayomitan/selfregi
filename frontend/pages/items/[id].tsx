import React, { useState, useEffect, useRef, useCallback } from "react";
import { useRouter } from 'next/router'
import axios from 'axios';
import Link from 'next/link';
import Image from 'next/image'
import Item from '../../models/item'
import { getItemDisplayOn } from "../../apiClient/item"



const Category = () => {
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }

  const [item, setItem] = useState<Item>()
  const router = useRouter()
  const pathId = router.query.id
  console.log(pathId)
  const getItem = useCallback(async (pathId: any) => {
    if (pathId === undefined) return 
    await getItemDisplayOn(pathId)
      .then((res: any) => {
        console.log(res.data)
        setItem(res.data)
      })
  }, [])
  
  useEffect(()  => {
    getItem(pathId);
  }, [pathId])

  return (
    <>
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>
      <ul>
        {/* {item.map((v, i) =>
          <li key={i}>
            <Link href={{
              pathname: "/items/[item_id]",
              query: {item_id: v.id}
            }}>
              <p>{v.name}</p>
            </Link>
          </li>
        )} */}
      </ul>
    </>
  )
}

export default Category