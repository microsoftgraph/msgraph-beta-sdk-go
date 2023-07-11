package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EnrollmentCompletionPageConfigurationPolicySetItem a class containing the properties used for Windows10EnrollmentCompletionPageConfiguration PolicySetItem.
type Windows10EnrollmentCompletionPageConfigurationPolicySetItem struct {
    PolicySetItem
}
// NewWindows10EnrollmentCompletionPageConfigurationPolicySetItem instantiates a new windows10EnrollmentCompletionPageConfigurationPolicySetItem and sets the default values.
func NewWindows10EnrollmentCompletionPageConfigurationPolicySetItem()(*Windows10EnrollmentCompletionPageConfigurationPolicySetItem) {
    m := &Windows10EnrollmentCompletionPageConfigurationPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.windows10EnrollmentCompletionPageConfigurationPolicySetItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10EnrollmentCompletionPageConfigurationPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EnrollmentCompletionPageConfigurationPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EnrollmentCompletionPageConfigurationPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10EnrollmentCompletionPageConfigurationPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. Priority of the Windows10EnrollmentCompletionPageConfigurationPolicySetItem.
func (m *Windows10EnrollmentCompletionPageConfigurationPolicySetItem) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Windows10EnrollmentCompletionPageConfigurationPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPriority sets the priority property value. Priority of the Windows10EnrollmentCompletionPageConfigurationPolicySetItem.
func (m *Windows10EnrollmentCompletionPageConfigurationPolicySetItem) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// Windows10EnrollmentCompletionPageConfigurationPolicySetItemable 
type Windows10EnrollmentCompletionPageConfigurationPolicySetItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicySetItemable
    GetPriority()(*int32)
    SetPriority(value *int32)()
}
