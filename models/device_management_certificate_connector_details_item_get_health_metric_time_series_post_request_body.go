package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody provides operations to call the getHealthMetricTimeSeries method.
type DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The timeSeries property
    timeSeries TimeSeriesParameterable
}
// NewDeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody instantiates a new DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody and sets the default values.
func NewDeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody()(*DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) {
    m := &DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["timeSeries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeSeriesParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeSeries(val.(TimeSeriesParameterable))
        }
        return nil
    }
    return res
}
// GetTimeSeries gets the timeSeries property value. The timeSeries property
func (m *DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) GetTimeSeries()(TimeSeriesParameterable) {
    return m.timeSeries
}
// Serialize serializes information the current object
func (m *DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("timeSeries", m.GetTimeSeries())
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
func (m *DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetTimeSeries sets the timeSeries property value. The timeSeries property
func (m *DeviceManagementCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) SetTimeSeries(value TimeSeriesParameterable)() {
    m.timeSeries = value
}
