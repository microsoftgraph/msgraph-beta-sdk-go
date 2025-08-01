// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SharePointMigrationJobStartEvent struct {
    SharePointMigrationEvent
}
// NewSharePointMigrationJobStartEvent instantiates a new SharePointMigrationJobStartEvent and sets the default values.
func NewSharePointMigrationJobStartEvent()(*SharePointMigrationJobStartEvent) {
    m := &SharePointMigrationJobStartEvent{
        SharePointMigrationEvent: *NewSharePointMigrationEvent(),
    }
    return m
}
// CreateSharePointMigrationJobStartEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSharePointMigrationJobStartEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointMigrationJobStartEvent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SharePointMigrationJobStartEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SharePointMigrationEvent.GetFieldDeserializers()
    res["isRestarted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRestarted(val)
        }
        return nil
    }
    res["totalRetryCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRetryCount(val)
        }
        return nil
    }
    return res
}
// GetIsRestarted gets the isRestarted property value. The isRestarted property
// returns a *bool when successful
func (m *SharePointMigrationJobStartEvent) GetIsRestarted()(*bool) {
    val, err := m.GetBackingStore().Get("isRestarted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTotalRetryCount gets the totalRetryCount property value. The totalRetryCount property
// returns a *int32 when successful
func (m *SharePointMigrationJobStartEvent) GetTotalRetryCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalRetryCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SharePointMigrationJobStartEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SharePointMigrationEvent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isRestarted", m.GetIsRestarted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalRetryCount", m.GetTotalRetryCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsRestarted sets the isRestarted property value. The isRestarted property
func (m *SharePointMigrationJobStartEvent) SetIsRestarted(value *bool)() {
    err := m.GetBackingStore().Set("isRestarted", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalRetryCount sets the totalRetryCount property value. The totalRetryCount property
func (m *SharePointMigrationJobStartEvent) SetTotalRetryCount(value *int32)() {
    err := m.GetBackingStore().Set("totalRetryCount", value)
    if err != nil {
        panic(err)
    }
}
type SharePointMigrationJobStartEventable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SharePointMigrationEventable
    GetIsRestarted()(*bool)
    GetTotalRetryCount()(*int32)
    SetIsRestarted(value *bool)()
    SetTotalRetryCount(value *int32)()
}
