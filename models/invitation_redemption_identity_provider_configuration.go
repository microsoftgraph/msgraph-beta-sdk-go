package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// InvitationRedemptionIdentityProviderConfiguration 
type InvitationRedemptionIdentityProviderConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewInvitationRedemptionIdentityProviderConfiguration instantiates a new invitationRedemptionIdentityProviderConfiguration and sets the default values.
func NewInvitationRedemptionIdentityProviderConfiguration()(*InvitationRedemptionIdentityProviderConfiguration) {
    m := &InvitationRedemptionIdentityProviderConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInvitationRedemptionIdentityProviderConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInvitationRedemptionIdentityProviderConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.defaultInvitationRedemptionIdentityProviderConfiguration":
                        return NewDefaultInvitationRedemptionIdentityProviderConfiguration(), nil
                }
            }
        }
    }
    return NewInvitationRedemptionIdentityProviderConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InvitationRedemptionIdentityProviderConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *InvitationRedemptionIdentityProviderConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFallbackIdentityProvider gets the fallbackIdentityProvider property value. The fallback identity provider to be used in case no primary identity provider can be used for guest invitation redemption. Possible values are: defaultConfiguredIdp, emailOneTimePasscode, or microsoftAccount.
func (m *InvitationRedemptionIdentityProviderConfiguration) GetFallbackIdentityProvider()(*InvitationRedemptionIdentityProviderConfiguration_fallbackIdentityProvider) {
    val, err := m.GetBackingStore().Get("fallbackIdentityProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*InvitationRedemptionIdentityProviderConfiguration_fallbackIdentityProvider)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InvitationRedemptionIdentityProviderConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fallbackIdentityProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInvitationRedemptionIdentityProviderConfiguration_fallbackIdentityProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFallbackIdentityProvider(val.(*InvitationRedemptionIdentityProviderConfiguration_fallbackIdentityProvider))
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
    res["primaryIdentityProviderPrecedenceOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseInvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*InvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder))
                }
            }
            m.SetPrimaryIdentityProviderPrecedenceOrder(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InvitationRedemptionIdentityProviderConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrimaryIdentityProviderPrecedenceOrder gets the primaryIdentityProviderPrecedenceOrder property value. Collection of identity providers in priority order of preference to be used for guest invitation redemption. Possible values are: azureActiveDirectory, externalFederation, or socialIdentityProviders.
func (m *InvitationRedemptionIdentityProviderConfiguration) GetPrimaryIdentityProviderPrecedenceOrder()([]InvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder) {
    val, err := m.GetBackingStore().Get("primaryIdentityProviderPrecedenceOrder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]InvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InvitationRedemptionIdentityProviderConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFallbackIdentityProvider() != nil {
        cast := (*m.GetFallbackIdentityProvider()).String()
        err := writer.WriteStringValue("fallbackIdentityProvider", &cast)
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
    if m.GetPrimaryIdentityProviderPrecedenceOrder() != nil {
        err := writer.WriteCollectionOfStringValues("primaryIdentityProviderPrecedenceOrder", SerializeInvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder(m.GetPrimaryIdentityProviderPrecedenceOrder()))
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
func (m *InvitationRedemptionIdentityProviderConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *InvitationRedemptionIdentityProviderConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFallbackIdentityProvider sets the fallbackIdentityProvider property value. The fallback identity provider to be used in case no primary identity provider can be used for guest invitation redemption. Possible values are: defaultConfiguredIdp, emailOneTimePasscode, or microsoftAccount.
func (m *InvitationRedemptionIdentityProviderConfiguration) SetFallbackIdentityProvider(value *InvitationRedemptionIdentityProviderConfiguration_fallbackIdentityProvider)() {
    err := m.GetBackingStore().Set("fallbackIdentityProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InvitationRedemptionIdentityProviderConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryIdentityProviderPrecedenceOrder sets the primaryIdentityProviderPrecedenceOrder property value. Collection of identity providers in priority order of preference to be used for guest invitation redemption. Possible values are: azureActiveDirectory, externalFederation, or socialIdentityProviders.
func (m *InvitationRedemptionIdentityProviderConfiguration) SetPrimaryIdentityProviderPrecedenceOrder(value []InvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder)() {
    err := m.GetBackingStore().Set("primaryIdentityProviderPrecedenceOrder", value)
    if err != nil {
        panic(err)
    }
}
// InvitationRedemptionIdentityProviderConfigurationable 
type InvitationRedemptionIdentityProviderConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFallbackIdentityProvider()(*InvitationRedemptionIdentityProviderConfiguration_fallbackIdentityProvider)
    GetOdataType()(*string)
    GetPrimaryIdentityProviderPrecedenceOrder()([]InvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFallbackIdentityProvider(value *InvitationRedemptionIdentityProviderConfiguration_fallbackIdentityProvider)()
    SetOdataType(value *string)()
    SetPrimaryIdentityProviderPrecedenceOrder(value []InvitationRedemptionIdentityProviderConfiguration_primaryIdentityProviderPrecedenceOrder)()
}
