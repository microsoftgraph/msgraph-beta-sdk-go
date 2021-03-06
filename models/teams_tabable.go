package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsTabable 
type TeamsTabable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(TeamsTabConfigurationable)
    GetDisplayName()(*string)
    GetMessageId()(*string)
    GetSortOrderIndex()(*string)
    GetTeamsApp()(TeamsAppable)
    GetTeamsAppId()(*string)
    GetWebUrl()(*string)
    SetConfiguration(value TeamsTabConfigurationable)()
    SetDisplayName(value *string)()
    SetMessageId(value *string)()
    SetSortOrderIndex(value *string)()
    SetTeamsApp(value TeamsAppable)()
    SetTeamsAppId(value *string)()
    SetWebUrl(value *string)()
}
