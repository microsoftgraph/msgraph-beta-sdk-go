package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationCount 
type UserRegistrationCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Provides the registration count for your tenant.
    registrationCount *int64
    // Represents the status of user registration. Possible values are: registered, enabled, capable, and mfaRegistered.
    registrationStatus *RegistrationStatusType
}
// NewUserRegistrationCount instantiates a new userRegistrationCount and sets the default values.
func NewUserRegistrationCount()(*UserRegistrationCount) {
    m := &UserRegistrationCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserRegistrationCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserRegistrationCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserRegistrationCount(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationCount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["registrationCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationCount(val)
        }
        return nil
    }
    res["registrationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistrationStatusType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationStatus(val.(*RegistrationStatusType))
        }
        return nil
    }
    return res
}
// GetRegistrationCount gets the registrationCount property value. Provides the registration count for your tenant.
func (m *UserRegistrationCount) GetRegistrationCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.registrationCount
    }
}
// GetRegistrationStatus gets the registrationStatus property value. Represents the status of user registration. Possible values are: registered, enabled, capable, and mfaRegistered.
func (m *UserRegistrationCount) GetRegistrationStatus()(*RegistrationStatusType) {
    if m == nil {
        return nil
    } else {
        return m.registrationStatus
    }
}
// Serialize serializes information the current object
func (m *UserRegistrationCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("registrationCount", m.GetRegistrationCount())
        if err != nil {
            return err
        }
    }
    if m.GetRegistrationStatus() != nil {
        cast := (*m.GetRegistrationStatus()).String()
        err := writer.WriteStringValue("registrationStatus", &cast)
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
func (m *UserRegistrationCount) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRegistrationCount sets the registrationCount property value. Provides the registration count for your tenant.
func (m *UserRegistrationCount) SetRegistrationCount(value *int64)() {
    if m != nil {
        m.registrationCount = value
    }
}
// SetRegistrationStatus sets the registrationStatus property value. Represents the status of user registration. Possible values are: registered, enabled, capable, and mfaRegistered.
func (m *UserRegistrationCount) SetRegistrationStatus(value *RegistrationStatusType)() {
    if m != nil {
        m.registrationStatus = value
    }
}
