package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookTaskGroupable 
type OutlookTaskGroupable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeKey()(*string)
    GetGroupKey()(*string)
    GetIsDefaultGroup()(*bool)
    GetName()(*string)
    GetTaskFolders()([]OutlookTaskFolderable)
    SetChangeKey(value *string)()
    SetGroupKey(value *string)()
    SetIsDefaultGroup(value *bool)()
    SetName(value *string)()
    SetTaskFolders(value []OutlookTaskFolderable)()
}
