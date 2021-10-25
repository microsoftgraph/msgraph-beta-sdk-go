package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GroupPolicyDefinitionFile struct {
    Entity
    definitions []GroupPolicyDefinition;
    description *string;
    displayName *string;
    languageCodes []string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    policyType *GroupPolicyType;
    revision *string;
    targetNamespace *string;
    targetPrefix *string;
}
func NewGroupPolicyDefinitionFile()(*GroupPolicyDefinitionFile) {
    m := &GroupPolicyDefinitionFile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GroupPolicyDefinitionFile) GetDefinitions()([]GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definitions
    }
}
func (m *GroupPolicyDefinitionFile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *GroupPolicyDefinitionFile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GroupPolicyDefinitionFile) GetLanguageCodes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.languageCodes
    }
}
func (m *GroupPolicyDefinitionFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *GroupPolicyDefinitionFile) GetPolicyType()(*GroupPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
func (m *GroupPolicyDefinitionFile) GetRevision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.revision
    }
}
func (m *GroupPolicyDefinitionFile) GetTargetNamespace()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetNamespace
    }
}
func (m *GroupPolicyDefinitionFile) GetTargetPrefix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetPrefix
    }
}
func (m *GroupPolicyDefinitionFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinition() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyDefinition))
        }
        m.SetDefinitions(res)
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
    res["languageCodes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetLanguageCodes(res)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["policyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyType)
        if err != nil {
            return err
        }
        cast := val.(GroupPolicyType)
        m.SetPolicyType(&cast)
        return nil
    }
    res["revision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRevision(val)
        return nil
    }
    res["targetNamespace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetNamespace(val)
        return nil
    }
    res["targetPrefix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetPrefix(val)
        return nil
    }
    return res
}
func (m *GroupPolicyDefinitionFile) IsNil()(bool) {
    return m == nil
}
func (m *GroupPolicyDefinitionFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
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
        err = writer.WriteCollectionOfStringValues("languageCodes", m.GetLanguageCodes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyType() != nil {
        cast := m.GetPolicyType().String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("revision", m.GetRevision())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetNamespace", m.GetTargetNamespace())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetPrefix", m.GetTargetPrefix())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GroupPolicyDefinitionFile) SetDefinitions(value []GroupPolicyDefinition)() {
    m.definitions = value
}
func (m *GroupPolicyDefinitionFile) SetDescription(value *string)() {
    m.description = value
}
func (m *GroupPolicyDefinitionFile) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GroupPolicyDefinitionFile) SetLanguageCodes(value []string)() {
    m.languageCodes = value
}
func (m *GroupPolicyDefinitionFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *GroupPolicyDefinitionFile) SetPolicyType(value *GroupPolicyType)() {
    m.policyType = value
}
func (m *GroupPolicyDefinitionFile) SetRevision(value *string)() {
    m.revision = value
}
func (m *GroupPolicyDefinitionFile) SetTargetNamespace(value *string)() {
    m.targetNamespace = value
}
func (m *GroupPolicyDefinitionFile) SetTargetPrefix(value *string)() {
    m.targetPrefix = value
}
