package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// VpnTrafficRule vPN Traffic Rule definition.
type VpnTrafficRule struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewVpnTrafficRule instantiates a new VpnTrafficRule and sets the default values.
func NewVpnTrafficRule()(*VpnTrafficRule) {
    m := &VpnTrafficRule{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVpnTrafficRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVpnTrafficRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnTrafficRule(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VpnTrafficRule) GetAdditionalData()(map[string]any) {
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
// GetAppId gets the appId property value. App identifier, if this traffic rule is triggered by an app.
// returns a *string when successful
func (m *VpnTrafficRule) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppType gets the appType property value. Indicates the type of app that a VPN traffic rule is associated with.
// returns a *VpnTrafficRuleAppType when successful
func (m *VpnTrafficRule) GetAppType()(*VpnTrafficRuleAppType) {
    val, err := m.GetBackingStore().Get("appType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnTrafficRuleAppType)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *VpnTrafficRule) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetClaims gets the claims property value. Claims associated with this traffic rule.
// returns a *string when successful
func (m *VpnTrafficRule) GetClaims()(*string) {
    val, err := m.GetBackingStore().Get("claims")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VpnTrafficRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["appType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnTrafficRuleAppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppType(val.(*VpnTrafficRuleAppType))
        }
        return nil
    }
    res["claims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClaims(val)
        }
        return nil
    }
    res["localAddressRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIPv4RangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IPv4Rangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IPv4Rangeable)
                }
            }
            m.SetLocalAddressRanges(res)
        }
        return nil
    }
    res["localPortRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNumberRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NumberRangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NumberRangeable)
                }
            }
            m.SetLocalPortRanges(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["protocols"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocols(val)
        }
        return nil
    }
    res["remoteAddressRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIPv4RangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IPv4Rangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IPv4Rangeable)
                }
            }
            m.SetRemoteAddressRanges(res)
        }
        return nil
    }
    res["remotePortRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNumberRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NumberRangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NumberRangeable)
                }
            }
            m.SetRemotePortRanges(res)
        }
        return nil
    }
    res["routingPolicyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnTrafficRuleRoutingPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoutingPolicyType(val.(*VpnTrafficRuleRoutingPolicyType))
        }
        return nil
    }
    res["vpnTrafficDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnTrafficDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnTrafficDirection(val.(*VpnTrafficDirection))
        }
        return nil
    }
    return res
}
// GetLocalAddressRanges gets the localAddressRanges property value. Local address range. This collection can contain a maximum of 500 elements.
// returns a []IPv4Rangeable when successful
func (m *VpnTrafficRule) GetLocalAddressRanges()([]IPv4Rangeable) {
    val, err := m.GetBackingStore().Get("localAddressRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IPv4Rangeable)
    }
    return nil
}
// GetLocalPortRanges gets the localPortRanges property value. Local port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
// returns a []NumberRangeable when successful
func (m *VpnTrafficRule) GetLocalPortRanges()([]NumberRangeable) {
    val, err := m.GetBackingStore().Get("localPortRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NumberRangeable)
    }
    return nil
}
// GetName gets the name property value. Name.
// returns a *string when successful
func (m *VpnTrafficRule) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *VpnTrafficRule) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProtocols gets the protocols property value. Protocols (0-255). Valid values 0 to 255
// returns a *int32 when successful
func (m *VpnTrafficRule) GetProtocols()(*int32) {
    val, err := m.GetBackingStore().Get("protocols")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRemoteAddressRanges gets the remoteAddressRanges property value. Remote address range. This collection can contain a maximum of 500 elements.
// returns a []IPv4Rangeable when successful
func (m *VpnTrafficRule) GetRemoteAddressRanges()([]IPv4Rangeable) {
    val, err := m.GetBackingStore().Get("remoteAddressRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IPv4Rangeable)
    }
    return nil
}
// GetRemotePortRanges gets the remotePortRanges property value. Remote port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
// returns a []NumberRangeable when successful
func (m *VpnTrafficRule) GetRemotePortRanges()([]NumberRangeable) {
    val, err := m.GetBackingStore().Get("remotePortRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NumberRangeable)
    }
    return nil
}
// GetRoutingPolicyType gets the routingPolicyType property value. Specifies the routing policy for a VPN traffic rule.
// returns a *VpnTrafficRuleRoutingPolicyType when successful
func (m *VpnTrafficRule) GetRoutingPolicyType()(*VpnTrafficRuleRoutingPolicyType) {
    val, err := m.GetBackingStore().Get("routingPolicyType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnTrafficRuleRoutingPolicyType)
    }
    return nil
}
// GetVpnTrafficDirection gets the vpnTrafficDirection property value. Specify whether the rule applies to inbound traffic or outbound traffic.
// returns a *VpnTrafficDirection when successful
func (m *VpnTrafficRule) GetVpnTrafficDirection()(*VpnTrafficDirection) {
    val, err := m.GetBackingStore().Get("vpnTrafficDirection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VpnTrafficDirection)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VpnTrafficRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetAppType() != nil {
        cast := (*m.GetAppType()).String()
        err := writer.WriteStringValue("appType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("claims", m.GetClaims())
        if err != nil {
            return err
        }
    }
    if m.GetLocalAddressRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalAddressRanges()))
        for i, v := range m.GetLocalAddressRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("localAddressRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalPortRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalPortRanges()))
        for i, v := range m.GetLocalPortRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("localPortRanges", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteInt32Value("protocols", m.GetProtocols())
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAddressRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoteAddressRanges()))
        for i, v := range m.GetRemoteAddressRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("remoteAddressRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemotePortRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemotePortRanges()))
        for i, v := range m.GetRemotePortRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("remotePortRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoutingPolicyType() != nil {
        cast := (*m.GetRoutingPolicyType()).String()
        err := writer.WriteStringValue("routingPolicyType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVpnTrafficDirection() != nil {
        cast := (*m.GetVpnTrafficDirection()).String()
        err := writer.WriteStringValue("vpnTrafficDirection", &cast)
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
func (m *VpnTrafficRule) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAppId sets the appId property value. App identifier, if this traffic rule is triggered by an app.
func (m *VpnTrafficRule) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppType sets the appType property value. Indicates the type of app that a VPN traffic rule is associated with.
func (m *VpnTrafficRule) SetAppType(value *VpnTrafficRuleAppType)() {
    err := m.GetBackingStore().Set("appType", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *VpnTrafficRule) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetClaims sets the claims property value. Claims associated with this traffic rule.
func (m *VpnTrafficRule) SetClaims(value *string)() {
    err := m.GetBackingStore().Set("claims", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalAddressRanges sets the localAddressRanges property value. Local address range. This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) SetLocalAddressRanges(value []IPv4Rangeable)() {
    err := m.GetBackingStore().Set("localAddressRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalPortRanges sets the localPortRanges property value. Local port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) SetLocalPortRanges(value []NumberRangeable)() {
    err := m.GetBackingStore().Set("localPortRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name.
func (m *VpnTrafficRule) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VpnTrafficRule) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProtocols sets the protocols property value. Protocols (0-255). Valid values 0 to 255
func (m *VpnTrafficRule) SetProtocols(value *int32)() {
    err := m.GetBackingStore().Set("protocols", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteAddressRanges sets the remoteAddressRanges property value. Remote address range. This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) SetRemoteAddressRanges(value []IPv4Rangeable)() {
    err := m.GetBackingStore().Set("remoteAddressRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetRemotePortRanges sets the remotePortRanges property value. Remote port range can be set only when protocol is either TCP or UDP (6 or 17). This collection can contain a maximum of 500 elements.
func (m *VpnTrafficRule) SetRemotePortRanges(value []NumberRangeable)() {
    err := m.GetBackingStore().Set("remotePortRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetRoutingPolicyType sets the routingPolicyType property value. Specifies the routing policy for a VPN traffic rule.
func (m *VpnTrafficRule) SetRoutingPolicyType(value *VpnTrafficRuleRoutingPolicyType)() {
    err := m.GetBackingStore().Set("routingPolicyType", value)
    if err != nil {
        panic(err)
    }
}
// SetVpnTrafficDirection sets the vpnTrafficDirection property value. Specify whether the rule applies to inbound traffic or outbound traffic.
func (m *VpnTrafficRule) SetVpnTrafficDirection(value *VpnTrafficDirection)() {
    err := m.GetBackingStore().Set("vpnTrafficDirection", value)
    if err != nil {
        panic(err)
    }
}
type VpnTrafficRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetAppType()(*VpnTrafficRuleAppType)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetClaims()(*string)
    GetLocalAddressRanges()([]IPv4Rangeable)
    GetLocalPortRanges()([]NumberRangeable)
    GetName()(*string)
    GetOdataType()(*string)
    GetProtocols()(*int32)
    GetRemoteAddressRanges()([]IPv4Rangeable)
    GetRemotePortRanges()([]NumberRangeable)
    GetRoutingPolicyType()(*VpnTrafficRuleRoutingPolicyType)
    GetVpnTrafficDirection()(*VpnTrafficDirection)
    SetAppId(value *string)()
    SetAppType(value *VpnTrafficRuleAppType)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetClaims(value *string)()
    SetLocalAddressRanges(value []IPv4Rangeable)()
    SetLocalPortRanges(value []NumberRangeable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetProtocols(value *int32)()
    SetRemoteAddressRanges(value []IPv4Rangeable)()
    SetRemotePortRanges(value []NumberRangeable)()
    SetRoutingPolicyType(value *VpnTrafficRuleRoutingPolicyType)()
    SetVpnTrafficDirection(value *VpnTrafficDirection)()
}
