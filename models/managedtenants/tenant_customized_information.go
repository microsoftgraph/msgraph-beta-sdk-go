package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type TenantCustomizedInformation struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewTenantCustomizedInformation instantiates a new TenantCustomizedInformation and sets the default values.
func NewTenantCustomizedInformation()(*TenantCustomizedInformation) {
    m := &TenantCustomizedInformation{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateTenantCustomizedInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTenantCustomizedInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantCustomizedInformation(), nil
}
// GetBusinessRelationship gets the businessRelationship property value. Describes the relationship between the Managed Services Provider and the managed tenant; for example, Managed, Co-managed, Licensing. The maximum length is 250 characters. Optional.
// returns a *string when successful
func (m *TenantCustomizedInformation) GetBusinessRelationship()(*string) {
    val, err := m.GetBackingStore().Get("businessRelationship")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetComplianceRequirements gets the complianceRequirements property value. Contains the compliance requirements for the customer tenant; for example, HIPPA, NIST, CMMC. The maximum length is 250 characters per compliance requirement. Optional.
// returns a []string when successful
func (m *TenantCustomizedInformation) GetComplianceRequirements()([]string) {
    val, err := m.GetBackingStore().Get("complianceRequirements")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetContacts gets the contacts property value. The collection of contacts for the managed tenant. Optional.
// returns a []TenantContactInformationable when successful
func (m *TenantCustomizedInformation) GetContacts()([]TenantContactInformationable) {
    val, err := m.GetBackingStore().Get("contacts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TenantContactInformationable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the managed tenant. Required. Read-only.
// returns a *string when successful
func (m *TenantCustomizedInformation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TenantCustomizedInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessRelationship"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessRelationship(val)
        }
        return nil
    }
    res["complianceRequirements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetComplianceRequirements(res)
        }
        return nil
    }
    res["contacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantContactInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantContactInformationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TenantContactInformationable)
                }
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
    res["managedServicesPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetManagedServicesPlans(res)
        }
        return nil
    }
    res["note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
        }
        return nil
    }
    res["noteLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoteLastModifiedDateTime(val)
        }
        return nil
    }
    res["partnerRelationshipManagerUserIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPartnerRelationshipManagerUserIds(res)
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
// GetManagedServicesPlans gets the managedServicesPlans property value. This is the Managed Services Plans for the customer tenant that the Managed Services Provider manages. The maximum length is 250 characters per managed service plan. Optional.
// returns a []string when successful
func (m *TenantCustomizedInformation) GetManagedServicesPlans()([]string) {
    val, err := m.GetBackingStore().Get("managedServicesPlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetNote gets the note property value. A field for the Managed Services Provider technician to input custom text to share notes between technicians within the Managed Service Providers. The maximum length is 5000 characters. Optional.
// returns a *string when successful
func (m *TenantCustomizedInformation) GetNote()(*string) {
    val, err := m.GetBackingStore().Get("note")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNoteLastModifiedDateTime gets the noteLastModifiedDateTime property value. The date on which the note field of this entity was last modified. Optional.
// returns a *Time when successful
func (m *TenantCustomizedInformation) GetNoteLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("noteLastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPartnerRelationshipManagerUserIds gets the partnerRelationshipManagerUserIds property value. The list of Entra user IDs for users in the Managed Services Provider that manage the relationship with the managed tenant. Optional.
// returns a []string when successful
func (m *TenantCustomizedInformation) GetPartnerRelationshipManagerUserIds()([]string) {
    val, err := m.GetBackingStore().Get("partnerRelationshipManagerUserIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The Microsoft Entra tenant identifier for the managed tenant. Optional. Read-only.
// returns a *string when successful
func (m *TenantCustomizedInformation) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWebsite gets the website property value. The website for the managed tenant. Required.
// returns a *string when successful
func (m *TenantCustomizedInformation) GetWebsite()(*string) {
    val, err := m.GetBackingStore().Get("website")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TenantCustomizedInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("businessRelationship", m.GetBusinessRelationship())
        if err != nil {
            return err
        }
    }
    if m.GetComplianceRequirements() != nil {
        err = writer.WriteCollectionOfStringValues("complianceRequirements", m.GetComplianceRequirements())
        if err != nil {
            return err
        }
    }
    if m.GetContacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContacts()))
        for i, v := range m.GetContacts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetManagedServicesPlans() != nil {
        err = writer.WriteCollectionOfStringValues("managedServicesPlans", m.GetManagedServicesPlans())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("note", m.GetNote())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("noteLastModifiedDateTime", m.GetNoteLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerRelationshipManagerUserIds() != nil {
        err = writer.WriteCollectionOfStringValues("partnerRelationshipManagerUserIds", m.GetPartnerRelationshipManagerUserIds())
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
// SetBusinessRelationship sets the businessRelationship property value. Describes the relationship between the Managed Services Provider and the managed tenant; for example, Managed, Co-managed, Licensing. The maximum length is 250 characters. Optional.
func (m *TenantCustomizedInformation) SetBusinessRelationship(value *string)() {
    err := m.GetBackingStore().Set("businessRelationship", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceRequirements sets the complianceRequirements property value. Contains the compliance requirements for the customer tenant; for example, HIPPA, NIST, CMMC. The maximum length is 250 characters per compliance requirement. Optional.
func (m *TenantCustomizedInformation) SetComplianceRequirements(value []string)() {
    err := m.GetBackingStore().Set("complianceRequirements", value)
    if err != nil {
        panic(err)
    }
}
// SetContacts sets the contacts property value. The collection of contacts for the managed tenant. Optional.
func (m *TenantCustomizedInformation) SetContacts(value []TenantContactInformationable)() {
    err := m.GetBackingStore().Set("contacts", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the managed tenant. Required. Read-only.
func (m *TenantCustomizedInformation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedServicesPlans sets the managedServicesPlans property value. This is the Managed Services Plans for the customer tenant that the Managed Services Provider manages. The maximum length is 250 characters per managed service plan. Optional.
func (m *TenantCustomizedInformation) SetManagedServicesPlans(value []string)() {
    err := m.GetBackingStore().Set("managedServicesPlans", value)
    if err != nil {
        panic(err)
    }
}
// SetNote sets the note property value. A field for the Managed Services Provider technician to input custom text to share notes between technicians within the Managed Service Providers. The maximum length is 5000 characters. Optional.
func (m *TenantCustomizedInformation) SetNote(value *string)() {
    err := m.GetBackingStore().Set("note", value)
    if err != nil {
        panic(err)
    }
}
// SetNoteLastModifiedDateTime sets the noteLastModifiedDateTime property value. The date on which the note field of this entity was last modified. Optional.
func (m *TenantCustomizedInformation) SetNoteLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("noteLastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerRelationshipManagerUserIds sets the partnerRelationshipManagerUserIds property value. The list of Entra user IDs for users in the Managed Services Provider that manage the relationship with the managed tenant. Optional.
func (m *TenantCustomizedInformation) SetPartnerRelationshipManagerUserIds(value []string)() {
    err := m.GetBackingStore().Set("partnerRelationshipManagerUserIds", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The Microsoft Entra tenant identifier for the managed tenant. Optional. Read-only.
func (m *TenantCustomizedInformation) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetWebsite sets the website property value. The website for the managed tenant. Required.
func (m *TenantCustomizedInformation) SetWebsite(value *string)() {
    err := m.GetBackingStore().Set("website", value)
    if err != nil {
        panic(err)
    }
}
type TenantCustomizedInformationable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusinessRelationship()(*string)
    GetComplianceRequirements()([]string)
    GetContacts()([]TenantContactInformationable)
    GetDisplayName()(*string)
    GetManagedServicesPlans()([]string)
    GetNote()(*string)
    GetNoteLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPartnerRelationshipManagerUserIds()([]string)
    GetTenantId()(*string)
    GetWebsite()(*string)
    SetBusinessRelationship(value *string)()
    SetComplianceRequirements(value []string)()
    SetContacts(value []TenantContactInformationable)()
    SetDisplayName(value *string)()
    SetManagedServicesPlans(value []string)()
    SetNote(value *string)()
    SetNoteLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPartnerRelationshipManagerUserIds(value []string)()
    SetTenantId(value *string)()
    SetWebsite(value *string)()
}
