package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DataSharingConsent struct {
    Entity
    grantDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    granted *bool;
    grantedByUpn *string;
    grantedByUserId *string;
    serviceDisplayName *string;
    termsUrl *string;
}
func NewDataSharingConsent()(*DataSharingConsent) {
    m := &DataSharingConsent{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DataSharingConsent) GetGrantDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.grantDateTime
    }
}
func (m *DataSharingConsent) GetGranted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.granted
    }
}
func (m *DataSharingConsent) GetGrantedByUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grantedByUpn
    }
}
func (m *DataSharingConsent) GetGrantedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grantedByUserId
    }
}
func (m *DataSharingConsent) GetServiceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceDisplayName
    }
}
func (m *DataSharingConsent) GetTermsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.termsUrl
    }
}
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
func (m *DataSharingConsent) SetGrantDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.grantDateTime = value
}
func (m *DataSharingConsent) SetGranted(value *bool)() {
    m.granted = value
}
func (m *DataSharingConsent) SetGrantedByUpn(value *string)() {
    m.grantedByUpn = value
}
func (m *DataSharingConsent) SetGrantedByUserId(value *string)() {
    m.grantedByUserId = value
}
func (m *DataSharingConsent) SetServiceDisplayName(value *string)() {
    m.serviceDisplayName = value
}
func (m *DataSharingConsent) SetTermsUrl(value *string)() {
    m.termsUrl = value
}
