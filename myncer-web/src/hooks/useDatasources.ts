import { listDatasources } from "@/generated_grpc/myncer/datasource-DatasourceService_connectquery"
import { useQuery } from "@connectrpc/connect-query"

export const useDatasources = () => {
  const { data, isLoading } = useQuery(listDatasources)

  return {
    datasources: data?.datasources || [],
    loading: isLoading,
  }
}
