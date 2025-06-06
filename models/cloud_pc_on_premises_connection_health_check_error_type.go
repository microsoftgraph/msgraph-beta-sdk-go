// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
type CloudPcOnPremisesConnectionHealthCheckErrorType int

const (
    DNSCHECKFQDNNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE CloudPcOnPremisesConnectionHealthCheckErrorType = iota
    DNSCHECKNAMEWITHINVALIDCHARACTER_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    DNSCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKFQDNNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKINCORRECTCREDENTIALS_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKORGANIZATIONALUNITNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKORGANIZATIONALUNITINCORRECTFORMAT_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKCOMPUTEROBJECTALREADYEXISTS_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKACCESSDENIED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKCREDENTIALSEXPIRED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKACCOUNTLOCKEDORDISABLED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKACCOUNTQUOTAEXCEEDED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKSERVERNOTOPERATIONAL_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ADJOINCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ENDPOINTCONNECTIVITYCHECKCLOUDPCURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ENDPOINTCONNECTIVITYCHECKWVDURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ENDPOINTCONNECTIVITYCHECKINTUNEURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ENDPOINTCONNECTIVITYCHECKAZUREADURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ENDPOINTCONNECTIVITYCHECKLOCALEURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ENDPOINTCONNECTIVITYCHECKVMAGENTENDPOINTCOMMUNICATIONERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    ENDPOINTCONNECTIVITYCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    AZUREADDEVICESYNCCHECKDEVICENOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    AZUREADDEVICESYNCCHECKLONGSYNCCIRCLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    AZUREADDEVICESYNCCHECKCONNECTDISABLED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    AZUREADDEVICESYNCCHECKDURATIONEXCEEDED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    AZUREADDEVICESYNCCHECKSCPNOTCONFIGURED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    AZUREADDEVICESYNCCHECKTRANSIENTSERVICEERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    AZUREADDEVICESYNCCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKNOSUBNETIP_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKSUBSCRIPTIONDISABLED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKAZUREPOLICYVIOLATION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKSUBSCRIPTIONNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKSUBSCRIPTIONTRANSFERRED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKGENERALSUBSCRIPTIONERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKUNSUPPORTEDVNETREGION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKRESOURCEGROUPINVALID_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKVNETINVALID_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKSUBNETINVALID_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKRESOURCEGROUPBEINGDELETED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKVNETBEINGMOVED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKSUBNETDELEGATIONFAILED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKSUBNETWITHEXTERNALRESOURCES_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKRESOURCEGROUPLOCKEDFORREADONLY_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKRESOURCEGROUPLOCKEDFORDELETE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKNOINTUNEREADERROLEERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKINTUNEDEFAULTWINDOWSRESTRICTIONVIOLATION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKINTUNECUSTOMWINDOWSRESTRICTIONVIOLATION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKDEPLOYMENTQUOTALIMITREACHED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKMISSINGREGISTRATIONFORLOCATION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKTRANSIENTSERVICEERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    RESOURCEAVAILABILITYCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    PERMISSIONCHECKNOSUBSCRIPTIONREADERROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    PERMISSIONCHECKNORESOURCEGROUPOWNERROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    PERMISSIONCHECKNOVNETCONTRIBUTORROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    PERMISSIONCHECKNORESOURCEGROUPNETWORKCONTRIBUTORROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    PERMISSIONCHECKNOWINDOWS365NETWORKUSERROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    PERMISSIONCHECKNOWINDOWS365NETWORKINTERFACECONTRIBUTORROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    PERMISSIONCHECKTRANSIENTSERVICEERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    PERMISSIONCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    UDPCONNECTIVITYCHECKSTUNURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    UDPCONNECTIVITYCHECKTURNURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    UDPCONNECTIVITYCHECKURLSNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    UDPCONNECTIVITYCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    INTERNALSERVERERRORDEPLOYMENTCANCELED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    INTERNALSERVERERRORALLOCATERESOURCEFAILED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    INTERNALSERVERERRORVMDEPLOYMENTTIMEOUT_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    INTERNALSERVERERRORUNABLETORUNDSCSCRIPT_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    SSOCHECKKERBEROSCONFIGURATIONERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    INTERNALSERVERUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
)

