// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type EducationAiFeedbackContentSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEducationAiFeedbackContentSettings instantiates a new EducationAiFeedbackContentSettings and sets the default values.
func NewEducationAiFeedbackContentSettings()(*EducationAiFeedbackContentSettings) {
    m := &EducationAiFeedbackContentSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationAiFeedbackContentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationAiFeedbackContentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAiFeedbackContentSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EducationAiFeedbackContentSettings) GetAdditionalData()(map[string]any) {
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
// returns a BackingStore when successful
func (m *EducationAiFeedbackContentSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationAiFeedbackContentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isMessageClarityEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMessageClarityEnabled(val)
        }
        return nil
    }
    res["isQualityOfInformationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsQualityOfInformationEnabled(val)
        }
        return nil
    }
    res["isSpeechOrganizationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSpeechOrganizationEnabled(val)
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
// GetIsMessageClarityEnabled gets the isMessageClarityEnabled property value. Indicates whether the student should receive feedback on their message clarity from the AI feedback.
// returns a *bool when successful
func (m *EducationAiFeedbackContentSettings) GetIsMessageClarityEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isMessageClarityEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsQualityOfInformationEnabled gets the isQualityOfInformationEnabled property value. Indicates whether the student should receive feedback on their quality of information from the AI feedback.
// returns a *bool when successful
func (m *EducationAiFeedbackContentSettings) GetIsQualityOfInformationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isQualityOfInformationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSpeechOrganizationEnabled gets the isSpeechOrganizationEnabled property value. Indicates whether the student should receive feedback on their speech organization from the AI feedback.
// returns a *bool when successful
func (m *EducationAiFeedbackContentSettings) GetIsSpeechOrganizationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSpeechOrganizationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *EducationAiFeedbackContentSettings) GetOdataType()(*string) {
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
func (m *EducationAiFeedbackContentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isMessageClarityEnabled", m.GetIsMessageClarityEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isQualityOfInformationEnabled", m.GetIsQualityOfInformationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSpeechOrganizationEnabled", m.GetIsSpeechOrganizationEnabled())
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
func (m *EducationAiFeedbackContentSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EducationAiFeedbackContentSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsMessageClarityEnabled sets the isMessageClarityEnabled property value. Indicates whether the student should receive feedback on their message clarity from the AI feedback.
func (m *EducationAiFeedbackContentSettings) SetIsMessageClarityEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isMessageClarityEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsQualityOfInformationEnabled sets the isQualityOfInformationEnabled property value. Indicates whether the student should receive feedback on their quality of information from the AI feedback.
func (m *EducationAiFeedbackContentSettings) SetIsQualityOfInformationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isQualityOfInformationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSpeechOrganizationEnabled sets the isSpeechOrganizationEnabled property value. Indicates whether the student should receive feedback on their speech organization from the AI feedback.
func (m *EducationAiFeedbackContentSettings) SetIsSpeechOrganizationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSpeechOrganizationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationAiFeedbackContentSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type EducationAiFeedbackContentSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsMessageClarityEnabled()(*bool)
    GetIsQualityOfInformationEnabled()(*bool)
    GetIsSpeechOrganizationEnabled()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsMessageClarityEnabled(value *bool)()
    SetIsQualityOfInformationEnabled(value *bool)()
    SetIsSpeechOrganizationEnabled(value *bool)()
    SetOdataType(value *string)()
}
