package getoffice365activeuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetOffice365ActiveUserDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    assignedProducts []string;
    deletedDate *string;
    displayName *string;
    exchangeLastActivityDate *string;
    exchangeLicenseAssignDate *string;
    hasExchangeLicense *bool;
    hasOneDriveLicense *bool;
    hasSharePointLicense *bool;
    hasSkypeForBusinessLicense *bool;
    hasTeamsLicense *bool;
    hasYammerLicense *bool;
    isDeleted *bool;
    oneDriveLastActivityDate *string;
    oneDriveLicenseAssignDate *string;
    reportRefreshDate *string;
    sharePointLastActivityDate *string;
    sharePointLicenseAssignDate *string;
    skypeForBusinessLastActivityDate *string;
    skypeForBusinessLicenseAssignDate *string;
    teamsLastActivityDate *string;
    teamsLicenseAssignDate *string;
    userPrincipalName *string;
    yammerLastActivityDate *string;
    yammerLicenseAssignDate *string;
}
func NewGetOffice365ActiveUserDetailWithPeriod()(*GetOffice365ActiveUserDetailWithPeriod) {
    m := &GetOffice365ActiveUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetExchangeLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeLastActivityDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetExchangeLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeLicenseAssignDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasExchangeLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasExchangeLicense
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasOneDriveLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasOneDriveLicense
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasSharePointLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasSharePointLicense
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasSkypeForBusinessLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasSkypeForBusinessLicense
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasTeamsLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasTeamsLicense
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetHasYammerLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasYammerLicense
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetOneDriveLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveLastActivityDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetOneDriveLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveLicenseAssignDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetSharePointLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sharePointLastActivityDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetSharePointLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sharePointLicenseAssignDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetSkypeForBusinessLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessLastActivityDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetSkypeForBusinessLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessLicenseAssignDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetTeamsLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsLastActivityDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetTeamsLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsLicenseAssignDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetYammerLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yammerLastActivityDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetYammerLicenseAssignDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yammerLicenseAssignDate
    }
}
func (m *GetOffice365ActiveUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedProducts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssignedProducts(res)
        return nil
    }
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeletedDate(val)
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
    res["exchangeLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExchangeLastActivityDate(val)
        return nil
    }
    res["exchangeLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExchangeLicenseAssignDate(val)
        return nil
    }
    res["hasExchangeLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasExchangeLicense(val)
        return nil
    }
    res["hasOneDriveLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasOneDriveLicense(val)
        return nil
    }
    res["hasSharePointLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasSharePointLicense(val)
        return nil
    }
    res["hasSkypeForBusinessLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasSkypeForBusinessLicense(val)
        return nil
    }
    res["hasTeamsLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasTeamsLicense(val)
        return nil
    }
    res["hasYammerLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasYammerLicense(val)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["oneDriveLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOneDriveLastActivityDate(val)
        return nil
    }
    res["oneDriveLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOneDriveLicenseAssignDate(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["sharePointLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSharePointLastActivityDate(val)
        return nil
    }
    res["sharePointLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSharePointLicenseAssignDate(val)
        return nil
    }
    res["skypeForBusinessLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSkypeForBusinessLastActivityDate(val)
        return nil
    }
    res["skypeForBusinessLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSkypeForBusinessLicenseAssignDate(val)
        return nil
    }
    res["teamsLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTeamsLastActivityDate(val)
        return nil
    }
    res["teamsLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTeamsLicenseAssignDate(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["yammerLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetYammerLastActivityDate(val)
        return nil
    }
    res["yammerLicenseAssignDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetYammerLicenseAssignDate(val)
        return nil
    }
    return res
}
func (m *GetOffice365ActiveUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
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
func (m *GetOffice365ActiveUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetExchangeLastActivityDate(value *string)() {
    m.exchangeLastActivityDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetExchangeLicenseAssignDate(value *string)() {
    m.exchangeLicenseAssignDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasExchangeLicense(value *bool)() {
    m.hasExchangeLicense = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasOneDriveLicense(value *bool)() {
    m.hasOneDriveLicense = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasSharePointLicense(value *bool)() {
    m.hasSharePointLicense = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasSkypeForBusinessLicense(value *bool)() {
    m.hasSkypeForBusinessLicense = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasTeamsLicense(value *bool)() {
    m.hasTeamsLicense = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetHasYammerLicense(value *bool)() {
    m.hasYammerLicense = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetOneDriveLastActivityDate(value *string)() {
    m.oneDriveLastActivityDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetOneDriveLicenseAssignDate(value *string)() {
    m.oneDriveLicenseAssignDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetSharePointLastActivityDate(value *string)() {
    m.sharePointLastActivityDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetSharePointLicenseAssignDate(value *string)() {
    m.sharePointLicenseAssignDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetSkypeForBusinessLastActivityDate(value *string)() {
    m.skypeForBusinessLastActivityDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetSkypeForBusinessLicenseAssignDate(value *string)() {
    m.skypeForBusinessLicenseAssignDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetTeamsLastActivityDate(value *string)() {
    m.teamsLastActivityDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetTeamsLicenseAssignDate(value *string)() {
    m.teamsLicenseAssignDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetYammerLastActivityDate(value *string)() {
    m.yammerLastActivityDate = value
}
func (m *GetOffice365ActiveUserDetailWithPeriod) SetYammerLicenseAssignDate(value *string)() {
    m.yammerLicenseAssignDate = value
}
