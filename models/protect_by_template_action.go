package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProtectByTemplateAction struct {
    InformationProtectionAction
}
// NewProtectByTemplateAction instantiates a new ProtectByTemplateAction and sets the default values.
func NewProtectByTemplateAction()(*ProtectByTemplateAction) {
    m := &ProtectByTemplateAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.protectByTemplateAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProtectByTemplateActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectByTemplateActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectByTemplateAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectByTemplateAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    return res
}
// GetTemplateId gets the templateId property value. The GUID of the Azure Information Protection template to apply to the information.
// returns a *string when successful
func (m *ProtectByTemplateAction) GetTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("templateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProtectByTemplateAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTemplateId sets the templateId property value. The GUID of the Azure Information Protection template to apply to the information.
func (m *ProtectByTemplateAction) SetTemplateId(value *string)() {
    err := m.GetBackingStore().Set("templateId", value)
    if err != nil {
        panic(err)
    }
}
type ProtectByTemplateActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTemplateId()(*string)
    SetTemplateId(value *string)()
}
