package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CustomerVoiceSettings 
type CustomerVoiceSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCustomerVoiceSettings instantiates a new customerVoiceSettings and sets the default values.
func NewCustomerVoiceSettings()(*CustomerVoiceSettings) {
    m := &CustomerVoiceSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomerVoiceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomerVoiceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomerVoiceSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomerVoiceSettings) GetAdditionalData()(map[string]any) {
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
func (m *CustomerVoiceSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomerVoiceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isInOrgFormsPhishingScanEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInOrgFormsPhishingScanEnabled(val)
        }
        return nil
    }
    res["isRecordIdentityByDefaultEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRecordIdentityByDefaultEnabled(val)
        }
        return nil
    }
    res["isRestrictedSurveyAccessEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRestrictedSurveyAccessEnabled(val)
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
    return res
}
// GetIsInOrgFormsPhishingScanEnabled gets the isInOrgFormsPhishingScanEnabled property value. Controls whether phishing protection is run on forms created by users, blocking the creation of forms if common phishing questions are detected.
func (m *CustomerVoiceSettings) GetIsInOrgFormsPhishingScanEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isInOrgFormsPhishingScanEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsRecordIdentityByDefaultEnabled gets the isRecordIdentityByDefaultEnabled property value. Controls whether the names of users who fill out forms are recorded.
func (m *CustomerVoiceSettings) GetIsRecordIdentityByDefaultEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isRecordIdentityByDefaultEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsRestrictedSurveyAccessEnabled gets the isRestrictedSurveyAccessEnabled property value. Controls whether only users inside your organization can submit a response.
func (m *CustomerVoiceSettings) GetIsRestrictedSurveyAccessEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isRestrictedSurveyAccessEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomerVoiceSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomerVoiceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isInOrgFormsPhishingScanEnabled", m.GetIsInOrgFormsPhishingScanEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRecordIdentityByDefaultEnabled", m.GetIsRecordIdentityByDefaultEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRestrictedSurveyAccessEnabled", m.GetIsRestrictedSurveyAccessEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *CustomerVoiceSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CustomerVoiceSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsInOrgFormsPhishingScanEnabled sets the isInOrgFormsPhishingScanEnabled property value. Controls whether phishing protection is run on forms created by users, blocking the creation of forms if common phishing questions are detected.
func (m *CustomerVoiceSettings) SetIsInOrgFormsPhishingScanEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isInOrgFormsPhishingScanEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRecordIdentityByDefaultEnabled sets the isRecordIdentityByDefaultEnabled property value. Controls whether the names of users who fill out forms are recorded.
func (m *CustomerVoiceSettings) SetIsRecordIdentityByDefaultEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isRecordIdentityByDefaultEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRestrictedSurveyAccessEnabled sets the isRestrictedSurveyAccessEnabled property value. Controls whether only users inside your organization can submit a response.
func (m *CustomerVoiceSettings) SetIsRestrictedSurveyAccessEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isRestrictedSurveyAccessEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomerVoiceSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// CustomerVoiceSettingsable 
type CustomerVoiceSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsInOrgFormsPhishingScanEnabled()(*bool)
    GetIsRecordIdentityByDefaultEnabled()(*bool)
    GetIsRestrictedSurveyAccessEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsInOrgFormsPhishingScanEnabled(value *bool)()
    SetIsRecordIdentityByDefaultEnabled(value *bool)()
    SetIsRestrictedSurveyAccessEnabled(value *bool)()
    SetOdataType(value *string)()
}
