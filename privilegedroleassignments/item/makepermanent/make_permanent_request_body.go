package makepermanent

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new makePermanentRequestBody and sets the default values.
func NewMakePermanentRequestBody()(*MakePermanentRequestBody) {
    m := &MakePermanentRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MakePermanentRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the reason property value. 
func (m *MakePermanentRequestBody) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// Gets the ticketNumber property value. 
func (m *MakePermanentRequestBody) GetTicketNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ticketNumber
    }
}
// Gets the ticketSystem property value. 
func (m *MakePermanentRequestBody) GetTicketSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ticketSystem
    }
}
// The deserialization information for the current model
func (m *MakePermanentRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReason(val)
        return nil
    }
    res["ticketNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTicketNumber(val)
        return nil
    }
    res["ticketSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTicketSystem(val)
        return nil
    }
    return res
}
func (m *MakePermanentRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *MakePermanentRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the reason property value. 
// Parameters:
//  - value : Value to set for the reason property.
func (m *MakePermanentRequestBody) SetReason(value *string)() {
    m.reason = value
}
// Sets the ticketNumber property value. 
// Parameters:
//  - value : Value to set for the ticketNumber property.
func (m *MakePermanentRequestBody) SetTicketNumber(value *string)() {
    m.ticketNumber = value
}
// Sets the ticketSystem property value. 
// Parameters:
//  - value : Value to set for the ticketSystem property.
func (m *MakePermanentRequestBody) SetTicketSystem(value *string)() {
    m.ticketSystem = value
}