func (i CloudPcOnPremisesConnectionHealthCheckErrorType) String() string {
    return []string{"dnsCheckFqdnNotFound", "dnsCheckNameWithInvalidCharacter", "dnsCheckUnknownError", "adJoinCheckFqdnNotFound", "adJoinCheckIncorrectCredentials", "adJoinCheckOrganizationalUnitNotFound", "adJoinCheckOrganizationalUnitIncorrectFormat", "adJoinCheckComputerObjectAlreadyExists", "adJoinCheckAccessDenied", "adJoinCheckCredentialsExpired", "adJoinCheckAccountLockedOrDisabled", "adJoinCheckAccountQuotaExceeded", "adJoinCheckServerNotOperational", "adJoinCheckUnknownError", "endpointConnectivityCheckCloudPcUrlNotAllowListed", "endpointConnectivityCheckWVDUrlNotAllowListed", "endpointConnectivityCheckIntuneUrlNotAllowListed", "endpointConnectivityCheckAzureADUrlNotAllowListed", "endpointConnectivityCheckLocaleUrlNotAllowListed", "endpointConnectivityCheckVMAgentEndPointCommunicationError", "endpointConnectivityCheckUnknownError", "azureAdDeviceSyncCheckDeviceNotFound", "azureAdDeviceSyncCheckLongSyncCircle", "azureAdDeviceSyncCheckConnectDisabled", "azureAdDeviceSyncCheckDurationExceeded", "azureAdDeviceSyncCheckScpNotConfigured", "azureAdDeviceSyncCheckTransientServiceError", "azureAdDeviceSyncCheckUnknownError", "resourceAvailabilityCheckNoSubnetIP", "resourceAvailabilityCheckSubscriptionDisabled", "resourceAvailabilityCheckAzurePolicyViolation", "resourceAvailabilityCheckSubscriptionNotFound", "resourceAvailabilityCheckSubscriptionTransferred", "resourceAvailabilityCheckGeneralSubscriptionError", "resourceAvailabilityCheckUnsupportedVNetRegion", "resourceAvailabilityCheckResourceGroupInvalid", "resourceAvailabilityCheckVNetInvalid", "resourceAvailabilityCheckSubnetInvalid", "resourceAvailabilityCheckResourceGroupBeingDeleted", "resourceAvailabilityCheckVNetBeingMoved", "resourceAvailabilityCheckSubnetDelegationFailed", "resourceAvailabilityCheckSubnetWithExternalResources", "resourceAvailabilityCheckResourceGroupLockedForReadonly", "resourceAvailabilityCheckResourceGroupLockedForDelete", "resourceAvailabilityCheckNoIntuneReaderRoleError", "resourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation", "resourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation", "resourceAvailabilityCheckDeploymentQuotaLimitReached", "resourceAvailabilityCheckMissingRegistrationForLocation", "resourceAvailabilityCheckTransientServiceError", "resourceAvailabilityCheckUnknownError", "permissionCheckNoSubscriptionReaderRole", "permissionCheckNoResourceGroupOwnerRole", "permissionCheckNoVNetContributorRole", "permissionCheckNoResourceGroupNetworkContributorRole", "permissionCheckNoWindows365NetworkUserRole", "permissionCheckNoWindows365NetworkInterfaceContributorRole", "permissionCheckTransientServiceError", "permissionCheckUnknownError", "udpConnectivityCheckStunUrlNotAllowListed", "udpConnectivityCheckTurnUrlNotAllowListed", "udpConnectivityCheckUrlsNotAllowListed", "udpConnectivityCheckUnknownError", "internalServerErrorDeploymentCanceled", "internalServerErrorAllocateResourceFailed", "internalServerErrorVMDeploymentTimeout", "internalServerErrorUnableToRunDscScript", "ssoCheckKerberosConfigurationError", "internalServerUnknownError", "unknownFutureValue"}[i]
}
func ParseCloudPcOnPremisesConnectionHealthCheckErrorType(v string) (any, error) {
    result := DNSCHECKFQDNNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
    switch v {
        case "dnsCheckFqdnNotFound":
            result = DNSCHECKFQDNNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "dnsCheckNameWithInvalidCharacter":
            result = DNSCHECKNAMEWITHINVALIDCHARACTER_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "dnsCheckUnknownError":
            result = DNSCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckFqdnNotFound":
            result = ADJOINCHECKFQDNNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckIncorrectCredentials":
            result = ADJOINCHECKINCORRECTCREDENTIALS_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckOrganizationalUnitNotFound":
            result = ADJOINCHECKORGANIZATIONALUNITNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckOrganizationalUnitIncorrectFormat":
            result = ADJOINCHECKORGANIZATIONALUNITINCORRECTFORMAT_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckComputerObjectAlreadyExists":
            result = ADJOINCHECKCOMPUTEROBJECTALREADYEXISTS_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckAccessDenied":
            result = ADJOINCHECKACCESSDENIED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckCredentialsExpired":
            result = ADJOINCHECKCREDENTIALSEXPIRED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckAccountLockedOrDisabled":
            result = ADJOINCHECKACCOUNTLOCKEDORDISABLED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckAccountQuotaExceeded":
            result = ADJOINCHECKACCOUNTQUOTAEXCEEDED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckServerNotOperational":
            result = ADJOINCHECKSERVERNOTOPERATIONAL_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "adJoinCheckUnknownError":
            result = ADJOINCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "endpointConnectivityCheckCloudPcUrlNotAllowListed":
            result = ENDPOINTCONNECTIVITYCHECKCLOUDPCURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "endpointConnectivityCheckWVDUrlNotAllowListed":
            result = ENDPOINTCONNECTIVITYCHECKWVDURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "endpointConnectivityCheckIntuneUrlNotAllowListed":
            result = ENDPOINTCONNECTIVITYCHECKINTUNEURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "endpointConnectivityCheckAzureADUrlNotAllowListed":
            result = ENDPOINTCONNECTIVITYCHECKAZUREADURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "endpointConnectivityCheckLocaleUrlNotAllowListed":
            result = ENDPOINTCONNECTIVITYCHECKLOCALEURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "endpointConnectivityCheckVMAgentEndPointCommunicationError":
            result = ENDPOINTCONNECTIVITYCHECKVMAGENTENDPOINTCOMMUNICATIONERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "endpointConnectivityCheckUnknownError":
            result = ENDPOINTCONNECTIVITYCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "azureAdDeviceSyncCheckDeviceNotFound":
            result = AZUREADDEVICESYNCCHECKDEVICENOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "azureAdDeviceSyncCheckLongSyncCircle":
            result = AZUREADDEVICESYNCCHECKLONGSYNCCIRCLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "azureAdDeviceSyncCheckConnectDisabled":
            result = AZUREADDEVICESYNCCHECKCONNECTDISABLED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "azureAdDeviceSyncCheckDurationExceeded":
            result = AZUREADDEVICESYNCCHECKDURATIONEXCEEDED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "azureAdDeviceSyncCheckScpNotConfigured":
            result = AZUREADDEVICESYNCCHECKSCPNOTCONFIGURED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "azureAdDeviceSyncCheckTransientServiceError":
            result = AZUREADDEVICESYNCCHECKTRANSIENTSERVICEERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "azureAdDeviceSyncCheckUnknownError":
            result = AZUREADDEVICESYNCCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckNoSubnetIP":
            result = RESOURCEAVAILABILITYCHECKNOSUBNETIP_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckSubscriptionDisabled":
            result = RESOURCEAVAILABILITYCHECKSUBSCRIPTIONDISABLED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckAzurePolicyViolation":
            result = RESOURCEAVAILABILITYCHECKAZUREPOLICYVIOLATION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckSubscriptionNotFound":
            result = RESOURCEAVAILABILITYCHECKSUBSCRIPTIONNOTFOUND_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckSubscriptionTransferred":
            result = RESOURCEAVAILABILITYCHECKSUBSCRIPTIONTRANSFERRED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckGeneralSubscriptionError":
            result = RESOURCEAVAILABILITYCHECKGENERALSUBSCRIPTIONERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckUnsupportedVNetRegion":
            result = RESOURCEAVAILABILITYCHECKUNSUPPORTEDVNETREGION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckResourceGroupInvalid":
            result = RESOURCEAVAILABILITYCHECKRESOURCEGROUPINVALID_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckVNetInvalid":
            result = RESOURCEAVAILABILITYCHECKVNETINVALID_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckSubnetInvalid":
            result = RESOURCEAVAILABILITYCHECKSUBNETINVALID_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckResourceGroupBeingDeleted":
            result = RESOURCEAVAILABILITYCHECKRESOURCEGROUPBEINGDELETED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckVNetBeingMoved":
            result = RESOURCEAVAILABILITYCHECKVNETBEINGMOVED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckSubnetDelegationFailed":
            result = RESOURCEAVAILABILITYCHECKSUBNETDELEGATIONFAILED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckSubnetWithExternalResources":
            result = RESOURCEAVAILABILITYCHECKSUBNETWITHEXTERNALRESOURCES_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckResourceGroupLockedForReadonly":
            result = RESOURCEAVAILABILITYCHECKRESOURCEGROUPLOCKEDFORREADONLY_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckResourceGroupLockedForDelete":
            result = RESOURCEAVAILABILITYCHECKRESOURCEGROUPLOCKEDFORDELETE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckNoIntuneReaderRoleError":
            result = RESOURCEAVAILABILITYCHECKNOINTUNEREADERROLEERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckIntuneDefaultWindowsRestrictionViolation":
            result = RESOURCEAVAILABILITYCHECKINTUNEDEFAULTWINDOWSRESTRICTIONVIOLATION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckIntuneCustomWindowsRestrictionViolation":
            result = RESOURCEAVAILABILITYCHECKINTUNECUSTOMWINDOWSRESTRICTIONVIOLATION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckDeploymentQuotaLimitReached":
            result = RESOURCEAVAILABILITYCHECKDEPLOYMENTQUOTALIMITREACHED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckMissingRegistrationForLocation":
            result = RESOURCEAVAILABILITYCHECKMISSINGREGISTRATIONFORLOCATION_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckTransientServiceError":
            result = RESOURCEAVAILABILITYCHECKTRANSIENTSERVICEERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "resourceAvailabilityCheckUnknownError":
            result = RESOURCEAVAILABILITYCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "permissionCheckNoSubscriptionReaderRole":
            result = PERMISSIONCHECKNOSUBSCRIPTIONREADERROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "permissionCheckNoResourceGroupOwnerRole":
            result = PERMISSIONCHECKNORESOURCEGROUPOWNERROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "permissionCheckNoVNetContributorRole":
            result = PERMISSIONCHECKNOVNETCONTRIBUTORROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "permissionCheckNoResourceGroupNetworkContributorRole":
            result = PERMISSIONCHECKNORESOURCEGROUPNETWORKCONTRIBUTORROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "permissionCheckNoWindows365NetworkUserRole":
            result = PERMISSIONCHECKNOWINDOWS365NETWORKUSERROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "permissionCheckNoWindows365NetworkInterfaceContributorRole":
            result = PERMISSIONCHECKNOWINDOWS365NETWORKINTERFACECONTRIBUTORROLE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "permissionCheckTransientServiceError":
            result = PERMISSIONCHECKTRANSIENTSERVICEERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "permissionCheckUnknownError":
            result = PERMISSIONCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "udpConnectivityCheckStunUrlNotAllowListed":
            result = UDPCONNECTIVITYCHECKSTUNURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "udpConnectivityCheckTurnUrlNotAllowListed":
            result = UDPCONNECTIVITYCHECKTURNURLNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "udpConnectivityCheckUrlsNotAllowListed":
            result = UDPCONNECTIVITYCHECKURLSNOTALLOWLISTED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "udpConnectivityCheckUnknownError":
            result = UDPCONNECTIVITYCHECKUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "internalServerErrorDeploymentCanceled":
            result = INTERNALSERVERERRORDEPLOYMENTCANCELED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "internalServerErrorAllocateResourceFailed":
            result = INTERNALSERVERERRORALLOCATERESOURCEFAILED_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "internalServerErrorVMDeploymentTimeout":
            result = INTERNALSERVERERRORVMDEPLOYMENTTIMEOUT_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "internalServerErrorUnableToRunDscScript":
            result = INTERNALSERVERERRORUNABLETORUNDSCSCRIPT_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "ssoCheckKerberosConfigurationError":
            result = SSOCHECKKERBEROSCONFIGURATIONERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "internalServerUnknownError":
            result = INTERNALSERVERUNKNOWNERROR_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONHEALTHCHECKERRORTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcOnPremisesConnectionHealthCheckErrorType(values []CloudPcOnPremisesConnectionHealthCheckErrorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcOnPremisesConnectionHealthCheckErrorType) isMultiValue() bool {
    return false
}
