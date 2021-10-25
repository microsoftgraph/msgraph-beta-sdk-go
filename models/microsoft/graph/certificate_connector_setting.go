package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CertificateConnectorSetting struct {
    additionalData map[string]interface{};
    certExpiryTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    connectorVersion *string;
    enrollmentError *string;
    lastConnectorConnectionTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastUploadVersion *int64;
    status *int32;
}
func NewCertificateConnectorSetting()(*CertificateConnectorSetting) {
    m := &CertificateConnectorSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CertificateConnectorSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CertificateConnectorSetting) GetCertExpiryTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certExpiryTime
    }
}
func (m *CertificateConnectorSetting) GetConnectorVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorVersion
    }
}
func (m *CertificateConnectorSetting) GetEnrollmentError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentError
    }
}
func (m *CertificateConnectorSetting) GetLastConnectorConnectionTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectorConnectionTime
    }
}
func (m *CertificateConnectorSetting) GetLastUploadVersion()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.lastUploadVersion
    }
}
func (m *CertificateConnectorSetting) GetStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *CertificateConnectorSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certExpiryTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCertExpiryTime(val)
        return nil
    }
    res["connectorVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConnectorVersion(val)
        return nil
    }
    res["enrollmentError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEnrollmentError(val)
        return nil
    }
    res["lastConnectorConnectionTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastConnectorConnectionTime(val)
        return nil
    }
    res["lastUploadVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLastUploadVersion(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *CertificateConnectorSetting) IsNil()(bool) {
    return m == nil
}
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
func (m *CertificateConnectorSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CertificateConnectorSetting) SetCertExpiryTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certExpiryTime = value
}
func (m *CertificateConnectorSetting) SetConnectorVersion(value *string)() {
    m.connectorVersion = value
}
func (m *CertificateConnectorSetting) SetEnrollmentError(value *string)() {
    m.enrollmentError = value
}
func (m *CertificateConnectorSetting) SetLastConnectorConnectionTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectorConnectionTime = value
}
func (m *CertificateConnectorSetting) SetLastUploadVersion(value *int64)() {
    m.lastUploadVersion = value
}
func (m *CertificateConnectorSetting) SetStatus(value *int32)() {
    m.status = value
}
