package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceConfigurationConflictSummaryable 
type DeviceConfigurationConflictSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConflictingDeviceConfigurations()([]SettingSourceable)
    GetContributingSettings()([]string)
    GetDeviceCheckinsImpacted()(*int32)
    SetConflictingDeviceConfigurations(value []SettingSourceable)()
    SetContributingSettings(value []string)()
    SetDeviceCheckinsImpacted(value *int32)()
}
