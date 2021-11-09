package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsCloudManagementDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total number of  co-managed devices.
    coManagedDeviceCount *int32;
    // The count of intune devices that are not autopilot registerd.
    intuneDeviceCount *int32;
    // Total count of tenant attach devices.
    tenantAttachDeviceCount *int32;
}
// Instantiates a new userExperienceAnalyticsCloudManagementDevicesSummary and sets the default values.
func NewUserExperienceAnalyticsCloudManagementDevicesSummary()(*UserExperienceAnalyticsCloudManagementDevicesSummary) {
    m := &UserExperienceAnalyticsCloudManagementDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the coManagedDeviceCount property value. Total number of  co-managed devices.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetCoManagedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coManagedDeviceCount
    }
}
// Gets the intuneDeviceCount property value. The count of intune devices that are not autopilot registerd.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetIntuneDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.intuneDeviceCount
    }
}
// Gets the tenantAttachDeviceCount property value. Total count of tenant attach devices.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetTenantAttachDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.tenantAttachDeviceCount
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["coManagedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoManagedDeviceCount(val)
        }
        return nil
    }
    res["intuneDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneDeviceCount(val)
        }
        return nil
    }
    res["tenantAttachDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAttachDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("coManagedDeviceCount", m.GetCoManagedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("intuneDeviceCount", m.GetIntuneDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("tenantAttachDeviceCount", m.GetTenantAttachDeviceCount())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the coManagedDeviceCount property value. Total number of  co-managed devices.
// Parameters:
//  - value : Value to set for the coManagedDeviceCount property.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetCoManagedDeviceCount(value *int32)() {
    m.coManagedDeviceCount = value
}
// Sets the intuneDeviceCount property value. The count of intune devices that are not autopilot registerd.
// Parameters:
//  - value : Value to set for the intuneDeviceCount property.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetIntuneDeviceCount(value *int32)() {
    m.intuneDeviceCount = value
}
// Sets the tenantAttachDeviceCount property value. Total count of tenant attach devices.
// Parameters:
//  - value : Value to set for the tenantAttachDeviceCount property.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetTenantAttachDeviceCount(value *int32)() {
    m.tenantAttachDeviceCount = value
}
