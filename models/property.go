package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Property 
type Property struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The aliases property
    aliases []string
    // The isQueryable property
    isQueryable *bool
    // The isRefinable property
    isRefinable *bool
    // The isRetrievable property
    isRetrievable *bool
    // The isSearchable property
    isSearchable *bool
    // The labels property
    labels []string
    // The name property
    name *string
    // The OdataType property
    odataType *string
    // The type property
    type_escaped *PropertyType
}
// NewProperty instantiates a new property and sets the default values.
func NewProperty()(*Property) {
    m := &Property{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.property";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Property) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAliases gets the aliases property value. The aliases property
func (m *Property) GetAliases()([]string) {
    if m == nil {
        return nil
    } else {
        return m.aliases
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Property) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aliases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAliases(res)
        }
        return nil
    }
    res["isQueryable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsQueryable(val)
        }
        return nil
    }
    res["isRefinable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRefinable(val)
        }
        return nil
    }
    res["isRetrievable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRetrievable(val)
        }
        return nil
    }
    res["isSearchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
        }
        return nil
    }
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLabels(res)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePropertyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*PropertyType))
        }
        return nil
    }
    return res
}
// GetIsQueryable gets the isQueryable property value. The isQueryable property
func (m *Property) GetIsQueryable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isQueryable
    }
}
// GetIsRefinable gets the isRefinable property value. The isRefinable property
func (m *Property) GetIsRefinable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRefinable
    }
}
// GetIsRetrievable gets the isRetrievable property value. The isRetrievable property
func (m *Property) GetIsRetrievable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRetrievable
    }
}
// GetIsSearchable gets the isSearchable property value. The isSearchable property
func (m *Property) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// GetLabels gets the labels property value. The labels property
func (m *Property) GetLabels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.labels
    }
}
// GetName gets the name property value. The name property
func (m *Property) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Property) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetType gets the type property value. The type property
func (m *Property) GetType()(*PropertyType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *Property) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAliases() != nil {
        err := writer.WriteCollectionOfStringValues("aliases", m.GetAliases())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isQueryable", m.GetIsQueryable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRefinable", m.GetIsRefinable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRetrievable", m.GetIsRetrievable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    if m.GetLabels() != nil {
        err := writer.WriteCollectionOfStringValues("labels", m.GetLabels())
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *Property) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAliases sets the aliases property value. The aliases property
func (m *Property) SetAliases(value []string)() {
    if m != nil {
        m.aliases = value
    }
}
// SetIsQueryable sets the isQueryable property value. The isQueryable property
func (m *Property) SetIsQueryable(value *bool)() {
    if m != nil {
        m.isQueryable = value
    }
}
// SetIsRefinable sets the isRefinable property value. The isRefinable property
func (m *Property) SetIsRefinable(value *bool)() {
    if m != nil {
        m.isRefinable = value
    }
}
// SetIsRetrievable sets the isRetrievable property value. The isRetrievable property
func (m *Property) SetIsRetrievable(value *bool)() {
    if m != nil {
        m.isRetrievable = value
    }
}
// SetIsSearchable sets the isSearchable property value. The isSearchable property
func (m *Property) SetIsSearchable(value *bool)() {
    if m != nil {
        m.isSearchable = value
    }
}
// SetLabels sets the labels property value. The labels property
func (m *Property) SetLabels(value []string)() {
    if m != nil {
        m.labels = value
    }
}
// SetName sets the name property value. The name property
func (m *Property) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Property) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetType sets the type property value. The type property
func (m *Property) SetType(value *PropertyType)() {
    if m != nil {
        m.type_escaped = value
    }
}
