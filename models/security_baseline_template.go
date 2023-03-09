package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineTemplate 
type SecurityBaselineTemplate struct {
    DeviceManagementTemplate
}
// NewSecurityBaselineTemplate instantiates a new SecurityBaselineTemplate and sets the default values.
func NewSecurityBaselineTemplate()(*SecurityBaselineTemplate) {
    m := &SecurityBaselineTemplate{
        DeviceManagementTemplate: *NewDeviceManagementTemplate(),
    }
    odataTypeValue := "#microsoft.graph.securityBaselineTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSecurityBaselineTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineTemplate(), nil
}
// GetCategoryDeviceStateSummaries gets the categoryDeviceStateSummaries property value. The security baseline per category device state summary
func (m *SecurityBaselineTemplate) GetCategoryDeviceStateSummaries()([]SecurityBaselineCategoryStateSummaryable) {
    val, err := m.GetBackingStore().Get("categoryDeviceStateSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityBaselineCategoryStateSummaryable)
    }
    return nil
}
// GetDeviceStates gets the deviceStates property value. The security baseline device states
func (m *SecurityBaselineTemplate) GetDeviceStates()([]SecurityBaselineDeviceStateable) {
    val, err := m.GetBackingStore().Get("deviceStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityBaselineDeviceStateable)
    }
    return nil
}
// GetDeviceStateSummary gets the deviceStateSummary property value. The security baseline device state summary
func (m *SecurityBaselineTemplate) GetDeviceStateSummary()(SecurityBaselineStateSummaryable) {
    val, err := m.GetBackingStore().Get("deviceStateSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SecurityBaselineStateSummaryable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementTemplate.GetFieldDeserializers()
    res["categoryDeviceStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineCategoryStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineCategoryStateSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityBaselineCategoryStateSummaryable)
            }
            m.SetCategoryDeviceStateSummaries(res)
        }
        return nil
    }
    res["deviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineDeviceStateable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityBaselineDeviceStateable)
            }
            m.SetDeviceStates(res)
        }
        return nil
    }
    res["deviceStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityBaselineStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStateSummary(val.(SecurityBaselineStateSummaryable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SecurityBaselineTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategoryDeviceStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategoryDeviceStateSummaries()))
        for i, v := range m.GetCategoryDeviceStateSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("categoryDeviceStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("categoryDeviceStateSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceStates sets the deviceStates property value. The security baseline device states
func (m *SecurityBaselineTemplate) SetDeviceStates(value []SecurityBaselineDeviceStateable)() {
    err := m.GetBackingStore().Set("deviceStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceStateSummary sets the deviceStateSummary property value. The security baseline device state summary
func (m *SecurityBaselineTemplate) SetDeviceStateSummary(value SecurityBaselineStateSummaryable)() {
    err := m.GetBackingStore().Set("deviceStateSummary", value)
    if err != nil {
        panic(err)
    }
}
// SecurityBaselineTemplateable 
type SecurityBaselineTemplateable interface {
    DeviceManagementTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategoryDeviceStateSummaries()([]SecurityBaselineCategoryStateSummaryable)
    GetDeviceStates()([]SecurityBaselineDeviceStateable)
    GetDeviceStateSummary()(SecurityBaselineStateSummaryable)
    SetCategoryDeviceStateSummaries(value []SecurityBaselineCategoryStateSummaryable)()
    SetDeviceStates(value []SecurityBaselineDeviceStateable)()
    SetDeviceStateSummary(value SecurityBaselineStateSummaryable)()
}
