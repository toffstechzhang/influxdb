import {Layer} from 'src/minard'
import {getGroupKey} from 'src/minard/utils/getGroupKey'

export const getBarFill = (
  {scales, aesthetics, table}: Layer,
  i: number
): string => {
  const fillScale = scales.fill
  const values = aesthetics.fill.map(colKey => table.columns[colKey][i])
  const groupKey = getGroupKey(values)
  const fill = fillScale(groupKey)

  return fill
}
