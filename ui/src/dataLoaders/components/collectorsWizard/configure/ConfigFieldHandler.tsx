// Libraries
import React, {PureComponent, ChangeEvent} from 'react'
import {connect} from 'react-redux'
import _ from 'lodash'

// Components
import ConfigFieldSwitcher from 'src/dataLoaders/components/configureStep/streaming/ConfigFieldSwitcher'

// Actions
import {
  updateTelegrafPluginConfig,
  addConfigValue,
  removeConfigValue,
  setConfigArrayValue,
} from 'src/dataLoaders/actions/dataLoaders'

// Types
import {
  TelegrafPlugin,
  ConfigFields,
  ConfigFieldType,
} from 'src/types/v2/dataLoaders'

interface OwnProps {
  configFields: ConfigFields
  telegrafPlugin: TelegrafPlugin
}

interface DispatchProps {
  onUpdateTelegrafPluginConfig: typeof updateTelegrafPluginConfig
  onAddConfigValue: typeof addConfigValue
  onRemoveConfigValue: typeof removeConfigValue
  onSetConfigArrayValue: typeof setConfigArrayValue
}

type Props = OwnProps & DispatchProps

export class ConfigFieldHandler extends PureComponent<Props> {
  public render() {
    return <div>{this.formFields}</div>
  }

  private get formFields(): JSX.Element[] | JSX.Element {
    const {configFields, telegrafPlugin, onSetConfigArrayValue} = this.props

    if (!configFields) {
      return <p data-test={'no-config'}>No configuration required.</p>
    }

    return Object.entries(configFields).map(
      ([fieldName, {type: fieldType, isRequired}], i) => {
        return (
          <ConfigFieldSwitcher
            key={fieldName}
            fieldName={fieldName}
            fieldType={fieldType}
            index={i}
            onChange={this.handleUpdateConfigField}
            value={this.getFieldValue(telegrafPlugin, fieldName, fieldType)}
            isRequired={isRequired}
            addTagValue={this.handleAddConfigFieldValue}
            removeTagValue={this.handleRemoveConfigFieldValue}
            onSetConfigArrayValue={onSetConfigArrayValue}
            telegrafPluginName={telegrafPlugin.name}
          />
        )
      }
    )
  }

  private handleAddConfigFieldValue = (
    value: string,
    fieldName: string
  ): void => {
    const {onAddConfigValue, telegrafPlugin} = this.props

    onAddConfigValue(telegrafPlugin.name, fieldName, value)
  }

  private handleRemoveConfigFieldValue = (value: string, fieldName: string) => {
    const {onRemoveConfigValue, telegrafPlugin} = this.props

    onRemoveConfigValue(telegrafPlugin.name, fieldName, value)
  }

  private getFieldValue(
    telegrafPlugin: TelegrafPlugin,
    fieldName: string,
    fieldType: ConfigFieldType
  ): string | string[] {
    let defaultEmpty: string | string[]
    if (
      fieldType === ConfigFieldType.String ||
      fieldType === ConfigFieldType.Uri
    ) {
      defaultEmpty = ''
    } else {
      defaultEmpty = []
    }
    return _.get(telegrafPlugin, `plugin.config.${fieldName}`, defaultEmpty)
  }

  private handleUpdateConfigField = (e: ChangeEvent<HTMLInputElement>) => {
    const {onUpdateTelegrafPluginConfig, telegrafPlugin} = this.props
    const {name, value} = e.target

    onUpdateTelegrafPluginConfig(telegrafPlugin.name, name, value)
  }
}

const mdtp: DispatchProps = {
  onUpdateTelegrafPluginConfig: updateTelegrafPluginConfig,
  onAddConfigValue: addConfigValue,
  onRemoveConfigValue: removeConfigValue,
  onSetConfigArrayValue: setConfigArrayValue,
}

export default connect<null, DispatchProps, OwnProps>(
  null,
  mdtp
)(ConfigFieldHandler)
