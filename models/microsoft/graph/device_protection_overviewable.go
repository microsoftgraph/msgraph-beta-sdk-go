package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceProtectionOverviewable 
type DeviceProtectionOverviewable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCleanDeviceCount()(*int32)
    GetCriticalFailuresDeviceCount()(*int32)
    GetInactiveThreatAgentDeviceCount()(*int32)
    GetPendingFullScanDeviceCount()(*int32)
    GetPendingManualStepsDeviceCount()(*int32)
    GetPendingOfflineScanDeviceCount()(*int32)
    GetPendingQuickScanDeviceCount()(*int32)
    GetPendingRestartDeviceCount()(*int32)
    GetPendingSignatureUpdateDeviceCount()(*int32)
    GetTotalReportedDeviceCount()(*int32)
    GetUnknownStateThreatAgentDeviceCount()(*int32)
    SetCleanDeviceCount(value *int32)()
    SetCriticalFailuresDeviceCount(value *int32)()
    SetInactiveThreatAgentDeviceCount(value *int32)()
    SetPendingFullScanDeviceCount(value *int32)()
    SetPendingManualStepsDeviceCount(value *int32)()
    SetPendingOfflineScanDeviceCount(value *int32)()
    SetPendingQuickScanDeviceCount(value *int32)()
    SetPendingRestartDeviceCount(value *int32)()
    SetPendingSignatureUpdateDeviceCount(value *int32)()
    SetTotalReportedDeviceCount(value *int32)()
    SetUnknownStateThreatAgentDeviceCount(value *int32)()
}
