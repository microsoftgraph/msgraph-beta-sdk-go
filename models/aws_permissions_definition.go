package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsPermissionsDefinition 
type AwsPermissionsDefinition struct {
    PermissionsDefinition
}
// NewAwsPermissionsDefinition instantiates a new awsPermissionsDefinition and sets the default values.
func NewAwsPermissionsDefinition()(*AwsPermissionsDefinition) {
    m := &AwsPermissionsDefinition{
        PermissionsDefinition: *NewPermissionsDefinition(),
    }
    odataTypeValue := "#microsoft.graph.awsPermissionsDefinition"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsPermissionsDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsPermissionsDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsPermissionsDefinition(), nil
}
// GetActionInfo gets the actionInfo property value. The actionInfo property
func (m *AwsPermissionsDefinition) GetActionInfo()(AwsPermissionsDefinitionActionable) {
    val, err := m.GetBackingStore().Get("actionInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsPermissionsDefinitionActionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsPermissionsDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinition.GetFieldDeserializers()
    res["actionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsPermissionsDefinitionActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionInfo(val.(AwsPermissionsDefinitionActionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AwsPermissionsDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetActionInfo sets the actionInfo property value. The actionInfo property
func (m *AwsPermissionsDefinition) SetActionInfo(value AwsPermissionsDefinitionActionable)() {
    err := m.GetBackingStore().Set("actionInfo", value)
    if err != nil {
        panic(err)
    }
}
// AwsPermissionsDefinitionable 
type AwsPermissionsDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionable
    GetActionInfo()(AwsPermissionsDefinitionActionable)
    SetActionInfo(value AwsPermissionsDefinitionActionable)()
}
