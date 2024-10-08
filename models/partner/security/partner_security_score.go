package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type PartnerSecurityScore struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// PartnerSecurityScore_PartnerSecurityScore_currentScore composed type wrapper for classes float32, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric, string
type PartnerSecurityScore_PartnerSecurityScore_currentScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPartnerSecurityScore_PartnerSecurityScore_currentScore instantiates a new PartnerSecurityScore_PartnerSecurityScore_currentScore and sets the default values.
func NewPartnerSecurityScore_PartnerSecurityScore_currentScore()(*PartnerSecurityScore_PartnerSecurityScore_currentScore) {
    m := &PartnerSecurityScore_PartnerSecurityScore_currentScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreatePartnerSecurityScore_PartnerSecurityScore_currentScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePartnerSecurityScore_PartnerSecurityScore_currentScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewPartnerSecurityScore_PartnerSecurityScore_currentScore()
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
    if val, err := parseNode.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
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
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) GetString()(*string) {
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
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
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
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *PartnerSecurityScore_PartnerSecurityScore_currentScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// PartnerSecurityScore_PartnerSecurityScore_maxScore composed type wrapper for classes float32, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric, string
type PartnerSecurityScore_PartnerSecurityScore_maxScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPartnerSecurityScore_PartnerSecurityScore_maxScore instantiates a new PartnerSecurityScore_PartnerSecurityScore_maxScore and sets the default values.
func NewPartnerSecurityScore_PartnerSecurityScore_maxScore()(*PartnerSecurityScore_PartnerSecurityScore_maxScore) {
    m := &PartnerSecurityScore_PartnerSecurityScore_maxScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreatePartnerSecurityScore_PartnerSecurityScore_maxScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePartnerSecurityScore_PartnerSecurityScore_maxScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewPartnerSecurityScore_PartnerSecurityScore_maxScore()
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
    if val, err := parseNode.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
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
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) GetString()(*string) {
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
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
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
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *PartnerSecurityScore_PartnerSecurityScore_maxScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type PartnerSecurityScore_PartnerSecurityScore_currentScoreable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)()
    SetString(value *string)()
}
type PartnerSecurityScore_PartnerSecurityScore_maxScoreable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)()
    SetString(value *string)()
}
// NewPartnerSecurityScore instantiates a new PartnerSecurityScore and sets the default values.
func NewPartnerSecurityScore()(*PartnerSecurityScore) {
    m := &PartnerSecurityScore{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreatePartnerSecurityScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePartnerSecurityScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPartnerSecurityScore(), nil
}
// GetCurrentScore gets the currentScore property value. The current security score for the partner.
// returns a PartnerSecurityScore_PartnerSecurityScore_currentScoreable when successful
func (m *PartnerSecurityScore) GetCurrentScore()(PartnerSecurityScore_PartnerSecurityScore_currentScoreable) {
    val, err := m.GetBackingStore().Get("currentScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PartnerSecurityScore_PartnerSecurityScore_currentScoreable)
    }
    return nil
}
// GetCustomerInsights gets the customerInsights property value. Contains customer-specific information for certain requirements.
// returns a []CustomerInsightable when successful
func (m *PartnerSecurityScore) GetCustomerInsights()([]CustomerInsightable) {
    val, err := m.GetBackingStore().Get("customerInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomerInsightable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PartnerSecurityScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["currentScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePartnerSecurityScore_PartnerSecurityScore_currentScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentScore(val.(*PartnerSecurityScore_PartnerSecurityScore_currentScore))
        }
        return nil
    }
    res["customerInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomerInsightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomerInsightable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomerInsightable)
                }
            }
            m.SetCustomerInsights(res)
        }
        return nil
    }
    res["history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityScoreHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityScoreHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecurityScoreHistoryable)
                }
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["lastRefreshDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshDateTime(val)
        }
        return nil
    }
    res["maxScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePartnerSecurityScore_PartnerSecurityScore_maxScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val.(*PartnerSecurityScore_PartnerSecurityScore_maxScore))
        }
        return nil
    }
    res["requirements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityRequirementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityRequirementable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecurityRequirementable)
                }
            }
            m.SetRequirements(res)
        }
        return nil
    }
    res["updatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedDateTime(val)
        }
        return nil
    }
    return res
}
// GetHistory gets the history property value. Contains a list of recent score changes.
// returns a []SecurityScoreHistoryable when successful
func (m *PartnerSecurityScore) GetHistory()([]SecurityScoreHistoryable) {
    val, err := m.GetBackingStore().Get("history")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityScoreHistoryable)
    }
    return nil
}
// GetLastRefreshDateTime gets the lastRefreshDateTime property value. The last time the data was checked.
// returns a *Time when successful
func (m *PartnerSecurityScore) GetLastRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastRefreshDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMaxScore gets the maxScore property value. The maximum score possible.
// returns a PartnerSecurityScore_PartnerSecurityScore_maxScoreable when successful
func (m *PartnerSecurityScore) GetMaxScore()(PartnerSecurityScore_PartnerSecurityScore_maxScoreable) {
    val, err := m.GetBackingStore().Get("maxScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PartnerSecurityScore_PartnerSecurityScore_maxScoreable)
    }
    return nil
}
// GetRequirements gets the requirements property value. Contains the list of security requirements that make up the score.
// returns a []SecurityRequirementable when successful
func (m *PartnerSecurityScore) GetRequirements()([]SecurityRequirementable) {
    val, err := m.GetBackingStore().Get("requirements")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityRequirementable)
    }
    return nil
}
// GetUpdatedDateTime gets the updatedDateTime property value. The last time the security score or related properties changed.
// returns a *Time when successful
func (m *PartnerSecurityScore) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("updatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PartnerSecurityScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("currentScore", m.GetCurrentScore())
        if err != nil {
            return err
        }
    }
    if m.GetCustomerInsights() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomerInsights()))
        for i, v := range m.GetCustomerInsights() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customerInsights", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshDateTime", m.GetLastRefreshDateTime())
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
    if m.GetRequirements() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequirements()))
        for i, v := range m.GetRequirements() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("requirements", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCurrentScore sets the currentScore property value. The current security score for the partner.
func (m *PartnerSecurityScore) SetCurrentScore(value PartnerSecurityScore_PartnerSecurityScore_currentScoreable)() {
    err := m.GetBackingStore().Set("currentScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomerInsights sets the customerInsights property value. Contains customer-specific information for certain requirements.
func (m *PartnerSecurityScore) SetCustomerInsights(value []CustomerInsightable)() {
    err := m.GetBackingStore().Set("customerInsights", value)
    if err != nil {
        panic(err)
    }
}
// SetHistory sets the history property value. Contains a list of recent score changes.
func (m *PartnerSecurityScore) SetHistory(value []SecurityScoreHistoryable)() {
    err := m.GetBackingStore().Set("history", value)
    if err != nil {
        panic(err)
    }
}
// SetLastRefreshDateTime sets the lastRefreshDateTime property value. The last time the data was checked.
func (m *PartnerSecurityScore) SetLastRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastRefreshDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxScore sets the maxScore property value. The maximum score possible.
func (m *PartnerSecurityScore) SetMaxScore(value PartnerSecurityScore_PartnerSecurityScore_maxScoreable)() {
    err := m.GetBackingStore().Set("maxScore", value)
    if err != nil {
        panic(err)
    }
}
// SetRequirements sets the requirements property value. Contains the list of security requirements that make up the score.
func (m *PartnerSecurityScore) SetRequirements(value []SecurityRequirementable)() {
    err := m.GetBackingStore().Set("requirements", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdatedDateTime sets the updatedDateTime property value. The last time the security score or related properties changed.
func (m *PartnerSecurityScore) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("updatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
type PartnerSecurityScoreable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentScore()(PartnerSecurityScore_PartnerSecurityScore_currentScoreable)
    GetCustomerInsights()([]CustomerInsightable)
    GetHistory()([]SecurityScoreHistoryable)
    GetLastRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaxScore()(PartnerSecurityScore_PartnerSecurityScore_maxScoreable)
    GetRequirements()([]SecurityRequirementable)
    GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCurrentScore(value PartnerSecurityScore_PartnerSecurityScore_currentScoreable)()
    SetCustomerInsights(value []CustomerInsightable)()
    SetHistory(value []SecurityScoreHistoryable)()
    SetLastRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaxScore(value PartnerSecurityScore_PartnerSecurityScore_maxScoreable)()
    SetRequirements(value []SecurityRequirementable)()
    SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
