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

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"./node_modules/react/jsx-dev-runtime.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"./node_modules/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! next/router */ \"./node_modules/next/router.js\");\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(next_router__WEBPACK_IMPORTED_MODULE_2__);\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! axios */ \"./node_modules/axios/index.js\");\n/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! next/link */ \"./node_modules/next/link.js\");\n/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(next_link__WEBPACK_IMPORTED_MODULE_3__);\n/* harmony import */ var _apiClient_item__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../../apiClient/item */ \"./apiClient/item.ts\");\n\nvar _s = $RefreshSig$();\n\n\n\n\n\nconst Category = ()=>{\n    _s();\n    axios__WEBPACK_IMPORTED_MODULE_5__[\"default\"].defaults.headers.common = {\n        \"Content-Type\": \"application/json\"\n    };\n    const [item, setItem] = (0,react__WEBPACK_IMPORTED_MODULE_1__.useState)([]);\n    const router = (0,next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter)();\n    const pathId = router.query.id;\n    const getItems = (0,react__WEBPACK_IMPORTED_MODULE_1__.useCallback)(async (pathId)=>{\n        await (0,_apiClient_item__WEBPACK_IMPORTED_MODULE_4__.getItem)(pathId).then((res)=>{\n            setItems(res.data);\n        });\n    }, []);\n    (0,react__WEBPACK_IMPORTED_MODULE_1__.useEffect)(()=>{\n        getItems(pathId);\n    }, [\n        pathId\n    ]);\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.Fragment, {\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h1\", {\n                className: \"text-3xl font-bold underline\",\n                children: \"Hello world!\"\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 35,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"ul\", {\n                children: items.map((v, i)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"li\", {\n                        children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)((next_link__WEBPACK_IMPORTED_MODULE_3___default()), {\n                            href: {\n                                pathname: \"/items/[item_id]\",\n                                query: {\n                                    item_id: v.id\n                                }\n                            },\n                            children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"p\", {\n                                children: v.name\n                            }, void 0, false, {\n                                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                                lineNumber: 45,\n                                columnNumber: 15\n                            }, undefined)\n                        }, void 0, false, {\n                            fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                            lineNumber: 41,\n                            columnNumber: 13\n                        }, undefined)\n                    }, i, false, {\n                        fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                        lineNumber: 40,\n                        columnNumber: 11\n                    }, undefined))\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 38,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true);\n};\n_s(Category, \"m3CCcZ1KcTfDYAKTNzgCUgKY3Qw=\", false, function() {\n    return [\n        next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter\n    ];\n});\n_c = Category;\n/* harmony default export */ __webpack_exports__[\"default\"] = (Category);\nvar _c;\n$RefreshReg$(_c, \"Category\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9pdGVtcy9baWRdLnRzeC5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7OztBQUFBOztBQUF3RTtBQUNqQztBQUNiO0FBQ0c7QUFHaUI7QUFJOUMsTUFBTVEsV0FBVyxJQUFNOztJQUNyQkgscUVBQTZCLEdBQUc7UUFDOUIsZ0JBQWdCO0lBR2xCO0lBRUEsTUFBTSxDQUFDTyxNQUFNQyxRQUFRLEdBQUdaLCtDQUFRQSxDQUFTLEVBQUU7SUFDM0MsTUFBTWEsU0FBU1Ysc0RBQVNBO0lBQ3hCLE1BQU1XLFNBQVNELE9BQU9FLEtBQUssQ0FBQ0MsRUFBRTtJQUM5QixNQUFNQyxXQUFXZixrREFBV0EsQ0FBQyxPQUFPWSxTQUFnQjtRQUNsRCxNQUFNUix3REFBT0EsQ0FBQ1EsUUFDWEksSUFBSSxDQUFDLENBQUNDLE1BQWE7WUFDbEJDLFNBQVNELElBQUlFLElBQUk7UUFDbkI7SUFDSixHQUFHLEVBQUU7SUFFTHBCLGdEQUFTQSxDQUFDLElBQU87UUFDZmdCLFNBQVNIO0lBRVgsR0FBRztRQUFDQTtLQUFPO0lBRVgscUJBQ0U7OzBCQUNFLDhEQUFDUTtnQkFBR0MsV0FBVTswQkFBK0I7Ozs7OzswQkFHN0MsOERBQUNDOzBCQUNFQyxNQUFNQyxHQUFHLENBQUMsQ0FBQ0MsR0FBR0Msa0JBQ2IsOERBQUNDO2tDQUNDLDRFQUFDeEIsa0RBQUlBOzRCQUFDeUIsTUFBTTtnQ0FDVkMsVUFBVTtnQ0FDVmhCLE9BQU87b0NBQUNpQixTQUFTTCxFQUFFWCxFQUFFO2dDQUFBOzRCQUN2QjtzQ0FDRSw0RUFBQ2lCOzBDQUFHTixFQUFFTyxJQUFJOzs7Ozs7Ozs7Ozt1QkFMTE47Ozs7Ozs7Ozs7OztBQVluQjtHQXpDTXJCOztRQVFXSixrREFBU0E7OztLQVJwQkk7QUEyQ04sK0RBQWVBLFFBQVFBLEVBQUEiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vcGFnZXMvaXRlbXMvW2lkXS50c3g/ODM3MCJdLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgUmVhY3QsIHsgdXNlU3RhdGUsIHVzZUVmZmVjdCwgdXNlUmVmLCB1c2VDYWxsYmFjayB9IGZyb20gXCJyZWFjdFwiO1xuaW1wb3J0IHsgdXNlUm91dGVyIH0gZnJvbSAnbmV4dC9yb3V0ZXInXG5pbXBvcnQgYXhpb3MgZnJvbSAnYXhpb3MnO1xuaW1wb3J0IExpbmsgZnJvbSAnbmV4dC9saW5rJztcbmltcG9ydCBJbWFnZSBmcm9tICduZXh0L2ltYWdlJ1xuaW1wb3J0IEl0ZW0gZnJvbSAnLi4vLi4vbW9kZWxzL2l0ZW0nXG5pbXBvcnQgeyBnZXRJdGVtIH0gZnJvbSBcIi4uLy4uL2FwaUNsaWVudC9pdGVtXCJcblxuXG5cbmNvbnN0IENhdGVnb3J5ID0gKCkgPT4ge1xuICBheGlvcy5kZWZhdWx0cy5oZWFkZXJzLmNvbW1vbiA9IHtcbiAgICAnQ29udGVudC1UeXBlJzogJ2FwcGxpY2F0aW9uL2pzb24nLFxuICAgIC8vICdYLVJlcXVlc3RlZC1XaXRoJzogJ1hNTEh0dHBSZXF1ZXN0JyxcbiAgICAvLyBcIlgtQ1NSRi1Ub2tlblwiOiBjc3JmVG9rZW5cbiAgfVxuXG4gIGNvbnN0IFtpdGVtLCBzZXRJdGVtXSA9IHVzZVN0YXRlPEl0ZW1bXT4oW10pXG4gIGNvbnN0IHJvdXRlciA9IHVzZVJvdXRlcigpXG4gIGNvbnN0IHBhdGhJZCA9IHJvdXRlci5xdWVyeS5pZFxuICBjb25zdCBnZXRJdGVtcyA9IHVzZUNhbGxiYWNrKGFzeW5jIChwYXRoSWQ6IGFueSkgPT4ge1xuICAgIGF3YWl0IGdldEl0ZW0ocGF0aElkKVxuICAgICAgLnRoZW4oKHJlczogYW55KSA9PiB7XG4gICAgICAgIHNldEl0ZW1zKHJlcy5kYXRhKVxuICAgICAgfSlcbiAgfSwgW10pXG4gIFxuICB1c2VFZmZlY3QoKCkgID0+IHtcbiAgICBnZXRJdGVtcyhwYXRoSWQpO1xuXG4gIH0sIFtwYXRoSWRdKVxuXG4gIHJldHVybiAoXG4gICAgPD5cbiAgICAgIDxoMSBjbGFzc05hbWU9XCJ0ZXh0LTN4bCBmb250LWJvbGQgdW5kZXJsaW5lXCI+XG4gICAgICAgIEhlbGxvIHdvcmxkIVxuICAgICAgPC9oMT5cbiAgICAgIDx1bD5cbiAgICAgICAge2l0ZW1zLm1hcCgodiwgaSkgPT5cbiAgICAgICAgICA8bGkga2V5PXtpfT5cbiAgICAgICAgICAgIDxMaW5rIGhyZWY9e3tcbiAgICAgICAgICAgICAgcGF0aG5hbWU6IFwiL2l0ZW1zL1tpdGVtX2lkXVwiLFxuICAgICAgICAgICAgICBxdWVyeToge2l0ZW1faWQ6IHYuaWR9XG4gICAgICAgICAgICB9fT5cbiAgICAgICAgICAgICAgPHA+e3YubmFtZX08L3A+XG4gICAgICAgICAgICA8L0xpbms+XG4gICAgICAgICAgPC9saT5cbiAgICAgICAgKX1cbiAgICAgIDwvdWw+XG4gICAgPC8+XG4gIClcbn1cblxuZXhwb3J0IGRlZmF1bHQgQ2F0ZWdvcnkiXSwibmFtZXMiOlsiUmVhY3QiLCJ1c2VTdGF0ZSIsInVzZUVmZmVjdCIsInVzZUNhbGxiYWNrIiwidXNlUm91dGVyIiwiYXhpb3MiLCJMaW5rIiwiZ2V0SXRlbSIsIkNhdGVnb3J5IiwiZGVmYXVsdHMiLCJoZWFkZXJzIiwiY29tbW9uIiwiaXRlbSIsInNldEl0ZW0iLCJyb3V0ZXIiLCJwYXRoSWQiLCJxdWVyeSIsImlkIiwiZ2V0SXRlbXMiLCJ0aGVuIiwicmVzIiwic2V0SXRlbXMiLCJkYXRhIiwiaDEiLCJjbGFzc05hbWUiLCJ1bCIsIml0ZW1zIiwibWFwIiwidiIsImkiLCJsaSIsImhyZWYiLCJwYXRobmFtZSIsIml0ZW1faWQiLCJwIiwibmFtZSJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./pages/items/[id].tsx\n"));

/***/ })

});