package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new configManagerCollection and sets the default values.
func NewConfigManagerCollection()(*ConfigManagerCollection) {
    m := &ConfigManagerCollection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the collectionIdentifier property value. The collection identifier in SCCM.
func (m *ConfigManagerCollection) GetCollectionIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.collectionIdentifier
    }
}
// Gets the createdDateTime property value. The created date.
func (m *ConfigManagerCollection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the displayName property value. The DisplayName.
func (m *ConfigManagerCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the hierarchyIdentifier property value. The Hierarchy Identifier.
func (m *ConfigManagerCollection) GetHierarchyIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hierarchyIdentifier
    }
}
// Gets the hierarchyName property value. The HierarchyName.
func (m *ConfigManagerCollection) GetHierarchyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hierarchyName
    }
}
// Gets the lastModifiedDateTime property value. The last modified date.
func (m *ConfigManagerCollection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// The deserialization information for the current model
func (m *ConfigManagerCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["collectionIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCollectionIdentifier(val)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["hierarchyIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHierarchyIdentifier(val)
        return nil
    }
    res["hierarchyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHierarchyName(val)
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
    return res
}
func (m *ConfigManagerCollection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the collectionIdentifier property value. The collection identifier in SCCM.
// Parameters:
//  - value : Value to set for the collectionIdentifier property.
func (m *ConfigManagerCollection) SetCollectionIdentifier(value *string)() {
    m.collectionIdentifier = value
}
// Sets the createdDateTime property value. The created date.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ConfigManagerCollection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the displayName property value. The DisplayName.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ConfigManagerCollection) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the hierarchyIdentifier property value. The Hierarchy Identifier.
// Parameters:
//  - value : Value to set for the hierarchyIdentifier property.
func (m *ConfigManagerCollection) SetHierarchyIdentifier(value *string)() {
    m.hierarchyIdentifier = value
}
// Sets the hierarchyName property value. The HierarchyName.
// Parameters:
//  - value : Value to set for the hierarchyName property.
func (m *ConfigManagerCollection) SetHierarchyName(value *string)() {
    m.hierarchyName = value
}
// Sets the lastModifiedDateTime property value. The last modified date.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *ConfigManagerCollection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
