package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MatchingLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    applicationMode *ApplicationMode;
    // 
    description *string;
    // 
    displayName *string;
    // 
    id *string;
    // 
    isEndpointProtectionEnabled *bool;
    // 
    labelActions []LabelActionBase;
    // 
    name *string;
    // 
    policyTip *string;
    // 
    priority *int32;
    // 
    toolTip *string;
}
// Instantiates a new matchingLabel and sets the default values.
func NewMatchingLabel()(*MatchingLabel) {
    m := &MatchingLabel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MatchingLabel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the applicationMode property value. 
func (m *MatchingLabel) GetApplicationMode()(*ApplicationMode) {
    if m == nil {
        return nil
    } else {
        return m.applicationMode
    }
}
// Gets the description property value. 
func (m *MatchingLabel) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. 
func (m *MatchingLabel) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the id property value. 
func (m *MatchingLabel) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the isEndpointProtectionEnabled property value. 
func (m *MatchingLabel) GetIsEndpointProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEndpointProtectionEnabled
    }
}
// Gets the labelActions property value. 
func (m *MatchingLabel) GetLabelActions()([]LabelActionBase) {
    if m == nil {
        return nil
    } else {
        return m.labelActions
    }
}
// Gets the name property value. 
func (m *MatchingLabel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the policyTip property value. 
func (m *MatchingLabel) GetPolicyTip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyTip
    }
}
// Gets the priority property value. 
func (m *MatchingLabel) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// Gets the toolTip property value. 
func (m *MatchingLabel) GetToolTip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.toolTip
    }
}
// The deserialization information for the current model
func (m *MatchingLabel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationMode)
        if err != nil {
            return err
        }
        cast := val.(ApplicationMode)
        m.SetApplicationMode(&cast)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["isEndpointProtectionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEndpointProtectionEnabled(val)
        return nil
    }
    res["labelActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLabelActionBase() })
        if err != nil {
            return err
        }
        res := make([]LabelActionBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*LabelActionBase))
        }
        m.SetLabelActions(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["policyTip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyTip(val)
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPriority(val)
        return nil
    }
    res["toolTip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetToolTip(val)
        return nil
    }
    return res
}
func (m *MatchingLabel) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MatchingLabel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetApplicationMode() != nil {
        cast := m.GetApplicationMode().String()
        err := writer.WriteStringValue("applicationMode", &cast)
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
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEndpointProtectionEnabled", m.GetIsEndpointProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLabelActions()))
        for i, v := range m.GetLabelActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("labelActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyTip", m.GetPolicyTip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("toolTip", m.GetToolTip())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *MatchingLabel) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the applicationMode property value. 
// Parameters:
//  - value : Value to set for the applicationMode property.
func (m *MatchingLabel) SetApplicationMode(value *ApplicationMode)() {
    m.applicationMode = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *MatchingLabel) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MatchingLabel) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the id property value. 
// Parameters:
//  - value : Value to set for the id property.
func (m *MatchingLabel) SetId(value *string)() {
    m.id = value
}
// Sets the isEndpointProtectionEnabled property value. 
// Parameters:
//  - value : Value to set for the isEndpointProtectionEnabled property.
func (m *MatchingLabel) SetIsEndpointProtectionEnabled(value *bool)() {
    m.isEndpointProtectionEnabled = value
}
// Sets the labelActions property value. 
// Parameters:
//  - value : Value to set for the labelActions property.
func (m *MatchingLabel) SetLabelActions(value []LabelActionBase)() {
    m.labelActions = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *MatchingLabel) SetName(value *string)() {
    m.name = value
}
// Sets the policyTip property value. 
// Parameters:
//  - value : Value to set for the policyTip property.
func (m *MatchingLabel) SetPolicyTip(value *string)() {
    m.policyTip = value
}
// Sets the priority property value. 
// Parameters:
//  - value : Value to set for the priority property.
func (m *MatchingLabel) SetPriority(value *int32)() {
    m.priority = value
}
// Sets the toolTip property value. 
// Parameters:
//  - value : Value to set for the toolTip property.
func (m *MatchingLabel) SetToolTip(value *string)() {
    m.toolTip = value
}
