package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SingleResourceAzurePermissionsDefinition struct {
    PermissionsDefinition
}
// NewSingleResourceAzurePermissionsDefinition instantiates a new SingleResourceAzurePermissionsDefinition and sets the default values.
func NewSingleResourceAzurePermissionsDefinition()(*SingleResourceAzurePermissionsDefinition) {
    m := &SingleResourceAzurePermissionsDefinition{
        PermissionsDefinition: *NewPermissionsDefinition(),
    }
    odataTypeValue := "#microsoft.graph.singleResourceAzurePermissionsDefinition"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSingleResourceAzurePermissionsDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSingleResourceAzurePermissionsDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSingleResourceAzurePermissionsDefinition(), nil
}
// GetActionInfo gets the actionInfo property value. The actionInfo property
// returns a AzurePermissionsDefinitionActionable when successful
func (m *SingleResourceAzurePermissionsDefinition) GetActionInfo()(AzurePermissionsDefinitionActionable) {
    val, err := m.GetBackingStore().Get("actionInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AzurePermissionsDefinitionActionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SingleResourceAzurePermissionsDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinition.GetFieldDeserializers()
    res["actionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzurePermissionsDefinitionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionInfo(val.(AzurePermissionsDefinitionActionable))
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
// returns a *string when successful
func (m *SingleResourceAzurePermissionsDefinition) GetResourceId()(*string) {
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
func (m *SingleResourceAzurePermissionsDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SingleResourceAzurePermissionsDefinition) SetActionInfo(value AzurePermissionsDefinitionActionable)() {
    err := m.GetBackingStore().Set("actionInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceId sets the resourceId property value. Identifier for the resource.
func (m *SingleResourceAzurePermissionsDefinition) SetResourceId(value *string)() {
    err := m.GetBackingStore().Set("resourceId", value)
    if err != nil {
        panic(err)
    }
}
type SingleResourceAzurePermissionsDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionable
    GetActionInfo()(AzurePermissionsDefinitionActionable)
    GetResourceId()(*string)
    SetActionInfo(value AzurePermissionsDefinitionActionable)()
    SetResourceId(value *string)()
}
