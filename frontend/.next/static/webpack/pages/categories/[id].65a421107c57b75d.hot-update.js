"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
self["webpackHotUpdate_N_E"]("pages/categories/[id]",{

/***/ "./pages/categories/[id].tsx":
/*!***********************************!*\
  !*** ./pages/categories/[id].tsx ***!
  \***********************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"./node_modules/react/jsx-dev-runtime.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"./node_modules/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! next/router */ \"./node_modules/next/router.js\");\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(next_router__WEBPACK_IMPORTED_MODULE_2__);\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! axios */ \"./node_modules/axios/index.js\");\n/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! next/link */ \"./node_modules/next/link.js\");\n/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(next_link__WEBPACK_IMPORTED_MODULE_3__);\n/* harmony import */ var _apiClient_item__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../../apiClient/item */ \"./apiClient/item.ts\");\n\nvar _s = $RefreshSig$();\n\n\n\n\n\nconst Category = ()=>{\n    _s();\n    axios__WEBPACK_IMPORTED_MODULE_5__[\"default\"].defaults.headers.common = {\n        \"Content-Type\": \"application/json\"\n    };\n    const [items, setItems] = (0,react__WEBPACK_IMPORTED_MODULE_1__.useState)([]);\n    const router = (0,next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter)();\n    const pathId = router.query.id;\n    const getItems = (0,react__WEBPACK_IMPORTED_MODULE_1__.useCallback)(async (pathId)=>{\n        await (0,_apiClient_item__WEBPACK_IMPORTED_MODULE_4__.getItemsDisplayOn)(pathId).then((res)=>{\n            setItems(res.data);\n        });\n    }, []);\n    (0,react__WEBPACK_IMPORTED_MODULE_1__.useEffect)(()=>{\n        getItems(pathId);\n    }, [\n        pathId\n    ]);\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.Fragment, {\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h1\", {\n                className: \"text-3xl font-bold underline\",\n                children: \"Hello world!\"\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/categories/[id].tsx\",\n                lineNumber: 34,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"ul\", {\n                children: items.map((v, i)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"li\", {\n                        children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)((next_link__WEBPACK_IMPORTED_MODULE_3___default()), {\n                            href: {\n                                pathname: \"/items/[item_id]\",\n                                query: {\n                                    item_id: v.id\n                                }\n                            },\n                            children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"a\", {\n                                children: v.name\n                            }, void 0, false, {\n                                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/categories/[id].tsx\",\n                                lineNumber: 44,\n                                columnNumber: 15\n                            }, undefined)\n                        }, void 0, false, {\n                            fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/categories/[id].tsx\",\n                            lineNumber: 40,\n                            columnNumber: 13\n                        }, undefined)\n                    }, i, false, {\n                        fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/categories/[id].tsx\",\n                        lineNumber: 39,\n                        columnNumber: 11\n                    }, undefined))\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/categories/[id].tsx\",\n                lineNumber: 37,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true);\n};\n_s(Category, \"F6vrTcaRMsgW/V4ljnr9mbaciec=\", false, function() {\n    return [\n        next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter\n    ];\n});\n_c = Category;\n/* harmony default export */ __webpack_exports__[\"default\"] = (Category);\nvar _c;\n$RefreshReg$(_c, \"Category\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9jYXRlZ29yaWVzL1tpZF0udHN4LmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7O0FBQUE7O0FBQXdFO0FBQ2pDO0FBQ2I7QUFDRztBQUcyQjtBQUl4RCxNQUFNUSxXQUFXLElBQU07O0lBQ3JCSCxxRUFBNkIsR0FBRztRQUM5QixnQkFBZ0I7SUFHbEI7SUFDQSxNQUFNLENBQUNPLE9BQU9DLFNBQVMsR0FBR1osK0NBQVFBLENBQVMsRUFBRTtJQUM3QyxNQUFNYSxTQUFTVixzREFBU0E7SUFDeEIsTUFBTVcsU0FBU0QsT0FBT0UsS0FBSyxDQUFDQyxFQUFFO0lBQzlCLE1BQU1DLFdBQVdmLGtEQUFXQSxDQUFDLE9BQU9ZLFNBQWdCO1FBQ2xELE1BQU1SLGtFQUFpQkEsQ0FBQ1EsUUFDckJJLElBQUksQ0FBQyxDQUFDQyxNQUFhO1lBQ2xCUCxTQUFTTyxJQUFJQyxJQUFJO1FBQ25CO0lBQ0osR0FBRyxFQUFFO0lBRUxuQixnREFBU0EsQ0FBQyxJQUFPO1FBQ2ZnQixTQUFTSDtJQUVYLEdBQUc7UUFBQ0E7S0FBTztJQUVYLHFCQUNFOzswQkFDRSw4REFBQ087Z0JBQUdDLFdBQVU7MEJBQStCOzs7Ozs7MEJBRzdDLDhEQUFDQzswQkFDRVosTUFBTWEsR0FBRyxDQUFDLENBQUNDLEdBQUdDLGtCQUNiLDhEQUFDQztrQ0FDQyw0RUFBQ3RCLGtEQUFJQTs0QkFBQ3VCLE1BQU07Z0NBQ1ZDLFVBQVU7Z0NBQ1ZkLE9BQU87b0NBQUNlLFNBQVNMLEVBQUVULEVBQUU7Z0NBQUE7NEJBQ3ZCO3NDQUNFLDRFQUFDZTswQ0FBR04sRUFBRU8sSUFBSTs7Ozs7Ozs7Ozs7dUJBTExOOzs7Ozs7Ozs7Ozs7QUFZbkI7R0F4Q01uQjs7UUFPV0osa0RBQVNBOzs7S0FQcEJJO0FBMENOLCtEQUFlQSxRQUFRQSxFQUFBIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vX05fRS8uL3BhZ2VzL2NhdGVnb3JpZXMvW2lkXS50c3g/YWVlYSJdLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgUmVhY3QsIHsgdXNlU3RhdGUsIHVzZUVmZmVjdCwgdXNlUmVmLCB1c2VDYWxsYmFjayB9IGZyb20gXCJyZWFjdFwiO1xuaW1wb3J0IHsgdXNlUm91dGVyIH0gZnJvbSAnbmV4dC9yb3V0ZXInXG5pbXBvcnQgYXhpb3MgZnJvbSAnYXhpb3MnO1xuaW1wb3J0IExpbmsgZnJvbSAnbmV4dC9saW5rJztcbmltcG9ydCBJbWFnZSBmcm9tICduZXh0L2ltYWdlJ1xuaW1wb3J0IEl0ZW0gZnJvbSAnLi4vLi4vbW9kZWxzL2l0ZW0nXG5pbXBvcnQgeyBnZXRJdGVtc0Rpc3BsYXlPbiB9IGZyb20gXCIuLi8uLi9hcGlDbGllbnQvaXRlbVwiXG5cblxuXG5jb25zdCBDYXRlZ29yeSA9ICgpID0+IHtcbiAgYXhpb3MuZGVmYXVsdHMuaGVhZGVycy5jb21tb24gPSB7XG4gICAgJ0NvbnRlbnQtVHlwZSc6ICdhcHBsaWNhdGlvbi9qc29uJyxcbiAgICAvLyAnWC1SZXF1ZXN0ZWQtV2l0aCc6ICdYTUxIdHRwUmVxdWVzdCcsXG4gICAgLy8gXCJYLUNTUkYtVG9rZW5cIjogY3NyZlRva2VuXG4gIH1cbiAgY29uc3QgW2l0ZW1zLCBzZXRJdGVtc10gPSB1c2VTdGF0ZTxJdGVtW10+KFtdKVxuICBjb25zdCByb3V0ZXIgPSB1c2VSb3V0ZXIoKVxuICBjb25zdCBwYXRoSWQgPSByb3V0ZXIucXVlcnkuaWRcbiAgY29uc3QgZ2V0SXRlbXMgPSB1c2VDYWxsYmFjayhhc3luYyAocGF0aElkOiBhbnkpID0+IHtcbiAgICBhd2FpdCBnZXRJdGVtc0Rpc3BsYXlPbihwYXRoSWQpXG4gICAgICAudGhlbigocmVzOiBhbnkpID0+IHtcbiAgICAgICAgc2V0SXRlbXMocmVzLmRhdGEpXG4gICAgICB9KVxuICB9LCBbXSlcbiAgXG4gIHVzZUVmZmVjdCgoKSAgPT4ge1xuICAgIGdldEl0ZW1zKHBhdGhJZCk7XG5cbiAgfSwgW3BhdGhJZF0pXG5cbiAgcmV0dXJuIChcbiAgICA8PlxuICAgICAgPGgxIGNsYXNzTmFtZT1cInRleHQtM3hsIGZvbnQtYm9sZCB1bmRlcmxpbmVcIj5cbiAgICAgICAgSGVsbG8gd29ybGQhXG4gICAgICA8L2gxPlxuICAgICAgPHVsPlxuICAgICAgICB7aXRlbXMubWFwKCh2LCBpKSA9PlxuICAgICAgICAgIDxsaSBrZXk9e2l9PlxuICAgICAgICAgICAgPExpbmsgaHJlZj17e1xuICAgICAgICAgICAgICBwYXRobmFtZTogXCIvaXRlbXMvW2l0ZW1faWRdXCIsXG4gICAgICAgICAgICAgIHF1ZXJ5OiB7aXRlbV9pZDogdi5pZH1cbiAgICAgICAgICAgIH19PlxuICAgICAgICAgICAgICA8YT57di5uYW1lfTwvYT5cbiAgICAgICAgICAgIDwvTGluaz5cbiAgICAgICAgICA8L2xpPlxuICAgICAgICApfVxuICAgICAgPC91bD5cbiAgICA8Lz5cbiAgKVxufVxuXG5leHBvcnQgZGVmYXVsdCBDYXRlZ29yeSJdLCJuYW1lcyI6WyJSZWFjdCIsInVzZVN0YXRlIiwidXNlRWZmZWN0IiwidXNlQ2FsbGJhY2siLCJ1c2VSb3V0ZXIiLCJheGlvcyIsIkxpbmsiLCJnZXRJdGVtc0Rpc3BsYXlPbiIsIkNhdGVnb3J5IiwiZGVmYXVsdHMiLCJoZWFkZXJzIiwiY29tbW9uIiwiaXRlbXMiLCJzZXRJdGVtcyIsInJvdXRlciIsInBhdGhJZCIsInF1ZXJ5IiwiaWQiLCJnZXRJdGVtcyIsInRoZW4iLCJyZXMiLCJkYXRhIiwiaDEiLCJjbGFzc05hbWUiLCJ1bCIsIm1hcCIsInYiLCJpIiwibGkiLCJocmVmIiwicGF0aG5hbWUiLCJpdGVtX2lkIiwiYSIsIm5hbWUiXSwic291cmNlUm9vdCI6IiJ9\n//# sourceURL=webpack-internal:///./pages/categories/[id].tsx\n"));

/***/ })

});