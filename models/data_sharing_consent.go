package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataSharingConsent 
type DataSharingConsent struct {
    Entity
    // The time consent was granted for this account
    grantDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The granted state for the data sharing consent
    granted *bool
    // The Upn of the user that granted consent for this account
    grantedByUpn *string
    // The UserId of the user that granted consent for this account
    grantedByUserId *string
    // The display name of the service work flow
    serviceDisplayName *string
    // The TermsUrl for the data sharing consent
    termsUrl *string
}
// NewDataSharingConsent instantiates a new dataSharingConsent and sets the default values.
func NewDataSharingConsent()(*DataSharingConsent) {
    m := &DataSharingConsent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDataSharingConsentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataSharingConsentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataSharingConsent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataSharingConsent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["grantDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrantDateTime(val)
        }
        return nil
    }
    res["granted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGranted(val)
        }
        return nil
    }
    res["grantedByUpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrantedByUpn(val)
        }
        return nil
    }
    res["grantedByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrantedByUserId(val)
        }
        return nil
    }
    res["serviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceDisplayName(val)
        }
        return nil
    }
    res["termsUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsUrl(val)
        }
        return nil
    }
    return res
}
// GetGrantDateTime gets the grantDateTime property value. The time consent was granted for this account
func (m *DataSharingConsent) GetGrantDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.grantDateTime
    }
}
// GetGranted gets the granted property value. The granted state for the data sharing consent
func (m *DataSharingConsent) GetGranted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.granted
    }
}
// GetGrantedByUpn gets the grantedByUpn property value. The Upn of the user that granted consent for this account
func (m *DataSharingConsent) GetGrantedByUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grantedByUpn
    }
}
// GetGrantedByUserId gets the grantedByUserId property value. The UserId of the user that granted consent for this account
func (m *DataSharingConsent) GetGrantedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.grantedByUserId
    }
}
// GetServiceDisplayName gets the serviceDisplayName property value. The display name of the service work flow
func (m *DataSharingConsent) GetServiceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceDisplayName
    }
}
// GetTermsUrl gets the termsUrl property value. The TermsUrl for the data sharing consent
func (m *DataSharingConsent) GetTermsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.termsUrl
    }
}
// Serialize serializes information the current object
func (m *DataSharingConsent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetGrantDateTime sets the grantDateTime property value. The time consent was granted for this account
func (m *DataSharingConsent) SetGrantDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.grantDateTime = value
    }
}
// SetGranted sets the granted property value. The granted state for the data sharing consent
func (m *DataSharingConsent) SetGranted(value *bool)() {
    if m != nil {
        m.granted = value
    }
}
// SetGrantedByUpn sets the grantedByUpn property value. The Upn of the user that granted consent for this account
func (m *DataSharingConsent) SetGrantedByUpn(value *string)() {
    if m != nil {
        m.grantedByUpn = value
    }
}
// SetGrantedByUserId sets the grantedByUserId property value. The UserId of the user that granted consent for this account
func (m *DataSharingConsent) SetGrantedByUserId(value *string)() {
    if m != nil {
        m.grantedByUserId = value
    }
}
// SetServiceDisplayName sets the serviceDisplayName property value. The display name of the service work flow
func (m *DataSharingConsent) SetServiceDisplayName(value *string)() {
    if m != nil {
        m.serviceDisplayName = value
    }
}
// SetTermsUrl sets the termsUrl property value. The TermsUrl for the data sharing consent
func (m *DataSharingConsent) SetTermsUrl(value *string)() {
    if m != nil {
        m.termsUrl = value
    }
}
