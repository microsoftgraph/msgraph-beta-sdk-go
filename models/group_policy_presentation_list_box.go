package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationListBox represents an ADMX listBox element and an ADMX list element.
type GroupPolicyPresentationListBox struct {
    GroupPolicyUploadedPresentation
}
// NewGroupPolicyPresentationListBox instantiates a new groupPolicyPresentationListBox and sets the default values.
func NewGroupPolicyPresentationListBox()(*GroupPolicyPresentationListBox) {
    m := &GroupPolicyPresentationListBox{
        GroupPolicyUploadedPresentation: *NewGroupPolicyUploadedPresentation(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyPresentationListBox"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupPolicyPresentationListBoxFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyPresentationListBoxFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyPresentationListBox(), nil
}
// GetExplicitValue gets the explicitValue property value. If this option is specified true the user must specify the registry subkey value and the registry subkey name. The list box shows two columns, one for the name and one for the data. The default value is false.
func (m *GroupPolicyPresentationListBox) GetExplicitValue()(*bool) {
    val, err := m.GetBackingStore().Get("explicitValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationListBox) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyUploadedPresentation.GetFieldDeserializers()
    res["explicitValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExplicitValue(val)
        }
        return nil
    }
    res["valuePrefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValuePrefix(val)
        }
        return nil
    }
    return res
}
// GetValuePrefix gets the valuePrefix property value. The valuePrefix property
func (m *GroupPolicyPresentationListBox) GetValuePrefix()(*string) {
    val, err := m.GetBackingStore().Get("valuePrefix")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationListBox) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GroupPolicyUploadedPresentation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("explicitValue", m.GetExplicitValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valuePrefix", m.GetValuePrefix())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExplicitValue sets the explicitValue property value. If this option is specified true the user must specify the registry subkey value and the registry subkey name. The list box shows two columns, one for the name and one for the data. The default value is false.
func (m *GroupPolicyPresentationListBox) SetExplicitValue(value *bool)() {
    err := m.GetBackingStore().Set("explicitValue", value)
    if err != nil {
        panic(err)
    }
}
// SetValuePrefix sets the valuePrefix property value. The valuePrefix property
func (m *GroupPolicyPresentationListBox) SetValuePrefix(value *string)() {
    err := m.GetBackingStore().Set("valuePrefix", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyPresentationListBoxable 
type GroupPolicyPresentationListBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExplicitValue()(*bool)
    GetValuePrefix()(*string)
    SetExplicitValue(value *bool)()
    SetValuePrefix(value *string)()
}
