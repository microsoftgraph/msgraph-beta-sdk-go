package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DataSharingConsent struct {
    Entity
    // The time consent was granted for this account
    grantDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The granted state for the data sharing consent
    granted *bool;
    // The Upn of the user that granted consent for this account
    grantedByUpn *string;
    // The UserId of the user that granted consent for this account
    grantedByUserId *string;
    // The display name of the service work flow
    serviceDisplayName *string;
    // The TermsUrl for the data sharing consent
    termsUrl *string;
}
// Instantiates a new dataSharingConsent and sets the default values.
func NewDataSharingConsent()(*DataSharingConsent) {
    m := &DataSharingConsent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the grantDateTime property value. The time consent was granted for this account
func (m *DataSharingConsent) GetGrantDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.grantDateTime
    }
}
// Gets the granted property value. The granted state for the data sharing consent
func (m *DataSharingConsent) GetGranted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.granted
    }
}
// Gets the grantedByUpn property value. The Upn of the user that granted consent for this account
func (m *DataSharingConsent) GetGrantedByUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grantedByUpn
    }
}
// Gets the grantedByUserId property value. The UserId of the user that granted consent for this account
func (m *DataSharingConsent) GetGrantedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grantedByUserId
    }
}
// Gets the serviceDisplayName property value. The display name of the service work flow
func (m *DataSharingConsent) GetServiceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceDisplayName
    }
}
// Gets the termsUrl property value. The TermsUrl for the data sharing consent
func (m *DataSharingConsent) GetTermsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.termsUrl
    }
}
// The deserialization information for the current model
func (m *DataSharingConsent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["grantDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetGrantDateTime(val)
        return nil
    }
    res["granted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetGranted(val)
        return nil
    }
    res["grantedByUpn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGrantedByUpn(val)
        return nil
    }
    res["grantedByUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGrantedByUserId(val)
        return nil
    }
    res["serviceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServiceDisplayName(val)
        return nil
    }
    res["termsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTermsUrl(val)
        return nil
    }
    return res
}
func (m *DataSharingConsent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DataSharingConsent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("grantDateTime", m.GetGrantDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("granted", m.GetGranted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("grantedByUpn", m.GetGrantedByUpn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("grantedByUserId", m.GetGrantedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceDisplayName", m.GetServiceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("termsUrl", m.GetTermsUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the grantDateTime property value. The time consent was granted for this account
// Parameters:
//  - value : Value to set for the grantDateTime property.
func (m *DataSharingConsent) SetGrantDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.grantDateTime = value
}
// Sets the granted property value. The granted state for the data sharing consent
// Parameters:
//  - value : Value to set for the granted property.
func (m *DataSharingConsent) SetGranted(value *bool)() {
    m.granted = value
}
// Sets the grantedByUpn property value. The Upn of the user that granted consent for this account
// Parameters:
//  - value : Value to set for the grantedByUpn property.
func (m *DataSharingConsent) SetGrantedByUpn(value *string)() {
    m.grantedByUpn = value
}
// Sets the grantedByUserId property value. The UserId of the user that granted consent for this account
// Parameters:
//  - value : Value to set for the grantedByUserId property.
func (m *DataSharingConsent) SetGrantedByUserId(value *string)() {
    m.grantedByUserId = value
}
// Sets the serviceDisplayName property value. The display name of the service work flow
// Parameters:
//  - value : Value to set for the serviceDisplayName property.
func (m *DataSharingConsent) SetServiceDisplayName(value *string)() {
    m.serviceDisplayName = value
}
// Sets the termsUrl property value. The TermsUrl for the data sharing consent
// Parameters:
//  - value : Value to set for the termsUrl property.
func (m *DataSharingConsent) SetTermsUrl(value *string)() {
    m.termsUrl = value
}
