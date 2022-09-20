package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantCustomizedInformation provides operations to manage the collection of activityStatistics entities.
type TenantCustomizedInformation struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The collection of contacts for the managed tenant. Optional.
    contacts []TenantContactInformationable
    // The display name for the managed tenant. Required. Read-only.
    displayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string
    // The website for the managed tenant. Required.
    website *string
}
// NewTenantCustomizedInformation instantiates a new tenantCustomizedInformation and sets the default values.
func NewTenantCustomizedInformation()(*TenantCustomizedInformation) {
    m := &TenantCustomizedInformation{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.tenantCustomizedInformation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTenantCustomizedInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantCustomizedInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantCustomizedInformation(), nil
}
// GetContacts gets the contacts property value. The collection of contacts for the managed tenant. Optional.
func (m *TenantCustomizedInformation) GetContacts()([]TenantContactInformationable) {
    return m.contacts
}
// GetDisplayName gets the displayName property value. The display name for the managed tenant. Required. Read-only.
func (m *TenantCustomizedInformation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantCustomizedInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["website"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    return m.tenantId
}
// GetWebsite gets the website property value. The website for the managed tenant. Required.
func (m *TenantCustomizedInformation) GetWebsite()(*string) {
    return m.website
}
// Serialize serializes information the current object
func (m *TenantCustomizedInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContacts()))
        for i, v := range m.GetContacts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    m.contacts = value
}
// SetDisplayName sets the displayName property value. The display name for the managed tenant. Required. Read-only.
func (m *TenantCustomizedInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *TenantCustomizedInformation) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetWebsite sets the website property value. The website for the managed tenant. Required.
func (m *TenantCustomizedInformation) SetWebsite(value *string)() {
    m.website = value
}
