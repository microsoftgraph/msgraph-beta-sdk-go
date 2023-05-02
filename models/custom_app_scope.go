package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomAppScope 
type CustomAppScope struct {
    AppScope
}
// NewCustomAppScope instantiates a new CustomAppScope and sets the default values.
func NewCustomAppScope()(*CustomAppScope) {
    m := &CustomAppScope{
        AppScope: *NewAppScope(),
    }
    odataTypeValue := "#microsoft.graph.customAppScope"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomAppScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomAppScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomAppScope(), nil
}
// GetCustomAttributes gets the customAttributes property value. The customAttributes property
func (m *CustomAppScope) GetCustomAttributes()(CustomAppScopeAttributesDictionaryable) {
    val, err := m.GetBackingStore().Get("customAttributes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomAppScopeAttributesDictionaryable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomAppScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppScope.GetFieldDeserializers()
    res["customAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomAppScopeAttributesDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomAttributes(val.(CustomAppScopeAttributesDictionaryable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CustomAppScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppScope.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customAttributes", m.GetCustomAttributes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomAttributes sets the customAttributes property value. The customAttributes property
func (m *CustomAppScope) SetCustomAttributes(value CustomAppScopeAttributesDictionaryable)() {
    err := m.GetBackingStore().Set("customAttributes", value)
    if err != nil {
        panic(err)
    }
}
// CustomAppScopeable 
type CustomAppScopeable interface {
    AppScopeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomAttributes()(CustomAppScopeAttributesDictionaryable)
    SetCustomAttributes(value CustomAppScopeAttributesDictionaryable)()
}
