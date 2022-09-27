package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineTemplate 
type SecurityBaselineTemplate struct {
    DeviceManagementTemplate
    // The security baseline per category device state summary
    categoryDeviceStateSummaries []SecurityBaselineCategoryStateSummaryable
    // The security baseline device states
    deviceStates []SecurityBaselineDeviceStateable
    // The security baseline device state summary
    deviceStateSummary SecurityBaselineStateSummaryable
}
// NewSecurityBaselineTemplate instantiates a new SecurityBaselineTemplate and sets the default values.
func NewSecurityBaselineTemplate()(*SecurityBaselineTemplate) {
    m := &SecurityBaselineTemplate{
        DeviceManagementTemplate: *NewDeviceManagementTemplate(),
    }
    odataTypeValue := "#microsoft.graph.securityBaselineTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSecurityBaselineTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineTemplate(), nil
}
// GetCategoryDeviceStateSummaries gets the categoryDeviceStateSummaries property value. The security baseline per category device state summary
func (m *SecurityBaselineTemplate) GetCategoryDeviceStateSummaries()([]SecurityBaselineCategoryStateSummaryable) {
    return m.categoryDeviceStateSummaries
}
// GetDeviceStates gets the deviceStates property value. The security baseline device states
func (m *SecurityBaselineTemplate) GetDeviceStates()([]SecurityBaselineDeviceStateable) {
    return m.deviceStates
}
// GetDeviceStateSummary gets the deviceStateSummary property value. The security baseline device state summary
func (m *SecurityBaselineTemplate) GetDeviceStateSummary()(SecurityBaselineStateSummaryable) {
    return m.deviceStateSummary
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementTemplate.GetFieldDeserializers()
    res["categoryDeviceStateSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSecurityBaselineCategoryStateSummaryFromDiscriminatorValue , m.SetCategoryDeviceStateSummaries)
    res["deviceStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSecurityBaselineDeviceStateFromDiscriminatorValue , m.SetDeviceStates)
    res["deviceStateSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSecurityBaselineStateSummaryFromDiscriminatorValue , m.SetDeviceStateSummary)
    return res
}
// Serialize serializes information the current object
func (m *SecurityBaselineTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategoryDeviceStateSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCategoryDeviceStateSummaries())
        err = writer.WriteCollectionOfObjectValues("categoryDeviceStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceStates())
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStateSummary", m.GetDeviceStateSummary())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategoryDeviceStateSummaries sets the categoryDeviceStateSummaries property value. The security baseline per category device state summary
func (m *SecurityBaselineTemplate) SetCategoryDeviceStateSummaries(value []SecurityBaselineCategoryStateSummaryable)() {
    m.categoryDeviceStateSummaries = value
}
// SetDeviceStates sets the deviceStates property value. The security baseline device states
func (m *SecurityBaselineTemplate) SetDeviceStates(value []SecurityBaselineDeviceStateable)() {
    m.deviceStates = value
}
// SetDeviceStateSummary sets the deviceStateSummary property value. The security baseline device state summary
func (m *SecurityBaselineTemplate) SetDeviceStateSummary(value SecurityBaselineStateSummaryable)() {
    m.deviceStateSummary = value
}
