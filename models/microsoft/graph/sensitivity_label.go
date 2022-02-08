package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SensitivityLabel 
type SensitivityLabel struct {
    Entity
    // 
    applicableTo *SensitivityLabelTarget;
    // 
    applicationMode *ApplicationMode;
    // 
    assignedPolicies []LabelPolicy;
    // 
    autoLabeling *AutoLabeling;
    // 
    description *string;
    // 
    displayName *string;
    // 
    isDefault *bool;
    // 
    isEndpointProtectionEnabled *bool;
    // 
    labelActions []LabelActionBase;
    // 
    name *string;
    // 
    priority *int32;
    // 
    sublabels []SensitivityLabel;
    // 
    toolTip *string;
}
// NewSensitivityLabel instantiates a new sensitivityLabel and sets the default values.
func NewSensitivityLabel()(*SensitivityLabel) {
    m := &SensitivityLabel{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicableTo gets the applicableTo property value. 
func (m *SensitivityLabel) GetApplicableTo()(*SensitivityLabelTarget) {
    if m == nil {
        return nil
    } else {
        return m.applicableTo
    }
}
// GetApplicationMode gets the applicationMode property value. 
func (m *SensitivityLabel) GetApplicationMode()(*ApplicationMode) {
    if m == nil {
        return nil
    } else {
        return m.applicationMode
    }
}
// GetAssignedPolicies gets the assignedPolicies property value. 
func (m *SensitivityLabel) GetAssignedPolicies()([]LabelPolicy) {
    if m == nil {
        return nil
    } else {
        return m.assignedPolicies
    }
}
// GetAutoLabeling gets the autoLabeling property value. 
func (m *SensitivityLabel) GetAutoLabeling()(*AutoLabeling) {
    if m == nil {
        return nil
    } else {
        return m.autoLabeling
    }
}
// GetDescription gets the description property value. 
func (m *SensitivityLabel) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. 
func (m *SensitivityLabel) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsDefault gets the isDefault property value. 
func (m *SensitivityLabel) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetIsEndpointProtectionEnabled gets the isEndpointProtectionEnabled property value. 
func (m *SensitivityLabel) GetIsEndpointProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEndpointProtectionEnabled
    }
}
// GetLabelActions gets the labelActions property value. 
func (m *SensitivityLabel) GetLabelActions()([]LabelActionBase) {
    if m == nil {
        return nil
    } else {
        return m.labelActions
    }
}
// GetName gets the name property value. 
func (m *SensitivityLabel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPriority gets the priority property value. 
func (m *SensitivityLabel) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetSublabels gets the sublabels property value. 
func (m *SensitivityLabel) GetSublabels()([]SensitivityLabel) {
    if m == nil {
        return nil
    } else {
        return m.sublabels
    }
}
// GetToolTip gets the toolTip property value. 
func (m *SensitivityLabel) GetToolTip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.toolTip
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityLabel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivityLabelTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableTo(val.(*SensitivityLabelTarget))
        }
        return nil
    }
    res["applicationMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationMode(val.(*ApplicationMode))
        }
        return nil
    }
    res["assignedPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLabelPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LabelPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*LabelPolicy))
            }
            m.SetAssignedPolicies(res)
        }
        return nil
    }
    res["autoLabeling"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAutoLabeling() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoLabeling(val.(*AutoLabeling))
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
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isEndpointProtectionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEndpointProtectionEnabled(val)
        }
        return nil
    }
    res["labelActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLabelActionBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LabelActionBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*LabelActionBase))
            }
            m.SetLabelActions(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["sublabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitivityLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabel, len(val))
            for i, v := range val {
                res[i] = *(v.(*SensitivityLabel))
            }
            m.SetSublabels(res)
        }
        return nil
    }
    res["toolTip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToolTip(val)
        }
        return nil
    }
    return res
}
func (m *SensitivityLabel) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SensitivityLabel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableTo() != nil {
        cast := (*m.GetApplicableTo()).String()
        err = writer.WriteStringValue("applicableTo", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicationMode() != nil {
        cast := (*m.GetApplicationMode()).String()
        err = writer.WriteStringValue("applicationMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignedPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedPolicies()))
        for i, v := range m.GetAssignedPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("autoLabeling", m.GetAutoLabeling())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEndpointProtectionEnabled", m.GetIsEndpointProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetLabelActions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLabelActions()))
        for i, v := range m.GetLabelActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("labelActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetSublabels() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSublabels()))
        for i, v := range m.GetSublabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sublabels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("toolTip", m.GetToolTip())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableTo sets the applicableTo property value. 
func (m *SensitivityLabel) SetApplicableTo(value *SensitivityLabelTarget)() {
    if m != nil {
        m.applicableTo = value
    }
}
// SetApplicationMode sets the applicationMode property value. 
func (m *SensitivityLabel) SetApplicationMode(value *ApplicationMode)() {
    if m != nil {
        m.applicationMode = value
    }
}
// SetAssignedPolicies sets the assignedPolicies property value. 
func (m *SensitivityLabel) SetAssignedPolicies(value []LabelPolicy)() {
    if m != nil {
        m.assignedPolicies = value
    }
}
// SetAutoLabeling sets the autoLabeling property value. 
func (m *SensitivityLabel) SetAutoLabeling(value *AutoLabeling)() {
    if m != nil {
        m.autoLabeling = value
    }
}
// SetDescription sets the description property value. 
func (m *SensitivityLabel) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *SensitivityLabel) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsDefault sets the isDefault property value. 
func (m *SensitivityLabel) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetIsEndpointProtectionEnabled sets the isEndpointProtectionEnabled property value. 
func (m *SensitivityLabel) SetIsEndpointProtectionEnabled(value *bool)() {
    if m != nil {
        m.isEndpointProtectionEnabled = value
    }
}
// SetLabelActions sets the labelActions property value. 
func (m *SensitivityLabel) SetLabelActions(value []LabelActionBase)() {
    if m != nil {
        m.labelActions = value
    }
}
// SetName sets the name property value. 
func (m *SensitivityLabel) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPriority sets the priority property value. 
func (m *SensitivityLabel) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetSublabels sets the sublabels property value. 
func (m *SensitivityLabel) SetSublabels(value []SensitivityLabel)() {
    if m != nil {
        m.sublabels = value
    }
}
// SetToolTip sets the toolTip property value. 
func (m *SensitivityLabel) SetToolTip(value *string)() {
    if m != nil {
        m.toolTip = value
    }
}
