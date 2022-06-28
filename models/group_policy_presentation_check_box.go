package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationCheckBox 
type GroupPolicyPresentationCheckBox struct {
    GroupPolicyPresentation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Default value for the check box. The default value is false.
    defaultChecked *bool
}
// NewGroupPolicyPresentationCheckBox instantiates a new GroupPolicyPresentationCheckBox and sets the default values.
func NewGroupPolicyPresentationCheckBox()(*GroupPolicyPresentationCheckBox) {
    m := &GroupPolicyPresentationCheckBox{
        GroupPolicyPresentation: *NewGroupPolicyPresentation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupPolicyPresentationCheckBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationCheckBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationCheckBox(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyPresentationCheckBox) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultChecked gets the defaultChecked property value. Default value for the check box. The default value is false.
func (m *GroupPolicyPresentationCheckBox) GetDefaultChecked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defaultChecked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationCheckBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyPresentation.GetFieldDeserializers()
    res["defaultChecked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultChecked(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationCheckBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("defaultChecked", m.GetDefaultChecked())
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
func (m *GroupPolicyPresentationCheckBox) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultChecked sets the defaultChecked property value. Default value for the check box. The default value is false.
func (m *GroupPolicyPresentationCheckBox) SetDefaultChecked(value *bool)() {
    if m != nil {
        m.defaultChecked = value
    }
}
