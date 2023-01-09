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

/***/ "./apiClient/index.ts":
/*!****************************!*\
  !*** ./apiClient/index.ts ***!
  \****************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"ApiClient\": function() { return /* binding */ ApiClient; }\n/* harmony export */ });\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! axios */ \"./node_modules/axios/index.js\");\n\nclass ApiClient {\n    async get(path) {\n        let params = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : {};\n        try {\n            const result = await this.axiosInstance.get(path, {\n                params\n            });\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    async post(path) {\n        let params = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : {};\n        try {\n            const result = await this.axiosInstance.post(path, {\n                params\n            });\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    async put(path) {\n        let params = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : {};\n        try {\n            const result = await this.axiosInstance.put(path, {\n                params\n            });\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    async patch(path) {\n        let params = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : {};\n        try {\n            const result = await this.axiosInstance.patch(path, {\n                params\n            });\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    async delete(path) {\n        try {\n            const result = await this.axiosInstance.delete(path);\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    successPromise(data) {\n        return Promise.resolve({\n            data,\n            isSuccess: true\n        });\n    }\n    failurePromise(error) {\n        return Promise.reject({\n            error,\n            isSuccess: false\n        });\n    }\n    constructor(baseURL = \"\"){\n        this.axiosInstance = axios__WEBPACK_IMPORTED_MODULE_0__[\"default\"].create({\n            baseURL\n        });\n        this.axiosInstance.interceptors.request.use(async (config)=>{\n            if (localStorage.getItem(\"token\")) {\n                config.headers.authorization = \"Bearer \" + localStorage.getItem(\"token\");\n            }\n            return config;\n        }, (err)=>{\n            return Promise.reject(err);\n        });\n    }\n}\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,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LCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vYXBpQ2xpZW50L2luZGV4LnRzPzQ5MjkiXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IGF4aW9zLCB7QXhpb3NFcnJvciwgQXhpb3NJbnN0YW5jZSwgQXhpb3NSZXF1ZXN0Q29uZmlnfSBmcm9tICdheGlvcydcblxuZXhwb3J0IGNsYXNzIEFwaUNsaWVudCB7XG4gIGF4aW9zSW5zdGFuY2U6IEF4aW9zSW5zdGFuY2VcbiAgY29uc3RydWN0b3IoYmFzZVVSTCA9ICcnKSB7XG4gICAgdGhpcy5heGlvc0luc3RhbmNlID0gYXhpb3MuY3JlYXRlKHtcbiAgICAgIGJhc2VVUkxcbiAgICB9KVxuXG4gICAgdGhpcy5heGlvc0luc3RhbmNlLmludGVyY2VwdG9ycy5yZXF1ZXN0LnVzZShcbiAgICAgIGFzeW5jIChjb25maWc6IEF4aW9zUmVxdWVzdENvbmZpZyB8IGFueSkgPT4ge1xuICAgICAgICBpZiAobG9jYWxTdG9yYWdlLmdldEl0ZW0oJ3Rva2VuJykpIHtcbiAgICAgICAgICBjb25maWcuaGVhZGVycy5hdXRob3JpemF0aW9uID0gJ0JlYXJlciAnICsgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ3Rva2VuJylcbiAgICAgICAgfVxuICAgICAgICByZXR1cm4gY29uZmlnXG4gICAgICB9LCAoZXJyOiBBeGlvc0Vycm9yKSA9PiB7XG4gICAgICAgIHJldHVybiBQcm9taXNlLnJlamVjdChlcnIpXG4gICAgICB9XG4gICAgKVxuICB9XG5cbiAgYXN5bmMgZ2V0PFQgPSBvYmplY3Q+KHBhdGg6IHN0cmluZywgcGFyYW1zOiBvYmplY3Q9e30pOiBQcm9taXNlPEF4aW9zUmVzcG9uc2U8VD4+IHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgcmVzdWx0ID0gYXdhaXQgdGhpcy5heGlvc0luc3RhbmNlLmdldChwYXRoLCB7XG4gICAgICAgIHBhcmFtc1xuICAgICAgfSlcbiAgICAgIHJldHVybiB0aGlzLnN1Y2Nlc3NQcm9taXNlKHJlc3VsdC5kYXRhKVxuICAgIH0gY2F0Y2goZXJyOiBhbnkpIHtcbiAgICAgIHJldHVybiB0aGlzLmZhaWx1cmVQcm9taXNlPFQ+KGVycilcbiAgICB9XG4gIH1cblxuICBhc3luYyBwb3N0PFQgPSBvYmplY3Q+KHBhdGg6IHN0cmluZywgcGFyYW1zOiBvYmplY3Q9e30pOiBQcm9taXNlPEF4aW9zUmVzcG9uc2U8VD4+IHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgcmVzdWx0ID0gYXdhaXQgdGhpcy5heGlvc0luc3RhbmNlLnBvc3QocGF0aCwge1xuICAgICAgICBwYXJhbXNcbiAgICAgIH0pXG4gICAgICByZXR1cm4gdGhpcy5zdWNjZXNzUHJvbWlzZShyZXN1bHQuZGF0YSlcbiAgICB9IGNhdGNoKGVycjogYW55KSB7XG4gICAgICByZXR1cm4gdGhpcy5mYWlsdXJlUHJvbWlzZTxUPihlcnIpXG4gICAgfVxuICB9XG5cbiAgYXN5bmMgcHV0PFQgPSBvYmplY3Q+KHBhdGg6IHN0cmluZywgcGFyYW1zOiBvYmplY3Q9e30pOiBQcm9taXNlPEF4aW9zUmVzcG9uc2U8VD4+IHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgcmVzdWx0ID0gYXdhaXQgdGhpcy5heGlvc0luc3RhbmNlLnB1dChwYXRoLCB7XG4gICAgICAgIHBhcmFtc1xuICAgICAgfSlcbiAgICAgIHJldHVybiB0aGlzLnN1Y2Nlc3NQcm9taXNlKHJlc3VsdC5kYXRhKVxuICAgIH0gY2F0Y2goZXJyOiBhbnkpIHtcbiAgICAgIHJldHVybiB0aGlzLmZhaWx1cmVQcm9taXNlPFQ+KGVycilcbiAgICB9XG4gIH1cblxuICBhc3luYyBwYXRjaDxUID0gb2JqZWN0PihwYXRoOiBzdHJpbmcsIHBhcmFtczogb2JqZWN0PXt9KTogUHJvbWlzZTxBeGlvc1Jlc3BvbnNlPFQ+PiB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHJlc3VsdCA9IGF3YWl0IHRoaXMuYXhpb3NJbnN0YW5jZS5wYXRjaChwYXRoLCB7XG4gICAgICAgIHBhcmFtc1xuICAgICAgfSlcbiAgICAgIHJldHVybiB0aGlzLnN1Y2Nlc3NQcm9taXNlKHJlc3VsdC5kYXRhKVxuICAgIH0gY2F0Y2goZXJyOiBhbnkpIHtcbiAgICAgIHJldHVybiB0aGlzLmZhaWx1cmVQcm9taXNlPFQ+KGVycilcbiAgICB9XG4gIH1cblxuICBhc3luYyBkZWxldGU8VCA9IG9iamVjdD4ocGF0aDogc3RyaW5nKTogUHJvbWlzZTxBeGlvc1Jlc3BvbnNlPFQ+PiB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHJlc3VsdCA9IGF3YWl0IHRoaXMuYXhpb3NJbnN0YW5jZS5kZWxldGUocGF0aClcbiAgICAgIHJldHVybiB0aGlzLnN1Y2Nlc3NQcm9taXNlKHJlc3VsdC5kYXRhKVxuICAgIH0gY2F0Y2goZXJyOiBhbnkpIHtcbiAgICAgIHJldHVybiB0aGlzLmZhaWx1cmVQcm9taXNlPFQ+KGVycilcbiAgICB9XG4gIH1cblxuICBwcml2YXRlIHN1Y2Nlc3NQcm9taXNlPFQ+KGRhdGE6IFQpOiBQcm9taXNlPEF4aW9zUmVzcG9uc2U8VD4+IHtcbiAgICByZXR1cm4gUHJvbWlzZS5yZXNvbHZlPEF4aW9zUmVzcG9uc2U8VD4+ICh7XG4gICAgICBkYXRhLFxuICAgICAgaXNTdWNjZXNzOiB0cnVlLFxuICAgIH0pXG4gIH1cblxuICBwcml2YXRlIGZhaWx1cmVQcm9taXNlPFQ+KGVycm9yOiBBeGlvc0Vycm9yKTogUHJvbWlzZTxBeGlvc1Jlc3BvbnNlPFQ+PiB7XG4gICAgcmV0dXJuIFByb21pc2UucmVqZWN0PEF4aW9zUmVzcG9uc2U8VD4+ICh7XG4gICAgICBlcnJvcixcbiAgICAgIGlzU3VjY2VzczogZmFsc2UsXG4gICAgfSlcbiAgfVxufVxuXG5cblxudHlwZSBBeGlvc1Jlc3BvbnNlPFQ+ID0ge1xuICBkYXRhPzogVFxuICBlcnJvcj86IEF4aW9zRXJyb3JcbiAgaXNTdWNjZXNzOiBib29sZWFuXG59Il0sIm5hbWVzIjpbImF4aW9zIiwiQXBpQ2xpZW50IiwiZ2V0IiwicGF0aCIsInBhcmFtcyIsInJlc3VsdCIsImF4aW9zSW5zdGFuY2UiLCJzdWNjZXNzUHJvbWlzZSIsImRhdGEiLCJlcnIiLCJmYWlsdXJlUHJvbWlzZSIsInBvc3QiLCJwdXQiLCJwYXRjaCIsImRlbGV0ZSIsIlByb21pc2UiLCJyZXNvbHZlIiwiaXNTdWNjZXNzIiwiZXJyb3IiLCJyZWplY3QiLCJjb25zdHJ1Y3RvciIsImJhc2VVUkwiLCJjcmVhdGUiLCJpbnRlcmNlcHRvcnMiLCJyZXF1ZXN0IiwidXNlIiwiY29uZmlnIiwibG9jYWxTdG9yYWdlIiwiZ2V0SXRlbSIsImhlYWRlcnMiLCJhdXRob3JpemF0aW9uIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./apiClient/index.ts\n"));

