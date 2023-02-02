package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody 
type CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The timeSeries property
    timeSeries ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSeriesParameterable
}
// NewCertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody instantiates a new CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody and sets the default values.
func NewCertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody()(*CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody) {
    m := &CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["timeSeries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeSeriesParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeSeries(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSeriesParameterable))
        }
        return nil
    }
    return res
}
// GetTimeSeries gets the timeSeries property value. The timeSeries property
func (m *CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody) GetTimeSeries()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSeriesParameterable) {
    return m.timeSeries
}
// Serialize serializes information the current object
func (m *CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTimeSeries sets the timeSeries property value. The timeSeries property
func (m *CertificateConnectorDetailsItemMicrosoftGraphGetHealthMetricTimeSeriesGetHealthMetricTimeSeriesPostRequestBody) SetTimeSeries(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSeriesParameterable)() {
    m.timeSeries = value
}
