package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserRegistrationMethodSummary 
type UserRegistrationMethodSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Total number of users in the tenant.
    totalUserCount *int64
    // Number of users registered for each authentication method.
    userRegistrationMethodCounts []UserRegistrationMethodCountable
    // User role type. Possible values are: all, privilegedAdmin, admin, user.
    userRoles *IncludedUserRoles
    // User type. Possible values are: all, member, guest.
    userTypes *IncludedUserTypes
}
// NewUserRegistrationMethodSummary instantiates a new UserRegistrationMethodSummary and sets the default values.
func NewUserRegistrationMethodSummary()(*UserRegistrationMethodSummary) {
    m := &UserRegistrationMethodSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserRegistrationMethodSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserRegistrationMethodSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserRegistrationMethodSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationMethodSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationMethodSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["totalUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUserCount(val)
        }
        return nil
    }
    res["userRegistrationMethodCounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserRegistrationMethodCountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserRegistrationMethodCountable, len(val))
            for i, v := range val {
                res[i] = v.(UserRegistrationMethodCountable)
            }
            m.SetUserRegistrationMethodCounts(res)
        }
        return nil
    }
    res["userRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncludedUserRoles)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRoles(val.(*IncludedUserRoles))
        }
        return nil
    }
    res["userTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncludedUserTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserTypes(val.(*IncludedUserTypes))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserRegistrationMethodSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetTotalUserCount gets the totalUserCount property value. Total number of users in the tenant.
func (m *UserRegistrationMethodSummary) GetTotalUserCount()(*int64) {
    return m.totalUserCount
}
// GetUserRegistrationMethodCounts gets the userRegistrationMethodCounts property value. Number of users registered for each authentication method.
func (m *UserRegistrationMethodSummary) GetUserRegistrationMethodCounts()([]UserRegistrationMethodCountable) {
    return m.userRegistrationMethodCounts
}
// GetUserRoles gets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationMethodSummary) GetUserRoles()(*IncludedUserRoles) {
    return m.userRoles
}
// GetUserTypes gets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationMethodSummary) GetUserTypes()(*IncludedUserTypes) {
    return m.userTypes
}
// Serialize serializes information the current object
func (m *UserRegistrationMethodSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetUserRegistrationMethodCounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserRegistrationMethodCounts()))
        for i, v := range m.GetUserRegistrationMethodCounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("userRegistrationMethodCounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserRoles() != nil {
        cast := (*m.GetUserRoles()).String()
        err := writer.WriteStringValue("userRoles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserTypes() != nil {
        cast := (*m.GetUserTypes()).String()
        err := writer.WriteStringValue("userTypes", &cast)
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
func (m *UserRegistrationMethodSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserRegistrationMethodSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTotalUserCount sets the totalUserCount property value. Total number of users in the tenant.
func (m *UserRegistrationMethodSummary) SetTotalUserCount(value *int64)() {
    m.totalUserCount = value
}
// SetUserRegistrationMethodCounts sets the userRegistrationMethodCounts property value. Number of users registered for each authentication method.
func (m *UserRegistrationMethodSummary) SetUserRegistrationMethodCounts(value []UserRegistrationMethodCountable)() {
    m.userRegistrationMethodCounts = value
}
// SetUserRoles sets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationMethodSummary) SetUserRoles(value *IncludedUserRoles)() {
    m.userRoles = value
}
// SetUserTypes sets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationMethodSummary) SetUserTypes(value *IncludedUserTypes)() {
    m.userTypes = value
}
