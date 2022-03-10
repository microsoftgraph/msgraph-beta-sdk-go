package windowsupdates

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeploymentStateable 
type DeploymentStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetReasons()([]DeploymentStateReasonable)
    GetRequestedValue()(*RequestedDeploymentStateValue)
    GetValue()(*DeploymentStateValue)
    SetReasons(value []DeploymentStateReasonable)()
    SetRequestedValue(value *RequestedDeploymentStateValue)()
    SetValue(value *DeploymentStateValue)()
}
