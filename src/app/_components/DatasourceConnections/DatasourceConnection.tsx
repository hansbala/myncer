import { type DATASOURCE_SCHEMA } from '~/app/_core/clientSideDatasourceSchema'
import { toTitleCase } from '~/core/utils/strings'
import ComponentLoading from '../Loading/ComponentLoading'

interface DatasourceConnectionProps {
  datasourceSchema: DATASOURCE_SCHEMA
  isAuthenticatedLoading: boolean
  isAuthenticationUrlLoading: boolean
  authenticationUrl?: string
  isAuthenticated?: boolean
}

export default function DatasourceConnection({
  isAuthenticatedLoading,
  isAuthenticationUrlLoading,
  datasourceSchema,
  isAuthenticated,
  authenticationUrl,
}: DatasourceConnectionProps) {
  const IconComponent = datasourceSchema.clientIcon

  if (isAuthenticatedLoading || isAuthenticationUrlLoading) {
    return <ComponentLoading />
  }

  return (
    <div className="flex flex-col rounded-lg border p-6 shadow-md">
      <div className="mb-2 flex flex-row items-center justify-between">
        <IconComponent className="h-12 w-12" />
        {isAuthenticated ? (
          <button className="btn btn-disabled">Connected</button>
        ) : (
          <button
            className="btn btn-primary"
            onClick={() => window.open(authenticationUrl, '_blank')}
          >
            Connect
          </button>
        )}
      </div>
      <h2 className="mb-2">{toTitleCase(datasourceSchema.name)}</h2>
      <p>{datasourceSchema.description}</p>
    </div>
  )
}
