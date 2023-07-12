package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsVppApp contains properties and inherited properties for MacOS Volume-Purchased Program (VPP) Apps.
type MacOsVppApp struct {
    MobileApp
    // The OdataType property
    OdataType *string
}
// NewMacOsVppApp instantiates a new macOsVppApp and sets the default values.
func NewMacOsVppApp()(*MacOsVppApp) {
    m := &MacOsVppApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.macOsVppApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOsVppAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOsVppAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOsVppApp(), nil
}
// GetAppStoreUrl gets the appStoreUrl property value. The store URL.
func (m *MacOsVppApp) GetAppStoreUrl()(*string) {
    val, err := m.GetBackingStore().Get("appStoreUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses assigned to this app.
func (m *MacOsVppApp) GetAssignedLicenses()([]MacOsVppAppAssignedLicenseable) {
    val, err := m.GetBackingStore().Get("assignedLicenses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOsVppAppAssignedLicenseable)
    }
    return nil
}
// GetBundleId gets the bundleId property value. The Identity Name.
func (m *MacOsVppApp) GetBundleId()(*string) {
    val, err := m.GetBackingStore().Get("bundleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOsVppApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreUrl(val)
        }
        return nil
    }
    res["assignedLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOsVppAppAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOsVppAppAssignedLicenseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOsVppAppAssignedLicenseable)
                }
            }
            m.SetAssignedLicenses(res)
        }
        return nil
    }
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
        }
        return nil
    }
    res["licensingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVppLicensingTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicensingType(val.(VppLicensingTypeable))
        }
        return nil
    }
    res["releaseDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    res["revokeLicenseActionResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOsVppAppRevokeLicensesActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOsVppAppRevokeLicensesActionResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOsVppAppRevokeLicensesActionResultable)
                }
            }
            m.SetRevokeLicenseActionResults(res)
        }
        return nil
    }
    res["totalLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicenseCount(val)
        }
        return nil
    }
    res["usedLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedLicenseCount(val)
        }
        return nil
    }
    res["vppTokenAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVppTokenAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenAccountType(val.(*VppTokenAccountType))
        }
        return nil
    }
    res["vppTokenAppleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenAppleId(val)
        }
        return nil
    }
    res["vppTokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenId(val)
        }
        return nil
    }
    res["vppTokenOrganizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenOrganizationName(val)
        }
        return nil
    }
    return res
}
// GetLicensingType gets the licensingType property value. The supported License Type.
func (m *MacOsVppApp) GetLicensingType()(VppLicensingTypeable) {
    val, err := m.GetBackingStore().Get("licensingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VppLicensingTypeable)
    }
    return nil
}
// GetReleaseDateTime gets the releaseDateTime property value. The VPP application release date and time.
func (m *MacOsVppApp) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("releaseDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRevokeLicenseActionResults gets the revokeLicenseActionResults property value. Results of revoke license actions on this app.
func (m *MacOsVppApp) GetRevokeLicenseActionResults()([]MacOsVppAppRevokeLicensesActionResultable) {
    val, err := m.GetBackingStore().Get("revokeLicenseActionResults")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOsVppAppRevokeLicensesActionResultable)
    }
    return nil
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of VPP licenses.
func (m *MacOsVppApp) GetTotalLicenseCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalLicenseCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *MacOsVppApp) GetUsedLicenseCount()(*int32) {
    val, err := m.GetBackingStore().Get("usedLicenseCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetVppTokenAccountType gets the vppTokenAccountType property value. Possible types of an Apple Volume Purchase Program token.
func (m *MacOsVppApp) GetVppTokenAccountType()(*VppTokenAccountType) {
    val, err := m.GetBackingStore().Get("vppTokenAccountType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VppTokenAccountType)
    }
    return nil
}
// GetVppTokenAppleId gets the vppTokenAppleId property value. The Apple Id associated with the given Apple Volume Purchase Program Token.
func (m *MacOsVppApp) GetVppTokenAppleId()(*string) {
    val, err := m.GetBackingStore().Get("vppTokenAppleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVppTokenId gets the vppTokenId property value. Identifier of the VPP token associated with this app.
func (m *MacOsVppApp) GetVppTokenId()(*string) {
    val, err := m.GetBackingStore().Get("vppTokenId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVppTokenOrganizationName gets the vppTokenOrganizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *MacOsVppApp) GetVppTokenOrganizationName()(*string) {
    val, err := m.GetBackingStore().Get("vppTokenOrganizationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOsVppApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    if m.GetAssignedLicenses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedLicenses()))
        for i, v := range m.GetAssignedLicenses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bundleId", m.GetBundleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("licensingType", m.GetLicensingType())
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
    if m.GetRevokeLicenseActionResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRevokeLicenseActionResults()))
        for i, v := range m.GetRevokeLicenseActionResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("revokeLicenseActionResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicenseCount", m.GetTotalLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usedLicenseCount", m.GetUsedLicenseCount())
        if err != nil {
            return err
        }
    }
    if m.GetVppTokenAccountType() != nil {
        cast := (*m.GetVppTokenAccountType()).String()
        err = writer.WriteStringValue("vppTokenAccountType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vppTokenAppleId", m.GetVppTokenAppleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vppTokenId", m.GetVppTokenId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vppTokenOrganizationName", m.GetVppTokenOrganizationName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppStoreUrl sets the appStoreUrl property value. The store URL.
func (m *MacOsVppApp) SetAppStoreUrl(value *string)() {
    err := m.GetBackingStore().Set("appStoreUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses assigned to this app.
func (m *MacOsVppApp) SetAssignedLicenses(value []MacOsVppAppAssignedLicenseable)() {
    err := m.GetBackingStore().Set("assignedLicenses", value)
    if err != nil {
        panic(err)
    }
}
// SetBundleId sets the bundleId property value. The Identity Name.
func (m *MacOsVppApp) SetBundleId(value *string)() {
    err := m.GetBackingStore().Set("bundleId", value)
    if err != nil {
        panic(err)
    }
}
// SetLicensingType sets the licensingType property value. The supported License Type.
func (m *MacOsVppApp) SetLicensingType(value VppLicensingTypeable)() {
    err := m.GetBackingStore().Set("licensingType", value)
    if err != nil {
        panic(err)
    }
}
// SetReleaseDateTime sets the releaseDateTime property value. The VPP application release date and time.
func (m *MacOsVppApp) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("releaseDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRevokeLicenseActionResults sets the revokeLicenseActionResults property value. Results of revoke license actions on this app.
func (m *MacOsVppApp) SetRevokeLicenseActionResults(value []MacOsVppAppRevokeLicensesActionResultable)() {
    err := m.GetBackingStore().Set("revokeLicenseActionResults", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of VPP licenses.
func (m *MacOsVppApp) SetTotalLicenseCount(value *int32)() {
    err := m.GetBackingStore().Set("totalLicenseCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *MacOsVppApp) SetUsedLicenseCount(value *int32)() {
    err := m.GetBackingStore().Set("usedLicenseCount", value)
    if err != nil {
        panic(err)
    }
}
// SetVppTokenAccountType sets the vppTokenAccountType property value. Possible types of an Apple Volume Purchase Program token.
func (m *MacOsVppApp) SetVppTokenAccountType(value *VppTokenAccountType)() {
    err := m.GetBackingStore().Set("vppTokenAccountType", value)
    if err != nil {
        panic(err)
    }
}
// SetVppTokenAppleId sets the vppTokenAppleId property value. The Apple Id associated with the given Apple Volume Purchase Program Token.
func (m *MacOsVppApp) SetVppTokenAppleId(value *string)() {
    err := m.GetBackingStore().Set("vppTokenAppleId", value)
    if err != nil {
        panic(err)
    }
}
// SetVppTokenId sets the vppTokenId property value. Identifier of the VPP token associated with this app.
func (m *MacOsVppApp) SetVppTokenId(value *string)() {
    err := m.GetBackingStore().Set("vppTokenId", value)
    if err != nil {
        panic(err)
    }
}
// SetVppTokenOrganizationName sets the vppTokenOrganizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *MacOsVppApp) SetVppTokenOrganizationName(value *string)() {
    err := m.GetBackingStore().Set("vppTokenOrganizationName", value)
    if err != nil {
        panic(err)
    }
}
// MacOsVppAppable 
type MacOsVppAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppStoreUrl()(*string)
    GetAssignedLicenses()([]MacOsVppAppAssignedLicenseable)
    GetBundleId()(*string)
    GetLicensingType()(VppLicensingTypeable)
    GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRevokeLicenseActionResults()([]MacOsVppAppRevokeLicensesActionResultable)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    GetVppTokenAccountType()(*VppTokenAccountType)
    GetVppTokenAppleId()(*string)
    GetVppTokenId()(*string)
    GetVppTokenOrganizationName()(*string)
    SetAppStoreUrl(value *string)()
    SetAssignedLicenses(value []MacOsVppAppAssignedLicenseable)()
    SetBundleId(value *string)()
    SetLicensingType(value VppLicensingTypeable)()
    SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRevokeLicenseActionResults(value []MacOsVppAppRevokeLicensesActionResultable)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
    SetVppTokenAccountType(value *VppTokenAccountType)()
    SetVppTokenAppleId(value *string)()
    SetVppTokenId(value *string)()
    SetVppTokenOrganizationName(value *string)()
}
