package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTroubleshootingEvent event representing an general failure.
type DeviceManagementTroubleshootingEvent struct {
    Entity
}
// NewDeviceManagementTroubleshootingEvent instantiates a new deviceManagementTroubleshootingEvent and sets the default values.
func NewDeviceManagementTroubleshootingEvent()(*DeviceManagementTroubleshootingEvent) {
    m := &DeviceManagementTroubleshootingEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.appleVppTokenTroubleshootingEvent":
                        return NewAppleVppTokenTroubleshootingEvent(), nil
                    case "#microsoft.graph.enrollmentTroubleshootingEvent":
                        return NewEnrollmentTroubleshootingEvent(), nil
                    case "#microsoft.graph.mobileAppTroubleshootingEvent":
                        return NewMobileAppTroubleshootingEvent(), nil
                }
            }
        }
    }
    return NewDeviceManagementTroubleshootingEvent(), nil
}
// GetAdditionalInformation gets the additionalInformation property value. A set of string key and string value pairs which provides additional information on the Troubleshooting event
func (m *DeviceManagementTroubleshootingEvent) GetAdditionalInformation()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("additionalInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetCorrelationId gets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) GetCorrelationId()(*string) {
    val, err := m.GetBackingStore().Get("correlationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEventDateTime gets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("eventDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEventName gets the eventName property value. Event Name corresponding to the Troubleshooting Event. It is an Optional field
func (m *DeviceManagementTroubleshootingEvent) GetEventName()(*string) {
    val, err := m.GetBackingStore().Get("eventName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTroubleshootingEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValuePairable)
                }
            }
            m.SetAdditionalInformation(res)
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["eventName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
        }
        return nil
    }
    res["troubleshootingErrorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTroubleshootingErrorDetails(val.(DeviceManagementTroubleshootingErrorDetailsable))
        }
        return nil
    }
    return res
}
// GetTroubleshootingErrorDetails gets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *DeviceManagementTroubleshootingEvent) GetTroubleshootingErrorDetails()(DeviceManagementTroubleshootingErrorDetailsable) {
    val, err := m.GetBackingStore().Get("troubleshootingErrorDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementTroubleshootingErrorDetailsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTroubleshootingEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalInformation()))
        for i, v := range m.GetAdditionalInformation() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
func (m *DeviceManagementTroubleshootingEvent) SetAdditionalInformation(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("additionalInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationId sets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) SetCorrelationId(value *string)() {
    err := m.GetBackingStore().Set("correlationId", value)
    if err != nil {
        panic(err)
    }
}
// SetEventDateTime sets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("eventDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEventName sets the eventName property value. Event Name corresponding to the Troubleshooting Event. It is an Optional field
func (m *DeviceManagementTroubleshootingEvent) SetEventName(value *string)() {
    err := m.GetBackingStore().Set("eventName", value)
    if err != nil {
        panic(err)
    }
}
// SetTroubleshootingErrorDetails sets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *DeviceManagementTroubleshootingEvent) SetTroubleshootingErrorDetails(value DeviceManagementTroubleshootingErrorDetailsable)() {
    err := m.GetBackingStore().Set("troubleshootingErrorDetails", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementTroubleshootingEventable 
type DeviceManagementTroubleshootingEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalInformation()([]KeyValuePairable)
    GetCorrelationId()(*string)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEventName()(*string)
    GetTroubleshootingErrorDetails()(DeviceManagementTroubleshootingErrorDetailsable)
    SetAdditionalInformation(value []KeyValuePairable)()
    SetCorrelationId(value *string)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEventName(value *string)()
    SetTroubleshootingErrorDetails(value DeviceManagementTroubleshootingErrorDetailsable)()
}
