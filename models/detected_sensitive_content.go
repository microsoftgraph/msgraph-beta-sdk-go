package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedSensitiveContent 
type DetectedSensitiveContent struct {
    DetectedSensitiveContentBase
}
// NewDetectedSensitiveContent instantiates a new detectedSensitiveContent and sets the default values.
func NewDetectedSensitiveContent()(*DetectedSensitiveContent) {
    m := &DetectedSensitiveContent{
        DetectedSensitiveContentBase: *NewDetectedSensitiveContentBase(),
    }
    return m
}
// CreateDetectedSensitiveContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDetectedSensitiveContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.machineLearningDetectedSensitiveContent":
                        return NewMachineLearningDetectedSensitiveContent(), nil
                }
            }
        }
    }
    return NewDetectedSensitiveContent(), nil
}
// GetClassificationAttributes gets the classificationAttributes property value. The classificationAttributes property
func (m *DetectedSensitiveContent) GetClassificationAttributes()([]ClassificationAttributeable) {
    val, err := m.GetBackingStore().Get("classificationAttributes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ClassificationAttributeable)
    }
    return nil
}
// GetClassificationMethod gets the classificationMethod property value. The classificationMethod property
func (m *DetectedSensitiveContent) GetClassificationMethod()(*DetectedSensitiveContent_classificationMethod) {
    val, err := m.GetBackingStore().Get("classificationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DetectedSensitiveContent_classificationMethod)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectedSensitiveContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DetectedSensitiveContentBase.GetFieldDeserializers()
    res["classificationAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClassificationAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassificationAttributeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ClassificationAttributeable)
                }
            }
            m.SetClassificationAttributes(res)
        }
        return nil
    }
    res["classificationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDetectedSensitiveContent_classificationMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassificationMethod(val.(*DetectedSensitiveContent_classificationMethod))
        }
        return nil
    }
    res["matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitiveContentLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveContentLocationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SensitiveContentLocationable)
                }
            }
            m.SetMatches(res)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDetectedSensitiveContent_scope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*DetectedSensitiveContent_scope))
        }
        return nil
    }
    res["sensitiveTypeSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDetectedSensitiveContent_sensitiveTypeSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitiveTypeSource(val.(*DetectedSensitiveContent_sensitiveTypeSource))
        }
        return nil
    }
    return res
}
// GetMatches gets the matches property value. The matches property
func (m *DetectedSensitiveContent) GetMatches()([]SensitiveContentLocationable) {
    val, err := m.GetBackingStore().Get("matches")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SensitiveContentLocationable)
    }
    return nil
}
// GetScope gets the scope property value. The scope property
func (m *DetectedSensitiveContent) GetScope()(*DetectedSensitiveContent_scope) {
    val, err := m.GetBackingStore().Get("scope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DetectedSensitiveContent_scope)
    }
    return nil
}
// GetSensitiveTypeSource gets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *DetectedSensitiveContent) GetSensitiveTypeSource()(*DetectedSensitiveContent_sensitiveTypeSource) {
    val, err := m.GetBackingStore().Get("sensitiveTypeSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DetectedSensitiveContent_sensitiveTypeSource)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DetectedSensitiveContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DetectedSensitiveContentBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassificationAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassificationAttributes()))
        for i, v := range m.GetClassificationAttributes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("classificationAttributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetClassificationMethod() != nil {
        cast := (*m.GetClassificationMethod()).String()
        err = writer.WriteStringValue("classificationMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatches()))
        for i, v := range m.GetMatches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("matches", cast)
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
    return nil
}
// SetClassificationAttributes sets the classificationAttributes property value. The classificationAttributes property
func (m *DetectedSensitiveContent) SetClassificationAttributes(value []ClassificationAttributeable)() {
    err := m.GetBackingStore().Set("classificationAttributes", value)
    if err != nil {
        panic(err)
    }
}
// SetClassificationMethod sets the classificationMethod property value. The classificationMethod property
func (m *DetectedSensitiveContent) SetClassificationMethod(value *DetectedSensitiveContent_classificationMethod)() {
    err := m.GetBackingStore().Set("classificationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetMatches sets the matches property value. The matches property
func (m *DetectedSensitiveContent) SetMatches(value []SensitiveContentLocationable)() {
    err := m.GetBackingStore().Set("matches", value)
    if err != nil {
        panic(err)
    }
}
// SetScope sets the scope property value. The scope property
func (m *DetectedSensitiveContent) SetScope(value *DetectedSensitiveContent_scope)() {
    err := m.GetBackingStore().Set("scope", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitiveTypeSource sets the sensitiveTypeSource property value. The sensitiveTypeSource property
func (m *DetectedSensitiveContent) SetSensitiveTypeSource(value *DetectedSensitiveContent_sensitiveTypeSource)() {
    err := m.GetBackingStore().Set("sensitiveTypeSource", value)
    if err != nil {
        panic(err)
    }
}
// DetectedSensitiveContentable 
type DetectedSensitiveContentable interface {
    DetectedSensitiveContentBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClassificationAttributes()([]ClassificationAttributeable)
    GetClassificationMethod()(*DetectedSensitiveContent_classificationMethod)
    GetMatches()([]SensitiveContentLocationable)
    GetScope()(*DetectedSensitiveContent_scope)
    GetSensitiveTypeSource()(*DetectedSensitiveContent_sensitiveTypeSource)
    SetClassificationAttributes(value []ClassificationAttributeable)()
    SetClassificationMethod(value *DetectedSensitiveContent_classificationMethod)()
    SetMatches(value []SensitiveContentLocationable)()
    SetScope(value *DetectedSensitiveContent_scope)()
    SetSensitiveTypeSource(value *DetectedSensitiveContent_sensitiveTypeSource)()
}
