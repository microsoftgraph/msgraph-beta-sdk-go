package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody provides operations to call the getHealthMetricTimeSeries method.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The metricName property
    metricName *string
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody instantiates a new DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) {
    m := &DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["metricName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetricName(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetMetricName gets the metricName property value. The metricName property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) GetMetricName()(*string) {
    return m.metricName
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("metricName", m.GetMetricName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetMetricName sets the metricName property value. The metricName property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) SetMetricName(value *string)() {
    m.metricName = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricTimeSeriesPostRequestBody) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
