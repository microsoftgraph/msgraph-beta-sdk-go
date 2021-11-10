package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SymantecCodeSigningCertificate struct {
    Entity
    // The Windows Symantec Code-Signing Certificate in the raw data format.
    content []byte;
    // The Cert Expiration Date.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Issuer value for the cert.
    issuer *string;
    // The Issuer Name for the cert.
    issuerName *string;
    // The Password required for .pfx file.
    password *string;
    // The Cert Status Provisioned or not Provisioned. Possible values are: notProvisioned, provisioned.
    status *CertificateStatus;
    // The Subject value for the cert.
    subject *string;
    // The Subject Name for the cert.
    subjectName *string;
    // The Type of the CodeSigning Cert as Symantec Cert.
    uploadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new symantecCodeSigningCertificate and sets the default values.
func NewSymantecCodeSigningCertificate()(*SymantecCodeSigningCertificate) {
    m := &SymantecCodeSigningCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the content property value. The Windows Symantec Code-Signing Certificate in the raw data format.
func (m *SymantecCodeSigningCertificate) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the expirationDateTime property value. The Cert Expiration Date.
func (m *SymantecCodeSigningCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the issuer property value. The Issuer value for the cert.
func (m *SymantecCodeSigningCertificate) GetIssuer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuer
    }
}
// Gets the issuerName property value. The Issuer Name for the cert.
func (m *SymantecCodeSigningCertificate) GetIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerName
    }
}
// Gets the password property value. The Password required for .pfx file.
func (m *SymantecCodeSigningCertificate) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// Gets the status property value. The Cert Status Provisioned or not Provisioned. Possible values are: notProvisioned, provisioned.
func (m *SymantecCodeSigningCertificate) GetStatus()(*CertificateStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the subject property value. The Subject value for the cert.
func (m *SymantecCodeSigningCertificate) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Gets the subjectName property value. The Subject Name for the cert.
func (m *SymantecCodeSigningCertificate) GetSubjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectName
    }
}
// Gets the uploadDateTime property value. The Type of the CodeSigning Cert as Symantec Cert.
func (m *SymantecCodeSigningCertificate) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.uploadDateTime
    }
}
// The deserialization information for the current model
func (m *SymantecCodeSigningCertificate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["issuer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerName(val)
        }
        return nil
    }
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CertificateStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["subjectName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectName(val)
        }
        return nil
    }
    res["uploadDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadDateTime(val)
        }
        return nil
    }
    return res
}
func (m *SymantecCodeSigningCertificate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SymantecCodeSigningCertificate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerName", m.GetIssuerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectName", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("uploadDateTime", m.GetUploadDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the content property value. The Windows Symantec Code-Signing Certificate in the raw data format.
// Parameters:
//  - value : Value to set for the content property.
func (m *SymantecCodeSigningCertificate) SetContent(value []byte)() {
    m.content = value
}
// Sets the expirationDateTime property value. The Cert Expiration Date.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *SymantecCodeSigningCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the issuer property value. The Issuer value for the cert.
// Parameters:
//  - value : Value to set for the issuer property.
func (m *SymantecCodeSigningCertificate) SetIssuer(value *string)() {
    m.issuer = value
}
// Sets the issuerName property value. The Issuer Name for the cert.
// Parameters:
//  - value : Value to set for the issuerName property.
func (m *SymantecCodeSigningCertificate) SetIssuerName(value *string)() {
    m.issuerName = value
}
// Sets the password property value. The Password required for .pfx file.
// Parameters:
//  - value : Value to set for the password property.
func (m *SymantecCodeSigningCertificate) SetPassword(value *string)() {
    m.password = value
}
// Sets the status property value. The Cert Status Provisioned or not Provisioned. Possible values are: notProvisioned, provisioned.
// Parameters:
//  - value : Value to set for the status property.
func (m *SymantecCodeSigningCertificate) SetStatus(value *CertificateStatus)() {
    m.status = value
}
// Sets the subject property value. The Subject value for the cert.
// Parameters:
//  - value : Value to set for the subject property.
func (m *SymantecCodeSigningCertificate) SetSubject(value *string)() {
    m.subject = value
}
// Sets the subjectName property value. The Subject Name for the cert.
// Parameters:
//  - value : Value to set for the subjectName property.
func (m *SymantecCodeSigningCertificate) SetSubjectName(value *string)() {
    m.subjectName = value
}
// Sets the uploadDateTime property value. The Type of the CodeSigning Cert as Symantec Cert.
// Parameters:
//  - value : Value to set for the uploadDateTime property.
func (m *SymantecCodeSigningCertificate) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadDateTime = value
}
