package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigManagerPolicySummary a ConfigManager policy summary.
type ConfigManagerPolicySummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of devices evaluated to be compliant by the policy.
    compliantDeviceCount *int32
    // The number of devices that have have been remediated by the policy.
    enforcedDeviceCount *int32
    // The number of devices that failed to be evaluated by the policy.
    failedDeviceCount *int32
    // The number of devices evaluated to be noncompliant by the policy.
    nonCompliantDeviceCount *int32
    // The OdataType property
    odataType *string
    // The number of devices that have acknowledged the policy but are pending evaluation.
    pendingDeviceCount *int32
    // The number of devices targeted by the policy.
    targetedDeviceCount *int32
}
// NewConfigManagerPolicySummary instantiates a new configManagerPolicySummary and sets the default values.
func NewConfigManagerPolicySummary()(*ConfigManagerPolicySummary) {
    m := &ConfigManagerPolicySummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.configManagerPolicySummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConfigManagerPolicySummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigManagerPolicySummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigManagerPolicySummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigManagerPolicySummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. The number of devices evaluated to be compliant by the policy.
func (m *ConfigManagerPolicySummary) GetCompliantDeviceCount()(*int32) {
    return m.compliantDeviceCount
}
// GetEnforcedDeviceCount gets the enforcedDeviceCount property value. The number of devices that have have been remediated by the policy.
func (m *ConfigManagerPolicySummary) GetEnforcedDeviceCount()(*int32) {
    return m.enforcedDeviceCount
}
// GetFailedDeviceCount gets the failedDeviceCount property value. The number of devices that failed to be evaluated by the policy.
func (m *ConfigManagerPolicySummary) GetFailedDeviceCount()(*int32) {
    return m.failedDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigManagerPolicySummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compliantDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCompliantDeviceCount)
    res["enforcedDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetEnforcedDeviceCount)
    res["failedDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedDeviceCount)
    res["nonCompliantDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNonCompliantDeviceCount)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["pendingDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPendingDeviceCount)
    res["targetedDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTargetedDeviceCount)
    return res
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. The number of devices evaluated to be noncompliant by the policy.
func (m *ConfigManagerPolicySummary) GetNonCompliantDeviceCount()(*int32) {
    return m.nonCompliantDeviceCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConfigManagerPolicySummary) GetOdataType()(*string) {
    return m.odataType
}
// GetPendingDeviceCount gets the pendingDeviceCount property value. The number of devices that have acknowledged the policy but are pending evaluation.
func (m *ConfigManagerPolicySummary) GetPendingDeviceCount()(*int32) {
    return m.pendingDeviceCount
}
// GetTargetedDeviceCount gets the targetedDeviceCount property value. The number of devices targeted by the policy.
func (m *ConfigManagerPolicySummary) GetTargetedDeviceCount()(*int32) {
    return m.targetedDeviceCount
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
func (m *ConfigManagerPolicySummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. The number of devices evaluated to be compliant by the policy.
func (m *ConfigManagerPolicySummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// SetEnforcedDeviceCount sets the enforcedDeviceCount property value. The number of devices that have have been remediated by the policy.
func (m *ConfigManagerPolicySummary) SetEnforcedDeviceCount(value *int32)() {
    m.enforcedDeviceCount = value
}
// SetFailedDeviceCount sets the failedDeviceCount property value. The number of devices that failed to be evaluated by the policy.
func (m *ConfigManagerPolicySummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. The number of devices evaluated to be noncompliant by the policy.
func (m *ConfigManagerPolicySummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConfigManagerPolicySummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPendingDeviceCount sets the pendingDeviceCount property value. The number of devices that have acknowledged the policy but are pending evaluation.
func (m *ConfigManagerPolicySummary) SetPendingDeviceCount(value *int32)() {
    m.pendingDeviceCount = value
}
// SetTargetedDeviceCount sets the targetedDeviceCount property value. The number of devices targeted by the policy.
func (m *ConfigManagerPolicySummary) SetTargetedDeviceCount(value *int32)() {
    m.targetedDeviceCount = value
}
