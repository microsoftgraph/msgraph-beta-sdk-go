package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody provides operations to call the makePermanent method.
type PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The reason property
    reason *string
    // The ticketNumber property
    ticketNumber *string
    // The ticketSystem property
    ticketSystem *string
}
// NewPrivilegedRoleAssignmentsItemMakePermanentPostRequestBody instantiates a new PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody and sets the default values.
func NewPrivilegedRoleAssignmentsItemMakePermanentPostRequestBody()(*PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) {
    m := &PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrivilegedRoleAssignmentsItemMakePermanentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleAssignmentsItemMakePermanentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRoleAssignmentsItemMakePermanentPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["ticketNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketNumber(val)
        }
        return nil
    }
    res["ticketSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketSystem(val)
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) GetReason()(*string) {
    return m.reason
}
// GetTicketNumber gets the ticketNumber property value. The ticketNumber property
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) GetTicketNumber()(*string) {
    return m.ticketNumber
}
// GetTicketSystem gets the ticketSystem property value. The ticketSystem property
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) GetTicketSystem()(*string) {
    return m.ticketSystem
}
// Serialize serializes information the current object
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ticketNumber", m.GetTicketNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ticketSystem", m.GetTicketSystem())
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
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetReason sets the reason property value. The reason property
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) SetReason(value *string)() {
    m.reason = value
}
// SetTicketNumber sets the ticketNumber property value. The ticketNumber property
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) SetTicketNumber(value *string)() {
    m.ticketNumber = value
}
// SetTicketSystem sets the ticketSystem property value. The ticketSystem property
func (m *PrivilegedRoleAssignmentsItemMakePermanentPostRequestBody) SetTicketSystem(value *string)() {
    m.ticketSystem = value
}
