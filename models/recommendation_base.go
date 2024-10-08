package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type RecommendationBase struct {
    Entity
}
// RecommendationBase_RecommendationBase_currentScore composed type wrapper for classes float64, ReferenceNumeric, string
type RecommendationBase_RecommendationBase_currentScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRecommendationBase_RecommendationBase_currentScore instantiates a new RecommendationBase_RecommendationBase_currentScore and sets the default values.
func NewRecommendationBase_RecommendationBase_currentScore()(*RecommendationBase_RecommendationBase_currentScore) {
    m := &RecommendationBase_RecommendationBase_currentScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateRecommendationBase_RecommendationBase_currentScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecommendationBase_RecommendationBase_currentScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewRecommendationBase_RecommendationBase_currentScore()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *RecommendationBase_RecommendationBase_currentScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *RecommendationBase_RecommendationBase_currentScore) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecommendationBase_RecommendationBase_currentScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *RecommendationBase_RecommendationBase_currentScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *RecommendationBase_RecommendationBase_currentScore) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *RecommendationBase_RecommendationBase_currentScore) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RecommendationBase_RecommendationBase_currentScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *RecommendationBase_RecommendationBase_currentScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *RecommendationBase_RecommendationBase_currentScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *RecommendationBase_RecommendationBase_currentScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *RecommendationBase_RecommendationBase_currentScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// RecommendationBase_RecommendationBase_maxScore composed type wrapper for classes float64, ReferenceNumeric, string
type RecommendationBase_RecommendationBase_maxScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRecommendationBase_RecommendationBase_maxScore instantiates a new RecommendationBase_RecommendationBase_maxScore and sets the default values.
func NewRecommendationBase_RecommendationBase_maxScore()(*RecommendationBase_RecommendationBase_maxScore) {
    m := &RecommendationBase_RecommendationBase_maxScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateRecommendationBase_RecommendationBase_maxScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecommendationBase_RecommendationBase_maxScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewRecommendationBase_RecommendationBase_maxScore()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *RecommendationBase_RecommendationBase_maxScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *RecommendationBase_RecommendationBase_maxScore) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecommendationBase_RecommendationBase_maxScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *RecommendationBase_RecommendationBase_maxScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *RecommendationBase_RecommendationBase_maxScore) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *RecommendationBase_RecommendationBase_maxScore) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RecommendationBase_RecommendationBase_maxScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *RecommendationBase_RecommendationBase_maxScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *RecommendationBase_RecommendationBase_maxScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *RecommendationBase_RecommendationBase_maxScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *RecommendationBase_RecommendationBase_maxScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type RecommendationBase_RecommendationBase_currentScoreable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type RecommendationBase_RecommendationBase_maxScoreable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewRecommendationBase instantiates a new RecommendationBase and sets the default values.
