package updaterequest

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// UpdateRequestRequestBody provides operations to call the updateRequest method.
type UpdateRequestRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignmentState property
    assignmentState *string
    // The decision property
    decision *string
    // The reason property
    reason *string
    // The schedule property
    schedule ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceScheduleable
}
// NewUpdateRequestRequestBody instantiates a new updateRequestRequestBody and sets the default values.
func NewUpdateRequestRequestBody()(*UpdateRequestRequestBody) {
    m := &UpdateRequestRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateRequestRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateRequestRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateRequestRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateRequestRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignmentState gets the assignmentState property value. The assignmentState property
func (m *UpdateRequestRequestBody) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
// GetDecision gets the decision property value. The decision property
func (m *UpdateRequestRequestBody) GetDecision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decision
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateRequestRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceScheduleable))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
func (m *UpdateRequestRequestBody) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetSchedule gets the schedule property value. The schedule property
func (m *UpdateRequestRequestBody) GetSchedule()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// Serialize serializes information the current object
func (m *UpdateRequestRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UpdateRequestRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignmentState sets the assignmentState property value. The assignmentState property
func (m *UpdateRequestRequestBody) SetAssignmentState(value *string)() {
    if m != nil {
        m.assignmentState = value
    }
}
// SetDecision sets the decision property value. The decision property
func (m *UpdateRequestRequestBody) SetDecision(value *string)() {
    if m != nil {
        m.decision = value
    }
}
// SetReason sets the reason property value. The reason property
func (m *UpdateRequestRequestBody) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
// SetSchedule sets the schedule property value. The schedule property
func (m *UpdateRequestRequestBody) SetSchedule(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceScheduleable)() {
    if m != nil {
        m.schedule = value
    }
}
