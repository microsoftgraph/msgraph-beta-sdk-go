package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementTroubleshootingEvent 
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
// NewDeviceManagementTroubleshootingEvent instantiates a new deviceManagementTroubleshootingEvent and sets the default values.
func NewDeviceManagementTroubleshootingEvent()(*DeviceManagementTroubleshootingEvent) {
    m := &DeviceManagementTroubleshootingEvent{
        Entity: *NewEntity(),
    }
    return m
}
// GetAdditionalInformation gets the additionalInformation property value. A set of string key and string value pairs which provides additional information on the Troubleshooting event
func (m *DeviceManagementTroubleshootingEvent) GetAdditionalInformation()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// GetCorrelationId gets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// GetEventDateTime gets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// GetEventName gets the eventName property value. Event Name corresponding to the Troubleshooting Event. It is an Optional field
func (m *DeviceManagementTroubleshootingEvent) GetEventName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventName
    }
}
// GetTroubleshootingErrorDetails gets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *DeviceManagementTroubleshootingEvent) GetTroubleshootingErrorDetails()(*DeviceManagementTroubleshootingErrorDetails) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingErrorDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTroubleshootingEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetAdditionalInformation(res)
        }
        return nil
    }
    res["correlationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["eventName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
        }
        return nil
    }
    res["troubleshootingErrorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingErrorDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTroubleshootingErrorDetails(val.(*DeviceManagementTroubleshootingErrorDetails))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementTroubleshootingEvent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTroubleshootingEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalInformation() != nil {
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
// SetAdditionalInformation sets the additionalInformation property value. A set of string key and string value pairs which provides additional information on the Troubleshooting event
func (m *DeviceManagementTroubleshootingEvent) SetAdditionalInformation(value []KeyValuePair)() {
    if m != nil {
        m.additionalInformation = value
    }
}
// SetCorrelationId sets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) SetCorrelationId(value *string)() {
    if m != nil {
        m.correlationId = value
    }
}
// SetEventDateTime sets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.eventDateTime = value
    }
}
// SetEventName sets the eventName property value. Event Name corresponding to the Troubleshooting Event. It is an Optional field
func (m *DeviceManagementTroubleshootingEvent) SetEventName(value *string)() {
    if m != nil {
        m.eventName = value
    }
}
// SetTroubleshootingErrorDetails sets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *DeviceManagementTroubleshootingEvent) SetTroubleshootingErrorDetails(value *DeviceManagementTroubleshootingErrorDetails)() {
    if m != nil {
        m.troubleshootingErrorDetails = value
    }
}
