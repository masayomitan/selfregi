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

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"./node_modules/react/jsx-dev-runtime.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"./node_modules/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! next/router */ \"./node_modules/next/router.js\");\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(next_router__WEBPACK_IMPORTED_MODULE_2__);\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! axios */ \"./node_modules/axios/index.js\");\n\nvar _s = $RefreshSig$();\n\n\n\nconst Category = ()=>{\n    _s();\n    axios__WEBPACK_IMPORTED_MODULE_3__[\"default\"].defaults.headers.common = {\n        \"Content-Type\": \"application/json\"\n    };\n    const [item, setItem] = (0,react__WEBPACK_IMPORTED_MODULE_1__.useState)();\n    const router = (0,next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter)();\n    const pathId = router.query.id;\n    console.log(pathId);\n    const getItem = (pathId)=>{\n        getItem(pathId).then((res)=>{\n            console.log(res.data);\n            setItem(res.data);\n        });\n    };\n    (0,react__WEBPACK_IMPORTED_MODULE_1__.useEffect)(()=>{\n        getItem(pathId);\n    }, [\n        pathId\n    ]);\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.Fragment, {\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h1\", {\n                className: \"text-3xl font-bold underline\",\n                children: \"Hello world!\"\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 36,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"ul\", {}, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 39,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true);\n};\n_s(Category, \"5kvJp9WwviAjjq+vXPHKNVovkRs=\", false, function() {\n    return [\n        next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter\n    ];\n});\n_c = Category;\n/* harmony default export */ __webpack_exports__[\"default\"] = (Category);\nvar _c;\n$RefreshReg$(_c, \"Category\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9pdGVtcy9baWRdLnRzeC5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7OztBQUFBOztBQUF3RTtBQUNqQztBQUNiO0FBUTFCLE1BQU1LLFdBQVcsSUFBTTs7SUFDckJELHFFQUE2QixHQUFHO1FBQzlCLGdCQUFnQjtJQUdsQjtJQUVBLE1BQU0sQ0FBQ0ssTUFBTUMsUUFBUSxHQUFHVCwrQ0FBUUE7SUFDaEMsTUFBTVUsU0FBU1Isc0RBQVNBO0lBQ3hCLE1BQU1TLFNBQVNELE9BQU9FLEtBQUssQ0FBQ0MsRUFBRTtJQUM5QkMsUUFBUUMsR0FBRyxDQUFDSjtJQUNaLE1BQU1LLFVBQVUsQ0FBQ0wsU0FBZ0I7UUFDL0JLLFFBQVFMLFFBQ0xNLElBQUksQ0FBQyxDQUFDQyxNQUFhO1lBQ2xCSixRQUFRQyxHQUFHLENBQUNHLElBQUlDLElBQUk7WUFDcEJWLFFBQVFTLElBQUlDLElBQUk7UUFDbEI7SUFDSjtJQUVBbEIsZ0RBQVNBLENBQUMsSUFBTztRQUNmZSxRQUFRTDtJQUNWLEdBQUc7UUFBQ0E7S0FBTztJQUVYLHFCQUNFOzswQkFDRSw4REFBQ1M7Z0JBQUdDLFdBQVU7MEJBQStCOzs7Ozs7MEJBRzdDLDhEQUFDQzs7Ozs7OztBQWNQO0dBMUNNbEI7O1FBUVdGLGtEQUFTQTs7O0tBUnBCRTtBQTRDTiwrREFBZUEsUUFBUUEsRUFBQSIsInNvdXJjZXMiOlsid2VicGFjazovL19OX0UvLi9wYWdlcy9pdGVtcy9baWRdLnRzeD84MzcwIl0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCBSZWFjdCwgeyB1c2VTdGF0ZSwgdXNlRWZmZWN0LCB1c2VSZWYsIHVzZUNhbGxiYWNrIH0gZnJvbSBcInJlYWN0XCI7XG5pbXBvcnQgeyB1c2VSb3V0ZXIgfSBmcm9tICduZXh0L3JvdXRlcidcbmltcG9ydCBheGlvcyBmcm9tICdheGlvcyc7XG5pbXBvcnQgTGluayBmcm9tICduZXh0L2xpbmsnO1xuaW1wb3J0IEltYWdlIGZyb20gJ25leHQvaW1hZ2UnXG5pbXBvcnQgSXRlbSBmcm9tICcuLi8uLi9tb2RlbHMvaXRlbSdcbmltcG9ydCB7IGdldEl0ZW1hIH0gZnJvbSBcIi4uLy4uL2FwaUNsaWVudC9pdGVtXCJcblxuXG5cbmNvbnN0IENhdGVnb3J5ID0gKCkgPT4ge1xuICBheGlvcy5kZWZhdWx0cy5oZWFkZXJzLmNvbW1vbiA9IHtcbiAgICAnQ29udGVudC1UeXBlJzogJ2FwcGxpY2F0aW9uL2pzb24nLFxuICAgIC8vICdYLVJlcXVlc3RlZC1XaXRoJzogJ1hNTEh0dHBSZXF1ZXN0JyxcbiAgICAvLyBcIlgtQ1NSRi1Ub2tlblwiOiBjc3JmVG9rZW5cbiAgfVxuXG4gIGNvbnN0IFtpdGVtLCBzZXRJdGVtXSA9IHVzZVN0YXRlPEl0ZW0+KClcbiAgY29uc3Qgcm91dGVyID0gdXNlUm91dGVyKClcbiAgY29uc3QgcGF0aElkID0gcm91dGVyLnF1ZXJ5LmlkXG4gIGNvbnNvbGUubG9nKHBhdGhJZClcbiAgY29uc3QgZ2V0SXRlbSA9IChwYXRoSWQ6IGFueSkgPT4ge1xuICAgIGdldEl0ZW0ocGF0aElkKVxuICAgICAgLnRoZW4oKHJlczogYW55KSA9PiB7XG4gICAgICAgIGNvbnNvbGUubG9nKHJlcy5kYXRhKVxuICAgICAgICBzZXRJdGVtKHJlcy5kYXRhKVxuICAgICAgfSlcbiAgfVxuICBcbiAgdXNlRWZmZWN0KCgpICA9PiB7XG4gICAgZ2V0SXRlbShwYXRoSWQpO1xuICB9LCBbcGF0aElkXSlcblxuICByZXR1cm4gKFxuICAgIDw+XG4gICAgICA8aDEgY2xhc3NOYW1lPVwidGV4dC0zeGwgZm9udC1ib2xkIHVuZGVybGluZVwiPlxuICAgICAgICBIZWxsbyB3b3JsZCFcbiAgICAgIDwvaDE+XG4gICAgICA8dWw+XG4gICAgICAgIHsvKiB7aXRlbS5tYXAoKHYsIGkpID0+XG4gICAgICAgICAgPGxpIGtleT17aX0+XG4gICAgICAgICAgICA8TGluayBocmVmPXt7XG4gICAgICAgICAgICAgIHBhdGhuYW1lOiBcIi9pdGVtcy9baXRlbV9pZF1cIixcbiAgICAgICAgICAgICAgcXVlcnk6IHtpdGVtX2lkOiB2LmlkfVxuICAgICAgICAgICAgfX0+XG4gICAgICAgICAgICAgIDxwPnt2Lm5hbWV9PC9wPlxuICAgICAgICAgICAgPC9MaW5rPlxuICAgICAgICAgIDwvbGk+XG4gICAgICAgICl9ICovfVxuICAgICAgPC91bD5cbiAgICA8Lz5cbiAgKVxufVxuXG5leHBvcnQgZGVmYXVsdCBDYXRlZ29yeSJdLCJuYW1lcyI6WyJSZWFjdCIsInVzZVN0YXRlIiwidXNlRWZmZWN0IiwidXNlUm91dGVyIiwiYXhpb3MiLCJDYXRlZ29yeSIsImRlZmF1bHRzIiwiaGVhZGVycyIsImNvbW1vbiIsIml0ZW0iLCJzZXRJdGVtIiwicm91dGVyIiwicGF0aElkIiwicXVlcnkiLCJpZCIsImNvbnNvbGUiLCJsb2ciLCJnZXRJdGVtIiwidGhlbiIsInJlcyIsImRhdGEiLCJoMSIsImNsYXNzTmFtZSIsInVsIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./pages/items/[id].tsx\n"));

/***/ })

});