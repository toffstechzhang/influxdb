// Libraries
import React, {PureComponent} from 'react'
import {connect} from 'react-redux'
import _ from 'lodash'

// Components
import {Form} from 'src/clockface'
import ConfigFieldHandler from 'src/dataLoaders/components/collectorsWizard/configure/ConfigFieldHandler'
import FancyScrollbar from 'src/shared/components/fancy_scrollbar/FancyScrollbar'

// Actions
import {
  setActiveTelegrafPlugin,
  setPluginConfiguration,
} from 'src/dataLoaders/actions/dataLoaders'

// Types
import {TelegrafPlugin, ConfigFields} from 'src/types/v2/dataLoaders'
import OnboardingButtons from 'src/onboarding/components/OnboardingButtons'
import {AppState} from 'src/types/v2'

interface OwnProps {
  telegrafPlugin: TelegrafPlugin
  configFields: ConfigFields
}

interface DispatchProps {
  onSetActiveTelegrafPlugin: typeof setActiveTelegrafPlugin
  onSetPluginConfiguration: typeof setPluginConfiguration
}

interface StateProps {
  telegrafPlugins: TelegrafPlugin[]
}

type Props = OwnProps & StateProps & DispatchProps

export class PluginConfigForm extends PureComponent<Props> {
  public render() {
    const {configFields, telegrafPlugin} = this.props
    return (
      <Form onSubmit={this.handleSubmitForm} className="data-loading--form">
        <FancyScrollbar
          autoHide={false}
          className="data-loading--scroll-content"
        >
          <div>
            <h3 className="wizard-step--title">
              {_.startCase(telegrafPlugin.name)}
            </h3>
            <h5 className="wizard-step--sub-title">
              For more information about this plugin, see{' '}
              <a
                target="_blank"
                data-test="docs-link"
                href={`https://github.com/influxdata/telegraf/tree/master/plugins/inputs/${
                  telegrafPlugin.name
                }`}
              >
                Documentation
              </a>
            </h5>
          </div>
          <ConfigFieldHandler
            configFields={configFields}
            telegrafPlugin={telegrafPlugin}
          />
        </FancyScrollbar>
        <OnboardingButtons
          autoFocusNext={this.autoFocus}
          nextButtonText={'Done'}
          className="data-loading--button-container"
        />
      </Form>
    )
  }

  private get autoFocus(): boolean {
    const {configFields} = this.props
    return !configFields
  }

  private handleSubmitForm = () => {
    const {
      telegrafPlugins,
      onSetPluginConfiguration,
      onSetActiveTelegrafPlugin,
    } = this.props

    const activeTelegrafPlugin = telegrafPlugins.find(tp => tp.active)
    if (!!activeTelegrafPlugin) {
      onSetPluginConfiguration(activeTelegrafPlugin.name)
    }

    onSetActiveTelegrafPlugin('')
  }
}

const mstp = ({
  dataLoading: {
    dataLoaders: {telegrafPlugins},
  },
}: AppState): StateProps => ({
  telegrafPlugins,
})

const mdtp: DispatchProps = {
  onSetActiveTelegrafPlugin: setActiveTelegrafPlugin,
  onSetPluginConfiguration: setPluginConfiguration,
}

export default connect<StateProps, DispatchProps, OwnProps>(
  mstp,
  mdtp
)(PluginConfigForm)
