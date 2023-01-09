import React, { useState, useEffect, useRef, useCallback } from "react";
import { useRouter } from 'next/router'
import axios from 'axios';
import Link from 'next/link';
import Image from 'next/image'
import Item from '../../models/item'
import { getItemsDisplayOn } from "../../apiClient/item"



const Category = () => {
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  const [items, setItems] = useState<Item[]>([])
  const router = useRouter()
  const pathId = router.query.id
  const getItems = useCallback(async (pathId: any) => {
    if (pathId === undefined) return 
    await getItemsDisplayOn(pathId)
      .then((res: any) => {
          setItems(res.data)
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
        {items.map((v, i) =>
          <li key={i}>
            <Link href={{
              pathname: "/items/[item_id]",
              query: {item_id: v.id}
            }}>
              <p>{v.name}</p>
            </Link>
          </li>
        )}
      </ul>
    </>
  )
}

export default Category