package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentRestrictionsConfigurationPolicySetItem a class containing the properties used for enrollment restriction PolicySetItem.
type EnrollmentRestrictionsConfigurationPolicySetItem struct {
    PolicySetItem
}
// NewEnrollmentRestrictionsConfigurationPolicySetItem instantiates a new enrollmentRestrictionsConfigurationPolicySetItem and sets the default values.
func NewEnrollmentRestrictionsConfigurationPolicySetItem()(*EnrollmentRestrictionsConfigurationPolicySetItem) {
    m := &EnrollmentRestrictionsConfigurationPolicySetItem{
        PolicySetItem: *NewPolicySetItem(),
    }
    odataTypeValue := "#microsoft.graph.enrollmentRestrictionsConfigurationPolicySetItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEnrollmentRestrictionsConfigurationPolicySetItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollmentRestrictionsConfigurationPolicySetItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentRestrictionsConfigurationPolicySetItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicySetItem.GetFieldDeserializers()
    res["limit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimit(val)
        }
        return nil
    }
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
// GetLimit gets the limit property value. Limit of the EnrollmentRestrictionsConfigurationPolicySetItem.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) GetLimit()(*int32) {
    val, err := m.GetBackingStore().Get("limit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPriority gets the priority property value. Priority of the EnrollmentRestrictionsConfigurationPolicySetItem.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) GetPriority()(*int32) {
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
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicySetItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("limit", m.GetLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLimit sets the limit property value. Limit of the EnrollmentRestrictionsConfigurationPolicySetItem.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) SetLimit(value *int32)() {
    err := m.GetBackingStore().Set("limit", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. Priority of the EnrollmentRestrictionsConfigurationPolicySetItem.
func (m *EnrollmentRestrictionsConfigurationPolicySetItem) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// EnrollmentRestrictionsConfigurationPolicySetItemable 
type EnrollmentRestrictionsConfigurationPolicySetItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicySetItemable
    GetLimit()(*int32)
    GetPriority()(*int32)
    SetLimit(value *int32)()
    SetPriority(value *int32)()
}
