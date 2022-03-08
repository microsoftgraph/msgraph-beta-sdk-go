package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceOverviewable 
type ManagedDeviceOverviewable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceExchangeAccessStateSummary()(DeviceExchangeAccessStateSummaryable)
    GetDeviceOperatingSystemSummary()(DeviceOperatingSystemSummaryable)
    GetDualEnrolledDeviceCount()(*int32)
    GetEnrolledDeviceCount()(*int32)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceModelsAndManufacturers()(ManagedDeviceModelsAndManufacturersable)
    GetMdmEnrolledCount()(*int32)
    SetDeviceExchangeAccessStateSummary(value DeviceExchangeAccessStateSummaryable)()
    SetDeviceOperatingSystemSummary(value DeviceOperatingSystemSummaryable)()
    SetDualEnrolledDeviceCount(value *int32)()
    SetEnrolledDeviceCount(value *int32)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceModelsAndManufacturers(value ManagedDeviceModelsAndManufacturersable)()
    SetMdmEnrolledCount(value *int32)()
}
