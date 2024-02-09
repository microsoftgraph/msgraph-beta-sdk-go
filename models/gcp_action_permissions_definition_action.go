package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GcpActionPermissionsDefinitionAction struct {
    GcpPermissionsDefinitionAction
}
// NewGcpActionPermissionsDefinitionAction instantiates a new GcpActionPermissionsDefinitionAction and sets the default values.
func NewGcpActionPermissionsDefinitionAction()(*GcpActionPermissionsDefinitionAction) {
    m := &GcpActionPermissionsDefinitionAction{
        GcpPermissionsDefinitionAction: *NewGcpPermissionsDefinitionAction(),
    }
    odataTypeValue := "#microsoft.graph.gcpActionPermissionsDefinitionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGcpActionPermissionsDefinitionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGcpActionPermissionsDefinitionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpActionPermissionsDefinitionAction(), nil
}
// GetActions gets the actions property value. List of actions.
// returns a []string when successful
func (m *GcpActionPermissionsDefinitionAction) GetActions()([]string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GcpActionPermissionsDefinitionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GcpPermissionsDefinitionAction.GetFieldDeserializers()
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
func (m *GcpActionPermissionsDefinitionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GcpPermissionsDefinitionAction.Serialize(writer)
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
// SetActions sets the actions property value. List of actions.
func (m *GcpActionPermissionsDefinitionAction) SetActions(value []string)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
type GcpActionPermissionsDefinitionActionable interface {
    GcpPermissionsDefinitionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]string)
    SetActions(value []string)()
}
