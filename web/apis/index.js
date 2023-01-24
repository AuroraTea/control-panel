import service from './service'

export const listCategories = params => service.get('/categories', { params })

export const listNetworkInterfaces = params => service.get('/network-interfaces', { params })

export const putNetworkInterfaceIP = (name, params) => service.put(`/network-interface/${name}/ipv4`, params)

export const putComputerName = params => service.put('/computer-name', params)