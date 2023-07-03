package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SequentialActivationRenewalsAlertConfiguration 
type SequentialActivationRenewalsAlertConfiguration struct {
    UnifiedRoleManagementAlertConfiguration
}
// NewSequentialActivationRenewalsAlertConfiguration instantiates a new SequentialActivationRenewalsAlertConfiguration and sets the default values.
func NewSequentialActivationRenewalsAlertConfiguration()(*SequentialActivationRenewalsAlertConfiguration) {
    m := &SequentialActivationRenewalsAlertConfiguration{
        UnifiedRoleManagementAlertConfiguration: *NewUnifiedRoleManagementAlertConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.sequentialActivationRenewalsAlertConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSequentialActivationRenewalsAlertConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSequentialActivationRenewalsAlertConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSequentialActivationRenewalsAlertConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SequentialActivationRenewalsAlertConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementAlertConfiguration.GetFieldDeserializers()
    res["sequentialActivationCounterThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequentialActivationCounterThreshold(val)
        }
        return nil
    }
    res["timeIntervalBetweenActivations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeIntervalBetweenActivations(val)
        }
        return nil
    }
    return res
}
// GetSequentialActivationCounterThreshold gets the sequentialActivationCounterThreshold property value. The minimum number of activations within the timeIntervalBetweenActivations period to trigger an alert.
func (m *SequentialActivationRenewalsAlertConfiguration) GetSequentialActivationCounterThreshold()(*int32) {
    val, err := m.GetBackingStore().Get("sequentialActivationCounterThreshold")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTimeIntervalBetweenActivations gets the timeIntervalBetweenActivations property value. Time interval between activations to trigger an alert.
func (m *SequentialActivationRenewalsAlertConfiguration) GetTimeIntervalBetweenActivations()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("timeIntervalBetweenActivations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SequentialActivationRenewalsAlertConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementAlertConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("sequentialActivationCounterThreshold", m.GetSequentialActivationCounterThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("timeIntervalBetweenActivations", m.GetTimeIntervalBetweenActivations())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSequentialActivationCounterThreshold sets the sequentialActivationCounterThreshold property value. The minimum number of activations within the timeIntervalBetweenActivations period to trigger an alert.
func (m *SequentialActivationRenewalsAlertConfiguration) SetSequentialActivationCounterThreshold(value *int32)() {
    err := m.GetBackingStore().Set("sequentialActivationCounterThreshold", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeIntervalBetweenActivations sets the timeIntervalBetweenActivations property value. Time interval between activations to trigger an alert.
func (m *SequentialActivationRenewalsAlertConfiguration) SetTimeIntervalBetweenActivations(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("timeIntervalBetweenActivations", value)
    if err != nil {
        panic(err)
    }
}
// SequentialActivationRenewalsAlertConfigurationable 
type SequentialActivationRenewalsAlertConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementAlertConfigurationable
    GetSequentialActivationCounterThreshold()(*int32)
    GetTimeIntervalBetweenActivations()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetSequentialActivationCounterThreshold(value *int32)()
    SetTimeIntervalBetweenActivations(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}
