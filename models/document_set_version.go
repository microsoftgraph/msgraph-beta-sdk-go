package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentSetVersion 
type DocumentSetVersion struct {
    ListItemVersion
    // The comment property
    comment *string
    // The createdBy property
    createdBy IdentitySetable
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The items property
    items []DocumentSetVersionItemable
    // The shouldCaptureMinorVersion property
    shouldCaptureMinorVersion *bool
}
// NewDocumentSetVersion instantiates a new documentSetVersion and sets the default values.
func NewDocumentSetVersion()(*DocumentSetVersion) {
    m := &DocumentSetVersion{
        ListItemVersion: *NewListItemVersion(),
    }
    return m
}
// CreateDocumentSetVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentSetVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentSetVersion(), nil
}
// GetComment gets the comment property value. The comment property
func (m *DocumentSetVersion) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *DocumentSetVersion) GetCreatedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *DocumentSetVersion) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentSetVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ListItemVersion.GetFieldDeserializers()
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDocumentSetVersionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DocumentSetVersionItemable, len(val))
            for i, v := range val {
                res[i] = v.(DocumentSetVersionItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["shouldCaptureMinorVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldCaptureMinorVersion(val)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
func (m *DocumentSetVersion) GetItems()([]DocumentSetVersionItemable) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetShouldCaptureMinorVersion gets the shouldCaptureMinorVersion property value. The shouldCaptureMinorVersion property
func (m *DocumentSetVersion) GetShouldCaptureMinorVersion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shouldCaptureMinorVersion
    }
}
// Serialize serializes information the current object
func (m *DocumentSetVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ListItemVersion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shouldCaptureMinorVersion", m.GetShouldCaptureMinorVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComment sets the comment property value. The comment property
func (m *DocumentSetVersion) SetComment(value *string)() {
    if m != nil {
        m.comment = value
    }
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *DocumentSetVersion) SetCreatedBy(value IdentitySetable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *DocumentSetVersion) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetItems sets the items property value. The items property
func (m *DocumentSetVersion) SetItems(value []DocumentSetVersionItemable)() {
    if m != nil {
        m.items = value
    }
}
// SetShouldCaptureMinorVersion sets the shouldCaptureMinorVersion property value. The shouldCaptureMinorVersion property
func (m *DocumentSetVersion) SetShouldCaptureMinorVersion(value *bool)() {
    if m != nil {
        m.shouldCaptureMinorVersion = value
    }
}
