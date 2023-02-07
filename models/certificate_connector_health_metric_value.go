package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateConnectorHealthMetricValue metric snapshot value returned in response to a GetHealthMetricTimeSeries request.
type CertificateConnectorHealthMetricValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Timestamp for this metric data-point.
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Count of failed requests/operations.
    failureCount *int64
    // The OdataType property
    odataType *string
    // Count of successful requests/operations.
    successCount *int64
}
// NewCertificateConnectorHealthMetricValue instantiates a new certificateConnectorHealthMetricValue and sets the default values.
func NewCertificateConnectorHealthMetricValue()(*CertificateConnectorHealthMetricValue) {
    m := &CertificateConnectorHealthMetricValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCertificateConnectorHealthMetricValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateConnectorHealthMetricValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateConnectorHealthMetricValue(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateConnectorHealthMetricValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDateTime gets the dateTime property value. Timestamp for this metric data-point.
func (m *CertificateConnectorHealthMetricValue) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateTime
}
// GetFailureCount gets the failureCount property value. Count of failed requests/operations.
func (m *CertificateConnectorHealthMetricValue) GetFailureCount()(*int64) {
    return m.failureCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateConnectorHealthMetricValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val)
        }
        return nil
    }
    res["failureCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureCount(val)
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
    res["successCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessCount(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CertificateConnectorHealthMetricValue) GetOdataType()(*string) {
    return m.odataType
}
// GetSuccessCount gets the successCount property value. Count of successful requests/operations.
func (m *CertificateConnectorHealthMetricValue) GetSuccessCount()(*int64) {
    return m.successCount
}
// Serialize serializes information the current object
func (m *CertificateConnectorHealthMetricValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("dateTime", m.GetDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("failureCount", m.GetFailureCount())
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
        err := writer.WriteInt64Value("successCount", m.GetSuccessCount())
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
func (m *CertificateConnectorHealthMetricValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDateTime sets the dateTime property value. Timestamp for this metric data-point.
func (m *CertificateConnectorHealthMetricValue) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateTime = value
}
// SetFailureCount sets the failureCount property value. Count of failed requests/operations.
func (m *CertificateConnectorHealthMetricValue) SetFailureCount(value *int64)() {
    m.failureCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CertificateConnectorHealthMetricValue) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSuccessCount sets the successCount property value. Count of successful requests/operations.
func (m *CertificateConnectorHealthMetricValue) SetSuccessCount(value *int64)() {
    m.successCount = value
}
