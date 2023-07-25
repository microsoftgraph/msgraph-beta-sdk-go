package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoveProtectionAction 
type RemoveProtectionAction struct {
    InformationProtectionAction
}
// NewRemoveProtectionAction instantiates a new removeProtectionAction and sets the default values.
func NewRemoveProtectionAction()(*RemoveProtectionAction) {
    m := &RemoveProtectionAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.removeProtectionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRemoveProtectionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoveProtectionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoveProtectionAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoveProtectionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RemoveProtectionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RemoveProtectionActionable 
type RemoveProtectionActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
