package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Certification 
type Certification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // URL that shows certification details for the application.
    certificationDetailsUrl *string;
    // The timestamp when the current certification for the application will expire.
    certificationExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates whether the application is certified by Microsoft.
    isCertifiedByMicrosoft *bool;
    // Indicates whether the application has been self-attested by the application developer or the publisher.
    isPublisherAttested *bool;
    // The timestamp when the certification for the application was most recently added or updated.
    lastCertificationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewCertification instantiates a new certification and sets the default values.
func NewCertification()(*Certification) {
    m := &Certification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Certification) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificationDetailsUrl gets the certificationDetailsUrl property value. URL that shows certification details for the application.
func (m *Certification) GetCertificationDetailsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificationDetailsUrl
    }
}
// GetCertificationExpirationDateTime gets the certificationExpirationDateTime property value. The timestamp when the current certification for the application will expire.
func (m *Certification) GetCertificationExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.certificationExpirationDateTime
    }
}
// GetIsCertifiedByMicrosoft gets the isCertifiedByMicrosoft property value. Indicates whether the application is certified by Microsoft.
func (m *Certification) GetIsCertifiedByMicrosoft()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCertifiedByMicrosoft
    }
}
// GetIsPublisherAttested gets the isPublisherAttested property value. Indicates whether the application has been self-attested by the application developer or the publisher.
func (m *Certification) GetIsPublisherAttested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPublisherAttested
    }
}
// GetLastCertificationDateTime gets the lastCertificationDateTime property value. The timestamp when the certification for the application was most recently added or updated.
func (m *Certification) GetLastCertificationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCertificationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Certification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificationDetailsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationDetailsUrl(val)
        }
        return nil
    }
    res["certificationExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationExpirationDateTime(val)
        }
        return nil
    }
    res["isCertifiedByMicrosoft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCertifiedByMicrosoft(val)
        }
        return nil
    }
    res["isPublisherAttested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPublisherAttested(val)
        }
        return nil
    }
    res["lastCertificationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCertificationDateTime(val)
        }
        return nil
    }
    return res
}
func (m *Certification) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Certification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("certificationDetailsUrl", m.GetCertificationDetailsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("certificationExpirationDateTime", m.GetCertificationExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCertifiedByMicrosoft", m.GetIsCertifiedByMicrosoft())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPublisherAttested", m.GetIsPublisherAttested())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastCertificationDateTime", m.GetLastCertificationDateTime())
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
func (m *Certification) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificationDetailsUrl sets the certificationDetailsUrl property value. URL that shows certification details for the application.
func (m *Certification) SetCertificationDetailsUrl(value *string)() {
    if m != nil {
        m.certificationDetailsUrl = value
    }
}
// SetCertificationExpirationDateTime sets the certificationExpirationDateTime property value. The timestamp when the current certification for the application will expire.
func (m *Certification) SetCertificationExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.certificationExpirationDateTime = value
    }
}
// SetIsCertifiedByMicrosoft sets the isCertifiedByMicrosoft property value. Indicates whether the application is certified by Microsoft.
func (m *Certification) SetIsCertifiedByMicrosoft(value *bool)() {
    if m != nil {
        m.isCertifiedByMicrosoft = value
    }
}
// SetIsPublisherAttested sets the isPublisherAttested property value. Indicates whether the application has been self-attested by the application developer or the publisher.
func (m *Certification) SetIsPublisherAttested(value *bool)() {
    if m != nil {
        m.isPublisherAttested = value
    }
}
// SetLastCertificationDateTime sets the lastCertificationDateTime property value. The timestamp when the certification for the application was most recently added or updated.
func (m *Certification) SetLastCertificationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastCertificationDateTime = value
    }
}
