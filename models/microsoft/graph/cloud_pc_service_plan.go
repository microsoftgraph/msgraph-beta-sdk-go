package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcServicePlan struct {
    Entity
    // The name for the service plan. Read-only.
    displayName *string;
    // The size of the RAM in GB. Read-only.
    ramInGB *int32;
    // The size of the OS Disk in GB. Read-only.
    storageInGB *int32;
    // The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
    type_escaped *CloudPcServicePlanType;
    // The size of the user profile disk in GB. Read-only.
    userProfileInGB *int32;
    // The number of vCPUs. Read-only.
    vCpuCount *int32;
}
// Instantiates a new cloudPcServicePlan and sets the default values.
func NewCloudPcServicePlan()(*CloudPcServicePlan) {
    m := &CloudPcServicePlan{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The name for the service plan. Read-only.
func (m *CloudPcServicePlan) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the ramInGB property value. The size of the RAM in GB. Read-only.
func (m *CloudPcServicePlan) GetRamInGB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ramInGB
    }
}
// Gets the storageInGB property value. The size of the OS Disk in GB. Read-only.
func (m *CloudPcServicePlan) GetStorageInGB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.storageInGB
    }
}
// Gets the type_escaped property value. The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
func (m *CloudPcServicePlan) GetType_escaped()(*CloudPcServicePlanType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the userProfileInGB property value. The size of the user profile disk in GB. Read-only.
func (m *CloudPcServicePlan) GetUserProfileInGB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.userProfileInGB
    }
}
// Gets the vCpuCount property value. The number of vCPUs. Read-only.
func (m *CloudPcServicePlan) GetVCpuCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.vCpuCount
    }
}
// The deserialization information for the current model
func (m *CloudPcServicePlan) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["ramInGB"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRamInGB(val)
        return nil
    }
    res["storageInGB"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetStorageInGB(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcServicePlanType)
        if err != nil {
            return err
        }
        cast := val.(CloudPcServicePlanType)
        m.SetType_escaped(&cast)
        return nil
    }
    res["userProfileInGB"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUserProfileInGB(val)
        return nil
    }
    res["vCpuCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetVCpuCount(val)
        return nil
    }
    return res
}
func (m *CloudPcServicePlan) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcServicePlan) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ramInGB", m.GetRamInGB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("storageInGB", m.GetStorageInGB())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("userProfileInGB", m.GetUserProfileInGB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("vCpuCount", m.GetVCpuCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The name for the service plan. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CloudPcServicePlan) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the ramInGB property value. The size of the RAM in GB. Read-only.
// Parameters:
//  - value : Value to set for the ramInGB property.
func (m *CloudPcServicePlan) SetRamInGB(value *int32)() {
    m.ramInGB = value
}
// Sets the storageInGB property value. The size of the OS Disk in GB. Read-only.
// Parameters:
//  - value : Value to set for the storageInGB property.
func (m *CloudPcServicePlan) SetStorageInGB(value *int32)() {
    m.storageInGB = value
}
// Sets the type_escaped property value. The type of the service plan. Possible values are: enterprise, business, unknownFutureValue. Read-only.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *CloudPcServicePlan) SetType_escaped(value *CloudPcServicePlanType)() {
    m.type_escaped = value
}
// Sets the userProfileInGB property value. The size of the user profile disk in GB. Read-only.
// Parameters:
//  - value : Value to set for the userProfileInGB property.
func (m *CloudPcServicePlan) SetUserProfileInGB(value *int32)() {
    m.userProfileInGB = value
}
// Sets the vCpuCount property value. The number of vCPUs. Read-only.
// Parameters:
//  - value : Value to set for the vCpuCount property.
func (m *CloudPcServicePlan) SetVCpuCount(value *int32)() {
    m.vCpuCount = value
}
