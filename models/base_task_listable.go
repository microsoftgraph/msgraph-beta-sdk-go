package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseTaskListable 
type BaseTaskListable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetExtensions()([]Extensionable)
    GetTasks()([]BaseTaskable)
    SetDisplayName(value *string)()
    SetExtensions(value []Extensionable)()
    SetTasks(value []BaseTaskable)()
}
