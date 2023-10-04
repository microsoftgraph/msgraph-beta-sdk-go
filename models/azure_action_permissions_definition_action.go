package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureActionPermissionsDefinitionAction 
type AzureActionPermissionsDefinitionAction struct {
    AzurePermissionsDefinitionAction
}
// NewAzureActionPermissionsDefinitionAction instantiates a new azureActionPermissionsDefinitionAction and sets the default values.
func NewAzureActionPermissionsDefinitionAction()(*AzureActionPermissionsDefinitionAction) {
    m := &AzureActionPermissionsDefinitionAction{
        AzurePermissionsDefinitionAction: *NewAzurePermissionsDefinitionAction(),
    }
    odataTypeValue := "#microsoft.graph.azureActionPermissionsDefinitionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureActionPermissionsDefinitionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureActionPermissionsDefinitionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureActionPermissionsDefinitionAction(), nil
}
// GetActions gets the actions property value. The actions property
func (m *AzureActionPermissionsDefinitionAction) GetActions()([]string) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureActionPermissionsDefinitionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AzurePermissionsDefinitionAction.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AzureActionPermissionsDefinitionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AzurePermissionsDefinitionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActions() != nil {
        err = writer.WriteCollectionOfStringValues("actions", m.GetActions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The actions property
func (m *AzureActionPermissionsDefinitionAction) SetActions(value []string)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// AzureActionPermissionsDefinitionActionable 
type AzureActionPermissionsDefinitionActionable interface {
    AzurePermissionsDefinitionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]string)
    SetActions(value []string)()
}
