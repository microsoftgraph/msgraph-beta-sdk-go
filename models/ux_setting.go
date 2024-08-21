package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UxSetting struct {
    Entity
}
// NewUxSetting instantiates a new UxSetting and sets the default values.
func NewUxSetting()(*UxSetting) {
    m := &UxSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUxSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUxSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUxSetting(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UxSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["restrictNonAdminAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNonAdminSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictNonAdminAccess(val.(*NonAdminSetting))
        }
        return nil
    }
    return res
}
// GetRestrictNonAdminAccess gets the restrictNonAdminAccess property value. The restrictNonAdminAccess property
// returns a *NonAdminSetting when successful
func (m *UxSetting) GetRestrictNonAdminAccess()(*NonAdminSetting) {
    val, err := m.GetBackingStore().Get("restrictNonAdminAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NonAdminSetting)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UxSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRestrictNonAdminAccess() != nil {
        cast := (*m.GetRestrictNonAdminAccess()).String()
        err = writer.WriteStringValue("restrictNonAdminAccess", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRestrictNonAdminAccess sets the restrictNonAdminAccess property value. The restrictNonAdminAccess property
func (m *UxSetting) SetRestrictNonAdminAccess(value *NonAdminSetting)() {
    err := m.GetBackingStore().Set("restrictNonAdminAccess", value)
    if err != nil {
        panic(err)
    }
}
type UxSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRestrictNonAdminAccess()(*NonAdminSetting)
    SetRestrictNonAdminAccess(value *NonAdminSetting)()
}
