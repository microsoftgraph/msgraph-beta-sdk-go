package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CloudPcDisasterRecoveryCapability struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudPcDisasterRecoveryCapability instantiates a new CloudPcDisasterRecoveryCapability and sets the default values.
func NewCloudPcDisasterRecoveryCapability()(*CloudPcDisasterRecoveryCapability) {
    m := &CloudPcDisasterRecoveryCapability{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcDisasterRecoveryCapabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcDisasterRecoveryCapabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDisasterRecoveryCapability(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloudPcDisasterRecoveryCapability) GetAdditionalData()(map[string]any) {
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
func (m *CloudPcDisasterRecoveryCapability) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCapabilityType gets the capabilityType property value. The disaster recovery action that can be performed for the Cloud PC. The possible values are: none, failover, failback, unknownFutureValue.
// returns a *CloudPcDisasterRecoveryCapabilityType when successful
func (m *CloudPcDisasterRecoveryCapability) GetCapabilityType()(*CloudPcDisasterRecoveryCapabilityType) {
    val, err := m.GetBackingStore().Get("capabilityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcDisasterRecoveryCapabilityType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcDisasterRecoveryCapability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capabilityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDisasterRecoveryCapabilityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilityType(val.(*CloudPcDisasterRecoveryCapabilityType))
        }
        return nil
    }
    res["licenseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDisasterRecoveryLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseType(val.(*CloudPcDisasterRecoveryLicenseType))
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
    res["primaryRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryRegion(val)
        }
        return nil
    }
    res["secondaryRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryRegion(val)
        }
        return nil
    }
    return res
}
// GetLicenseType gets the licenseType property value. The disaster recovery license type that provides the capability.
// returns a *CloudPcDisasterRecoveryLicenseType when successful
func (m *CloudPcDisasterRecoveryCapability) GetLicenseType()(*CloudPcDisasterRecoveryLicenseType) {
    val, err := m.GetBackingStore().Get("licenseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcDisasterRecoveryLicenseType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CloudPcDisasterRecoveryCapability) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrimaryRegion gets the primaryRegion property value. The primary and mainly used region where the Cloud PC is located.
// returns a *string when successful
func (m *CloudPcDisasterRecoveryCapability) GetPrimaryRegion()(*string) {
    val, err := m.GetBackingStore().Get("primaryRegion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecondaryRegion gets the secondaryRegion property value. The secondary region to which the Cloud PC can be failed over during a regional outage.
// returns a *string when successful
func (m *CloudPcDisasterRecoveryCapability) GetSecondaryRegion()(*string) {
    val, err := m.GetBackingStore().Get("secondaryRegion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcDisasterRecoveryCapability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCapabilityType() != nil {
        cast := (*m.GetCapabilityType()).String()
        err := writer.WriteStringValue("capabilityType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseType() != nil {
        cast := (*m.GetLicenseType()).String()
        err := writer.WriteStringValue("licenseType", &cast)
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
        err := writer.WriteStringValue("primaryRegion", m.GetPrimaryRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secondaryRegion", m.GetSecondaryRegion())
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
func (m *CloudPcDisasterRecoveryCapability) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CloudPcDisasterRecoveryCapability) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCapabilityType sets the capabilityType property value. The disaster recovery action that can be performed for the Cloud PC. The possible values are: none, failover, failback, unknownFutureValue.
func (m *CloudPcDisasterRecoveryCapability) SetCapabilityType(value *CloudPcDisasterRecoveryCapabilityType)() {
    err := m.GetBackingStore().Set("capabilityType", value)
    if err != nil {
        panic(err)
    }
}
// SetLicenseType sets the licenseType property value. The disaster recovery license type that provides the capability.
func (m *CloudPcDisasterRecoveryCapability) SetLicenseType(value *CloudPcDisasterRecoveryLicenseType)() {
    err := m.GetBackingStore().Set("licenseType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcDisasterRecoveryCapability) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryRegion sets the primaryRegion property value. The primary and mainly used region where the Cloud PC is located.
func (m *CloudPcDisasterRecoveryCapability) SetPrimaryRegion(value *string)() {
    err := m.GetBackingStore().Set("primaryRegion", value)
    if err != nil {
        panic(err)
    }
}
// SetSecondaryRegion sets the secondaryRegion property value. The secondary region to which the Cloud PC can be failed over during a regional outage.
func (m *CloudPcDisasterRecoveryCapability) SetSecondaryRegion(value *string)() {
    err := m.GetBackingStore().Set("secondaryRegion", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcDisasterRecoveryCapabilityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCapabilityType()(*CloudPcDisasterRecoveryCapabilityType)
    GetLicenseType()(*CloudPcDisasterRecoveryLicenseType)
    GetOdataType()(*string)
    GetPrimaryRegion()(*string)
    GetSecondaryRegion()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCapabilityType(value *CloudPcDisasterRecoveryCapabilityType)()
    SetLicenseType(value *CloudPcDisasterRecoveryLicenseType)()
    SetOdataType(value *string)()
    SetPrimaryRegion(value *string)()
    SetSecondaryRegion(value *string)()
}
