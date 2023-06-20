package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// SystemCredentialPreferences 
type SystemCredentialPreferences struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSystemCredentialPreferences instantiates a new systemCredentialPreferences and sets the default values.
func NewSystemCredentialPreferences()(*SystemCredentialPreferences) {
    m := &SystemCredentialPreferences{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSystemCredentialPreferencesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSystemCredentialPreferencesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSystemCredentialPreferences(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SystemCredentialPreferences) GetAdditionalData()(map[string]any) {
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
func (m *SystemCredentialPreferences) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExcludeTargets gets the excludeTargets property value. Users and groups excluded from the preferred authentication method experience of the system.
func (m *SystemCredentialPreferences) GetExcludeTargets()([]ExcludeTargetable) {
    val, err := m.GetBackingStore().Get("excludeTargets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExcludeTargetable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SystemCredentialPreferences) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExcludeTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExcludeTargetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExcludeTargetable)
                }
            }
            m.SetExcludeTargets(res)
        }
        return nil
    }
    res["includeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIncludeTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IncludeTargetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IncludeTargetable)
                }
            }
            m.SetIncludeTargets(res)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedConfigState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AdvancedConfigState))
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. Users and groups included in the preferred authentication method experience of the system.
func (m *SystemCredentialPreferences) GetIncludeTargets()([]IncludeTargetable) {
    val, err := m.GetBackingStore().Get("includeTargets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IncludeTargetable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SystemCredentialPreferences) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetState gets the state property value. The state property
func (m *SystemCredentialPreferences) GetState()(*AdvancedConfigState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AdvancedConfigState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SystemCredentialPreferences) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExcludeTargets()))
        for i, v := range m.GetExcludeTargets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("excludeTargets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("includeTargets", cast)
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *SystemCredentialPreferences) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *SystemCredentialPreferences) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExcludeTargets sets the excludeTargets property value. Users and groups excluded from the preferred authentication method experience of the system.
func (m *SystemCredentialPreferences) SetExcludeTargets(value []ExcludeTargetable)() {
    err := m.GetBackingStore().Set("excludeTargets", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeTargets sets the includeTargets property value. Users and groups included in the preferred authentication method experience of the system.
func (m *SystemCredentialPreferences) SetIncludeTargets(value []IncludeTargetable)() {
    err := m.GetBackingStore().Set("includeTargets", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SystemCredentialPreferences) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *SystemCredentialPreferences) SetState(value *AdvancedConfigState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SystemCredentialPreferencesable 
type SystemCredentialPreferencesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExcludeTargets()([]ExcludeTargetable)
    GetIncludeTargets()([]IncludeTargetable)
    GetOdataType()(*string)
    GetState()(*AdvancedConfigState)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExcludeTargets(value []ExcludeTargetable)()
    SetIncludeTargets(value []IncludeTargetable)()
    SetOdataType(value *string)()
    SetState(value *AdvancedConfigState)()
}
