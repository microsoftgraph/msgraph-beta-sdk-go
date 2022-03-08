package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementIntentInfo provides operations to manage the tenantRelationship singleton.
type ManagementIntentInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The display name for the management intent. Optional. Read-only.
    managementIntentDisplayName *string;
    // The identifier for the management intent. Required. Read-only.
    managementIntentId *string;
    // The collection of management template information associated with the management intent. Optional. Read-only.
    managementTemplates []ManagementTemplateDetailedInfoable;
}
// NewManagementIntentInfo instantiates a new managementIntentInfo and sets the default values.
func NewManagementIntentInfo()(*ManagementIntentInfo) {
    m := &ManagementIntentInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementIntentInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementIntentInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagementIntentInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementIntentInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementIntentInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementIntentDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementIntentDisplayName(val)
        }
        return nil
    }
    res["managementIntentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementIntentId(val)
        }
        return nil
    }
    res["managementTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateDetailedInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateDetailedInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateDetailedInfoable)
            }
            m.SetManagementTemplates(res)
        }
        return nil
    }
    return res
}
// GetManagementIntentDisplayName gets the managementIntentDisplayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) GetManagementIntentDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementIntentDisplayName
    }
}
// GetManagementIntentId gets the managementIntentId property value. The identifier for the management intent. Required. Read-only.
func (m *ManagementIntentInfo) GetManagementIntentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementIntentId
    }
}
// GetManagementTemplates gets the managementTemplates property value. The collection of management template information associated with the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) GetManagementTemplates()([]ManagementTemplateDetailedInfoable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
func (m *ManagementIntentInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementIntentInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementIntentDisplayName", m.GetManagementIntentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementIntentId", m.GetManagementIntentId())
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplates()))
        for i, v := range m.GetManagementTemplates() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("managementTemplates", cast)
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
func (m *ManagementIntentInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManagementIntentDisplayName sets the managementIntentDisplayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) SetManagementIntentDisplayName(value *string)() {
    if m != nil {
        m.managementIntentDisplayName = value
    }
}
// SetManagementIntentId sets the managementIntentId property value. The identifier for the management intent. Required. Read-only.
func (m *ManagementIntentInfo) SetManagementIntentId(value *string)() {
    if m != nil {
        m.managementIntentId = value
    }
}
// SetManagementTemplates sets the managementTemplates property value. The collection of management template information associated with the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) SetManagementTemplates(value []ManagementTemplateDetailedInfoable)() {
    if m != nil {
        m.managementTemplates = value
    }
}
