package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConfigManagerCollection 
type ConfigManagerCollection struct {
    Entity
    // The collection identifier in SCCM.
    collectionIdentifier *string;
    // The created date.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The DisplayName.
    displayName *string;
    // The Hierarchy Identifier.
    hierarchyIdentifier *string;
    // The HierarchyName.
    hierarchyName *string;
    // The last modified date.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewConfigManagerCollection instantiates a new configManagerCollection and sets the default values.
func NewConfigManagerCollection()(*ConfigManagerCollection) {
    m := &ConfigManagerCollection{
        Entity: *NewEntity(),
    }
    return m
}
// GetCollectionIdentifier gets the collectionIdentifier property value. The collection identifier in SCCM.
func (m *ConfigManagerCollection) GetCollectionIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.collectionIdentifier
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The created date.
func (m *ConfigManagerCollection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. The DisplayName.
func (m *ConfigManagerCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetHierarchyIdentifier gets the hierarchyIdentifier property value. The Hierarchy Identifier.
func (m *ConfigManagerCollection) GetHierarchyIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hierarchyIdentifier
    }
}
// GetHierarchyName gets the hierarchyName property value. The HierarchyName.
func (m *ConfigManagerCollection) GetHierarchyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hierarchyName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last modified date.
func (m *ConfigManagerCollection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigManagerCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["collectionIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollectionIdentifier(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["hierarchyIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHierarchyIdentifier(val)
        }
        return nil
    }
    res["hierarchyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHierarchyName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *ConfigManagerCollection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConfigManagerCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("collectionIdentifier", m.GetCollectionIdentifier())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hierarchyIdentifier", m.GetHierarchyIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hierarchyName", m.GetHierarchyName())
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
    return nil
}
// SetCollectionIdentifier sets the collectionIdentifier property value. The collection identifier in SCCM.
func (m *ConfigManagerCollection) SetCollectionIdentifier(value *string)() {
    if m != nil {
        m.collectionIdentifier = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The created date.
func (m *ConfigManagerCollection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. The DisplayName.
func (m *ConfigManagerCollection) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHierarchyIdentifier sets the hierarchyIdentifier property value. The Hierarchy Identifier.
func (m *ConfigManagerCollection) SetHierarchyIdentifier(value *string)() {
    if m != nil {
        m.hierarchyIdentifier = value
    }
}
// SetHierarchyName sets the hierarchyName property value. The HierarchyName.
func (m *ConfigManagerCollection) SetHierarchyName(value *string)() {
    if m != nil {
        m.hierarchyName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last modified date.
func (m *ConfigManagerCollection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
