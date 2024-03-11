'use client'

import { DATASOURCE_SCHEMAS } from '~/app/_core/clientSideDatasourceSchema'
import { type DATASOURCE } from '~/core/datasources'

interface DatasourceDropdownProps {
  setDatasource: (datasource: DATASOURCE) => void
  datasource?: DATASOURCE
  className?: string
}
export default function DatasourceDropdown({
  datasource: propsDatasource,
  setDatasource: setDatasourceProps,
  className: classNameProps,
}: DatasourceDropdownProps) {
  return (
    <div className={classNameProps}>
      <details className="dropdown">
        <summary className="btn m-1">
          {propsDatasource ?? 'Choose a datasource'}{' '}
        </summary>

        <ul className="menu dropdown-content z-[1] w-52 rounded-box bg-base-100 p-2 shadow">
          {Object.entries(DATASOURCE_SCHEMAS).map(
            ([datasource, datasourceSchema]) => {
              return (
                <li key={datasource}>
                  <a
                    onClick={() =>
                      setDatasourceProps(datasourceSchema.datasource)
                    }
                  >
                    {datasource}
                  </a>
                </li>
              )
            }
          )}
        </ul>
      </details>
    </div>
  )
}
