package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AdminConsent admin consent information.
type AdminConsent struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAdminConsent instantiates a new adminConsent and sets the default values.
func NewAdminConsent()(*AdminConsent) {
    m := &AdminConsent{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAdminConsentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminConsentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminConsent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminConsent) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AdminConsent) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminConsent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["shareAPNSData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdminConsentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareAPNSData(val.(*AdminConsentState))
        }
        return nil
    }
    res["shareUserExperienceAnalyticsData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdminConsentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareUserExperienceAnalyticsData(val.(*AdminConsentState))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AdminConsent) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetShareAPNSData gets the shareAPNSData property value. Admin consent state.
func (m *AdminConsent) GetShareAPNSData()(*AdminConsentState) {
    val, err := m.GetBackingStore().Get("shareAPNSData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AdminConsentState)
    }
    return nil
}
// GetShareUserExperienceAnalyticsData gets the shareUserExperienceAnalyticsData property value. Admin consent state.
func (m *AdminConsent) GetShareUserExperienceAnalyticsData()(*AdminConsentState) {
    val, err := m.GetBackingStore().Get("shareUserExperienceAnalyticsData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AdminConsentState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AdminConsent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetShareAPNSData() != nil {
        cast := (*m.GetShareAPNSData()).String()
        err := writer.WriteStringValue("shareAPNSData", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetShareUserExperienceAnalyticsData() != nil {
        cast := (*m.GetShareUserExperienceAnalyticsData()).String()
        err := writer.WriteStringValue("shareUserExperienceAnalyticsData", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminConsent) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AdminConsent) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AdminConsent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetShareAPNSData sets the shareAPNSData property value. Admin consent state.
func (m *AdminConsent) SetShareAPNSData(value *AdminConsentState)() {
    err := m.GetBackingStore().Set("shareAPNSData", value)
    if err != nil {
        panic(err)
    }
}
// SetShareUserExperienceAnalyticsData sets the shareUserExperienceAnalyticsData property value. Admin consent state.
func (m *AdminConsent) SetShareUserExperienceAnalyticsData(value *AdminConsentState)() {
    err := m.GetBackingStore().Set("shareUserExperienceAnalyticsData", value)
    if err != nil {
        panic(err)
    }
}
// AdminConsentable 
type AdminConsentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetShareAPNSData()(*AdminConsentState)
    GetShareUserExperienceAnalyticsData()(*AdminConsentState)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetShareAPNSData(value *AdminConsentState)()
    SetShareUserExperienceAnalyticsData(value *AdminConsentState)()
}
