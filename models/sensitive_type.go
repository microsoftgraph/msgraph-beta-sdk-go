package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitiveType provides operations to manage the collection of accessReviewDecision entities.
type SensitiveType struct {
    Entity
    // The classificationMethod property
    classificationMethod *ClassificationMethod
    // The description property
    description *string
    // The name property
    name *string
    // The publisherName property
    publisherName *string
    // The rulePackageId property
    rulePackageId *string
    // The rulePackageType property
    rulePackageType *string
    // The scope property
    scope *SensitiveTypeScope
    // The sensitiveTypeSource property
    sensitiveTypeSource *SensitiveTypeSource
    // The state property
    state *string
}
// NewSensitiveType instantiates a new sensitiveType and sets the default values.
func NewSensitiveType()(*SensitiveType) {
    m := &SensitiveType{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.sensitiveType";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSensitiveTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitiveTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitiveType(), nil
}
// GetClassificationMethod gets the classificationMethod property value. The classificationMethod property
func (m *SensitiveType) GetClassificationMethod()(*ClassificationMethod) {
    return m.classificationMethod
}
// GetDescription gets the description property value. The description property
func (m *SensitiveType) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitiveType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classificationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseClassificationMethod , m.SetClassificationMethod)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["publisherName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublisherName)
    res["rulePackageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRulePackageId)
    res["rulePackageType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRulePackageType)
    res["scope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSensitiveTypeScope , m.SetScope)
    res["sensitiveTypeSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSensitiveTypeSource , m.SetSensitiveTypeSource)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetState)
    return res
}
// GetName gets the name property value. The name property
func (m *SensitiveType) GetName()(*string) {
    return m.name
}
// GetPublisherName gets the publisherName property value. The publisherName property
func (m *SensitiveType) GetPublisherName()(*string) {
    return m.publisherName
}
// GetRulePackageId gets the rulePackageId property value. The rulePackageId property
func (m *SensitiveType) GetRulePackageId()(*string) {
    return m.rulePackageId
}
// GetRulePackageType gets the rulePackageType property value. The rulePackageType property
func (m *SensitiveType) GetRulePackageType()(*string) {
    return m.rulePackageType
}
// GetScope gets the scope property value. The scope property
func (m *SensitiveType) GetScope()(*SensitiveTypeScope) {
    return m.scope
}
// GetSensitiveTypeSource gets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *SensitiveType) GetSensitiveTypeSource()(*SensitiveTypeSource) {
    return m.sensitiveTypeSource
}
// GetState gets the state property value. The state property
func (m *SensitiveType) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *SensitiveType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassificationMethod() != nil {
        cast := (*m.GetClassificationMethod()).String()
        err = writer.WriteStringValue("classificationMethod", &cast)
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherName", m.GetPublisherName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rulePackageId", m.GetRulePackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rulePackageType", m.GetRulePackageType())
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := (*m.GetScope()).String()
        err = writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeSource() != nil {
        cast := (*m.GetSensitiveTypeSource()).String()
        err = writer.WriteStringValue("sensitiveTypeSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassificationMethod sets the classificationMethod property value. The classificationMethod property
func (m *SensitiveType) SetClassificationMethod(value *ClassificationMethod)() {
    m.classificationMethod = value
}
// SetDescription sets the description property value. The description property
func (m *SensitiveType) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. The name property
func (m *SensitiveType) SetName(value *string)() {
    m.name = value
}
// SetPublisherName sets the publisherName property value. The publisherName property
func (m *SensitiveType) SetPublisherName(value *string)() {
    m.publisherName = value
}
// SetRulePackageId sets the rulePackageId property value. The rulePackageId property
func (m *SensitiveType) SetRulePackageId(value *string)() {
    m.rulePackageId = value
}
// SetRulePackageType sets the rulePackageType property value. The rulePackageType property
func (m *SensitiveType) SetRulePackageType(value *string)() {
    m.rulePackageType = value
}
// SetScope sets the scope property value. The scope property
func (m *SensitiveType) SetScope(value *SensitiveTypeScope)() {
    m.scope = value
}
// SetSensitiveTypeSource sets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *SensitiveType) SetSensitiveTypeSource(value *SensitiveTypeSource)() {
    m.sensitiveTypeSource = value
}
// SetState sets the state property value. The state property
func (m *SensitiveType) SetState(value *string)() {
    m.state = value
}
