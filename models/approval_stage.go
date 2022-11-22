package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalStage 
type ApprovalStage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of days that a request can be pending a response before it is automatically denied.
    approvalStageTimeOutInDays *int32
    // If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.  When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection.
    escalationApprovers []UserSetable
    // If escalation is required, the time a request can be pending a response from a primary approver.
    escalationTimeInMinutes *int32
    // Indicates whether the approver is required to provide a justification for approving a request.
    isApproverJustificationRequired *bool
    // If true, then one or more escalation approvers are configured in this approval stage.
    isEscalationEnabled *bool
    // The OdataType property
    odataType *string
    // The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection.
    primaryApprovers []UserSetable
}
// NewApprovalStage instantiates a new approvalStage and sets the default values.
func NewApprovalStage()(*ApprovalStage) {
    m := &ApprovalStage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApprovalStageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalStageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalStage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalStage) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApprovalStageTimeOutInDays gets the approvalStageTimeOutInDays property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *ApprovalStage) GetApprovalStageTimeOutInDays()(*int32) {
    return m.approvalStageTimeOutInDays
}
// GetEscalationApprovers gets the escalationApprovers property value. If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.  When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection.
func (m *ApprovalStage) GetEscalationApprovers()([]UserSetable) {
    return m.escalationApprovers
}
// GetEscalationTimeInMinutes gets the escalationTimeInMinutes property value. If escalation is required, the time a request can be pending a response from a primary approver.
func (m *ApprovalStage) GetEscalationTimeInMinutes()(*int32) {
    return m.escalationTimeInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalStage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["approvalStageTimeOutInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetApprovalStageTimeOutInDays)
    res["escalationApprovers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserSetFromDiscriminatorValue , m.SetEscalationApprovers)
    res["escalationTimeInMinutes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetEscalationTimeInMinutes)
    res["isApproverJustificationRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsApproverJustificationRequired)
    res["isEscalationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEscalationEnabled)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["primaryApprovers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserSetFromDiscriminatorValue , m.SetPrimaryApprovers)
    return res
}
// GetIsApproverJustificationRequired gets the isApproverJustificationRequired property value. Indicates whether the approver is required to provide a justification for approving a request.
func (m *ApprovalStage) GetIsApproverJustificationRequired()(*bool) {
    return m.isApproverJustificationRequired
}
// GetIsEscalationEnabled gets the isEscalationEnabled property value. If true, then one or more escalation approvers are configured in this approval stage.
func (m *ApprovalStage) GetIsEscalationEnabled()(*bool) {
    return m.isEscalationEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ApprovalStage) GetOdataType()(*string) {
    return m.odataType
}
// GetPrimaryApprovers gets the primaryApprovers property value. The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection.
func (m *ApprovalStage) GetPrimaryApprovers()([]UserSetable) {
    return m.primaryApprovers
}
// Serialize serializes information the current object
func (m *ApprovalStage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("approvalStageTimeOutInDays", m.GetApprovalStageTimeOutInDays())
        if err != nil {
            return err
        }
    }
    if m.GetEscalationApprovers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEscalationApprovers())
        err := writer.WriteCollectionOfObjectValues("escalationApprovers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("escalationTimeInMinutes", m.GetEscalationTimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApproverJustificationRequired", m.GetIsApproverJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEscalationEnabled", m.GetIsEscalationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPrimaryApprovers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPrimaryApprovers())
        err := writer.WriteCollectionOfObjectValues("primaryApprovers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalStage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApprovalStageTimeOutInDays sets the approvalStageTimeOutInDays property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *ApprovalStage) SetApprovalStageTimeOutInDays(value *int32)() {
    m.approvalStageTimeOutInDays = value
}
// SetEscalationApprovers sets the escalationApprovers property value. If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.  When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection.
func (m *ApprovalStage) SetEscalationApprovers(value []UserSetable)() {
    m.escalationApprovers = value
}
// SetEscalationTimeInMinutes sets the escalationTimeInMinutes property value. If escalation is required, the time a request can be pending a response from a primary approver.
func (m *ApprovalStage) SetEscalationTimeInMinutes(value *int32)() {
    m.escalationTimeInMinutes = value
}
// SetIsApproverJustificationRequired sets the isApproverJustificationRequired property value. Indicates whether the approver is required to provide a justification for approving a request.
func (m *ApprovalStage) SetIsApproverJustificationRequired(value *bool)() {
    m.isApproverJustificationRequired = value
}
// SetIsEscalationEnabled sets the isEscalationEnabled property value. If true, then one or more escalation approvers are configured in this approval stage.
func (m *ApprovalStage) SetIsEscalationEnabled(value *bool)() {
    m.isEscalationEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ApprovalStage) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrimaryApprovers sets the primaryApprovers property value. The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection.
func (m *ApprovalStage) SetPrimaryApprovers(value []UserSetable)() {
    m.primaryApprovers = value
}
