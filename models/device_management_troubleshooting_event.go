package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTroubleshootingEvent event representing an general failure.
type DeviceManagementTroubleshootingEvent struct {
    Entity
    // A set of string key and string value pairs which provides additional information on the Troubleshooting event
    additionalInformation []KeyValuePairable
    // Id used for tracing the failure in the service.
    correlationId *string
    // Time when the event occurred .
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Event Name corresponding to the Troubleshooting Event. It is an Optional field
    eventName *string
    // Object containing detailed information about the error and its remediation.
    troubleshootingErrorDetails DeviceManagementTroubleshootingErrorDetailsable
}
// NewDeviceManagementTroubleshootingEvent instantiates a new deviceManagementTroubleshootingEvent and sets the default values.
func NewDeviceManagementTroubleshootingEvent()(*DeviceManagementTroubleshootingEvent) {
    m := &DeviceManagementTroubleshootingEvent{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementTroubleshootingEvent";
    m.SetOdataType(&odataTypeValue);
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
    return m.additionalInformation
}
// GetCorrelationId gets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetEventDateTime gets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventDateTime
}
// GetEventName gets the eventName property value. Event Name corresponding to the Troubleshooting Event. It is an Optional field
func (m *DeviceManagementTroubleshootingEvent) GetEventName()(*string) {
    return m.eventName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTroubleshootingEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetAdditionalInformation)
    res["correlationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCorrelationId)
    res["eventDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEventDateTime)
    res["eventName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEventName)
    res["troubleshootingErrorDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementTroubleshootingErrorDetailsFromDiscriminatorValue , m.SetTroubleshootingErrorDetails)
    return res
}
// GetTroubleshootingErrorDetails gets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *DeviceManagementTroubleshootingEvent) GetTroubleshootingErrorDetails()(DeviceManagementTroubleshootingErrorDetailsable) {
    return m.troubleshootingErrorDetails
}
// Serialize serializes information the current object
func (m *DeviceManagementTroubleshootingEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalInformation() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAdditionalInformation())
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
    m.additionalInformation = value
}
// SetCorrelationId sets the correlationId property value. Id used for tracing the failure in the service.
func (m *DeviceManagementTroubleshootingEvent) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetEventDateTime sets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementTroubleshootingEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// SetEventName sets the eventName property value. Event Name corresponding to the Troubleshooting Event. It is an Optional field
func (m *DeviceManagementTroubleshootingEvent) SetEventName(value *string)() {
    m.eventName = value
}
// SetTroubleshootingErrorDetails sets the troubleshootingErrorDetails property value. Object containing detailed information about the error and its remediation.
func (m *DeviceManagementTroubleshootingEvent) SetTroubleshootingErrorDetails(value DeviceManagementTroubleshootingErrorDetailsable)() {
    m.troubleshootingErrorDetails = value
}
