package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ThreatSubmissionRoot 
type ThreatSubmissionRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
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
    val, err := m.GetBackingStore().Get("emailThreats")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EmailThreatSubmissionable)
    }
    return nil
}
// GetEmailThreatSubmissionPolicies gets the emailThreatSubmissionPolicies property value. The emailThreatSubmissionPolicies property
func (m *ThreatSubmissionRoot) GetEmailThreatSubmissionPolicies()([]EmailThreatSubmissionPolicyable) {
    val, err := m.GetBackingStore().Get("emailThreatSubmissionPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EmailThreatSubmissionPolicyable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThreatSubmissionRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailThreats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailThreatSubmissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailThreatSubmissionable, len(val))
            for i, v := range val {
                res[i] = v.(EmailThreatSubmissionable)
            }
            m.SetEmailThreats(res)
        }
        return nil
    }
    res["emailThreatSubmissionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailThreatSubmissionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailThreatSubmissionPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(EmailThreatSubmissionPolicyable)
            }
            m.SetEmailThreatSubmissionPolicies(res)
        }
        return nil
    }
    res["fileThreats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFileThreatSubmissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FileThreatSubmissionable, len(val))
            for i, v := range val {
                res[i] = v.(FileThreatSubmissionable)
            }
            m.SetFileThreats(res)
        }
        return nil
    }
    res["urlThreats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUrlThreatSubmissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UrlThreatSubmissionable, len(val))
            for i, v := range val {
                res[i] = v.(UrlThreatSubmissionable)
            }
            m.SetUrlThreats(res)
        }
        return nil
    }
    return res
}
// GetFileThreats gets the fileThreats property value. The fileThreats property
func (m *ThreatSubmissionRoot) GetFileThreats()([]FileThreatSubmissionable) {
    val, err := m.GetBackingStore().Get("fileThreats")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FileThreatSubmissionable)
    }
    return nil
}
// GetUrlThreats gets the urlThreats property value. The urlThreats property
func (m *ThreatSubmissionRoot) GetUrlThreats()([]UrlThreatSubmissionable) {
    val, err := m.GetBackingStore().Get("urlThreats")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UrlThreatSubmissionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ThreatSubmissionRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEmailThreats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailThreats()))
        for i, v := range m.GetEmailThreats() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emailThreats", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmailThreatSubmissionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailThreatSubmissionPolicies()))
        for i, v := range m.GetEmailThreatSubmissionPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emailThreatSubmissionPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFileThreats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFileThreats()))
        for i, v := range m.GetFileThreats() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("fileThreats", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUrlThreats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUrlThreats()))
        for i, v := range m.GetUrlThreats() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("urlThreats", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailThreats sets the emailThreats property value. The emailThreats property
func (m *ThreatSubmissionRoot) SetEmailThreats(value []EmailThreatSubmissionable)() {
    err := m.GetBackingStore().Set("emailThreats", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailThreatSubmissionPolicies sets the emailThreatSubmissionPolicies property value. The emailThreatSubmissionPolicies property
func (m *ThreatSubmissionRoot) SetEmailThreatSubmissionPolicies(value []EmailThreatSubmissionPolicyable)() {
    err := m.GetBackingStore().Set("emailThreatSubmissionPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetFileThreats sets the fileThreats property value. The fileThreats property
func (m *ThreatSubmissionRoot) SetFileThreats(value []FileThreatSubmissionable)() {
    err := m.GetBackingStore().Set("fileThreats", value)
    if err != nil {
        panic(err)
    }
}
// SetUrlThreats sets the urlThreats property value. The urlThreats property
func (m *ThreatSubmissionRoot) SetUrlThreats(value []UrlThreatSubmissionable)() {
    err := m.GetBackingStore().Set("urlThreats", value)
    if err != nil {
        panic(err)
    }
}
// ThreatSubmissionRootable 
type ThreatSubmissionRootable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailThreats()([]EmailThreatSubmissionable)
    GetEmailThreatSubmissionPolicies()([]EmailThreatSubmissionPolicyable)
    GetFileThreats()([]FileThreatSubmissionable)
    GetUrlThreats()([]UrlThreatSubmissionable)
    SetEmailThreats(value []EmailThreatSubmissionable)()
    SetEmailThreatSubmissionPolicies(value []EmailThreatSubmissionPolicyable)()
    SetFileThreats(value []FileThreatSubmissionable)()
    SetUrlThreats(value []UrlThreatSubmissionable)()
}
