package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyObjectFile 
type GroupPolicyObjectFile struct {
    Entity
    // The Group Policy Object file content.
    content *string
    // The date and time at which the GroupPolicy was first uploaded.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Group Policy Object GUID from GPO Xml content
    groupPolicyObjectId *string
    // The date and time at which the GroupPolicyObjectFile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The distinguished name of the OU.
    ouDistinguishedName *string
}
// NewGroupPolicyObjectFile instantiates a new GroupPolicyObjectFile and sets the default values.
func NewGroupPolicyObjectFile()(*GroupPolicyObjectFile) {
    m := &GroupPolicyObjectFile{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyObjectFile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyObjectFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyObjectFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyObjectFile(), nil
}
// GetContent gets the content property value. The Group Policy Object file content.
func (m *GroupPolicyObjectFile) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time at which the GroupPolicy was first uploaded.
func (m *GroupPolicyObjectFile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyObjectFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
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
    res["groupPolicyObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyObjectId(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["ouDistinguishedName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuDistinguishedName(val)
        }
        return nil
    }
    return res
}
// GetGroupPolicyObjectId gets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyObjectFile) GetGroupPolicyObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyObjectId
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time at which the GroupPolicyObjectFile was last modified.
func (m *GroupPolicyObjectFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetOuDistinguishedName gets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyObjectFile) GetOuDistinguishedName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ouDistinguishedName
    }
}
// Serialize serializes information the current object
func (m *GroupPolicyObjectFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
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
    {
        err = writer.WriteStringValue("groupPolicyObjectId", m.GetGroupPolicyObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ouDistinguishedName", m.GetOuDistinguishedName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The Group Policy Object file content.
func (m *GroupPolicyObjectFile) SetContent(value *string)() {
    if m != nil {
        m.content = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time at which the GroupPolicy was first uploaded.
func (m *GroupPolicyObjectFile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetGroupPolicyObjectId sets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyObjectFile) SetGroupPolicyObjectId(value *string)() {
    if m != nil {
        m.groupPolicyObjectId = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time at which the GroupPolicyObjectFile was last modified.
func (m *GroupPolicyObjectFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetOuDistinguishedName sets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyObjectFile) SetOuDistinguishedName(value *string)() {
    if m != nil {
        m.ouDistinguishedName = value
    }
}
