package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody provides operations to call the getHealthMetrics method.
type DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The metricNames property
    metricNames []string
}
// NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody instantiates a new DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody) {
    m := &DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["metricNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMetricNames(res)
        }
        return nil
    }
    return res
}
// GetMetricNames gets the metricNames property value. The metricNames property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody) GetMetricNames()([]string) {
    return m.metricNames
}
// Serialize serializes information the current object
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMetricNames() != nil {
        err := writer.WriteCollectionOfStringValues("metricNames", m.GetMetricNames())
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
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMetricNames sets the metricNames property value. The metricNames property
func (m *DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBody) SetMetricNames(value []string)() {
    m.metricNames = value
}
