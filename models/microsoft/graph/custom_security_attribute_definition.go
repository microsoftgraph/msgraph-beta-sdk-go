package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomSecurityAttributeDefinition 
type CustomSecurityAttributeDefinition struct {
    Entity
    // 
    allowedValues []AllowedValue;
    // 
    attributeSet *string;
    // 
    description *string;
    // 
    isCollection *bool;
    // 
    isSearchable *bool;
    // 
    name *string;
    // 
    status *string;
    // 
    type_escaped *string;
    // 
    usePreDefinedValuesOnly *bool;
}
// NewCustomSecurityAttributeDefinition instantiates a new customSecurityAttributeDefinition and sets the default values.
func NewCustomSecurityAttributeDefinition()(*CustomSecurityAttributeDefinition) {
    m := &CustomSecurityAttributeDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllowedValues gets the allowedValues property value. 
func (m *CustomSecurityAttributeDefinition) GetAllowedValues()([]AllowedValue) {
    if m == nil {
        return nil
    } else {
        return m.allowedValues
    }
}
// GetAttributeSet gets the attributeSet property value. 
func (m *CustomSecurityAttributeDefinition) GetAttributeSet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attributeSet
    }
}
// GetDescription gets the description property value. 
func (m *CustomSecurityAttributeDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetIsCollection gets the isCollection property value. 
func (m *CustomSecurityAttributeDefinition) GetIsCollection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCollection
    }
}
// GetIsSearchable gets the isSearchable property value. 
func (m *CustomSecurityAttributeDefinition) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// GetName gets the name property value. 
func (m *CustomSecurityAttributeDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetStatus gets the status property value. 
func (m *CustomSecurityAttributeDefinition) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetType_escaped gets the type_escaped property value. 
func (m *CustomSecurityAttributeDefinition) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUsePreDefinedValuesOnly gets the usePreDefinedValuesOnly property value. 
func (m *CustomSecurityAttributeDefinition) GetUsePreDefinedValuesOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usePreDefinedValuesOnly
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomSecurityAttributeDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAllowedValue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AllowedValue, len(val))
            for i, v := range val {
                res[i] = *(v.(*AllowedValue))
            }
            m.SetAllowedValues(res)
        }
        return nil
    }
    res["attributeSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributeSet(val)
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
    res["isCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCollection(val)
        }
        return nil
    }
    res["isSearchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
        }
        return nil
    }
    res["usePreDefinedValuesOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsePreDefinedValuesOnly(val)
        }
        return nil
    }
    return res
}
func (m *CustomSecurityAttributeDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CustomSecurityAttributeDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedValues()))
        for i, v := range m.GetAllowedValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("allowedValues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("attributeSet", m.GetAttributeSet())
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
        err = writer.WriteBoolValue("isCollection", m.GetIsCollection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
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
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usePreDefinedValuesOnly", m.GetUsePreDefinedValuesOnly())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedValues sets the allowedValues property value. 
func (m *CustomSecurityAttributeDefinition) SetAllowedValues(value []AllowedValue)() {
    m.allowedValues = value
}
// SetAttributeSet sets the attributeSet property value. 
func (m *CustomSecurityAttributeDefinition) SetAttributeSet(value *string)() {
    m.attributeSet = value
}
// SetDescription sets the description property value. 
func (m *CustomSecurityAttributeDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetIsCollection sets the isCollection property value. 
func (m *CustomSecurityAttributeDefinition) SetIsCollection(value *bool)() {
    m.isCollection = value
}
// SetIsSearchable sets the isSearchable property value. 
func (m *CustomSecurityAttributeDefinition) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// SetName sets the name property value. 
func (m *CustomSecurityAttributeDefinition) SetName(value *string)() {
    m.name = value
}
// SetStatus sets the status property value. 
func (m *CustomSecurityAttributeDefinition) SetStatus(value *string)() {
    m.status = value
}
// SetType_escaped sets the type_escaped property value. 
func (m *CustomSecurityAttributeDefinition) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// SetUsePreDefinedValuesOnly sets the usePreDefinedValuesOnly property value. 
func (m *CustomSecurityAttributeDefinition) SetUsePreDefinedValuesOnly(value *bool)() {
    m.usePreDefinedValuesOnly = value
}
