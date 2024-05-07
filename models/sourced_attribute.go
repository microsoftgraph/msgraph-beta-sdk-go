package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SourcedAttribute struct {
    CustomClaimAttributeBase
}
// NewSourcedAttribute instantiates a new SourcedAttribute and sets the default values.
func NewSourcedAttribute()(*SourcedAttribute) {
    m := &SourcedAttribute{
        CustomClaimAttributeBase: *NewCustomClaimAttributeBase(),
    }
    odataTypeValue := "#microsoft.graph.sourcedAttribute"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSourcedAttributeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSourcedAttributeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSourcedAttribute(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SourcedAttribute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimAttributeBase.GetFieldDeserializers()
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isExtensionAttribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExtensionAttribute(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *SourcedAttribute) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsExtensionAttribute gets the isExtensionAttribute property value. The isExtensionAttribute property
// returns a *bool when successful
func (m *SourcedAttribute) GetIsExtensionAttribute()(*bool) {
    val, err := m.GetBackingStore().Get("isExtensionAttribute")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSource gets the source property value. The source property
// returns a *string when successful
func (m *SourcedAttribute) GetSource()(*string) {
    val, err := m.GetBackingStore().Get("source")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SourcedAttribute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimAttributeBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isExtensionAttribute", m.GetIsExtensionAttribute())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetId sets the id property value. The id property
func (m *SourcedAttribute) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExtensionAttribute sets the isExtensionAttribute property value. The isExtensionAttribute property
func (m *SourcedAttribute) SetIsExtensionAttribute(value *bool)() {
    err := m.GetBackingStore().Set("isExtensionAttribute", value)
    if err != nil {
        panic(err)
    }
}
// SetSource sets the source property value. The source property
func (m *SourcedAttribute) SetSource(value *string)() {
    err := m.GetBackingStore().Set("source", value)
    if err != nil {
        panic(err)
    }
}
type SourcedAttributeable interface {
    CustomClaimAttributeBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetIsExtensionAttribute()(*bool)
    GetSource()(*string)
    SetId(value *string)()
    SetIsExtensionAttribute(value *bool)()
    SetSource(value *string)()
}
