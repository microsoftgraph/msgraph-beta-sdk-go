package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ManagedIdentity 
type ManagedIdentity struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManagedIdentity instantiates a new managedIdentity and sets the default values.
func NewManagedIdentity()(*ManagedIdentity) {
    m := &ManagedIdentity{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedIdentity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedIdentity) GetAdditionalData()(map[string]any) {
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
// GetAssociatedResourceId gets the associatedResourceId property value. The ARM resource ID of the Azure resource associated with the managed identity used for sign in.
func (m *ManagedIdentity) GetAssociatedResourceId()(*string) {
    val, err := m.GetBackingStore().Get("associatedResourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ManagedIdentity) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFederatedTokenId gets the federatedTokenId property value. The unique ID of the federated token.
func (m *ManagedIdentity) GetFederatedTokenId()(*string) {
    val, err := m.GetBackingStore().Get("federatedTokenId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFederatedTokenIssuer gets the federatedTokenIssuer property value. The issuer of the federated token.
func (m *ManagedIdentity) GetFederatedTokenIssuer()(*string) {
    val, err := m.GetBackingStore().Get("federatedTokenIssuer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["associatedResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssociatedResourceId(val)
        }
        return nil
    }
    res["federatedTokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedTokenId(val)
        }
        return nil
    }
    res["federatedTokenIssuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedTokenIssuer(val)
        }
        return nil
    }
    res["msiType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMsiType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMsiType(val.(*MsiType))
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
// GetMsiType gets the msiType property value. The possible values are: none, userAssigned, systemAssigned, unknownFutureValue.
func (m *ManagedIdentity) GetMsiType()(*MsiType) {
    val, err := m.GetBackingStore().Get("msiType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MsiType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedIdentity) GetOdataType()(*string) {
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
func (m *ManagedIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("associatedResourceId", m.GetAssociatedResourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("federatedTokenId", m.GetFederatedTokenId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("federatedTokenIssuer", m.GetFederatedTokenIssuer())
        if err != nil {
            return err
        }
    }
    if m.GetMsiType() != nil {
        cast := (*m.GetMsiType()).String()
        err := writer.WriteStringValue("msiType", &cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedIdentity) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssociatedResourceId sets the associatedResourceId property value. The ARM resource ID of the Azure resource associated with the managed identity used for sign in.
func (m *ManagedIdentity) SetAssociatedResourceId(value *string)() {
    err := m.GetBackingStore().Set("associatedResourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ManagedIdentity) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFederatedTokenId sets the federatedTokenId property value. The unique ID of the federated token.
func (m *ManagedIdentity) SetFederatedTokenId(value *string)() {
    err := m.GetBackingStore().Set("federatedTokenId", value)
    if err != nil {
        panic(err)
    }
}
// SetFederatedTokenIssuer sets the federatedTokenIssuer property value. The issuer of the federated token.
func (m *ManagedIdentity) SetFederatedTokenIssuer(value *string)() {
    err := m.GetBackingStore().Set("federatedTokenIssuer", value)
    if err != nil {
        panic(err)
    }
}
// SetMsiType sets the msiType property value. The possible values are: none, userAssigned, systemAssigned, unknownFutureValue.
func (m *ManagedIdentity) SetMsiType(value *MsiType)() {
    err := m.GetBackingStore().Set("msiType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedIdentity) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// ManagedIdentityable 
type ManagedIdentityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssociatedResourceId()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFederatedTokenId()(*string)
    GetFederatedTokenIssuer()(*string)
    GetMsiType()(*MsiType)
    GetOdataType()(*string)
    SetAssociatedResourceId(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFederatedTokenId(value *string)()
    SetFederatedTokenIssuer(value *string)()
    SetMsiType(value *MsiType)()
    SetOdataType(value *string)()
}
