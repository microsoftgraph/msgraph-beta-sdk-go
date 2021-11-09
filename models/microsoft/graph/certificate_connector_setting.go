package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new certificateConnectorSetting and sets the default values.
func NewCertificateConnectorSetting()(*CertificateConnectorSetting) {
    m := &CertificateConnectorSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateConnectorSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the certExpiryTime property value. Certificate expire time
func (m *CertificateConnectorSetting) GetCertExpiryTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certExpiryTime
    }
}
// Gets the connectorVersion property value. Version of certificate connector
func (m *CertificateConnectorSetting) GetConnectorVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorVersion
    }
}
// Gets the enrollmentError property value. Certificate connector enrollment error
func (m *CertificateConnectorSetting) GetEnrollmentError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentError
    }
}
// Gets the lastConnectorConnectionTime property value. Last time certificate connector connected
func (m *CertificateConnectorSetting) GetLastConnectorConnectionTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectorConnectionTime
    }
}
// Gets the lastUploadVersion property value. Version of last uploaded certificate connector
func (m *CertificateConnectorSetting) GetLastUploadVersion()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.lastUploadVersion
    }
}
// Gets the status property value. Certificate connector status
func (m *CertificateConnectorSetting) GetStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CertificateConnectorSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the certExpiryTime property value. Certificate expire time
// Parameters:
//  - value : Value to set for the certExpiryTime property.
func (m *CertificateConnectorSetting) SetCertExpiryTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certExpiryTime = value
}
// Sets the connectorVersion property value. Version of certificate connector
// Parameters:
//  - value : Value to set for the connectorVersion property.
func (m *CertificateConnectorSetting) SetConnectorVersion(value *string)() {
    m.connectorVersion = value
}
// Sets the enrollmentError property value. Certificate connector enrollment error
// Parameters:
//  - value : Value to set for the enrollmentError property.
func (m *CertificateConnectorSetting) SetEnrollmentError(value *string)() {
    m.enrollmentError = value
}
// Sets the lastConnectorConnectionTime property value. Last time certificate connector connected
// Parameters:
//  - value : Value to set for the lastConnectorConnectionTime property.
func (m *CertificateConnectorSetting) SetLastConnectorConnectionTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectorConnectionTime = value
}
// Sets the lastUploadVersion property value. Version of last uploaded certificate connector
// Parameters:
//  - value : Value to set for the lastUploadVersion property.
func (m *CertificateConnectorSetting) SetLastUploadVersion(value *int64)() {
    m.lastUploadVersion = value
}
// Sets the status property value. Certificate connector status
// Parameters:
//  - value : Value to set for the status property.
func (m *CertificateConnectorSetting) SetStatus(value *int32)() {
    m.status = value
}
