package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ConfigManagerPolicySummary a ConfigManager policy summary.
type ConfigManagerPolicySummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewConfigManagerPolicySummary instantiates a new configManagerPolicySummary and sets the default values.
func NewConfigManagerPolicySummary()(*ConfigManagerPolicySummary) {
    m := &ConfigManagerPolicySummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConfigManagerPolicySummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigManagerPolicySummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigManagerPolicySummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigManagerPolicySummary) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ConfigManagerPolicySummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. The number of devices evaluated to be compliant by the policy.
func (m *ConfigManagerPolicySummary) GetCompliantDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("compliantDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEnforcedDeviceCount gets the enforcedDeviceCount property value. The number of devices that have have been remediated by the policy.
func (m *ConfigManagerPolicySummary) GetEnforcedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("enforcedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFailedDeviceCount gets the failedDeviceCount property value. The number of devices that failed to be evaluated by the policy.
func (m *ConfigManagerPolicySummary) GetFailedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("failedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigManagerPolicySummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantDeviceCount(val)
        }
        return nil
    }
    res["enforcedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcedDeviceCount(val)
        }
        return nil
    }
    res["failedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    res["nonCompliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantDeviceCount(val)
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
    res["pendingDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingDeviceCount(val)
        }
        return nil
    }
    res["targetedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. The number of devices evaluated to be noncompliant by the policy.
func (m *ConfigManagerPolicySummary) GetNonCompliantDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("nonCompliantDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConfigManagerPolicySummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPendingDeviceCount gets the pendingDeviceCount property value. The number of devices that have acknowledged the policy but are pending evaluation.
func (m *ConfigManagerPolicySummary) GetPendingDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("pendingDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTargetedDeviceCount gets the targetedDeviceCount property value. The number of devices targeted by the policy.
func (m *ConfigManagerPolicySummary) GetTargetedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("targetedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ConfigManagerPolicySummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("compliantDeviceCount", m.GetCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("enforcedDeviceCount", m.GetEnforcedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("nonCompliantDeviceCount", m.GetNonCompliantDeviceCount())
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
        err := writer.WriteInt32Value("pendingDeviceCount", m.GetPendingDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("targetedDeviceCount", m.GetTargetedDeviceCount())
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
func (m *ConfigManagerPolicySummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ConfigManagerPolicySummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. The number of devices evaluated to be compliant by the policy.
func (m *ConfigManagerPolicySummary) SetCompliantDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("compliantDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetEnforcedDeviceCount sets the enforcedDeviceCount property value. The number of devices that have have been remediated by the policy.
func (m *ConfigManagerPolicySummary) SetEnforcedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("enforcedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetFailedDeviceCount sets the failedDeviceCount property value. The number of devices that failed to be evaluated by the policy.
func (m *ConfigManagerPolicySummary) SetFailedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("failedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. The number of devices evaluated to be noncompliant by the policy.
func (m *ConfigManagerPolicySummary) SetNonCompliantDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("nonCompliantDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConfigManagerPolicySummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPendingDeviceCount sets the pendingDeviceCount property value. The number of devices that have acknowledged the policy but are pending evaluation.
func (m *ConfigManagerPolicySummary) SetPendingDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("pendingDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedDeviceCount sets the targetedDeviceCount property value. The number of devices targeted by the policy.
func (m *ConfigManagerPolicySummary) SetTargetedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("targetedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// ConfigManagerPolicySummaryable 
type ConfigManagerPolicySummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompliantDeviceCount()(*int32)
    GetEnforcedDeviceCount()(*int32)
    GetFailedDeviceCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetOdataType()(*string)
    GetPendingDeviceCount()(*int32)
    GetTargetedDeviceCount()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompliantDeviceCount(value *int32)()
    SetEnforcedDeviceCount(value *int32)()
    SetFailedDeviceCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetOdataType(value *string)()
    SetPendingDeviceCount(value *int32)()
    SetTargetedDeviceCount(value *int32)()
}
