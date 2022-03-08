package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPCable 
type CloudPCable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAadDeviceId()(*string)
    GetDisplayName()(*string)
    GetGracePeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetImageDisplayName()(*string)
    GetLastLoginResult()(CloudPcLoginResultable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastRemoteActionResult()(CloudPcRemoteActionResultable)
    GetManagedDeviceId()(*string)
    GetManagedDeviceName()(*string)
    GetOnPremisesConnectionName()(*string)
    GetOsVersion()(*CloudPcOperatingSystem)
    GetProvisioningPolicyId()(*string)
    GetProvisioningPolicyName()(*string)
    GetServicePlanId()(*string)
    GetServicePlanName()(*string)
    GetServicePlanType()(*CloudPcServicePlanType)
    GetStatus()(*CloudPcStatus)
    GetStatusDetails()(CloudPcStatusDetailsable)
    GetUserAccountType()(*CloudPcUserAccountType)
    GetUserPrincipalName()(*string)
    SetAadDeviceId(value *string)()
    SetDisplayName(value *string)()
    SetGracePeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetImageDisplayName(value *string)()
    SetLastLoginResult(value CloudPcLoginResultable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastRemoteActionResult(value CloudPcRemoteActionResultable)()
    SetManagedDeviceId(value *string)()
    SetManagedDeviceName(value *string)()
    SetOnPremisesConnectionName(value *string)()
    SetOsVersion(value *CloudPcOperatingSystem)()
    SetProvisioningPolicyId(value *string)()
    SetProvisioningPolicyName(value *string)()
    SetServicePlanId(value *string)()
    SetServicePlanName(value *string)()
    SetServicePlanType(value *CloudPcServicePlanType)()
    SetStatus(value *CloudPcStatus)()
    SetStatusDetails(value CloudPcStatusDetailsable)()
    SetUserAccountType(value *CloudPcUserAccountType)()
    SetUserPrincipalName(value *string)()
}
