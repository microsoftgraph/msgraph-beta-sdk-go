package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AdvancedThreatProtectionOnboardingStateSummaryable 
type AdvancedThreatProtectionOnboardingStateSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdvancedThreatProtectionOnboardingDeviceSettingStates()([]AdvancedThreatProtectionOnboardingDeviceSettingStateable)
    GetCompliantDeviceCount()(*int32)
    GetConflictDeviceCount()(*int32)
    GetErrorDeviceCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetNotAssignedDeviceCount()(*int32)
    GetRemediatedDeviceCount()(*int32)
    GetUnknownDeviceCount()(*int32)
    SetAdvancedThreatProtectionOnboardingDeviceSettingStates(value []AdvancedThreatProtectionOnboardingDeviceSettingStateable)()
    SetCompliantDeviceCount(value *int32)()
    SetConflictDeviceCount(value *int32)()
    SetErrorDeviceCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetNotAssignedDeviceCount(value *int32)()
    SetRemediatedDeviceCount(value *int32)()
    SetUnknownDeviceCount(value *int32)()
}