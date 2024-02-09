package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationIdentityCreationConfiguration struct {
    EducationIdentitySynchronizationConfiguration
}
// NewEducationIdentityCreationConfiguration instantiates a new EducationIdentityCreationConfiguration and sets the default values.
func NewEducationIdentityCreationConfiguration()(*EducationIdentityCreationConfiguration) {
    m := &EducationIdentityCreationConfiguration{
        EducationIdentitySynchronizationConfiguration: *NewEducationIdentitySynchronizationConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.educationIdentityCreationConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationIdentityCreationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationIdentityCreationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationIdentityCreationConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationIdentityCreationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationIdentitySynchronizationConfiguration.GetFieldDeserializers()
    res["userDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationIdentityDomainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationIdentityDomainable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationIdentityDomainable)
                }
            }
            m.SetUserDomains(res)
        }
        return nil
    }
    return res
}
// GetUserDomains gets the userDomains property value. The userDomains property
// returns a []EducationIdentityDomainable when successful
func (m *EducationIdentityCreationConfiguration) GetUserDomains()([]EducationIdentityDomainable) {
    val, err := m.GetBackingStore().Get("userDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EducationIdentityDomainable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationIdentityCreationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationIdentitySynchronizationConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserDomains()))
        for i, v := range m.GetUserDomains() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userDomains", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserDomains sets the userDomains property value. The userDomains property
func (m *EducationIdentityCreationConfiguration) SetUserDomains(value []EducationIdentityDomainable)() {
    err := m.GetBackingStore().Set("userDomains", value)
    if err != nil {
        panic(err)
    }
}
type EducationIdentityCreationConfigurationable interface {
    EducationIdentitySynchronizationConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUserDomains()([]EducationIdentityDomainable)
    SetUserDomains(value []EducationIdentityDomainable)()
}
