"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
self["webpackHotUpdate_N_E"]("pages/items/[id]",{

/***/ "./pages/items/[id].tsx":
/*!******************************!*\
  !*** ./pages/items/[id].tsx ***!
  \******************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"./node_modules/react/jsx-dev-runtime.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"./node_modules/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! next/router */ \"./node_modules/next/router.js\");\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(next_router__WEBPACK_IMPORTED_MODULE_2__);\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! axios */ \"./node_modules/axios/index.js\");\n/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! next/link */ \"./node_modules/next/link.js\");\n/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(next_link__WEBPACK_IMPORTED_MODULE_3__);\n/* harmony import */ var _apiClient_item__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../../apiClient/item */ \"./apiClient/item.ts\");\n\nvar _s = $RefreshSig$();\n\n\n\n\n\nconst Category = ()=>{\n    _s();\n    axios__WEBPACK_IMPORTED_MODULE_5__[\"default\"].defaults.headers.common = {\n        \"Content-Type\": \"application/json\"\n    };\n    const [item, setItem] = (0,react__WEBPACK_IMPORTED_MODULE_1__.useState)([]);\n    const router = (0,next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter)();\n    const pathId = router.query.id;\n    const getItems = (0,react__WEBPACK_IMPORTED_MODULE_1__.useCallback)(async (pathId)=>{\n        await (0,_apiClient_item__WEBPACK_IMPORTED_MODULE_4__.getItem)(pathId).then((res)=>{\n            setItem(res.data);\n        });\n    }, []);\n    (0,react__WEBPACK_IMPORTED_MODULE_1__.useEffect)(()=>{\n        getItems(pathId);\n    }, [\n        pathId\n    ]);\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.Fragment, {\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h1\", {\n                className: \"text-3xl font-bold underline\",\n                children: \"Hello world!\"\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 35,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"ul\", {\n                children: items.map((v, i)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"li\", {\n                        children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)((next_link__WEBPACK_IMPORTED_MODULE_3___default()), {\n                            href: {\n                                pathname: \"/items/[item_id]\",\n                                query: {\n                                    item_id: v.id\n                                }\n                            },\n                            children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"p\", {\n                                children: v.name\n                            }, void 0, false, {\n                                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                                lineNumber: 45,\n                                columnNumber: 15\n                            }, undefined)\n                        }, void 0, false, {\n                            fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                            lineNumber: 41,\n                            columnNumber: 13\n                        }, undefined)\n                    }, i, false, {\n                        fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                        lineNumber: 40,\n                        columnNumber: 11\n                    }, undefined))\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 38,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true);\n};\n_s(Category, \"m3CCcZ1KcTfDYAKTNzgCUgKY3Qw=\", false, function() {\n    return [\n        next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter\n    ];\n});\n_c = Category;\n/* harmony default export */ __webpack_exports__[\"default\"] = (Category);\nvar _c;\n$RefreshReg$(_c, \"Category\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9pdGVtcy9baWRdLnRzeC5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7OztBQUFBOztBQUF3RTtBQUNqQztBQUNiO0FBQ0c7QUFHaUI7QUFJOUMsTUFBTVEsV0FBVyxJQUFNOztJQUNyQkgscUVBQTZCLEdBQUc7UUFDOUIsZ0JBQWdCO0lBR2xCO0lBRUEsTUFBTSxDQUFDTyxNQUFNQyxRQUFRLEdBQUdaLCtDQUFRQSxDQUFPLEVBQUU7SUFDekMsTUFBTWEsU0FBU1Ysc0RBQVNBO0lBQ3hCLE1BQU1XLFNBQVNELE9BQU9FLEtBQUssQ0FBQ0MsRUFBRTtJQUM5QixNQUFNQyxXQUFXZixrREFBV0EsQ0FBQyxPQUFPWSxTQUFnQjtRQUNsRCxNQUFNUix3REFBT0EsQ0FBQ1EsUUFDWEksSUFBSSxDQUFDLENBQUNDLE1BQWE7WUFDbEJQLFFBQVFPLElBQUlDLElBQUk7UUFDbEI7SUFDSixHQUFHLEVBQUU7SUFFTG5CLGdEQUFTQSxDQUFDLElBQU87UUFDZmdCLFNBQVNIO0lBRVgsR0FBRztRQUFDQTtLQUFPO0lBRVgscUJBQ0U7OzBCQUNFLDhEQUFDTztnQkFBR0MsV0FBVTswQkFBK0I7Ozs7OzswQkFHN0MsOERBQUNDOzBCQUNFQyxNQUFNQyxHQUFHLENBQUMsQ0FBQ0MsR0FBR0Msa0JBQ2IsOERBQUNDO2tDQUNDLDRFQUFDdkIsa0RBQUlBOzRCQUFDd0IsTUFBTTtnQ0FDVkMsVUFBVTtnQ0FDVmYsT0FBTztvQ0FBQ2dCLFNBQVNMLEVBQUVWLEVBQUU7Z0NBQUE7NEJBQ3ZCO3NDQUNFLDRFQUFDZ0I7MENBQUdOLEVBQUVPLElBQUk7Ozs7Ozs7Ozs7O3VCQUxMTjs7Ozs7Ozs7Ozs7O0FBWW5CO0dBekNNcEI7O1FBUVdKLGtEQUFTQTs7O0tBUnBCSTtBQTJDTiwrREFBZUEsUUFBUUEsRUFBQSIsInNvdXJjZXMiOlsid2VicGFjazovL19OX0UvLi9wYWdlcy9pdGVtcy9baWRdLnRzeD84MzcwIl0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCBSZWFjdCwgeyB1c2VTdGF0ZSwgdXNlRWZmZWN0LCB1c2VSZWYsIHVzZUNhbGxiYWNrIH0gZnJvbSBcInJlYWN0XCI7XG5pbXBvcnQgeyB1c2VSb3V0ZXIgfSBmcm9tICduZXh0L3JvdXRlcidcbmltcG9ydCBheGlvcyBmcm9tICdheGlvcyc7XG5pbXBvcnQgTGluayBmcm9tICduZXh0L2xpbmsnO1xuaW1wb3J0IEltYWdlIGZyb20gJ25leHQvaW1hZ2UnXG5pbXBvcnQgSXRlbSBmcm9tICcuLi8uLi9tb2RlbHMvaXRlbSdcbmltcG9ydCB7IGdldEl0ZW0gfSBmcm9tIFwiLi4vLi4vYXBpQ2xpZW50L2l0ZW1cIlxuXG5cblxuY29uc3QgQ2F0ZWdvcnkgPSAoKSA9PiB7XG4gIGF4aW9zLmRlZmF1bHRzLmhlYWRlcnMuY29tbW9uID0ge1xuICAgICdDb250ZW50LVR5cGUnOiAnYXBwbGljYXRpb24vanNvbicsXG4gICAgLy8gJ1gtUmVxdWVzdGVkLVdpdGgnOiAnWE1MSHR0cFJlcXVlc3QnLFxuICAgIC8vIFwiWC1DU1JGLVRva2VuXCI6IGNzcmZUb2tlblxuICB9XG5cbiAgY29uc3QgW2l0ZW0sIHNldEl0ZW1dID0gdXNlU3RhdGU8SXRlbT4oW10pXG4gIGNvbnN0IHJvdXRlciA9IHVzZVJvdXRlcigpXG4gIGNvbnN0IHBhdGhJZCA9IHJvdXRlci5xdWVyeS5pZFxuICBjb25zdCBnZXRJdGVtcyA9IHVzZUNhbGxiYWNrKGFzeW5jIChwYXRoSWQ6IGFueSkgPT4ge1xuICAgIGF3YWl0IGdldEl0ZW0ocGF0aElkKVxuICAgICAgLnRoZW4oKHJlczogYW55KSA9PiB7XG4gICAgICAgIHNldEl0ZW0ocmVzLmRhdGEpXG4gICAgICB9KVxuICB9LCBbXSlcbiAgXG4gIHVzZUVmZmVjdCgoKSAgPT4ge1xuICAgIGdldEl0ZW1zKHBhdGhJZCk7XG5cbiAgfSwgW3BhdGhJZF0pXG5cbiAgcmV0dXJuIChcbiAgICA8PlxuICAgICAgPGgxIGNsYXNzTmFtZT1cInRleHQtM3hsIGZvbnQtYm9sZCB1bmRlcmxpbmVcIj5cbiAgICAgICAgSGVsbG8gd29ybGQhXG4gICAgICA8L2gxPlxuICAgICAgPHVsPlxuICAgICAgICB7aXRlbXMubWFwKCh2LCBpKSA9PlxuICAgICAgICAgIDxsaSBrZXk9e2l9PlxuICAgICAgICAgICAgPExpbmsgaHJlZj17e1xuICAgICAgICAgICAgICBwYXRobmFtZTogXCIvaXRlbXMvW2l0ZW1faWRdXCIsXG4gICAgICAgICAgICAgIHF1ZXJ5OiB7aXRlbV9pZDogdi5pZH1cbiAgICAgICAgICAgIH19PlxuICAgICAgICAgICAgICA8cD57di5uYW1lfTwvcD5cbiAgICAgICAgICAgIDwvTGluaz5cbiAgICAgICAgICA8L2xpPlxuICAgICAgICApfVxuICAgICAgPC91bD5cbiAgICA8Lz5cbiAgKVxufVxuXG5leHBvcnQgZGVmYXVsdCBDYXRlZ29yeSJdLCJuYW1lcyI6WyJSZWFjdCIsInVzZVN0YXRlIiwidXNlRWZmZWN0IiwidXNlQ2FsbGJhY2siLCJ1c2VSb3V0ZXIiLCJheGlvcyIsIkxpbmsiLCJnZXRJdGVtIiwiQ2F0ZWdvcnkiLCJkZWZhdWx0cyIsImhlYWRlcnMiLCJjb21tb24iLCJpdGVtIiwic2V0SXRlbSIsInJvdXRlciIsInBhdGhJZCIsInF1ZXJ5IiwiaWQiLCJnZXRJdGVtcyIsInRoZW4iLCJyZXMiLCJkYXRhIiwiaDEiLCJjbGFzc05hbWUiLCJ1bCIsIml0ZW1zIiwibWFwIiwidiIsImkiLCJsaSIsImhyZWYiLCJwYXRobmFtZSIsIml0ZW1faWQiLCJwIiwibmFtZSJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./pages/items/[id].tsx\n"));

/***/ })

});