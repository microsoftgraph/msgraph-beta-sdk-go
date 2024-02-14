package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationCheckBox represents an ADMX checkBox element and an ADMX boolean element.
type GroupPolicyPresentationCheckBox struct {
    GroupPolicyUploadedPresentation
}
// NewGroupPolicyPresentationCheckBox instantiates a new GroupPolicyPresentationCheckBox and sets the default values.
func NewGroupPolicyPresentationCheckBox()(*GroupPolicyPresentationCheckBox) {
    m := &GroupPolicyPresentationCheckBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationCheckBox"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupPolicyPresentationCheckBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupPolicyPresentationCheckBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationCheckBox(), nil
}
// GetDefaultChecked gets the defaultChecked property value. Default value for the check box. The default value is false.
// returns a *bool when successful
func (m *GroupPolicyPresentationCheckBox) GetDefaultChecked()(*bool) {
    val, err := m.GetBackingStore().Get("defaultChecked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupPolicyPresentationCheckBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
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
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("defaultChecked", m.GetDefaultChecked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultChecked sets the defaultChecked property value. Default value for the check box. The default value is false.
func (m *GroupPolicyPresentationCheckBox) SetDefaultChecked(value *bool)() {
    err := m.GetBackingStore().Set("defaultChecked", value)
    if err != nil {
        panic(err)
    }
}
type GroupPolicyPresentationCheckBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultChecked()(*bool)
    SetDefaultChecked(value *bool)()
}
