package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CatalogEntry 
type CatalogEntry struct {
    Entity
    // The date on which the content is no longer available to deploy using the service. Read-only.
    deployableUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The display name of the content. Read-only.
    displayName *string;
    // The release date for the content. Read-only.
    releaseDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewCatalogEntry instantiates a new catalogEntry and sets the default values.
func NewCatalogEntry()(*CatalogEntry) {
    m := &CatalogEntry{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeployableUntilDateTime gets the deployableUntilDateTime property value. The date on which the content is no longer available to deploy using the service. Read-only.
func (m *CatalogEntry) GetDeployableUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deployableUntilDateTime
    }
}
// GetDisplayName gets the displayName property value. The display name of the content. Read-only.
func (m *CatalogEntry) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetReleaseDateTime gets the releaseDateTime property value. The release date for the content. Read-only.
func (m *CatalogEntry) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.releaseDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CatalogEntry) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deployableUntilDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployableUntilDateTime(val)
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
    res["releaseDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    return res
}
func (m *CatalogEntry) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CatalogEntry) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("deployableUntilDateTime", m.GetDeployableUntilDateTime())
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
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeployableUntilDateTime sets the deployableUntilDateTime property value. The date on which the content is no longer available to deploy using the service. Read-only.
func (m *CatalogEntry) SetDeployableUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deployableUntilDateTime = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the content. Read-only.
func (m *CatalogEntry) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetReleaseDateTime sets the releaseDateTime property value. The release date for the content. Read-only.
func (m *CatalogEntry) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.releaseDateTime = value
    }
}
