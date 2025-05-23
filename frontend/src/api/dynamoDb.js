import { get, post, del } from './request';

const listTables = async (prefix, limit) => {
  const suffix = prefix === '' ? '' : `/${encodeURIComponent(prefix)}`

  return get(`/api/dynamodb-list/${limit}${suffix}`);
}

const deleteTable = async (tableName) => {
  return del(`/api/dynamodb/${encodeURIComponent(tableName)}`)
}

const createTable = async (payload) => {
  return post(`/api/dynamodb`, payload);
}

const scanTable = async (payload) => {
  return post(`/api/scan_dynamodb`, payload);
}

export { listTables, deleteTable, createTable, scanTable };
