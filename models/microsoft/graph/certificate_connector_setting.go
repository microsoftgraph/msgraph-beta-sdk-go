package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CertificateConnectorSetting 
type CertificateConnectorSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Certificate expire time
    certExpiryTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Version of certificate connector
    connectorVersion *string;
    // Certificate connector enrollment error
    enrollmentError *string;
    // Last time certificate connector connected
    lastConnectorConnectionTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Version of last uploaded certificate connector
    lastUploadVersion *int64;
    // Certificate connector status
    status *int32;
}
// NewCertificateConnectorSetting instantiates a new certificateConnectorSetting and sets the default values.
func NewCertificateConnectorSetting()(*CertificateConnectorSetting) {
    m := &CertificateConnectorSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateConnectorSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertExpiryTime gets the certExpiryTime property value. Certificate expire time
func (m *CertificateConnectorSetting) GetCertExpiryTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certExpiryTime
    }
}
// GetConnectorVersion gets the connectorVersion property value. Version of certificate connector
func (m *CertificateConnectorSetting) GetConnectorVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorVersion
    }
}
// GetEnrollmentError gets the enrollmentError property value. Certificate connector enrollment error
func (m *CertificateConnectorSetting) GetEnrollmentError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentError
    }
}
// GetLastConnectorConnectionTime gets the lastConnectorConnectionTime property value. Last time certificate connector connected
func (m *CertificateConnectorSetting) GetLastConnectorConnectionTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectorConnectionTime
    }
}
// GetLastUploadVersion gets the lastUploadVersion property value. Version of last uploaded certificate connector
func (m *CertificateConnectorSetting) GetLastUploadVersion()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.lastUploadVersion
    }
}
// GetStatus gets the status property value. Certificate connector status
func (m *CertificateConnectorSetting) GetStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateConnectorSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certExpiryTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertExpiryTime(val)
        }
        return nil
    }
    res["connectorVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorVersion(val)
        }
        return nil
    }
    res["enrollmentError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentError(val)
        }
        return nil
    }
    res["lastConnectorConnectionTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastConnectorConnectionTime(val)
        }
        return nil
    }
    res["lastUploadVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUploadVersion(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
func (m *CertificateConnectorSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CertificateConnectorSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("certExpiryTime", m.GetCertExpiryTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectorVersion", m.GetConnectorVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("enrollmentError", m.GetEnrollmentError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastConnectorConnectionTime", m.GetLastConnectorConnectionTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("lastUploadVersion", m.GetLastUploadVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("status", m.GetStatus())
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
func (m *CertificateConnectorSetting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertExpiryTime sets the certExpiryTime property value. Certificate expire time
func (m *CertificateConnectorSetting) SetCertExpiryTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.certExpiryTime = value
    }
}
// SetConnectorVersion sets the connectorVersion property value. Version of certificate connector
func (m *CertificateConnectorSetting) SetConnectorVersion(value *string)() {
    if m != nil {
        m.connectorVersion = value
    }
}
// SetEnrollmentError sets the enrollmentError property value. Certificate connector enrollment error
func (m *CertificateConnectorSetting) SetEnrollmentError(value *string)() {
    if m != nil {
        m.enrollmentError = value
    }
}
// SetLastConnectorConnectionTime sets the lastConnectorConnectionTime property value. Last time certificate connector connected
func (m *CertificateConnectorSetting) SetLastConnectorConnectionTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastConnectorConnectionTime = value
    }
}
// SetLastUploadVersion sets the lastUploadVersion property value. Version of last uploaded certificate connector
func (m *CertificateConnectorSetting) SetLastUploadVersion(value *int64)() {
    if m != nil {
        m.lastUploadVersion = value
    }
}
// SetStatus sets the status property value. Certificate connector status
func (m *CertificateConnectorSetting) SetStatus(value *int32)() {
    if m != nil {
        m.status = value
    }
}
