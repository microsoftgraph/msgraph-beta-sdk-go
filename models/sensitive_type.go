package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitiveType provides operations to manage the collection of accessReview entities.
type SensitiveType struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSensitiveTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitiveTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitiveType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClassificationMethod gets the classificationMethod property value. The classificationMethod property
func (m *SensitiveType) GetClassificationMethod()(*ClassificationMethod) {
    if m == nil {
        return nil
    } else {
        return m.classificationMethod
    }
}
// GetDescription gets the description property value. The description property
func (m *SensitiveType) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitiveType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classificationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClassificationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassificationMethod(val.(*ClassificationMethod))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["publisherName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    res["rulePackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulePackageId(val)
        }
        return nil
    }
    res["rulePackageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulePackageType(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*SensitiveTypeScope))
        }
        return nil
    }
    res["sensitiveTypeSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitiveTypeSource(val.(*SensitiveTypeSource))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *SensitiveType) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPublisherName gets the publisherName property value. The publisherName property
func (m *SensitiveType) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// GetRulePackageId gets the rulePackageId property value. The rulePackageId property
func (m *SensitiveType) GetRulePackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageId
    }
}
// GetRulePackageType gets the rulePackageType property value. The rulePackageType property
func (m *SensitiveType) GetRulePackageType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rulePackageType
    }
}
// GetScope gets the scope property value. The scope property
func (m *SensitiveType) GetScope()(*SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetSensitiveTypeSource gets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *SensitiveType) GetSensitiveTypeSource()(*SensitiveTypeSource) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeSource
    }
}
// GetState gets the state property value. The state property
func (m *SensitiveType) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SensitiveType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClassificationMethod sets the classificationMethod property value. The classificationMethod property
func (m *SensitiveType) SetClassificationMethod(value *ClassificationMethod)() {
    if m != nil {
        m.classificationMethod = value
    }
}
// SetDescription sets the description property value. The description property
func (m *SensitiveType) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetName sets the name property value. The name property
func (m *SensitiveType) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPublisherName sets the publisherName property value. The publisherName property
func (m *SensitiveType) SetPublisherName(value *string)() {
    if m != nil {
        m.publisherName = value
    }
}
// SetRulePackageId sets the rulePackageId property value. The rulePackageId property
func (m *SensitiveType) SetRulePackageId(value *string)() {
    if m != nil {
        m.rulePackageId = value
    }
}
// SetRulePackageType sets the rulePackageType property value. The rulePackageType property
func (m *SensitiveType) SetRulePackageType(value *string)() {
    if m != nil {
        m.rulePackageType = value
    }
}
// SetScope sets the scope property value. The scope property
func (m *SensitiveType) SetScope(value *SensitiveTypeScope)() {
    if m != nil {
        m.scope = value
    }
}
// SetSensitiveTypeSource sets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *SensitiveType) SetSensitiveTypeSource(value *SensitiveTypeSource)() {
    if m != nil {
        m.sensitiveTypeSource = value
    }
}
// SetState sets the state property value. The state property
func (m *SensitiveType) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