func NewRecommendationBase()(*RecommendationBase) {
    m := &RecommendationBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRecommendationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRecommendationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.recommendation":
                        return NewRecommendation(), nil
                }
            }
        }
    }
    return NewRecommendationBase(), nil
}
// GetActionSteps gets the actionSteps property value. List of actions to take to complete a recommendation.
// returns a []ActionStepable when successful
func (m *RecommendationBase) GetActionSteps()([]ActionStepable) {
    val, err := m.GetBackingStore().Get("actionSteps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ActionStepable)
    }
    return nil
}
// GetBenefits gets the benefits property value. An explanation of why completing the recommendation will benefit you. Corresponds to the Value section of a recommendation shown in the Microsoft Entra admin center.
// returns a *string when successful
func (m *RecommendationBase) GetBenefits()(*string) {
    val, err := m.GetBackingStore().Get("benefits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCategory gets the category property value. The category property
// returns a *RecommendationCategory when successful
func (m *RecommendationBase) GetCategory()(*RecommendationCategory) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationCategory)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the recommendation was detected as applicable to your directory.
// returns a *Time when successful
func (m *RecommendationBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCurrentScore gets the currentScore property value. The number of points the tenant has attained. Only applies to recommendations with category set to identitySecureScore.
// returns a RecommendationBase_RecommendationBase_currentScoreable when successful
func (m *RecommendationBase) GetCurrentScore()(RecommendationBase_RecommendationBase_currentScoreable) {
    val, err := m.GetBackingStore().Get("currentScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RecommendationBase_RecommendationBase_currentScoreable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The title of the recommendation.
// returns a *string when successful
func (m *RecommendationBase) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFeatureAreas gets the featureAreas property value. The directory feature that the recommendation is related to.
// returns a []RecommendationFeatureAreas when successful
func (m *RecommendationBase) GetFeatureAreas()([]RecommendationFeatureAreas) {
    val, err := m.GetBackingStore().Get("featureAreas")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RecommendationFeatureAreas)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RecommendationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActionStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActionStepable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActionStepable)
                }
            }
            m.SetActionSteps(res)
        }
        return nil
    }
    res["benefits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBenefits(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*RecommendationCategory))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["currentScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecommendationBase_RecommendationBase_currentScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentScore(val.(*RecommendationBase_RecommendationBase_currentScore))
        }
        return nil
    }
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
    res["featureAreas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRecommendationFeatureAreas)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendationFeatureAreas, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*RecommendationFeatureAreas))
                }
            }
            m.SetFeatureAreas(res)
        }
        return nil
    }
    res["impactedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImpactedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImpactedResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ImpactedResourceable)
                }
            }
            m.SetImpactedResources(res)
        }
        return nil
    }
    res["impactStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactStartDateTime(val)
        }
        return nil
    }
    res["impactType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactType(val)
        }
        return nil
    }
    res["insights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsights(val)
        }
        return nil
    }
    res["lastCheckedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckedDateTime(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["maxScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecommendationBase_RecommendationBase_maxScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val.(*RecommendationBase_RecommendationBase_maxScore))
        }
        return nil
    }
    res["postponeUntilDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostponeUntilDateTime(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val.(*RecommendationPriority))
        }
        return nil
    }
    res["recommendationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendationType(val.(*RecommendationType))
        }
        return nil
    }
    res["releaseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseType(val)
        }
        return nil
    }
    res["remediationImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationImpact(val)
        }
        return nil
    }
    res["requiredLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredLicenses)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredLicenses(val.(*RequiredLicenses))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecommendationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*RecommendationStatus))
        }
        return nil
    }
    return res
}
// GetImpactedResources gets the impactedResources property value. The list of directory objects associated with the recommendation.
// returns a []ImpactedResourceable when successful
func (m *RecommendationBase) GetImpactedResources()([]ImpactedResourceable) {
    val, err := m.GetBackingStore().Get("impactedResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ImpactedResourceable)
    }
    return nil
}
// GetImpactStartDateTime gets the impactStartDateTime property value. The future date and time when a recommendation should be completed.
// returns a *Time when successful
func (m *RecommendationBase) GetImpactStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("impactStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetImpactType gets the impactType property value. Indicates the scope of impact of a recommendation. Tenant level indicates that the recommendation impacts the whole tenant. Other possible values include users, applications.
// returns a *string when successful
func (m *RecommendationBase) GetImpactType()(*string) {
    val, err := m.GetBackingStore().Get("impactType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInsights gets the insights property value. Describes why a recommendation uniquely applies to your directory. Corresponds to the Description section of a recommendation shown in the Microsoft Entra admin center.
// returns a *string when successful
func (m *RecommendationBase) GetInsights()(*string) {
    val, err := m.GetBackingStore().Get("insights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastCheckedDateTime gets the lastCheckedDateTime property value. The most recent date and time a recommendation was deemed applicable to your directory.
// returns a *Time when successful
func (m *RecommendationBase) GetLastCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCheckedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. Name of the user who last updated the status of the recommendation.
// returns a *string when successful
func (m *RecommendationBase) GetLastModifiedBy()(*string) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the status of a recommendation was last updated.
// returns a *Time when successful
func (m *RecommendationBase) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMaxScore gets the maxScore property value. The maximum number of points attainable. Only applies to recommendations with category set to identitySecureScore.
// returns a RecommendationBase_RecommendationBase_maxScoreable when successful
func (m *RecommendationBase) GetMaxScore()(RecommendationBase_RecommendationBase_maxScoreable) {
    val, err := m.GetBackingStore().Get("maxScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RecommendationBase_RecommendationBase_maxScoreable)
    }
    return nil
}
// GetPostponeUntilDateTime gets the postponeUntilDateTime property value. The future date and time when the status of a postponed recommendation will be active again.
// returns a *Time when successful
func (m *RecommendationBase) GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("postponeUntilDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPriority gets the priority property value. The priority property
// returns a *RecommendationPriority when successful
func (m *RecommendationBase) GetPriority()(*RecommendationPriority) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationPriority)
    }
    return nil
}
// GetRecommendationType gets the recommendationType property value. Friendly shortname to identify the recommendation. The possible values are: adfsAppsMigration, enableDesktopSSO, enablePHS, enableProvisioning, switchFromPerUserMFA, tenantMFA, thirdPartyApps, turnOffPerUserMFA, useAuthenticatorApp, useMyApps, staleApps, staleAppCreds, applicationCredentialExpiry, servicePrincipalKeyExpiry, adminMFAV2, blockLegacyAuthentication, integratedApps, mfaRegistrationV2, pwagePolicyNew, passwordHashSync, oneAdmin, roleOverlap, selfServicePasswordReset, signinRiskPolicy, userRiskPolicy, verifyAppPublisher, privateLinkForAAD, appRoleAssignmentsGroups, appRoleAssignmentsUsers, managedIdentity, overprivilegedApps, unknownFutureValue, longLivedCredentials, aadConnectDeprecated, adalToMsalMigration, ownerlessApps, inactiveGuests, aadGraphDeprecationApplication, aadGraphDeprecationServicePrincipal, mfaServerDeprecation. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: longLivedCredentials , aadConnectDeprecated , adalToMsalMigration , ownerlessApps , inactiveGuests , aadGraphDeprecationApplication , aadGraphDeprecationServicePrincipal , mfaServerDeprecation.
// returns a *RecommendationType when successful
func (m *RecommendationBase) GetRecommendationType()(*RecommendationType) {
    val, err := m.GetBackingStore().Get("recommendationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationType)
    }
    return nil
}
// GetReleaseType gets the releaseType property value. The current release type of the recommendation. The possible values are: preview, generallyAvailable, unknownFutureValue.
// returns a *string when successful
func (m *RecommendationBase) GetReleaseType()(*string) {
    val, err := m.GetBackingStore().Get("releaseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationImpact gets the remediationImpact property value. Description of the impact on users of the remediation. Only applies to recommendations with category set to identitySecureScore.
// returns a *string when successful
func (m *RecommendationBase) GetRemediationImpact()(*string) {
    val, err := m.GetBackingStore().Get("remediationImpact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequiredLicenses gets the requiredLicenses property value. The required licenses to view the recommendation. The possible values are: notApplicable, microsoftEntraIdFree, microsoftEntraIdP1, microsoftEntraIdP2, microsoftEntraIdGovernance, microsoftEntraWorkloadId, unknownFutureValue.
// returns a *RequiredLicenses when successful
func (m *RecommendationBase) GetRequiredLicenses()(*RequiredLicenses) {
    val, err := m.GetBackingStore().Get("requiredLicenses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RequiredLicenses)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *RecommendationStatus when successful
func (m *RecommendationBase) GetStatus()(*RecommendationStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RecommendationStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RecommendationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActionSteps()))
        for i, v := range m.GetActionSteps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("actionSteps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("benefits", m.GetBenefits())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("currentScore", m.GetCurrentScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetFeatureAreas() != nil {
        err = writer.WriteCollectionOfStringValues("featureAreas", SerializeRecommendationFeatureAreas(m.GetFeatureAreas()))
        if err != nil {
            return err
        }
    }
    if m.GetImpactedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImpactedResources()))
        for i, v := range m.GetImpactedResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("impactedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("impactStartDateTime", m.GetImpactStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("impactType", m.GetImpactType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("insights", m.GetInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckedDateTime", m.GetLastCheckedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("maxScore", m.GetMaxScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("postponeUntilDateTime", m.GetPostponeUntilDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        cast := (*m.GetPriority()).String()
        err = writer.WriteStringValue("priority", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecommendationType() != nil {
        cast := (*m.GetRecommendationType()).String()
        err = writer.WriteStringValue("recommendationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("releaseType", m.GetReleaseType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remediationImpact", m.GetRemediationImpact())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredLicenses() != nil {
        cast := (*m.GetRequiredLicenses()).String()
        err = writer.WriteStringValue("requiredLicenses", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionSteps sets the actionSteps property value. List of actions to take to complete a recommendation.
func (m *RecommendationBase) SetActionSteps(value []ActionStepable)() {
    err := m.GetBackingStore().Set("actionSteps", value)
    if err != nil {
        panic(err)
    }
}
// SetBenefits sets the benefits property value. An explanation of why completing the recommendation will benefit you. Corresponds to the Value section of a recommendation shown in the Microsoft Entra admin center.
func (m *RecommendationBase) SetBenefits(value *string)() {
    err := m.GetBackingStore().Set("benefits", value)
    if err != nil {
        panic(err)
    }
}
// SetCategory sets the category property value. The category property
func (m *RecommendationBase) SetCategory(value *RecommendationCategory)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the recommendation was detected as applicable to your directory.
func (m *RecommendationBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCurrentScore sets the currentScore property value. The number of points the tenant has attained. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) SetCurrentScore(value RecommendationBase_RecommendationBase_currentScoreable)() {
    err := m.GetBackingStore().Set("currentScore", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The title of the recommendation.
func (m *RecommendationBase) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFeatureAreas sets the featureAreas property value. The directory feature that the recommendation is related to.
func (m *RecommendationBase) SetFeatureAreas(value []RecommendationFeatureAreas)() {
    err := m.GetBackingStore().Set("featureAreas", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactedResources sets the impactedResources property value. The list of directory objects associated with the recommendation.
func (m *RecommendationBase) SetImpactedResources(value []ImpactedResourceable)() {
    err := m.GetBackingStore().Set("impactedResources", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactStartDateTime sets the impactStartDateTime property value. The future date and time when a recommendation should be completed.
func (m *RecommendationBase) SetImpactStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("impactStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactType sets the impactType property value. Indicates the scope of impact of a recommendation. Tenant level indicates that the recommendation impacts the whole tenant. Other possible values include users, applications.
func (m *RecommendationBase) SetImpactType(value *string)() {
    err := m.GetBackingStore().Set("impactType", value)
    if err != nil {
        panic(err)
    }
}
// SetInsights sets the insights property value. Describes why a recommendation uniquely applies to your directory. Corresponds to the Description section of a recommendation shown in the Microsoft Entra admin center.
func (m *RecommendationBase) SetInsights(value *string)() {
    err := m.GetBackingStore().Set("insights", value)
    if err != nil {
        panic(err)
    }
}
// SetLastCheckedDateTime sets the lastCheckedDateTime property value. The most recent date and time a recommendation was deemed applicable to your directory.
func (m *RecommendationBase) SetLastCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCheckedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Name of the user who last updated the status of the recommendation.
func (m *RecommendationBase) SetLastModifiedBy(value *string)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the status of a recommendation was last updated.
func (m *RecommendationBase) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxScore sets the maxScore property value. The maximum number of points attainable. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) SetMaxScore(value RecommendationBase_RecommendationBase_maxScoreable)() {
    err := m.GetBackingStore().Set("maxScore", value)
    if err != nil {
        panic(err)
    }
}
// SetPostponeUntilDateTime sets the postponeUntilDateTime property value. The future date and time when the status of a postponed recommendation will be active again.
func (m *RecommendationBase) SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("postponeUntilDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority property
func (m *RecommendationBase) SetPriority(value *RecommendationPriority)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetRecommendationType sets the recommendationType property value. Friendly shortname to identify the recommendation. The possible values are: adfsAppsMigration, enableDesktopSSO, enablePHS, enableProvisioning, switchFromPerUserMFA, tenantMFA, thirdPartyApps, turnOffPerUserMFA, useAuthenticatorApp, useMyApps, staleApps, staleAppCreds, applicationCredentialExpiry, servicePrincipalKeyExpiry, adminMFAV2, blockLegacyAuthentication, integratedApps, mfaRegistrationV2, pwagePolicyNew, passwordHashSync, oneAdmin, roleOverlap, selfServicePasswordReset, signinRiskPolicy, userRiskPolicy, verifyAppPublisher, privateLinkForAAD, appRoleAssignmentsGroups, appRoleAssignmentsUsers, managedIdentity, overprivilegedApps, unknownFutureValue, longLivedCredentials, aadConnectDeprecated, adalToMsalMigration, ownerlessApps, inactiveGuests, aadGraphDeprecationApplication, aadGraphDeprecationServicePrincipal, mfaServerDeprecation. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: longLivedCredentials , aadConnectDeprecated , adalToMsalMigration , ownerlessApps , inactiveGuests , aadGraphDeprecationApplication , aadGraphDeprecationServicePrincipal , mfaServerDeprecation.
func (m *RecommendationBase) SetRecommendationType(value *RecommendationType)() {
    err := m.GetBackingStore().Set("recommendationType", value)
    if err != nil {
        panic(err)
    }
}
// SetReleaseType sets the releaseType property value. The current release type of the recommendation. The possible values are: preview, generallyAvailable, unknownFutureValue.
func (m *RecommendationBase) SetReleaseType(value *string)() {
    err := m.GetBackingStore().Set("releaseType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationImpact sets the remediationImpact property value. Description of the impact on users of the remediation. Only applies to recommendations with category set to identitySecureScore.
func (m *RecommendationBase) SetRemediationImpact(value *string)() {
    err := m.GetBackingStore().Set("remediationImpact", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredLicenses sets the requiredLicenses property value. The required licenses to view the recommendation. The possible values are: notApplicable, microsoftEntraIdFree, microsoftEntraIdP1, microsoftEntraIdP2, microsoftEntraIdGovernance, microsoftEntraWorkloadId, unknownFutureValue.
func (m *RecommendationBase) SetRequiredLicenses(value *RequiredLicenses)() {
    err := m.GetBackingStore().Set("requiredLicenses", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *RecommendationBase) SetStatus(value *RecommendationStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
type RecommendationBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionSteps()([]ActionStepable)
    GetBenefits()(*string)
    GetCategory()(*RecommendationCategory)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentScore()(RecommendationBase_RecommendationBase_currentScoreable)
    GetDisplayName()(*string)
    GetFeatureAreas()([]RecommendationFeatureAreas)
    GetImpactedResources()([]ImpactedResourceable)
    GetImpactStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetImpactType()(*string)
    GetInsights()(*string)
    GetLastCheckedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedBy()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaxScore()(RecommendationBase_RecommendationBase_maxScoreable)
    GetPostponeUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPriority()(*RecommendationPriority)
    GetRecommendationType()(*RecommendationType)
    GetReleaseType()(*string)
    GetRemediationImpact()(*string)
    GetRequiredLicenses()(*RequiredLicenses)
    GetStatus()(*RecommendationStatus)
    SetActionSteps(value []ActionStepable)()
    SetBenefits(value *string)()
    SetCategory(value *RecommendationCategory)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentScore(value RecommendationBase_RecommendationBase_currentScoreable)()
    SetDisplayName(value *string)()
    SetFeatureAreas(value []RecommendationFeatureAreas)()
    SetImpactedResources(value []ImpactedResourceable)()
    SetImpactStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetImpactType(value *string)()
    SetInsights(value *string)()
    SetLastCheckedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedBy(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaxScore(value RecommendationBase_RecommendationBase_maxScoreable)()
    SetPostponeUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPriority(value *RecommendationPriority)()
    SetRecommendationType(value *RecommendationType)()
    SetReleaseType(value *string)()
    SetRemediationImpact(value *string)()
    SetRequiredLicenses(value *RequiredLicenses)()
    SetStatus(value *RecommendationStatus)()
}
