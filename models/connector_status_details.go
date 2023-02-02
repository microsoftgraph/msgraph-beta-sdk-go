package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectorStatusDetails represent connector status
type ConnectorStatusDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Connector Instance Id
    connectorInstanceId *string
    // Connectors name for connector status
    connectorName *ConnectorName
    // Event datetime
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Connector health state for connector status
    status *ConnectorHealthState
}
// NewConnectorStatusDetails instantiates a new connectorStatusDetails and sets the default values.
func NewConnectorStatusDetails()(*ConnectorStatusDetails) {
    m := &ConnectorStatusDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConnectorStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectorStatusDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectorStatusDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConnectorStatusDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConnectorInstanceId gets the connectorInstanceId property value. Connector Instance Id
func (m *ConnectorStatusDetails) GetConnectorInstanceId()(*string) {
    return m.connectorInstanceId
}
// GetConnectorName gets the connectorName property value. Connectors name for connector status
func (m *ConnectorStatusDetails) GetConnectorName()(*ConnectorName) {
    return m.connectorName
}
// GetEventDateTime gets the eventDateTime property value. Event datetime
func (m *ConnectorStatusDetails) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectorStatusDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["connectorInstanceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorInstanceId(val)
        }
        return nil
    }
    res["connectorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorName(val.(*ConnectorName))
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ConnectorHealthState))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConnectorStatusDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. Connector health state for connector status
func (m *ConnectorStatusDetails) GetStatus()(*ConnectorHealthState) {
    return m.status
}
// Serialize serializes information the current object
func (m *ConnectorStatusDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("connectorInstanceId", m.GetConnectorInstanceId())
        if err != nil {
            return err
        }
    }
    if m.GetConnectorName() != nil {
        cast := (*m.GetConnectorName()).String()
        err := writer.WriteStringValue("connectorName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *ConnectorStatusDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConnectorInstanceId sets the connectorInstanceId property value. Connector Instance Id
func (m *ConnectorStatusDetails) SetConnectorInstanceId(value *string)() {
    m.connectorInstanceId = value
}
// SetConnectorName sets the connectorName property value. Connectors name for connector status
func (m *ConnectorStatusDetails) SetConnectorName(value *ConnectorName)() {
    m.connectorName = value
}
// SetEventDateTime sets the eventDateTime property value. Event datetime
func (m *ConnectorStatusDetails) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConnectorStatusDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. Connector health state for connector status
func (m *ConnectorStatusDetails) SetStatus(value *ConnectorHealthState)() {
    m.status = value
}
