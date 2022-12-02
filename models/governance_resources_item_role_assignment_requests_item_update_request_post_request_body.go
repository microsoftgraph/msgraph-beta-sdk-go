package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody provides operations to call the updateRequest method.
type GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignmentState property
    assignmentState *string
    // The decision property
    decision *string
    // The reason property
    reason *string
    // The schedule property
    schedule GovernanceScheduleable
}
// NewGovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody instantiates a new GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody and sets the default values.
func NewGovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody()(*GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) {
    m := &GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentState gets the assignmentState property value. The assignmentState property
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) GetAssignmentState()(*string) {
    return m.assignmentState
}
// GetDecision gets the decision property value. The decision property
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) GetDecision()(*string) {
    return m.decision
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentState(val)
        }
        return nil
    }
    res["decision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecision(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(GovernanceScheduleable))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) GetReason()(*string) {
    return m.reason
}
// GetSchedule gets the schedule property value. The schedule property
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) GetSchedule()(GovernanceScheduleable) {
    return m.schedule
}
// Serialize serializes information the current object
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignmentState", m.GetAssignmentState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("decision", m.GetDecision())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("schedule", m.GetSchedule())
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
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentState sets the assignmentState property value. The assignmentState property
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
// SetDecision sets the decision property value. The decision property
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) SetDecision(value *string)() {
    m.decision = value
}
// SetReason sets the reason property value. The reason property
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) SetReason(value *string)() {
    m.reason = value
}
// SetSchedule sets the schedule property value. The schedule property
func (m *GovernanceResourcesItemRoleAssignmentRequestsItemUpdateRequestPostRequestBody) SetSchedule(value GovernanceScheduleable)() {
    m.schedule = value
}
