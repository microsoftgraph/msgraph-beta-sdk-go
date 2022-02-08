package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LabelingOptions 
type LabelingOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Possible values are: standard, privileged, auto.
    assignmentMethod *AssignmentMethod;
    // The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
    downgradeJustification *DowngradeJustification;
    // Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
    extendedProperties []KeyValuePair;
    // The GUID of the label that should be applied to the information.
    labelId *string;
}
// NewLabelingOptions instantiates a new labelingOptions and sets the default values.
func NewLabelingOptions()(*LabelingOptions) {
    m := &LabelingOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LabelingOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignmentMethod gets the assignmentMethod property value. Possible values are: standard, privileged, auto.
func (m *LabelingOptions) GetAssignmentMethod()(*AssignmentMethod) {
    if m == nil {
        return nil
    } else {
        return m.assignmentMethod
    }
}
// GetDowngradeJustification gets the downgradeJustification property value. The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
func (m *LabelingOptions) GetDowngradeJustification()(*DowngradeJustification) {
    if m == nil {
        return nil
    } else {
        return m.downgradeJustification
    }
}
// GetExtendedProperties gets the extendedProperties property value. Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
func (m *LabelingOptions) GetExtendedProperties()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.extendedProperties
    }
}
// GetLabelId gets the labelId property value. The GUID of the label that should be applied to the information.
func (m *LabelingOptions) GetLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.labelId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelingOptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentMethod(val.(*AssignmentMethod))
        }
        return nil
    }
    res["downgradeJustification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDowngradeJustification() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeJustification(val.(*DowngradeJustification))
        }
        return nil
    }
    res["extendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetExtendedProperties(res)
        }
        return nil
    }
    res["labelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelId(val)
        }
        return nil
    }
    return res
}
func (m *LabelingOptions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LabelingOptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := (*m.GetAssignmentMethod()).String()
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
    if m.GetExtendedProperties() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LabelingOptions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignmentMethod sets the assignmentMethod property value. Possible values are: standard, privileged, auto.
func (m *LabelingOptions) SetAssignmentMethod(value *AssignmentMethod)() {
    if m != nil {
        m.assignmentMethod = value
    }
}
// SetDowngradeJustification sets the downgradeJustification property value. The downgrade justification object that indicates if downgrade was justified and, if so, the reason.
func (m *LabelingOptions) SetDowngradeJustification(value *DowngradeJustification)() {
    if m != nil {
        m.downgradeJustification = value
    }
}
// SetExtendedProperties sets the extendedProperties property value. Extended properties will be parsed and returned in the standard MIP labeled metadata format as part of the label information.
func (m *LabelingOptions) SetExtendedProperties(value []KeyValuePair)() {
    if m != nil {
        m.extendedProperties = value
    }
}
// SetLabelId sets the labelId property value. The GUID of the label that should be applied to the information.
func (m *LabelingOptions) SetLabelId(value *string)() {
    if m != nil {
        m.labelId = value
    }
}
