package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type QualityUpdateCveSeverityInformation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore composed type wrapper for classes float64, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric, string
type QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewQualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore instantiates a new QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore and sets the default values.
func NewQualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore()(*QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) {
    m := &QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateQualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateQualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewQualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore()
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
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) GetDouble()(*float64) {
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
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric) {
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
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) GetString()(*string) {
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
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)()
    SetString(value *string)()
}
// NewQualityUpdateCveSeverityInformation instantiates a new QualityUpdateCveSeverityInformation and sets the default values.
func NewQualityUpdateCveSeverityInformation()(*QualityUpdateCveSeverityInformation) {
    m := &QualityUpdateCveSeverityInformation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateQualityUpdateCveSeverityInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateQualityUpdateCveSeverityInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQualityUpdateCveSeverityInformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *QualityUpdateCveSeverityInformation) GetAdditionalData()(map[string]any) {
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
func (m *QualityUpdateCveSeverityInformation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExploitedCves gets the exploitedCves property value. The exploitedCves property
// returns a []CveInformationable when successful
func (m *QualityUpdateCveSeverityInformation) GetExploitedCves()([]CveInformationable) {
    val, err := m.GetBackingStore().Get("exploitedCves")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CveInformationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *QualityUpdateCveSeverityInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exploitedCves"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCveInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CveInformationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CveInformationable)
                }
            }
            m.SetExploitedCves(res)
        }
        return nil
    }
    res["maxBaseScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateQualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxBaseScore(val.(*QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScore))
        }
        return nil
    }
    res["maxSeverity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCveSeverityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSeverity(val.(*CveSeverityLevel))
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
// GetMaxBaseScore gets the maxBaseScore property value. Highest base score that occurs of any CVE addressed by the quality update. Read-only.
// returns a QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreable when successful
func (m *QualityUpdateCveSeverityInformation) GetMaxBaseScore()(QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreable) {
    val, err := m.GetBackingStore().Get("maxBaseScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreable)
    }
    return nil
}
// GetMaxSeverity gets the maxSeverity property value. The maxSeverity property
// returns a *CveSeverityLevel when successful
func (m *QualityUpdateCveSeverityInformation) GetMaxSeverity()(*CveSeverityLevel) {
    val, err := m.GetBackingStore().Get("maxSeverity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CveSeverityLevel)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *QualityUpdateCveSeverityInformation) GetOdataType()(*string) {
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
func (m *QualityUpdateCveSeverityInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExploitedCves() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExploitedCves()))
        for i, v := range m.GetExploitedCves() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("exploitedCves", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("maxBaseScore", m.GetMaxBaseScore())
        if err != nil {
            return err
        }
    }
    if m.GetMaxSeverity() != nil {
        cast := (*m.GetMaxSeverity()).String()
        err := writer.WriteStringValue("maxSeverity", &cast)
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
func (m *QualityUpdateCveSeverityInformation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *QualityUpdateCveSeverityInformation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExploitedCves sets the exploitedCves property value. The exploitedCves property
func (m *QualityUpdateCveSeverityInformation) SetExploitedCves(value []CveInformationable)() {
    err := m.GetBackingStore().Set("exploitedCves", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxBaseScore sets the maxBaseScore property value. Highest base score that occurs of any CVE addressed by the quality update. Read-only.
func (m *QualityUpdateCveSeverityInformation) SetMaxBaseScore(value QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreable)() {
    err := m.GetBackingStore().Set("maxBaseScore", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxSeverity sets the maxSeverity property value. The maxSeverity property
func (m *QualityUpdateCveSeverityInformation) SetMaxSeverity(value *CveSeverityLevel)() {
    err := m.GetBackingStore().Set("maxSeverity", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *QualityUpdateCveSeverityInformation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type QualityUpdateCveSeverityInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExploitedCves()([]CveInformationable)
    GetMaxBaseScore()(QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreable)
    GetMaxSeverity()(*CveSeverityLevel)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExploitedCves(value []CveInformationable)()
    SetMaxBaseScore(value QualityUpdateCveSeverityInformation_QualityUpdateCveSeverityInformation_maxBaseScoreable)()
    SetMaxSeverity(value *CveSeverityLevel)()
    SetOdataType(value *string)()
}
