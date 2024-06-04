package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EnrollmentTimeDeviceMembershipTargetResult the EnrollmentTimeDeviceMembershipTargetResult entity represents the results of the set/get EnrollmentTimeDeviceMembershipTarget request. The set/get EnrollmentTimeDeviceMembershipTarget API validates the device membership targets specified by the admin to ensure that they exist, that they are of the proper type, and any other target requirements are met such as that the Intune Device Provisioning First Party App is an owner of the target. Failures other than validation will result in 500 else validationSucceeded will be true or false if any of the validation fails for EnrollmentTimeDeviceMembershipTarget.
type EnrollmentTimeDeviceMembershipTargetResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEnrollmentTimeDeviceMembershipTargetResult instantiates a new EnrollmentTimeDeviceMembershipTargetResult and sets the default values.
func NewEnrollmentTimeDeviceMembershipTargetResult()(*EnrollmentTimeDeviceMembershipTargetResult) {
    m := &EnrollmentTimeDeviceMembershipTargetResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnrollmentTimeDeviceMembershipTargetResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnrollmentTimeDeviceMembershipTargetResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentTimeDeviceMembershipTargetResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnrollmentTimeDeviceMembershipTargetResult) GetAdditionalData()(map[string]any) {
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
func (m *EnrollmentTimeDeviceMembershipTargetResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEnrollmentTimeDeviceMembershipTargetValidationStatuses gets the enrollmentTimeDeviceMembershipTargetValidationStatuses property value. A list of validation status of the memberships targetted to profile. This collection can contain a maximum of 1 elements.
// returns a []EnrollmentTimeDeviceMembershipTargetStatusable when successful
func (m *EnrollmentTimeDeviceMembershipTargetResult) GetEnrollmentTimeDeviceMembershipTargetValidationStatuses()([]EnrollmentTimeDeviceMembershipTargetStatusable) {
    val, err := m.GetBackingStore().Get("enrollmentTimeDeviceMembershipTargetValidationStatuses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EnrollmentTimeDeviceMembershipTargetStatusable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnrollmentTimeDeviceMembershipTargetResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentTimeDeviceMembershipTargetValidationStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnrollmentTimeDeviceMembershipTargetStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnrollmentTimeDeviceMembershipTargetStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EnrollmentTimeDeviceMembershipTargetStatusable)
                }
            }
            m.SetEnrollmentTimeDeviceMembershipTargetValidationStatuses(res)
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
    res["validationSucceeded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidationSucceeded(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *EnrollmentTimeDeviceMembershipTargetResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValidationSucceeded gets the validationSucceeded property value. Indicates if validations succeeded for the device membership target. When 'true', the device membership target validation found no issues. When 'false', the device membership target validation found issues. default - false
// returns a *bool when successful
func (m *EnrollmentTimeDeviceMembershipTargetResult) GetValidationSucceeded()(*bool) {
    val, err := m.GetBackingStore().Get("validationSucceeded")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EnrollmentTimeDeviceMembershipTargetResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnrollmentTimeDeviceMembershipTargetValidationStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnrollmentTimeDeviceMembershipTargetValidationStatuses()))
        for i, v := range m.GetEnrollmentTimeDeviceMembershipTargetValidationStatuses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("enrollmentTimeDeviceMembershipTargetValidationStatuses", cast)
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
        err := writer.WriteBoolValue("validationSucceeded", m.GetValidationSucceeded())
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
func (m *EnrollmentTimeDeviceMembershipTargetResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EnrollmentTimeDeviceMembershipTargetResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEnrollmentTimeDeviceMembershipTargetValidationStatuses sets the enrollmentTimeDeviceMembershipTargetValidationStatuses property value. A list of validation status of the memberships targetted to profile. This collection can contain a maximum of 1 elements.
func (m *EnrollmentTimeDeviceMembershipTargetResult) SetEnrollmentTimeDeviceMembershipTargetValidationStatuses(value []EnrollmentTimeDeviceMembershipTargetStatusable)() {
    err := m.GetBackingStore().Set("enrollmentTimeDeviceMembershipTargetValidationStatuses", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EnrollmentTimeDeviceMembershipTargetResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetValidationSucceeded sets the validationSucceeded property value. Indicates if validations succeeded for the device membership target. When 'true', the device membership target validation found no issues. When 'false', the device membership target validation found issues. default - false
func (m *EnrollmentTimeDeviceMembershipTargetResult) SetValidationSucceeded(value *bool)() {
    err := m.GetBackingStore().Set("validationSucceeded", value)
    if err != nil {
        panic(err)
    }
}
type EnrollmentTimeDeviceMembershipTargetResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEnrollmentTimeDeviceMembershipTargetValidationStatuses()([]EnrollmentTimeDeviceMembershipTargetStatusable)
    GetOdataType()(*string)
    GetValidationSucceeded()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEnrollmentTimeDeviceMembershipTargetValidationStatuses(value []EnrollmentTimeDeviceMembershipTargetStatusable)()
    SetOdataType(value *string)()
    SetValidationSucceeded(value *bool)()
}
