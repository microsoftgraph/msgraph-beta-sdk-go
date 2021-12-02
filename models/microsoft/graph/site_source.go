package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SiteSource 
type SiteSource struct {
    DataSource
    // 
    site *Site;
}
// NewSiteSource instantiates a new siteSource and sets the default values.
func NewSiteSource()(*SiteSource) {
    m := &SiteSource{
        DataSource: *NewDataSource(),
    }
    return m
}
// GetSite gets the site property value. 
func (m *SiteSource) GetSite()(*Site) {
    if m == nil {
        return nil
    } else {
        return m.site
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SiteSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DataSource.GetFieldDeserializers()
    res["site"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSite() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val.(*Site))
        }
        return nil
    }
    return res
}
func (m *SiteSource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SiteSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DataSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("site", m.GetSite())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSite sets the site property value. 
func (m *SiteSource) SetSite(value *Site)() {
    if m != nil {
        m.site = value
    }
}
