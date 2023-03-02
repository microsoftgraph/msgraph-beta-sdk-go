package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// VppTokenLicenseSummary license summary of a given app in a token.
type VppTokenLicenseSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewVppTokenLicenseSummary instantiates a new vppTokenLicenseSummary and sets the default values.
func NewVppTokenLicenseSummary()(*VppTokenLicenseSummary) {
    m := &VppTokenLicenseSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVppTokenLicenseSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVppTokenLicenseSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVppTokenLicenseSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VppTokenLicenseSummary) GetAdditionalData()(map[string]any) {
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
// GetAppleId gets the appleId property value. The Apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppTokenLicenseSummary) GetAppleId()(*string) {
    val, err := m.GetBackingStore().Get("appleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAvailableLicenseCount gets the availableLicenseCount property value. The number of VPP licenses available.
func (m *VppTokenLicenseSummary) GetAvailableLicenseCount()(*int32) {
    val, err := m.GetBackingStore().Get("availableLicenseCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *VppTokenLicenseSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VppTokenLicenseSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleId(val)
        }
        return nil
    }
    res["availableLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableLicenseCount(val)
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
    res["organizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationName(val)
        }
        return nil
    }
    res["usedLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedLicenseCount(val)
        }
        return nil
    }
    res["vppTokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VppTokenLicenseSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrganizationName gets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token.
func (m *VppTokenLicenseSummary) GetOrganizationName()(*string) {
    val, err := m.GetBackingStore().Get("organizationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *VppTokenLicenseSummary) GetUsedLicenseCount()(*int32) {
    val, err := m.GetBackingStore().Get("usedLicenseCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetVppTokenId gets the vppTokenId property value. Identifier of the VPP token.
func (m *VppTokenLicenseSummary) GetVppTokenId()(*string) {
    val, err := m.GetBackingStore().Get("vppTokenId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VppTokenLicenseSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("availableLicenseCount", m.GetAvailableLicenseCount())
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
        err := writer.WriteStringValue("organizationName", m.GetOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("usedLicenseCount", m.GetUsedLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vppTokenId", m.GetVppTokenId())
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
func (m *VppTokenLicenseSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAppleId sets the appleId property value. The Apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppTokenLicenseSummary) SetAppleId(value *string)() {
    err := m.GetBackingStore().Set("appleId", value)
    if err != nil {
        panic(err)
    }
}
// SetAvailableLicenseCount sets the availableLicenseCount property value. The number of VPP licenses available.
func (m *VppTokenLicenseSummary) SetAvailableLicenseCount(value *int32)() {
    err := m.GetBackingStore().Set("availableLicenseCount", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *VppTokenLicenseSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VppTokenLicenseSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationName sets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token.
func (m *VppTokenLicenseSummary) SetOrganizationName(value *string)() {
    err := m.GetBackingStore().Set("organizationName", value)
    if err != nil {
        panic(err)
    }
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *VppTokenLicenseSummary) SetUsedLicenseCount(value *int32)() {
    err := m.GetBackingStore().Set("usedLicenseCount", value)
    if err != nil {
        panic(err)
    }
}
// SetVppTokenId sets the vppTokenId property value. Identifier of the VPP token.
func (m *VppTokenLicenseSummary) SetVppTokenId(value *string)() {
    err := m.GetBackingStore().Set("vppTokenId", value)
    if err != nil {
        panic(err)
    }
}
// VppTokenLicenseSummaryable 
type VppTokenLicenseSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleId()(*string)
    GetAvailableLicenseCount()(*int32)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetOrganizationName()(*string)
    GetUsedLicenseCount()(*int32)
    GetVppTokenId()(*string)
    SetAppleId(value *string)()
    SetAvailableLicenseCount(value *int32)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetOrganizationName(value *string)()
    SetUsedLicenseCount(value *int32)()
    SetVppTokenId(value *string)()
}
