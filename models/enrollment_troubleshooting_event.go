package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentTroubleshootingEvent 
type EnrollmentTroubleshootingEvent struct {
    DeviceManagementTroubleshootingEvent
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Azure AD device identifier.
    deviceId *string
    // Type of the enrollment. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount, azureAdJoinUsingAzureVmExtension, androidEnterpriseDedicatedDevice, androidEnterpriseFullyManaged, androidEnterpriseCorporateWorkProfile.
    enrollmentType *DeviceEnrollmentType
    // Highlevel failure category. Possible values are: unknown, authentication, authorization, accountValidation, userValidation, deviceNotSupported, inMaintenance, badRequest, featureNotSupported, enrollmentRestrictionsEnforced, clientDisconnected, userAbandonment.
    failureCategory *DeviceEnrollmentFailureReason
    // Detailed failure reason.
    failureReason *string
    // Device identifier created or collected by Intune.
    managedDeviceIdentifier *string
    // Operating System.
    operatingSystem *string
    // OS Version.
    osVersion *string
    // Identifier for the user that tried to enroll the device.
    userId *string
}
// NewEnrollmentTroubleshootingEvent instantiates a new EnrollmentTroubleshootingEvent and sets the default values.
func NewEnrollmentTroubleshootingEvent()(*EnrollmentTroubleshootingEvent) {
    m := &EnrollmentTroubleshootingEvent{
        DeviceManagementTroubleshootingEvent: *NewDeviceManagementTroubleshootingEvent(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEnrollmentTroubleshootingEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollmentTroubleshootingEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentTroubleshootingEvent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnrollmentTroubleshootingEvent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceId gets the deviceId property value. Azure AD device identifier.
func (m *EnrollmentTroubleshootingEvent) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetEnrollmentType gets the enrollmentType property value. Type of the enrollment. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount, azureAdJoinUsingAzureVmExtension, androidEnterpriseDedicatedDevice, androidEnterpriseFullyManaged, androidEnterpriseCorporateWorkProfile.
func (m *EnrollmentTroubleshootingEvent) GetEnrollmentType()(*DeviceEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentType
    }
}
// GetFailureCategory gets the failureCategory property value. Highlevel failure category. Possible values are: unknown, authentication, authorization, accountValidation, userValidation, deviceNotSupported, inMaintenance, badRequest, featureNotSupported, enrollmentRestrictionsEnforced, clientDisconnected, userAbandonment.
func (m *EnrollmentTroubleshootingEvent) GetFailureCategory()(*DeviceEnrollmentFailureReason) {
    if m == nil {
        return nil
    } else {
        return m.failureCategory
    }
}
// GetFailureReason gets the failureReason property value. Detailed failure reason.
func (m *EnrollmentTroubleshootingEvent) GetFailureReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failureReason
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentTroubleshootingEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementTroubleshootingEvent.GetFieldDeserializers()
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["enrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val.(*DeviceEnrollmentType))
        }
        return nil
    }
    res["failureCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentFailureReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureCategory(val.(*DeviceEnrollmentFailureReason))
        }
        return nil
    }
    res["failureReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureReason(val)
        }
        return nil
    }
    res["managedDeviceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceIdentifier(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceIdentifier gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *EnrollmentTroubleshootingEvent) GetManagedDeviceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceIdentifier
    }
}
// GetOperatingSystem gets the operatingSystem property value. Operating System.
func (m *EnrollmentTroubleshootingEvent) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// GetOsVersion gets the osVersion property value. OS Version.
func (m *EnrollmentTroubleshootingEvent) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetUserId gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *EnrollmentTroubleshootingEvent) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Serialize serializes information the current object
func (m *EnrollmentTroubleshootingEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementTroubleshootingEvent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentType() != nil {
        cast := (*m.GetEnrollmentType()).String()
        err = writer.WriteStringValue("enrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFailureCategory() != nil {
        cast := (*m.GetFailureCategory()).String()
        err = writer.WriteStringValue("failureCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("failureReason", m.GetFailureReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceIdentifier", m.GetManagedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnrollmentTroubleshootingEvent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceId sets the deviceId property value. Azure AD device identifier.
func (m *EnrollmentTroubleshootingEvent) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetEnrollmentType sets the enrollmentType property value. Type of the enrollment. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount, azureAdJoinUsingAzureVmExtension, androidEnterpriseDedicatedDevice, androidEnterpriseFullyManaged, androidEnterpriseCorporateWorkProfile.
func (m *EnrollmentTroubleshootingEvent) SetEnrollmentType(value *DeviceEnrollmentType)() {
    if m != nil {
        m.enrollmentType = value
    }
}
// SetFailureCategory sets the failureCategory property value. Highlevel failure category. Possible values are: unknown, authentication, authorization, accountValidation, userValidation, deviceNotSupported, inMaintenance, badRequest, featureNotSupported, enrollmentRestrictionsEnforced, clientDisconnected, userAbandonment.
func (m *EnrollmentTroubleshootingEvent) SetFailureCategory(value *DeviceEnrollmentFailureReason)() {
    if m != nil {
        m.failureCategory = value
    }
}
// SetFailureReason sets the failureReason property value. Detailed failure reason.
func (m *EnrollmentTroubleshootingEvent) SetFailureReason(value *string)() {
    if m != nil {
        m.failureReason = value
    }
}
// SetManagedDeviceIdentifier sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *EnrollmentTroubleshootingEvent) SetManagedDeviceIdentifier(value *string)() {
    if m != nil {
        m.managedDeviceIdentifier = value
    }
}
// SetOperatingSystem sets the operatingSystem property value. Operating System.
func (m *EnrollmentTroubleshootingEvent) SetOperatingSystem(value *string)() {
    if m != nil {
        m.operatingSystem = value
    }
}
// SetOsVersion sets the osVersion property value. OS Version.
func (m *EnrollmentTroubleshootingEvent) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetUserId sets the userId property value. Identifier for the user that tried to enroll the device.
func (m *EnrollmentTroubleshootingEvent) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
