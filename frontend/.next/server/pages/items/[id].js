"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
(() => {
var exports = {};
exports.id = "pages/items/[id]";
exports.ids = ["pages/items/[id]"];
exports.modules = {

/***/ "./apiClient/index.ts":
/*!****************************!*\
  !*** ./apiClient/index.ts ***!
  \****************************/
/***/ ((module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.a(module, async (__webpack_handle_async_dependencies__, __webpack_async_result__) => { try {\n__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"ApiClient\": () => (/* binding */ ApiClient)\n/* harmony export */ });\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! axios */ \"axios\");\nvar __webpack_async_dependencies__ = __webpack_handle_async_dependencies__([axios__WEBPACK_IMPORTED_MODULE_0__]);\naxios__WEBPACK_IMPORTED_MODULE_0__ = (__webpack_async_dependencies__.then ? (await __webpack_async_dependencies__)() : __webpack_async_dependencies__)[0];\n\nclass ApiClient {\n    constructor(baseURL = \"\"){\n        this.axiosInstance = axios__WEBPACK_IMPORTED_MODULE_0__[\"default\"].create({\n            baseURL\n        });\n        this.axiosInstance.interceptors.request.use(async (config)=>{\n            if (localStorage.getItem(\"token\")) {\n                config.headers.authorization = \"Bearer \" + localStorage.getItem(\"token\");\n            }\n            return config;\n        }, (err)=>{\n            return Promise.reject(err);\n        });\n    }\n    async get(path, params = {}) {\n        try {\n            const result = await this.axiosInstance.get(path, {\n                params\n            });\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    async post(path, params = {}) {\n        try {\n            const result = await this.axiosInstance.post(path, {\n                params\n            });\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    async put(path, params = {}) {\n        try {\n            const result = await this.axiosInstance.put(path, {\n                params\n            });\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    async patch(path, params = {}) {\n        try {\n            const result = await this.axiosInstance.patch(path, {\n                params\n            });\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    async delete(path) {\n        try {\n            const result = await this.axiosInstance.delete(path);\n            return this.successPromise(result.data);\n        } catch (err) {\n            return this.failurePromise(err);\n        }\n    }\n    successPromise(data) {\n        return Promise.resolve({\n            data,\n            isSuccess: true\n        });\n    }\n    failurePromise(error) {\n        return Promise.reject({\n            error,\n            isSuccess: false\n        });\n    }\n}\n\n__webpack_async_result__();\n} catch(e) { __webpack_async_result__(e); } });//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9hcGlDbGllbnQvaW5kZXgudHMuanMiLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7QUFBMEU7QUFFbkUsTUFBTUM7SUFFWEMsWUFBWUMsVUFBVSxFQUFFLENBQUU7UUFDeEIsSUFBSSxDQUFDQyxhQUFhLEdBQUdKLG9EQUFZLENBQUM7WUFDaENHO1FBQ0Y7UUFFQSxJQUFJLENBQUNDLGFBQWEsQ0FBQ0UsWUFBWSxDQUFDQyxPQUFPLENBQUNDLEdBQUcsQ0FDekMsT0FBT0MsU0FBcUM7WUFDMUMsSUFBSUMsYUFBYUMsT0FBTyxDQUFDLFVBQVU7Z0JBQ2pDRixPQUFPRyxPQUFPLENBQUNDLGFBQWEsR0FBRyxZQUFZSCxhQUFhQyxPQUFPLENBQUM7WUFDbEUsQ0FBQztZQUNELE9BQU9GO1FBQ1QsR0FBRyxDQUFDSyxNQUFvQjtZQUN0QixPQUFPQyxRQUFRQyxNQUFNLENBQUNGO1FBQ3hCO0lBRUo7SUFFQSxNQUFNRyxJQUFnQkMsSUFBWSxFQUFFQyxTQUFlLENBQUMsQ0FBQyxFQUE2QjtRQUNoRixJQUFJO1lBQ0YsTUFBTUMsU0FBUyxNQUFNLElBQUksQ0FBQ2hCLGFBQWEsQ0FBQ2EsR0FBRyxDQUFDQyxNQUFNO2dCQUNoREM7WUFDRjtZQUNBLE9BQU8sSUFBSSxDQUFDRSxjQUFjLENBQUNELE9BQU9FLElBQUk7UUFDeEMsRUFBRSxPQUFNUixLQUFVO1lBQ2hCLE9BQU8sSUFBSSxDQUFDUyxjQUFjLENBQUlUO1FBQ2hDO0lBQ0Y7SUFFQSxNQUFNVSxLQUFpQk4sSUFBWSxFQUFFQyxTQUFlLENBQUMsQ0FBQyxFQUE2QjtRQUNqRixJQUFJO1lBQ0YsTUFBTUMsU0FBUyxNQUFNLElBQUksQ0FBQ2hCLGFBQWEsQ0FBQ29CLElBQUksQ0FBQ04sTUFBTTtnQkFDakRDO1lBQ0Y7WUFDQSxPQUFPLElBQUksQ0FBQ0UsY0FBYyxDQUFDRCxPQUFPRSxJQUFJO1FBQ3hDLEVBQUUsT0FBTVIsS0FBVTtZQUNoQixPQUFPLElBQUksQ0FBQ1MsY0FBYyxDQUFJVDtRQUNoQztJQUNGO0lBRUEsTUFBTVcsSUFBZ0JQLElBQVksRUFBRUMsU0FBZSxDQUFDLENBQUMsRUFBNkI7UUFDaEYsSUFBSTtZQUNGLE1BQU1DLFNBQVMsTUFBTSxJQUFJLENBQUNoQixhQUFhLENBQUNxQixHQUFHLENBQUNQLE1BQU07Z0JBQ2hEQztZQUNGO1lBQ0EsT0FBTyxJQUFJLENBQUNFLGNBQWMsQ0FBQ0QsT0FBT0UsSUFBSTtRQUN4QyxFQUFFLE9BQU1SLEtBQVU7WUFDaEIsT0FBTyxJQUFJLENBQUNTLGNBQWMsQ0FBSVQ7UUFDaEM7SUFDRjtJQUVBLE1BQU1ZLE1BQWtCUixJQUFZLEVBQUVDLFNBQWUsQ0FBQyxDQUFDLEVBQTZCO1FBQ2xGLElBQUk7WUFDRixNQUFNQyxTQUFTLE1BQU0sSUFBSSxDQUFDaEIsYUFBYSxDQUFDc0IsS0FBSyxDQUFDUixNQUFNO2dCQUNsREM7WUFDRjtZQUNBLE9BQU8sSUFBSSxDQUFDRSxjQUFjLENBQUNELE9BQU9FLElBQUk7UUFDeEMsRUFBRSxPQUFNUixLQUFVO1lBQ2hCLE9BQU8sSUFBSSxDQUFDUyxjQUFjLENBQUlUO1FBQ2hDO0lBQ0Y7SUFFQSxNQUFNYSxPQUFtQlQsSUFBWSxFQUE2QjtRQUNoRSxJQUFJO1lBQ0YsTUFBTUUsU0FBUyxNQUFNLElBQUksQ0FBQ2hCLGFBQWEsQ0FBQ3VCLE1BQU0sQ0FBQ1Q7WUFDL0MsT0FBTyxJQUFJLENBQUNHLGNBQWMsQ0FBQ0QsT0FBT0UsSUFBSTtRQUN4QyxFQUFFLE9BQU1SLEtBQVU7WUFDaEIsT0FBTyxJQUFJLENBQUNTLGNBQWMsQ0FBSVQ7UUFDaEM7SUFDRjtJQUVRTyxlQUFrQkMsSUFBTyxFQUE2QjtRQUM1RCxPQUFPUCxRQUFRYSxPQUFPLENBQW9CO1lBQ3hDTjtZQUNBTyxXQUFXLElBQUk7UUFDakI7SUFDRjtJQUVRTixlQUFrQk8sS0FBaUIsRUFBNkI7UUFDdEUsT0FBT2YsUUFBUUMsTUFBTSxDQUFvQjtZQUN2Q2M7WUFDQUQsV0FBVyxLQUFLO1FBQ2xCO0lBQ0Y7QUFDRixDQUFDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vbXktYXBwLy4vYXBpQ2xpZW50L2luZGV4LnRzPzQ5MjkiXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IGF4aW9zLCB7QXhpb3NFcnJvciwgQXhpb3NJbnN0YW5jZSwgQXhpb3NSZXF1ZXN0Q29uZmlnfSBmcm9tICdheGlvcydcblxuZXhwb3J0IGNsYXNzIEFwaUNsaWVudCB7XG4gIGF4aW9zSW5zdGFuY2U6IEF4aW9zSW5zdGFuY2VcbiAgY29uc3RydWN0b3IoYmFzZVVSTCA9ICcnKSB7XG4gICAgdGhpcy5heGlvc0luc3RhbmNlID0gYXhpb3MuY3JlYXRlKHtcbiAgICAgIGJhc2VVUkxcbiAgICB9KVxuXG4gICAgdGhpcy5heGlvc0luc3RhbmNlLmludGVyY2VwdG9ycy5yZXF1ZXN0LnVzZShcbiAgICAgIGFzeW5jIChjb25maWc6IEF4aW9zUmVxdWVzdENvbmZpZyB8IGFueSkgPT4ge1xuICAgICAgICBpZiAobG9jYWxTdG9yYWdlLmdldEl0ZW0oJ3Rva2VuJykpIHtcbiAgICAgICAgICBjb25maWcuaGVhZGVycy5hdXRob3JpemF0aW9uID0gJ0JlYXJlciAnICsgbG9jYWxTdG9yYWdlLmdldEl0ZW0oJ3Rva2VuJylcbiAgICAgICAgfVxuICAgICAgICByZXR1cm4gY29uZmlnXG4gICAgICB9LCAoZXJyOiBBeGlvc0Vycm9yKSA9PiB7XG4gICAgICAgIHJldHVybiBQcm9taXNlLnJlamVjdChlcnIpXG4gICAgICB9XG4gICAgKVxuICB9XG5cbiAgYXN5bmMgZ2V0PFQgPSBvYmplY3Q+KHBhdGg6IHN0cmluZywgcGFyYW1zOiBvYmplY3Q9e30pOiBQcm9taXNlPEF4aW9zUmVzcG9uc2U8VD4+IHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgcmVzdWx0ID0gYXdhaXQgdGhpcy5heGlvc0luc3RhbmNlLmdldChwYXRoLCB7XG4gICAgICAgIHBhcmFtc1xuICAgICAgfSlcbiAgICAgIHJldHVybiB0aGlzLnN1Y2Nlc3NQcm9taXNlKHJlc3VsdC5kYXRhKVxuICAgIH0gY2F0Y2goZXJyOiBhbnkpIHtcbiAgICAgIHJldHVybiB0aGlzLmZhaWx1cmVQcm9taXNlPFQ+KGVycilcbiAgICB9XG4gIH1cblxuICBhc3luYyBwb3N0PFQgPSBvYmplY3Q+KHBhdGg6IHN0cmluZywgcGFyYW1zOiBvYmplY3Q9e30pOiBQcm9taXNlPEF4aW9zUmVzcG9uc2U8VD4+IHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgcmVzdWx0ID0gYXdhaXQgdGhpcy5heGlvc0luc3RhbmNlLnBvc3QocGF0aCwge1xuICAgICAgICBwYXJhbXNcbiAgICAgIH0pXG4gICAgICByZXR1cm4gdGhpcy5zdWNjZXNzUHJvbWlzZShyZXN1bHQuZGF0YSlcbiAgICB9IGNhdGNoKGVycjogYW55KSB7XG4gICAgICByZXR1cm4gdGhpcy5mYWlsdXJlUHJvbWlzZTxUPihlcnIpXG4gICAgfVxuICB9XG5cbiAgYXN5bmMgcHV0PFQgPSBvYmplY3Q+KHBhdGg6IHN0cmluZywgcGFyYW1zOiBvYmplY3Q9e30pOiBQcm9taXNlPEF4aW9zUmVzcG9uc2U8VD4+IHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgcmVzdWx0ID0gYXdhaXQgdGhpcy5heGlvc0luc3RhbmNlLnB1dChwYXRoLCB7XG4gICAgICAgIHBhcmFtc1xuICAgICAgfSlcbiAgICAgIHJldHVybiB0aGlzLnN1Y2Nlc3NQcm9taXNlKHJlc3VsdC5kYXRhKVxuICAgIH0gY2F0Y2goZXJyOiBhbnkpIHtcbiAgICAgIHJldHVybiB0aGlzLmZhaWx1cmVQcm9taXNlPFQ+KGVycilcbiAgICB9XG4gIH1cblxuICBhc3luYyBwYXRjaDxUID0gb2JqZWN0PihwYXRoOiBzdHJpbmcsIHBhcmFtczogb2JqZWN0PXt9KTogUHJvbWlzZTxBeGlvc1Jlc3BvbnNlPFQ+PiB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHJlc3VsdCA9IGF3YWl0IHRoaXMuYXhpb3NJbnN0YW5jZS5wYXRjaChwYXRoLCB7XG4gICAgICAgIHBhcmFtc1xuICAgICAgfSlcbiAgICAgIHJldHVybiB0aGlzLnN1Y2Nlc3NQcm9taXNlKHJlc3VsdC5kYXRhKVxuICAgIH0gY2F0Y2goZXJyOiBhbnkpIHtcbiAgICAgIHJldHVybiB0aGlzLmZhaWx1cmVQcm9taXNlPFQ+KGVycilcbiAgICB9XG4gIH1cblxuICBhc3luYyBkZWxldGU8VCA9IG9iamVjdD4ocGF0aDogc3RyaW5nKTogUHJvbWlzZTxBeGlvc1Jlc3BvbnNlPFQ+PiB7XG4gICAgdHJ5IHtcbiAgICAgIGNvbnN0IHJlc3VsdCA9IGF3YWl0IHRoaXMuYXhpb3NJbnN0YW5jZS5kZWxldGUocGF0aClcbiAgICAgIHJldHVybiB0aGlzLnN1Y2Nlc3NQcm9taXNlKHJlc3VsdC5kYXRhKVxuICAgIH0gY2F0Y2goZXJyOiBhbnkpIHtcbiAgICAgIHJldHVybiB0aGlzLmZhaWx1cmVQcm9taXNlPFQ+KGVycilcbiAgICB9XG4gIH1cblxuICBwcml2YXRlIHN1Y2Nlc3NQcm9taXNlPFQ+KGRhdGE6IFQpOiBQcm9taXNlPEF4aW9zUmVzcG9uc2U8VD4+IHtcbiAgICByZXR1cm4gUHJvbWlzZS5yZXNvbHZlPEF4aW9zUmVzcG9uc2U8VD4+ICh7XG4gICAgICBkYXRhLFxuICAgICAgaXNTdWNjZXNzOiB0cnVlLFxuICAgIH0pXG4gIH1cblxuICBwcml2YXRlIGZhaWx1cmVQcm9taXNlPFQ+KGVycm9yOiBBeGlvc0Vycm9yKTogUHJvbWlzZTxBeGlvc1Jlc3BvbnNlPFQ+PiB7XG4gICAgcmV0dXJuIFByb21pc2UucmVqZWN0PEF4aW9zUmVzcG9uc2U8VD4+ICh7XG4gICAgICBlcnJvcixcbiAgICAgIGlzU3VjY2VzczogZmFsc2UsXG4gICAgfSlcbiAgfVxufVxuXG5cblxudHlwZSBBeGlvc1Jlc3BvbnNlPFQ+ID0ge1xuICBkYXRhPzogVFxuICBlcnJvcj86IEF4aW9zRXJyb3JcbiAgaXNTdWNjZXNzOiBib29sZWFuXG59Il0sIm5hbWVzIjpbImF4aW9zIiwiQXBpQ2xpZW50IiwiY29uc3RydWN0b3IiLCJiYXNlVVJMIiwiYXhpb3NJbnN0YW5jZSIsImNyZWF0ZSIsImludGVyY2VwdG9ycyIsInJlcXVlc3QiLCJ1c2UiLCJjb25maWciLCJsb2NhbFN0b3JhZ2UiLCJnZXRJdGVtIiwiaGVhZGVycyIsImF1dGhvcml6YXRpb24iLCJlcnIiLCJQcm9taXNlIiwicmVqZWN0IiwiZ2V0IiwicGF0aCIsInBhcmFtcyIsInJlc3VsdCIsInN1Y2Nlc3NQcm9taXNlIiwiZGF0YSIsImZhaWx1cmVQcm9taXNlIiwicG9zdCIsInB1dCIsInBhdGNoIiwiZGVsZXRlIiwicmVzb2x2ZSIsImlzU3VjY2VzcyIsImVycm9yIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./apiClient/index.ts\n");

/***/ }),

/***/ "./apiClient/item.ts":
/*!***************************!*\
  !*** ./apiClient/item.ts ***!
  \***************************/
/***/ ((module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.a(module, async (__webpack_handle_async_dependencies__, __webpack_async_result__) => { try {\n__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"getItemDisplayOn\": () => (/* binding */ getItemDisplayOn),\n/* harmony export */   \"getItemsDisplayOn\": () => (/* binding */ getItemsDisplayOn),\n/* harmony export */   \"postItem\": () => (/* binding */ postItem)\n/* harmony export */ });\n/* harmony import */ var ___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! . */ \"./apiClient/index.ts\");\nvar __webpack_async_dependencies__ = __webpack_handle_async_dependencies__([___WEBPACK_IMPORTED_MODULE_0__]);\n___WEBPACK_IMPORTED_MODULE_0__ = (__webpack_async_dependencies__.then ? (await __webpack_async_dependencies__)() : __webpack_async_dependencies__)[0];\n\nconst apiBaseUrl = \"http://localhost:8000\";\nconst apiClient = new ___WEBPACK_IMPORTED_MODULE_0__.ApiClient(apiBaseUrl);\nconst getItemDisplayOn = (itemId)=>{\n    return apiClient.get(\"/api/items/get/\" + itemId);\n};\nconst getItemsDisplayOn = (categoryId)=>{\n    return apiClient.get(\"/api/categories/\" + categoryId + \"/items/get\");\n};\nconst postItem = (payload)=>{\n    return apiClient.post(\"/api/admin/items/post\", payload);\n};\n\n__webpack_async_result__();\n} catch(e) { __webpack_async_result__(e); } });//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9hcGlDbGllbnQvaXRlbS50cy5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7O0FBQTJCO0FBRTNCLE1BQU1DLGFBQWFDLHVCQUFnQztBQUNuRCxNQUFNRyxZQUFZLElBQUlMLHdDQUFTQSxDQUFDQztBQUV6QixNQUFNSyxtQkFBbUIsQ0FBQ0MsU0FBbUI7SUFDbEQsT0FBT0YsVUFBVUcsR0FBRyxDQUNsQixvQkFBb0JEO0FBRXhCLEVBQUM7QUFFTSxNQUFNRSxvQkFBb0IsQ0FBQ0MsYUFBdUI7SUFDdkQsT0FBT0wsVUFBVUcsR0FBRyxDQUNsQixxQkFBcUJFLGFBQWE7QUFFdEMsRUFBQztBQUVNLE1BQU1DLFdBQVcsQ0FBQ0MsVUFBaUI7SUFDdEMsT0FBT1AsVUFBVVEsSUFBSSxDQUNuQix5QkFDQUQ7QUFFTixFQUFDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vbXktYXBwLy4vYXBpQ2xpZW50L2l0ZW0udHM/MGIwYiJdLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQge0FwaUNsaWVudH0gZnJvbSAnLidcblxuY29uc3QgYXBpQmFzZVVybCA9IHByb2Nlc3MuZW52Lk5FWFRfUFVCTElDX0JBU0VfVVJMXG5jb25zdCBhcGlDbGllbnQgPSBuZXcgQXBpQ2xpZW50KGFwaUJhc2VVcmwpXG5cbmV4cG9ydCBjb25zdCBnZXRJdGVtRGlzcGxheU9uID0gKGl0ZW1JZDogc3RyaW5nKSA9PiB7XG4gIHJldHVybiBhcGlDbGllbnQuZ2V0KFxuICAgICcvYXBpL2l0ZW1zL2dldC8nICsgaXRlbUlkLFxuICApXG59XG5cbmV4cG9ydCBjb25zdCBnZXRJdGVtc0Rpc3BsYXlPbiA9IChjYXRlZ29yeUlkOiBzdHJpbmcpID0+IHtcbiAgcmV0dXJuIGFwaUNsaWVudC5nZXQoXG4gICAgJy9hcGkvY2F0ZWdvcmllcy8nICsgY2F0ZWdvcnlJZCArICcvaXRlbXMvZ2V0JyxcbiAgKVxufVxuXG5leHBvcnQgY29uc3QgcG9zdEl0ZW0gPSAocGF5bG9hZDogYW55KSA9PiB7XG4gICAgcmV0dXJuIGFwaUNsaWVudC5wb3N0KFxuICAgICAgJy9hcGkvYWRtaW4vaXRlbXMvcG9zdCcsXG4gICAgICBwYXlsb2FkXG4gICAgKVxufSJdLCJuYW1lcyI6WyJBcGlDbGllbnQiLCJhcGlCYXNlVXJsIiwicHJvY2VzcyIsImVudiIsIk5FWFRfUFVCTElDX0JBU0VfVVJMIiwiYXBpQ2xpZW50IiwiZ2V0SXRlbURpc3BsYXlPbiIsIml0ZW1JZCIsImdldCIsImdldEl0ZW1zRGlzcGxheU9uIiwiY2F0ZWdvcnlJZCIsInBvc3RJdGVtIiwicGF5bG9hZCIsInBvc3QiXSwic291cmNlUm9vdCI6IiJ9\n//# sourceURL=webpack-internal:///./apiClient/item.ts\n");

/***/ }),

/***/ "./pages/items/[id].tsx":
/*!******************************!*\
  !*** ./pages/items/[id].tsx ***!
  \******************************/
/***/ ((module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.a(module, async (__webpack_handle_async_dependencies__, __webpack_async_result__) => { try {\n__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"default\": () => (__WEBPACK_DEFAULT_EXPORT__)\n/* harmony export */ });\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"react/jsx-dev-runtime\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"react\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! next/router */ \"next/router\");\n/* harmony import */ var next_router__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(next_router__WEBPACK_IMPORTED_MODULE_2__);\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! axios */ \"axios\");\n/* harmony import */ var _apiClient_item__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../../apiClient/item */ \"./apiClient/item.ts\");\nvar __webpack_async_dependencies__ = __webpack_handle_async_dependencies__([axios__WEBPACK_IMPORTED_MODULE_3__, _apiClient_item__WEBPACK_IMPORTED_MODULE_4__]);\n([axios__WEBPACK_IMPORTED_MODULE_3__, _apiClient_item__WEBPACK_IMPORTED_MODULE_4__] = __webpack_async_dependencies__.then ? (await __webpack_async_dependencies__)() : __webpack_async_dependencies__);\n\n\n\n\n\nconst Category = ()=>{\n    axios__WEBPACK_IMPORTED_MODULE_3__[\"default\"].defaults.headers.common = {\n        \"Content-Type\": \"application/json\"\n    };\n    const [item, setItem] = (0,react__WEBPACK_IMPORTED_MODULE_1__.useState)();\n    const router = (0,next_router__WEBPACK_IMPORTED_MODULE_2__.useRouter)();\n    const pathId = router.query.id;\n    console.log(pathId);\n    const getItem = (0,react__WEBPACK_IMPORTED_MODULE_1__.useCallback)(async (pathId)=>{\n        if (pathId === undefined) return;\n        await (0,_apiClient_item__WEBPACK_IMPORTED_MODULE_4__.getItemDisplayOn)(pathId).then((res)=>{\n            console.log(res.data);\n            setItem(res.data);\n        });\n    }, []);\n    (0,react__WEBPACK_IMPORTED_MODULE_1__.useEffect)(()=>{\n        getItem(pathId);\n    }, [\n        pathId\n    ]);\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.Fragment, {\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h1\", {\n                className: \"text-3xl font-bold underline\",\n                children: \"Hello world!\"\n            }, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 37,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"ul\", {}, void 0, false, {\n                fileName: \"/Users/kawasakimasato/selfregi/frontend/pages/items/[id].tsx\",\n                lineNumber: 40,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true);\n};\n/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (Category);\n\n__webpack_async_result__();\n} catch(e) { __webpack_async_result__(e); } });//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9wYWdlcy9pdGVtcy9baWRdLnRzeC5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7Ozs7Ozs7QUFBQTtBQUF3RTtBQUNqQztBQUNiO0FBSTZCO0FBSXZELE1BQU1PLFdBQVcsSUFBTTtJQUNyQkYscUVBQTZCLEdBQUc7UUFDOUIsZ0JBQWdCO0lBR2xCO0lBRUEsTUFBTSxDQUFDTSxNQUFNQyxRQUFRLEdBQUdYLCtDQUFRQTtJQUNoQyxNQUFNWSxTQUFTVCxzREFBU0E7SUFDeEIsTUFBTVUsU0FBU0QsT0FBT0UsS0FBSyxDQUFDQyxFQUFFO0lBQzlCQyxRQUFRQyxHQUFHLENBQUNKO0lBQ1osTUFBTUssVUFBVWhCLGtEQUFXQSxDQUFDLE9BQU9XLFNBQWdCO1FBQ2pELElBQUlBLFdBQVdNLFdBQVc7UUFDMUIsTUFBTWQsaUVBQWdCQSxDQUFDUSxRQUNwQk8sSUFBSSxDQUFDLENBQUNDLE1BQWE7WUFDbEJMLFFBQVFDLEdBQUcsQ0FBQ0ksSUFBSUMsSUFBSTtZQUNwQlgsUUFBUVUsSUFBSUMsSUFBSTtRQUNsQjtJQUNKLEdBQUcsRUFBRTtJQUVMckIsZ0RBQVNBLENBQUMsSUFBTztRQUNmaUIsUUFBUUw7SUFDVixHQUFHO1FBQUNBO0tBQU87SUFFWCxxQkFDRTs7MEJBQ0UsOERBQUNVO2dCQUFHQyxXQUFVOzBCQUErQjs7Ozs7OzBCQUc3Qyw4REFBQ0M7Ozs7Ozs7QUFjUDtBQUVBLGlFQUFlbkIsUUFBUUEsRUFBQSIsInNvdXJjZXMiOlsid2VicGFjazovL215LWFwcC8uL3BhZ2VzL2l0ZW1zL1tpZF0udHN4PzgzNzAiXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IFJlYWN0LCB7IHVzZVN0YXRlLCB1c2VFZmZlY3QsIHVzZVJlZiwgdXNlQ2FsbGJhY2sgfSBmcm9tIFwicmVhY3RcIjtcbmltcG9ydCB7IHVzZVJvdXRlciB9IGZyb20gJ25leHQvcm91dGVyJ1xuaW1wb3J0IGF4aW9zIGZyb20gJ2F4aW9zJztcbmltcG9ydCBMaW5rIGZyb20gJ25leHQvbGluayc7XG5pbXBvcnQgSW1hZ2UgZnJvbSAnbmV4dC9pbWFnZSdcbmltcG9ydCBJdGVtIGZyb20gJy4uLy4uL21vZGVscy9pdGVtJ1xuaW1wb3J0IHsgZ2V0SXRlbURpc3BsYXlPbiB9IGZyb20gXCIuLi8uLi9hcGlDbGllbnQvaXRlbVwiXG5cblxuXG5jb25zdCBDYXRlZ29yeSA9ICgpID0+IHtcbiAgYXhpb3MuZGVmYXVsdHMuaGVhZGVycy5jb21tb24gPSB7XG4gICAgJ0NvbnRlbnQtVHlwZSc6ICdhcHBsaWNhdGlvbi9qc29uJyxcbiAgICAvLyAnWC1SZXF1ZXN0ZWQtV2l0aCc6ICdYTUxIdHRwUmVxdWVzdCcsXG4gICAgLy8gXCJYLUNTUkYtVG9rZW5cIjogY3NyZlRva2VuXG4gIH1cblxuICBjb25zdCBbaXRlbSwgc2V0SXRlbV0gPSB1c2VTdGF0ZTxJdGVtPigpXG4gIGNvbnN0IHJvdXRlciA9IHVzZVJvdXRlcigpXG4gIGNvbnN0IHBhdGhJZCA9IHJvdXRlci5xdWVyeS5pZFxuICBjb25zb2xlLmxvZyhwYXRoSWQpXG4gIGNvbnN0IGdldEl0ZW0gPSB1c2VDYWxsYmFjayhhc3luYyAocGF0aElkOiBhbnkpID0+IHtcbiAgICBpZiAocGF0aElkID09PSB1bmRlZmluZWQpIHJldHVybiBcbiAgICBhd2FpdCBnZXRJdGVtRGlzcGxheU9uKHBhdGhJZClcbiAgICAgIC50aGVuKChyZXM6IGFueSkgPT4ge1xuICAgICAgICBjb25zb2xlLmxvZyhyZXMuZGF0YSlcbiAgICAgICAgc2V0SXRlbShyZXMuZGF0YSlcbiAgICAgIH0pXG4gIH0sIFtdKVxuICBcbiAgdXNlRWZmZWN0KCgpICA9PiB7XG4gICAgZ2V0SXRlbShwYXRoSWQpO1xuICB9LCBbcGF0aElkXSlcblxuICByZXR1cm4gKFxuICAgIDw+XG4gICAgICA8aDEgY2xhc3NOYW1lPVwidGV4dC0zeGwgZm9udC1ib2xkIHVuZGVybGluZVwiPlxuICAgICAgICBIZWxsbyB3b3JsZCFcbiAgICAgIDwvaDE+XG4gICAgICA8dWw+XG4gICAgICAgIHsvKiB7aXRlbS5tYXAoKHYsIGkpID0+XG4gICAgICAgICAgPGxpIGtleT17aX0+XG4gICAgICAgICAgICA8TGluayBocmVmPXt7XG4gICAgICAgICAgICAgIHBhdGhuYW1lOiBcIi9pdGVtcy9baXRlbV9pZF1cIixcbiAgICAgICAgICAgICAgcXVlcnk6IHtpdGVtX2lkOiB2LmlkfVxuICAgICAgICAgICAgfX0+XG4gICAgICAgICAgICAgIDxwPnt2Lm5hbWV9PC9wPlxuICAgICAgICAgICAgPC9MaW5rPlxuICAgICAgICAgIDwvbGk+XG4gICAgICAgICl9ICovfVxuICAgICAgPC91bD5cbiAgICA8Lz5cbiAgKVxufVxuXG5leHBvcnQgZGVmYXVsdCBDYXRlZ29yeSJdLCJuYW1lcyI6WyJSZWFjdCIsInVzZVN0YXRlIiwidXNlRWZmZWN0IiwidXNlQ2FsbGJhY2siLCJ1c2VSb3V0ZXIiLCJheGlvcyIsImdldEl0ZW1EaXNwbGF5T24iLCJDYXRlZ29yeSIsImRlZmF1bHRzIiwiaGVhZGVycyIsImNvbW1vbiIsIml0ZW0iLCJzZXRJdGVtIiwicm91dGVyIiwicGF0aElkIiwicXVlcnkiLCJpZCIsImNvbnNvbGUiLCJsb2ciLCJnZXRJdGVtIiwidW5kZWZpbmVkIiwidGhlbiIsInJlcyIsImRhdGEiLCJoMSIsImNsYXNzTmFtZSIsInVsIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./pages/items/[id].tsx\n");

/***/ }),

/***/ "next/router":
/*!******************************!*\
  !*** external "next/router" ***!
  \******************************/
/***/ ((module) => {

module.exports = require("next/router");

/***/ }),

/***/ "react":
/*!************************!*\
  !*** external "react" ***!
  \************************/
/***/ ((module) => {

module.exports = require("react");

/***/ }),

/***/ "react/jsx-dev-runtime":
/*!****************************************!*\
  !*** external "react/jsx-dev-runtime" ***!
  \****************************************/
/***/ ((module) => {

module.exports = require("react/jsx-dev-runtime");

/***/ }),

/***/ "axios":
/*!************************!*\
  !*** external "axios" ***!
  \************************/
/***/ ((module) => {

module.exports = import("axios");;

/***/ })

};
;

// load runtime
var __webpack_require__ = require("../../webpack-runtime.js");
__webpack_require__.C(exports);
var __webpack_exec__ = (moduleId) => (__webpack_require__(__webpack_require__.s = moduleId))
var __webpack_exports__ = (__webpack_exec__("./pages/items/[id].tsx"));
module.exports = __webpack_exports__;

})();