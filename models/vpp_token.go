package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppToken you purchase multiple licenses for iOS apps through the Apple Volume Purchase Program for Business or Education. This involves setting up an Apple VPP account from the Apple website and uploading the Apple VPP Business or Education token to Intune. You can then synchronize your volume purchase information with Intune and track your volume-purchased app use. You can upload multiple Apple VPP Business or Education tokens.
type VppToken struct {
    Entity
    // The apple Id associated with the given Apple Volume Purchase Program Token.
    appleId *string
    // Whether or not apps for the VPP token will be automatically updated.
    automaticallyUpdateApps *bool
    // Admin consent to allow claiming token management from external MDM.
    claimTokenManagementFromExternalMdm *bool
    // Whether or not apps for the VPP token will be automatically updated.
    countryOrRegion *string
    // Consent granted for data sharing with the Apple Volume Purchase Program.
    dataSharingConsentGranted *bool
    // An admin specified token friendly name.
    displayName *string
    // The expiration date time of the Apple Volume Purchase Program Token.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last modification date time associated with the Apple Volume Purchase Program Token.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Possible sync statuses associated with an Apple Volume Purchase Program token.
    lastSyncStatus *VppTokenSyncStatus
    // Token location returned from Apple VPP.
    locationName *string
    // The organization associated with the Apple Volume Purchase Program Token
    organizationName *string
    // Role Scope Tags IDs assigned to this entity.
    roleScopeTagIds []string
    // Possible states associated with an Apple Volume Purchase Program token.
    state *VppTokenState
    // The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
    token *string
    // The collection of statuses of the actions performed on the Apple Volume Purchase Program Token.
    tokenActionResults []VppTokenActionResultable
    // Possible types of an Apple Volume Purchase Program token.
    vppTokenAccountType *VppTokenAccountType
}
// NewVppToken instantiates a new vppToken and sets the default values.
func NewVppToken()(*VppToken) {
    m := &VppToken{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.vppToken";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateVppTokenFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVppTokenFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVppToken(), nil
}
// GetAppleId gets the appleId property value. The apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppToken) GetAppleId()(*string) {
    return m.appleId
}
// GetAutomaticallyUpdateApps gets the automaticallyUpdateApps property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) GetAutomaticallyUpdateApps()(*bool) {
    return m.automaticallyUpdateApps
}
// GetClaimTokenManagementFromExternalMdm gets the claimTokenManagementFromExternalMdm property value. Admin consent to allow claiming token management from external MDM.
func (m *VppToken) GetClaimTokenManagementFromExternalMdm()(*bool) {
    return m.claimTokenManagementFromExternalMdm
}
// GetCountryOrRegion gets the countryOrRegion property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) GetCountryOrRegion()(*string) {
    return m.countryOrRegion
}
// GetDataSharingConsentGranted gets the dataSharingConsentGranted property value. Consent granted for data sharing with the Apple Volume Purchase Program.
func (m *VppToken) GetDataSharingConsentGranted()(*bool) {
    return m.dataSharingConsentGranted
}
// GetDisplayName gets the displayName property value. An admin specified token friendly name.
func (m *VppToken) GetDisplayName()(*string) {
    return m.displayName
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration date time of the Apple Volume Purchase Program Token.
func (m *VppToken) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VppToken) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppleId)
    res["automaticallyUpdateApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAutomaticallyUpdateApps)
    res["claimTokenManagementFromExternalMdm"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetClaimTokenManagementFromExternalMdm)
    res["countryOrRegion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCountryOrRegion)
    res["dataSharingConsentGranted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDataSharingConsentGranted)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["expirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExpirationDateTime)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["lastSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastSyncDateTime)
    res["lastSyncStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVppTokenSyncStatus , m.SetLastSyncStatus)
    res["locationName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLocationName)
    res["organizationName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOrganizationName)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVppTokenState , m.SetState)
    res["token"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetToken)
    res["tokenActionResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateVppTokenActionResultFromDiscriminatorValue , m.SetTokenActionResults)
    res["vppTokenAccountType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVppTokenAccountType , m.SetVppTokenAccountType)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modification date time associated with the Apple Volume Purchase Program Token.
func (m *VppToken) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
func (m *VppToken) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetLastSyncStatus gets the lastSyncStatus property value. Possible sync statuses associated with an Apple Volume Purchase Program token.
func (m *VppToken) GetLastSyncStatus()(*VppTokenSyncStatus) {
    return m.lastSyncStatus
}
// GetLocationName gets the locationName property value. Token location returned from Apple VPP.
func (m *VppToken) GetLocationName()(*string) {
    return m.locationName
}
// GetOrganizationName gets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *VppToken) GetOrganizationName()(*string) {
    return m.organizationName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. Role Scope Tags IDs assigned to this entity.
func (m *VppToken) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetState gets the state property value. Possible states associated with an Apple Volume Purchase Program token.
func (m *VppToken) GetState()(*VppTokenState) {
    return m.state
}
// GetToken gets the token property value. The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
func (m *VppToken) GetToken()(*string) {
    return m.token
}
// GetTokenActionResults gets the tokenActionResults property value. The collection of statuses of the actions performed on the Apple Volume Purchase Program Token.
func (m *VppToken) GetTokenActionResults()([]VppTokenActionResultable) {
    return m.tokenActionResults
}
// GetVppTokenAccountType gets the vppTokenAccountType property value. Possible types of an Apple Volume Purchase Program token.
func (m *VppToken) GetVppTokenAccountType()(*VppTokenAccountType) {
    return m.vppTokenAccountType
}
// Serialize serializes information the current object
func (m *VppToken) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("automaticallyUpdateApps", m.GetAutomaticallyUpdateApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("claimTokenManagementFromExternalMdm", m.GetClaimTokenManagementFromExternalMdm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryOrRegion", m.GetCountryOrRegion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dataSharingConsentGranted", m.GetDataSharingConsentGranted())
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
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLastSyncStatus() != nil {
        cast := (*m.GetLastSyncStatus()).String()
        err = writer.WriteStringValue("lastSyncStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("locationName", m.GetLocationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationName", m.GetOrganizationName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("token", m.GetToken())
        if err != nil {
            return err
        }
    }
    if m.GetTokenActionResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTokenActionResults())
        err = writer.WriteCollectionOfObjectValues("tokenActionResults", cast)
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
    return nil
}
// SetAppleId sets the appleId property value. The apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppToken) SetAppleId(value *string)() {
    m.appleId = value
}
// SetAutomaticallyUpdateApps sets the automaticallyUpdateApps property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) SetAutomaticallyUpdateApps(value *bool)() {
    m.automaticallyUpdateApps = value
}
// SetClaimTokenManagementFromExternalMdm sets the claimTokenManagementFromExternalMdm property value. Admin consent to allow claiming token management from external MDM.
func (m *VppToken) SetClaimTokenManagementFromExternalMdm(value *bool)() {
    m.claimTokenManagementFromExternalMdm = value
}
// SetCountryOrRegion sets the countryOrRegion property value. Whether or not apps for the VPP token will be automatically updated.
func (m *VppToken) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// SetDataSharingConsentGranted sets the dataSharingConsentGranted property value. Consent granted for data sharing with the Apple Volume Purchase Program.
func (m *VppToken) SetDataSharingConsentGranted(value *bool)() {
    m.dataSharingConsentGranted = value
}
// SetDisplayName sets the displayName property value. An admin specified token friendly name.
func (m *VppToken) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration date time of the Apple Volume Purchase Program Token.
func (m *VppToken) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modification date time associated with the Apple Volume Purchase Program Token.
func (m *VppToken) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
func (m *VppToken) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetLastSyncStatus sets the lastSyncStatus property value. Possible sync statuses associated with an Apple Volume Purchase Program token.
func (m *VppToken) SetLastSyncStatus(value *VppTokenSyncStatus)() {
    m.lastSyncStatus = value
}
// SetLocationName sets the locationName property value. Token location returned from Apple VPP.
func (m *VppToken) SetLocationName(value *string)() {
    m.locationName = value
}
// SetOrganizationName sets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token
func (m *VppToken) SetOrganizationName(value *string)() {
    m.organizationName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. Role Scope Tags IDs assigned to this entity.
func (m *VppToken) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetState sets the state property value. Possible states associated with an Apple Volume Purchase Program token.
func (m *VppToken) SetState(value *VppTokenState)() {
    m.state = value
}
// SetToken sets the token property value. The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
func (m *VppToken) SetToken(value *string)() {
    m.token = value
}
// SetTokenActionResults sets the tokenActionResults property value. The collection of statuses of the actions performed on the Apple Volume Purchase Program Token.
func (m *VppToken) SetTokenActionResults(value []VppTokenActionResultable)() {
    m.tokenActionResults = value
}
// SetVppTokenAccountType sets the vppTokenAccountType property value. Possible types of an Apple Volume Purchase Program token.
func (m *VppToken) SetVppTokenAccountType(value *VppTokenAccountType)() {
    m.vppTokenAccountType = value
}
