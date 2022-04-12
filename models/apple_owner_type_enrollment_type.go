package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleOwnerTypeEnrollmentType 
type AppleOwnerTypeEnrollmentType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enrollment type. Possible values are: unknown, device, user.
    enrollmentType *AppleUserInitiatedEnrollmentType
    // The owner type. Possible values are: unknown, company, personal.
    ownerType *ManagedDeviceOwnerType
}
// NewAppleOwnerTypeEnrollmentType instantiates a new appleOwnerTypeEnrollmentType and sets the default values.
func NewAppleOwnerTypeEnrollmentType()(*AppleOwnerTypeEnrollmentType) {
    m := &AppleOwnerTypeEnrollmentType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleOwnerTypeEnrollmentType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppleOwnerTypeEnrollmentType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnrollmentType gets the enrollmentType property value. The enrollment type. Possible values are: unknown, device, user.
func (m *AppleOwnerTypeEnrollmentType) GetEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleOwnerTypeEnrollmentType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleUserInitiatedEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val.(*AppleUserInitiatedEnrollmentType))
        }
        return nil
    }
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    return res
}
// GetOwnerType gets the ownerType property value. The owner type. Possible values are: unknown, company, personal.
func (m *AppleOwnerTypeEnrollmentType) GetOwnerType()(*ManagedDeviceOwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
// Serialize serializes information the current object
func (m *AppleOwnerTypeEnrollmentType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnrollmentType() != nil {
        cast := (*m.GetEnrollmentType()).String()
        err := writer.WriteStringValue("enrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err := writer.WriteStringValue("ownerType", &cast)
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
func (m *AppleOwnerTypeEnrollmentType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnrollmentType sets the enrollmentType property value. The enrollment type. Possible values are: unknown, device, user.
func (m *AppleOwnerTypeEnrollmentType) SetEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    if m != nil {
        m.enrollmentType = value
    }
}
// SetOwnerType sets the ownerType property value. The owner type. Possible values are: unknown, company, personal.
func (m *AppleOwnerTypeEnrollmentType) SetOwnerType(value *ManagedDeviceOwnerType)() {
    if m != nil {
        m.ownerType = value
    }
}
