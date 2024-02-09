package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccessPackageSubject struct {
    Entity
}
// NewAccessPackageSubject instantiates a new AccessPackageSubject and sets the default values.
func NewAccessPackageSubject()(*AccessPackageSubject) {
    m := &AccessPackageSubject{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageSubjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccessPackageSubjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageSubject(), nil
}
// GetAltSecId gets the altSecId property value. Not Supported.
// returns a *string when successful
func (m *AccessPackageSubject) GetAltSecId()(*string) {
    val, err := m.GetBackingStore().Get("altSecId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCleanupScheduledDateTime gets the cleanupScheduledDateTime property value. The date and time the subject is marked to be blocked from sign in or deleted. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.
// returns a *Time when successful
func (m *AccessPackageSubject) GetCleanupScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("cleanupScheduledDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetConnectedOrganization gets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
// returns a ConnectedOrganizationable when successful
func (m *AccessPackageSubject) GetConnectedOrganization()(ConnectedOrganizationable) {
    val, err := m.GetBackingStore().Get("connectedOrganization")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ConnectedOrganizationable)
    }
    return nil
}
// GetConnectedOrganizationId gets the connectedOrganizationId property value. The identifier of the connected organization of the subject.
// returns a *string when successful
func (m *AccessPackageSubject) GetConnectedOrganizationId()(*string) {
    val, err := m.GetBackingStore().Get("connectedOrganizationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the subject.
// returns a *string when successful
func (m *AccessPackageSubject) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmail gets the email property value. The email address of the subject.
// returns a *string when successful
func (m *AccessPackageSubject) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
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
func (m *AccessPackageSubject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["altSecId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltSecId(val)
        }
        return nil
    }
    res["cleanupScheduledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCleanupScheduledDateTime(val)
        }
        return nil
    }
    res["connectedOrganization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConnectedOrganizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedOrganization(val.(ConnectedOrganizationable))
        }
        return nil
    }
    res["connectedOrganizationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedOrganizationId(val)
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["objectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectId(val)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
        }
        return nil
    }
    res["principalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalName(val)
        }
        return nil
    }
    res["subjectLifecycle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageSubjectLifecycle)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectLifecycle(val.(*AccessPackageSubjectLifecycle))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetObjectId gets the objectId property value. The object identifier of the subject. null if the subject isn't yet a user in the tenant. Alternate key.
// returns a *string when successful
func (m *AccessPackageSubject) GetObjectId()(*string) {
    val, err := m.GetBackingStore().Get("objectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. The onPremisesSecurityIdentifier property
// returns a *string when successful
func (m *AccessPackageSubject) GetOnPremisesSecurityIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("onPremisesSecurityIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrincipalName gets the principalName property value. The principal name, if known, of the subject.
// returns a *string when successful
func (m *AccessPackageSubject) GetPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("principalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubjectLifecycle gets the subjectLifecycle property value. The lifecycle of the subject user, if a guest. The possible values are: notDefined, notGoverned, governed, unknownFutureValue.
// returns a *AccessPackageSubjectLifecycle when successful
func (m *AccessPackageSubject) GetSubjectLifecycle()(*AccessPackageSubjectLifecycle) {
    val, err := m.GetBackingStore().Get("subjectLifecycle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AccessPackageSubjectLifecycle)
    }
    return nil
}
// GetTypeEscaped gets the type property value. The resource type of the subject.
// returns a *string when successful
func (m *AccessPackageSubject) GetTypeEscaped()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageSubject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("altSecId", m.GetAltSecId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("cleanupScheduledDateTime", m.GetCleanupScheduledDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectedOrganization", m.GetConnectedOrganization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("connectedOrganizationId", m.GetConnectedOrganizationId())
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
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("objectId", m.GetObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalName", m.GetPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectLifecycle() != nil {
        cast := (*m.GetSubjectLifecycle()).String()
        err = writer.WriteStringValue("subjectLifecycle", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAltSecId sets the altSecId property value. Not Supported.
func (m *AccessPackageSubject) SetAltSecId(value *string)() {
    err := m.GetBackingStore().Set("altSecId", value)
    if err != nil {
        panic(err)
    }
}
// SetCleanupScheduledDateTime sets the cleanupScheduledDateTime property value. The date and time the subject is marked to be blocked from sign in or deleted. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.
func (m *AccessPackageSubject) SetCleanupScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("cleanupScheduledDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectedOrganization sets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
func (m *AccessPackageSubject) SetConnectedOrganization(value ConnectedOrganizationable)() {
    err := m.GetBackingStore().Set("connectedOrganization", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectedOrganizationId sets the connectedOrganizationId property value. The identifier of the connected organization of the subject.
func (m *AccessPackageSubject) SetConnectedOrganizationId(value *string)() {
    err := m.GetBackingStore().Set("connectedOrganizationId", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the subject.
func (m *AccessPackageSubject) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEmail sets the email property value. The email address of the subject.
func (m *AccessPackageSubject) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
// SetObjectId sets the objectId property value. The object identifier of the subject. null if the subject isn't yet a user in the tenant. Alternate key.
func (m *AccessPackageSubject) SetObjectId(value *string)() {
    err := m.GetBackingStore().Set("objectId", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. The onPremisesSecurityIdentifier property
func (m *AccessPackageSubject) SetOnPremisesSecurityIdentifier(value *string)() {
    err := m.GetBackingStore().Set("onPremisesSecurityIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetPrincipalName sets the principalName property value. The principal name, if known, of the subject.
func (m *AccessPackageSubject) SetPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("principalName", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectLifecycle sets the subjectLifecycle property value. The lifecycle of the subject user, if a guest. The possible values are: notDefined, notGoverned, governed, unknownFutureValue.
func (m *AccessPackageSubject) SetSubjectLifecycle(value *AccessPackageSubjectLifecycle)() {
    err := m.GetBackingStore().Set("subjectLifecycle", value)
    if err != nil {
        panic(err)
    }
}
// SetTypeEscaped sets the type property value. The resource type of the subject.
func (m *AccessPackageSubject) SetTypeEscaped(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
type AccessPackageSubjectable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAltSecId()(*string)
    GetCleanupScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConnectedOrganization()(ConnectedOrganizationable)
    GetConnectedOrganizationId()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetObjectId()(*string)
    GetOnPremisesSecurityIdentifier()(*string)
    GetPrincipalName()(*string)
    GetSubjectLifecycle()(*AccessPackageSubjectLifecycle)
    GetTypeEscaped()(*string)
    SetAltSecId(value *string)()
    SetCleanupScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConnectedOrganization(value ConnectedOrganizationable)()
    SetConnectedOrganizationId(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetObjectId(value *string)()
    SetOnPremisesSecurityIdentifier(value *string)()
    SetPrincipalName(value *string)()
    SetSubjectLifecycle(value *AccessPackageSubjectLifecycle)()
    SetTypeEscaped(value *string)()
}
