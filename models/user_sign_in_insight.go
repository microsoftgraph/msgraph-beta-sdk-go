package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSignInInsight 
type UserSignInInsight struct {
    GovernanceInsight
    // Indicates when the user last signed in
    lastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewUserSignInInsight instantiates a new UserSignInInsight and sets the default values.
func NewUserSignInInsight()(*UserSignInInsight) {
    m := &UserSignInInsight{
        GovernanceInsight: *NewGovernanceInsight(),
    }
    odataTypeValue := "#microsoft.graph.userSignInInsight";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserSignInInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSignInInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSignInInsight(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSignInInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GovernanceInsight.GetFieldDeserializers()
    res["lastSignInDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastSignInDateTime)
    return res
}
// GetLastSignInDateTime gets the lastSignInDateTime property value. Indicates when the user last signed in
func (m *UserSignInInsight) GetLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSignInDateTime
}
// Serialize serializes information the current object
func (m *UserSignInInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GovernanceInsight.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastSignInDateTime", m.GetLastSignInDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastSignInDateTime sets the lastSignInDateTime property value. Indicates when the user last signed in
func (m *UserSignInInsight) SetLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSignInDateTime = value
}
