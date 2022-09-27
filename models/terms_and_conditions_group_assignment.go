package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TermsAndConditionsGroupAssignment a termsAndConditionsGroupAssignment entity represents the assignment of a given Terms and Conditions (T&C) policy to a given group. Users in the group will be required to accept the terms in order to have devices enrolled into Intune.
type TermsAndConditionsGroupAssignment struct {
    Entity
    // Unique identifier of a group that the T&C policy is assigned to.
    targetGroupId *string
    // Navigation link to the terms and conditions that are assigned.
    termsAndConditions TermsAndConditionsable
}
// NewTermsAndConditionsGroupAssignment instantiates a new termsAndConditionsGroupAssignment and sets the default values.
func NewTermsAndConditionsGroupAssignment()(*TermsAndConditionsGroupAssignment) {
    m := &TermsAndConditionsGroupAssignment{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.termsAndConditionsGroupAssignment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTermsAndConditionsGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTermsAndConditionsGroupAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTermsAndConditionsGroupAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermsAndConditionsGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetGroupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTargetGroupId)
    res["termsAndConditions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTermsAndConditionsFromDiscriminatorValue , m.SetTermsAndConditions)
    return res
}
// GetTargetGroupId gets the targetGroupId property value. Unique identifier of a group that the T&C policy is assigned to.
func (m *TermsAndConditionsGroupAssignment) GetTargetGroupId()(*string) {
    return m.targetGroupId
}
// GetTermsAndConditions gets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsGroupAssignment) GetTermsAndConditions()(TermsAndConditionsable) {
    return m.termsAndConditions
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
    m.targetGroupId = value
}
// SetTermsAndConditions sets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsGroupAssignment) SetTermsAndConditions(value TermsAndConditionsable)() {
    m.termsAndConditions = value
}
