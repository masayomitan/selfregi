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

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"./node_modules/react/jsx-dev-runtime.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"./node_modules/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! next/router */ \"./node_modules/next/router.js\");\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(next_router__WEBPACK_IMPORTED_MODULE_2__);\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! axios */ \"./node_modules/axios/index.js\");\n/* harmony import */ var _apiClient_item__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../../apiClient/item */ \"./apiClient/item.ts\");\n\nvar _s = $RefreshSig$();\n\n\n\n// import Category from '../../models/category'\n\nconst Category = ()=>{\n    _s();\n    axios__WEBPACK_IMPORTED_MODULE_4__[\"default\"].defaults.headers.common = {\n        \"Content-Type\": \"application/json\"\n    };\n    // const [categories, setCategories] = useState<Category[]>([])\n    const router = (0,next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter)();\n    console.log(router.query.id);\n    const getItems = (0,react__WEBPACK_IMPORTED_MODULE_1__.useCallback)(async ()=>{\n        await (0,_apiClient_item__WEBPACK_IMPORTED_MODULE_3__.getItemsDisplayOn)().then((res)=>{\n            console.log(res);\n        });\n    }, []);\n    (0,react__WEBPACK_IMPORTED_MODULE_1__.useEffect)(()=>{\n        getItems();\n    }, []);\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.Fragment, {\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h1\", {\n                className: \"text-3xl font-bold underline\",\n                children: \"Hello world!\"\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/categories/[id].tsx\",\n                lineNumber: 35,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"ul\", {}, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/categories/[id].tsx\",\n                lineNumber: 38,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true);\n};\n_s(Category, \"Q42Gfu38xnSolt1oOR+nWW1Mcx8=\", false, function() {\n    return [\n        next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter\n    ];\n});\n_c = Category;\n/* harmony default export */ __webpack_exports__[\"default\"] = (Category);\nvar _c;\n$RefreshReg$(_c, \"Category\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9jYXRlZ29yaWVzL1tpZF0udHN4LmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7OztBQUFBOztBQUF3RTtBQUNqQztBQUNiO0FBRzFCLCtDQUErQztBQUNTO0FBSXhELE1BQU1NLFdBQVcsSUFBTTs7SUFDckJGLHFFQUE2QixHQUFHO1FBQzlCLGdCQUFnQjtJQUdsQjtJQUNBLCtEQUErRDtJQUMvRCxNQUFNTSxTQUFTUCxzREFBU0E7SUFFeEJRLFFBQVFDLEdBQUcsQ0FBQ0YsT0FBT0csS0FBSyxDQUFDQyxFQUFFO0lBQzNCLE1BQU1DLFdBQVdiLGtEQUFXQSxDQUFDLFVBQVk7UUFDdkMsTUFBTUcsa0VBQWlCQSxHQUNwQlcsSUFBSSxDQUFDLENBQUNDLE1BQVE7WUFDYk4sUUFBUUMsR0FBRyxDQUFDSztRQUNkO0lBQ0osR0FBRyxFQUFFO0lBRUxoQixnREFBU0EsQ0FBQyxJQUFPO1FBQ2ZjO0lBRUYsR0FBRyxFQUFFO0lBRUwscUJBQ0U7OzBCQUNFLDhEQUFDRztnQkFBR0MsV0FBVTswQkFBK0I7Ozs7OzswQkFHN0MsOERBQUNDOzs7Ozs7O0FBY1A7R0F6Q01kOztRQU9XSCxrREFBU0E7OztLQVBwQkc7QUEyQ04sK0RBQWVBLFFBQVFBLEVBQUEiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vcGFnZXMvY2F0ZWdvcmllcy9baWRdLnRzeD9hZWVhIl0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCBSZWFjdCwgeyB1c2VTdGF0ZSwgdXNlRWZmZWN0LCB1c2VSZWYsIHVzZUNhbGxiYWNrIH0gZnJvbSBcInJlYWN0XCI7XG5pbXBvcnQgeyB1c2VSb3V0ZXIgfSBmcm9tICduZXh0L3JvdXRlcidcbmltcG9ydCBheGlvcyBmcm9tICdheGlvcyc7XG5pbXBvcnQgTGluayBmcm9tICduZXh0L2xpbmsnO1xuaW1wb3J0IEltYWdlIGZyb20gJ25leHQvaW1hZ2UnXG4vLyBpbXBvcnQgQ2F0ZWdvcnkgZnJvbSAnLi4vLi4vbW9kZWxzL2NhdGVnb3J5J1xuaW1wb3J0IHsgZ2V0SXRlbXNEaXNwbGF5T24gfSBmcm9tIFwiLi4vLi4vYXBpQ2xpZW50L2l0ZW1cIlxuXG5cblxuY29uc3QgQ2F0ZWdvcnkgPSAoKSA9PiB7XG4gIGF4aW9zLmRlZmF1bHRzLmhlYWRlcnMuY29tbW9uID0ge1xuICAgICdDb250ZW50LVR5cGUnOiAnYXBwbGljYXRpb24vanNvbicsXG4gICAgLy8gJ1gtUmVxdWVzdGVkLVdpdGgnOiAnWE1MSHR0cFJlcXVlc3QnLFxuICAgIC8vIFwiWC1DU1JGLVRva2VuXCI6IGNzcmZUb2tlblxuICB9XG4gIC8vIGNvbnN0IFtjYXRlZ29yaWVzLCBzZXRDYXRlZ29yaWVzXSA9IHVzZVN0YXRlPENhdGVnb3J5W10+KFtdKVxuICBjb25zdCByb3V0ZXIgPSB1c2VSb3V0ZXIoKVxuICBcbiAgY29uc29sZS5sb2cocm91dGVyLnF1ZXJ5LmlkKVxuICBjb25zdCBnZXRJdGVtcyA9IHVzZUNhbGxiYWNrKGFzeW5jICgpID0+IHtcbiAgICBhd2FpdCBnZXRJdGVtc0Rpc3BsYXlPbigpXG4gICAgICAudGhlbigocmVzKSA9PiB7XG4gICAgICAgIGNvbnNvbGUubG9nKHJlcylcbiAgICAgIH0pXG4gIH0sIFtdKVxuICBcbiAgdXNlRWZmZWN0KCgpICA9PiB7XG4gICAgZ2V0SXRlbXMoKTtcblxuICB9LCBbXSlcblxuICByZXR1cm4gKFxuICAgIDw+XG4gICAgICA8aDEgY2xhc3NOYW1lPVwidGV4dC0zeGwgZm9udC1ib2xkIHVuZGVybGluZVwiPlxuICAgICAgICBIZWxsbyB3b3JsZCFcbiAgICAgIDwvaDE+XG4gICAgICA8dWw+XG4gICAgICAgIHsvKiB7Y2F0ZWdvcmllcy5tYXAoKHYsIGkpID0+XG4gICAgICAgICAgPGxpIGtleT17aX0+XG4gICAgICAgICAgICA8TGluayBocmVmPXt7XG4gICAgICAgICAgICAgIHBhdGhuYW1lOiBcIi9jYXRlZ29yaWVzL1tjYXRlZ29yeV9pZF1cIixcbiAgICAgICAgICAgICAgcXVlcnk6IHtjYXRlZ29yeV9pZDogdi5pZH1cbiAgICAgICAgICAgIH19PlxuICAgICAgICAgICAgICA8YT57di5uYW1lfTwvYT5cbiAgICAgICAgICAgIDwvTGluaz5cbiAgICAgICAgICA8L2xpPlxuICAgICAgICApfSAqL31cbiAgICAgIDwvdWw+XG4gICAgPC8+XG4gIClcbn1cblxuZXhwb3J0IGRlZmF1bHQgQ2F0ZWdvcnkiXSwibmFtZXMiOlsiUmVhY3QiLCJ1c2VFZmZlY3QiLCJ1c2VDYWxsYmFjayIsInVzZVJvdXRlciIsImF4aW9zIiwiZ2V0SXRlbXNEaXNwbGF5T24iLCJDYXRlZ29yeSIsImRlZmF1bHRzIiwiaGVhZGVycyIsImNvbW1vbiIsInJvdXRlciIsImNvbnNvbGUiLCJsb2ciLCJxdWVyeSIsImlkIiwiZ2V0SXRlbXMiLCJ0aGVuIiwicmVzIiwiaDEiLCJjbGFzc05hbWUiLCJ1bCJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./pages/categories/[id].tsx\n"));

/***/ })

});