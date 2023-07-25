package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectByTemplateAction 
type ProtectByTemplateAction struct {
    InformationProtectionAction
}
// NewProtectByTemplateAction instantiates a new protectByTemplateAction and sets the default values.
func NewProtectByTemplateAction()(*ProtectByTemplateAction) {
    m := &ProtectByTemplateAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.protectByTemplateAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProtectByTemplateActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectByTemplateActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectByTemplateAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetTemplateId gets the templateId property value. The unique identifier for a protection template in Microsoft Purview Information Protection to apply to the content.
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
// SetTemplateId sets the templateId property value. The unique identifier for a protection template in Microsoft Purview Information Protection to apply to the content.
func (m *ProtectByTemplateAction) SetTemplateId(value *string)() {
    err := m.GetBackingStore().Set("templateId", value)
    if err != nil {
        panic(err)
    }
}
// ProtectByTemplateActionable 
type ProtectByTemplateActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTemplateId()(*string)
    SetTemplateId(value *string)()
}
