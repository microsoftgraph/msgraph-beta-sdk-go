package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type WorkloadAction struct {
    // The unique identifier for the workload action. Required. Read-only.
    actionId *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
    category *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadActionCategory;
    // The description for the workload action. Optional. Read-only.
    description *string;
    // The display name for the workload action. Optional. Read-only.
    displayName *string;
    // 
    licenses []string;
    // The service associated with workload action. Optional. Read-only.
    service *string;
    // The collection of settings associated with the workload action. Optional. Read-only.
    settings []Setting;
}
// Instantiates a new workloadAction and sets the default values.
func NewWorkloadAction()(*WorkloadAction) {
    m := &WorkloadAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadAction) GetActionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionId
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the category property value. The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
func (m *WorkloadAction) GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadActionCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the description property value. The description for the workload action. Optional. Read-only.
func (m *WorkloadAction) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for the workload action. Optional. Read-only.
func (m *WorkloadAction) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the licenses property value. 
func (m *WorkloadAction) GetLicenses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.licenses
    }
}
// Gets the service property value. The service associated with workload action. Optional. Read-only.
func (m *WorkloadAction) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// Gets the settings property value. The collection of settings associated with the workload action. Optional. Read-only.
func (m *WorkloadAction) GetSettings()([]Setting) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// The deserialization information for the current model
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
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseWorkloadActionCategory)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadActionCategory)
            m.SetCategory(&cast)
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Setting, len(val))
            for i, v := range val {
                res[i] = *(v.(*Setting))
            }
            m.SetSettings(res)
        }
        return nil
    }
    return res
}
func (m *WorkloadAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkloadAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionId", m.GetActionId())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := m.GetCategory().String()
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
    {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the actionId property value. The unique identifier for the workload action. Required. Read-only.
// Parameters:
//  - value : Value to set for the actionId property.
func (m *WorkloadAction) SetActionId(value *string)() {
    m.actionId = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WorkloadAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the category property value. The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
// Parameters:
//  - value : Value to set for the category property.
func (m *WorkloadAction) SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadActionCategory)() {
    m.category = value
}
// Sets the description property value. The description for the workload action. Optional. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *WorkloadAction) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for the workload action. Optional. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *WorkloadAction) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the licenses property value. 
// Parameters:
//  - value : Value to set for the licenses property.
func (m *WorkloadAction) SetLicenses(value []string)() {
    m.licenses = value
}
// Sets the service property value. The service associated with workload action. Optional. Read-only.
// Parameters:
//  - value : Value to set for the service property.
func (m *WorkloadAction) SetService(value *string)() {
    m.service = value
}
// Sets the settings property value. The collection of settings associated with the workload action. Optional. Read-only.
// Parameters:
//  - value : Value to set for the settings property.
func (m *WorkloadAction) SetSettings(value []Setting)() {
    m.settings = value
}
