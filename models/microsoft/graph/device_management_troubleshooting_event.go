package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementTroubleshootingEvent struct {
    Entity
    // A set of string key and string value pairs which provides additional information on the Troubleshooting event
    additionalInformation []KeyValuePair;
    // Id used for tracing the failure in the service.
    correlationId *string;
    // Time when the event occurred .
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Event Name corresponding to the Troubleshooting Event. It is an Optional field
    eventName *string;
    // Object containing detailed information about the error and its remediation.
    troubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails;
}
// Instantiates a new deviceManagementTroubleshootingEvent and sets the default values.
func NewDeviceManagementTroubleshootingEvent()(*DeviceManagementTroubleshootingEvent) {
    m := &DeviceManagementTroubleshootingEvent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the additionalInformation property value. A set of string key and string value pairs which provides additional information on the Troubleshooting event
func (m *DeviceManagementTroubleshootingEvent) GetAdditionalInformation()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// Gets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// Gets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// Gets the eventName property value. Event Name corresponding to the Troubleshooting Event. It is an Optional field
func (m *DeviceManagementTroubleshootingEvent) GetEventName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventName
    }
}
// Gets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *DeviceManagementTroubleshootingEvent) GetTroubleshootingErrorDetails()(*DeviceManagementTroubleshootingErrorDetails) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingErrorDetails
    }
}
// The deserialization information for the current model
func (m *DeviceManagementTroubleshootingEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetAdditionalInformation(res)
        return nil
    }
    res["correlationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCorrelationId(val)
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEventDateTime(val)
        return nil
    }
    res["eventName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEventName(val)
        return nil
    }
    res["troubleshootingErrorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingErrorDetails() })
        if err != nil {
            return err
        }
        m.SetTroubleshootingErrorDetails(val.(*DeviceManagementTroubleshootingErrorDetails))
        return nil
    }
    return res
}
func (m *DeviceManagementTroubleshootingEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementTroubleshootingEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalInformation()))
        for i, v := range m.GetAdditionalInformation() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("additionalInformation", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventName", m.GetEventName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("troubleshootingErrorDetails", m.GetTroubleshootingErrorDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the additionalInformation property value. A set of string key and string value pairs which provides additional information on the Troubleshooting event
// Parameters:
//  - value : Value to set for the additionalInformation property.
func (m *DeviceManagementTroubleshootingEvent) SetAdditionalInformation(value []KeyValuePair)() {
    m.additionalInformation = value
}
// Sets the correlationId property value. Id used for tracing the failure in the service.
// Parameters:
//  - value : Value to set for the correlationId property.
func (m *DeviceManagementTroubleshootingEvent) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// Sets the eventDateTime property value. Time when the event occurred .
// Parameters:
//  - value : Value to set for the eventDateTime property.
func (m *DeviceManagementTroubleshootingEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// Sets the eventName property value. Event Name corresponding to the Troubleshooting Event. It is an Optional field
// Parameters:
//  - value : Value to set for the eventName property.
func (m *DeviceManagementTroubleshootingEvent) SetEventName(value *string)() {
    m.eventName = value
}
// Sets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
// Parameters:
//  - value : Value to set for the troubleshootingErrorDetails property.
func (m *DeviceManagementTroubleshootingEvent) SetTroubleshootingErrorDetails(value *DeviceManagementTroubleshootingErrorDetails)() {
    m.troubleshootingErrorDetails = value
}
