package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SingleResourceGcpPermissionsDefinition 
type SingleResourceGcpPermissionsDefinition struct {
    PermissionsDefinition
}
// NewSingleResourceGcpPermissionsDefinition instantiates a new singleResourceGcpPermissionsDefinition and sets the default values.
func NewSingleResourceGcpPermissionsDefinition()(*SingleResourceGcpPermissionsDefinition) {
    m := &SingleResourceGcpPermissionsDefinition{
        PermissionsDefinition: *NewPermissionsDefinition(),
    }
    odataTypeValue := "#microsoft.graph.singleResourceGcpPermissionsDefinition"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSingleResourceGcpPermissionsDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSingleResourceGcpPermissionsDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSingleResourceGcpPermissionsDefinition(), nil
}
// GetActionInfo gets the actionInfo property value. The actionInfo property
func (m *SingleResourceGcpPermissionsDefinition) GetActionInfo()(GcpPermissionsDefinitionActionable) {
    val, err := m.GetBackingStore().Get("actionInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GcpPermissionsDefinitionActionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SingleResourceGcpPermissionsDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinition.GetFieldDeserializers()
    res["actionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGcpPermissionsDefinitionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionInfo(val.(GcpPermissionsDefinitionActionable))
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
// GetResourceId gets the resourceId property value. Identifier for the resource.
func (m *SingleResourceGcpPermissionsDefinition) GetResourceId()(*string) {
    val, err := m.GetBackingStore().Get("resourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SingleResourceGcpPermissionsDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PermissionsDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("actionInfo", m.GetActionInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionInfo sets the actionInfo property value. The actionInfo property
func (m *SingleResourceGcpPermissionsDefinition) SetActionInfo(value GcpPermissionsDefinitionActionable)() {
    err := m.GetBackingStore().Set("actionInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceId sets the resourceId property value. Identifier for the resource.
func (m *SingleResourceGcpPermissionsDefinition) SetResourceId(value *string)() {
    err := m.GetBackingStore().Set("resourceId", value)
    if err != nil {
        panic(err)
    }
}
// SingleResourceGcpPermissionsDefinitionable 
type SingleResourceGcpPermissionsDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionable
    GetActionInfo()(GcpPermissionsDefinitionActionable)
    GetResourceId()(*string)
    SetActionInfo(value GcpPermissionsDefinitionActionable)()
    SetResourceId(value *string)()
}
