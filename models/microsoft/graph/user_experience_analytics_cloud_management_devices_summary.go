package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsCloudManagementDevicesSummary struct {
    additionalData map[string]interface{};
    coManagedDeviceCount *int32;
    intuneDeviceCount *int32;
    tenantAttachDeviceCount *int32;
}
func NewUserExperienceAnalyticsCloudManagementDevicesSummary()(*UserExperienceAnalyticsCloudManagementDevicesSummary) {
    m := &UserExperienceAnalyticsCloudManagementDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetCoManagedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coManagedDeviceCount
    }
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetIntuneDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.intuneDeviceCount
    }
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetTenantAttachDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.tenantAttachDeviceCount
    }
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["coManagedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCoManagedDeviceCount(val)
        return nil
    }
    res["intuneDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIntuneDeviceCount(val)
        return nil
    }
    res["tenantAttachDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTenantAttachDeviceCount(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) IsNil()(bool) {
    return m == nil
}
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
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetCoManagedDeviceCount(value *int32)() {
    m.coManagedDeviceCount = value
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetIntuneDeviceCount(value *int32)() {
    m.intuneDeviceCount = value
}
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetTenantAttachDeviceCount(value *int32)() {
    m.tenantAttachDeviceCount = value
}
