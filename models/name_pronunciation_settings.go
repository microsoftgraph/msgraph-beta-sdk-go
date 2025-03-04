package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NamePronunciationSettings struct {
    Entity
}
// NewNamePronunciationSettings instantiates a new NamePronunciationSettings and sets the default values.
func NewNamePronunciationSettings()(*NamePronunciationSettings) {
    m := &NamePronunciationSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateNamePronunciationSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamePronunciationSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamePronunciationSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamePronunciationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEnabledInOrganization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledInOrganization(val)
        }
        return nil
    }
    return res
}
// GetIsEnabledInOrganization gets the isEnabledInOrganization property value. true to enable name pronunciation in the organization; otherwise, false. The default value is false.
// returns a *bool when successful
func (m *NamePronunciationSettings) GetIsEnabledInOrganization()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabledInOrganization")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *NamePronunciationSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabledInOrganization", m.GetIsEnabledInOrganization())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsEnabledInOrganization sets the isEnabledInOrganization property value. true to enable name pronunciation in the organization; otherwise, false. The default value is false.
func (m *NamePronunciationSettings) SetIsEnabledInOrganization(value *bool)() {
    err := m.GetBackingStore().Set("isEnabledInOrganization", value)
    if err != nil {
        panic(err)
    }
}
type NamePronunciationSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsEnabledInOrganization()(*bool)
    SetIsEnabledInOrganization(value *bool)()
}
