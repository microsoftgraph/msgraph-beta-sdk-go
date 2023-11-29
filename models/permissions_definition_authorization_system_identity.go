package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PermissionsDefinitionAuthorizationSystemIdentity 
type PermissionsDefinitionAuthorizationSystemIdentity struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPermissionsDefinitionAuthorizationSystemIdentity instantiates a new permissionsDefinitionAuthorizationSystemIdentity and sets the default values.
func NewPermissionsDefinitionAuthorizationSystemIdentity()(*PermissionsDefinitionAuthorizationSystemIdentity) {
    m := &PermissionsDefinitionAuthorizationSystemIdentity{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePermissionsDefinitionAuthorizationSystemIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionsDefinitionAuthorizationSystemIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionsDefinitionAuthorizationSystemIdentity(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PermissionsDefinitionAuthorizationSystemIdentity) GetAdditionalData()(map[string]any) {
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
func (m *PermissionsDefinitionAuthorizationSystemIdentity) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExternalId gets the externalId property value. Unique ID of the identity within the external system. Prefixed with rsn: if this is a SAML or ED user in AWS. Alternate key.
func (m *PermissionsDefinitionAuthorizationSystemIdentity) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionsDefinitionAuthorizationSystemIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["identityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionsDefinitionIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityType(val.(*PermissionsDefinitionIdentityType))
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsDefinitionIdentitySourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(PermissionsDefinitionIdentitySourceable))
        }
        return nil
    }
    return res
}
// GetIdentityType gets the identityType property value. The identityType property
func (m *PermissionsDefinitionAuthorizationSystemIdentity) GetIdentityType()(*PermissionsDefinitionIdentityType) {
    val, err := m.GetBackingStore().Get("identityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PermissionsDefinitionIdentityType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PermissionsDefinitionAuthorizationSystemIdentity) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSource gets the source property value. The source property
func (m *PermissionsDefinitionAuthorizationSystemIdentity) GetSource()(PermissionsDefinitionIdentitySourceable) {
    val, err := m.GetBackingStore().Get("source")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsDefinitionIdentitySourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PermissionsDefinitionAuthorizationSystemIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetIdentityType() != nil {
        cast := (*m.GetIdentityType()).String()
        err := writer.WriteStringValue("identityType", &cast)
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
        err := writer.WriteObjectValue("source", m.GetSource())
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
func (m *PermissionsDefinitionAuthorizationSystemIdentity) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PermissionsDefinitionAuthorizationSystemIdentity) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExternalId sets the externalId property value. Unique ID of the identity within the external system. Prefixed with rsn: if this is a SAML or ED user in AWS. Alternate key.
func (m *PermissionsDefinitionAuthorizationSystemIdentity) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityType sets the identityType property value. The identityType property
func (m *PermissionsDefinitionAuthorizationSystemIdentity) SetIdentityType(value *PermissionsDefinitionIdentityType)() {
    err := m.GetBackingStore().Set("identityType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PermissionsDefinitionAuthorizationSystemIdentity) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSource sets the source property value. The source property
func (m *PermissionsDefinitionAuthorizationSystemIdentity) SetSource(value PermissionsDefinitionIdentitySourceable)() {
    err := m.GetBackingStore().Set("source", value)
    if err != nil {
        panic(err)
    }
}
// PermissionsDefinitionAuthorizationSystemIdentityable 
type PermissionsDefinitionAuthorizationSystemIdentityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExternalId()(*string)
    GetIdentityType()(*PermissionsDefinitionIdentityType)
    GetOdataType()(*string)
    GetSource()(PermissionsDefinitionIdentitySourceable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExternalId(value *string)()
    SetIdentityType(value *PermissionsDefinitionIdentityType)()
    SetOdataType(value *string)()
    SetSource(value PermissionsDefinitionIdentitySourceable)()
}
