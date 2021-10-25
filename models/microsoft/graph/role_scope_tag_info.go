package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RoleScopeTagInfo struct {
    additionalData map[string]interface{};
    displayName *string;
    roleScopeTagId *string;
}
func NewRoleScopeTagInfo()(*RoleScopeTagInfo) {
    m := &RoleScopeTagInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RoleScopeTagInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RoleScopeTagInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *RoleScopeTagInfo) GetRoleScopeTagId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagId
    }
}
func (m *RoleScopeTagInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["roleScopeTagId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleScopeTagId(val)
        return nil
    }
    return res
}
func (m *RoleScopeTagInfo) IsNil()(bool) {
    return m == nil
}
func (m *RoleScopeTagInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roleScopeTagId", m.GetRoleScopeTagId())
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
func (m *RoleScopeTagInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RoleScopeTagInfo) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *RoleScopeTagInfo) SetRoleScopeTagId(value *string)() {
    m.roleScopeTagId = value
}
