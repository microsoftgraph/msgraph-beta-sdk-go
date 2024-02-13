package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type RemoteNetworkConnectivityConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRemoteNetworkConnectivityConfiguration instantiates a new RemoteNetworkConnectivityConfiguration and sets the default values.
func NewRemoteNetworkConnectivityConfiguration()(*RemoteNetworkConnectivityConfiguration) {
    m := &RemoteNetworkConnectivityConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRemoteNetworkConnectivityConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRemoteNetworkConnectivityConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteNetworkConnectivityConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RemoteNetworkConnectivityConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *RemoteNetworkConnectivityConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RemoteNetworkConnectivityConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectivityConfigurationLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectivityConfigurationLinkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConnectivityConfigurationLinkable)
                }
            }
            m.SetLinks(res)
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
    res["remoteNetworkId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteNetworkId(val)
        }
        return nil
    }
    res["remoteNetworkName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteNetworkName(val)
        }
        return nil
    }
    return res
}
// GetLinks gets the links property value. The links property
// returns a []ConnectivityConfigurationLinkable when successful
func (m *RemoteNetworkConnectivityConfiguration) GetLinks()([]ConnectivityConfigurationLinkable) {
    val, err := m.GetBackingStore().Get("links")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConnectivityConfigurationLinkable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *RemoteNetworkConnectivityConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemoteNetworkId gets the remoteNetworkId property value. The remoteNetworkId property
// returns a *string when successful
func (m *RemoteNetworkConnectivityConfiguration) GetRemoteNetworkId()(*string) {
    val, err := m.GetBackingStore().Get("remoteNetworkId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemoteNetworkName gets the remoteNetworkName property value. The remoteNetworkName property
// returns a *string when successful
func (m *RemoteNetworkConnectivityConfiguration) GetRemoteNetworkName()(*string) {
    val, err := m.GetBackingStore().Get("remoteNetworkName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RemoteNetworkConnectivityConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
        for i, v := range m.GetLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("links", cast)
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
        err := writer.WriteStringValue("remoteNetworkId", m.GetRemoteNetworkId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remoteNetworkName", m.GetRemoteNetworkName())
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
func (m *RemoteNetworkConnectivityConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *RemoteNetworkConnectivityConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetLinks sets the links property value. The links property
func (m *RemoteNetworkConnectivityConfiguration) SetLinks(value []ConnectivityConfigurationLinkable)() {
    err := m.GetBackingStore().Set("links", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RemoteNetworkConnectivityConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteNetworkId sets the remoteNetworkId property value. The remoteNetworkId property
func (m *RemoteNetworkConnectivityConfiguration) SetRemoteNetworkId(value *string)() {
    err := m.GetBackingStore().Set("remoteNetworkId", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteNetworkName sets the remoteNetworkName property value. The remoteNetworkName property
func (m *RemoteNetworkConnectivityConfiguration) SetRemoteNetworkName(value *string)() {
    err := m.GetBackingStore().Set("remoteNetworkName", value)
    if err != nil {
        panic(err)
    }
}
type RemoteNetworkConnectivityConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetLinks()([]ConnectivityConfigurationLinkable)
    GetOdataType()(*string)
    GetRemoteNetworkId()(*string)
    GetRemoteNetworkName()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetLinks(value []ConnectivityConfigurationLinkable)()
    SetOdataType(value *string)()
    SetRemoteNetworkId(value *string)()
    SetRemoteNetworkName(value *string)()
}
