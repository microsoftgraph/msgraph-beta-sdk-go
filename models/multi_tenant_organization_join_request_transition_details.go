package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MultiTenantOrganizationJoinRequestTransitionDetails 
type MultiTenantOrganizationJoinRequestTransitionDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMultiTenantOrganizationJoinRequestTransitionDetails instantiates a new multiTenantOrganizationJoinRequestTransitionDetails and sets the default values.
func NewMultiTenantOrganizationJoinRequestTransitionDetails()(*MultiTenantOrganizationJoinRequestTransitionDetails) {
    m := &MultiTenantOrganizationJoinRequestTransitionDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMultiTenantOrganizationJoinRequestTransitionDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMultiTenantOrganizationJoinRequestTransitionDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMultiTenantOrganizationJoinRequestTransitionDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetAdditionalData()(map[string]any) {
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
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDesiredMemberState gets the desiredMemberState property value. State of the tenant in the multitenant organization currently being processed. The possible values are: pending, active, removed, unknownFutureValue. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetDesiredMemberState()(*MultiTenantOrganizationJoinRequestTransitionDetails_desiredMemberState) {
    val, err := m.GetBackingStore().Get("desiredMemberState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MultiTenantOrganizationJoinRequestTransitionDetails_desiredMemberState)
    }
    return nil
}
// GetDetails gets the details property value. Details that explain the processing status if any. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetDetails()(*string) {
    val, err := m.GetBackingStore().Get("details")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["desiredMemberState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationJoinRequestTransitionDetails_desiredMemberState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDesiredMemberState(val.(*MultiTenantOrganizationJoinRequestTransitionDetails_desiredMemberState))
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMultiTenantOrganizationJoinRequestTransitionDetails_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*MultiTenantOrganizationJoinRequestTransitionDetails_status))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) GetStatus()(*MultiTenantOrganizationJoinRequestTransitionDetails_status) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MultiTenantOrganizationJoinRequestTransitionDetails_status)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDesiredMemberState() != nil {
        cast := (*m.GetDesiredMemberState()).String()
        err := writer.WriteStringValue("desiredMemberState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("details", m.GetDetails())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDesiredMemberState sets the desiredMemberState property value. State of the tenant in the multitenant organization currently being processed. The possible values are: pending, active, removed, unknownFutureValue. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetDesiredMemberState(value *MultiTenantOrganizationJoinRequestTransitionDetails_desiredMemberState)() {
    err := m.GetBackingStore().Set("desiredMemberState", value)
    if err != nil {
        panic(err)
    }
}
// SetDetails sets the details property value. Details that explain the processing status if any. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetDetails(value *string)() {
    err := m.GetBackingStore().Set("details", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Processing state of the asynchronous job. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue. Read-only.
func (m *MultiTenantOrganizationJoinRequestTransitionDetails) SetStatus(value *MultiTenantOrganizationJoinRequestTransitionDetails_status)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// MultiTenantOrganizationJoinRequestTransitionDetailsable 
type MultiTenantOrganizationJoinRequestTransitionDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDesiredMemberState()(*MultiTenantOrganizationJoinRequestTransitionDetails_desiredMemberState)
    GetDetails()(*string)
    GetOdataType()(*string)
    GetStatus()(*MultiTenantOrganizationJoinRequestTransitionDetails_status)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDesiredMemberState(value *MultiTenantOrganizationJoinRequestTransitionDetails_desiredMemberState)()
    SetDetails(value *string)()
    SetOdataType(value *string)()
    SetStatus(value *MultiTenantOrganizationJoinRequestTransitionDetails_status)()
}
