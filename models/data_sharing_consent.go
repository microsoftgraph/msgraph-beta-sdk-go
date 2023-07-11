package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataSharingConsent data sharing consent information.
type DataSharingConsent struct {
    Entity
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    val, err := m.GetBackingStore().Get("grantDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetGranted gets the granted property value. The granted state for the data sharing consent
func (m *DataSharingConsent) GetGranted()(*bool) {
    val, err := m.GetBackingStore().Get("granted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetGrantedByUpn gets the grantedByUpn property value. The Upn of the user that granted consent for this account
func (m *DataSharingConsent) GetGrantedByUpn()(*string) {
    val, err := m.GetBackingStore().Get("grantedByUpn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetGrantedByUserId gets the grantedByUserId property value. The UserId of the user that granted consent for this account
func (m *DataSharingConsent) GetGrantedByUserId()(*string) {
    val, err := m.GetBackingStore().Get("grantedByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DataSharingConsent) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServiceDisplayName gets the serviceDisplayName property value. The display name of the service work flow
func (m *DataSharingConsent) GetServiceDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("serviceDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTermsUrl gets the termsUrl property value. The TermsUrl for the data sharing consent
func (m *DataSharingConsent) GetTermsUrl()(*string) {
    val, err := m.GetBackingStore().Get("termsUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    err := m.GetBackingStore().Set("grantDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetGranted sets the granted property value. The granted state for the data sharing consent
func (m *DataSharingConsent) SetGranted(value *bool)() {
    err := m.GetBackingStore().Set("granted", value)
    if err != nil {
        panic(err)
    }
}
// SetGrantedByUpn sets the grantedByUpn property value. The Upn of the user that granted consent for this account
func (m *DataSharingConsent) SetGrantedByUpn(value *string)() {
    err := m.GetBackingStore().Set("grantedByUpn", value)
    if err != nil {
        panic(err)
    }
}
// SetGrantedByUserId sets the grantedByUserId property value. The UserId of the user that granted consent for this account
func (m *DataSharingConsent) SetGrantedByUserId(value *string)() {
    err := m.GetBackingStore().Set("grantedByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DataSharingConsent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceDisplayName sets the serviceDisplayName property value. The display name of the service work flow
func (m *DataSharingConsent) SetServiceDisplayName(value *string)() {
    err := m.GetBackingStore().Set("serviceDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTermsUrl sets the termsUrl property value. The TermsUrl for the data sharing consent
func (m *DataSharingConsent) SetTermsUrl(value *string)() {
    err := m.GetBackingStore().Set("termsUrl", value)
    if err != nil {
        panic(err)
    }
}
// DataSharingConsentable 
type DataSharingConsentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGrantDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGranted()(*bool)
    GetGrantedByUpn()(*string)
    GetGrantedByUserId()(*string)
    GetOdataType()(*string)
    GetServiceDisplayName()(*string)
    GetTermsUrl()(*string)
    SetGrantDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGranted(value *bool)()
    SetGrantedByUpn(value *string)()
    SetGrantedByUserId(value *string)()
    SetOdataType(value *string)()
    SetServiceDisplayName(value *string)()
    SetTermsUrl(value *string)()
}
