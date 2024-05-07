package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type CustomClaimTransformation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCustomClaimTransformation instantiates a new CustomClaimTransformation and sets the default values.
func NewCustomClaimTransformation()(*CustomClaimTransformation) {
    m := &CustomClaimTransformation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomClaimTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomClaimTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.containsTransformation":
                        return NewContainsTransformation(), nil
                    case "#microsoft.graph.endsWithTransformation":
                        return NewEndsWithTransformation(), nil
                    case "#microsoft.graph.extractAlphaTransformation":
                        return NewExtractAlphaTransformation(), nil
                    case "#microsoft.graph.extractMailPrefixTransformation":
                        return NewExtractMailPrefixTransformation(), nil
                    case "#microsoft.graph.extractNumberTransformation":
                        return NewExtractNumberTransformation(), nil
                    case "#microsoft.graph.extractTransformation":
                        return NewExtractTransformation(), nil
                    case "#microsoft.graph.ifEmptyTransformation":
                        return NewIfEmptyTransformation(), nil
                    case "#microsoft.graph.ifNotEmptyTransformation":
                        return NewIfNotEmptyTransformation(), nil
                    case "#microsoft.graph.joinTransformation":
                        return NewJoinTransformation(), nil
                    case "#microsoft.graph.regexReplaceTransformation":
                        return NewRegexReplaceTransformation(), nil
                    case "#microsoft.graph.startsWithTransformation":
                        return NewStartsWithTransformation(), nil
                    case "#microsoft.graph.substringTransformation":
                        return NewSubstringTransformation(), nil
                    case "#microsoft.graph.toLowercaseTransformation":
                        return NewToLowercaseTransformation(), nil
                    case "#microsoft.graph.toUppercaseTransformation":
                        return NewToUppercaseTransformation(), nil
                    case "#microsoft.graph.trimTransformation":
                        return NewTrimTransformation(), nil
                }
            }
        }
    }
    return NewCustomClaimTransformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CustomClaimTransformation) GetAdditionalData()(map[string]any) {
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
func (m *CustomClaimTransformation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomClaimTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["input"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransformationAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInput(val.(TransformationAttributeable))
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
    return res
}
// GetInput gets the input property value. The input property
// returns a TransformationAttributeable when successful
func (m *CustomClaimTransformation) GetInput()(TransformationAttributeable) {
    val, err := m.GetBackingStore().Get("input")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TransformationAttributeable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *CustomClaimTransformation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomClaimTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("input", m.GetInput())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomClaimTransformation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *CustomClaimTransformation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetInput sets the input property value. The input property
func (m *CustomClaimTransformation) SetInput(value TransformationAttributeable)() {
    err := m.GetBackingStore().Set("input", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomClaimTransformation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type CustomClaimTransformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetInput()(TransformationAttributeable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetInput(value TransformationAttributeable)()
    SetOdataType(value *string)()
}
