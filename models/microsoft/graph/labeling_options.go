package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LabelingOptions struct {
    additionalData map[string]interface{};
    assignmentMethod *AssignmentMethod;
    downgradeJustification *DowngradeJustification;
    extendedProperties []KeyValuePair;
    labelId *string;
}
func NewLabelingOptions()(*LabelingOptions) {
    m := &LabelingOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LabelingOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LabelingOptions) GetAssignmentMethod()(*AssignmentMethod) {
    if m == nil {
        return nil
    } else {
        return m.assignmentMethod
    }
}
func (m *LabelingOptions) GetDowngradeJustification()(*DowngradeJustification) {
    if m == nil {
        return nil
    } else {
        return m.downgradeJustification
    }
}
func (m *LabelingOptions) GetExtendedProperties()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.extendedProperties
    }
}
func (m *LabelingOptions) GetLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.labelId
    }
}
func (m *LabelingOptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentMethod)
        if err != nil {
            return err
        }
        cast := val.(AssignmentMethod)
        m.SetAssignmentMethod(&cast)
        return nil
    }
    res["downgradeJustification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDowngradeJustification() })
        if err != nil {
            return err
        }
        m.SetDowngradeJustification(val.(*DowngradeJustification))
        return nil
    }
    res["extendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetExtendedProperties(res)
        return nil
    }
    res["labelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLabelId(val)
        return nil
    }
    return res
}
func (m *LabelingOptions) IsNil()(bool) {
    return m == nil
}
func (m *LabelingOptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := m.GetAssignmentMethod().String()
        err := writer.WriteStringValue("assignmentMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("downgradeJustification", m.GetDowngradeJustification())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtendedProperties()))
        for i, v := range m.GetExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("extendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labelId", m.GetLabelId())
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
func (m *LabelingOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LabelingOptions) SetAssignmentMethod(value *AssignmentMethod)() {
    m.assignmentMethod = value
}
func (m *LabelingOptions) SetDowngradeJustification(value *DowngradeJustification)() {
    m.downgradeJustification = value
}
func (m *LabelingOptions) SetExtendedProperties(value []KeyValuePair)() {
    m.extendedProperties = value
}
func (m *LabelingOptions) SetLabelId(value *string)() {
    m.labelId = value
}
