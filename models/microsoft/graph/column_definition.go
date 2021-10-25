package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ColumnDefinition struct {
    Entity
    boolean *BooleanColumn;
    calculated *CalculatedColumn;
    choice *ChoiceColumn;
    columnGroup *string;
    contentApprovalStatus *ContentApprovalStatusColumn;
    currency *CurrencyColumn;
    dateTime *DateTimeColumn;
    defaultValue *DefaultColumnValue;
    description *string;
    displayName *string;
    enforceUniqueValues *bool;
    geolocation *GeolocationColumn;
    hidden *bool;
    hyperlinkOrPicture *HyperlinkOrPictureColumn;
    indexed *bool;
    isDeletable *bool;
    isReorderable *bool;
    isSealed *bool;
    lookup *LookupColumn;
    name *string;
    number *NumberColumn;
    personOrGroup *PersonOrGroupColumn;
    propagateChanges *bool;
    readOnly *bool;
    required *bool;
    sourceColumn *ColumnDefinition;
    term *TermColumn;
    text *TextColumn;
    thumbnail *ThumbnailColumn;
    type_escaped *ColumnTypes;
    validation *ColumnValidation;
}
func NewColumnDefinition()(*ColumnDefinition) {
    m := &ColumnDefinition{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ColumnDefinition) GetBoolean()(*BooleanColumn) {
    if m == nil {
        return nil
    } else {
        return m.boolean
    }
}
func (m *ColumnDefinition) GetCalculated()(*CalculatedColumn) {
    if m == nil {
        return nil
    } else {
        return m.calculated
    }
}
func (m *ColumnDefinition) GetChoice()(*ChoiceColumn) {
    if m == nil {
        return nil
    } else {
        return m.choice
    }
}
func (m *ColumnDefinition) GetColumnGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.columnGroup
    }
}
func (m *ColumnDefinition) GetContentApprovalStatus()(*ContentApprovalStatusColumn) {
    if m == nil {
        return nil
    } else {
        return m.contentApprovalStatus
    }
}
func (m *ColumnDefinition) GetCurrency()(*CurrencyColumn) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
func (m *ColumnDefinition) GetDateTime()(*DateTimeColumn) {
    if m == nil {
        return nil
    } else {
        return m.dateTime
    }
}
func (m *ColumnDefinition) GetDefaultValue()(*DefaultColumnValue) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
func (m *ColumnDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *ColumnDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ColumnDefinition) GetEnforceUniqueValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enforceUniqueValues
    }
}
func (m *ColumnDefinition) GetGeolocation()(*GeolocationColumn) {
    if m == nil {
        return nil
    } else {
        return m.geolocation
    }
}
func (m *ColumnDefinition) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
func (m *ColumnDefinition) GetHyperlinkOrPicture()(*HyperlinkOrPictureColumn) {
    if m == nil {
        return nil
    } else {
        return m.hyperlinkOrPicture
    }
}
func (m *ColumnDefinition) GetIndexed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.indexed
    }
}
func (m *ColumnDefinition) GetIsDeletable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeletable
    }
}
func (m *ColumnDefinition) GetIsReorderable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReorderable
    }
}
func (m *ColumnDefinition) GetIsSealed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSealed
    }
}
func (m *ColumnDefinition) GetLookup()(*LookupColumn) {
    if m == nil {
        return nil
    } else {
        return m.lookup
    }
}
func (m *ColumnDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ColumnDefinition) GetNumber()(*NumberColumn) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
func (m *ColumnDefinition) GetPersonOrGroup()(*PersonOrGroupColumn) {
    if m == nil {
        return nil
    } else {
        return m.personOrGroup
    }
}
func (m *ColumnDefinition) GetPropagateChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateChanges
    }
}
func (m *ColumnDefinition) GetReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.readOnly
    }
}
func (m *ColumnDefinition) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
func (m *ColumnDefinition) GetSourceColumn()(*ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.sourceColumn
    }
}
func (m *ColumnDefinition) GetTerm()(*TermColumn) {
    if m == nil {
        return nil
    } else {
        return m.term
    }
}
func (m *ColumnDefinition) GetText()(*TextColumn) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
func (m *ColumnDefinition) GetThumbnail()(*ThumbnailColumn) {
    if m == nil {
        return nil
    } else {
        return m.thumbnail
    }
}
func (m *ColumnDefinition) GetType_escaped()(*ColumnTypes) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *ColumnDefinition) GetValidation()(*ColumnValidation) {
    if m == nil {
        return nil
    } else {
        return m.validation
    }
}
func (m *ColumnDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["boolean"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBooleanColumn() })
        if err != nil {
            return err
        }
        m.SetBoolean(val.(*BooleanColumn))
        return nil
    }
    res["calculated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalculatedColumn() })
        if err != nil {
            return err
        }
        m.SetCalculated(val.(*CalculatedColumn))
        return nil
    }
    res["choice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChoiceColumn() })
        if err != nil {
            return err
        }
        m.SetChoice(val.(*ChoiceColumn))
        return nil
    }
    res["columnGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColumnGroup(val)
        return nil
    }
    res["contentApprovalStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentApprovalStatusColumn() })
        if err != nil {
            return err
        }
        m.SetContentApprovalStatus(val.(*ContentApprovalStatusColumn))
        return nil
    }
    res["currency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrencyColumn() })
        if err != nil {
            return err
        }
        m.SetCurrency(val.(*CurrencyColumn))
        return nil
    }
    res["dateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeColumn() })
        if err != nil {
            return err
        }
        m.SetDateTime(val.(*DateTimeColumn))
        return nil
    }
    res["defaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDefaultColumnValue() })
        if err != nil {
            return err
        }
        m.SetDefaultValue(val.(*DefaultColumnValue))
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
    res["enforceUniqueValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnforceUniqueValues(val)
        return nil
    }
    res["geolocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeolocationColumn() })
        if err != nil {
            return err
        }
        m.SetGeolocation(val.(*GeolocationColumn))
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
    res["hyperlinkOrPicture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHyperlinkOrPictureColumn() })
        if err != nil {
            return err
        }
        m.SetHyperlinkOrPicture(val.(*HyperlinkOrPictureColumn))
        return nil
    }
    res["indexed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIndexed(val)
        return nil
    }
    res["isDeletable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeletable(val)
        return nil
    }
    res["isReorderable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsReorderable(val)
        return nil
    }
    res["isSealed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSealed(val)
        return nil
    }
    res["lookup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLookupColumn() })
        if err != nil {
            return err
        }
        m.SetLookup(val.(*LookupColumn))
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
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNumberColumn() })
        if err != nil {
            return err
        }
        m.SetNumber(val.(*NumberColumn))
        return nil
    }
    res["personOrGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonOrGroupColumn() })
        if err != nil {
            return err
        }
        m.SetPersonOrGroup(val.(*PersonOrGroupColumn))
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
    res["required"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequired(val)
        return nil
    }
    res["sourceColumn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        m.SetSourceColumn(val.(*ColumnDefinition))
        return nil
    }
    res["term"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermColumn() })
        if err != nil {
            return err
        }
        m.SetTerm(val.(*TermColumn))
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTextColumn() })
        if err != nil {
            return err
        }
        m.SetText(val.(*TextColumn))
        return nil
    }
    res["thumbnail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnailColumn() })
        if err != nil {
            return err
        }
        m.SetThumbnail(val.(*ThumbnailColumn))
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseColumnTypes)
        if err != nil {
            return err
        }
        cast := val.(ColumnTypes)
        m.SetType_escaped(&cast)
        return nil
    }
    res["validation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnValidation() })
        if err != nil {
            return err
        }
        m.SetValidation(val.(*ColumnValidation))
        return nil
    }
    return res
}
func (m *ColumnDefinition) IsNil()(bool) {
    return m == nil
}
func (m *ColumnDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("boolean", m.GetBoolean())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calculated", m.GetCalculated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("choice", m.GetChoice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("columnGroup", m.GetColumnGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentApprovalStatus", m.GetContentApprovalStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dateTime", m.GetDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultValue", m.GetDefaultValue())
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
        err = writer.WriteBoolValue("enforceUniqueValues", m.GetEnforceUniqueValues())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("geolocation", m.GetGeolocation())
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
        err = writer.WriteObjectValue("hyperlinkOrPicture", m.GetHyperlinkOrPicture())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("indexed", m.GetIndexed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeletable", m.GetIsDeletable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReorderable", m.GetIsReorderable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSealed", m.GetIsSealed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lookup", m.GetLookup())
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
        err = writer.WriteObjectValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("personOrGroup", m.GetPersonOrGroup())
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
        err = writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceColumn", m.GetSourceColumn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("term", m.GetTerm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("thumbnail", m.GetThumbnail())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("validation", m.GetValidation())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ColumnDefinition) SetBoolean(value *BooleanColumn)() {
    m.boolean = value
}
func (m *ColumnDefinition) SetCalculated(value *CalculatedColumn)() {
    m.calculated = value
}
func (m *ColumnDefinition) SetChoice(value *ChoiceColumn)() {
    m.choice = value
}
func (m *ColumnDefinition) SetColumnGroup(value *string)() {
    m.columnGroup = value
}
func (m *ColumnDefinition) SetContentApprovalStatus(value *ContentApprovalStatusColumn)() {
    m.contentApprovalStatus = value
}
func (m *ColumnDefinition) SetCurrency(value *CurrencyColumn)() {
    m.currency = value
}
func (m *ColumnDefinition) SetDateTime(value *DateTimeColumn)() {
    m.dateTime = value
}
func (m *ColumnDefinition) SetDefaultValue(value *DefaultColumnValue)() {
    m.defaultValue = value
}
func (m *ColumnDefinition) SetDescription(value *string)() {
    m.description = value
}
func (m *ColumnDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ColumnDefinition) SetEnforceUniqueValues(value *bool)() {
    m.enforceUniqueValues = value
}
func (m *ColumnDefinition) SetGeolocation(value *GeolocationColumn)() {
    m.geolocation = value
}
func (m *ColumnDefinition) SetHidden(value *bool)() {
    m.hidden = value
}
func (m *ColumnDefinition) SetHyperlinkOrPicture(value *HyperlinkOrPictureColumn)() {
    m.hyperlinkOrPicture = value
}
func (m *ColumnDefinition) SetIndexed(value *bool)() {
    m.indexed = value
}
func (m *ColumnDefinition) SetIsDeletable(value *bool)() {
    m.isDeletable = value
}
func (m *ColumnDefinition) SetIsReorderable(value *bool)() {
    m.isReorderable = value
}
func (m *ColumnDefinition) SetIsSealed(value *bool)() {
    m.isSealed = value
}
func (m *ColumnDefinition) SetLookup(value *LookupColumn)() {
    m.lookup = value
}
func (m *ColumnDefinition) SetName(value *string)() {
    m.name = value
}
func (m *ColumnDefinition) SetNumber(value *NumberColumn)() {
    m.number = value
}
func (m *ColumnDefinition) SetPersonOrGroup(value *PersonOrGroupColumn)() {
    m.personOrGroup = value
}
func (m *ColumnDefinition) SetPropagateChanges(value *bool)() {
    m.propagateChanges = value
}
func (m *ColumnDefinition) SetReadOnly(value *bool)() {
    m.readOnly = value
}
func (m *ColumnDefinition) SetRequired(value *bool)() {
    m.required = value
}
func (m *ColumnDefinition) SetSourceColumn(value *ColumnDefinition)() {
    m.sourceColumn = value
}
func (m *ColumnDefinition) SetTerm(value *TermColumn)() {
    m.term = value
}
func (m *ColumnDefinition) SetText(value *TextColumn)() {
    m.text = value
}
func (m *ColumnDefinition) SetThumbnail(value *ThumbnailColumn)() {
    m.thumbnail = value
}
func (m *ColumnDefinition) SetType_escaped(value *ColumnTypes)() {
    m.type_escaped = value
}
func (m *ColumnDefinition) SetValidation(value *ColumnValidation)() {
    m.validation = value
}
