import {extent, ticks} from 'd3-array'
import {scaleLinear, scaleOrdinal} from 'd3-scale'
import {produce} from 'immer'
import chroma from 'chroma-js'

import {
  PlotEnv,
  Layer,
  PLOT_PADDING,
  TICK_CHAR_WIDTH,
  TICK_CHAR_HEIGHT,
  TICK_PADDING_RIGHT,
  TICK_PADDING_TOP,
} from 'src/minard'
import {PlotAction} from 'src/minard/utils/plotEnvActions'
import {getGroupKey} from 'src/minard/utils/getGroupKey'

export const INITIAL_PLOT_ENV: PlotEnv = {
  width: 0,
  height: 0,
  innerWidth: 0,
  innerHeight: 0,
  defaults: {
    table: {columns: {}, columnTypes: {}},
    aesthetics: {},
    scales: {},
  },
  layers: {},
  xTicks: [],
  yTicks: [],
  xDomain: [],
  yDomain: [],
  hoverX: null,
  hoverY: null,
  margins: {
    top: PLOT_PADDING,
    right: PLOT_PADDING,
    bottom: PLOT_PADDING,
    left: PLOT_PADDING,
  },
  dispatch: () => {},
}

export const plotEnvReducer = (state: PlotEnv, action: PlotAction): PlotEnv =>
  produce(state, draftState => {
    switch (action.type) {
      case 'REGISTER_LAYER': {
        const {layerKey, table, aesthetics, colors} = action.payload

        draftState.layers[layerKey] = {table, aesthetics, colors, scales: {}}

        computeXYDomain(draftState)
        computeXYLayout(draftState)
        computeFillScales(draftState)

        return
      }

      case 'UNREGISTER_LAYER': {
        const {layerKey} = action.payload

        delete draftState.layers[layerKey]

        computeXYDomain(draftState)
        computeXYLayout(draftState)
        computeFillScales(draftState)

        return
      }

      case 'SET_DIMENSIONS': {
        const {width, height} = action.payload

        draftState.width = width
        draftState.height = height

        computeXYLayout(draftState)

        return
      }

      case 'SET_TABLE': {
        draftState.defaults.table = action.payload.table

        return
      }

      case 'SET_COLORS': {
        const {colors, layerKey} = action.payload

        if (layerKey) {
          draftState.layers[layerKey].colors = colors
        } else {
          draftState.defaults.colors = colors
        }

        computeFillScales(draftState)

        return
      }
    }
  })

const getCols = (state: PlotEnv, aestheticNames: string[]): any[][] => {
  const {defaults, layers} = state

  const cols = []

  for (const layer of [defaults, ...Object.values(layers)]) {
    for (const aestheticName of aestheticNames) {
      if (layer.aesthetics[aestheticName]) {
        const colName = layer.aesthetics[aestheticName]
        const col = layer.table
          ? layer.table.columns[colName]
          : defaults.table.columns[colName]

        cols.push(col)
      }
    }
  }

  return cols
}

const flatten = (arrays: any[][]): any[] => [].concat(...arrays)

// TODO: Memoize computation by comparing to previous state
const computeXYDomain = (draftState: PlotEnv): void => {
  draftState.xDomain = extent(
    flatten(getCols(draftState, ['x', 'xMin', 'xMax']).map(col => extent(col)))
  )

  draftState.yDomain = extent(
    flatten(getCols(draftState, ['y', 'yMin', 'yMax']).map(col => extent(col)))
  )
}

const getTicks = ([d0, d1]: number[], length: number): string[] => {
  const approxTickWidth =
    Math.max(String(d0).length, String(d1).length) * TICK_CHAR_WIDTH

  const TICK_DENSITY = 0.3
  const numTicks = Math.round((length / approxTickWidth) * TICK_DENSITY)
  const result = ticks(d0, d1, numTicks).map(t => String(t))

  return result
}

const computeXYLayout = (draftState: PlotEnv): void => {
  const {width, height, xDomain, yDomain} = draftState

  draftState.xTicks = getTicks(xDomain, width)
  draftState.yTicks = getTicks(yDomain, height)

  const yTickWidth =
    Math.max(...draftState.yTicks.map(t => t.length)) * TICK_CHAR_WIDTH

  const margins = {
    top: PLOT_PADDING,
    right: PLOT_PADDING,
    bottom: TICK_CHAR_HEIGHT + TICK_PADDING_TOP + PLOT_PADDING,
    left: yTickWidth + TICK_PADDING_RIGHT + PLOT_PADDING,
  }

  const innerWidth = width - margins.left - margins.right
  const innerHeight = height - margins.top - margins.bottom

  draftState.margins = margins
  draftState.innerWidth = innerWidth
  draftState.innerHeight = innerHeight

  draftState.defaults.scales.x = scaleLinear()
    .domain(xDomain)
    .range([0, innerWidth])

  draftState.defaults.scales.y = scaleLinear()
    .domain(yDomain)
    .range([innerHeight, 0])
}

const getColorScale = (domain: any[], colors: string[]) => {
  const range = chroma
    .scale(colors)
    .mode('lch')
    .colors(domain.length)

  const scale = scaleOrdinal()
    .domain(domain)
    .range(range)

  return scale
}

const getFillDomain = ({table, aesthetics}: Layer): string[] => {
  const fillColKeys = aesthetics.fill

  if (!fillColKeys.length) {
    return []
  }

  const fillDomain = new Set()
  const n = Object.values(table.columns)[0].length

  for (let i = 0; i < n; i++) {
    fillDomain.add(getGroupKey(fillColKeys.map(k => table.columns[k][i])))
  }

  return [...fillDomain].sort()
}

const computeFillScales = (draftState: PlotEnv) => {
  const layers = Object.values(draftState.layers)

  layers
    .filter(
      layer => layer.aesthetics.fill && layer.colors && layer.colors.length
    )
    .forEach(layer => {
      layer.scales.fill = getColorScale(getFillDomain(layer), layer.colors)
    })
}
