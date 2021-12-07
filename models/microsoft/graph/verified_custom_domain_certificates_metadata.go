package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VerifiedCustomDomainCertificatesMetadata 
type VerifiedCustomDomainCertificatesMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The expiry date of the custom domain certificate. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    expiryDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The issue date of the custom domain. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    issueDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The issuer name of the custom domain certificate.
    issuerName *string;
    // The subject name of the custom domain certificate.
    subjectName *string;
    // The thumbprint associated with the custom domain certificate.
    thumbprint *string;
}
// NewVerifiedCustomDomainCertificatesMetadata instantiates a new verifiedCustomDomainCertificatesMetadata and sets the default values.
func NewVerifiedCustomDomainCertificatesMetadata()(*VerifiedCustomDomainCertificatesMetadata) {
    m := &VerifiedCustomDomainCertificatesMetadata{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifiedCustomDomainCertificatesMetadata) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExpiryDate gets the expiryDate property value. The expiry date of the custom domain certificate. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VerifiedCustomDomainCertificatesMetadata) GetExpiryDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiryDate
    }
}
// GetIssueDate gets the issueDate property value. The issue date of the custom domain. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VerifiedCustomDomainCertificatesMetadata) GetIssueDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.issueDate
    }
}
// GetIssuerName gets the issuerName property value. The issuer name of the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) GetIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerName
    }
}
// GetSubjectName gets the subjectName property value. The subject name of the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) GetSubjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectName
    }
}
// GetThumbprint gets the thumbprint property value. The thumbprint associated with the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) GetThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbprint
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifiedCustomDomainCertificatesMetadata) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expiryDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryDate(val)
        }
        return nil
    }
    res["issueDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueDate(val)
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
    res["thumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbprint(val)
        }
        return nil
    }
    return res
}
func (m *VerifiedCustomDomainCertificatesMetadata) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *VerifiedCustomDomainCertificatesMetadata) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expiryDate", m.GetExpiryDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("issueDate", m.GetIssueDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuerName", m.GetIssuerName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subjectName", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbprint", m.GetThumbprint())
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
func (m *VerifiedCustomDomainCertificatesMetadata) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExpiryDate sets the expiryDate property value. The expiry date of the custom domain certificate. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VerifiedCustomDomainCertificatesMetadata) SetExpiryDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expiryDate = value
    }
}
// SetIssueDate sets the issueDate property value. The issue date of the custom domain. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VerifiedCustomDomainCertificatesMetadata) SetIssueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.issueDate = value
    }
}
// SetIssuerName sets the issuerName property value. The issuer name of the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) SetIssuerName(value *string)() {
    if m != nil {
        m.issuerName = value
    }
}
// SetSubjectName sets the subjectName property value. The subject name of the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) SetSubjectName(value *string)() {
    if m != nil {
        m.subjectName = value
    }
}
// SetThumbprint sets the thumbprint property value. The thumbprint associated with the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) SetThumbprint(value *string)() {
    if m != nil {
        m.thumbprint = value
    }
}
