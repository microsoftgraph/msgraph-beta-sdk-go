package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriverUpdateCatalogEntryable 
type DriverUpdateCatalogEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SoftwareUpdateCatalogEntryable
    GetDescription()(*string)
    GetDriverClass()(*string)
    GetManufacturer()(*string)
    GetProvider()(*string)
    GetSetupInformationFile()(*string)
    GetVersion()(*string)
    GetVersionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDescription(value *string)()
    SetDriverClass(value *string)()
    SetManufacturer(value *string)()
    SetProvider(value *string)()
    SetSetupInformationFile(value *string)()
    SetVersion(value *string)()
    SetVersionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
