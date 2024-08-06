package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationalBrandingTheme struct {
    Entity
}
// NewOrganizationalBrandingTheme instantiates a new OrganizationalBrandingTheme and sets the default values.
func NewOrganizationalBrandingTheme()(*OrganizationalBrandingTheme) {
    m := &OrganizationalBrandingTheme{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOrganizationalBrandingThemeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationalBrandingThemeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalBrandingTheme(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationalBrandingTheme) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isDefaultTheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultTheme(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetIsDefaultTheme gets the isDefaultTheme property value. The isDefaultTheme property
// returns a *bool when successful
func (m *OrganizationalBrandingTheme) GetIsDefaultTheme()(*bool) {
    val, err := m.GetBackingStore().Get("isDefaultTheme")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OrganizationalBrandingTheme) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OrganizationalBrandingTheme) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefaultTheme", m.GetIsDefaultTheme())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefaultTheme sets the isDefaultTheme property value. The isDefaultTheme property
func (m *OrganizationalBrandingTheme) SetIsDefaultTheme(value *bool)() {
    err := m.GetBackingStore().Set("isDefaultTheme", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *OrganizationalBrandingTheme) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
type OrganizationalBrandingThemeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsDefaultTheme()(*bool)
    GetName()(*string)
    SetIsDefaultTheme(value *bool)()
    SetName(value *string)()
}
