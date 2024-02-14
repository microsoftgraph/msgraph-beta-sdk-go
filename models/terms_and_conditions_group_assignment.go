package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TermsAndConditionsGroupAssignment a termsAndConditionsGroupAssignment entity represents the assignment of a given Terms and Conditions (T&C) policy to a given group. Users in the group will be required to accept the terms in order to have devices enrolled into Intune.
type TermsAndConditionsGroupAssignment struct {
    Entity
}
// NewTermsAndConditionsGroupAssignment instantiates a new TermsAndConditionsGroupAssignment and sets the default values.
func NewTermsAndConditionsGroupAssignment()(*TermsAndConditionsGroupAssignment) {
    m := &TermsAndConditionsGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTermsAndConditionsGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTermsAndConditionsGroupAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTermsAndConditionsGroupAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TermsAndConditionsGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    res["termsAndConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermsAndConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsAndConditions(val.(TermsAndConditionsable))
        }
        return nil
    }
    return res
}
// GetTargetGroupId gets the targetGroupId property value. Unique identifier of a group that the T&C policy is assigned to.
// returns a *string when successful
func (m *TermsAndConditionsGroupAssignment) GetTargetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("targetGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTermsAndConditions gets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
// returns a TermsAndConditionsable when successful
func (m *TermsAndConditionsGroupAssignment) GetTermsAndConditions()(TermsAndConditionsable) {
    val, err := m.GetBackingStore().Get("termsAndConditions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TermsAndConditionsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TermsAndConditionsGroupAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("termsAndConditions", m.GetTermsAndConditions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetGroupId sets the targetGroupId property value. Unique identifier of a group that the T&C policy is assigned to.
func (m *TermsAndConditionsGroupAssignment) SetTargetGroupId(value *string)() {
    err := m.GetBackingStore().Set("targetGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetTermsAndConditions sets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsGroupAssignment) SetTermsAndConditions(value TermsAndConditionsable)() {
    err := m.GetBackingStore().Set("termsAndConditions", value)
    if err != nil {
        panic(err)
    }
}
type TermsAndConditionsGroupAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTargetGroupId()(*string)
    GetTermsAndConditions()(TermsAndConditionsable)
    SetTargetGroupId(value *string)()
    SetTermsAndConditions(value TermsAndConditionsable)()
}
