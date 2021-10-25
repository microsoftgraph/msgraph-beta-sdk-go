package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DocumentSet struct {
    additionalData map[string]interface{};
    allowedContentTypes []ContentTypeInfo;
    defaultContents []DocumentSetContent;
    propagateWelcomePageChanges *bool;
    sharedColumns []ColumnDefinition;
    shouldPrefixNameToFile *bool;
    welcomePageColumns []ColumnDefinition;
    welcomePageUrl *string;
}
func NewDocumentSet()(*DocumentSet) {
    m := &DocumentSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DocumentSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DocumentSet) GetAllowedContentTypes()([]ContentTypeInfo) {
    if m == nil {
        return nil
    } else {
        return m.allowedContentTypes
    }
}
func (m *DocumentSet) GetDefaultContents()([]DocumentSetContent) {
    if m == nil {
        return nil
    } else {
        return m.defaultContents
    }
}
func (m *DocumentSet) GetPropagateWelcomePageChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateWelcomePageChanges
    }
}
func (m *DocumentSet) GetSharedColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.sharedColumns
    }
}
func (m *DocumentSet) GetShouldPrefixNameToFile()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shouldPrefixNameToFile
    }
}
func (m *DocumentSet) GetWelcomePageColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.welcomePageColumns
    }
}
func (m *DocumentSet) GetWelcomePageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.welcomePageUrl
    }
}
func (m *DocumentSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedContentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeInfo() })
        if err != nil {
            return err
        }
        res := make([]ContentTypeInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*ContentTypeInfo))
        }
        m.SetAllowedContentTypes(res)
        return nil
    }
    res["defaultContents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSetContent() })
        if err != nil {
            return err
        }
        res := make([]DocumentSetContent, len(val))
        for i, v := range val {
            res[i] = *(v.(*DocumentSetContent))
        }
        m.SetDefaultContents(res)
        return nil
    }
    res["propagateWelcomePageChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPropagateWelcomePageChanges(val)
        return nil
    }
    res["sharedColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetSharedColumns(res)
        return nil
    }
    res["shouldPrefixNameToFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShouldPrefixNameToFile(val)
        return nil
    }
    res["welcomePageColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetWelcomePageColumns(res)
        return nil
    }
    res["welcomePageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWelcomePageUrl(val)
        return nil
    }
    return res
}
func (m *DocumentSet) IsNil()(bool) {
    return m == nil
}
func (m *DocumentSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedContentTypes()))
        for i, v := range m.GetAllowedContentTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("allowedContentTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefaultContents()))
        for i, v := range m.GetDefaultContents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("defaultContents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("propagateWelcomePageChanges", m.GetPropagateWelcomePageChanges())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSharedColumns()))
        for i, v := range m.GetSharedColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("sharedColumns", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("shouldPrefixNameToFile", m.GetShouldPrefixNameToFile())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWelcomePageColumns()))
        for i, v := range m.GetWelcomePageColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("welcomePageColumns", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("welcomePageUrl", m.GetWelcomePageUrl())
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
func (m *DocumentSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DocumentSet) SetAllowedContentTypes(value []ContentTypeInfo)() {
    m.allowedContentTypes = value
}
func (m *DocumentSet) SetDefaultContents(value []DocumentSetContent)() {
    m.defaultContents = value
}
func (m *DocumentSet) SetPropagateWelcomePageChanges(value *bool)() {
    m.propagateWelcomePageChanges = value
}
func (m *DocumentSet) SetSharedColumns(value []ColumnDefinition)() {
    m.sharedColumns = value
}
func (m *DocumentSet) SetShouldPrefixNameToFile(value *bool)() {
    m.shouldPrefixNameToFile = value
}
func (m *DocumentSet) SetWelcomePageColumns(value []ColumnDefinition)() {
    m.welcomePageColumns = value
}
func (m *DocumentSet) SetWelcomePageUrl(value *string)() {
    m.welcomePageUrl = value
}
