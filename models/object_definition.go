package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectDefinition 
type ObjectDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The attributes property
    attributes []AttributeDefinitionable
    // The metadata property
    metadata []MetadataEntryable
    // The name property
    name *string
    // The OdataType property
    odataType *string
    // The supportedApis property
    supportedApis []string
}
// NewObjectDefinition instantiates a new objectDefinition and sets the default values.
func NewObjectDefinition()(*ObjectDefinition) {
    m := &ObjectDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.objectDefinition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateObjectDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectDefinition) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttributes gets the attributes property value. The attributes property
func (m *ObjectDefinition) GetAttributes()([]AttributeDefinitionable) {
    return m.attributes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAttributeDefinitionFromDiscriminatorValue , m.SetAttributes)
    res["metadata"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMetadataEntryFromDiscriminatorValue , m.SetMetadata)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["supportedApis"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSupportedApis)
    return res
}
// GetMetadata gets the metadata property value. The metadata property
func (m *ObjectDefinition) GetMetadata()([]MetadataEntryable) {
    return m.metadata
}
// GetName gets the name property value. The name property
func (m *ObjectDefinition) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ObjectDefinition) GetOdataType()(*string) {
    return m.odataType
}
// GetSupportedApis gets the supportedApis property value. The supportedApis property
func (m *ObjectDefinition) GetSupportedApis()([]string) {
    return m.supportedApis
}
// Serialize serializes information the current object
func (m *ObjectDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttributes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAttributes())
        err := writer.WriteCollectionOfObjectValues("attributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMetadata())
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
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
    if m.GetSupportedApis() != nil {
        err := writer.WriteCollectionOfStringValues("supportedApis", m.GetSupportedApis())
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
func (m *ObjectDefinition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttributes sets the attributes property value. The attributes property
func (m *ObjectDefinition) SetAttributes(value []AttributeDefinitionable)() {
    m.attributes = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *ObjectDefinition) SetMetadata(value []MetadataEntryable)() {
    m.metadata = value
}
// SetName sets the name property value. The name property
func (m *ObjectDefinition) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ObjectDefinition) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSupportedApis sets the supportedApis property value. The supportedApis property
func (m *ObjectDefinition) SetSupportedApis(value []string)() {
    m.supportedApis = value
}
