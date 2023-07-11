package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanConfigurationLocalization 
type PlannerPlanConfigurationLocalization struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewPlannerPlanConfigurationLocalization instantiates a new plannerPlanConfigurationLocalization and sets the default values.
func NewPlannerPlanConfigurationLocalization()(*PlannerPlanConfigurationLocalization) {
    m := &PlannerPlanConfigurationLocalization{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanConfigurationLocalizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanConfigurationLocalization(), nil
}
// GetBuckets gets the buckets property value. Localized names for configured buckets in the plan configuration.
func (m *PlannerPlanConfigurationLocalization) GetBuckets()([]PlannerPlanConfigurationBucketLocalizationable) {
    val, err := m.GetBackingStore().Get("buckets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerPlanConfigurationBucketLocalizationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanConfigurationLocalization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buckets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanConfigurationBucketLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanConfigurationBucketLocalizationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PlannerPlanConfigurationBucketLocalizationable)
                }
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["planTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanTitle(val)
        }
        return nil
    }
    return res
}
// GetLanguageTag gets the languageTag property value. The language code associated with the localized names in this object.
func (m *PlannerPlanConfigurationLocalization) GetLanguageTag()(*string) {
    val, err := m.GetBackingStore().Get("languageTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlanTitle gets the planTitle property value. Localized title of the plan.
func (m *PlannerPlanConfigurationLocalization) GetPlanTitle()(*string) {
    val, err := m.GetBackingStore().Get("planTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerPlanConfigurationLocalization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBuckets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBuckets()))
        for i, v := range m.GetBuckets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("planTitle", m.GetPlanTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuckets sets the buckets property value. Localized names for configured buckets in the plan configuration.
func (m *PlannerPlanConfigurationLocalization) SetBuckets(value []PlannerPlanConfigurationBucketLocalizationable)() {
    err := m.GetBackingStore().Set("buckets", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguageTag sets the languageTag property value. The language code associated with the localized names in this object.
func (m *PlannerPlanConfigurationLocalization) SetLanguageTag(value *string)() {
    err := m.GetBackingStore().Set("languageTag", value)
    if err != nil {
        panic(err)
    }
}
// SetPlanTitle sets the planTitle property value. Localized title of the plan.
func (m *PlannerPlanConfigurationLocalization) SetPlanTitle(value *string)() {
    err := m.GetBackingStore().Set("planTitle", value)
    if err != nil {
        panic(err)
    }
}
// PlannerPlanConfigurationLocalizationable 
type PlannerPlanConfigurationLocalizationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuckets()([]PlannerPlanConfigurationBucketLocalizationable)
    GetLanguageTag()(*string)
    GetPlanTitle()(*string)
    SetBuckets(value []PlannerPlanConfigurationBucketLocalizationable)()
    SetLanguageTag(value *string)()
    SetPlanTitle(value *string)()
}
