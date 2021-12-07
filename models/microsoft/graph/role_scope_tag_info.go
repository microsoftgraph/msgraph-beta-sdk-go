package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleScopeTagInfo 
type RoleScopeTagInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Scope Tag Display name.
    displayName *string;
    // Scope Tag Id.
    roleScopeTagId *string;
}
// NewRoleScopeTagInfo instantiates a new roleScopeTagInfo and sets the default values.
func NewRoleScopeTagInfo()(*RoleScopeTagInfo) {
    m := &RoleScopeTagInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleScopeTagInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Scope Tag Display name.
func (m *RoleScopeTagInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetRoleScopeTagId gets the roleScopeTagId property value. Scope Tag Id.
func (m *RoleScopeTagInfo) GetRoleScopeTagId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleScopeTagInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["roleScopeTagId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleScopeTagId(val)
        }
        return nil
    }
    return res
}
func (m *RoleScopeTagInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RoleScopeTagInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Scope Tag Display name.
func (m *RoleScopeTagInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetRoleScopeTagId sets the roleScopeTagId property value. Scope Tag Id.
func (m *RoleScopeTagInfo) SetRoleScopeTagId(value *string)() {
    if m != nil {
        m.roleScopeTagId = value
    }
}
