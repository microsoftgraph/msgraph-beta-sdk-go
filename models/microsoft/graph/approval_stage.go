package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ApprovalStage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of days that a request can be pending a response before it is automatically denied.
    approvalStageTimeOutInDays *int32;
    // If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.
    escalationApprovers []UserSet;
    // If escalation is required, the time a request can be pending a response from a primary approver.
    escalationTimeInMinutes *int32;
    // Indicates whether the approver is required to provide a justification for approving a request.
    isApproverJustificationRequired *bool;
    // If true, then one or more escalation approvers are configured in this approval stage.
    isEscalationEnabled *bool;
    // The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.
    primaryApprovers []UserSet;
}
// Instantiates a new approvalStage and sets the default values.
func NewApprovalStage()(*ApprovalStage) {
    m := &ApprovalStage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalStage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the approvalStageTimeOutInDays property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *ApprovalStage) GetApprovalStageTimeOutInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.approvalStageTimeOutInDays
    }
}
// Gets the escalationApprovers property value. If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.
func (m *ApprovalStage) GetEscalationApprovers()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.escalationApprovers
    }
}
// Gets the escalationTimeInMinutes property value. If escalation is required, the time a request can be pending a response from a primary approver.
func (m *ApprovalStage) GetEscalationTimeInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.escalationTimeInMinutes
    }
}
// Gets the isApproverJustificationRequired property value. Indicates whether the approver is required to provide a justification for approving a request.
func (m *ApprovalStage) GetIsApproverJustificationRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApproverJustificationRequired
    }
}
// Gets the isEscalationEnabled property value. If true, then one or more escalation approvers are configured in this approval stage.
func (m *ApprovalStage) GetIsEscalationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEscalationEnabled
    }
}
// Gets the primaryApprovers property value. The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.
func (m *ApprovalStage) GetPrimaryApprovers()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.primaryApprovers
    }
}
// The deserialization information for the current model
func (m *ApprovalStage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approvalStageTimeOutInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetApprovalStageTimeOutInDays(val)
        return nil
    }
    res["escalationApprovers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        res := make([]UserSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSet))
        }
        m.SetEscalationApprovers(res)
        return nil
    }
    res["escalationTimeInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEscalationTimeInMinutes(val)
        return nil
    }
    res["isApproverJustificationRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsApproverJustificationRequired(val)
        return nil
    }
    res["isEscalationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEscalationEnabled(val)
        return nil
    }
    res["primaryApprovers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        res := make([]UserSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSet))
        }
        m.SetPrimaryApprovers(res)
        return nil
    }
    return res
}
func (m *ApprovalStage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ApprovalStage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("approvalStageTimeOutInDays", m.GetApprovalStageTimeOutInDays())
        if err != nil {
            return err
        }
    }
    {
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
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ApprovalStage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the approvalStageTimeOutInDays property value. The number of days that a request can be pending a response before it is automatically denied.
// Parameters:
//  - value : Value to set for the approvalStageTimeOutInDays property.
func (m *ApprovalStage) SetApprovalStageTimeOutInDays(value *int32)() {
    m.approvalStageTimeOutInDays = value
}
// Sets the escalationApprovers property value. If escalation is enabled and the primary approvers do not respond before the escalation time, the escalationApprovers are the users who will be asked to approve requests. This can be a collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.
// Parameters:
//  - value : Value to set for the escalationApprovers property.
func (m *ApprovalStage) SetEscalationApprovers(value []UserSet)() {
    m.escalationApprovers = value
}
// Sets the escalationTimeInMinutes property value. If escalation is required, the time a request can be pending a response from a primary approver.
// Parameters:
//  - value : Value to set for the escalationTimeInMinutes property.
func (m *ApprovalStage) SetEscalationTimeInMinutes(value *int32)() {
    m.escalationTimeInMinutes = value
}
// Sets the isApproverJustificationRequired property value. Indicates whether the approver is required to provide a justification for approving a request.
// Parameters:
//  - value : Value to set for the isApproverJustificationRequired property.
func (m *ApprovalStage) SetIsApproverJustificationRequired(value *bool)() {
    m.isApproverJustificationRequired = value
}
// Sets the isEscalationEnabled property value. If true, then one or more escalation approvers are configured in this approval stage.
// Parameters:
//  - value : Value to set for the isEscalationEnabled property.
func (m *ApprovalStage) SetIsEscalationEnabled(value *bool)() {
    m.isEscalationEnabled = value
}
// Sets the primaryApprovers property value. The users who will be asked to approve requests. A collection of singleUser, groupMembers, requestorManager, internalSponsors and externalSponsors.
// Parameters:
//  - value : Value to set for the primaryApprovers property.
func (m *ApprovalStage) SetPrimaryApprovers(value []UserSet)() {
    m.primaryApprovers = value
}
