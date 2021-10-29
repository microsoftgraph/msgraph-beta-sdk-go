package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GroupPolicyObjectFile struct {
    Entity
    // The Group Policy Object file content.
    content *string;
    // The date and time at which the GroupPolicy was first uploaded.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Group Policy Object GUID from GPO Xml content
    groupPolicyObjectId *string;
    // The date and time at which the GroupPolicyObjectFile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The distinguished name of the OU.
    ouDistinguishedName *string;
}
// Instantiates a new groupPolicyObjectFile and sets the default values.
func NewGroupPolicyObjectFile()(*GroupPolicyObjectFile) {
    m := &GroupPolicyObjectFile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the content property value. The Group Policy Object file content.
func (m *GroupPolicyObjectFile) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the createdDateTime property value. The date and time at which the GroupPolicy was first uploaded.
func (m *GroupPolicyObjectFile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyObjectFile) GetGroupPolicyObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyObjectId
    }
}
// Gets the lastModifiedDateTime property value. The date and time at which the GroupPolicyObjectFile was last modified.
func (m *GroupPolicyObjectFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyObjectFile) GetOuDistinguishedName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ouDistinguishedName
    }
}
// The deserialization information for the current model
func (m *GroupPolicyObjectFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["groupPolicyObjectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupPolicyObjectId(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["ouDistinguishedName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOuDistinguishedName(val)
        return nil
    }
    return res
}
func (m *GroupPolicyObjectFile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GroupPolicyObjectFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the content property value. The Group Policy Object file content.
// Parameters:
//  - value : Value to set for the content property.
func (m *GroupPolicyObjectFile) SetContent(value *string)() {
    m.content = value
}
// Sets the createdDateTime property value. The date and time at which the GroupPolicy was first uploaded.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *GroupPolicyObjectFile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
// Parameters:
//  - value : Value to set for the groupPolicyObjectId property.
func (m *GroupPolicyObjectFile) SetGroupPolicyObjectId(value *string)() {
    m.groupPolicyObjectId = value
}
// Sets the lastModifiedDateTime property value. The date and time at which the GroupPolicyObjectFile was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *GroupPolicyObjectFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the ouDistinguishedName property value. The distinguished name of the OU.
// Parameters:
//  - value : Value to set for the ouDistinguishedName property.
func (m *GroupPolicyObjectFile) SetOuDistinguishedName(value *string)() {
    m.ouDistinguishedName = value
}
