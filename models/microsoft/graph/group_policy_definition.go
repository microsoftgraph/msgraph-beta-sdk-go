package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GroupPolicyDefinition struct {
    Entity
    category *GroupPolicyCategory;
    categoryPath *string;
    classType *GroupPolicyDefinitionClassType;
    definitionFile *GroupPolicyDefinitionFile;
    displayName *string;
    explainText *string;
    groupPolicyCategoryId *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    policyType *GroupPolicyType;
    presentations []GroupPolicyPresentation;
    supportedOn *string;
}
func NewGroupPolicyDefinition()(*GroupPolicyDefinition) {
    m := &GroupPolicyDefinition{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GroupPolicyDefinition) GetCategory()(*GroupPolicyCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *GroupPolicyDefinition) GetCategoryPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.categoryPath
    }
}
func (m *GroupPolicyDefinition) GetClassType()(*GroupPolicyDefinitionClassType) {
    if m == nil {
        return nil
    } else {
        return m.classType
    }
}
func (m *GroupPolicyDefinition) GetDefinitionFile()(*GroupPolicyDefinitionFile) {
    if m == nil {
        return nil
    } else {
        return m.definitionFile
    }
}
func (m *GroupPolicyDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GroupPolicyDefinition) GetExplainText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.explainText
    }
}
func (m *GroupPolicyDefinition) GetGroupPolicyCategoryId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyCategoryId
    }
}
func (m *GroupPolicyDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *GroupPolicyDefinition) GetPolicyType()(*GroupPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
func (m *GroupPolicyDefinition) GetPresentations()([]GroupPolicyPresentation) {
    if m == nil {
        return nil
    } else {
        return m.presentations
    }
}
func (m *GroupPolicyDefinition) GetSupportedOn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportedOn
    }
}
func (m *GroupPolicyDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyCategory() })
        if err != nil {
            return err
        }
        m.SetCategory(val.(*GroupPolicyCategory))
        return nil
    }
    res["categoryPath"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategoryPath(val)
        return nil
    }
    res["classType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyDefinitionClassType)
        if err != nil {
            return err
        }
        cast := val.(GroupPolicyDefinitionClassType)
        m.SetClassType(&cast)
        return nil
    }
    res["definitionFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinitionFile() })
        if err != nil {
            return err
        }
        m.SetDefinitionFile(val.(*GroupPolicyDefinitionFile))
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
    res["explainText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExplainText(val)
        return nil
    }
    res["groupPolicyCategoryId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupPolicyCategoryId(val)
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
    res["presentations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyPresentation() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyPresentation, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyPresentation))
        }
        m.SetPresentations(res)
        return nil
    }
    res["supportedOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSupportedOn(val)
        return nil
    }
    return res
}
func (m *GroupPolicyDefinition) IsNil()(bool) {
    return m == nil
}
func (m *GroupPolicyDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("categoryPath", m.GetCategoryPath())
        if err != nil {
            return err
        }
    }
    if m.GetClassType() != nil {
        cast := m.GetClassType().String()
        err = writer.WriteStringValue("classType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("definitionFile", m.GetDefinitionFile())
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
        err = writer.WriteStringValue("explainText", m.GetExplainText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupPolicyCategoryId", m.GetGroupPolicyCategoryId())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPresentations()))
        for i, v := range m.GetPresentations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("presentations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("supportedOn", m.GetSupportedOn())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GroupPolicyDefinition) SetCategory(value *GroupPolicyCategory)() {
    m.category = value
}
func (m *GroupPolicyDefinition) SetCategoryPath(value *string)() {
    m.categoryPath = value
}
func (m *GroupPolicyDefinition) SetClassType(value *GroupPolicyDefinitionClassType)() {
    m.classType = value
}
func (m *GroupPolicyDefinition) SetDefinitionFile(value *GroupPolicyDefinitionFile)() {
    m.definitionFile = value
}
func (m *GroupPolicyDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GroupPolicyDefinition) SetExplainText(value *string)() {
    m.explainText = value
}
func (m *GroupPolicyDefinition) SetGroupPolicyCategoryId(value *string)() {
    m.groupPolicyCategoryId = value
}
func (m *GroupPolicyDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *GroupPolicyDefinition) SetPolicyType(value *GroupPolicyType)() {
    m.policyType = value
}
func (m *GroupPolicyDefinition) SetPresentations(value []GroupPolicyPresentation)() {
    m.presentations = value
}
func (m *GroupPolicyDefinition) SetSupportedOn(value *string)() {
    m.supportedOn = value
}
