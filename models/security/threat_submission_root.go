package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ThreatSubmissionRoot 
type ThreatSubmissionRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The emailThreats property
    emailThreats []EmailThreatSubmissionable
    // The emailThreatSubmissionPolicies property
    emailThreatSubmissionPolicies []EmailThreatSubmissionPolicyable
    // The fileThreats property
    fileThreats []FileThreatSubmissionable
    // The urlThreats property
    urlThreats []UrlThreatSubmissionable
}
// NewThreatSubmissionRoot instantiates a new threatSubmissionRoot and sets the default values.
func NewThreatSubmissionRoot()(*ThreatSubmissionRoot) {
    m := &ThreatSubmissionRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateThreatSubmissionRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThreatSubmissionRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThreatSubmissionRoot(), nil
}
// GetEmailThreats gets the emailThreats property value. The emailThreats property
func (m *ThreatSubmissionRoot) GetEmailThreats()([]EmailThreatSubmissionable) {
    return m.emailThreats
}
// GetEmailThreatSubmissionPolicies gets the emailThreatSubmissionPolicies property value. The emailThreatSubmissionPolicies property
func (m *ThreatSubmissionRoot) GetEmailThreatSubmissionPolicies()([]EmailThreatSubmissionPolicyable) {
    return m.emailThreatSubmissionPolicies
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThreatSubmissionRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailThreats"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEmailThreatSubmissionFromDiscriminatorValue , m.SetEmailThreats)
    res["emailThreatSubmissionPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEmailThreatSubmissionPolicyFromDiscriminatorValue , m.SetEmailThreatSubmissionPolicies)
    res["fileThreats"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateFileThreatSubmissionFromDiscriminatorValue , m.SetFileThreats)
    res["urlThreats"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUrlThreatSubmissionFromDiscriminatorValue , m.SetUrlThreats)
    return res
}
// GetFileThreats gets the fileThreats property value. The fileThreats property
func (m *ThreatSubmissionRoot) GetFileThreats()([]FileThreatSubmissionable) {
    return m.fileThreats
}
// GetUrlThreats gets the urlThreats property value. The urlThreats property
func (m *ThreatSubmissionRoot) GetUrlThreats()([]UrlThreatSubmissionable) {
    return m.urlThreats
}
// Serialize serializes information the current object
func (m *ThreatSubmissionRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEmailThreats() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEmailThreats())
        err = writer.WriteCollectionOfObjectValues("emailThreats", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmailThreatSubmissionPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEmailThreatSubmissionPolicies())
        err = writer.WriteCollectionOfObjectValues("emailThreatSubmissionPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileThreats() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFileThreats())
        err = writer.WriteCollectionOfObjectValues("fileThreats", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUrlThreats() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUrlThreats())
        err = writer.WriteCollectionOfObjectValues("urlThreats", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailThreats sets the emailThreats property value. The emailThreats property
func (m *ThreatSubmissionRoot) SetEmailThreats(value []EmailThreatSubmissionable)() {
    m.emailThreats = value
}
// SetEmailThreatSubmissionPolicies sets the emailThreatSubmissionPolicies property value. The emailThreatSubmissionPolicies property
func (m *ThreatSubmissionRoot) SetEmailThreatSubmissionPolicies(value []EmailThreatSubmissionPolicyable)() {
    m.emailThreatSubmissionPolicies = value
}
// SetFileThreats sets the fileThreats property value. The fileThreats property
func (m *ThreatSubmissionRoot) SetFileThreats(value []FileThreatSubmissionable)() {
    m.fileThreats = value
}
// SetUrlThreats sets the urlThreats property value. The urlThreats property
func (m *ThreatSubmissionRoot) SetUrlThreats(value []UrlThreatSubmissionable)() {
    m.urlThreats = value
}
