package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type UserTrainingCompletionSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserTrainingCompletionSummary instantiates a new UserTrainingCompletionSummary and sets the default values.
func NewUserTrainingCompletionSummary()(*UserTrainingCompletionSummary) {
    m := &UserTrainingCompletionSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserTrainingCompletionSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserTrainingCompletionSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserTrainingCompletionSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserTrainingCompletionSummary) GetAdditionalData()(map[string]any) {
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
func (m *UserTrainingCompletionSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompletedUsersCount gets the completedUsersCount property value. The completedUsersCount property
// returns a *int32 when successful
func (m *UserTrainingCompletionSummary) GetCompletedUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("completedUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserTrainingCompletionSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completedUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedUsersCount(val)
        }
        return nil
    }
    res["inProgressUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInProgressUsersCount(val)
        }
        return nil
    }
    res["notCompletedUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotCompletedUsersCount(val)
        }
        return nil
    }
    res["notStartedUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotStartedUsersCount(val)
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
    res["previouslyAssignedUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviouslyAssignedUsersCount(val)
        }
        return nil
    }
    return res
}
// GetInProgressUsersCount gets the inProgressUsersCount property value. The inProgressUsersCount property
// returns a *int32 when successful
func (m *UserTrainingCompletionSummary) GetInProgressUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("inProgressUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotCompletedUsersCount gets the notCompletedUsersCount property value. The notCompletedUsersCount property
// returns a *int32 when successful
func (m *UserTrainingCompletionSummary) GetNotCompletedUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("notCompletedUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotStartedUsersCount gets the notStartedUsersCount property value. The notStartedUsersCount property
// returns a *int32 when successful
func (m *UserTrainingCompletionSummary) GetNotStartedUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("notStartedUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *UserTrainingCompletionSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreviouslyAssignedUsersCount gets the previouslyAssignedUsersCount property value. The previouslyAssignedUsersCount property
// returns a *int32 when successful
func (m *UserTrainingCompletionSummary) GetPreviouslyAssignedUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("previouslyAssignedUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserTrainingCompletionSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("completedUsersCount", m.GetCompletedUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inProgressUsersCount", m.GetInProgressUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("notCompletedUsersCount", m.GetNotCompletedUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("notStartedUsersCount", m.GetNotStartedUsersCount())
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
        err := writer.WriteInt32Value("previouslyAssignedUsersCount", m.GetPreviouslyAssignedUsersCount())
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
func (m *UserTrainingCompletionSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *UserTrainingCompletionSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompletedUsersCount sets the completedUsersCount property value. The completedUsersCount property
func (m *UserTrainingCompletionSummary) SetCompletedUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("completedUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInProgressUsersCount sets the inProgressUsersCount property value. The inProgressUsersCount property
func (m *UserTrainingCompletionSummary) SetInProgressUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("inProgressUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotCompletedUsersCount sets the notCompletedUsersCount property value. The notCompletedUsersCount property
func (m *UserTrainingCompletionSummary) SetNotCompletedUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("notCompletedUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotStartedUsersCount sets the notStartedUsersCount property value. The notStartedUsersCount property
func (m *UserTrainingCompletionSummary) SetNotStartedUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("notStartedUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserTrainingCompletionSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviouslyAssignedUsersCount sets the previouslyAssignedUsersCount property value. The previouslyAssignedUsersCount property
func (m *UserTrainingCompletionSummary) SetPreviouslyAssignedUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("previouslyAssignedUsersCount", value)
    if err != nil {
        panic(err)
    }
}
type UserTrainingCompletionSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompletedUsersCount()(*int32)
    GetInProgressUsersCount()(*int32)
    GetNotCompletedUsersCount()(*int32)
    GetNotStartedUsersCount()(*int32)
    GetOdataType()(*string)
    GetPreviouslyAssignedUsersCount()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompletedUsersCount(value *int32)()
    SetInProgressUsersCount(value *int32)()
    SetNotCompletedUsersCount(value *int32)()
    SetNotStartedUsersCount(value *int32)()
    SetOdataType(value *string)()
    SetPreviouslyAssignedUsersCount(value *int32)()
}
