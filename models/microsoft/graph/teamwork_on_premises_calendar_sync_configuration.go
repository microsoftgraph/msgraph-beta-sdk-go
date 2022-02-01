package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkOnPremisesCalendarSyncConfiguration 
type TeamworkOnPremisesCalendarSyncConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The fully qualified domain name (FQDN) of the Skype for Business Server. Use the Exchange domain if the Skype for Business SIP domain is different from the Exchange domain of the user.
    domain *string;
    // The domain and username of the console device, for example, Seattle/RanierConf.
    domainUserName *string;
    // The Simple Mail Transfer Protocol (SMTP) address of the user account. This is only required if a different user principal name (UPN) is used to sign in to Exchange other than Microsoft Teams and Skype for Business. This is a common scenario in a hybrid environment where an on-premises Exchange server is used.
    smtpAddress *string;
}
// NewTeamworkOnPremisesCalendarSyncConfiguration instantiates a new teamworkOnPremisesCalendarSyncConfiguration and sets the default values.
func NewTeamworkOnPremisesCalendarSyncConfiguration()(*TeamworkOnPremisesCalendarSyncConfiguration) {
    m := &TeamworkOnPremisesCalendarSyncConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDomain gets the domain property value. The fully qualified domain name (FQDN) of the Skype for Business Server. Use the Exchange domain if the Skype for Business SIP domain is different from the Exchange domain of the user.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domain
    }
}
// GetDomainUserName gets the domainUserName property value. The domain and username of the console device, for example, Seattle/RanierConf.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetDomainUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainUserName
    }
}
// GetSmtpAddress gets the smtpAddress property value. The Simple Mail Transfer Protocol (SMTP) address of the user account. This is only required if a different user principal name (UPN) is used to sign in to Exchange other than Microsoft Teams and Skype for Business. This is a common scenario in a hybrid environment where an on-premises Exchange server is used.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetSmtpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.smtpAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["domain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
    res["domainUserName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainUserName(val)
        }
        return nil
    }
    res["smtpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmtpAddress(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkOnPremisesCalendarSyncConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkOnPremisesCalendarSyncConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainUserName", m.GetDomainUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("smtpAddress", m.GetSmtpAddress())
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
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDomain sets the domain property value. The fully qualified domain name (FQDN) of the Skype for Business Server. Use the Exchange domain if the Skype for Business SIP domain is different from the Exchange domain of the user.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetDomain(value *string)() {
    if m != nil {
        m.domain = value
    }
}
// SetDomainUserName sets the domainUserName property value. The domain and username of the console device, for example, Seattle/RanierConf.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetDomainUserName(value *string)() {
    if m != nil {
        m.domainUserName = value
    }
}
// SetSmtpAddress sets the smtpAddress property value. The Simple Mail Transfer Protocol (SMTP) address of the user account. This is only required if a different user principal name (UPN) is used to sign in to Exchange other than Microsoft Teams and Skype for Business. This is a common scenario in a hybrid environment where an on-premises Exchange server is used.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetSmtpAddress(value *string)() {
    if m != nil {
        m.smtpAddress = value
    }
}
