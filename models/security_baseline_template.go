package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineTemplate 
type SecurityBaselineTemplate struct {
    DeviceManagementTemplate
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecurityBaselineTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityBaselineTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategoryDeviceStateSummaries gets the categoryDeviceStateSummaries property value. The security baseline per category device state summary
func (m *SecurityBaselineTemplate) GetCategoryDeviceStateSummaries()([]SecurityBaselineCategoryStateSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.categoryDeviceStateSummaries
    }
}
// GetDeviceStates gets the deviceStates property value. The security baseline device states
func (m *SecurityBaselineTemplate) GetDeviceStates()([]SecurityBaselineDeviceStateable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// GetDeviceStateSummary gets the deviceStateSummary property value. The security baseline device state summary
func (m *SecurityBaselineTemplate) GetDeviceStateSummary()(SecurityBaselineStateSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStateSummary
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityBaselineTemplate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategoryDeviceStateSummaries sets the categoryDeviceStateSummaries property value. The security baseline per category device state summary
func (m *SecurityBaselineTemplate) SetCategoryDeviceStateSummaries(value []SecurityBaselineCategoryStateSummaryable)() {
    if m != nil {
        m.categoryDeviceStateSummaries = value
    }
}
// SetDeviceStates sets the deviceStates property value. The security baseline device states
func (m *SecurityBaselineTemplate) SetDeviceStates(value []SecurityBaselineDeviceStateable)() {
    if m != nil {
        m.deviceStates = value
    }
}
// SetDeviceStateSummary sets the deviceStateSummary property value. The security baseline device state summary
func (m *SecurityBaselineTemplate) SetDeviceStateSummary(value SecurityBaselineStateSummaryable)() {
    if m != nil {
        m.deviceStateSummary = value
    }
}
