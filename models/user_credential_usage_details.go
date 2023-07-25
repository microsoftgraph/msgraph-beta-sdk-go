package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserCredentialUsageDetails 
type UserCredentialUsageDetails struct {
    Entity
}
// NewUserCredentialUsageDetails instantiates a new userCredentialUsageDetails and sets the default values.
func NewUserCredentialUsageDetails()(*UserCredentialUsageDetails) {
    m := &UserCredentialUsageDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserCredentialUsageDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserCredentialUsageDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserCredentialUsageDetails(), nil
}
// GetAuthMethod gets the authMethod property value. The authMethod property
func (m *UserCredentialUsageDetails) GetAuthMethod()(*UsageAuthMethod) {
    val, err := m.GetBackingStore().Get("authMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UsageAuthMethod)
    }
    return nil
}
// GetEventDateTime gets the eventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *UserCredentialUsageDetails) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("eventDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFailureReason gets the failureReason property value. Provides the failure reason for the corresponding reset or registration workflow.
func (m *UserCredentialUsageDetails) GetFailureReason()(*string) {
    val, err := m.GetBackingStore().Get("failureReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFeature gets the feature property value. The feature property
func (m *UserCredentialUsageDetails) GetFeature()(*FeatureType) {
    val, err := m.GetBackingStore().Get("feature")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FeatureType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserCredentialUsageDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUsageAuthMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthMethod(val.(*UsageAuthMethod))
        }
        return nil
    }
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["failureReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureReason(val)
        }
        return nil
    }
    res["feature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFeatureType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeature(val.(*FeatureType))
        }
        return nil
    }
    res["isSuccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSuccess(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetIsSuccess gets the isSuccess property value. Indicates success or failure of the workflow.
func (m *UserCredentialUsageDetails) GetIsSuccess()(*bool) {
    val, err := m.GetBackingStore().Get("isSuccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserDisplayName gets the userDisplayName property value. User name of the user performing the reset or registration workflow.
func (m *UserCredentialUsageDetails) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name of the user performing the reset or registration workflow.
func (m *UserCredentialUsageDetails) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserCredentialUsageDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthMethod() != nil {
        cast := (*m.GetAuthMethod()).String()
        err = writer.WriteStringValue("authMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("failureReason", m.GetFailureReason())
        if err != nil {
            return err
        }
    }
    if m.GetFeature() != nil {
        cast := (*m.GetFeature()).String()
        err = writer.WriteStringValue("feature", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSuccess", m.GetIsSuccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthMethod sets the authMethod property value. The authMethod property
func (m *UserCredentialUsageDetails) SetAuthMethod(value *UsageAuthMethod)() {
    err := m.GetBackingStore().Set("authMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetEventDateTime sets the eventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *UserCredentialUsageDetails) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("eventDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFailureReason sets the failureReason property value. Provides the failure reason for the corresponding reset or registration workflow.
func (m *UserCredentialUsageDetails) SetFailureReason(value *string)() {
    err := m.GetBackingStore().Set("failureReason", value)
    if err != nil {
        panic(err)
    }
}
// SetFeature sets the feature property value. The feature property
func (m *UserCredentialUsageDetails) SetFeature(value *FeatureType)() {
    err := m.GetBackingStore().Set("feature", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSuccess sets the isSuccess property value. Indicates success or failure of the workflow.
func (m *UserCredentialUsageDetails) SetIsSuccess(value *bool)() {
    err := m.GetBackingStore().Set("isSuccess", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDisplayName sets the userDisplayName property value. User name of the user performing the reset or registration workflow.
func (m *UserCredentialUsageDetails) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name of the user performing the reset or registration workflow.
func (m *UserCredentialUsageDetails) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// UserCredentialUsageDetailsable 
type UserCredentialUsageDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthMethod()(*UsageAuthMethod)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailureReason()(*string)
    GetFeature()(*FeatureType)
    GetIsSuccess()(*bool)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetAuthMethod(value *UsageAuthMethod)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailureReason(value *string)()
    SetFeature(value *FeatureType)()
    SetIsSuccess(value *bool)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
