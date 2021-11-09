package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new customSecurityAttributeDefinition and sets the default values.
func NewCustomSecurityAttributeDefinition()(*CustomSecurityAttributeDefinition) {
    m := &CustomSecurityAttributeDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allowedValues property value. 
func (m *CustomSecurityAttributeDefinition) GetAllowedValues()([]AllowedValue) {
    if m == nil {
        return nil
    } else {
        return m.allowedValues
    }
}
// Gets the attributeSet property value. 
func (m *CustomSecurityAttributeDefinition) GetAttributeSet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attributeSet
    }
}
// Gets the description property value. 
func (m *CustomSecurityAttributeDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the isCollection property value. 
func (m *CustomSecurityAttributeDefinition) GetIsCollection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCollection
    }
}
// Gets the isSearchable property value. 
func (m *CustomSecurityAttributeDefinition) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// Gets the name property value. 
func (m *CustomSecurityAttributeDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the status property value. 
func (m *CustomSecurityAttributeDefinition) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the type_escaped property value. 
func (m *CustomSecurityAttributeDefinition) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the usePreDefinedValuesOnly property value. 
func (m *CustomSecurityAttributeDefinition) GetUsePreDefinedValuesOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usePreDefinedValuesOnly
    }
}
// The deserialization information for the current model
func (m *CustomSecurityAttributeDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAllowedValue() })
        if err != nil {
            return err
        }
        res := make([]AllowedValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*AllowedValue))
        }
        m.SetAllowedValues(res)
        return nil
    }
    res["attributeSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAttributeSet(val)
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
    res["isCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCollection(val)
        return nil
    }
    res["isSearchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSearchable(val)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    res["usePreDefinedValuesOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsePreDefinedValuesOnly(val)
        return nil
    }
    return res
}
func (m *CustomSecurityAttributeDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the allowedValues property value. 
// Parameters:
//  - value : Value to set for the allowedValues property.
func (m *CustomSecurityAttributeDefinition) SetAllowedValues(value []AllowedValue)() {
    m.allowedValues = value
}
// Sets the attributeSet property value. 
// Parameters:
//  - value : Value to set for the attributeSet property.
func (m *CustomSecurityAttributeDefinition) SetAttributeSet(value *string)() {
    m.attributeSet = value
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *CustomSecurityAttributeDefinition) SetDescription(value *string)() {
    m.description = value
}
// Sets the isCollection property value. 
// Parameters:
//  - value : Value to set for the isCollection property.
func (m *CustomSecurityAttributeDefinition) SetIsCollection(value *bool)() {
    m.isCollection = value
}
// Sets the isSearchable property value. 
// Parameters:
//  - value : Value to set for the isSearchable property.
func (m *CustomSecurityAttributeDefinition) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *CustomSecurityAttributeDefinition) SetName(value *string)() {
    m.name = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *CustomSecurityAttributeDefinition) SetStatus(value *string)() {
    m.status = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *CustomSecurityAttributeDefinition) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the usePreDefinedValuesOnly property value. 
// Parameters:
//  - value : Value to set for the usePreDefinedValuesOnly property.
func (m *CustomSecurityAttributeDefinition) SetUsePreDefinedValuesOnly(value *bool)() {
    m.usePreDefinedValuesOnly = value
}
