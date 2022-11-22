package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMapping 
type AttributeMapping struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Default value to be used in case the source property was evaluated to null. Optional.
    defaultValue *string
    // For internal use only.
    exportMissingReferences *bool
    // The flowBehavior property
    flowBehavior *AttributeFlowBehavior
    // The flowType property
    flowType *AttributeFlowType
    // If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
    matchingPriority *int32
    // The OdataType property
    odataType *string
    // Defines how a value should be extracted (or transformed) from the source object.
    source AttributeMappingSourceable
    // Name of the attribute on the target object.
    targetAttributeName *string
}
// NewAttributeMapping instantiates a new attributeMapping and sets the default values.
func NewAttributeMapping()(*AttributeMapping) {
    m := &AttributeMapping{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttributeMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMapping(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttributeMapping) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultValue gets the defaultValue property value. Default value to be used in case the source property was evaluated to null. Optional.
func (m *AttributeMapping) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetExportMissingReferences gets the exportMissingReferences property value. For internal use only.
func (m *AttributeMapping) GetExportMissingReferences()(*bool) {
    return m.exportMissingReferences
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultValue)
    res["exportMissingReferences"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetExportMissingReferences)
    res["flowBehavior"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAttributeFlowBehavior , m.SetFlowBehavior)
    res["flowType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAttributeFlowType , m.SetFlowType)
    res["matchingPriority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMatchingPriority)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["source"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAttributeMappingSourceFromDiscriminatorValue , m.SetSource)
    res["targetAttributeName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTargetAttributeName)
    return res
}
// GetFlowBehavior gets the flowBehavior property value. The flowBehavior property
func (m *AttributeMapping) GetFlowBehavior()(*AttributeFlowBehavior) {
    return m.flowBehavior
}
// GetFlowType gets the flowType property value. The flowType property
func (m *AttributeMapping) GetFlowType()(*AttributeFlowType) {
    return m.flowType
}
// GetMatchingPriority gets the matchingPriority property value. If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
func (m *AttributeMapping) GetMatchingPriority()(*int32) {
    return m.matchingPriority
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttributeMapping) GetOdataType()(*string) {
    return m.odataType
}
// GetSource gets the source property value. Defines how a value should be extracted (or transformed) from the source object.
func (m *AttributeMapping) GetSource()(AttributeMappingSourceable) {
    return m.source
}
// GetTargetAttributeName gets the targetAttributeName property value. Name of the attribute on the target object.
func (m *AttributeMapping) GetTargetAttributeName()(*string) {
    return m.targetAttributeName
}
// Serialize serializes information the current object
func (m *AttributeMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exportMissingReferences", m.GetExportMissingReferences())
        if err != nil {
            return err
        }
    }
    if m.GetFlowBehavior() != nil {
        cast := (*m.GetFlowBehavior()).String()
        err := writer.WriteStringValue("flowBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFlowType() != nil {
        cast := (*m.GetFlowType()).String()
        err := writer.WriteStringValue("flowType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("matchingPriority", m.GetMatchingPriority())
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
        err := writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetAttributeName", m.GetTargetAttributeName())
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
func (m *AttributeMapping) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultValue sets the defaultValue property value. Default value to be used in case the source property was evaluated to null. Optional.
func (m *AttributeMapping) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetExportMissingReferences sets the exportMissingReferences property value. For internal use only.
func (m *AttributeMapping) SetExportMissingReferences(value *bool)() {
    m.exportMissingReferences = value
}
// SetFlowBehavior sets the flowBehavior property value. The flowBehavior property
func (m *AttributeMapping) SetFlowBehavior(value *AttributeFlowBehavior)() {
    m.flowBehavior = value
}
// SetFlowType sets the flowType property value. The flowType property
func (m *AttributeMapping) SetFlowType(value *AttributeFlowType)() {
    m.flowType = value
}
// SetMatchingPriority sets the matchingPriority property value. If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
func (m *AttributeMapping) SetMatchingPriority(value *int32)() {
    m.matchingPriority = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttributeMapping) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSource sets the source property value. Defines how a value should be extracted (or transformed) from the source object.
func (m *AttributeMapping) SetSource(value AttributeMappingSourceable)() {
    m.source = value
}
// SetTargetAttributeName sets the targetAttributeName property value. Name of the attribute on the target object.
func (m *AttributeMapping) SetTargetAttributeName(value *string)() {
    m.targetAttributeName = value
}
