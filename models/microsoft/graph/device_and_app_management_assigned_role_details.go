package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceAndAppManagementAssignedRoleDetails struct {
    additionalData map[string]interface{};
    roleAssignmentIds []string;
    roleDefinitionIds []string;
}
func NewDeviceAndAppManagementAssignedRoleDetails()(*DeviceAndAppManagementAssignedRoleDetails) {
    m := &DeviceAndAppManagementAssignedRoleDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceAndAppManagementAssignedRoleDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceAndAppManagementAssignedRoleDetails) GetRoleAssignmentIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentIds
    }
}
func (m *DeviceAndAppManagementAssignedRoleDetails) GetRoleDefinitionIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionIds
    }
}
func (m *DeviceAndAppManagementAssignedRoleDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["roleAssignmentIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRoleAssignmentIds(res)
        return nil
    }
    res["roleDefinitionIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRoleDefinitionIds(res)
        return nil
    }
    return res
}
func (m *DeviceAndAppManagementAssignedRoleDetails) IsNil()(bool) {
    return m == nil
}
func (m *DeviceAndAppManagementAssignedRoleDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("roleAssignmentIds", m.GetRoleAssignmentIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("roleDefinitionIds", m.GetRoleDefinitionIds())
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
func (m *DeviceAndAppManagementAssignedRoleDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceAndAppManagementAssignedRoleDetails) SetRoleAssignmentIds(value []string)() {
    m.roleAssignmentIds = value
}
func (m *DeviceAndAppManagementAssignedRoleDetails) SetRoleDefinitionIds(value []string)() {
    m.roleDefinitionIds = value
}
