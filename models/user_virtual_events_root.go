package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserVirtualEventsRoot struct {
    Entity
}
// NewUserVirtualEventsRoot instantiates a new UserVirtualEventsRoot and sets the default values.
func NewUserVirtualEventsRoot()(*UserVirtualEventsRoot) {
    m := &UserVirtualEventsRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserVirtualEventsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserVirtualEventsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserVirtualEventsRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserVirtualEventsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["webinars"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventWebinarFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventWebinarable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventWebinarable)
                }
            }
            m.SetWebinars(res)
        }
        return nil
    }
    return res
}
// GetWebinars gets the webinars property value. The webinars property
// returns a []VirtualEventWebinarable when successful
func (m *UserVirtualEventsRoot) GetWebinars()([]VirtualEventWebinarable) {
    val, err := m.GetBackingStore().Get("webinars")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventWebinarable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserVirtualEventsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetWebinars() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebinars()))
        for i, v := range m.GetWebinars() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("webinars", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWebinars sets the webinars property value. The webinars property
func (m *UserVirtualEventsRoot) SetWebinars(value []VirtualEventWebinarable)() {
    err := m.GetBackingStore().Set("webinars", value)
    if err != nil {
        panic(err)
    }
}
type UserVirtualEventsRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetWebinars()([]VirtualEventWebinarable)
    SetWebinars(value []VirtualEventWebinarable)()
}
