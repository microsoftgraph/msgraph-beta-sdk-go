package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody 
type MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The metricNames property
    metricNames []string
}
// NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody instantiates a new MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody and sets the default values.
func NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody()(*MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody) {
    m := &MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody) GetMetricNames()([]string) {
    return m.metricNames
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMetricNames sets the metricNames property value. The metricNames property
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemMicrosoftGraphGetHealthMetricsGetHealthMetricsPostRequestBody) SetMetricNames(value []string)() {
    m.metricNames = value
}
