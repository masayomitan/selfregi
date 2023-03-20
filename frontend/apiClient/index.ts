import axios, {AxiosError, AxiosInstance, AxiosRequestConfig} from 'axios'

export class ApiClient {
  axiosInstance: AxiosInstance
  constructor(baseURL = '') {
    this.axiosInstance = axios.create({
      baseURL
    })

    this.axiosInstance.interceptors.request.use(
      async (config: AxiosRequestConfig | any) => {
        if (localStorage.getItem('token')) {
          config.headers.authorization = 'Bearer ' + localStorage.getItem('token')
        }
        return config
      }, (err: AxiosError) => {
        return Promise.reject(err)
      }
    )
  }

  async get<T=object>(path: string, params: object={}): Promise<AxiosResponse<T>> {
    try {
      const result = await this.axiosInstance.get(path, {
        params
      })
      return this.successPromise(result.data)
    } catch(err: any) {
      return this.failurePromise<T>(err)
    }
  }

  async post<T=object>(path: string, params: object={}): Promise<AxiosResponse<T>> {
    try {
      const result = await this.axiosInstance.post(path, params)
      return this.successPromise(result.data)
    } catch(err: any) {
      return this.failurePromise<T>(err)
    }
  }

  async put<T=object>(path: string, params: object={}): Promise<AxiosResponse<T>> {
    try {
      const result = await this.axiosInstance.put(path, params)
      return this.successPromise(result.data)
    } catch(err: any) {
      return this.failurePromise<T>(err)
    }
  }

  async patch<T=object>(path: string, params: object={}): Promise<AxiosResponse<T>> {
    try {
      const result = await this.axiosInstance.patch(path, {
        params
      })
      return this.successPromise(result.data)
    } catch(err: any) {
      return this.failurePromise<T>(err)
    }
  }

  async delete<T=object>(path: string): Promise<AxiosResponse<T>> {
    try {
      const result = await this.axiosInstance.delete(path)
      return this.successPromise(result.data)
    } catch(err: any) {
      return this.failurePromise<T>(err)
    }
  }

  private successPromise<T>(data: T): Promise<AxiosResponse<T>> {
    return Promise.resolve<AxiosResponse<T>> ({
      data,
      isSuccess: true,
    })
  }

  private failurePromise<T>(error: AxiosError): Promise<AxiosResponse<T>> {
    return Promise.reject<AxiosResponse<T>> ({
      error,
      isSuccess: false,
    })
  }
}



type AxiosResponse<T> = {
  data?: T
  error?: AxiosError
  isSuccess: boolean
}