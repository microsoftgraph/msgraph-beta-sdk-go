package getoffice365activeuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetOffice365ActiveUserDetailWithPeriod 
type GetOffice365ActiveUserDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // All the products assigned for the user.
    assignedProducts []string;
    // The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
    deletedDate *string;
    // The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates.
    displayName *string;
    // The date when user last read or sent email.
    exchangeLastActivityDate *string;
    // The last date when the user was assigned an Exchange license.
    exchangeLicenseAssignDate *string;
    // Whether the user has been assigned an Exchange license.
    hasExchangeLicense *bool;
    // Whether the user has been assigned a OneDrive license.
    hasOneDriveLicense *bool;
    // Whether the user has been assigned a SharePoint license.
    hasSharePointLicense *bool;
    // Whether the user has been assigned a Skype For Business license.
    hasSkypeForBusinessLicense *bool;
    // Whether the user has been assigned a Teams license.
    hasTeamsLicense *bool;
    // Whether the user has been assigned a Yammer license.
    hasYammerLicense *bool;
    // Whether this user has been deleted or soft deleted.
    isDeleted *bool;
    // The date when user last viewed or edited files, shared files internally or externally, or synced files.
    oneDriveLastActivityDate *string;
    // The last date when the user was assigned a OneDrive license.
    oneDriveLicenseAssignDate *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages.
    sharePointLastActivityDate *string;
    // The last date when the user was assigned a SharePoint license.
    sharePointLicenseAssignDate *string;
    // The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
    skypeForBusinessLastActivityDate *string;
    // The last date when the user was assigned a Skype For Business license.
    skypeForBusinessLicenseAssignDate *string;
    // The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls.
    teamsLastActivityDate *string;
    // The last date when the user was assigned a Teams license.
    teamsLicenseAssignDate *string;
    // The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
    userPrincipalName *string;
    // The date when user last posted, read, or liked message.
    yammerLastActivityDate *string;
    // The last date when the user was assigned a Yammer license.
    yammerLicenseAssignDate *string;
}
// NewGetOffice365ActiveUserDetailWithPeriod instantiates a new getOffice365ActiveUserDetailWithPeriod and sets the default values.
func NewGetOffice365ActiveUserDetailWithPeriod()(*GetOffice365ActiveUserDetailWithPeriod) {
    m := &GetOffice365ActiveUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAssignedProducts gets the assignedProducts property value. All the products assigned for the user.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// GetDeletedDate gets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// GetDisplayName gets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExchangeLastActivityDate gets the exchangeLastActivityDate property value. The date when user last read or sent email.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetExchangeLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeLastActivityDate
    }
}
// GetExchangeLicenseAssignDate gets the exchangeLicenseAssignDate property value. The last date when the user was assigned an Exchange license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetExchangeLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeLicenseAssignDate
    }
}
// GetHasExchangeLicense gets the hasExchangeLicense property value. Whether the user has been assigned an Exchange license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasExchangeLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasExchangeLicense
    }
}
// GetHasOneDriveLicense gets the hasOneDriveLicense property value. Whether the user has been assigned a OneDrive license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasOneDriveLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasOneDriveLicense
    }
}
// GetHasSharePointLicense gets the hasSharePointLicense property value. Whether the user has been assigned a SharePoint license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasSharePointLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasSharePointLicense
    }
}
// GetHasSkypeForBusinessLicense gets the hasSkypeForBusinessLicense property value. Whether the user has been assigned a Skype For Business license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasSkypeForBusinessLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasSkypeForBusinessLicense
    }
}
// GetHasTeamsLicense gets the hasTeamsLicense property value. Whether the user has been assigned a Teams license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasTeamsLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasTeamsLicense
    }
}
// GetHasYammerLicense gets the hasYammerLicense property value. Whether the user has been assigned a Yammer license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasYammerLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasYammerLicense
    }
}
// GetIsDeleted gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetOneDriveLastActivityDate gets the oneDriveLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, or synced files.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetOneDriveLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveLastActivityDate
    }
}
// GetOneDriveLicenseAssignDate gets the oneDriveLicenseAssignDate property value. The last date when the user was assigned a OneDrive license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetOneDriveLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveLicenseAssignDate
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSharePointLastActivityDate gets the sharePointLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetSharePointLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sharePointLastActivityDate
    }
}
// GetSharePointLicenseAssignDate gets the sharePointLicenseAssignDate property value. The last date when the user was assigned a SharePoint license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetSharePointLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sharePointLicenseAssignDate
    }
}
// GetSkypeForBusinessLastActivityDate gets the skypeForBusinessLastActivityDate property value. The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetSkypeForBusinessLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessLastActivityDate
    }
}
// GetSkypeForBusinessLicenseAssignDate gets the skypeForBusinessLicenseAssignDate property value. The last date when the user was assigned a Skype For Business license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetSkypeForBusinessLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessLicenseAssignDate
    }
}
// GetTeamsLastActivityDate gets the teamsLastActivityDate property value. The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetTeamsLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsLastActivityDate
    }
}
// GetTeamsLicenseAssignDate gets the teamsLicenseAssignDate property value. The last date when the user was assigned a Teams license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetTeamsLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsLicenseAssignDate
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetYammerLastActivityDate gets the yammerLastActivityDate property value. The date when user last posted, read, or liked message.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetYammerLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yammerLastActivityDate
    }
}
// GetYammerLicenseAssignDate gets the yammerLicenseAssignDate property value. The last date when the user was assigned a Yammer license.
func (m *GetOffice365ActiveUserDetailWithPeriod) GetYammerLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yammerLicenseAssignDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOffice365ActiveUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedProducts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignedProducts(res)
        }
        return nil
    }
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDate(val)
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
    res["exchangeLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastActivityDate(val)
        }
        return nil
    }
    res["exchangeLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLicenseAssignDate(val)
        }
        return nil
    }
    res["hasExchangeLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasExchangeLicense(val)
        }
        return nil
    }
    res["hasOneDriveLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasOneDriveLicense(val)
        }
        return nil
    }
    res["hasSharePointLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasSharePointLicense(val)
        }
        return nil
    }
    res["hasSkypeForBusinessLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasSkypeForBusinessLicense(val)
        }
        return nil
    }
    res["hasTeamsLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasTeamsLicense(val)
        }
        return nil
    }
    res["hasYammerLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasYammerLicense(val)
        }
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["oneDriveLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveLastActivityDate(val)
        }
        return nil
    }
    res["oneDriveLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveLicenseAssignDate(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharePointLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointLastActivityDate(val)
        }
        return nil
    }
    res["sharePointLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointLicenseAssignDate(val)
        }
        return nil
    }
    res["skypeForBusinessLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusinessLastActivityDate(val)
        }
        return nil
    }
    res["skypeForBusinessLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusinessLicenseAssignDate(val)
        }
        return nil
    }
    res["teamsLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsLastActivityDate(val)
        }
        return nil
    }
    res["teamsLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsLicenseAssignDate(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["yammerLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerLastActivityDate(val)
        }
        return nil
    }
    res["yammerLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
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
func (m *GetOffice365ActiveUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetOffice365ActiveUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("assignedProducts", m.GetAssignedProducts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deletedDate", m.GetDeletedDate())
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
        err = writer.WriteStringValue("exchangeLastActivityDate", m.GetExchangeLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("exchangeLicenseAssignDate", m.GetExchangeLicenseAssignDate())
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
        err = writer.WriteStringValue("oneDriveLastActivityDate", m.GetOneDriveLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("oneDriveLicenseAssignDate", m.GetOneDriveLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sharePointLastActivityDate", m.GetSharePointLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sharePointLicenseAssignDate", m.GetSharePointLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skypeForBusinessLastActivityDate", m.GetSkypeForBusinessLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skypeForBusinessLicenseAssignDate", m.GetSkypeForBusinessLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsLastActivityDate", m.GetTeamsLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsLicenseAssignDate", m.GetTeamsLicenseAssignDate())
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
        err = writer.WriteStringValue("yammerLastActivityDate", m.GetYammerLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yammerLicenseAssignDate", m.GetYammerLicenseAssignDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedProducts sets the assignedProducts property value. All the products assigned for the user.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    if m != nil {
        m.assignedProducts = value
    }
}
// SetDeletedDate sets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetDeletedDate(value *string)() {
    if m != nil {
        m.deletedDate = value
    }
}
// SetDisplayName sets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExchangeLastActivityDate sets the exchangeLastActivityDate property value. The date when user last read or sent email.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetExchangeLastActivityDate(value *string)() {
    if m != nil {
        m.exchangeLastActivityDate = value
    }
}
// SetExchangeLicenseAssignDate sets the exchangeLicenseAssignDate property value. The last date when the user was assigned an Exchange license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetExchangeLicenseAssignDate(value *string)() {
    if m != nil {
        m.exchangeLicenseAssignDate = value
    }
}
// SetHasExchangeLicense sets the hasExchangeLicense property value. Whether the user has been assigned an Exchange license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasExchangeLicense(value *bool)() {
    if m != nil {
        m.hasExchangeLicense = value
    }
}
// SetHasOneDriveLicense sets the hasOneDriveLicense property value. Whether the user has been assigned a OneDrive license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasOneDriveLicense(value *bool)() {
    if m != nil {
        m.hasOneDriveLicense = value
    }
}
// SetHasSharePointLicense sets the hasSharePointLicense property value. Whether the user has been assigned a SharePoint license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasSharePointLicense(value *bool)() {
    if m != nil {
        m.hasSharePointLicense = value
    }
}
// SetHasSkypeForBusinessLicense sets the hasSkypeForBusinessLicense property value. Whether the user has been assigned a Skype For Business license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasSkypeForBusinessLicense(value *bool)() {
    if m != nil {
        m.hasSkypeForBusinessLicense = value
    }
}
// SetHasTeamsLicense sets the hasTeamsLicense property value. Whether the user has been assigned a Teams license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasTeamsLicense(value *bool)() {
    if m != nil {
        m.hasTeamsLicense = value
    }
}
// SetHasYammerLicense sets the hasYammerLicense property value. Whether the user has been assigned a Yammer license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasYammerLicense(value *bool)() {
    if m != nil {
        m.hasYammerLicense = value
    }
}
// SetIsDeleted sets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    if m != nil {
        m.isDeleted = value
    }
}
// SetOneDriveLastActivityDate sets the oneDriveLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, or synced files.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetOneDriveLastActivityDate(value *string)() {
    if m != nil {
        m.oneDriveLastActivityDate = value
    }
}
// SetOneDriveLicenseAssignDate sets the oneDriveLicenseAssignDate property value. The last date when the user was assigned a OneDrive license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetOneDriveLicenseAssignDate(value *string)() {
    if m != nil {
        m.oneDriveLicenseAssignDate = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetSharePointLastActivityDate sets the sharePointLastActivityDate property value. The date when user last viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetSharePointLastActivityDate(value *string)() {
    if m != nil {
        m.sharePointLastActivityDate = value
    }
}
// SetSharePointLicenseAssignDate sets the sharePointLicenseAssignDate property value. The last date when the user was assigned a SharePoint license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetSharePointLicenseAssignDate(value *string)() {
    if m != nil {
        m.sharePointLicenseAssignDate = value
    }
}
// SetSkypeForBusinessLastActivityDate sets the skypeForBusinessLastActivityDate property value. The date when user last organized or participated in conferences, or joined peer-to-peer sessions.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetSkypeForBusinessLastActivityDate(value *string)() {
    if m != nil {
        m.skypeForBusinessLastActivityDate = value
    }
}
// SetSkypeForBusinessLicenseAssignDate sets the skypeForBusinessLicenseAssignDate property value. The last date when the user was assigned a Skype For Business license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetSkypeForBusinessLicenseAssignDate(value *string)() {
    if m != nil {
        m.skypeForBusinessLicenseAssignDate = value
    }
}
// SetTeamsLastActivityDate sets the teamsLastActivityDate property value. The date when user last posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetTeamsLastActivityDate(value *string)() {
    if m != nil {
        m.teamsLastActivityDate = value
    }
}
// SetTeamsLicenseAssignDate sets the teamsLicenseAssignDate property value. The last date when the user was assigned a Teams license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetTeamsLicenseAssignDate(value *string)() {
    if m != nil {
        m.teamsLicenseAssignDate = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetYammerLastActivityDate sets the yammerLastActivityDate property value. The date when user last posted, read, or liked message.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetYammerLastActivityDate(value *string)() {
    if m != nil {
        m.yammerLastActivityDate = value
    }
}
// SetYammerLicenseAssignDate sets the yammerLicenseAssignDate property value. The last date when the user was assigned a Yammer license.
func (m *GetOffice365ActiveUserDetailWithPeriod) SetYammerLicenseAssignDate(value *string)() {
    if m != nil {
        m.yammerLicenseAssignDate = value
    }
}
