package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type PermissionsDefinition struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPermissionsDefinition instantiates a new PermissionsDefinition and sets the default values.
func NewPermissionsDefinition()(*PermissionsDefinition) {
    m := &PermissionsDefinition{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePermissionsDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePermissionsDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.awsPermissionsDefinition":
                        return NewAwsPermissionsDefinition(), nil
                    case "#microsoft.graph.singleResourceAzurePermissionsDefinition":
                        return NewSingleResourceAzurePermissionsDefinition(), nil
                    case "#microsoft.graph.singleResourceGcpPermissionsDefinition":
                        return NewSingleResourceGcpPermissionsDefinition(), nil
                }
            }
        }
    }
    return NewPermissionsDefinition(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PermissionsDefinition) GetAdditionalData()(map[string]any) {
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
// GetAuthorizationSystemInfo gets the authorizationSystemInfo property value. The authorizationSystemInfo property
// returns a PermissionsDefinitionAuthorizationSystemable when successful
func (m *PermissionsDefinition) GetAuthorizationSystemInfo()(PermissionsDefinitionAuthorizationSystemable) {
    val, err := m.GetBackingStore().Get("authorizationSystemInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsDefinitionAuthorizationSystemable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *PermissionsDefinition) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PermissionsDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authorizationSystemInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsDefinitionAuthorizationSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizationSystemInfo(val.(PermissionsDefinitionAuthorizationSystemable))
        }
        return nil
    }
    res["identityInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsDefinitionAuthorizationSystemIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityInfo(val.(PermissionsDefinitionAuthorizationSystemIdentityable))
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
// GetIdentityInfo gets the identityInfo property value. The identityInfo property
// returns a PermissionsDefinitionAuthorizationSystemIdentityable when successful
func (m *PermissionsDefinition) GetIdentityInfo()(PermissionsDefinitionAuthorizationSystemIdentityable) {
    val, err := m.GetBackingStore().Get("identityInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsDefinitionAuthorizationSystemIdentityable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *PermissionsDefinition) GetOdataType()(*string) {
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
func (m *PermissionsDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authorizationSystemInfo", m.GetAuthorizationSystemInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("identityInfo", m.GetIdentityInfo())
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
func (m *PermissionsDefinition) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizationSystemInfo sets the authorizationSystemInfo property value. The authorizationSystemInfo property
func (m *PermissionsDefinition) SetAuthorizationSystemInfo(value PermissionsDefinitionAuthorizationSystemable)() {
    err := m.GetBackingStore().Set("authorizationSystemInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PermissionsDefinition) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIdentityInfo sets the identityInfo property value. The identityInfo property
func (m *PermissionsDefinition) SetIdentityInfo(value PermissionsDefinitionAuthorizationSystemIdentityable)() {
    err := m.GetBackingStore().Set("identityInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PermissionsDefinition) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type PermissionsDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorizationSystemInfo()(PermissionsDefinitionAuthorizationSystemable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIdentityInfo()(PermissionsDefinitionAuthorizationSystemIdentityable)
    GetOdataType()(*string)
    SetAuthorizationSystemInfo(value PermissionsDefinitionAuthorizationSystemable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIdentityInfo(value PermissionsDefinitionAuthorizationSystemIdentityable)()
    SetOdataType(value *string)()
}
