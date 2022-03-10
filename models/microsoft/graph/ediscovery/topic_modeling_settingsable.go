package ediscovery

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TopicModelingSettingsable 
type TopicModelingSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDynamicallyAdjustTopicCount()(*bool)
    GetIgnoreNumbers()(*bool)
    GetIsEnabled()(*bool)
    GetTopicCount()(*int32)
    SetDynamicallyAdjustTopicCount(value *bool)()
    SetIgnoreNumbers(value *bool)()
    SetIsEnabled(value *bool)()
    SetTopicCount(value *int32)()
}
