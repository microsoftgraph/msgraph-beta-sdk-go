package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// LocalConnectivityConfiguration 
type LocalConnectivityConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewLocalConnectivityConfiguration instantiates a new localConnectivityConfiguration and sets the default values.
func NewLocalConnectivityConfiguration()(*LocalConnectivityConfiguration) {
    m := &LocalConnectivityConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLocalConnectivityConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocalConnectivityConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocalConnectivityConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocalConnectivityConfiguration) GetAdditionalData()(map[string]any) {
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
// GetAsn gets the asn property value. The asn property
func (m *LocalConnectivityConfiguration) GetAsn()(*int32) {
    val, err := m.GetBackingStore().Get("asn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *LocalConnectivityConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBgpAddress gets the bgpAddress property value. The bgpAddress property
func (m *LocalConnectivityConfiguration) GetBgpAddress()(*string) {
    val, err := m.GetBackingStore().Get("bgpAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndpoint gets the endpoint property value. The endpoint property
func (m *LocalConnectivityConfiguration) GetEndpoint()(*string) {
    val, err := m.GetBackingStore().Get("endpoint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocalConnectivityConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["asn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAsn(val)
        }
        return nil
    }
    res["bgpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBgpAddress(val)
        }
        return nil
    }
    res["endpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpoint(val)
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
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val.(*Region))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LocalConnectivityConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRegion gets the region property value. The region property
func (m *LocalConnectivityConfiguration) GetRegion()(*Region) {
    val, err := m.GetBackingStore().Get("region")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Region)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LocalConnectivityConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("asn", m.GetAsn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bgpAddress", m.GetBgpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("endpoint", m.GetEndpoint())
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
    if m.GetRegion() != nil {
        cast := (*m.GetRegion()).String()
        err := writer.WriteStringValue("region", &cast)
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
func (m *LocalConnectivityConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAsn sets the asn property value. The asn property
func (m *LocalConnectivityConfiguration) SetAsn(value *int32)() {
    err := m.GetBackingStore().Set("asn", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *LocalConnectivityConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBgpAddress sets the bgpAddress property value. The bgpAddress property
func (m *LocalConnectivityConfiguration) SetBgpAddress(value *string)() {
    err := m.GetBackingStore().Set("bgpAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetEndpoint sets the endpoint property value. The endpoint property
func (m *LocalConnectivityConfiguration) SetEndpoint(value *string)() {
    err := m.GetBackingStore().Set("endpoint", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LocalConnectivityConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRegion sets the region property value. The region property
func (m *LocalConnectivityConfiguration) SetRegion(value *Region)() {
    err := m.GetBackingStore().Set("region", value)
    if err != nil {
        panic(err)
    }
}
// LocalConnectivityConfigurationable 
type LocalConnectivityConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAsn()(*int32)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBgpAddress()(*string)
    GetEndpoint()(*string)
    GetOdataType()(*string)
    GetRegion()(*Region)
    SetAsn(value *int32)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBgpAddress(value *string)()
    SetEndpoint(value *string)()
    SetOdataType(value *string)()
    SetRegion(value *Region)()
}