/***/ }),

/***/ "./apiClient/item.ts":
/*!***************************!*\
  !*** ./apiClient/item.ts ***!
  \***************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"getItemDisplayOn\": function() { return /* binding */ getItemDisplayOn; },\n/* harmony export */   \"getItemsDisplayOnFromCategoryId\": function() { return /* binding */ getItemsDisplayOnFromCategoryId; },\n/* harmony export */   \"postItem\": function() { return /* binding */ postItem; }\n/* harmony export */ });\n/* harmony import */ var ___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! . */ \"./apiClient/index.ts\");\n\nconst apiBaseUrl = \"http://localhost:8000\";\nconst apiClient = new ___WEBPACK_IMPORTED_MODULE_0__.ApiClient(apiBaseUrl);\nconst getItemDisplayOn = (itemId)=>{\n    return apiClient.get(\"/api/items/get/\" + itemId);\n};\nconst getItemsDisplayOnFromCategoryId = (categoryId)=>{\n    return apiClient.get(\"/api/categories/\" + categoryId + \"/items/get\");\n};\nconst postItem = (payload)=>{\n    return apiClient.post(\"/api/admin/items/post\", payload);\n};\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9hcGlDbGllbnQvaXRlbS50cy5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7O0FBQTJCO0FBRTNCLE1BQU1DLGFBQWFDLHVCQUFnQztBQUNuRCxNQUFNRyxZQUFZLElBQUlMLHdDQUFTQSxDQUFDQztBQUV6QixNQUFNSyxtQkFBbUIsQ0FBQ0MsU0FBbUI7SUFDbEQsT0FBT0YsVUFBVUcsR0FBRyxDQUNsQixvQkFBb0JEO0FBRXhCLEVBQUM7QUFFTSxNQUFNRSxrQ0FBa0MsQ0FBQ0MsYUFBdUI7SUFDckUsT0FBT0wsVUFBVUcsR0FBRyxDQUNsQixxQkFBcUJFLGFBQWE7QUFFdEMsRUFBQztBQUVNLE1BQU1DLFdBQVcsQ0FBQ0MsVUFBaUI7SUFDdEMsT0FBT1AsVUFBVVEsSUFBSSxDQUNuQix5QkFDQUQ7QUFFTixFQUFDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vX05fRS8uL2FwaUNsaWVudC9pdGVtLnRzPzBiMGIiXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHtBcGlDbGllbnR9IGZyb20gJy4nXG5cbmNvbnN0IGFwaUJhc2VVcmwgPSBwcm9jZXNzLmVudi5ORVhUX1BVQkxJQ19CQVNFX1VSTFxuY29uc3QgYXBpQ2xpZW50ID0gbmV3IEFwaUNsaWVudChhcGlCYXNlVXJsKVxuXG5leHBvcnQgY29uc3QgZ2V0SXRlbURpc3BsYXlPbiA9IChpdGVtSWQ6IHN0cmluZykgPT4ge1xuICByZXR1cm4gYXBpQ2xpZW50LmdldChcbiAgICAnL2FwaS9pdGVtcy9nZXQvJyArIGl0ZW1JZCxcbiAgKVxufVxuXG5leHBvcnQgY29uc3QgZ2V0SXRlbXNEaXNwbGF5T25Gcm9tQ2F0ZWdvcnlJZCA9IChjYXRlZ29yeUlkOiBzdHJpbmcpID0+IHtcbiAgcmV0dXJuIGFwaUNsaWVudC5nZXQoXG4gICAgJy9hcGkvY2F0ZWdvcmllcy8nICsgY2F0ZWdvcnlJZCArICcvaXRlbXMvZ2V0JyxcbiAgKVxufVxuXG5leHBvcnQgY29uc3QgcG9zdEl0ZW0gPSAocGF5bG9hZDogYW55KSA9PiB7XG4gICAgcmV0dXJuIGFwaUNsaWVudC5wb3N0KFxuICAgICAgJy9hcGkvYWRtaW4vaXRlbXMvcG9zdCcsXG4gICAgICBwYXlsb2FkXG4gICAgKVxufSJdLCJuYW1lcyI6WyJBcGlDbGllbnQiLCJhcGlCYXNlVXJsIiwicHJvY2VzcyIsImVudiIsIk5FWFRfUFVCTElDX0JBU0VfVVJMIiwiYXBpQ2xpZW50IiwiZ2V0SXRlbURpc3BsYXlPbiIsIml0ZW1JZCIsImdldCIsImdldEl0ZW1zRGlzcGxheU9uRnJvbUNhdGVnb3J5SWQiLCJjYXRlZ29yeUlkIiwicG9zdEl0ZW0iLCJwYXlsb2FkIiwicG9zdCJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./apiClient/item.ts\n"));

