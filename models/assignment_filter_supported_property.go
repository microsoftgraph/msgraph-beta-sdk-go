package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AssignmentFilterSupportedProperty represents the information about the property which is supported in crafting the rule of AssignmentFilter.
type AssignmentFilterSupportedProperty struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAssignmentFilterSupportedProperty instantiates a new AssignmentFilterSupportedProperty and sets the default values.
func NewAssignmentFilterSupportedProperty()(*AssignmentFilterSupportedProperty) {
    m := &AssignmentFilterSupportedProperty{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAssignmentFilterSupportedPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAssignmentFilterSupportedPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterSupportedProperty(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AssignmentFilterSupportedProperty) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AssignmentFilterSupportedProperty) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDataType gets the dataType property value. The data type of the property.
// returns a *string when successful
func (m *AssignmentFilterSupportedProperty) GetDataType()(*string) {
    val, err := m.GetBackingStore().Get("dataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
                if v != nil {
                    res[i] = *(v.(*AssignmentFilterOperator))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSupportedValues(res)
        }
        return nil
    }
    return res
}
// GetIsCollection gets the isCollection property value. Indicates whether the property is a collection type or not.
// returns a *bool when successful
func (m *AssignmentFilterSupportedProperty) GetIsCollection()(*bool) {
    val, err := m.GetBackingStore().Get("isCollection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetName gets the name property value. Name of the property.
// returns a *string when successful
func (m *AssignmentFilterSupportedProperty) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AssignmentFilterSupportedProperty) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPropertyRegexConstraint gets the propertyRegexConstraint property value. Regex string to do validation on the property value.
// returns a *string when successful
func (m *AssignmentFilterSupportedProperty) GetPropertyRegexConstraint()(*string) {
    val, err := m.GetBackingStore().Get("propertyRegexConstraint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSupportedOperators gets the supportedOperators property value. List of all supported operators on this property.
// returns a []AssignmentFilterOperator when successful
func (m *AssignmentFilterSupportedProperty) GetSupportedOperators()([]AssignmentFilterOperator) {
    val, err := m.GetBackingStore().Get("supportedOperators")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AssignmentFilterOperator)
    }
    return nil
}
// GetSupportedValues gets the supportedValues property value. List of all supported values for this property, empty if everything is supported.
// returns a []string when successful
func (m *AssignmentFilterSupportedProperty) GetSupportedValues()([]string) {
    val, err := m.GetBackingStore().Get("supportedValues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterSupportedProperty) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AssignmentFilterSupportedProperty) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDataType sets the dataType property value. The data type of the property.
func (m *AssignmentFilterSupportedProperty) SetDataType(value *string)() {
    err := m.GetBackingStore().Set("dataType", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCollection sets the isCollection property value. Indicates whether the property is a collection type or not.
func (m *AssignmentFilterSupportedProperty) SetIsCollection(value *bool)() {
    err := m.GetBackingStore().Set("isCollection", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name of the property.
func (m *AssignmentFilterSupportedProperty) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterSupportedProperty) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPropertyRegexConstraint sets the propertyRegexConstraint property value. Regex string to do validation on the property value.
func (m *AssignmentFilterSupportedProperty) SetPropertyRegexConstraint(value *string)() {
    err := m.GetBackingStore().Set("propertyRegexConstraint", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedOperators sets the supportedOperators property value. List of all supported operators on this property.
func (m *AssignmentFilterSupportedProperty) SetSupportedOperators(value []AssignmentFilterOperator)() {
    err := m.GetBackingStore().Set("supportedOperators", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedValues sets the supportedValues property value. List of all supported values for this property, empty if everything is supported.
func (m *AssignmentFilterSupportedProperty) SetSupportedValues(value []string)() {
    err := m.GetBackingStore().Set("supportedValues", value)
    if err != nil {
        panic(err)
    }
}
type AssignmentFilterSupportedPropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDataType()(*string)
    GetIsCollection()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetPropertyRegexConstraint()(*string)
    GetSupportedOperators()([]AssignmentFilterOperator)
    GetSupportedValues()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDataType(value *string)()
    SetIsCollection(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPropertyRegexConstraint(value *string)()
    SetSupportedOperators(value []AssignmentFilterOperator)()
    SetSupportedValues(value []string)()
}
