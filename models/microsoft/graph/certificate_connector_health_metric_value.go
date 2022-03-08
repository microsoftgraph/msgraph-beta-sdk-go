package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CertificateConnectorHealthMetricValue provides operations to call the getHealthMetricTimeSeries method.
type CertificateConnectorHealthMetricValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Timestamp for this metric data-point.
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Count of failed requests/operations.
    failureCount *int64;
    // Count of successful requests/operations.
    successCount *int64;
}
// NewCertificateConnectorHealthMetricValue instantiates a new certificateConnectorHealthMetricValue and sets the default values.
func NewCertificateConnectorHealthMetricValue()(*CertificateConnectorHealthMetricValue) {
    m := &CertificateConnectorHealthMetricValue{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCertificateConnectorHealthMetricValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateConnectorHealthMetricValueFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCertificateConnectorHealthMetricValue(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateConnectorHealthMetricValue) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDateTime gets the dateTime property value. Timestamp for this metric data-point.
func (m *CertificateConnectorHealthMetricValue) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dateTime
    }
}
// GetFailureCount gets the failureCount property value. Count of failed requests/operations.
func (m *CertificateConnectorHealthMetricValue) GetFailureCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failureCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateConnectorHealthMetricValue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val)
        }
        return nil
    }
    res["failureCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureCount(val)
        }
        return nil
    }
    res["successCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetSuccessCount gets the successCount property value. Count of successful requests/operations.
func (m *CertificateConnectorHealthMetricValue) GetSuccessCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successCount
    }
}
func (m *CertificateConnectorHealthMetricValue) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CertificateConnectorHealthMetricValue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *CertificateConnectorHealthMetricValue) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDateTime sets the dateTime property value. Timestamp for this metric data-point.
func (m *CertificateConnectorHealthMetricValue) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.dateTime = value
    }
}
// SetFailureCount sets the failureCount property value. Count of failed requests/operations.
func (m *CertificateConnectorHealthMetricValue) SetFailureCount(value *int64)() {
    if m != nil {
        m.failureCount = value
    }
}
// SetSuccessCount sets the successCount property value. Count of successful requests/operations.
func (m *CertificateConnectorHealthMetricValue) SetSuccessCount(value *int64)() {
    if m != nil {
        m.successCount = value
    }
}
