package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkloadAction provides operations to manage the tenantRelationship singleton.
type WorkloadAction struct {
    // The unique identifier for the workload action. Required. Read-only.
    actionId *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
    category *WorkloadActionCategory;
    // The description for the workload action. Optional. Read-only.
    description *string;
    // The display name for the workload action. Optional. Read-only.
    displayName *string;
    // 
    licenses []string;
    // The service associated with workload action. Optional. Read-only.
    service *string;
    // The collection of settings associated with the workload action. Optional. Read-only.
    settings []Settingable;
}
// NewWorkloadAction instantiates a new workloadAction and sets the default values.
func NewWorkloadAction()(*WorkloadAction) {
    m := &WorkloadAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkloadActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkloadActionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkloadAction(), nil
}
// GetActionId gets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadAction) GetActionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionId
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategory gets the category property value. The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
func (m *WorkloadAction) GetCategory()(*WorkloadActionCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetDescription gets the description property value. The description for the workload action. Optional. Read-only.
func (m *WorkloadAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the workload action. Optional. Read-only.
func (m *WorkloadAction) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkloadAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionId(val)
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkloadActionCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*WorkloadActionCategory))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["licenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLicenses(res)
        }
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Settingable, len(val))
            for i, v := range val {
                res[i] = v.(Settingable)
            }
            m.SetSettings(res)
        }
        return nil
    }
    return res
}
// GetLicenses gets the licenses property value. 
func (m *WorkloadAction) GetLicenses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.licenses
    }
}
// GetService gets the service property value. The service associated with workload action. Optional. Read-only.
func (m *WorkloadAction) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// GetSettings gets the settings property value. The collection of settings associated with the workload action. Optional. Read-only.
func (m *WorkloadAction) GetSettings()([]Settingable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *WorkloadAction) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkloadAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionId", m.GetActionId())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err := writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetLicenses() != nil {
        err := writer.WriteCollectionOfStringValues("licenses", m.GetLicenses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    if m.GetSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("settings", cast)
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
// SetActionId sets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadAction) SetActionId(value *string)() {
    if m != nil {
        m.actionId = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategory sets the category property value. The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
func (m *WorkloadAction) SetCategory(value *WorkloadActionCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetDescription sets the description property value. The description for the workload action. Optional. Read-only.
func (m *WorkloadAction) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the workload action. Optional. Read-only.
func (m *WorkloadAction) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLicenses sets the licenses property value. 
func (m *WorkloadAction) SetLicenses(value []string)() {
    if m != nil {
        m.licenses = value
    }
}
// SetService sets the service property value. The service associated with workload action. Optional. Read-only.
func (m *WorkloadAction) SetService(value *string)() {
    if m != nil {
        m.service = value
    }
}
// SetSettings sets the settings property value. The collection of settings associated with the workload action. Optional. Read-only.
func (m *WorkloadAction) SetSettings(value []Settingable)() {
    if m != nil {
        m.settings = value
    }
}
