package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody provides operations to call the getHealthMetricTimeSeries method.
type CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The timeSeries property
    timeSeries ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSeriesParameterable
}
// NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody instantiates a new CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody and sets the default values.
func NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody()(*CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) {
    m := &CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) GetTimeSeries()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSeriesParameterable) {
    return m.timeSeries
}
// Serialize serializes information the current object
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetTimeSeries sets the timeSeries property value. The timeSeries property
func (m *CertificateConnectorDetailsItemGetHealthMetricTimeSeriesPostRequestBody) SetTimeSeries(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeSeriesParameterable)() {
    m.timeSeries = value
}
