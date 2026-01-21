/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions, type InfiniteData  } from "@tanstack/vue-query";
import { useClient } from '../useClient';

export default function useProjectbitProjectbitV_1() {
  const client = useClient();

  type QueryParamsMethod = typeof client.ProjectbitProjectbitV_1.query.queryParams;
  type QueryParamsData = Awaited<ReturnType<QueryParamsMethod>>["data"];
  const QueryParams = ( options: Partial<UseQueryOptions<QueryParamsData>>) => {
    const key = { type: 'QueryParams',  };    
    return useQuery<QueryParamsData>({ queryKey: [key], queryFn: async () => {
      const res = await client.ProjectbitProjectbitV_1.query.queryParams();
        return res.data;
    }, ...options});
  }
  

  type QueryGetPostMethod = typeof client.ProjectbitProjectbitV_1.query.queryGetPost;
  type QueryGetPostData = Awaited<ReturnType<QueryGetPostMethod>>["data"];
  const QueryGetPost = (id: string,  options: Partial<UseQueryOptions<QueryGetPostData>>) => {
    const key = { type: 'QueryGetPost',  id };    
    return useQuery<QueryGetPostData>({ queryKey: [key], queryFn: async () => {
      const { id } = key
      const res = await client.ProjectbitProjectbitV_1.query.queryGetPost(id);
        return res.data;
    }, ...options});
  }
  
  type QueryListPostMethod = typeof client.ProjectbitProjectbitV_1.query.queryListPost;
  type QueryListPostData = Awaited<ReturnType<QueryListPostMethod>>["data"] & { pageParam: number };
  const QueryListPost = (query:  NonNullable<Parameters<QueryListPostMethod>[0]>, options:  Partial<UseInfiniteQueryOptions<QueryListPostData, unknown, InfiniteData<QueryListPostData,number>, Array<string | unknown>, number>> , perPage: number) => {
    const key = { type: 'QueryListPost', query };    
    return useInfiniteQuery<QueryListPostData, unknown, InfiniteData<QueryListPostData,number>, Array<string | unknown>, number>({ queryKey: [key], queryFn: async (context: {pageParam?: number}) => {
      const { pageParam=1 } = context;
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      const res = await client.ProjectbitProjectbitV_1.query.queryListPost(query ?? undefined);
        return { ...res.data, pageParam }; 
    }, ...options,
      initialPageParam: 1,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  

  type QuerySayHelloMethod = typeof client.ProjectbitProjectbitV_1.query.querySayHello;
  type QuerySayHelloData = Awaited<ReturnType<QuerySayHelloMethod>>["data"];
  const QuerySayHello = (name: string,  options: Partial<UseQueryOptions<QuerySayHelloData>>) => {
    const key = { type: 'QuerySayHello',  name };    
    return useQuery<QuerySayHelloData>({ queryKey: [key], queryFn: async () => {
      const { name } = key
      const res = await client.ProjectbitProjectbitV_1.query.querySayHello(name);
        return res.data;
    }, ...options});
  }
  
  return {QueryParams,QueryGetPost,QueryListPost,QuerySayHello,
  }
}
