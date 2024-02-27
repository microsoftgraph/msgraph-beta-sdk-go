package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationAssignmentSettings struct {
    Entity
}
// NewEducationAssignmentSettings instantiates a new EducationAssignmentSettings and sets the default values.
func NewEducationAssignmentSettings()(*EducationAssignmentSettings) {
    m := &EducationAssignmentSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentSettings(), nil
}
// GetDefaultGradingScheme gets the defaultGradingScheme property value. The defaultGradingScheme property
// returns a EducationGradingSchemeable when successful
func (m *EducationAssignmentSettings) GetDefaultGradingScheme()(EducationGradingSchemeable) {
    val, err := m.GetBackingStore().Get("defaultGradingScheme")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationGradingSchemeable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultGradingScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationGradingSchemeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultGradingScheme(val.(EducationGradingSchemeable))
        }
        return nil
    }
    res["gradingCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationGradingCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationGradingCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationGradingCategoryable)
                }
            }
            m.SetGradingCategories(res)
        }
        return nil
    }
    res["gradingSchemes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationGradingSchemeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationGradingSchemeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationGradingSchemeable)
                }
            }
            m.SetGradingSchemes(res)
        }
        return nil
    }
    res["submissionAnimationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmissionAnimationDisabled(val)
        }
        return nil
    }
    return res
}
// GetGradingCategories gets the gradingCategories property value. When set, enables users to weight assignments differently when computing a class average grade.
// returns a []EducationGradingCategoryable when successful
func (m *EducationAssignmentSettings) GetGradingCategories()([]EducationGradingCategoryable) {
    val, err := m.GetBackingStore().Get("gradingCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EducationGradingCategoryable)
    }
    return nil
}
// GetGradingSchemes gets the gradingSchemes property value. The gradingSchemes property
// returns a []EducationGradingSchemeable when successful
func (m *EducationAssignmentSettings) GetGradingSchemes()([]EducationGradingSchemeable) {
    val, err := m.GetBackingStore().Get("gradingSchemes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EducationGradingSchemeable)
    }
    return nil
}
// GetSubmissionAnimationDisabled gets the submissionAnimationDisabled property value. Indicates whether turn-in celebration animation will be shown. If true, the animation won't be shown. The default value is false.
// returns a *bool when successful
func (m *EducationAssignmentSettings) GetSubmissionAnimationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("submissionAnimationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultGradingScheme", m.GetDefaultGradingScheme())
        if err != nil {
            return err
        }
    }
    if m.GetGradingCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGradingCategories()))
        for i, v := range m.GetGradingCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("gradingCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGradingSchemes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGradingSchemes()))
        for i, v := range m.GetGradingSchemes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("gradingSchemes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("submissionAnimationDisabled", m.GetSubmissionAnimationDisabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultGradingScheme sets the defaultGradingScheme property value. The defaultGradingScheme property
func (m *EducationAssignmentSettings) SetDefaultGradingScheme(value EducationGradingSchemeable)() {
    err := m.GetBackingStore().Set("defaultGradingScheme", value)
    if err != nil {
        panic(err)
    }
}
// SetGradingCategories sets the gradingCategories property value. When set, enables users to weight assignments differently when computing a class average grade.
func (m *EducationAssignmentSettings) SetGradingCategories(value []EducationGradingCategoryable)() {
    err := m.GetBackingStore().Set("gradingCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetGradingSchemes sets the gradingSchemes property value. The gradingSchemes property
func (m *EducationAssignmentSettings) SetGradingSchemes(value []EducationGradingSchemeable)() {
    err := m.GetBackingStore().Set("gradingSchemes", value)
    if err != nil {
        panic(err)
    }
}
// SetSubmissionAnimationDisabled sets the submissionAnimationDisabled property value. Indicates whether turn-in celebration animation will be shown. If true, the animation won't be shown. The default value is false.
func (m *EducationAssignmentSettings) SetSubmissionAnimationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("submissionAnimationDisabled", value)
    if err != nil {
        panic(err)
    }
}
type EducationAssignmentSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultGradingScheme()(EducationGradingSchemeable)
    GetGradingCategories()([]EducationGradingCategoryable)
    GetGradingSchemes()([]EducationGradingSchemeable)
    GetSubmissionAnimationDisabled()(*bool)
    SetDefaultGradingScheme(value EducationGradingSchemeable)()
    SetGradingCategories(value []EducationGradingCategoryable)()
    SetGradingSchemes(value []EducationGradingSchemeable)()
    SetSubmissionAnimationDisabled(value *bool)()
}
