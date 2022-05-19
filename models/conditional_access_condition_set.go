package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessConditionSet 
type ConditionalAccessConditionSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Applications and user actions included in and excluded from the policy. Required.
    applications ConditionalAccessApplicationsable
    // Client applications (service principals and workload identities) included in and excluded from the policy. Either users or clientApplications is required.
    clientApplications ConditionalAccessClientApplicationsable
    // Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
    clientAppTypes []string
    // Devices in the policy.
    devices ConditionalAccessDevicesable
    // Device states in the policy.
    deviceStates ConditionalAccessDeviceStatesable
    // Locations included in and excluded from the policy.
    locations ConditionalAccessLocationsable
    // Platforms included in and excluded from the policy.
    platforms ConditionalAccessPlatformsable
    // Service principal risk levels included in the policy. Possible values are: low, medium, high, none, unknownFutureValue.
    servicePrincipalRiskLevels []string
    // Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
    signInRiskLevels []string
    // User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
    userRiskLevels []string
    // Users, groups, and roles included in and excluded from the policy. Either users or clientApplications is required.
    users ConditionalAccessUsersable
}
// NewConditionalAccessConditionSet instantiates a new conditionalAccessConditionSet and sets the default values.
func NewConditionalAccessConditionSet()(*ConditionalAccessConditionSet) {
    m := &ConditionalAccessConditionSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessConditionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessConditionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessConditionSet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessConditionSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplications gets the applications property value. Applications and user actions included in and excluded from the policy. Required.
func (m *ConditionalAccessConditionSet) GetApplications()(ConditionalAccessApplicationsable) {
    if m == nil {
        return nil
    } else {
        return m.applications
    }
}
// GetClientApplications gets the clientApplications property value. Client applications (service principals and workload identities) included in and excluded from the policy. Either users or clientApplications is required.
func (m *ConditionalAccessConditionSet) GetClientApplications()(ConditionalAccessClientApplicationsable) {
    if m == nil {
        return nil
    } else {
        return m.clientApplications
    }
}
// GetClientAppTypes gets the clientAppTypes property value. Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
func (m *ConditionalAccessConditionSet) GetClientAppTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.clientAppTypes
    }
}
// GetDevices gets the devices property value. Devices in the policy.
func (m *ConditionalAccessConditionSet) GetDevices()(ConditionalAccessDevicesable) {
    if m == nil {
        return nil
    } else {
        return m.devices
    }
}
// GetDeviceStates gets the deviceStates property value. Device states in the policy.
func (m *ConditionalAccessConditionSet) GetDeviceStates()(ConditionalAccessDeviceStatesable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessConditionSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplications(val.(ConditionalAccessApplicationsable))
        }
        return nil
    }
    res["clientApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessClientApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientApplications(val.(ConditionalAccessClientApplicationsable))
        }
        return nil
    }
    res["clientAppTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetClientAppTypes(res)
        }
        return nil
    }
    res["devices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessDevicesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevices(val.(ConditionalAccessDevicesable))
        }
        return nil
    }
    res["deviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessDeviceStatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStates(val.(ConditionalAccessDeviceStatesable))
        }
        return nil
    }
    res["locations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessLocationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocations(val.(ConditionalAccessLocationsable))
        }
        return nil
    }
    res["platforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessPlatformsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatforms(val.(ConditionalAccessPlatformsable))
        }
        return nil
    }
    res["servicePrincipalRiskLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetServicePrincipalRiskLevels(res)
        }
        return nil
    }
    res["signInRiskLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSignInRiskLevels(res)
        }
        return nil
    }
    res["userRiskLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUserRiskLevels(res)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessUsersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsers(val.(ConditionalAccessUsersable))
        }
        return nil
    }
    return res
}
// GetLocations gets the locations property value. Locations included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) GetLocations()(ConditionalAccessLocationsable) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
// GetPlatforms gets the platforms property value. Platforms included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) GetPlatforms()(ConditionalAccessPlatformsable) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
// GetServicePrincipalRiskLevels gets the servicePrincipalRiskLevels property value. Service principal risk levels included in the policy. Possible values are: low, medium, high, none, unknownFutureValue.
func (m *ConditionalAccessConditionSet) GetServicePrincipalRiskLevels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalRiskLevels
    }
}
// GetSignInRiskLevels gets the signInRiskLevels property value. Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) GetSignInRiskLevels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.signInRiskLevels
    }
}
// GetUserRiskLevels gets the userRiskLevels property value. User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) GetUserRiskLevels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userRiskLevels
    }
}
// GetUsers gets the users property value. Users, groups, and roles included in and excluded from the policy. Either users or clientApplications is required.
func (m *ConditionalAccessConditionSet) GetUsers()(ConditionalAccessUsersable) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
// Serialize serializes information the current object
func (m *ConditionalAccessConditionSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applications", m.GetApplications())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("clientApplications", m.GetClientApplications())
        if err != nil {
            return err
        }
    }
    if m.GetClientAppTypes() != nil {
        err := writer.WriteCollectionOfStringValues("clientAppTypes", m.GetClientAppTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("devices", m.GetDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceStates", m.GetDeviceStates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("locations", m.GetLocations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("platforms", m.GetPlatforms())
        if err != nil {
            return err
        }
    }
    if m.GetServicePrincipalRiskLevels() != nil {
        err := writer.WriteCollectionOfStringValues("servicePrincipalRiskLevels", m.GetServicePrincipalRiskLevels())
        if err != nil {
            return err
        }
    }
    if m.GetSignInRiskLevels() != nil {
        err := writer.WriteCollectionOfStringValues("signInRiskLevels", m.GetSignInRiskLevels())
        if err != nil {
            return err
        }
    }
    if m.GetUserRiskLevels() != nil {
        err := writer.WriteCollectionOfStringValues("userRiskLevels", m.GetUserRiskLevels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("users", m.GetUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessConditionSet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplications sets the applications property value. Applications and user actions included in and excluded from the policy. Required.
func (m *ConditionalAccessConditionSet) SetApplications(value ConditionalAccessApplicationsable)() {
    if m != nil {
        m.applications = value
    }
}
// SetClientApplications sets the clientApplications property value. Client applications (service principals and workload identities) included in and excluded from the policy. Either users or clientApplications is required.
func (m *ConditionalAccessConditionSet) SetClientApplications(value ConditionalAccessClientApplicationsable)() {
    if m != nil {
        m.clientApplications = value
    }
}
// SetClientAppTypes sets the clientAppTypes property value. Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
func (m *ConditionalAccessConditionSet) SetClientAppTypes(value []string)() {
    if m != nil {
        m.clientAppTypes = value
    }
}
// SetDevices sets the devices property value. Devices in the policy.
func (m *ConditionalAccessConditionSet) SetDevices(value ConditionalAccessDevicesable)() {
    if m != nil {
        m.devices = value
    }
}
// SetDeviceStates sets the deviceStates property value. Device states in the policy.
func (m *ConditionalAccessConditionSet) SetDeviceStates(value ConditionalAccessDeviceStatesable)() {
    if m != nil {
        m.deviceStates = value
    }
}
// SetLocations sets the locations property value. Locations included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) SetLocations(value ConditionalAccessLocationsable)() {
    if m != nil {
        m.locations = value
    }
}
// SetPlatforms sets the platforms property value. Platforms included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) SetPlatforms(value ConditionalAccessPlatformsable)() {
    if m != nil {
        m.platforms = value
    }
}
// SetServicePrincipalRiskLevels sets the servicePrincipalRiskLevels property value. Service principal risk levels included in the policy. Possible values are: low, medium, high, none, unknownFutureValue.
func (m *ConditionalAccessConditionSet) SetServicePrincipalRiskLevels(value []string)() {
    if m != nil {
        m.servicePrincipalRiskLevels = value
    }
}
// SetSignInRiskLevels sets the signInRiskLevels property value. Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) SetSignInRiskLevels(value []string)() {
    if m != nil {
        m.signInRiskLevels = value
    }
}
// SetUserRiskLevels sets the userRiskLevels property value. User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) SetUserRiskLevels(value []string)() {
    if m != nil {
        m.userRiskLevels = value
    }
}
// SetUsers sets the users property value. Users, groups, and roles included in and excluded from the policy. Either users or clientApplications is required.
func (m *ConditionalAccessConditionSet) SetUsers(value ConditionalAccessUsersable)() {
    if m != nil {
        m.users = value
    }
}