/***/ }),

/***/ "./pages/items/[id].tsx":
/*!******************************!*\
  !*** ./pages/items/[id].tsx ***!
  \******************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"./node_modules/react/jsx-dev-runtime.js\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"./node_modules/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! next/router */ \"./node_modules/next/router.js\");\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(next_router__WEBPACK_IMPORTED_MODULE_2__);\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! axios */ \"./node_modules/axios/index.js\");\n/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! next/link */ \"./node_modules/next/link.js\");\n/* harmony import */ var next_link__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(next_link__WEBPACK_IMPORTED_MODULE_3__);\n/* harmony import */ var _apiClient_item__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../../apiClient/item */ \"./apiClient/item.ts\");\n\nvar _s = $RefreshSig$();\n\n\n\n\n\nconst Category = ()=>{\n    _s();\n    axios__WEBPACK_IMPORTED_MODULE_5__[\"default\"].defaults.headers.common = {\n        \"Content-Type\": \"application/json\"\n    };\n    const [items, setItems] = (0,react__WEBPACK_IMPORTED_MODULE_1__.useState)([]);\n    const router = (0,next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter)();\n    const pathId = router.query.id;\n    const getItems = (0,react__WEBPACK_IMPORTED_MODULE_1__.useCallback)(async (pathId)=>{\n        await (0,_apiClient_item__WEBPACK_IMPORTED_MODULE_4__.getItemsDisplayOnFromCategoryId)(pathId).then((res)=>{\n            setItems(res.data);\n        });\n    }, []);\n    (0,react__WEBPACK_IMPORTED_MODULE_1__.useEffect)(()=>{\n        getItems(pathId);\n    }, [\n        pathId\n    ]);\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.Fragment, {\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h1\", {\n                className: \"text-3xl font-bold underline\",\n                children: \"Hello world!\"\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 34,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"ul\", {\n                children: items.map((v, i)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"li\", {\n                        children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)((next_link__WEBPACK_IMPORTED_MODULE_3___default()), {\n                            href: {\n                                pathname: \"/items/[item_id]\",\n                                query: {\n                                    item_id: v.id\n                                }\n                            },\n                            children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"p\", {\n                                children: v.name\n                            }, void 0, false, {\n                                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                                lineNumber: 44,\n                                columnNumber: 15\n                            }, undefined)\n                        }, void 0, false, {\n                            fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                            lineNumber: 40,\n                            columnNumber: 13\n                        }, undefined)\n                    }, i, false, {\n                        fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                        lineNumber: 39,\n                        columnNumber: 11\n                    }, undefined))\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 37,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true);\n};\n_s(Category, \"F6vrTcaRMsgW/V4ljnr9mbaciec=\", false, function() {\n    return [\n        next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter\n    ];\n});\n_c = Category;\n/* harmony default export */ __webpack_exports__[\"default\"] = (Category);\nvar _c;\n$RefreshReg$(_c, \"Category\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports on update so we can compare the boundary\n                // signatures.\n                module.hot.dispose(function (data) {\n                    data.prevExports = currentExports;\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevExports !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevExports !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9pdGVtcy9baWRdLnRzeC5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7OztBQUFBOztBQUF3RTtBQUNqQztBQUNiO0FBQ0c7QUFHeUM7QUFJdEUsTUFBTVEsV0FBVyxJQUFNOztJQUNyQkgscUVBQTZCLEdBQUc7UUFDOUIsZ0JBQWdCO0lBR2xCO0lBQ0EsTUFBTSxDQUFDTyxPQUFPQyxTQUFTLEdBQUdaLCtDQUFRQSxDQUFTLEVBQUU7SUFDN0MsTUFBTWEsU0FBU1Ysc0RBQVNBO0lBQ3hCLE1BQU1XLFNBQVNELE9BQU9FLEtBQUssQ0FBQ0MsRUFBRTtJQUM5QixNQUFNQyxXQUFXZixrREFBV0EsQ0FBQyxPQUFPWSxTQUFnQjtRQUNsRCxNQUFNUixnRkFBK0JBLENBQUNRLFFBQ25DSSxJQUFJLENBQUMsQ0FBQ0MsTUFBYTtZQUNsQlAsU0FBU08sSUFBSUMsSUFBSTtRQUNuQjtJQUNKLEdBQUcsRUFBRTtJQUVMbkIsZ0RBQVNBLENBQUMsSUFBTztRQUNmZ0IsU0FBU0g7SUFFWCxHQUFHO1FBQUNBO0tBQU87SUFFWCxxQkFDRTs7MEJBQ0UsOERBQUNPO2dCQUFHQyxXQUFVOzBCQUErQjs7Ozs7OzBCQUc3Qyw4REFBQ0M7MEJBQ0VaLE1BQU1hLEdBQUcsQ0FBQyxDQUFDQyxHQUFHQyxrQkFDYiw4REFBQ0M7a0NBQ0MsNEVBQUN0QixrREFBSUE7NEJBQUN1QixNQUFNO2dDQUNWQyxVQUFVO2dDQUNWZCxPQUFPO29DQUFDZSxTQUFTTCxFQUFFVCxFQUFFO2dDQUFBOzRCQUN2QjtzQ0FDRSw0RUFBQ2U7MENBQUdOLEVBQUVPLElBQUk7Ozs7Ozs7Ozs7O3VCQUxMTjs7Ozs7Ozs7Ozs7O0FBWW5CO0dBeENNbkI7O1FBT1dKLGtEQUFTQTs7O0tBUHBCSTtBQTBDTiwrREFBZUEsUUFBUUEsRUFBQSIsInNvdXJjZXMiOlsid2VicGFjazovL19OX0UvLi9wYWdlcy9pdGVtcy9baWRdLnRzeD84MzcwIl0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCBSZWFjdCwgeyB1c2VTdGF0ZSwgdXNlRWZmZWN0LCB1c2VSZWYsIHVzZUNhbGxiYWNrIH0gZnJvbSBcInJlYWN0XCI7XG5pbXBvcnQgeyB1c2VSb3V0ZXIgfSBmcm9tICduZXh0L3JvdXRlcidcbmltcG9ydCBheGlvcyBmcm9tICdheGlvcyc7XG5pbXBvcnQgTGluayBmcm9tICduZXh0L2xpbmsnO1xuaW1wb3J0IEltYWdlIGZyb20gJ25leHQvaW1hZ2UnXG5pbXBvcnQgSXRlbSBmcm9tICcuLi8uLi9tb2RlbHMvaXRlbSdcbmltcG9ydCB7IGdldEl0ZW1zRGlzcGxheU9uRnJvbUNhdGVnb3J5SWQgfSBmcm9tIFwiLi4vLi4vYXBpQ2xpZW50L2l0ZW1cIlxuXG5cblxuY29uc3QgQ2F0ZWdvcnkgPSAoKSA9PiB7XG4gIGF4aW9zLmRlZmF1bHRzLmhlYWRlcnMuY29tbW9uID0ge1xuICAgICdDb250ZW50LVR5cGUnOiAnYXBwbGljYXRpb24vanNvbicsXG4gICAgLy8gJ1gtUmVxdWVzdGVkLVdpdGgnOiAnWE1MSHR0cFJlcXVlc3QnLFxuICAgIC8vIFwiWC1DU1JGLVRva2VuXCI6IGNzcmZUb2tlblxuICB9XG4gIGNvbnN0IFtpdGVtcywgc2V0SXRlbXNdID0gdXNlU3RhdGU8SXRlbVtdPihbXSlcbiAgY29uc3Qgcm91dGVyID0gdXNlUm91dGVyKClcbiAgY29uc3QgcGF0aElkID0gcm91dGVyLnF1ZXJ5LmlkXG4gIGNvbnN0IGdldEl0ZW1zID0gdXNlQ2FsbGJhY2soYXN5bmMgKHBhdGhJZDogYW55KSA9PiB7XG4gICAgYXdhaXQgZ2V0SXRlbXNEaXNwbGF5T25Gcm9tQ2F0ZWdvcnlJZChwYXRoSWQpXG4gICAgICAudGhlbigocmVzOiBhbnkpID0+IHtcbiAgICAgICAgc2V0SXRlbXMocmVzLmRhdGEpXG4gICAgICB9KVxuICB9LCBbXSlcbiAgXG4gIHVzZUVmZmVjdCgoKSAgPT4ge1xuICAgIGdldEl0ZW1zKHBhdGhJZCk7XG5cbiAgfSwgW3BhdGhJZF0pXG5cbiAgcmV0dXJuIChcbiAgICA8PlxuICAgICAgPGgxIGNsYXNzTmFtZT1cInRleHQtM3hsIGZvbnQtYm9sZCB1bmRlcmxpbmVcIj5cbiAgICAgICAgSGVsbG8gd29ybGQhXG4gICAgICA8L2gxPlxuICAgICAgPHVsPlxuICAgICAgICB7aXRlbXMubWFwKCh2LCBpKSA9PlxuICAgICAgICAgIDxsaSBrZXk9e2l9PlxuICAgICAgICAgICAgPExpbmsgaHJlZj17e1xuICAgICAgICAgICAgICBwYXRobmFtZTogXCIvaXRlbXMvW2l0ZW1faWRdXCIsXG4gICAgICAgICAgICAgIHF1ZXJ5OiB7aXRlbV9pZDogdi5pZH1cbiAgICAgICAgICAgIH19PlxuICAgICAgICAgICAgICA8cD57di5uYW1lfTwvcD5cbiAgICAgICAgICAgIDwvTGluaz5cbiAgICAgICAgICA8L2xpPlxuICAgICAgICApfVxuICAgICAgPC91bD5cbiAgICA8Lz5cbiAgKVxufVxuXG5leHBvcnQgZGVmYXVsdCBDYXRlZ29yeSJdLCJuYW1lcyI6WyJSZWFjdCIsInVzZVN0YXRlIiwidXNlRWZmZWN0IiwidXNlQ2FsbGJhY2siLCJ1c2VSb3V0ZXIiLCJheGlvcyIsIkxpbmsiLCJnZXRJdGVtc0Rpc3BsYXlPbkZyb21DYXRlZ29yeUlkIiwiQ2F0ZWdvcnkiLCJkZWZhdWx0cyIsImhlYWRlcnMiLCJjb21tb24iLCJpdGVtcyIsInNldEl0ZW1zIiwicm91dGVyIiwicGF0aElkIiwicXVlcnkiLCJpZCIsImdldEl0ZW1zIiwidGhlbiIsInJlcyIsImRhdGEiLCJoMSIsImNsYXNzTmFtZSIsInVsIiwibWFwIiwidiIsImkiLCJsaSIsImhyZWYiLCJwYXRobmFtZSIsIml0ZW1faWQiLCJwIiwibmFtZSJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./pages/items/[id].tsx\n"));

/***/ })

});