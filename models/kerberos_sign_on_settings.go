package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// KerberosSignOnSettings 
type KerberosSignOnSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewKerberosSignOnSettings instantiates a new kerberosSignOnSettings and sets the default values.
func NewKerberosSignOnSettings()(*KerberosSignOnSettings) {
    m := &KerberosSignOnSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateKerberosSignOnSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKerberosSignOnSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKerberosSignOnSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KerberosSignOnSettings) GetAdditionalData()(map[string]any) {
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
func (m *KerberosSignOnSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KerberosSignOnSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["kerberosServicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosServicePrincipalName(val)
        }
        return nil
    }
    res["kerberosSignOnMappingAttributeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKerberosSignOnMappingAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnMappingAttributeType(val.(*KerberosSignOnMappingAttributeType))
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
// GetKerberosServicePrincipalName gets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) GetKerberosServicePrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("kerberosServicePrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKerberosSignOnMappingAttributeType gets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType) {
    val, err := m.GetBackingStore().Get("kerberosSignOnMappingAttributeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KerberosSignOnMappingAttributeType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *KerberosSignOnSettings) GetOdataType()(*string) {
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
func (m *KerberosSignOnSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("kerberosServicePrincipalName", m.GetKerberosServicePrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetKerberosSignOnMappingAttributeType() != nil {
        cast := (*m.GetKerberosSignOnMappingAttributeType()).String()
        err := writer.WriteStringValue("kerberosSignOnMappingAttributeType", &cast)
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
func (m *KerberosSignOnSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *KerberosSignOnSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetKerberosServicePrincipalName sets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) SetKerberosServicePrincipalName(value *string)() {
    err := m.GetBackingStore().Set("kerberosServicePrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetKerberosSignOnMappingAttributeType sets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)() {
    err := m.GetBackingStore().Set("kerberosSignOnMappingAttributeType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *KerberosSignOnSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// KerberosSignOnSettingsable 
type KerberosSignOnSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetKerberosServicePrincipalName()(*string)
    GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetKerberosServicePrincipalName(value *string)()
    SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)()
    SetOdataType(value *string)()
}
