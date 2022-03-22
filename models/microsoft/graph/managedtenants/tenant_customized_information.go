package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// TenantCustomizedInformation 
type TenantCustomizedInformation struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The collection of contacts for the managed tenant. Optional.
    contacts []TenantContactInformationable;
    // The display name for the managed tenant. Required. Read-only.
    displayName *string;
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string;
    // The website for the managed tenant. Required.
    website *string;
}
// NewTenantCustomizedInformation instantiates a new tenantCustomizedInformation and sets the default values.
func NewTenantCustomizedInformation()(*TenantCustomizedInformation) {
    m := &TenantCustomizedInformation{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateTenantCustomizedInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantCustomizedInformationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTenantCustomizedInformation(), nil
}
// GetContacts gets the contacts property value. The collection of contacts for the managed tenant. Optional.
func (m *TenantCustomizedInformation) GetContacts()([]TenantContactInformationable) {
    if m == nil {
        return nil
    } else {
        return m.contacts
    }
}
// GetDisplayName gets the displayName property value. The display name for the managed tenant. Required. Read-only.
func (m *TenantCustomizedInformation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantCustomizedInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contacts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantContactInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantContactInformationable, len(val))
            for i, v := range val {
                res[i] = v.(TenantContactInformationable)
            }
            m.SetContacts(res)
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
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["website"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebsite(val)
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *TenantCustomizedInformation) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetWebsite gets the website property value. The website for the managed tenant. Required.
func (m *TenantCustomizedInformation) GetWebsite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.website
    }
}
// Serialize serializes information the current object
func (m *TenantCustomizedInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContacts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContacts()))
        for i, v := range m.GetContacts() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("contacts", cast)
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
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("website", m.GetWebsite())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContacts sets the contacts property value. The collection of contacts for the managed tenant. Optional.
func (m *TenantCustomizedInformation) SetContacts(value []TenantContactInformationable)() {
    if m != nil {
        m.contacts = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the managed tenant. Required. Read-only.
func (m *TenantCustomizedInformation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *TenantCustomizedInformation) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetWebsite sets the website property value. The website for the managed tenant. Required.
func (m *TenantCustomizedInformation) SetWebsite(value *string)() {
    if m != nil {
        m.website = value
    }
}
