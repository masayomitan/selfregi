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

/***/ "./apiClient/item.ts":
/*!***************************!*\
  !*** ./apiClient/item.ts ***!
  \***************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"getItemDisplayOn\": function() { return /* binding */ getItemDisplayOn; },\n/* harmony export */   \"getItemsDisplayOn\": function() { return /* binding */ getItemsDisplayOn; },\n/* harmony export */   \"postItem\": function() { return /* binding */ postItem; }\n/* harmony export */ });\n/* harmony import */ var ___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! . */ \"./apiClient/index.ts\");\n\nconst apiBaseUrl = \"http://localhost:8000\";\nconst apiClient = new ___WEBPACK_IMPORTED_MODULE_0__.ApiClient(apiBaseUrl);\nconst getItemDisplayOn = (categoryId)=>{\n    return apiClient.get(\"/api/items/get/\");\n};\nconst getItemsDisplayOn = (categoryId)=>{\n    return apiClient.get(\"/api/categories/\" + categoryId + \"/items/get\");\n};\nconst postItem = (payload)=>{\n    return apiClient.post(\"/api/admin/items/post\", payload);\n};\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9hcGlDbGllbnQvaXRlbS50cy5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7O0FBQTJCO0FBRTNCLE1BQU1DLGFBQWFDLHVCQUFnQztBQUNuRCxNQUFNRyxZQUFZLElBQUlMLHdDQUFTQSxDQUFDQztBQUV6QixNQUFNSyxtQkFBbUIsQ0FBQ0MsYUFBdUI7SUFDdEQsT0FBT0YsVUFBVUcsR0FBRyxDQUNsQjtBQUVKLEVBQUM7QUFFTSxNQUFNQyxvQkFBb0IsQ0FBQ0YsYUFBdUI7SUFDdkQsT0FBT0YsVUFBVUcsR0FBRyxDQUNsQixxQkFBcUJELGFBQWE7QUFFdEMsRUFBQztBQUVNLE1BQU1HLFdBQVcsQ0FBQ0MsVUFBaUI7SUFDdEMsT0FBT04sVUFBVU8sSUFBSSxDQUNuQix5QkFDQUQ7QUFFTixFQUFDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vX05fRS8uL2FwaUNsaWVudC9pdGVtLnRzPzBiMGIiXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHtBcGlDbGllbnR9IGZyb20gJy4nXG5cbmNvbnN0IGFwaUJhc2VVcmwgPSBwcm9jZXNzLmVudi5ORVhUX1BVQkxJQ19CQVNFX1VSTFxuY29uc3QgYXBpQ2xpZW50ID0gbmV3IEFwaUNsaWVudChhcGlCYXNlVXJsKVxuXG5leHBvcnQgY29uc3QgZ2V0SXRlbURpc3BsYXlPbiA9IChjYXRlZ29yeUlkOiBzdHJpbmcpID0+IHtcbiAgcmV0dXJuIGFwaUNsaWVudC5nZXQoXG4gICAgJy9hcGkvaXRlbXMvZ2V0LycsXG4gIClcbn1cblxuZXhwb3J0IGNvbnN0IGdldEl0ZW1zRGlzcGxheU9uID0gKGNhdGVnb3J5SWQ6IHN0cmluZykgPT4ge1xuICByZXR1cm4gYXBpQ2xpZW50LmdldChcbiAgICAnL2FwaS9jYXRlZ29yaWVzLycgKyBjYXRlZ29yeUlkICsgJy9pdGVtcy9nZXQnLFxuICApXG59XG5cbmV4cG9ydCBjb25zdCBwb3N0SXRlbSA9IChwYXlsb2FkOiBhbnkpID0+IHtcbiAgICByZXR1cm4gYXBpQ2xpZW50LnBvc3QoXG4gICAgICAnL2FwaS9hZG1pbi9pdGVtcy9wb3N0JyxcbiAgICAgIHBheWxvYWRcbiAgICApXG59Il0sIm5hbWVzIjpbIkFwaUNsaWVudCIsImFwaUJhc2VVcmwiLCJwcm9jZXNzIiwiZW52IiwiTkVYVF9QVUJMSUNfQkFTRV9VUkwiLCJhcGlDbGllbnQiLCJnZXRJdGVtRGlzcGxheU9uIiwiY2F0ZWdvcnlJZCIsImdldCIsImdldEl0ZW1zRGlzcGxheU9uIiwicG9zdEl0ZW0iLCJwYXlsb2FkIiwicG9zdCJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./apiClient/item.ts\n"));

/***/ })

});