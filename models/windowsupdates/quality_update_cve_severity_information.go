package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// QualityUpdateCveSeverityInformation 
type QualityUpdateCveSeverityInformation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewQualityUpdateCveSeverityInformation instantiates a new qualityUpdateCveSeverityInformation and sets the default values.
func NewQualityUpdateCveSeverityInformation()(*QualityUpdateCveSeverityInformation) {
    m := &QualityUpdateCveSeverityInformation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateQualityUpdateCveSeverityInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQualityUpdateCveSeverityInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQualityUpdateCveSeverityInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *QualityUpdateCveSeverityInformation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExploitedCves gets the exploitedCves property value. The exploitedCves property
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
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxBaseScore(val)
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
// GetMaxBaseScore gets the maxBaseScore property value. The maxBaseScore property
func (m *QualityUpdateCveSeverityInformation) GetMaxBaseScore()(*float64) {
    val, err := m.GetBackingStore().Get("maxBaseScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetMaxSeverity gets the maxSeverity property value. The maxSeverity property
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
        err := writer.WriteFloat64Value("maxBaseScore", m.GetMaxBaseScore())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *QualityUpdateCveSeverityInformation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
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
// SetMaxBaseScore sets the maxBaseScore property value. The maxBaseScore property
func (m *QualityUpdateCveSeverityInformation) SetMaxBaseScore(value *float64)() {
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
// QualityUpdateCveSeverityInformationable 
type QualityUpdateCveSeverityInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExploitedCves()([]CveInformationable)
    GetMaxBaseScore()(*float64)
    GetMaxSeverity()(*CveSeverityLevel)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExploitedCves(value []CveInformationable)()
    SetMaxBaseScore(value *float64)()
    SetMaxSeverity(value *CveSeverityLevel)()
    SetOdataType(value *string)()
}
