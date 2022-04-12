package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterSupportedProperty represents the information about the property which is supported in crafting the rule of AssignmentFilter.
type AssignmentFilterSupportedProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The data type of the property.
    dataType *string
    // Indicates whether the property is a collection type or not.
    isCollection *bool
    // Name of the property.
    name *string
    // Regex string to do validation on the property value.
    propertyRegexConstraint *string
    // List of all supported operators on this property.
    supportedOperators []AssignmentFilterOperator
    // List of all supported values for this propery, empty if everything is supported.
    supportedValues []string
}
// NewAssignmentFilterSupportedProperty instantiates a new assignmentFilterSupportedProperty and sets the default values.
func NewAssignmentFilterSupportedProperty()(*AssignmentFilterSupportedProperty) {
    m := &AssignmentFilterSupportedProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignmentFilterSupportedPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterSupportedPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterSupportedProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterSupportedProperty) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDataType gets the dataType property value. The data type of the property.
func (m *AssignmentFilterSupportedProperty) GetDataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterSupportedProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dataType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val)
        }
        return nil
    }
    res["isCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCollection(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["propertyRegexConstraint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyRegexConstraint(val)
        }
        return nil
    }
    res["supportedOperators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAssignmentFilterOperator)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterOperator, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignmentFilterOperator))
            }
            m.SetSupportedOperators(res)
        }
        return nil
    }
    res["supportedValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedValues(res)
        }
        return nil
    }
    return res
}
// GetIsCollection gets the isCollection property value. Indicates whether the property is a collection type or not.
func (m *AssignmentFilterSupportedProperty) GetIsCollection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCollection
    }
}
// GetName gets the name property value. Name of the property.
func (m *AssignmentFilterSupportedProperty) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPropertyRegexConstraint gets the propertyRegexConstraint property value. Regex string to do validation on the property value.
func (m *AssignmentFilterSupportedProperty) GetPropertyRegexConstraint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.propertyRegexConstraint
    }
}
// GetSupportedOperators gets the supportedOperators property value. List of all supported operators on this property.
func (m *AssignmentFilterSupportedProperty) GetSupportedOperators()([]AssignmentFilterOperator) {
    if m == nil {
        return nil
    } else {
        return m.supportedOperators
    }
}
// GetSupportedValues gets the supportedValues property value. List of all supported values for this propery, empty if everything is supported.
func (m *AssignmentFilterSupportedProperty) GetSupportedValues()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedValues
    }
}
// Serialize serializes information the current object
func (m *AssignmentFilterSupportedProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dataType", m.GetDataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCollection", m.GetIsCollection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("propertyRegexConstraint", m.GetPropertyRegexConstraint())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedOperators() != nil {
        err := writer.WriteCollectionOfStringValues("supportedOperators", SerializeAssignmentFilterOperator(m.GetSupportedOperators()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedValues() != nil {
        err := writer.WriteCollectionOfStringValues("supportedValues", m.GetSupportedValues())
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
func (m *AssignmentFilterSupportedProperty) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDataType sets the dataType property value. The data type of the property.
func (m *AssignmentFilterSupportedProperty) SetDataType(value *string)() {
    if m != nil {
        m.dataType = value
    }
}
// SetIsCollection sets the isCollection property value. Indicates whether the property is a collection type or not.
func (m *AssignmentFilterSupportedProperty) SetIsCollection(value *bool)() {
    if m != nil {
        m.isCollection = value
    }
}
// SetName sets the name property value. Name of the property.
func (m *AssignmentFilterSupportedProperty) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPropertyRegexConstraint sets the propertyRegexConstraint property value. Regex string to do validation on the property value.
func (m *AssignmentFilterSupportedProperty) SetPropertyRegexConstraint(value *string)() {
    if m != nil {
        m.propertyRegexConstraint = value
    }
}
// SetSupportedOperators sets the supportedOperators property value. List of all supported operators on this property.
func (m *AssignmentFilterSupportedProperty) SetSupportedOperators(value []AssignmentFilterOperator)() {
    if m != nil {
        m.supportedOperators = value
    }
}
// SetSupportedValues sets the supportedValues property value. List of all supported values for this propery, empty if everything is supported.
func (m *AssignmentFilterSupportedProperty) SetSupportedValues(value []string)() {
    if m != nil {
        m.supportedValues = value
    }
}
