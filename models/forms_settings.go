package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// FormsSettings 
type FormsSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewFormsSettings instantiates a new formsSettings and sets the default values.
func NewFormsSettings()(*FormsSettings) {
    m := &FormsSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFormsSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFormsSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFormsSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FormsSettings) GetAdditionalData()(map[string]any) {
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
func (m *FormsSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FormsSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isBingImageSearchEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBingImageSearchEnabled(val)
        }
        return nil
    }
    res["isExternalSendFormEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternalSendFormEnabled(val)
        }
        return nil
    }
    res["isExternalShareCollaborationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternalShareCollaborationEnabled(val)
        }
        return nil
    }
    res["isExternalShareResultEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternalShareResultEnabled(val)
        }
        return nil
    }
    res["isExternalShareTemplateEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExternalShareTemplateEnabled(val)
        }
        return nil
    }
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
// GetIsBingImageSearchEnabled gets the isBingImageSearchEnabled property value. Controls whether users can add images from Bing to forms.
func (m *FormsSettings) GetIsBingImageSearchEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isBingImageSearchEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsExternalSendFormEnabled gets the isExternalSendFormEnabled property value. Controls whether users can send a link to a form to an external user.
func (m *FormsSettings) GetIsExternalSendFormEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isExternalSendFormEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsExternalShareCollaborationEnabled gets the isExternalShareCollaborationEnabled property value. Controls whether users can collaborate on a form layout and structure with an external user.
func (m *FormsSettings) GetIsExternalShareCollaborationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isExternalShareCollaborationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsExternalShareResultEnabled gets the isExternalShareResultEnabled property value. Controls whether users can share form results with external users.
func (m *FormsSettings) GetIsExternalShareResultEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isExternalShareResultEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsExternalShareTemplateEnabled gets the isExternalShareTemplateEnabled property value. Controls whether users can share form templates with external users.
func (m *FormsSettings) GetIsExternalShareTemplateEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isExternalShareTemplateEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsInOrgFormsPhishingScanEnabled gets the isInOrgFormsPhishingScanEnabled property value. Controls whether phishing protection is run on forms created by users, blocking the creation of forms if common phishing questions are detected.
func (m *FormsSettings) GetIsInOrgFormsPhishingScanEnabled()(*bool) {
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
func (m *FormsSettings) GetIsRecordIdentityByDefaultEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isRecordIdentityByDefaultEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *FormsSettings) GetOdataType()(*string) {
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
func (m *FormsSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isBingImageSearchEnabled", m.GetIsBingImageSearchEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isExternalSendFormEnabled", m.GetIsExternalSendFormEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isExternalShareCollaborationEnabled", m.GetIsExternalShareCollaborationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isExternalShareResultEnabled", m.GetIsExternalShareResultEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isExternalShareTemplateEnabled", m.GetIsExternalShareTemplateEnabled())
        if err != nil {
            return err
        }
    }
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
func (m *FormsSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *FormsSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsBingImageSearchEnabled sets the isBingImageSearchEnabled property value. Controls whether users can add images from Bing to forms.
func (m *FormsSettings) SetIsBingImageSearchEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isBingImageSearchEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExternalSendFormEnabled sets the isExternalSendFormEnabled property value. Controls whether users can send a link to a form to an external user.
func (m *FormsSettings) SetIsExternalSendFormEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isExternalSendFormEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExternalShareCollaborationEnabled sets the isExternalShareCollaborationEnabled property value. Controls whether users can collaborate on a form layout and structure with an external user.
func (m *FormsSettings) SetIsExternalShareCollaborationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isExternalShareCollaborationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExternalShareResultEnabled sets the isExternalShareResultEnabled property value. Controls whether users can share form results with external users.
func (m *FormsSettings) SetIsExternalShareResultEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isExternalShareResultEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExternalShareTemplateEnabled sets the isExternalShareTemplateEnabled property value. Controls whether users can share form templates with external users.
func (m *FormsSettings) SetIsExternalShareTemplateEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isExternalShareTemplateEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsInOrgFormsPhishingScanEnabled sets the isInOrgFormsPhishingScanEnabled property value. Controls whether phishing protection is run on forms created by users, blocking the creation of forms if common phishing questions are detected.
func (m *FormsSettings) SetIsInOrgFormsPhishingScanEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isInOrgFormsPhishingScanEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRecordIdentityByDefaultEnabled sets the isRecordIdentityByDefaultEnabled property value. Controls whether the names of users who fill out forms are recorded.
func (m *FormsSettings) SetIsRecordIdentityByDefaultEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isRecordIdentityByDefaultEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FormsSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// FormsSettingsable 
type FormsSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsBingImageSearchEnabled()(*bool)
    GetIsExternalSendFormEnabled()(*bool)
    GetIsExternalShareCollaborationEnabled()(*bool)
    GetIsExternalShareResultEnabled()(*bool)
    GetIsExternalShareTemplateEnabled()(*bool)
    GetIsInOrgFormsPhishingScanEnabled()(*bool)
    GetIsRecordIdentityByDefaultEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsBingImageSearchEnabled(value *bool)()
    SetIsExternalSendFormEnabled(value *bool)()
    SetIsExternalShareCollaborationEnabled(value *bool)()
    SetIsExternalShareResultEnabled(value *bool)()
    SetIsExternalShareTemplateEnabled(value *bool)()
    SetIsInOrgFormsPhishingScanEnabled(value *bool)()
    SetIsRecordIdentityByDefaultEnabled(value *bool)()
    SetOdataType(value *string)()
}
