package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OnPremisesPublishingSingleSignOn 
type OnPremisesPublishingSingleSignOn struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnPremisesPublishingSingleSignOn instantiates a new onPremisesPublishingSingleSignOn and sets the default values.
func NewOnPremisesPublishingSingleSignOn()(*OnPremisesPublishingSingleSignOn) {
    m := &OnPremisesPublishingSingleSignOn{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesPublishingSingleSignOnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesPublishingSingleSignOn(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishingSingleSignOn) GetAdditionalData()(map[string]any) {
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
func (m *OnPremisesPublishingSingleSignOn) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishingSingleSignOn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["kerberosSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKerberosSignOnSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnSettings(val.(KerberosSignOnSettingsable))
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
    res["singleSignOnMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnPremisesPublishingSingleSignOn_singleSignOnMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnMode(val.(*OnPremisesPublishingSingleSignOn_singleSignOnMode))
        }
        return nil
    }
    return res
}
// GetKerberosSignOnSettings gets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) GetKerberosSignOnSettings()(KerberosSignOnSettingsable) {
    val, err := m.GetBackingStore().Get("kerberosSignOnSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(KerberosSignOnSettingsable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesPublishingSingleSignOn) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSingleSignOnMode gets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased, oAuthToken.
func (m *OnPremisesPublishingSingleSignOn) GetSingleSignOnMode()(*OnPremisesPublishingSingleSignOn_singleSignOnMode) {
    val, err := m.GetBackingStore().Get("singleSignOnMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OnPremisesPublishingSingleSignOn_singleSignOnMode)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnPremisesPublishingSingleSignOn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("kerberosSignOnSettings", m.GetKerberosSignOnSettings())
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
    if m.GetSingleSignOnMode() != nil {
        cast := (*m.GetSingleSignOnMode()).String()
        err := writer.WriteStringValue("singleSignOnMode", &cast)
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
func (m *OnPremisesPublishingSingleSignOn) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OnPremisesPublishingSingleSignOn) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetKerberosSignOnSettings sets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) SetKerberosSignOnSettings(value KerberosSignOnSettingsable)() {
    err := m.GetBackingStore().Set("kerberosSignOnSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesPublishingSingleSignOn) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnMode sets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased, oAuthToken.
func (m *OnPremisesPublishingSingleSignOn) SetSingleSignOnMode(value *OnPremisesPublishingSingleSignOn_singleSignOnMode)() {
    err := m.GetBackingStore().Set("singleSignOnMode", value)
    if err != nil {
        panic(err)
    }
}
// OnPremisesPublishingSingleSignOnable 
type OnPremisesPublishingSingleSignOnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetKerberosSignOnSettings()(KerberosSignOnSettingsable)
    GetOdataType()(*string)
    GetSingleSignOnMode()(*OnPremisesPublishingSingleSignOn_singleSignOnMode)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetKerberosSignOnSettings(value KerberosSignOnSettingsable)()
    SetOdataType(value *string)()
    SetSingleSignOnMode(value *OnPremisesPublishingSingleSignOn_singleSignOnMode)()
}
