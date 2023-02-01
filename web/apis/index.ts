import service from './service'

// export const listCategories = params => service.get('/categories', { params })

export const listNetworkInterfaces = (): Promise<Array<string>> => service.get('/network-interfaces')

export const putNetworkInterfaceIP = (name: string, params: {
    name: string,
    ip: string,
    mask: string,
    gateway: string
}): Promise<string> => service.put(`/network-interface/${name}/ipv4`, params)

export const putComputerName = (params: {
    newName: string
}): Promise<string> => service.put('/computer-name', params)

export const getConfig = (): Promise<Record<string, any>> => service.get('/config')

export const putConfig = (params: Record<string, any>): Promise<Record<string, any>> => service.put('/config', params)