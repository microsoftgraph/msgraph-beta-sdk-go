package makepermanent

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MakePermanentRequestBody 
type MakePermanentRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    reason *string;
    // 
    ticketNumber *string;
    // 
    ticketSystem *string;
}
// NewMakePermanentRequestBody instantiates a new makePermanentRequestBody and sets the default values.
func NewMakePermanentRequestBody()(*MakePermanentRequestBody) {
    m := &MakePermanentRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MakePermanentRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetReason gets the reason property value. 
func (m *MakePermanentRequestBody) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetTicketNumber gets the ticketNumber property value. 
func (m *MakePermanentRequestBody) GetTicketNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ticketNumber
    }
}
// GetTicketSystem gets the ticketSystem property value. 
func (m *MakePermanentRequestBody) GetTicketSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ticketSystem
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MakePermanentRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["ticketNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketNumber(val)
        }
        return nil
    }
    res["ticketSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *MakePermanentRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MakePermanentRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *MakePermanentRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReason sets the reason property value. 
func (m *MakePermanentRequestBody) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
// SetTicketNumber sets the ticketNumber property value. 
func (m *MakePermanentRequestBody) SetTicketNumber(value *string)() {
    if m != nil {
        m.ticketNumber = value
    }
}
// SetTicketSystem sets the ticketSystem property value. 
func (m *MakePermanentRequestBody) SetTicketSystem(value *string)() {
    if m != nil {
        m.ticketSystem = value
    }
}
