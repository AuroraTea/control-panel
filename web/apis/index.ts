import service from './service'

// export const listCategories = params => service.get('/categories', { params })

export const getNetAdapters = (): Promise<Array<string>> => service.get('/net-adapters')

export const setIPv4 = (params: {
  netAdapter: string
  ip: string
  mask: string
  gateway: string
}): Promise<string> => service.put('/ipv4', params)

export const setComputerName =
  (params: { newName: string }): Promise<string> =>
    service.put('/computer-name', params)

export const getConfig = (): Promise<Record<string, any>> => service.get('/config')

export const setConfig =
  (params: Record<string, any>): Promise<Record<string, any>> =>
    service.put('/config', params)
