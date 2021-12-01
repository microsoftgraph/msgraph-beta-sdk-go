package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SymantecCodeSigningCertificate 
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
// NewSymantecCodeSigningCertificate instantiates a new symantecCodeSigningCertificate and sets the default values.
func NewSymantecCodeSigningCertificate()(*SymantecCodeSigningCertificate) {
    m := &SymantecCodeSigningCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// GetContent gets the content property value. The Windows Symantec Code-Signing Certificate in the raw data format.
func (m *SymantecCodeSigningCertificate) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The Cert Expiration Date.
func (m *SymantecCodeSigningCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetIssuer gets the issuer property value. The Issuer value for the cert.
func (m *SymantecCodeSigningCertificate) GetIssuer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuer
    }
}
// GetIssuerName gets the issuerName property value. The Issuer Name for the cert.
func (m *SymantecCodeSigningCertificate) GetIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerName
    }
}
// GetPassword gets the password property value. The Password required for .pfx file.
func (m *SymantecCodeSigningCertificate) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// GetStatus gets the status property value. The Cert Status Provisioned or not Provisioned. Possible values are: notProvisioned, provisioned.
func (m *SymantecCodeSigningCertificate) GetStatus()(*CertificateStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSubject gets the subject property value. The Subject value for the cert.
func (m *SymantecCodeSigningCertificate) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetSubjectName gets the subjectName property value. The Subject Name for the cert.
func (m *SymantecCodeSigningCertificate) GetSubjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectName
    }
}
// GetUploadDateTime gets the uploadDateTime property value. The Type of the CodeSigning Cert as Symantec Cert.
func (m *SymantecCodeSigningCertificate) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.uploadDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetContent sets the content property value. The Windows Symantec Code-Signing Certificate in the raw data format.
func (m *SymantecCodeSigningCertificate) SetContent(value []byte)() {
    if m != nil {
        m.content = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The Cert Expiration Date.
func (m *SymantecCodeSigningCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetIssuer sets the issuer property value. The Issuer value for the cert.
func (m *SymantecCodeSigningCertificate) SetIssuer(value *string)() {
    if m != nil {
        m.issuer = value
    }
}
// SetIssuerName sets the issuerName property value. The Issuer Name for the cert.
func (m *SymantecCodeSigningCertificate) SetIssuerName(value *string)() {
    if m != nil {
        m.issuerName = value
    }
}
// SetPassword sets the password property value. The Password required for .pfx file.
func (m *SymantecCodeSigningCertificate) SetPassword(value *string)() {
    if m != nil {
        m.password = value
    }
}
// SetStatus sets the status property value. The Cert Status Provisioned or not Provisioned. Possible values are: notProvisioned, provisioned.
func (m *SymantecCodeSigningCertificate) SetStatus(value *CertificateStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetSubject sets the subject property value. The Subject value for the cert.
func (m *SymantecCodeSigningCertificate) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
// SetSubjectName sets the subjectName property value. The Subject Name for the cert.
func (m *SymantecCodeSigningCertificate) SetSubjectName(value *string)() {
    if m != nil {
        m.subjectName = value
    }
}
// SetUploadDateTime sets the uploadDateTime property value. The Type of the CodeSigning Cert as Symantec Cert.
func (m *SymantecCodeSigningCertificate) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.uploadDateTime = value
    }
}
