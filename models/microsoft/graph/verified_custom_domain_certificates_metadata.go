package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type VerifiedCustomDomainCertificatesMetadata struct {
    additionalData map[string]interface{};
    expiryDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    issueDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    issuerName *string;
    subjectName *string;
    thumbprint *string;
}
func NewVerifiedCustomDomainCertificatesMetadata()(*VerifiedCustomDomainCertificatesMetadata) {
    m := &VerifiedCustomDomainCertificatesMetadata{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *VerifiedCustomDomainCertificatesMetadata) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *VerifiedCustomDomainCertificatesMetadata) GetExpiryDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiryDate
    }
}
func (m *VerifiedCustomDomainCertificatesMetadata) GetIssueDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.issueDate
    }
}
func (m *VerifiedCustomDomainCertificatesMetadata) GetIssuerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerName
    }
}
func (m *VerifiedCustomDomainCertificatesMetadata) GetSubjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectName
    }
}
func (m *VerifiedCustomDomainCertificatesMetadata) GetThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbprint
    }
}
func (m *VerifiedCustomDomainCertificatesMetadata) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["expiryDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpiryDate(val)
        return nil
    }
    res["issueDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetIssueDate(val)
        return nil
    }
    res["issuerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuerName(val)
        return nil
    }
    res["subjectName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubjectName(val)
        return nil
    }
    res["thumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThumbprint(val)
        return nil
    }
    return res
}
func (m *VerifiedCustomDomainCertificatesMetadata) IsNil()(bool) {
    return m == nil
}
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
func (m *VerifiedCustomDomainCertificatesMetadata) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *VerifiedCustomDomainCertificatesMetadata) SetExpiryDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiryDate = value
}
func (m *VerifiedCustomDomainCertificatesMetadata) SetIssueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.issueDate = value
}
func (m *VerifiedCustomDomainCertificatesMetadata) SetIssuerName(value *string)() {
    m.issuerName = value
}
func (m *VerifiedCustomDomainCertificatesMetadata) SetSubjectName(value *string)() {
    m.subjectName = value
}
func (m *VerifiedCustomDomainCertificatesMetadata) SetThumbprint(value *string)() {
    m.thumbprint = value
}
