package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// ManagementTemplateDetailedInfo 
type ManagementTemplateDetailedInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
    category *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory;
    // The display name for the management template. Required. Read-only.
    displayName *string;
    // The unique identifier for the management template. Required. Read-only.
    managementTemplateId *string;
    // 
    version *int32;
}
// NewManagementTemplateDetailedInfo instantiates a new managementTemplateDetailedInfo and sets the default values.
func NewManagementTemplateDetailedInfo()(*ManagementTemplateDetailedInfo) {
    m := &ManagementTemplateDetailedInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementTemplateDetailedInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategory gets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetDisplayName gets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetManagementTemplateId gets the managementTemplateId property value. The unique identifier for the management template. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) GetManagementTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateId
    }
}
// GetVersion gets the version property value. 
func (m *ManagementTemplateDetailedInfo) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateDetailedInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory))
        }
        return nil
    }
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
    res["managementTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *ManagementTemplateDetailedInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementTemplateDetailedInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err := writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("version", m.GetVersion())
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
func (m *ManagementTemplateDetailedInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategory sets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetManagementTemplateId sets the managementTemplateId property value. The unique identifier for the management template. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) SetManagementTemplateId(value *string)() {
    if m != nil {
        m.managementTemplateId = value
    }
}
// SetVersion sets the version property value. 
func (m *ManagementTemplateDetailedInfo) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
