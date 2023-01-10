package privilegedroles

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSelfActivatePostRequestBody 
type ItemSelfActivatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The duration property
    duration *string
    // The reason property
    reason *string
    // The ticketNumber property
    ticketNumber *string
    // The ticketSystem property
    ticketSystem *string
}
// NewItemSelfActivatePostRequestBody instantiates a new ItemSelfActivatePostRequestBody and sets the default values.
func NewItemSelfActivatePostRequestBody()(*ItemSelfActivatePostRequestBody) {
    m := &ItemSelfActivatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSelfActivatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSelfActivatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSelfActivatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSelfActivatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDuration gets the duration property value. The duration property
func (m *ItemSelfActivatePostRequestBody) GetDuration()(*string) {
    return m.duration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSelfActivatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
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
func (m *ItemSelfActivatePostRequestBody) GetReason()(*string) {
    return m.reason
}
// GetTicketNumber gets the ticketNumber property value. The ticketNumber property
func (m *ItemSelfActivatePostRequestBody) GetTicketNumber()(*string) {
    return m.ticketNumber
}
// GetTicketSystem gets the ticketSystem property value. The ticketSystem property
func (m *ItemSelfActivatePostRequestBody) GetTicketSystem()(*string) {
    return m.ticketSystem
}
// Serialize serializes information the current object
func (m *ItemSelfActivatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("duration", m.GetDuration())
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
func (m *ItemSelfActivatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDuration sets the duration property value. The duration property
func (m *ItemSelfActivatePostRequestBody) SetDuration(value *string)() {
    m.duration = value
}
// SetReason sets the reason property value. The reason property
func (m *ItemSelfActivatePostRequestBody) SetReason(value *string)() {
    m.reason = value
}
// SetTicketNumber sets the ticketNumber property value. The ticketNumber property
func (m *ItemSelfActivatePostRequestBody) SetTicketNumber(value *string)() {
    m.ticketNumber = value
}
// SetTicketSystem sets the ticketSystem property value. The ticketSystem property
func (m *ItemSelfActivatePostRequestBody) SetTicketSystem(value *string)() {
    m.ticketSystem = value
}
