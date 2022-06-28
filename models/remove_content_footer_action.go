package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoveContentFooterAction 
type RemoveContentFooterAction struct {
    InformationProtectionAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the UI element of the footer to be removed.
    uiElementNames []string
}
// NewRemoveContentFooterAction instantiates a new RemoveContentFooterAction and sets the default values.
func NewRemoveContentFooterAction()(*RemoveContentFooterAction) {
    m := &RemoveContentFooterAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRemoveContentFooterActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoveContentFooterActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoveContentFooterAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoveContentFooterAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoveContentFooterAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["uiElementNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUiElementNames(res)
        }
        return nil
    }
    return res
}
// GetUiElementNames gets the uiElementNames property value. The name of the UI element of the footer to be removed.
func (m *RemoveContentFooterAction) GetUiElementNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.uiElementNames
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoveContentFooterAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUiElementNames sets the uiElementNames property value. The name of the UI element of the footer to be removed.
func (m *RemoveContentFooterAction) SetUiElementNames(value []string)() {
    if m != nil {
        m.uiElementNames = value
    }
}
