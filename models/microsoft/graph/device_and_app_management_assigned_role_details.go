package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceAndAppManagementAssignedRoleDetails 
type DeviceAndAppManagementAssignedRoleDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Role Assignment IDs for the specifc Role Assignments assigned to a user. This property is read-only.
    roleAssignmentIds []string;
    // Role Definition IDs for the specifc Role Definitions assigned to a user. This property is read-only.
    roleDefinitionIds []string;
}
// NewDeviceAndAppManagementAssignedRoleDetails instantiates a new deviceAndAppManagementAssignedRoleDetails and sets the default values.
func NewDeviceAndAppManagementAssignedRoleDetails()(*DeviceAndAppManagementAssignedRoleDetails) {
    m := &DeviceAndAppManagementAssignedRoleDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAndAppManagementAssignedRoleDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetRoleAssignmentIds gets the roleAssignmentIds property value. Role Assignment IDs for the specifc Role Assignments assigned to a user. This property is read-only.
func (m *DeviceAndAppManagementAssignedRoleDetails) GetRoleAssignmentIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentIds
    }
}
// GetRoleDefinitionIds gets the roleDefinitionIds property value. Role Definition IDs for the specifc Role Definitions assigned to a user. This property is read-only.
func (m *DeviceAndAppManagementAssignedRoleDetails) GetRoleDefinitionIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitionIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementAssignedRoleDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["roleAssignmentIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleAssignmentIds(res)
        }
        return nil
    }
    res["roleDefinitionIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleDefinitionIds(res)
        }
        return nil
    }
    return res
}
func (m *DeviceAndAppManagementAssignedRoleDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignedRoleDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetRoleAssignmentIds() != nil {
        err := writer.WriteCollectionOfStringValues("roleAssignmentIds", m.GetRoleAssignmentIds())
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitionIds() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAndAppManagementAssignedRoleDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRoleAssignmentIds sets the roleAssignmentIds property value. Role Assignment IDs for the specifc Role Assignments assigned to a user. This property is read-only.
func (m *DeviceAndAppManagementAssignedRoleDetails) SetRoleAssignmentIds(value []string)() {
    if m != nil {
        m.roleAssignmentIds = value
    }
}
// SetRoleDefinitionIds sets the roleDefinitionIds property value. Role Definition IDs for the specifc Role Definitions assigned to a user. This property is read-only.
func (m *DeviceAndAppManagementAssignedRoleDetails) SetRoleDefinitionIds(value []string)() {
    if m != nil {
        m.roleDefinitionIds = value
    }
}
