package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcConnectivityEvent 
type CloudPcConnectivityEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.cloudPcConnectivityEvent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcConnectivityEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcConnectivityEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcConnectivityEvent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcConnectivityEvent) GetAdditionalData()(map[string]interface{}) {
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
    res["eventDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEventDateTime)
    res["eventName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEventName)
    res["eventResult"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcConnectivityEventResult , m.SetEventResult)
    res["eventType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcConnectivityEventType , m.SetEventType)
    res["message"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMessage)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
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
func (m *CloudPcConnectivityEvent) SetAdditionalData(value map[string]interface{})() {
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
