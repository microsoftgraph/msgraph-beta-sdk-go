package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Printerable 
type Printerable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    PrinterBaseable
    GetAcceptingJobs()(*bool)
    GetConnectors()([]PrintConnectorable)
    GetHasPhysicalDevice()(*bool)
    GetIsShared()(*bool)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetShare()(PrinterShareable)
    GetShares()([]PrinterShareable)
    GetTaskTriggers()([]PrintTaskTriggerable)
    SetAcceptingJobs(value *bool)()
    SetConnectors(value []PrintConnectorable)()
    SetHasPhysicalDevice(value *bool)()
    SetIsShared(value *bool)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetShare(value PrinterShareable)()
    SetShares(value []PrinterShareable)()
    SetTaskTriggers(value []PrintTaskTriggerable)()
}
