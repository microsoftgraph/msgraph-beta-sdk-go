package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryHoldPolicy 
type EdiscoveryHoldPolicy struct {
    PolicyBase
    // The contentQuery property
    contentQuery *string
    // The errors property
    errors []string
    // The isEnabled property
    isEnabled *bool
    // The siteSources property
    siteSources []SiteSourceable
    // The userSources property
    userSources []UserSourceable
}
// NewEdiscoveryHoldPolicy instantiates a new ediscoveryHoldPolicy and sets the default values.
func NewEdiscoveryHoldPolicy()(*EdiscoveryHoldPolicy) {
    m := &EdiscoveryHoldPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// CreateEdiscoveryHoldPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryHoldPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryHoldPolicy(), nil
}
// GetContentQuery gets the contentQuery property value. The contentQuery property
func (m *EdiscoveryHoldPolicy) GetContentQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentQuery
    }
}
// GetErrors gets the errors property value. The errors property
func (m *EdiscoveryHoldPolicy) GetErrors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryHoldPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["contentQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentQuery(val)
        }
        return nil
    }
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["siteSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSiteSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SiteSourceable, len(val))
            for i, v := range val {
                res[i] = v.(SiteSourceable)
            }
            m.SetSiteSources(res)
        }
        return nil
    }
    res["userSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSourceable, len(val))
            for i, v := range val {
                res[i] = v.(UserSourceable)
            }
            m.SetUserSources(res)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. The isEnabled property
func (m *EdiscoveryHoldPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetSiteSources gets the siteSources property value. The siteSources property
func (m *EdiscoveryHoldPolicy) GetSiteSources()([]SiteSourceable) {
    if m == nil {
        return nil
    } else {
        return m.siteSources
    }
}
// GetUserSources gets the userSources property value. The userSources property
func (m *EdiscoveryHoldPolicy) GetUserSources()([]UserSourceable) {
    if m == nil {
        return nil
    } else {
        return m.userSources
    }
}
// Serialize serializes information the current object
func (m *EdiscoveryHoldPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
        if err != nil {
            return err
        }
    }
    if m.GetErrors() != nil {
        err = writer.WriteCollectionOfStringValues("errors", m.GetErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetSiteSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSiteSources()))
        for i, v := range m.GetSiteSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("siteSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSources()))
        for i, v := range m.GetUserSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentQuery sets the contentQuery property value. The contentQuery property
func (m *EdiscoveryHoldPolicy) SetContentQuery(value *string)() {
    if m != nil {
        m.contentQuery = value
    }
}
// SetErrors sets the errors property value. The errors property
func (m *EdiscoveryHoldPolicy) SetErrors(value []string)() {
    if m != nil {
        m.errors = value
    }
}
// SetIsEnabled sets the isEnabled property value. The isEnabled property
func (m *EdiscoveryHoldPolicy) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetSiteSources sets the siteSources property value. The siteSources property
func (m *EdiscoveryHoldPolicy) SetSiteSources(value []SiteSourceable)() {
    if m != nil {
        m.siteSources = value
    }
}
// SetUserSources sets the userSources property value. The userSources property
func (m *EdiscoveryHoldPolicy) SetUserSources(value []UserSourceable)() {
    if m != nil {
        m.userSources = value
    }
}
