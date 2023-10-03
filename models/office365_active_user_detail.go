package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365ActiveUserDetail 
type Office365ActiveUserDetail struct {
    Entity
}
// NewOffice365ActiveUserDetail instantiates a new office365ActiveUserDetail and sets the default values.
func NewOffice365ActiveUserDetail()(*Office365ActiveUserDetail) {
    m := &Office365ActiveUserDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365ActiveUserDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365ActiveUserDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365ActiveUserDetail(), nil
}
// GetAssignedProducts gets the assignedProducts property value. All the products assigned for the user.
func (m *Office365ActiveUserDetail) GetAssignedProducts()([]string) {
    val, err := m.GetBackingStore().Get("assignedProducts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDeletedDate gets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user hasn't been deleted.
func (m *Office365ActiveUserDetail) GetDeletedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("deletedDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it can't be cleared during updates.
func (m *Office365ActiveUserDetail) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExchangeLastActivityDate gets the exchangeLastActivityDate property value. The date when user last read or sent email.
func (m *Office365ActiveUserDetail) GetExchangeLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("exchangeLastActivityDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetExchangeLicenseAssignDate gets the exchangeLicenseAssignDate property value. The last date when the user was assigned an Exchange license.
func (m *Office365ActiveUserDetail) GetExchangeLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("exchangeLicenseAssignDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365ActiveUserDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedProducts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAssignedProducts(res)
        }
        return nil
    }
    res["deletedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDate(val)
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
    res["exchangeLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastActivityDate(val)
        }
        return nil
    }
    res["exchangeLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLicenseAssignDate(val)
        }
        return nil
    }
    res["hasExchangeLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasExchangeLicense(val)
        }
        return nil
    }
    res["hasOneDriveLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasOneDriveLicense(val)
        }
        return nil
    }
    res["hasSharePointLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasSharePointLicense(val)
        }
        return nil
    }
    res["hasSkypeForBusinessLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasSkypeForBusinessLicense(val)
        }
        return nil
    }
    res["hasTeamsLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasTeamsLicense(val)
        }
        return nil
    }
    res["hasYammerLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasYammerLicense(val)
        }
        return nil
    }
    res["isDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["oneDriveLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveLastActivityDate(val)
        }
        return nil
    }
    res["oneDriveLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveLicenseAssignDate(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharePointLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointLastActivityDate(val)
        }
        return nil
    }
    res["sharePointLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointLicenseAssignDate(val)
        }
        return nil
    }
    res["skypeForBusinessLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusinessLastActivityDate(val)
        }
        return nil
    }
    res["skypeForBusinessLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusinessLicenseAssignDate(val)
        }
        return nil
    }
    res["teamsLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsLastActivityDate(val)
        }
        return nil
    }
    res["teamsLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsLicenseAssignDate(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["yammerLastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerLastActivityDate(val)
        }
        return nil
    }
    res["yammerLicenseAssignDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerLicenseAssignDate(val)
        }
        return nil
    }
    return res
}
// GetHasExchangeLicense gets the hasExchangeLicense property value. Whether the user has been assigned an Exchange license.
func (m *Office365ActiveUserDetail) GetHasExchangeLicense()(*bool) {
    val, err := m.GetBackingStore().Get("hasExchangeLicense")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHasOneDriveLicense gets the hasOneDriveLicense property value. Whether the user has been assigned a OneDrive license.
func (m *Office365ActiveUserDetail) GetHasOneDriveLicense()(*bool) {
    val, err := m.GetBackingStore().Get("hasOneDriveLicense")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHasSharePointLicense gets the hasSharePointLicense property value. Whether the user has been assigned a SharePoint license.
func (m *Office365ActiveUserDetail) GetHasSharePointLicense()(*bool) {
    val, err := m.GetBackingStore().Get("hasSharePointLicense")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHasSkypeForBusinessLicense gets the hasSkypeForBusinessLicense property value. Whether the user has been assigned a Skype For Business license.
func (m *Office365ActiveUserDetail) GetHasSkypeForBusinessLicense()(*bool) {
    val, err := m.GetBackingStore().Get("hasSkypeForBusinessLicense")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHasTeamsLicense gets the hasTeamsLicense property value. Whether the user has been assigned a Teams license.
func (m *Office365ActiveUserDetail) GetHasTeamsLicense()(*bool) {
    val, err := m.GetBackingStore().Get("hasTeamsLicense")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHasYammerLicense gets the hasYammerLicense property value. Whether the user has been assigned a Yammer license.
func (m *Office365ActiveUserDetail) GetHasYammerLicense()(*bool) {
    val, err := m.GetBackingStore().Get("hasYammerLicense")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsDeleted gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365ActiveUserDetail) GetIsDeleted()(*bool) {
    val, err := m.GetBackingStore().Get("isDeleted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOneDriveLastActivityDate gets the oneDriveLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, or synced files.
func (m *Office365ActiveUserDetail) GetOneDriveLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("oneDriveLastActivityDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetOneDriveLicenseAssignDate gets the oneDriveLicenseAssignDate property value. The last date when the user was assigned a OneDrive license.
func (m *Office365ActiveUserDetail) GetOneDriveLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("oneDriveLicenseAssignDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365ActiveUserDetail) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("reportRefreshDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetSharePointLastActivityDate gets the sharePointLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages.
func (m *Office365ActiveUserDetail) GetSharePointLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("sharePointLastActivityDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetSharePointLicenseAssignDate gets the sharePointLicenseAssignDate property value. The last date when the user was assigned a SharePoint license.
func (m *Office365ActiveUserDetail) GetSharePointLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("sharePointLicenseAssignDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetSkypeForBusinessLastActivityDate gets the skypeForBusinessLastActivityDate property value. The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
func (m *Office365ActiveUserDetail) GetSkypeForBusinessLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("skypeForBusinessLastActivityDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetSkypeForBusinessLicenseAssignDate gets the skypeForBusinessLicenseAssignDate property value. The last date when the user was assigned a Skype For Business license.
func (m *Office365ActiveUserDetail) GetSkypeForBusinessLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("skypeForBusinessLicenseAssignDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetTeamsLastActivityDate gets the teamsLastActivityDate property value. The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls.
func (m *Office365ActiveUserDetail) GetTeamsLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("teamsLastActivityDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetTeamsLicenseAssignDate gets the teamsLicenseAssignDate property value. The last date when the user was assigned a Teams license.
func (m *Office365ActiveUserDetail) GetTeamsLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("teamsLicenseAssignDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
func (m *Office365ActiveUserDetail) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetYammerLastActivityDate gets the yammerLastActivityDate property value. The date when user last posted, read, or liked message.
func (m *Office365ActiveUserDetail) GetYammerLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("yammerLastActivityDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetYammerLicenseAssignDate gets the yammerLicenseAssignDate property value. The last date when the user was assigned a Yammer license.
func (m *Office365ActiveUserDetail) GetYammerLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("yammerLicenseAssignDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Office365ActiveUserDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedProducts() != nil {
        err = writer.WriteCollectionOfStringValues("assignedProducts", m.GetAssignedProducts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("deletedDate", m.GetDeletedDate())
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
        err = writer.WriteDateOnlyValue("exchangeLastActivityDate", m.GetExchangeLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("exchangeLicenseAssignDate", m.GetExchangeLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasExchangeLicense", m.GetHasExchangeLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasOneDriveLicense", m.GetHasOneDriveLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasSharePointLicense", m.GetHasSharePointLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasSkypeForBusinessLicense", m.GetHasSkypeForBusinessLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasTeamsLicense", m.GetHasTeamsLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasYammerLicense", m.GetHasYammerLicense())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("oneDriveLastActivityDate", m.GetOneDriveLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("oneDriveLicenseAssignDate", m.GetOneDriveLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("sharePointLastActivityDate", m.GetSharePointLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("sharePointLicenseAssignDate", m.GetSharePointLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("skypeForBusinessLastActivityDate", m.GetSkypeForBusinessLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("skypeForBusinessLicenseAssignDate", m.GetSkypeForBusinessLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("teamsLastActivityDate", m.GetTeamsLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("teamsLicenseAssignDate", m.GetTeamsLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("yammerLastActivityDate", m.GetYammerLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("yammerLicenseAssignDate", m.GetYammerLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedProducts sets the assignedProducts property value. All the products assigned for the user.
func (m *Office365ActiveUserDetail) SetAssignedProducts(value []string)() {
    err := m.GetBackingStore().Set("assignedProducts", value)
    if err != nil {
        panic(err)
    }
}
// SetDeletedDate sets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user hasn't been deleted.
func (m *Office365ActiveUserDetail) SetDeletedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("deletedDate", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it can't be cleared during updates.
func (m *Office365ActiveUserDetail) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeLastActivityDate sets the exchangeLastActivityDate property value. The date when user last read or sent email.
func (m *Office365ActiveUserDetail) SetExchangeLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("exchangeLastActivityDate", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeLicenseAssignDate sets the exchangeLicenseAssignDate property value. The last date when the user was assigned an Exchange license.
func (m *Office365ActiveUserDetail) SetExchangeLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("exchangeLicenseAssignDate", value)
    if err != nil {
        panic(err)
    }
}
// SetHasExchangeLicense sets the hasExchangeLicense property value. Whether the user has been assigned an Exchange license.
func (m *Office365ActiveUserDetail) SetHasExchangeLicense(value *bool)() {
    err := m.GetBackingStore().Set("hasExchangeLicense", value)
    if err != nil {
        panic(err)
    }
}
// SetHasOneDriveLicense sets the hasOneDriveLicense property value. Whether the user has been assigned a OneDrive license.
func (m *Office365ActiveUserDetail) SetHasOneDriveLicense(value *bool)() {
    err := m.GetBackingStore().Set("hasOneDriveLicense", value)
    if err != nil {
        panic(err)
    }
}
// SetHasSharePointLicense sets the hasSharePointLicense property value. Whether the user has been assigned a SharePoint license.
func (m *Office365ActiveUserDetail) SetHasSharePointLicense(value *bool)() {
    err := m.GetBackingStore().Set("hasSharePointLicense", value)
    if err != nil {
        panic(err)
    }
}
// SetHasSkypeForBusinessLicense sets the hasSkypeForBusinessLicense property value. Whether the user has been assigned a Skype For Business license.
func (m *Office365ActiveUserDetail) SetHasSkypeForBusinessLicense(value *bool)() {
    err := m.GetBackingStore().Set("hasSkypeForBusinessLicense", value)
    if err != nil {
        panic(err)
    }
}
// SetHasTeamsLicense sets the hasTeamsLicense property value. Whether the user has been assigned a Teams license.
func (m *Office365ActiveUserDetail) SetHasTeamsLicense(value *bool)() {
    err := m.GetBackingStore().Set("hasTeamsLicense", value)
    if err != nil {
        panic(err)
    }
}
// SetHasYammerLicense sets the hasYammerLicense property value. Whether the user has been assigned a Yammer license.
func (m *Office365ActiveUserDetail) SetHasYammerLicense(value *bool)() {
    err := m.GetBackingStore().Set("hasYammerLicense", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDeleted sets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365ActiveUserDetail) SetIsDeleted(value *bool)() {
    err := m.GetBackingStore().Set("isDeleted", value)
    if err != nil {
        panic(err)
    }
}
// SetOneDriveLastActivityDate sets the oneDriveLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, or synced files.
func (m *Office365ActiveUserDetail) SetOneDriveLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("oneDriveLastActivityDate", value)
    if err != nil {
        panic(err)
    }
}
// SetOneDriveLicenseAssignDate sets the oneDriveLicenseAssignDate property value. The last date when the user was assigned a OneDrive license.
func (m *Office365ActiveUserDetail) SetOneDriveLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("oneDriveLicenseAssignDate", value)
    if err != nil {
        panic(err)
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365ActiveUserDetail) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("reportRefreshDate", value)
    if err != nil {
        panic(err)
    }
}
// SetSharePointLastActivityDate sets the sharePointLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages.
func (m *Office365ActiveUserDetail) SetSharePointLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("sharePointLastActivityDate", value)
    if err != nil {
        panic(err)
    }
}
// SetSharePointLicenseAssignDate sets the sharePointLicenseAssignDate property value. The last date when the user was assigned a SharePoint license.
func (m *Office365ActiveUserDetail) SetSharePointLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("sharePointLicenseAssignDate", value)
    if err != nil {
        panic(err)
    }
}
// SetSkypeForBusinessLastActivityDate sets the skypeForBusinessLastActivityDate property value. The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
func (m *Office365ActiveUserDetail) SetSkypeForBusinessLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("skypeForBusinessLastActivityDate", value)
    if err != nil {
        panic(err)
    }
}
// SetSkypeForBusinessLicenseAssignDate sets the skypeForBusinessLicenseAssignDate property value. The last date when the user was assigned a Skype For Business license.
func (m *Office365ActiveUserDetail) SetSkypeForBusinessLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("skypeForBusinessLicenseAssignDate", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsLastActivityDate sets the teamsLastActivityDate property value. The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls.
func (m *Office365ActiveUserDetail) SetTeamsLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("teamsLastActivityDate", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsLicenseAssignDate sets the teamsLicenseAssignDate property value. The last date when the user was assigned a Teams license.
func (m *Office365ActiveUserDetail) SetTeamsLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("teamsLicenseAssignDate", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
func (m *Office365ActiveUserDetail) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetYammerLastActivityDate sets the yammerLastActivityDate property value. The date when user last posted, read, or liked message.
func (m *Office365ActiveUserDetail) SetYammerLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("yammerLastActivityDate", value)
    if err != nil {
        panic(err)
    }
}
// SetYammerLicenseAssignDate sets the yammerLicenseAssignDate property value. The last date when the user was assigned a Yammer license.
func (m *Office365ActiveUserDetail) SetYammerLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("yammerLicenseAssignDate", value)
    if err != nil {
        panic(err)
    }
}
// Office365ActiveUserDetailable 
type Office365ActiveUserDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedProducts()([]string)
    GetDeletedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetDisplayName()(*string)
    GetExchangeLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetExchangeLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetHasExchangeLicense()(*bool)
    GetHasOneDriveLicense()(*bool)
    GetHasSharePointLicense()(*bool)
    GetHasSkypeForBusinessLicense()(*bool)
    GetHasTeamsLicense()(*bool)
    GetHasYammerLicense()(*bool)
    GetIsDeleted()(*bool)
    GetOneDriveLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetOneDriveLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSharePointLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSharePointLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSkypeForBusinessLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSkypeForBusinessLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTeamsLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTeamsLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetUserPrincipalName()(*string)
    GetYammerLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetYammerLicenseAssignDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetAssignedProducts(value []string)()
    SetDeletedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetDisplayName(value *string)()
    SetExchangeLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetExchangeLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetHasExchangeLicense(value *bool)()
    SetHasOneDriveLicense(value *bool)()
    SetHasSharePointLicense(value *bool)()
    SetHasSkypeForBusinessLicense(value *bool)()
    SetHasTeamsLicense(value *bool)()
    SetHasYammerLicense(value *bool)()
    SetIsDeleted(value *bool)()
    SetOneDriveLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetOneDriveLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSharePointLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSharePointLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSkypeForBusinessLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSkypeForBusinessLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTeamsLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTeamsLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetUserPrincipalName(value *string)()
    SetYammerLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetYammerLicenseAssignDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
