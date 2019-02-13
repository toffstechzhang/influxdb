// Libraries
import React, {PureComponent} from 'react'
import _ from 'lodash'

// Components
import {ErrorHandling} from 'src/shared/decorators/errors'
import EmptyDataSourceState from 'src/dataLoaders/components/configureStep/EmptyDataSourceState'
import Scraping from 'src/dataLoaders/components/configureStep/Scraping'

// Types
import {DataLoaderType} from 'src/types/v2/dataLoaders'
import {Bucket} from '@influxdata/influx'

export interface Props {
  dataLoaderType: DataLoaderType
  buckets: Bucket[]
  bucket: string
  org: string
  onClickNext: () => void
}

@ErrorHandling
class ConfigureDataSourceSwitcher extends PureComponent<Props> {
  public render() {
    const {dataLoaderType, onClickNext, buckets} = this.props

    switch (dataLoaderType) {
      case DataLoaderType.Scraping:
        return <Scraping onClickNext={onClickNext} buckets={buckets} />
      default:
        return <EmptyDataSourceState />
    }
  }
}

export default ConfigureDataSourceSwitcher
