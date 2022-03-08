package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcUserRoleScopeTagInfo provides operations to manage the deviceManagement singleton.
type CloudPcUserRoleScopeTagInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Scope tag display name.
    displayName *string;
    // Scope tag ID.
    roleScopeTagId *string;
}
// NewCloudPcUserRoleScopeTagInfo instantiates a new cloudPcUserRoleScopeTagInfo and sets the default values.
func NewCloudPcUserRoleScopeTagInfo()(*CloudPcUserRoleScopeTagInfo) {
    m := &CloudPcUserRoleScopeTagInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudPcUserRoleScopeTagInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcUserRoleScopeTagInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Scope tag display name.
func (m *CloudPcUserRoleScopeTagInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcUserRoleScopeTagInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
// GetRoleScopeTagId gets the roleScopeTagId property value. Scope tag ID.
func (m *CloudPcUserRoleScopeTagInfo) GetRoleScopeTagId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagId
    }
}
func (m *CloudPcUserRoleScopeTagInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcUserRoleScopeTagInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *CloudPcUserRoleScopeTagInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Scope tag display name.
func (m *CloudPcUserRoleScopeTagInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetRoleScopeTagId sets the roleScopeTagId property value. Scope tag ID.
func (m *CloudPcUserRoleScopeTagInfo) SetRoleScopeTagId(value *string)() {
    if m != nil {
        m.roleScopeTagId = value
    }
}
