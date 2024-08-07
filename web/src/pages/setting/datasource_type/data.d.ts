export interface TableListItem {
  id: number;
  name: string;
  description: string;
  sort:number;
  enable:number;
  gmt_created: date;
  gmt_updated: date;
}

export interface TableListPagination {
  total: number;
  pageSize: number;
  current: number;
}

export interface TableListData {
  list: TableListItem[];
  pagination: Partial<TableListPagination>;
}

export interface TableListParams {
  id?: number;
  name?: string;
  description?: string;
  pageSize?: number;
  currentPage?: number;
  filter?: { [key: string]: any[] };
  sorter?: { [key: string]: any };
}
