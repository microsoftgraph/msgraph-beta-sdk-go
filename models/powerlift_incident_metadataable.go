package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PowerliftIncidentMetadataable 
type PowerliftIncidentMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(*string)
    GetClientVersion()(*string)
    GetCreatedAtDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEasyId()(*string)
    GetFileNames()([]string)
    GetLocale()(*string)
    GetOdataType()(*string)
    GetPlatform()(*string)
    GetPowerliftId()(*UUID)
    SetApplication(value *string)()
    SetClientVersion(value *string)()
    SetCreatedAtDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEasyId(value *string)()
    SetFileNames(value []string)()
    SetLocale(value *string)()
    SetOdataType(value *string)()
    SetPlatform(value *string)()
    SetPowerliftId(value *UUID)()
}
