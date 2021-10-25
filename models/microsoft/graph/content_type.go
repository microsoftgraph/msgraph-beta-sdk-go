package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ContentType struct {
    Entity
    associatedHubsUrls []string;
    base *ContentType;
    baseTypes []ContentType;
    columnLinks []ColumnLink;
    columnPositions []ColumnDefinition;
    columns []ColumnDefinition;
    description *string;
    documentSet *DocumentSet;
    documentTemplate *DocumentSetContent;
    group *string;
    hidden *bool;
    inheritedFrom *ItemReference;
    isBuiltIn *bool;
    name *string;
    order *ContentTypeOrder;
    parentId *string;
    propagateChanges *bool;
    readOnly *bool;
    sealed *bool;
}
func NewContentType()(*ContentType) {
    m := &ContentType{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ContentType) GetAssociatedHubsUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.associatedHubsUrls
    }
}
func (m *ContentType) GetBase()(*ContentType) {
    if m == nil {
        return nil
    } else {
        return m.base
    }
}
func (m *ContentType) GetBaseTypes()([]ContentType) {
    if m == nil {
        return nil
    } else {
        return m.baseTypes
    }
}
func (m *ContentType) GetColumnLinks()([]ColumnLink) {
    if m == nil {
        return nil
    } else {
        return m.columnLinks
    }
}
func (m *ContentType) GetColumnPositions()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.columnPositions
    }
}
func (m *ContentType) GetColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
func (m *ContentType) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *ContentType) GetDocumentSet()(*DocumentSet) {
    if m == nil {
        return nil
    } else {
        return m.documentSet
    }
}
func (m *ContentType) GetDocumentTemplate()(*DocumentSetContent) {
    if m == nil {
        return nil
    } else {
        return m.documentTemplate
    }
}
func (m *ContentType) GetGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
func (m *ContentType) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
func (m *ContentType) GetInheritedFrom()(*ItemReference) {
    if m == nil {
        return nil
    } else {
        return m.inheritedFrom
    }
}
func (m *ContentType) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
func (m *ContentType) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ContentType) GetOrder()(*ContentTypeOrder) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
func (m *ContentType) GetParentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentId
    }
}
func (m *ContentType) GetPropagateChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateChanges
    }
}
func (m *ContentType) GetReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.readOnly
    }
}
func (m *ContentType) GetSealed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sealed
    }
}
func (m *ContentType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["associatedHubsUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssociatedHubsUrls(res)
        return nil
    }
    res["base"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentType() })
        if err != nil {
            return err
        }
        m.SetBase(val.(*ContentType))
        return nil
    }
    res["baseTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentType() })
        if err != nil {
            return err
        }
        res := make([]ContentType, len(val))
        for i, v := range val {
            res[i] = *(v.(*ContentType))
        }
        m.SetBaseTypes(res)
        return nil
    }
    res["columnLinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnLink() })
        if err != nil {
            return err
        }
        res := make([]ColumnLink, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnLink))
        }
        m.SetColumnLinks(res)
        return nil
    }
    res["columnPositions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetColumnPositions(res)
        return nil
    }
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetColumns(res)
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
    res["documentSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSet() })
        if err != nil {
            return err
        }
        m.SetDocumentSet(val.(*DocumentSet))
        return nil
    }
    res["documentTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSetContent() })
        if err != nil {
            return err
        }
        m.SetDocumentTemplate(val.(*DocumentSetContent))
        return nil
    }
    res["group"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroup(val)
        return nil
    }
    res["hidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHidden(val)
        return nil
    }
    res["inheritedFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemReference() })
        if err != nil {
            return err
        }
        m.SetInheritedFrom(val.(*ItemReference))
        return nil
    }
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsBuiltIn(val)
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
    res["order"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeOrder() })
        if err != nil {
            return err
        }
        m.SetOrder(val.(*ContentTypeOrder))
        return nil
    }
    res["parentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentId(val)
        return nil
    }
    res["propagateChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPropagateChanges(val)
        return nil
    }
    res["readOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetReadOnly(val)
        return nil
    }
    res["sealed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSealed(val)
        return nil
    }
    return res
}
func (m *ContentType) IsNil()(bool) {
    return m == nil
}
func (m *ContentType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("associatedHubsUrls", m.GetAssociatedHubsUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("base", m.GetBase())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBaseTypes()))
        for i, v := range m.GetBaseTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("baseTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumnLinks()))
        for i, v := range m.GetColumnLinks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columnLinks", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumnPositions()))
        for i, v := range m.GetColumnPositions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columnPositions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
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
        err = writer.WriteObjectValue("documentSet", m.GetDocumentSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("documentTemplate", m.GetDocumentTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hidden", m.GetHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inheritedFrom", m.GetInheritedFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
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
        err = writer.WriteObjectValue("order", m.GetOrder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentId", m.GetParentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("propagateChanges", m.GetPropagateChanges())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("readOnly", m.GetReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sealed", m.GetSealed())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ContentType) SetAssociatedHubsUrls(value []string)() {
    m.associatedHubsUrls = value
}
func (m *ContentType) SetBase(value *ContentType)() {
    m.base = value
}
func (m *ContentType) SetBaseTypes(value []ContentType)() {
    m.baseTypes = value
}
func (m *ContentType) SetColumnLinks(value []ColumnLink)() {
    m.columnLinks = value
}
func (m *ContentType) SetColumnPositions(value []ColumnDefinition)() {
    m.columnPositions = value
}
func (m *ContentType) SetColumns(value []ColumnDefinition)() {
    m.columns = value
}
func (m *ContentType) SetDescription(value *string)() {
    m.description = value
}
func (m *ContentType) SetDocumentSet(value *DocumentSet)() {
    m.documentSet = value
}
func (m *ContentType) SetDocumentTemplate(value *DocumentSetContent)() {
    m.documentTemplate = value
}
func (m *ContentType) SetGroup(value *string)() {
    m.group = value
}
func (m *ContentType) SetHidden(value *bool)() {
    m.hidden = value
}
func (m *ContentType) SetInheritedFrom(value *ItemReference)() {
    m.inheritedFrom = value
}
func (m *ContentType) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
func (m *ContentType) SetName(value *string)() {
    m.name = value
}
func (m *ContentType) SetOrder(value *ContentTypeOrder)() {
    m.order = value
}
func (m *ContentType) SetParentId(value *string)() {
    m.parentId = value
}
func (m *ContentType) SetPropagateChanges(value *bool)() {
    m.propagateChanges = value
}
func (m *ContentType) SetReadOnly(value *bool)() {
    m.readOnly = value
}
func (m *ContentType) SetSealed(value *bool)() {
    m.sealed = value
}
