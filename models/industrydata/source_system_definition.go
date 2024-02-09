package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type SourceSystemDefinition struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewSourceSystemDefinition instantiates a new SourceSystemDefinition and sets the default values.
func NewSourceSystemDefinition()(*SourceSystemDefinition) {
    m := &SourceSystemDefinition{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSourceSystemDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSourceSystemDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSourceSystemDefinition(), nil
}
// GetDisplayName gets the displayName property value. The name of the source system. Maximum supported length is 100 characters.
// returns a *string when successful
func (m *SourceSystemDefinition) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SourceSystemDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["userMatchingSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserMatchingSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserMatchingSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserMatchingSettingable)
                }
            }
            m.SetUserMatchingSettings(res)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorEscaped(val)
        }
        return nil
    }
    return res
}
// GetUserMatchingSettings gets the userMatchingSettings property value. A collection of user matching settings by roleGroup.
// returns a []UserMatchingSettingable when successful
func (m *SourceSystemDefinition) GetUserMatchingSettings()([]UserMatchingSettingable) {
    val, err := m.GetBackingStore().Get("userMatchingSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserMatchingSettingable)
    }
    return nil
}
// GetVendorEscaped gets the vendor property value. The name of the vendor who supplies the source system. Maximum supported length is 100 characters.
// returns a *string when successful
func (m *SourceSystemDefinition) GetVendorEscaped()(*string) {
    val, err := m.GetBackingStore().Get("vendorEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SourceSystemDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetUserMatchingSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserMatchingSettings()))
        for i, v := range m.GetUserMatchingSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userMatchingSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendor", m.GetVendorEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the source system. Maximum supported length is 100 characters.
func (m *SourceSystemDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserMatchingSettings sets the userMatchingSettings property value. A collection of user matching settings by roleGroup.
func (m *SourceSystemDefinition) SetUserMatchingSettings(value []UserMatchingSettingable)() {
    err := m.GetBackingStore().Set("userMatchingSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorEscaped sets the vendor property value. The name of the vendor who supplies the source system. Maximum supported length is 100 characters.
func (m *SourceSystemDefinition) SetVendorEscaped(value *string)() {
    err := m.GetBackingStore().Set("vendorEscaped", value)
    if err != nil {
        panic(err)
    }
}
type SourceSystemDefinitionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetUserMatchingSettings()([]UserMatchingSettingable)
    GetVendorEscaped()(*string)
    SetDisplayName(value *string)()
    SetUserMatchingSettings(value []UserMatchingSettingable)()
    SetVendorEscaped(value *string)()
}
