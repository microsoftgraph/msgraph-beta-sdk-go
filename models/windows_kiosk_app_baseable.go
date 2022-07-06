package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskAppBaseable 
type WindowsKioskAppBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppType()(*WindowsKioskAppType)
    GetAutoLaunch()(*bool)
    GetName()(*string)
    GetStartLayoutTileSize()(*WindowsAppStartLayoutTileSize)
    GetType()(*string)
    SetAppType(value *WindowsKioskAppType)()
    SetAutoLaunch(value *bool)()
    SetName(value *string)()
    SetStartLayoutTileSize(value *WindowsAppStartLayoutTileSize)()
    SetType(value *string)()
}
