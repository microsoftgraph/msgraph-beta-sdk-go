package selfactivate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SelfActivateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    duration *string;
    // 
    reason *string;
    // 
    ticketNumber *string;
    // 
    ticketSystem *string;
}
// Instantiates a new selfActivateRequestBody and sets the default values.
func NewSelfActivateRequestBody()(*SelfActivateRequestBody) {
    m := &SelfActivateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfActivateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the duration property value. 
func (m *SelfActivateRequestBody) GetDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// Gets the reason property value. 
func (m *SelfActivateRequestBody) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// Gets the ticketNumber property value. 
func (m *SelfActivateRequestBody) GetTicketNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ticketNumber
    }
}
// Gets the ticketSystem property value. 
func (m *SelfActivateRequestBody) GetTicketSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ticketSystem
    }
}
// The deserialization information for the current model
func (m *SelfActivateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDuration(val)
        return nil
    }
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
func (m *SelfActivateRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SelfActivateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SelfActivateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the duration property value. 
// Parameters:
//  - value : Value to set for the duration property.
func (m *SelfActivateRequestBody) SetDuration(value *string)() {
    m.duration = value
}
// Sets the reason property value. 
// Parameters:
//  - value : Value to set for the reason property.
func (m *SelfActivateRequestBody) SetReason(value *string)() {
    m.reason = value
}
// Sets the ticketNumber property value. 
// Parameters:
//  - value : Value to set for the ticketNumber property.
func (m *SelfActivateRequestBody) SetTicketNumber(value *string)() {
    m.ticketNumber = value
}
// Sets the ticketSystem property value. 
// Parameters:
//  - value : Value to set for the ticketSystem property.
func (m *SelfActivateRequestBody) SetTicketSystem(value *string)() {
    m.ticketSystem = value
}
