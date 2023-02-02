package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcConnectivityEvent 
type CloudPcConnectivityEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates the date and time when this event was created. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the event.
    eventName *string
    // The eventResult property
    eventResult *CloudPcConnectivityEventResult
    // The eventType property
    eventType *CloudPcConnectivityEventType
    // Additional message for this event.
    message *string
    // The OdataType property
    odataType *string
}
// NewCloudPcConnectivityEvent instantiates a new cloudPcConnectivityEvent and sets the default values.
func NewCloudPcConnectivityEvent()(*CloudPcConnectivityEvent) {
    m := &CloudPcConnectivityEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcConnectivityEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcConnectivityEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcConnectivityEvent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcConnectivityEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEventDateTime gets the eventDateTime property value. Indicates the date and time when this event was created. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcConnectivityEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventDateTime
}
// GetEventName gets the eventName property value. Name of the event.
func (m *CloudPcConnectivityEvent) GetEventName()(*string) {
    return m.eventName
}
// GetEventResult gets the eventResult property value. The eventResult property
func (m *CloudPcConnectivityEvent) GetEventResult()(*CloudPcConnectivityEventResult) {
    return m.eventResult
}
// GetEventType gets the eventType property value. The eventType property
func (m *CloudPcConnectivityEvent) GetEventType()(*CloudPcConnectivityEventType) {
    return m.eventType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcConnectivityEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["eventResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcConnectivityEventResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventResult(val.(*CloudPcConnectivityEventResult))
        }
        return nil
    }
    res["eventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcConnectivityEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventType(val.(*CloudPcConnectivityEventType))
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. Additional message for this event.
func (m *CloudPcConnectivityEvent) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcConnectivityEvent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CloudPcConnectivityEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventName", m.GetEventName())
        if err != nil {
            return err
        }
    }
    if m.GetEventResult() != nil {
        cast := (*m.GetEventResult()).String()
        err := writer.WriteStringValue("eventResult", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEventType() != nil {
        cast := (*m.GetEventType()).String()
        err := writer.WriteStringValue("eventType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *CloudPcConnectivityEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEventDateTime sets the eventDateTime property value. Indicates the date and time when this event was created. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcConnectivityEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// SetEventName sets the eventName property value. Name of the event.
func (m *CloudPcConnectivityEvent) SetEventName(value *string)() {
    m.eventName = value
}
// SetEventResult sets the eventResult property value. The eventResult property
func (m *CloudPcConnectivityEvent) SetEventResult(value *CloudPcConnectivityEventResult)() {
    m.eventResult = value
}
// SetEventType sets the eventType property value. The eventType property
func (m *CloudPcConnectivityEvent) SetEventType(value *CloudPcConnectivityEventType)() {
    m.eventType = value
}
// SetMessage sets the message property value. Additional message for this event.
func (m *CloudPcConnectivityEvent) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcConnectivityEvent) SetOdataType(value *string)() {
    m.odataType = value
}
