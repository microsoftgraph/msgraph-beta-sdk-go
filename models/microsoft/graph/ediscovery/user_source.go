package ediscovery

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserSource 
type UserSource struct {
    DataSource
    // Email address of the user's mailbox.
    email *string;
    // Specifies which sources are included in this group. Possible values are: mailbox, site.
    includedSources *SourceType;
    // The URL of the user's OneDrive for Business site. Read-only.
    siteWebUrl *string;
}
// NewUserSource instantiates a new userSource and sets the default values.
func NewUserSource()(*UserSource) {
    m := &UserSource{
        DataSource: *NewDataSource(),
    }
    return m
}
// CreateUserSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSourceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserSource(), nil
}
// GetEmail gets the email property value. Email address of the user's mailbox.
func (m *UserSource) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DataSource.GetFieldDeserializers()
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["includedSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludedSources(val.(*SourceType))
        }
        return nil
    }
    res["siteWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteWebUrl(val)
        }
        return nil
    }
    return res
}
// GetIncludedSources gets the includedSources property value. Specifies which sources are included in this group. Possible values are: mailbox, site.
func (m *UserSource) GetIncludedSources()(*SourceType) {
    if m == nil {
        return nil
    } else {
        return m.includedSources
    }
}
// GetSiteWebUrl gets the siteWebUrl property value. The URL of the user's OneDrive for Business site. Read-only.
func (m *UserSource) GetSiteWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteWebUrl
    }
}
// Serialize serializes information the current object
func (m *UserSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DataSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    if m.GetIncludedSources() != nil {
        cast := (*m.GetIncludedSources()).String()
        err = writer.WriteStringValue("includedSources", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("siteWebUrl", m.GetSiteWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmail sets the email property value. Email address of the user's mailbox.
func (m *UserSource) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetIncludedSources sets the includedSources property value. Specifies which sources are included in this group. Possible values are: mailbox, site.
func (m *UserSource) SetIncludedSources(value *SourceType)() {
    if m != nil {
        m.includedSources = value
    }
}
// SetSiteWebUrl sets the siteWebUrl property value. The URL of the user's OneDrive for Business site. Read-only.
func (m *UserSource) SetSiteWebUrl(value *string)() {
    if m != nil {
        m.siteWebUrl = value
    }
}
