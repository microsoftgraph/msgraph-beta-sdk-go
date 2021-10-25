package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RoleManagement struct {
    additionalData map[string]interface{};
    cloudPC *RbacApplicationMultiple;
    deviceManagement *RbacApplicationMultiple;
    directory *RbacApplication;
    entitlementManagement *RbacApplication;
}
func NewRoleManagement()(*RoleManagement) {
    m := &RoleManagement{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RoleManagement) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RoleManagement) GetCloudPC()(*RbacApplicationMultiple) {
    if m == nil {
        return nil
    } else {
        return m.cloudPC
    }
}
func (m *RoleManagement) GetDeviceManagement()(*RbacApplicationMultiple) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagement
    }
}
func (m *RoleManagement) GetDirectory()(*RbacApplication) {
    if m == nil {
        return nil
    } else {
        return m.directory
    }
}
func (m *RoleManagement) GetEntitlementManagement()(*RbacApplication) {
    if m == nil {
        return nil
    } else {
        return m.entitlementManagement
    }
}
func (m *RoleManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cloudPC"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRbacApplicationMultiple() })
        if err != nil {
            return err
        }
        m.SetCloudPC(val.(*RbacApplicationMultiple))
        return nil
    }
    res["deviceManagement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRbacApplicationMultiple() })
        if err != nil {
            return err
        }
        m.SetDeviceManagement(val.(*RbacApplicationMultiple))
        return nil
    }
    res["directory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRbacApplication() })
        if err != nil {
            return err
        }
        m.SetDirectory(val.(*RbacApplication))
        return nil
    }
    res["entitlementManagement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRbacApplication() })
        if err != nil {
            return err
        }
        m.SetEntitlementManagement(val.(*RbacApplication))
        return nil
    }
    return res
}
func (m *RoleManagement) IsNil()(bool) {
    return m == nil
}
func (m *RoleManagement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cloudPC", m.GetCloudPC())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceManagement", m.GetDeviceManagement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("entitlementManagement", m.GetEntitlementManagement())
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
func (m *RoleManagement) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RoleManagement) SetCloudPC(value *RbacApplicationMultiple)() {
    m.cloudPC = value
}
func (m *RoleManagement) SetDeviceManagement(value *RbacApplicationMultiple)() {
    m.deviceManagement = value
}
func (m *RoleManagement) SetDirectory(value *RbacApplication)() {
    m.directory = value
}
func (m *RoleManagement) SetEntitlementManagement(value *RbacApplication)() {
    m.entitlementManagement = value
}
