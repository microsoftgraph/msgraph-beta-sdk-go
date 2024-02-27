package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationGradingScheme struct {
    Entity
}
// NewEducationGradingScheme instantiates a new EducationGradingScheme and sets the default values.
func NewEducationGradingScheme()(*EducationGradingScheme) {
    m := &EducationGradingScheme{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationGradingSchemeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationGradingSchemeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationGradingScheme(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *EducationGradingScheme) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationGradingScheme) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["grades"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationGradingSchemeGradeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationGradingSchemeGradeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationGradingSchemeGradeable)
                }
            }
            m.SetGrades(res)
        }
        return nil
    }
    res["hidePointsDuringGrading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidePointsDuringGrading(val)
        }
        return nil
    }
    return res
}
// GetGrades gets the grades property value. The grades property
// returns a []EducationGradingSchemeGradeable when successful
func (m *EducationGradingScheme) GetGrades()([]EducationGradingSchemeGradeable) {
    val, err := m.GetBackingStore().Get("grades")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EducationGradingSchemeGradeable)
    }
    return nil
}
// GetHidePointsDuringGrading gets the hidePointsDuringGrading property value. The hidePointsDuringGrading property
// returns a *bool when successful
func (m *EducationGradingScheme) GetHidePointsDuringGrading()(*bool) {
    val, err := m.GetBackingStore().Get("hidePointsDuringGrading")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationGradingScheme) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetGrades() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGrades()))
        for i, v := range m.GetGrades() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("grades", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hidePointsDuringGrading", m.GetHidePointsDuringGrading())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *EducationGradingScheme) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetGrades sets the grades property value. The grades property
func (m *EducationGradingScheme) SetGrades(value []EducationGradingSchemeGradeable)() {
    err := m.GetBackingStore().Set("grades", value)
    if err != nil {
        panic(err)
    }
}
// SetHidePointsDuringGrading sets the hidePointsDuringGrading property value. The hidePointsDuringGrading property
func (m *EducationGradingScheme) SetHidePointsDuringGrading(value *bool)() {
    err := m.GetBackingStore().Set("hidePointsDuringGrading", value)
    if err != nil {
        panic(err)
    }
}
type EducationGradingSchemeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetGrades()([]EducationGradingSchemeGradeable)
    GetHidePointsDuringGrading()(*bool)
    SetDisplayName(value *string)()
    SetGrades(value []EducationGradingSchemeGradeable)()
    SetHidePointsDuringGrading(value *bool)()
}
