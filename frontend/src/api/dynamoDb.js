import { get, post, del } from './request';

const listTables = async (prefix, limit) => {
    return get(`/api/dynamodb/${encodeURIComponent(prefix)}/${limit}`);
}

const deleteTable = async (tableName) => {
  return del(`/api/dynamodb/${encodeURIComponent(tableName)}`)  
}

const createTable = async (payload) => {
  return post(`/api/dynamodb`, payload);
}

export { listTables, deleteTable, createTable };
