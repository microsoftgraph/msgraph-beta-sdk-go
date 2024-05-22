package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RegexReplaceTransformation struct {
    CustomClaimTransformation
}
// NewRegexReplaceTransformation instantiates a new RegexReplaceTransformation and sets the default values.
func NewRegexReplaceTransformation()(*RegexReplaceTransformation) {
    m := &RegexReplaceTransformation{
        CustomClaimTransformation: *NewCustomClaimTransformation(),
    }
    odataTypeValue := "#microsoft.graph.regexReplaceTransformation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRegexReplaceTransformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRegexReplaceTransformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegexReplaceTransformation(), nil
}
// GetAdditionalAttributes gets the additionalAttributes property value. Additional attributes that can be referenced within the replacement string.
// returns a []SourcedAttributeable when successful
func (m *RegexReplaceTransformation) GetAdditionalAttributes()([]SourcedAttributeable) {
    val, err := m.GetBackingStore().Get("additionalAttributes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SourcedAttributeable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RegexReplaceTransformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimTransformation.GetFieldDeserializers()
    res["additionalAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSourcedAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SourcedAttributeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SourcedAttributeable)
                }
            }
            m.SetAdditionalAttributes(res)
        }
        return nil
    }
    res["regex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegex(val)
        }
        return nil
    }
    res["replacement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplacement(val)
        }
        return nil
    }
    return res
}
// GetRegex gets the regex property value. The regular expression to be applied on the input directory attribute or constant.
// returns a *string when successful
func (m *RegexReplaceTransformation) GetRegex()(*string) {
    val, err := m.GetBackingStore().Get("regex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReplacement gets the replacement property value. The transformation output replacement pattern with regular expression output group and input parameter group reference.
// returns a *string when successful
func (m *RegexReplaceTransformation) GetReplacement()(*string) {
    val, err := m.GetBackingStore().Get("replacement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RegexReplaceTransformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimTransformation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalAttributes()))
        for i, v := range m.GetAdditionalAttributes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("additionalAttributes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("regex", m.GetRegex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("replacement", m.GetReplacement())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalAttributes sets the additionalAttributes property value. Additional attributes that can be referenced within the replacement string.
func (m *RegexReplaceTransformation) SetAdditionalAttributes(value []SourcedAttributeable)() {
    err := m.GetBackingStore().Set("additionalAttributes", value)
    if err != nil {
        panic(err)
    }
}
// SetRegex sets the regex property value. The regular expression to be applied on the input directory attribute or constant.
func (m *RegexReplaceTransformation) SetRegex(value *string)() {
    err := m.GetBackingStore().Set("regex", value)
    if err != nil {
        panic(err)
    }
}
// SetReplacement sets the replacement property value. The transformation output replacement pattern with regular expression output group and input parameter group reference.
func (m *RegexReplaceTransformation) SetReplacement(value *string)() {
    err := m.GetBackingStore().Set("replacement", value)
    if err != nil {
        panic(err)
    }
}
type RegexReplaceTransformationable interface {
    CustomClaimTransformationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalAttributes()([]SourcedAttributeable)
    GetRegex()(*string)
    GetReplacement()(*string)
    SetAdditionalAttributes(value []SourcedAttributeable)()
    SetRegex(value *string)()
    SetReplacement(value *string)()
}
