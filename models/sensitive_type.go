package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitiveType 
type SensitiveType struct {
    Entity
}
// NewSensitiveType instantiates a new sensitiveType and sets the default values.
func NewSensitiveType()(*SensitiveType) {
    m := &SensitiveType{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSensitiveTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitiveTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitiveType(), nil
}
// GetClassificationMethod gets the classificationMethod property value. The classificationMethod property
func (m *SensitiveType) GetClassificationMethod()(*ClassificationMethod) {
    val, err := m.GetBackingStore().Get("classificationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ClassificationMethod)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *SensitiveType) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublisherName gets the publisherName property value. The publisherName property
func (m *SensitiveType) GetPublisherName()(*string) {
    val, err := m.GetBackingStore().Get("publisherName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRulePackageId gets the rulePackageId property value. The rulePackageId property
func (m *SensitiveType) GetRulePackageId()(*string) {
    val, err := m.GetBackingStore().Get("rulePackageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRulePackageType gets the rulePackageType property value. The rulePackageType property
func (m *SensitiveType) GetRulePackageType()(*string) {
    val, err := m.GetBackingStore().Get("rulePackageType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScope gets the scope property value. The scope property
func (m *SensitiveType) GetScope()(*SensitiveTypeScope) {
    val, err := m.GetBackingStore().Get("scope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SensitiveTypeScope)
    }
    return nil
}
// GetSensitiveTypeSource gets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *SensitiveType) GetSensitiveTypeSource()(*SensitiveTypeSource) {
    val, err := m.GetBackingStore().Get("sensitiveTypeSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SensitiveTypeSource)
    }
    return nil
}
// GetState gets the state property value. The state property
func (m *SensitiveType) GetState()(*string) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("classificationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *SensitiveType) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *SensitiveType) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisherName sets the publisherName property value. The publisherName property
func (m *SensitiveType) SetPublisherName(value *string)() {
    err := m.GetBackingStore().Set("publisherName", value)
    if err != nil {
        panic(err)
    }
}
// SetRulePackageId sets the rulePackageId property value. The rulePackageId property
func (m *SensitiveType) SetRulePackageId(value *string)() {
    err := m.GetBackingStore().Set("rulePackageId", value)
    if err != nil {
        panic(err)
    }
}
// SetRulePackageType sets the rulePackageType property value. The rulePackageType property
func (m *SensitiveType) SetRulePackageType(value *string)() {
    err := m.GetBackingStore().Set("rulePackageType", value)
    if err != nil {
        panic(err)
    }
}
// SetScope sets the scope property value. The scope property
func (m *SensitiveType) SetScope(value *SensitiveTypeScope)() {
    err := m.GetBackingStore().Set("scope", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitiveTypeSource sets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *SensitiveType) SetSensitiveTypeSource(value *SensitiveTypeSource)() {
    err := m.GetBackingStore().Set("sensitiveTypeSource", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *SensitiveType) SetState(value *string)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SensitiveTypeable 
type SensitiveTypeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassificationMethod()(*ClassificationMethod)
    GetDescription()(*string)
    GetName()(*string)
    GetPublisherName()(*string)
    GetRulePackageId()(*string)
    GetRulePackageType()(*string)
    GetScope()(*SensitiveTypeScope)
    GetSensitiveTypeSource()(*SensitiveTypeSource)
    GetState()(*string)
    SetClassificationMethod(value *ClassificationMethod)()
    SetDescription(value *string)()
    SetName(value *string)()
    SetPublisherName(value *string)()
    SetRulePackageId(value *string)()
    SetRulePackageType(value *string)()
    SetScope(value *SensitiveTypeScope)()
    SetSensitiveTypeSource(value *SensitiveTypeSource)()
    SetState(value *string)()
}
