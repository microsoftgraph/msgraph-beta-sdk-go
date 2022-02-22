package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApprovalStage 
type ApprovalStage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of days that a request can be pending a response before it is automatically denied.
    approvalStageTimeOutInDays *int32;
    // If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.  When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection.
    escalationApprovers []UserSet;
    // If escalation is required, the time a request can be pending a response from a primary approver.
    escalationTimeInMinutes *int32;
    // Indicates whether the approver is required to provide a justification for approving a request.
    isApproverJustificationRequired *bool;
    // If true, then one or more escalation approvers are configured in this approval stage.
    isEscalationEnabled *bool;
    // The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection.
    primaryApprovers []UserSet;
}
// NewApprovalStage instantiates a new approvalStage and sets the default values.
func NewApprovalStage()(*ApprovalStage) {
    m := &ApprovalStage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalStage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApprovalStageTimeOutInDays gets the approvalStageTimeOutInDays property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *ApprovalStage) GetApprovalStageTimeOutInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.approvalStageTimeOutInDays
    }
}
// GetEscalationApprovers gets the escalationApprovers property value. If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.  When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection.
func (m *ApprovalStage) GetEscalationApprovers()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.escalationApprovers
    }
}
// GetEscalationTimeInMinutes gets the escalationTimeInMinutes property value. If escalation is required, the time a request can be pending a response from a primary approver.
func (m *ApprovalStage) GetEscalationTimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.escalationTimeInMinutes
    }
}
// GetIsApproverJustificationRequired gets the isApproverJustificationRequired property value. Indicates whether the approver is required to provide a justification for approving a request.
func (m *ApprovalStage) GetIsApproverJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApproverJustificationRequired
    }
}
// GetIsEscalationEnabled gets the isEscalationEnabled property value. If true, then one or more escalation approvers are configured in this approval stage.
func (m *ApprovalStage) GetIsEscalationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEscalationEnabled
    }
}
// GetPrimaryApprovers gets the primaryApprovers property value. The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection.
func (m *ApprovalStage) GetPrimaryApprovers()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.primaryApprovers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalStage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approvalStageTimeOutInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalStageTimeOutInDays(val)
        }
        return nil
    }
    res["escalationApprovers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSet, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserSet))
            }
            m.SetEscalationApprovers(res)
        }
        return nil
    }
    res["escalationTimeInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEscalationTimeInMinutes(val)
        }
        return nil
    }
    res["isApproverJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApproverJustificationRequired(val)
        }
        return nil
    }
    res["isEscalationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEscalationEnabled(val)
        }
        return nil
    }
    res["primaryApprovers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSet, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserSet))
            }
            m.SetPrimaryApprovers(res)
        }
        return nil
    }
    return res
}
func (m *ApprovalStage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApprovalStage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("approvalStageTimeOutInDays", m.GetApprovalStageTimeOutInDays())
        if err != nil {
            return err
        }
    }
    if m.GetEscalationApprovers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEscalationApprovers()))
        for i, v := range m.GetEscalationApprovers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    if m.GetPrimaryApprovers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPrimaryApprovers()))
        for i, v := range m.GetPrimaryApprovers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    if m != nil {
        m.additionalData = value
    }
}
// SetApprovalStageTimeOutInDays sets the approvalStageTimeOutInDays property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *ApprovalStage) SetApprovalStageTimeOutInDays(value *int32)() {
    if m != nil {
        m.approvalStageTimeOutInDays = value
    }
}
// SetEscalationApprovers sets the escalationApprovers property value. If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.  When creating or updating a policy, if there are no escalation approvers, or escalation approvers are not required for the stage, the value of this property should be an empty collection.
func (m *ApprovalStage) SetEscalationApprovers(value []UserSet)() {
    if m != nil {
        m.escalationApprovers = value
    }
}
// SetEscalationTimeInMinutes sets the escalationTimeInMinutes property value. If escalation is required, the time a request can be pending a response from a primary approver.
func (m *ApprovalStage) SetEscalationTimeInMinutes(value *int32)() {
    if m != nil {
        m.escalationTimeInMinutes = value
    }
}
// SetIsApproverJustificationRequired sets the isApproverJustificationRequired property value. Indicates whether the approver is required to provide a justification for approving a request.
func (m *ApprovalStage) SetIsApproverJustificationRequired(value *bool)() {
    if m != nil {
        m.isApproverJustificationRequired = value
    }
}
// SetIsEscalationEnabled sets the isEscalationEnabled property value. If true, then one or more escalation approvers are configured in this approval stage.
func (m *ApprovalStage) SetIsEscalationEnabled(value *bool)() {
    if m != nil {
        m.isEscalationEnabled = value
    }
}
// SetPrimaryApprovers sets the primaryApprovers property value. The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors. When creating or updating a policy, include at least one userSet in this collection.
func (m *ApprovalStage) SetPrimaryApprovers(value []UserSet)() {
    if m != nil {
        m.primaryApprovers = value
    }
}
