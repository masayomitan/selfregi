import React, { useState, useEffect, useRef, useCallback } from "react";
import axios from 'axios';
import Link from 'next/link';
import Image from 'next/image'

const Index = () => {
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    // 'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }

  return (
    <>
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>
    </>
  )
}

export default Index