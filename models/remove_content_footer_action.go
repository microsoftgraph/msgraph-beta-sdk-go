package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoveContentFooterAction 
type RemoveContentFooterAction struct {
    InformationProtectionAction
    // The name of the UI element of the footer to be removed.
    uiElementNames []string
}
// NewRemoveContentFooterAction instantiates a new RemoveContentFooterAction and sets the default values.
func NewRemoveContentFooterAction()(*RemoveContentFooterAction) {
    m := &RemoveContentFooterAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.removeContentFooterAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRemoveContentFooterActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoveContentFooterActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoveContentFooterAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoveContentFooterAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["uiElementNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetUiElementNames)
    return res
}
// GetUiElementNames gets the uiElementNames property value. The name of the UI element of the footer to be removed.
func (m *RemoveContentFooterAction) GetUiElementNames()([]string) {
    return m.uiElementNames
}
// Serialize serializes information the current object
func (m *RemoveContentFooterAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUiElementNames() != nil {
        err = writer.WriteCollectionOfStringValues("uiElementNames", m.GetUiElementNames())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUiElementNames sets the uiElementNames property value. The name of the UI element of the footer to be removed.
func (m *RemoveContentFooterAction) SetUiElementNames(value []string)() {
    m.uiElementNames = value
}
