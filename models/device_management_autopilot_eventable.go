package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAutopilotEventable 
type DeviceManagementAutopilotEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAccountSetupStatus()(*WindowsAutopilotDeploymentState)
    GetDeploymentDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDeploymentEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentState()(*WindowsAutopilotDeploymentState)
    GetDeploymentTotalDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDeviceId()(*string)
    GetDevicePreparationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDeviceRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceSerialNumber()(*string)
    GetDeviceSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDeviceSetupStatus()(*WindowsAutopilotDeploymentState)
    GetEnrollmentFailureDetails()(*string)
    GetEnrollmentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEnrollmentState()(*EnrollmentState)
    GetEnrollmentType()(*WindowsAutopilotEnrollmentType)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceName()(*string)
    GetOsVersion()(*string)
    GetPolicyStatusDetails()([]DeviceManagementAutopilotPolicyStatusDetailable)
    GetTargetedAppCount()(*int32)
    GetTargetedPolicyCount()(*int32)
    GetUserPrincipalName()(*string)
    GetWindows10EnrollmentCompletionPageConfigurationDisplayName()(*string)
    GetWindows10EnrollmentCompletionPageConfigurationId()(*string)
    GetWindowsAutopilotDeploymentProfileDisplayName()(*string)
    SetAccountSetupDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAccountSetupStatus(value *WindowsAutopilotDeploymentState)()
    SetDeploymentDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDeploymentEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentState(value *WindowsAutopilotDeploymentState)()
    SetDeploymentTotalDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDeviceId(value *string)()
    SetDevicePreparationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDeviceRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceSerialNumber(value *string)()
    SetDeviceSetupDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDeviceSetupStatus(value *WindowsAutopilotDeploymentState)()
    SetEnrollmentFailureDetails(value *string)()
    SetEnrollmentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEnrollmentState(value *EnrollmentState)()
    SetEnrollmentType(value *WindowsAutopilotEnrollmentType)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceName(value *string)()
    SetOsVersion(value *string)()
    SetPolicyStatusDetails(value []DeviceManagementAutopilotPolicyStatusDetailable)()
    SetTargetedAppCount(value *int32)()
    SetTargetedPolicyCount(value *int32)()
    SetUserPrincipalName(value *string)()
    SetWindows10EnrollmentCompletionPageConfigurationDisplayName(value *string)()
    SetWindows10EnrollmentCompletionPageConfigurationId(value *string)()
    SetWindowsAutopilotDeploymentProfileDisplayName(value *string)()
}
