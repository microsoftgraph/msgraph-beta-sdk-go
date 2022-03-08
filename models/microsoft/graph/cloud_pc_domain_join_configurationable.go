package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcDomainJoinConfigurationable 
type CloudPcDomainJoinConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetOnPremisesConnectionId()(*string)
    GetRegionName()(*string)
    GetType()(*CloudPcDomainJoinType)
    SetOnPremisesConnectionId(value *string)()
    SetRegionName(value *string)()
    SetType(value *CloudPcDomainJoinType)()
}
