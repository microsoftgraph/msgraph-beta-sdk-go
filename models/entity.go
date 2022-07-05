package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/tenantadmin"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
    iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb "github.com/microsoftgraph/msgraph-beta-sdk-go/models/search"
)

// Entity 
type Entity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The id property
    id *string
}
// NewEntity instantiates a new entity and sets the default values.
func NewEntity()(*Entity) {
    m := &Entity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEntityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.accessPackage":
                        return NewAccessPackage(), nil
                    case "#microsoft.graph.accessPackageAssignment":
                        return NewAccessPackageAssignment(), nil
                    case "#microsoft.graph.accessPackageAssignmentPolicy":
                        return NewAccessPackageAssignmentPolicy(), nil
                    case "#microsoft.graph.accessPackageAssignmentRequest":
                        return NewAccessPackageAssignmentRequest(), nil
                    case "#microsoft.graph.accessPackageAssignmentResourceRole":
                        return NewAccessPackageAssignmentResourceRole(), nil
                    case "#microsoft.graph.accessPackageCatalog":
                        return NewAccessPackageCatalog(), nil
                    case "#microsoft.graph.accessPackageResource":
                        return NewAccessPackageResource(), nil
                    case "#microsoft.graph.accessPackageResourceEnvironment":
                        return NewAccessPackageResourceEnvironment(), nil
                    case "#microsoft.graph.accessPackageResourceRequest":
                        return NewAccessPackageResourceRequest(), nil
                    case "#microsoft.graph.accessPackageResourceRole":
                        return NewAccessPackageResourceRole(), nil
                    case "#microsoft.graph.accessPackageResourceRoleScope":
                        return NewAccessPackageResourceRoleScope(), nil
                    case "#microsoft.graph.accessPackageResourceScope":
                        return NewAccessPackageResourceScope(), nil
                    case "#microsoft.graph.accessPackageSubject":
                        return NewAccessPackageSubject(), nil
                    case "#microsoft.graph.accessReview":
                        return NewAccessReview(), nil
                    case "#microsoft.graph.accessReviewDecision":
                        return NewAccessReviewDecision(), nil
                    case "#microsoft.graph.accessReviewHistoryDefinition":
                        return NewAccessReviewHistoryDefinition(), nil
                    case "#microsoft.graph.accessReviewHistoryInstance":
                        return NewAccessReviewHistoryInstance(), nil
                    case "#microsoft.graph.accessReviewInstance":
                        return NewAccessReviewInstance(), nil
                    case "#microsoft.graph.accessReviewInstanceDecisionItem":
                        return NewAccessReviewInstanceDecisionItem(), nil
                    case "#microsoft.graph.accessReviewPolicy":
                        return NewAccessReviewPolicy(), nil
                    case "#microsoft.graph.accessReviewReviewer":
                        return NewAccessReviewReviewer(), nil
                    case "#microsoft.graph.accessReviewScheduleDefinition":
                        return NewAccessReviewScheduleDefinition(), nil
                    case "#microsoft.graph.accessReviewSet":
                        return NewAccessReviewSet(), nil
                    case "#microsoft.graph.accessReviewStage":
                        return NewAccessReviewStage(), nil
                    case "#microsoft.graph.account":
                        return NewAccount(), nil
                    case "#microsoft.graph.activityHistoryItem":
                        return NewActivityHistoryItem(), nil
                    case "#microsoft.graph.activityStatistics":
                        return NewActivityStatistics(), nil
                    case "#microsoft.graph.adminConsentRequestPolicy":
                        return NewAdminConsentRequestPolicy(), nil
                    case "#microsoft.graph.adminReportSettings":
                        return NewAdminReportSettings(), nil
                    case "#microsoft.graph.advancedThreatProtectionOnboardingDeviceSettingState":
                        return NewAdvancedThreatProtectionOnboardingDeviceSettingState(), nil
                    case "#microsoft.graph.advancedThreatProtectionOnboardingStateSummary":
                        return NewAdvancedThreatProtectionOnboardingStateSummary(), nil
                    case "#microsoft.graph.agedAccountsPayable":
                        return NewAgedAccountsPayable(), nil
                    case "#microsoft.graph.agedAccountsReceivable":
                        return NewAgedAccountsReceivable(), nil
                    case "#microsoft.graph.agreement":
                        return NewAgreement(), nil
                    case "#microsoft.graph.agreementAcceptance":
                        return NewAgreementAcceptance(), nil
                    case "#microsoft.graph.agreementFileProperties":
                        return NewAgreementFileProperties(), nil
                    case "#microsoft.graph.alert":
                        return NewAlert(), nil
                    case "#microsoft.graph.allowedDataLocation":
                        return NewAllowedDataLocation(), nil
                    case "#microsoft.graph.allowedValue":
                        return NewAllowedValue(), nil
                    case "#microsoft.graph.androidDeviceComplianceLocalActionBase":
                        return NewAndroidDeviceComplianceLocalActionBase(), nil
                    case "#microsoft.graph.androidDeviceOwnerEnrollmentProfile":
                        return NewAndroidDeviceOwnerEnrollmentProfile(), nil
                    case "#microsoft.graph.androidForWorkAppConfigurationSchema":
                        return NewAndroidForWorkAppConfigurationSchema(), nil
                    case "#microsoft.graph.androidForWorkEnrollmentProfile":
                        return NewAndroidForWorkEnrollmentProfile(), nil
                    case "#microsoft.graph.androidForWorkSettings":
                        return NewAndroidForWorkSettings(), nil
                    case "#microsoft.graph.androidManagedStoreAccountEnterpriseSettings":
                        return NewAndroidManagedStoreAccountEnterpriseSettings(), nil
                    case "#microsoft.graph.androidManagedStoreAppConfigurationSchema":
                        return NewAndroidManagedStoreAppConfigurationSchema(), nil
                    case "#microsoft.graph.appConsentApprovalRoute":
                        return NewAppConsentApprovalRoute(), nil
                    case "#microsoft.graph.appConsentRequest":
                        return NewAppConsentRequest(), nil
                    case "#microsoft.graph.appleEnrollmentProfileAssignment":
                        return NewAppleEnrollmentProfileAssignment(), nil
                    case "#microsoft.graph.applePushNotificationCertificate":
                        return NewApplePushNotificationCertificate(), nil
                    case "#microsoft.graph.appleUserInitiatedEnrollmentProfile":
                        return NewAppleUserInitiatedEnrollmentProfile(), nil
                    case "#microsoft.graph.applicationSignInDetailedSummary":
                        return NewApplicationSignInDetailedSummary(), nil
                    case "#microsoft.graph.applicationSignInSummary":
                        return NewApplicationSignInSummary(), nil
                    case "#microsoft.graph.applicationTemplate":
                        return NewApplicationTemplate(), nil
                    case "#microsoft.graph.appLogCollectionRequest":
                        return NewAppLogCollectionRequest(), nil
                    case "#microsoft.graph.appRoleAssignment":
                        return NewAppRoleAssignment(), nil
                    case "#microsoft.graph.approval":
                        return NewApproval(), nil
                    case "#microsoft.graph.approvalStep":
                        return NewApprovalStep(), nil
                    case "#microsoft.graph.approvalWorkflowProvider":
                        return NewApprovalWorkflowProvider(), nil
                    case "#microsoft.graph.appScope":
                        return NewAppScope(), nil
                    case "#microsoft.graph.appVulnerabilityManagedDevice":
                        return NewAppVulnerabilityManagedDevice(), nil
                    case "#microsoft.graph.appVulnerabilityMobileApp":
                        return NewAppVulnerabilityMobileApp(), nil
                    case "#microsoft.graph.assignmentFilterEvaluationStatusDetails":
                        return NewAssignmentFilterEvaluationStatusDetails(), nil
                    case "#microsoft.graph.attachment":
                        return NewAttachment(), nil
                    case "#microsoft.graph.attachmentBase":
                        return NewAttachmentBase(), nil
                    case "#microsoft.graph.attachmentSession":
                        return NewAttachmentSession(), nil
                    case "#microsoft.graph.attackSimulationRoot":
                        return NewAttackSimulationRoot(), nil
                    case "#microsoft.graph.attendanceRecord":
                        return NewAttendanceRecord(), nil
                    case "#microsoft.graph.attributeMappingFunctionSchema":
                        return NewAttributeMappingFunctionSchema(), nil
                    case "#microsoft.graph.attributeSet":
                        return NewAttributeSet(), nil
                    case "#microsoft.graph.audioRoutingGroup":
                        return NewAudioRoutingGroup(), nil
                    case "#microsoft.graph.auditEvent":
                        return NewAuditEvent(), nil
                    case "#microsoft.graph.auditLogRoot":
                        return NewAuditLogRoot(), nil
                    case "#microsoft.graph.authentication":
                        return NewAuthentication(), nil
                    case "#microsoft.graph.authenticationContextClassReference":
                        return NewAuthenticationContextClassReference(), nil
                    case "#microsoft.graph.authenticationEventsPolicy":
                        return NewAuthenticationEventsPolicy(), nil
                    case "#microsoft.graph.authenticationFlowsPolicy":
                        return NewAuthenticationFlowsPolicy(), nil
                    case "#microsoft.graph.authenticationListener":
                        return NewAuthenticationListener(), nil
                    case "#microsoft.graph.authenticationMethod":
                        return NewAuthenticationMethod(), nil
                    case "#microsoft.graph.authenticationMethodConfiguration":
                        return NewAuthenticationMethodConfiguration(), nil
                    case "#microsoft.graph.authenticationMethodsPolicy":
                        return NewAuthenticationMethodsPolicy(), nil
                    case "#microsoft.graph.authenticationMethodsRoot":
                        return NewAuthenticationMethodsRoot(), nil
                    case "#microsoft.graph.authenticationMethodTarget":
                        return NewAuthenticationMethodTarget(), nil
                    case "#microsoft.graph.authoredNote":
                        return NewAuthoredNote(), nil
                    case "#microsoft.graph.azureADFeatureUsage":
                        return NewAzureADFeatureUsage(), nil
                    case "#microsoft.graph.azureADLicenseUsage":
                        return NewAzureADLicenseUsage(), nil
                    case "#microsoft.graph.azureADUserFeatureUsage":
                        return NewAzureADUserFeatureUsage(), nil
                    case "#microsoft.graph.b2cAuthenticationMethodsPolicy":
                        return NewB2cAuthenticationMethodsPolicy(), nil
                    case "#microsoft.graph.baseItem":
                        return NewBaseItem(), nil
                    case "#microsoft.graph.baseItemVersion":
                        return NewBaseItemVersion(), nil
                    case "#microsoft.graph.baseTask":
                        return NewBaseTask(), nil
                    case "#microsoft.graph.baseTaskList":
                        return NewBaseTaskList(), nil
                    case "#microsoft.graph.bitlocker":
                        return NewBitlocker(), nil
                    case "#microsoft.graph.bitlockerRecoveryKey":
                        return NewBitlockerRecoveryKey(), nil
                    case "#microsoft.graph.bookingAppointment":
                        return NewBookingAppointment(), nil
                    case "#microsoft.graph.bookingCurrency":
                        return NewBookingCurrency(), nil
                    case "#microsoft.graph.bookingCustomQuestion":
                        return NewBookingCustomQuestion(), nil
                    case "#microsoft.graph.bookingNamedEntity":
                        return NewBookingNamedEntity(), nil
                    case "#microsoft.graph.businessFlow":
                        return NewBusinessFlow(), nil
                    case "#microsoft.graph.businessFlowTemplate":
                        return NewBusinessFlowTemplate(), nil
                    case "#microsoft.graph.calendar":
                        return NewCalendar(), nil
                    case "#microsoft.graph.calendarGroup":
                        return NewCalendarGroup(), nil
                    case "#microsoft.graph.calendarPermission":
                        return NewCalendarPermission(), nil
                    case "#microsoft.graph.call":
                        return NewCall(), nil
                    case "#microsoft.graph.callRecords.callRecord":
                        return iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.NewCallRecord(), nil
                    case "#microsoft.graph.callRecords.segment":
                        return iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.NewSegment(), nil
                    case "#microsoft.graph.callRecords.session":
                        return iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.NewSession(), nil
                    case "#microsoft.graph.cartToClassAssociation":
                        return NewCartToClassAssociation(), nil
                    case "#microsoft.graph.certificateBasedAuthConfiguration":
                        return NewCertificateBasedAuthConfiguration(), nil
                    case "#microsoft.graph.certificateConnectorDetails":
                        return NewCertificateConnectorDetails(), nil
                    case "#microsoft.graph.changeTrackedEntity":
                        return NewChangeTrackedEntity(), nil
                    case "#microsoft.graph.channel":
                        return NewChannel(), nil
                    case "#microsoft.graph.chat":
                        return NewChat(), nil
                    case "#microsoft.graph.chatMessage":
                        return NewChatMessage(), nil
                    case "#microsoft.graph.chatMessageInfo":
                        return NewChatMessageInfo(), nil
                    case "#microsoft.graph.checklistItem":
                        return NewChecklistItem(), nil
                    case "#microsoft.graph.chromeOSOnboardingSettings":
                        return NewChromeOSOnboardingSettings(), nil
                    case "#microsoft.graph.cloudAppSecurityProfile":
                        return NewCloudAppSecurityProfile(), nil
                    case "#microsoft.graph.cloudCommunications":
                        return NewCloudCommunications(), nil
                    case "#microsoft.graph.cloudPC":
                        return NewCloudPC(), nil
                    case "#microsoft.graph.cloudPcAuditEvent":
                        return NewCloudPcAuditEvent(), nil
                    case "#microsoft.graph.cloudPCConnectivityIssue":
                        return NewCloudPCConnectivityIssue(), nil
                    case "#microsoft.graph.cloudPcDeviceImage":
                        return NewCloudPcDeviceImage(), nil
                    case "#microsoft.graph.cloudPcExternalPartnerSetting":
                        return NewCloudPcExternalPartnerSetting(), nil
                    case "#microsoft.graph.cloudPcGalleryImage":
                        return NewCloudPcGalleryImage(), nil
                    case "#microsoft.graph.cloudPcOnPremisesConnection":
                        return NewCloudPcOnPremisesConnection(), nil
                    case "#microsoft.graph.cloudPcOrganizationSettings":
                        return NewCloudPcOrganizationSettings(), nil
                    case "#microsoft.graph.cloudPcProvisioningPolicy":
                        return NewCloudPcProvisioningPolicy(), nil
                    case "#microsoft.graph.cloudPcProvisioningPolicyAssignment":
                        return NewCloudPcProvisioningPolicyAssignment(), nil
                    case "#microsoft.graph.cloudPcServicePlan":
                        return NewCloudPcServicePlan(), nil
                    case "#microsoft.graph.cloudPcSnapshot":
                        return NewCloudPcSnapshot(), nil
                    case "#microsoft.graph.cloudPcSupportedRegion":
                        return NewCloudPcSupportedRegion(), nil
                    case "#microsoft.graph.cloudPcUserSetting":
                        return NewCloudPcUserSetting(), nil
                    case "#microsoft.graph.cloudPcUserSettingAssignment":
                        return NewCloudPcUserSettingAssignment(), nil
                    case "#microsoft.graph.columnDefinition":
                        return NewColumnDefinition(), nil
                    case "#microsoft.graph.columnLink":
                        return NewColumnLink(), nil
                    case "#microsoft.graph.comanagementEligibleDevice":
                        return NewComanagementEligibleDevice(), nil
                    case "#microsoft.graph.command":
                        return NewCommand(), nil
                    case "#microsoft.graph.commsOperation":
                        return NewCommsOperation(), nil
                    case "#microsoft.graph.company":
                        return NewCompany(), nil
                    case "#microsoft.graph.companyInformation":
                        return NewCompanyInformation(), nil
                    case "#microsoft.graph.complianceManagementPartner":
                        return NewComplianceManagementPartner(), nil
                    case "#microsoft.graph.conditionalAccessPolicy":
                        return NewConditionalAccessPolicy(), nil
                    case "#microsoft.graph.conditionalAccessRoot":
                        return NewConditionalAccessRoot(), nil
                    case "#microsoft.graph.configManagerCollection":
                        return NewConfigManagerCollection(), nil
                    case "#microsoft.graph.connectedOrganization":
                        return NewConnectedOrganization(), nil
                    case "#microsoft.graph.connectionOperation":
                        return NewConnectionOperation(), nil
                    case "#microsoft.graph.connector":
                        return NewConnector(), nil
                    case "#microsoft.graph.connectorGroup":
                        return NewConnectorGroup(), nil
                    case "#microsoft.graph.contactFolder":
                        return NewContactFolder(), nil
                    case "#microsoft.graph.contactMergeSuggestions":
                        return NewContactMergeSuggestions(), nil
                    case "#microsoft.graph.contentSharingSession":
                        return NewContentSharingSession(), nil
                    case "#microsoft.graph.contentType":
                        return NewContentType(), nil
                    case "#microsoft.graph.continuousAccessEvaluationPolicy":
                        return NewContinuousAccessEvaluationPolicy(), nil
                    case "#microsoft.graph.conversation":
                        return NewConversation(), nil
                    case "#microsoft.graph.conversationMember":
                        return NewConversationMember(), nil
                    case "#microsoft.graph.conversationThread":
                        return NewConversationThread(), nil
                    case "#microsoft.graph.countryRegion":
                        return NewCountryRegion(), nil
                    case "#microsoft.graph.credentialUsageSummary":
                        return NewCredentialUsageSummary(), nil
                    case "#microsoft.graph.credentialUserRegistrationCount":
                        return NewCredentialUserRegistrationCount(), nil
                    case "#microsoft.graph.credentialUserRegistrationDetails":
                        return NewCredentialUserRegistrationDetails(), nil
                    case "#microsoft.graph.crossTenantAccessPolicyConfigurationDefault":
                        return NewCrossTenantAccessPolicyConfigurationDefault(), nil
                    case "#microsoft.graph.currency":
                        return NewCurrency(), nil
                    case "#microsoft.graph.customCalloutExtension":
                        return NewCustomCalloutExtension(), nil
                    case "#microsoft.graph.customer":
                        return NewCustomer(), nil
                    case "#microsoft.graph.customerPayment":
                        return NewCustomerPayment(), nil
                    case "#microsoft.graph.customerPaymentJournal":
                        return NewCustomerPaymentJournal(), nil
                    case "#microsoft.graph.customExtensionHandler":
                        return NewCustomExtensionHandler(), nil
                    case "#microsoft.graph.customSecurityAttributeDefinition":
                        return NewCustomSecurityAttributeDefinition(), nil
                    case "#microsoft.graph.dataClassificationService":
                        return NewDataClassificationService(), nil
                    case "#microsoft.graph.dataLossPreventionPolicy":
                        return NewDataLossPreventionPolicy(), nil
                    case "#microsoft.graph.dataPolicyOperation":
                        return NewDataPolicyOperation(), nil
                    case "#microsoft.graph.dataSharingConsent":
                        return NewDataSharingConsent(), nil
                    case "#microsoft.graph.defaultUserRoleOverride":
                        return NewDefaultUserRoleOverride(), nil
                    case "#microsoft.graph.delegatedAdminAccessAssignment":
                        return NewDelegatedAdminAccessAssignment(), nil
                    case "#microsoft.graph.delegatedAdminCustomer":
                        return NewDelegatedAdminCustomer(), nil
                    case "#microsoft.graph.delegatedAdminRelationship":
                        return NewDelegatedAdminRelationship(), nil
                    case "#microsoft.graph.delegatedAdminRelationshipOperation":
                        return NewDelegatedAdminRelationshipOperation(), nil
                    case "#microsoft.graph.delegatedAdminRelationshipRequest":
                        return NewDelegatedAdminRelationshipRequest(), nil
                    case "#microsoft.graph.delegatedAdminServiceManagementDetail":
                        return NewDelegatedAdminServiceManagementDetail(), nil
                    case "#microsoft.graph.delegatedPermissionClassification":
                        return NewDelegatedPermissionClassification(), nil
                    case "#microsoft.graph.deletedTeam":
                        return NewDeletedTeam(), nil
                    case "#microsoft.graph.depOnboardingSetting":
                        return NewDepOnboardingSetting(), nil
                    case "#microsoft.graph.detectedApp":
                        return NewDetectedApp(), nil
                    case "#microsoft.graph.deviceAndAppManagementAssignmentFilter":
                        return NewDeviceAndAppManagementAssignmentFilter(), nil
                    case "#microsoft.graph.deviceAppManagement":
                        return NewDeviceAppManagement(), nil
                    case "#microsoft.graph.deviceAppManagementTask":
                        return NewDeviceAppManagementTask(), nil
                    case "#microsoft.graph.deviceCategory":
                        return NewDeviceCategory(), nil
                    case "#microsoft.graph.deviceComplianceActionItem":
                        return NewDeviceComplianceActionItem(), nil
                    case "#microsoft.graph.deviceComplianceDeviceOverview":
                        return NewDeviceComplianceDeviceOverview(), nil
                    case "#microsoft.graph.deviceComplianceDeviceStatus":
                        return NewDeviceComplianceDeviceStatus(), nil
                    case "#microsoft.graph.deviceCompliancePolicy":
                        return NewDeviceCompliancePolicy(), nil
                    case "#microsoft.graph.deviceCompliancePolicyAssignment":
                        return NewDeviceCompliancePolicyAssignment(), nil
                    case "#microsoft.graph.deviceCompliancePolicyDeviceStateSummary":
                        return NewDeviceCompliancePolicyDeviceStateSummary(), nil
                    case "#microsoft.graph.deviceCompliancePolicyGroupAssignment":
                        return NewDeviceCompliancePolicyGroupAssignment(), nil
                    case "#microsoft.graph.deviceCompliancePolicySettingStateSummary":
                        return NewDeviceCompliancePolicySettingStateSummary(), nil
                    case "#microsoft.graph.deviceCompliancePolicyState":
                        return NewDeviceCompliancePolicyState(), nil
                    case "#microsoft.graph.deviceComplianceScheduledActionForRule":
                        return NewDeviceComplianceScheduledActionForRule(), nil
                    case "#microsoft.graph.deviceComplianceScript":
                        return NewDeviceComplianceScript(), nil
                    case "#microsoft.graph.deviceComplianceScriptDeviceState":
                        return NewDeviceComplianceScriptDeviceState(), nil
                    case "#microsoft.graph.deviceComplianceScriptRunSummary":
                        return NewDeviceComplianceScriptRunSummary(), nil
                    case "#microsoft.graph.deviceComplianceSettingState":
                        return NewDeviceComplianceSettingState(), nil
                    case "#microsoft.graph.deviceComplianceUserOverview":
                        return NewDeviceComplianceUserOverview(), nil
                    case "#microsoft.graph.deviceComplianceUserStatus":
                        return NewDeviceComplianceUserStatus(), nil
                    case "#microsoft.graph.deviceConfiguration":
                        return NewDeviceConfiguration(), nil
                    case "#microsoft.graph.deviceConfigurationAssignment":
                        return NewDeviceConfigurationAssignment(), nil
                    case "#microsoft.graph.deviceConfigurationConflictSummary":
                        return NewDeviceConfigurationConflictSummary(), nil
                    case "#microsoft.graph.deviceConfigurationDeviceOverview":
                        return NewDeviceConfigurationDeviceOverview(), nil
                    case "#microsoft.graph.deviceConfigurationDeviceStateSummary":
                        return NewDeviceConfigurationDeviceStateSummary(), nil
                    case "#microsoft.graph.deviceConfigurationDeviceStatus":
                        return NewDeviceConfigurationDeviceStatus(), nil
                    case "#microsoft.graph.deviceConfigurationGroupAssignment":
                        return NewDeviceConfigurationGroupAssignment(), nil
                    case "#microsoft.graph.deviceConfigurationState":
                        return NewDeviceConfigurationState(), nil
                    case "#microsoft.graph.deviceConfigurationUserOverview":
                        return NewDeviceConfigurationUserOverview(), nil
                    case "#microsoft.graph.deviceConfigurationUserStateSummary":
                        return NewDeviceConfigurationUserStateSummary(), nil
                    case "#microsoft.graph.deviceConfigurationUserStatus":
                        return NewDeviceConfigurationUserStatus(), nil
                    case "#microsoft.graph.deviceCustomAttributeShellScript":
                        return NewDeviceCustomAttributeShellScript(), nil
                    case "#microsoft.graph.deviceEnrollmentConfiguration":
                        return NewDeviceEnrollmentConfiguration(), nil
                    case "#microsoft.graph.deviceHealthScript":
                        return NewDeviceHealthScript(), nil
                    case "#microsoft.graph.deviceHealthScriptAssignment":
                        return NewDeviceHealthScriptAssignment(), nil
                    case "#microsoft.graph.deviceHealthScriptDeviceState":
                        return NewDeviceHealthScriptDeviceState(), nil
                    case "#microsoft.graph.deviceHealthScriptRunSummary":
                        return NewDeviceHealthScriptRunSummary(), nil
                    case "#microsoft.graph.deviceInstallState":
                        return NewDeviceInstallState(), nil
                    case "#microsoft.graph.deviceLogCollectionResponse":
                        return NewDeviceLogCollectionResponse(), nil
                    case "#microsoft.graph.deviceManagement":
                        return NewDeviceManagement(), nil
                    case "#microsoft.graph.deviceManagementAutopilotEvent":
                        return NewDeviceManagementAutopilotEvent(), nil
                    case "#microsoft.graph.deviceManagementAutopilotPolicyStatusDetail":
                        return NewDeviceManagementAutopilotPolicyStatusDetail(), nil
                    case "#microsoft.graph.deviceManagementCachedReportConfiguration":
                        return NewDeviceManagementCachedReportConfiguration(), nil
                    case "#microsoft.graph.deviceManagementComplianceActionItem":
                        return NewDeviceManagementComplianceActionItem(), nil
                    case "#microsoft.graph.deviceManagementCompliancePolicy":
                        return NewDeviceManagementCompliancePolicy(), nil
                    case "#microsoft.graph.deviceManagementComplianceScheduledActionForRule":
                        return NewDeviceManagementComplianceScheduledActionForRule(), nil
                    case "#microsoft.graph.deviceManagementConfigurationCategory":
                        return NewDeviceManagementConfigurationCategory(), nil
                    case "#microsoft.graph.deviceManagementConfigurationPolicy":
                        return NewDeviceManagementConfigurationPolicy(), nil
                    case "#microsoft.graph.deviceManagementConfigurationPolicyAssignment":
                        return NewDeviceManagementConfigurationPolicyAssignment(), nil
                    case "#microsoft.graph.deviceManagementConfigurationPolicyTemplate":
                        return NewDeviceManagementConfigurationPolicyTemplate(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSetting":
                        return NewDeviceManagementConfigurationSetting(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSettingDefinition":
                        return NewDeviceManagementConfigurationSettingDefinition(), nil
                    case "#microsoft.graph.deviceManagementConfigurationSettingTemplate":
                        return NewDeviceManagementConfigurationSettingTemplate(), nil
                    case "#microsoft.graph.deviceManagementDerivedCredentialSettings":
                        return NewDeviceManagementDerivedCredentialSettings(), nil
                    case "#microsoft.graph.deviceManagementDomainJoinConnector":
                        return NewDeviceManagementDomainJoinConnector(), nil
                    case "#microsoft.graph.deviceManagementExchangeConnector":
                        return NewDeviceManagementExchangeConnector(), nil
                    case "#microsoft.graph.deviceManagementExchangeOnPremisesPolicy":
                        return NewDeviceManagementExchangeOnPremisesPolicy(), nil
                    case "#microsoft.graph.deviceManagementExportJob":
                        return NewDeviceManagementExportJob(), nil
                    case "#microsoft.graph.deviceManagementIntent":
                        return NewDeviceManagementIntent(), nil
                    case "#microsoft.graph.deviceManagementIntentAssignment":
                        return NewDeviceManagementIntentAssignment(), nil
                    case "#microsoft.graph.deviceManagementIntentDeviceSettingStateSummary":
                        return NewDeviceManagementIntentDeviceSettingStateSummary(), nil
                    case "#microsoft.graph.deviceManagementIntentDeviceState":
                        return NewDeviceManagementIntentDeviceState(), nil
                    case "#microsoft.graph.deviceManagementIntentDeviceStateSummary":
                        return NewDeviceManagementIntentDeviceStateSummary(), nil
                    case "#microsoft.graph.deviceManagementIntentUserState":
                        return NewDeviceManagementIntentUserState(), nil
                    case "#microsoft.graph.deviceManagementIntentUserStateSummary":
                        return NewDeviceManagementIntentUserStateSummary(), nil
                    case "#microsoft.graph.deviceManagementPartner":
                        return NewDeviceManagementPartner(), nil
                    case "#microsoft.graph.deviceManagementReports":
                        return NewDeviceManagementReports(), nil
                    case "#microsoft.graph.deviceManagementResourceAccessProfileAssignment":
                        return NewDeviceManagementResourceAccessProfileAssignment(), nil
                    case "#microsoft.graph.deviceManagementResourceAccessProfileBase":
                        return NewDeviceManagementResourceAccessProfileBase(), nil
                    case "#microsoft.graph.deviceManagementReusablePolicySetting":
                        return NewDeviceManagementReusablePolicySetting(), nil
                    case "#microsoft.graph.deviceManagementScript":
                        return NewDeviceManagementScript(), nil
                    case "#microsoft.graph.deviceManagementScriptAssignment":
                        return NewDeviceManagementScriptAssignment(), nil
                    case "#microsoft.graph.deviceManagementScriptDeviceState":
                        return NewDeviceManagementScriptDeviceState(), nil
                    case "#microsoft.graph.deviceManagementScriptGroupAssignment":
                        return NewDeviceManagementScriptGroupAssignment(), nil
                    case "#microsoft.graph.deviceManagementScriptRunSummary":
                        return NewDeviceManagementScriptRunSummary(), nil
                    case "#microsoft.graph.deviceManagementScriptUserState":
                        return NewDeviceManagementScriptUserState(), nil
                    case "#microsoft.graph.deviceManagementSettingCategory":
                        return NewDeviceManagementSettingCategory(), nil
                    case "#microsoft.graph.deviceManagementSettingDefinition":
                        return NewDeviceManagementSettingDefinition(), nil
                    case "#microsoft.graph.deviceManagementSettingInstance":
                        return NewDeviceManagementSettingInstance(), nil
                    case "#microsoft.graph.deviceManagementTemplate":
                        return NewDeviceManagementTemplate(), nil
                    case "#microsoft.graph.deviceManagementTroubleshootingEvent":
                        return NewDeviceManagementTroubleshootingEvent(), nil
                    case "#microsoft.graph.deviceRegistrationPolicy":
                        return NewDeviceRegistrationPolicy(), nil
                    case "#microsoft.graph.deviceSetupConfiguration":
                        return NewDeviceSetupConfiguration(), nil
                    case "#microsoft.graph.deviceShellScript":
                        return NewDeviceShellScript(), nil
                    case "#microsoft.graph.dimension":
                        return NewDimension(), nil
                    case "#microsoft.graph.dimensionValue":
                        return NewDimensionValue(), nil
                    case "#microsoft.graph.directory":
                        return NewDirectory(), nil
                    case "#microsoft.graph.directoryAudit":
                        return NewDirectoryAudit(), nil
                    case "#microsoft.graph.directoryDefinition":
                        return NewDirectoryDefinition(), nil
                    case "#microsoft.graph.directoryObject":
                        return NewDirectoryObject(), nil
                    case "#microsoft.graph.directoryRoleAccessReviewPolicy":
                        return NewDirectoryRoleAccessReviewPolicy(), nil
                    case "#microsoft.graph.directorySetting":
                        return NewDirectorySetting(), nil
                    case "#microsoft.graph.document":
                        return NewDocument(), nil
                    case "#microsoft.graph.documentComment":
                        return NewDocumentComment(), nil
                    case "#microsoft.graph.documentCommentReply":
                        return NewDocumentCommentReply(), nil
                    case "#microsoft.graph.domain":
                        return NewDomain(), nil
                    case "#microsoft.graph.domainDnsRecord":
                        return NewDomainDnsRecord(), nil
                    case "#microsoft.graph.domainSecurityProfile":
                        return NewDomainSecurityProfile(), nil
                    case "#microsoft.graph.eBookInstallSummary":
                        return NewEBookInstallSummary(), nil
                    case "#microsoft.graph.ediscovery.case":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewCase_escaped(), nil
                    case "#microsoft.graph.ediscovery.caseOperation":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewCaseOperation(), nil
                    case "#microsoft.graph.ediscovery.caseSettings":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewCaseSettings(), nil
                    case "#microsoft.graph.ediscovery.dataSource":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewDataSource(), nil
                    case "#microsoft.graph.ediscovery.dataSourceContainer":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewDataSourceContainer(), nil
                    case "#microsoft.graph.ediscovery.ediscoveryroot":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewEdiscoveryroot(), nil
                    case "#microsoft.graph.ediscovery.legalHold":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewLegalHold(), nil
                    case "#microsoft.graph.ediscovery.reviewSet":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewReviewSet(), nil
                    case "#microsoft.graph.ediscovery.reviewSetQuery":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewReviewSetQuery(), nil
                    case "#microsoft.graph.ediscovery.sourceCollection":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewSourceCollection(), nil
                    case "#microsoft.graph.ediscovery.tag":
                        return ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.NewTag(), nil
                    case "#microsoft.graph.educationAssignment":
                        return NewEducationAssignment(), nil
                    case "#microsoft.graph.educationAssignmentDefaults":
                        return NewEducationAssignmentDefaults(), nil
                    case "#microsoft.graph.educationAssignmentResource":
                        return NewEducationAssignmentResource(), nil
                    case "#microsoft.graph.educationAssignmentSettings":
                        return NewEducationAssignmentSettings(), nil
                    case "#microsoft.graph.educationCategory":
                        return NewEducationCategory(), nil
                    case "#microsoft.graph.educationClass":
                        return NewEducationClass(), nil
                    case "#microsoft.graph.educationOrganization":
                        return NewEducationOrganization(), nil
                    case "#microsoft.graph.educationOutcome":
                        return NewEducationOutcome(), nil
                    case "#microsoft.graph.educationRubric":
                        return NewEducationRubric(), nil
                    case "#microsoft.graph.educationSubmission":
                        return NewEducationSubmission(), nil
                    case "#microsoft.graph.educationSubmissionResource":
                        return NewEducationSubmissionResource(), nil
                    case "#microsoft.graph.educationSynchronizationError":
                        return NewEducationSynchronizationError(), nil
                    case "#microsoft.graph.educationSynchronizationProfile":
                        return NewEducationSynchronizationProfile(), nil
                    case "#microsoft.graph.educationSynchronizationProfileStatus":
                        return NewEducationSynchronizationProfileStatus(), nil
                    case "#microsoft.graph.educationUser":
                        return NewEducationUser(), nil
                    case "#microsoft.graph.embeddedSIMActivationCodePool":
                        return NewEmbeddedSIMActivationCodePool(), nil
                    case "#microsoft.graph.embeddedSIMActivationCodePoolAssignment":
                        return NewEmbeddedSIMActivationCodePoolAssignment(), nil
                    case "#microsoft.graph.embeddedSIMDeviceState":
                        return NewEmbeddedSIMDeviceState(), nil
                    case "#microsoft.graph.employee":
                        return NewEmployee(), nil
                    case "#microsoft.graph.enrollmentConfigurationAssignment":
                        return NewEnrollmentConfigurationAssignment(), nil
                    case "#microsoft.graph.enrollmentProfile":
                        return NewEnrollmentProfile(), nil
                    case "#microsoft.graph.enterpriseCodeSigningCertificate":
                        return NewEnterpriseCodeSigningCertificate(), nil
                    case "#microsoft.graph.entitlementManagement":
                        return NewEntitlementManagement(), nil
                    case "#microsoft.graph.entitlementManagementSettings":
                        return NewEntitlementManagementSettings(), nil
                    case "#microsoft.graph.exactMatchDataStoreBase":
                        return NewExactMatchDataStoreBase(), nil
                    case "#microsoft.graph.exactMatchJobBase":
                        return NewExactMatchJobBase(), nil
                    case "#microsoft.graph.exactMatchUploadAgent":
                        return NewExactMatchUploadAgent(), nil
                    case "#microsoft.graph.extension":
                        return NewExtension(), nil
                    case "#microsoft.graph.external":
                        return NewExternal(), nil
                    case "#microsoft.graph.externalConnection":
                        return NewExternalConnection(), nil
                    case "#microsoft.graph.externalConnectors.connectionOperation":
                        return ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.NewConnectionOperation(), nil
                    case "#microsoft.graph.externalConnectors.connectionQuota":
                        return ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.NewConnectionQuota(), nil
                    case "#microsoft.graph.externalConnectors.externalConnection":
                        return ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.NewExternalConnection(), nil
                    case "#microsoft.graph.externalConnectors.externalGroup":
                        return ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.NewExternalGroup(), nil
                    case "#microsoft.graph.externalConnectors.externalItem":
                        return ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.NewExternalItem(), nil
                    case "#microsoft.graph.externalConnectors.identity":
                        return ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.NewIdentity(), nil
                    case "#microsoft.graph.externalConnectors.schema":
                        return ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.NewSchema(), nil
                    case "#microsoft.graph.externalDomainName":
                        return NewExternalDomainName(), nil
                    case "#microsoft.graph.externalGroup":
                        return NewExternalGroup(), nil
                    case "#microsoft.graph.externalItem":
                        return NewExternalItem(), nil
                    case "#microsoft.graph.featureRolloutPolicy":
                        return NewFeatureRolloutPolicy(), nil
                    case "#microsoft.graph.federatedIdentityCredential":
                        return NewFederatedIdentityCredential(), nil
                    case "#microsoft.graph.fieldValueSet":
                        return NewFieldValueSet(), nil
                    case "#microsoft.graph.fileClassificationRequest":
                        return NewFileClassificationRequest(), nil
                    case "#microsoft.graph.fileSecurityProfile":
                        return NewFileSecurityProfile(), nil
                    case "#microsoft.graph.filterOperatorSchema":
                        return NewFilterOperatorSchema(), nil
                    case "#microsoft.graph.generalLedgerEntry":
                        return NewGeneralLedgerEntry(), nil
                    case "#microsoft.graph.governanceInsight":
                        return NewGovernanceInsight(), nil
                    case "#microsoft.graph.governancePolicyTemplate":
                        return NewGovernancePolicyTemplate(), nil
                    case "#microsoft.graph.governanceResource":
                        return NewGovernanceResource(), nil
                    case "#microsoft.graph.governanceRoleAssignment":
                        return NewGovernanceRoleAssignment(), nil
                    case "#microsoft.graph.governanceRoleAssignmentRequest":
                        return NewGovernanceRoleAssignmentRequest(), nil
                    case "#microsoft.graph.governanceRoleDefinition":
                        return NewGovernanceRoleDefinition(), nil
                    case "#microsoft.graph.governanceRoleSetting":
                        return NewGovernanceRoleSetting(), nil
                    case "#microsoft.graph.governanceSubject":
                        return NewGovernanceSubject(), nil
                    case "#microsoft.graph.groupLifecyclePolicy":
                        return NewGroupLifecyclePolicy(), nil
                    case "#microsoft.graph.groupPolicyCategory":
                        return NewGroupPolicyCategory(), nil
                    case "#microsoft.graph.groupPolicyConfiguration":
                        return NewGroupPolicyConfiguration(), nil
                    case "#microsoft.graph.groupPolicyConfigurationAssignment":
                        return NewGroupPolicyConfigurationAssignment(), nil
                    case "#microsoft.graph.groupPolicyDefinition":
                        return NewGroupPolicyDefinition(), nil
                    case "#microsoft.graph.groupPolicyDefinitionFile":
                        return NewGroupPolicyDefinitionFile(), nil
                    case "#microsoft.graph.groupPolicyDefinitionValue":
                        return NewGroupPolicyDefinitionValue(), nil
                    case "#microsoft.graph.groupPolicyMigrationReport":
                        return NewGroupPolicyMigrationReport(), nil
                    case "#microsoft.graph.groupPolicyObjectFile":
                        return NewGroupPolicyObjectFile(), nil
                    case "#microsoft.graph.groupPolicyOperation":
                        return NewGroupPolicyOperation(), nil
                    case "#microsoft.graph.groupPolicyPresentation":
                        return NewGroupPolicyPresentation(), nil
                    case "#microsoft.graph.groupPolicyPresentationValue":
                        return NewGroupPolicyPresentationValue(), nil
                    case "#microsoft.graph.groupPolicySettingMapping":
                        return NewGroupPolicySettingMapping(), nil
                    case "#microsoft.graph.hostSecurityProfile":
                        return NewHostSecurityProfile(), nil
                    case "#microsoft.graph.identityApiConnector":
                        return NewIdentityApiConnector(), nil
                    case "#microsoft.graph.identityProvider":
                        return NewIdentityProvider(), nil
                    case "#microsoft.graph.identityProviderBase":
                        return NewIdentityProviderBase(), nil
                    case "#microsoft.graph.identityUserFlow":
                        return NewIdentityUserFlow(), nil
                    case "#microsoft.graph.identityUserFlowAttribute":
                        return NewIdentityUserFlowAttribute(), nil
                    case "#microsoft.graph.identityUserFlowAttributeAssignment":
                        return NewIdentityUserFlowAttributeAssignment(), nil
                    case "#microsoft.graph.importedAppleDeviceIdentity":
                        return NewImportedAppleDeviceIdentity(), nil
                    case "#microsoft.graph.importedDeviceIdentity":
                        return NewImportedDeviceIdentity(), nil
                    case "#microsoft.graph.importedWindowsAutopilotDeviceIdentity":
                        return NewImportedWindowsAutopilotDeviceIdentity(), nil
                    case "#microsoft.graph.importedWindowsAutopilotDeviceIdentityUpload":
                        return NewImportedWindowsAutopilotDeviceIdentityUpload(), nil
                    case "#microsoft.graph.inferenceClassification":
                        return NewInferenceClassification(), nil
                    case "#microsoft.graph.inferenceClassificationOverride":
                        return NewInferenceClassificationOverride(), nil
                    case "#microsoft.graph.informationProtection":
                        return NewInformationProtection(), nil
                    case "#microsoft.graph.informationProtectionLabel":
                        return NewInformationProtectionLabel(), nil
                    case "#microsoft.graph.informationProtectionPolicy":
                        return NewInformationProtectionPolicy(), nil
                    case "#microsoft.graph.insightsSettings":
                        return NewInsightsSettings(), nil
                    case "#microsoft.graph.intuneBrandingProfile":
                        return NewIntuneBrandingProfile(), nil
                    case "#microsoft.graph.intuneBrandingProfileAssignment":
                        return NewIntuneBrandingProfileAssignment(), nil
                    case "#microsoft.graph.invitation":
                        return NewInvitation(), nil
                    case "#microsoft.graph.iosLobAppProvisioningConfiguration":
                        return NewIosLobAppProvisioningConfiguration(), nil
                    case "#microsoft.graph.iosLobAppProvisioningConfigurationAssignment":
                        return NewIosLobAppProvisioningConfigurationAssignment(), nil
                    case "#microsoft.graph.iosUpdateDeviceStatus":
                        return NewIosUpdateDeviceStatus(), nil
                    case "#microsoft.graph.iosVppAppAssignedLicense":
                        return NewIosVppAppAssignedLicense(), nil
                    case "#microsoft.graph.ipSecurityProfile":
                        return NewIpSecurityProfile(), nil
                    case "#microsoft.graph.item":
                        return NewItem(), nil
                    case "#microsoft.graph.itemActivity":
                        return NewItemActivity(), nil
                    case "#microsoft.graph.itemActivityOLD":
                        return NewItemActivityOLD(), nil
                    case "#microsoft.graph.itemActivityStat":
                        return NewItemActivityStat(), nil
                    case "#microsoft.graph.itemAnalytics":
                        return NewItemAnalytics(), nil
                    case "#microsoft.graph.itemCategory":
                        return NewItemCategory(), nil
                    case "#microsoft.graph.itemFacet":
                        return NewItemFacet(), nil
                    case "#microsoft.graph.jobResponseBase":
                        return NewJobResponseBase(), nil
                    case "#microsoft.graph.journal":
                        return NewJournal(), nil
                    case "#microsoft.graph.journalLine":
                        return NewJournalLine(), nil
                    case "#microsoft.graph.licenseDetails":
                        return NewLicenseDetails(), nil
                    case "#microsoft.graph.linkedResource":
                        return NewLinkedResource(), nil
                    case "#microsoft.graph.linkedResource_v2":
                        return NewLinkedResource_v2(), nil
                    case "#microsoft.graph.localizedNotificationMessage":
                        return NewLocalizedNotificationMessage(), nil
                    case "#microsoft.graph.longRunningOperation":
                        return NewLongRunningOperation(), nil
                    case "#microsoft.graph.lookupResultRow":
                        return NewLookupResultRow(), nil
                    case "#microsoft.graph.macOSSoftwareUpdateAccountSummary":
                        return NewMacOSSoftwareUpdateAccountSummary(), nil
                    case "#microsoft.graph.macOSSoftwareUpdateCategorySummary":
                        return NewMacOSSoftwareUpdateCategorySummary(), nil
                    case "#microsoft.graph.macOSSoftwareUpdateStateSummary":
                        return NewMacOSSoftwareUpdateStateSummary(), nil
                    case "#microsoft.graph.macOsVppAppAssignedLicense":
                        return NewMacOsVppAppAssignedLicense(), nil
                    case "#microsoft.graph.mailFolder":
                        return NewMailFolder(), nil
                    case "#microsoft.graph.malwareStateForWindowsDevice":
                        return NewMalwareStateForWindowsDevice(), nil
                    case "#microsoft.graph.managedAllDeviceCertificateState":
                        return NewManagedAllDeviceCertificateState(), nil
                    case "#microsoft.graph.managedAppOperation":
                        return NewManagedAppOperation(), nil
                    case "#microsoft.graph.managedAppPolicy":
                        return NewManagedAppPolicy(), nil
                    case "#microsoft.graph.managedAppPolicyDeploymentSummary":
                        return NewManagedAppPolicyDeploymentSummary(), nil
                    case "#microsoft.graph.managedAppRegistration":
                        return NewManagedAppRegistration(), nil
                    case "#microsoft.graph.managedAppStatus":
                        return NewManagedAppStatus(), nil
                    case "#microsoft.graph.managedDevice":
                        return NewManagedDevice(), nil
                    case "#microsoft.graph.managedDeviceCertificateState":
                        return NewManagedDeviceCertificateState(), nil
                    case "#microsoft.graph.managedDeviceEncryptionState":
                        return NewManagedDeviceEncryptionState(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfiguration":
                        return NewManagedDeviceMobileAppConfiguration(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfigurationAssignment":
                        return NewManagedDeviceMobileAppConfigurationAssignment(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfigurationDeviceStatus":
                        return NewManagedDeviceMobileAppConfigurationDeviceStatus(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfigurationDeviceSummary":
                        return NewManagedDeviceMobileAppConfigurationDeviceSummary(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfigurationState":
                        return NewManagedDeviceMobileAppConfigurationState(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfigurationUserStatus":
                        return NewManagedDeviceMobileAppConfigurationUserStatus(), nil
                    case "#microsoft.graph.managedDeviceMobileAppConfigurationUserSummary":
                        return NewManagedDeviceMobileAppConfigurationUserSummary(), nil
                    case "#microsoft.graph.managedDeviceOverview":
                        return NewManagedDeviceOverview(), nil
                    case "#microsoft.graph.managedEBook":
                        return NewManagedEBook(), nil
                    case "#microsoft.graph.managedEBookAssignment":
                        return NewManagedEBookAssignment(), nil
                    case "#microsoft.graph.managedEBookCategory":
                        return NewManagedEBookCategory(), nil
                    case "#microsoft.graph.managedMobileApp":
                        return NewManagedMobileApp(), nil
                    case "#microsoft.graph.managedTenants.aggregatedPolicyCompliance":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewAggregatedPolicyCompliance(), nil
                    case "#microsoft.graph.managedTenants.auditEvent":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewAuditEvent(), nil
                    case "#microsoft.graph.managedTenants.cloudPcConnection":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewCloudPcConnection(), nil
                    case "#microsoft.graph.managedTenants.cloudPcDevice":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewCloudPcDevice(), nil
                    case "#microsoft.graph.managedTenants.cloudPcOverview":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewCloudPcOverview(), nil
                    case "#microsoft.graph.managedTenants.conditionalAccessPolicyCoverage":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewConditionalAccessPolicyCoverage(), nil
                    case "#microsoft.graph.managedTenants.credentialUserRegistrationsSummary":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewCredentialUserRegistrationsSummary(), nil
                    case "#microsoft.graph.managedTenants.deviceCompliancePolicySettingStateSummary":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewDeviceCompliancePolicySettingStateSummary(), nil
                    case "#microsoft.graph.managedTenants.managedDeviceCompliance":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagedDeviceCompliance(), nil
                    case "#microsoft.graph.managedTenants.managedDeviceComplianceTrend":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagedDeviceComplianceTrend(), nil
                    case "#microsoft.graph.managedTenants.managedTenant":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagedTenant(), nil
                    case "#microsoft.graph.managedTenants.managementAction":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagementAction(), nil
                    case "#microsoft.graph.managedTenants.managementActionTenantDeploymentStatus":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagementActionTenantDeploymentStatus(), nil
                    case "#microsoft.graph.managedTenants.managementIntent":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagementIntent(), nil
                    case "#microsoft.graph.managedTenants.managementTemplate":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagementTemplate(), nil
                    case "#microsoft.graph.managedTenants.managementTemplateCollection":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagementTemplateCollection(), nil
                    case "#microsoft.graph.managedTenants.managementTemplateStep":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagementTemplateStep(), nil
                    case "#microsoft.graph.managedTenants.managementTemplateStepDeployment":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagementTemplateStepDeployment(), nil
                    case "#microsoft.graph.managedTenants.managementTemplateStepVersion":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewManagementTemplateStepVersion(), nil
                    case "#microsoft.graph.managedTenants.tenant":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewTenant(), nil
                    case "#microsoft.graph.managedTenants.tenantCustomizedInformation":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewTenantCustomizedInformation(), nil
                    case "#microsoft.graph.managedTenants.tenantDetailedInformation":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewTenantDetailedInformation(), nil
                    case "#microsoft.graph.managedTenants.tenantGroup":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewTenantGroup(), nil
                    case "#microsoft.graph.managedTenants.tenantTag":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewTenantTag(), nil
                    case "#microsoft.graph.managedTenants.windowsDeviceMalwareState":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewWindowsDeviceMalwareState(), nil
                    case "#microsoft.graph.managedTenants.windowsProtectionState":
                        return i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.NewWindowsProtectionState(), nil
                    case "#microsoft.graph.meetingAttendanceReport":
                        return NewMeetingAttendanceReport(), nil
                    case "#microsoft.graph.meetingRegistrantBase":
                        return NewMeetingRegistrantBase(), nil
                    case "#microsoft.graph.meetingRegistrationBase":
                        return NewMeetingRegistrationBase(), nil
                    case "#microsoft.graph.meetingRegistrationQuestion":
                        return NewMeetingRegistrationQuestion(), nil
                    case "#microsoft.graph.mention":
                        return NewMention(), nil
                    case "#microsoft.graph.messageEvent":
                        return NewMessageEvent(), nil
                    case "#microsoft.graph.messageRecipient":
                        return NewMessageRecipient(), nil
                    case "#microsoft.graph.messageRule":
                        return NewMessageRule(), nil
                    case "#microsoft.graph.messageTrace":
                        return NewMessageTrace(), nil
                    case "#microsoft.graph.microsoftApplicationDataAccessSettings":
                        return NewMicrosoftApplicationDataAccessSettings(), nil
                    case "#microsoft.graph.microsoftTunnelConfiguration":
                        return NewMicrosoftTunnelConfiguration(), nil
                    case "#microsoft.graph.microsoftTunnelHealthThreshold":
                        return NewMicrosoftTunnelHealthThreshold(), nil
                    case "#microsoft.graph.microsoftTunnelServer":
                        return NewMicrosoftTunnelServer(), nil
                    case "#microsoft.graph.microsoftTunnelServerLogCollectionResponse":
                        return NewMicrosoftTunnelServerLogCollectionResponse(), nil
                    case "#microsoft.graph.microsoftTunnelSite":
                        return NewMicrosoftTunnelSite(), nil
                    case "#microsoft.graph.mobileApp":
                        return NewMobileApp(), nil
                    case "#microsoft.graph.mobileAppAssignment":
                        return NewMobileAppAssignment(), nil
                    case "#microsoft.graph.mobileAppCategory":
                        return NewMobileAppCategory(), nil
                    case "#microsoft.graph.mobileAppContent":
                        return NewMobileAppContent(), nil
                    case "#microsoft.graph.mobileAppContentFile":
                        return NewMobileAppContentFile(), nil
                    case "#microsoft.graph.mobileAppInstallStatus":
                        return NewMobileAppInstallStatus(), nil
                    case "#microsoft.graph.mobileAppInstallSummary":
                        return NewMobileAppInstallSummary(), nil
                    case "#microsoft.graph.mobileAppIntentAndState":
                        return NewMobileAppIntentAndState(), nil
                    case "#microsoft.graph.mobileAppProvisioningConfigGroupAssignment":
                        return NewMobileAppProvisioningConfigGroupAssignment(), nil
                    case "#microsoft.graph.mobileAppRelationship":
                        return NewMobileAppRelationship(), nil
                    case "#microsoft.graph.mobileContainedApp":
                        return NewMobileContainedApp(), nil
                    case "#microsoft.graph.mobileThreatDefenseConnector":
                        return NewMobileThreatDefenseConnector(), nil
                    case "#microsoft.graph.mobilityManagementPolicy":
                        return NewMobilityManagementPolicy(), nil
                    case "#microsoft.graph.multiValueLegacyExtendedProperty":
                        return NewMultiValueLegacyExtendedProperty(), nil
                    case "#microsoft.graph.namedLocation":
                        return NewNamedLocation(), nil
                    case "#microsoft.graph.ndesConnector":
                        return NewNdesConnector(), nil
                    case "#microsoft.graph.notification":
                        return NewNotification(), nil
                    case "#microsoft.graph.notificationMessageTemplate":
                        return NewNotificationMessageTemplate(), nil
                    case "#microsoft.graph.oAuth2PermissionGrant":
                        return NewOAuth2PermissionGrant(), nil
                    case "#microsoft.graph.oemWarrantyInformationOnboarding":
                        return NewOemWarrantyInformationOnboarding(), nil
                    case "#microsoft.graph.office365ActiveUserCounts":
                        return NewOffice365ActiveUserCounts(), nil
                    case "#microsoft.graph.office365ActiveUserDetail":
                        return NewOffice365ActiveUserDetail(), nil
                    case "#microsoft.graph.office365GroupsActivityCounts":
                        return NewOffice365GroupsActivityCounts(), nil
                    case "#microsoft.graph.office365GroupsActivityDetail":
                        return NewOffice365GroupsActivityDetail(), nil
                    case "#microsoft.graph.office365GroupsActivityFileCounts":
                        return NewOffice365GroupsActivityFileCounts(), nil
                    case "#microsoft.graph.office365GroupsActivityGroupCounts":
                        return NewOffice365GroupsActivityGroupCounts(), nil
                    case "#microsoft.graph.office365GroupsActivityStorage":
                        return NewOffice365GroupsActivityStorage(), nil
                    case "#microsoft.graph.office365ServicesUserCounts":
                        return NewOffice365ServicesUserCounts(), nil
                    case "#microsoft.graph.officeClientConfiguration":
                        return NewOfficeClientConfiguration(), nil
                    case "#microsoft.graph.officeClientConfigurationAssignment":
                        return NewOfficeClientConfigurationAssignment(), nil
                    case "#microsoft.graph.officeGraphInsights":
                        return NewOfficeGraphInsights(), nil
                    case "#microsoft.graph.onenote":
                        return NewOnenote(), nil
                    case "#microsoft.graph.onenoteEntityBaseModel":
                        return NewOnenoteEntityBaseModel(), nil
                    case "#microsoft.graph.onlineMeeting":
                        return NewOnlineMeeting(), nil
                    case "#microsoft.graph.onPremisesAgent":
                        return NewOnPremisesAgent(), nil
                    case "#microsoft.graph.onPremisesAgentGroup":
                        return NewOnPremisesAgentGroup(), nil
                    case "#microsoft.graph.onPremisesConditionalAccessSettings":
                        return NewOnPremisesConditionalAccessSettings(), nil
                    case "#microsoft.graph.onPremisesPublishingProfile":
                        return NewOnPremisesPublishingProfile(), nil
                    case "#microsoft.graph.operation":
                        return NewOperation(), nil
                    case "#microsoft.graph.organizationalBrandingProperties":
                        return NewOrganizationalBrandingProperties(), nil
                    case "#microsoft.graph.organizationSettings":
                        return NewOrganizationSettings(), nil
                    case "#microsoft.graph.outlookCategory":
                        return NewOutlookCategory(), nil
                    case "#microsoft.graph.outlookItem":
                        return NewOutlookItem(), nil
                    case "#microsoft.graph.outlookTaskFolder":
                        return NewOutlookTaskFolder(), nil
                    case "#microsoft.graph.outlookTaskGroup":
                        return NewOutlookTaskGroup(), nil
                    case "#microsoft.graph.outlookUser":
                        return NewOutlookUser(), nil
                    case "#microsoft.graph.participant":
                        return NewParticipant(), nil
                    case "#microsoft.graph.participantJoiningNotification":
                        return NewParticipantJoiningNotification(), nil
                    case "#microsoft.graph.participantLeftNotification":
                        return NewParticipantLeftNotification(), nil
                    case "#microsoft.graph.payloadResponse":
                        return NewPayloadResponse(), nil
                    case "#microsoft.graph.paymentMethod":
                        return NewPaymentMethod(), nil
                    case "#microsoft.graph.paymentTerm":
                        return NewPaymentTerm(), nil
                    case "#microsoft.graph.permission":
                        return NewPermission(), nil
                    case "#microsoft.graph.permissionGrantConditionSet":
                        return NewPermissionGrantConditionSet(), nil
                    case "#microsoft.graph.person":
                        return NewPerson(), nil
                    case "#microsoft.graph.picture":
                        return NewPicture(), nil
                    case "#microsoft.graph.pinnedChatMessageInfo":
                        return NewPinnedChatMessageInfo(), nil
                    case "#microsoft.graph.place":
                        return NewPlace(), nil
                    case "#microsoft.graph.planner":
                        return NewPlanner(), nil
                    case "#microsoft.graph.plannerDelta":
                        return NewPlannerDelta(), nil
                    case "#microsoft.graph.plannerGroup":
                        return NewPlannerGroup(), nil
                    case "#microsoft.graph.plannerRoster":
                        return NewPlannerRoster(), nil
                    case "#microsoft.graph.plannerRosterMember":
                        return NewPlannerRosterMember(), nil
                    case "#microsoft.graph.policySet":
                        return NewPolicySet(), nil
                    case "#microsoft.graph.policySetAssignment":
                        return NewPolicySetAssignment(), nil
                    case "#microsoft.graph.policySetItem":
                        return NewPolicySetItem(), nil
                    case "#microsoft.graph.presence":
                        return NewPresence(), nil
                    case "#microsoft.graph.presentation":
                        return NewPresentation(), nil
                    case "#microsoft.graph.printConnector":
                        return NewPrintConnector(), nil
                    case "#microsoft.graph.printDocument":
                        return NewPrintDocument(), nil
                    case "#microsoft.graph.printerBase":
                        return NewPrinterBase(), nil
                    case "#microsoft.graph.printJob":
                        return NewPrintJob(), nil
                    case "#microsoft.graph.printOperation":
                        return NewPrintOperation(), nil
                    case "#microsoft.graph.printService":
                        return NewPrintService(), nil
                    case "#microsoft.graph.printServiceEndpoint":
                        return NewPrintServiceEndpoint(), nil
                    case "#microsoft.graph.printTask":
                        return NewPrintTask(), nil
                    case "#microsoft.graph.printTaskDefinition":
                        return NewPrintTaskDefinition(), nil
                    case "#microsoft.graph.printTaskTrigger":
                        return NewPrintTaskTrigger(), nil
                    case "#microsoft.graph.printUsage":
                        return NewPrintUsage(), nil
                    case "#microsoft.graph.privilegedAccess":
                        return NewPrivilegedAccess(), nil
                    case "#microsoft.graph.privilegedApproval":
                        return NewPrivilegedApproval(), nil
                    case "#microsoft.graph.privilegedOperationEvent":
                        return NewPrivilegedOperationEvent(), nil
                    case "#microsoft.graph.privilegedRole":
                        return NewPrivilegedRole(), nil
                    case "#microsoft.graph.privilegedRoleAssignment":
                        return NewPrivilegedRoleAssignment(), nil
                    case "#microsoft.graph.privilegedRoleAssignmentRequest":
                        return NewPrivilegedRoleAssignmentRequest(), nil
                    case "#microsoft.graph.privilegedRoleSettings":
                        return NewPrivilegedRoleSettings(), nil
                    case "#microsoft.graph.privilegedRoleSummary":
                        return NewPrivilegedRoleSummary(), nil
                    case "#microsoft.graph.privilegedSignupStatus":
                        return NewPrivilegedSignupStatus(), nil
                    case "#microsoft.graph.profile":
                        return NewProfile(), nil
                    case "#microsoft.graph.profileCardProperty":
                        return NewProfileCardProperty(), nil
                    case "#microsoft.graph.profilePhoto":
                        return NewProfilePhoto(), nil
                    case "#microsoft.graph.program":
                        return NewProgram(), nil
                    case "#microsoft.graph.programControl":
                        return NewProgramControl(), nil
                    case "#microsoft.graph.programControlType":
                        return NewProgramControlType(), nil
                    case "#microsoft.graph.providerTenantSetting":
                        return NewProviderTenantSetting(), nil
                    case "#microsoft.graph.provisioningObjectSummary":
                        return NewProvisioningObjectSummary(), nil
                    case "#microsoft.graph.publishedResource":
                        return NewPublishedResource(), nil
                    case "#microsoft.graph.purchaseInvoice":
                        return NewPurchaseInvoice(), nil
                    case "#microsoft.graph.purchaseInvoiceLine":
                        return NewPurchaseInvoiceLine(), nil
                    case "#microsoft.graph.rbacApplication":
                        return NewRbacApplication(), nil
                    case "#microsoft.graph.rbacApplicationMultiple":
                        return NewRbacApplicationMultiple(), nil
                    case "#microsoft.graph.recommendation":
                        return NewRecommendation(), nil
                    case "#microsoft.graph.recommendationResource":
                        return NewRecommendationResource(), nil
                    case "#microsoft.graph.regionalAndLanguageSettings":
                        return NewRegionalAndLanguageSettings(), nil
                    case "#microsoft.graph.relyingPartyDetailedSummary":
                        return NewRelyingPartyDetailedSummary(), nil
                    case "#microsoft.graph.remoteActionAudit":
                        return NewRemoteActionAudit(), nil
                    case "#microsoft.graph.remoteAssistancePartner":
                        return NewRemoteAssistancePartner(), nil
                    case "#microsoft.graph.remoteAssistanceSettings":
                        return NewRemoteAssistanceSettings(), nil
                    case "#microsoft.graph.reportRoot":
                        return NewReportRoot(), nil
                    case "#microsoft.graph.request":
                        return NewRequest(), nil
                    case "#microsoft.graph.resourceOperation":
                        return NewResourceOperation(), nil
                    case "#microsoft.graph.restrictedAppsViolation":
                        return NewRestrictedAppsViolation(), nil
                    case "#microsoft.graph.riskDetection":
                        return NewRiskDetection(), nil
                    case "#microsoft.graph.riskyServicePrincipal":
                        return NewRiskyServicePrincipal(), nil
                    case "#microsoft.graph.riskyUser":
                        return NewRiskyUser(), nil
                    case "#microsoft.graph.roleAssignment":
                        return NewRoleAssignment(), nil
                    case "#microsoft.graph.roleDefinition":
                        return NewRoleDefinition(), nil
                    case "#microsoft.graph.roleScopeTag":
                        return NewRoleScopeTag(), nil
                    case "#microsoft.graph.roleScopeTagAutoAssignment":
                        return NewRoleScopeTagAutoAssignment(), nil
                    case "#microsoft.graph.salesCreditMemo":
                        return NewSalesCreditMemo(), nil
                    case "#microsoft.graph.salesCreditMemoLine":
                        return NewSalesCreditMemoLine(), nil
                    case "#microsoft.graph.salesInvoice":
                        return NewSalesInvoice(), nil
                    case "#microsoft.graph.salesInvoiceLine":
                        return NewSalesInvoiceLine(), nil
                    case "#microsoft.graph.salesOrder":
                        return NewSalesOrder(), nil
                    case "#microsoft.graph.salesOrderLine":
                        return NewSalesOrderLine(), nil
                    case "#microsoft.graph.salesQuote":
                        return NewSalesQuote(), nil
                    case "#microsoft.graph.salesQuoteLine":
                        return NewSalesQuoteLine(), nil
                    case "#microsoft.graph.schedule":
                        return NewSchedule(), nil
                    case "#microsoft.graph.schema":
                        return NewSchema(), nil
                    case "#microsoft.graph.schemaExtension":
                        return NewSchemaExtension(), nil
                    case "#microsoft.graph.scopedRoleMembership":
                        return NewScopedRoleMembership(), nil
                    case "#microsoft.graph.search.searchAnswer":
                        return iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.NewSearchAnswer(), nil
                    case "#microsoft.graph.searchEntity":
                        return NewSearchEntity(), nil
                    case "#microsoft.graph.secureScore":
                        return NewSecureScore(), nil
                    case "#microsoft.graph.secureScoreControlProfile":
                        return NewSecureScoreControlProfile(), nil
                    case "#microsoft.graph.security":
                        return NewSecurity(), nil
                    case "#microsoft.graph.security.alert":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewAlert(), nil
                    case "#microsoft.graph.security.case":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewCase_escaped(), nil
                    case "#microsoft.graph.security.caseOperation":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewCaseOperation(), nil
                    case "#microsoft.graph.security.casesRoot":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewCasesRoot(), nil
                    case "#microsoft.graph.security.dataSet":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewDataSet(), nil
                    case "#microsoft.graph.security.dataSource":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewDataSource(), nil
                    case "#microsoft.graph.security.dataSourceContainer":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewDataSourceContainer(), nil
                    case "#microsoft.graph.security.dispositionReviewStage":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewDispositionReviewStage(), nil
                    case "#microsoft.graph.security.ediscoveryCaseSettings":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewEdiscoveryCaseSettings(), nil
                    case "#microsoft.graph.security.file":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewFile(), nil
                    case "#microsoft.graph.security.incident":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewIncident(), nil
                    case "#microsoft.graph.security.informationProtection":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewInformationProtection(), nil
                    case "#microsoft.graph.security.informationProtectionPolicySetting":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewInformationProtectionPolicySetting(), nil
                    case "#microsoft.graph.security.labelsRoot":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewLabelsRoot(), nil
                    case "#microsoft.graph.security.policyBase":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewPolicyBase(), nil
                    case "#microsoft.graph.security.retentionEvent":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewRetentionEvent(), nil
                    case "#microsoft.graph.security.retentionEventType":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewRetentionEventType(), nil
                    case "#microsoft.graph.security.retentionLabel":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewRetentionLabel(), nil
                    case "#microsoft.graph.security.search":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewSearch(), nil
                    case "#microsoft.graph.security.security":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewSecurity(), nil
                    case "#microsoft.graph.security.sensitivityLabel":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewSensitivityLabel(), nil
                    case "#microsoft.graph.security.tag":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewTag(), nil
                    case "#microsoft.graph.security.triggersRoot":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewTriggersRoot(), nil
                    case "#microsoft.graph.security.triggerTypesRoot":
                        return i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.NewTriggerTypesRoot(), nil
                    case "#microsoft.graph.securityAction":
                        return NewSecurityAction(), nil
                    case "#microsoft.graph.securityBaselineDeviceState":
                        return NewSecurityBaselineDeviceState(), nil
                    case "#microsoft.graph.securityBaselineSettingState":
                        return NewSecurityBaselineSettingState(), nil
                    case "#microsoft.graph.securityBaselineState":
                        return NewSecurityBaselineState(), nil
                    case "#microsoft.graph.securityBaselineStateSummary":
                        return NewSecurityBaselineStateSummary(), nil
                    case "#microsoft.graph.securityReportsRoot":
                        return NewSecurityReportsRoot(), nil
                    case "#microsoft.graph.sensitiveType":
                        return NewSensitiveType(), nil
                    case "#microsoft.graph.sensitivityLabel":
                        return NewSensitivityLabel(), nil
                    case "#microsoft.graph.sensitivityPolicySettings":
                        return NewSensitivityPolicySettings(), nil
                    case "#microsoft.graph.serviceAnnouncement":
                        return NewServiceAnnouncement(), nil
                    case "#microsoft.graph.serviceAnnouncementAttachment":
                        return NewServiceAnnouncementAttachment(), nil
                    case "#microsoft.graph.serviceAnnouncementBase":
                        return NewServiceAnnouncementBase(), nil
                    case "#microsoft.graph.serviceHealth":
                        return NewServiceHealth(), nil
                    case "#microsoft.graph.servicePrincipalCreationConditionSet":
                        return NewServicePrincipalCreationConditionSet(), nil
                    case "#microsoft.graph.servicePrincipalRiskDetection":
                        return NewServicePrincipalRiskDetection(), nil
                    case "#microsoft.graph.settingStateDeviceSummary":
                        return NewSettingStateDeviceSummary(), nil
                    case "#microsoft.graph.sharedEmailDomain":
                        return NewSharedEmailDomain(), nil
                    case "#microsoft.graph.sharedEmailDomainInvitation":
                        return NewSharedEmailDomainInvitation(), nil
                    case "#microsoft.graph.sharedInsight":
                        return NewSharedInsight(), nil
                    case "#microsoft.graph.shipmentMethod":
                        return NewShipmentMethod(), nil
                    case "#microsoft.graph.sideLoadingKey":
                        return NewSideLoadingKey(), nil
                    case "#microsoft.graph.signIn":
                        return NewSignIn(), nil
                    case "#microsoft.graph.simulation":
                        return NewSimulation(), nil
                    case "#microsoft.graph.simulationAutomation":
                        return NewSimulationAutomation(), nil
                    case "#microsoft.graph.simulationAutomationRun":
                        return NewSimulationAutomationRun(), nil
                    case "#microsoft.graph.singleValueLegacyExtendedProperty":
                        return NewSingleValueLegacyExtendedProperty(), nil
                    case "#microsoft.graph.softwareUpdateStatusSummary":
                        return NewSoftwareUpdateStatusSummary(), nil
                    case "#microsoft.graph.strongAuthenticationDetail":
                        return NewStrongAuthenticationDetail(), nil
                    case "#microsoft.graph.strongAuthenticationPhoneAppDetail":
                        return NewStrongAuthenticationPhoneAppDetail(), nil
                    case "#microsoft.graph.subjectRightsRequest":
                        return NewSubjectRightsRequest(), nil
                    case "#microsoft.graph.subscribedSku":
                        return NewSubscribedSku(), nil
                    case "#microsoft.graph.subscription":
                        return NewSubscription(), nil
                    case "#microsoft.graph.symantecCodeSigningCertificate":
                        return NewSymantecCodeSigningCertificate(), nil
                    case "#microsoft.graph.synchronization":
                        return NewSynchronization(), nil
                    case "#microsoft.graph.synchronizationJob":
                        return NewSynchronizationJob(), nil
                    case "#microsoft.graph.synchronizationSchema":
                        return NewSynchronizationSchema(), nil
                    case "#microsoft.graph.synchronizationTemplate":
                        return NewSynchronizationTemplate(), nil
                    case "#microsoft.graph.targetedManagedAppPolicyAssignment":
                        return NewTargetedManagedAppPolicyAssignment(), nil
                    case "#microsoft.graph.tasks":
                        return NewTasks(), nil
                    case "#microsoft.graph.taxArea":
                        return NewTaxArea(), nil
                    case "#microsoft.graph.taxGroup":
                        return NewTaxGroup(), nil
                    case "#microsoft.graph.team":
                        return NewTeam(), nil
                    case "#microsoft.graph.teamInfo":
                        return NewTeamInfo(), nil
                    case "#microsoft.graph.teamsApp":
                        return NewTeamsApp(), nil
                    case "#microsoft.graph.teamsAppDefinition":
                        return NewTeamsAppDefinition(), nil
                    case "#microsoft.graph.teamsAppIcon":
                        return NewTeamsAppIcon(), nil
                    case "#microsoft.graph.teamsAppInstallation":
                        return NewTeamsAppInstallation(), nil
                    case "#microsoft.graph.teamsAppSettings":
                        return NewTeamsAppSettings(), nil
                    case "#microsoft.graph.teamsAsyncOperation":
                        return NewTeamsAsyncOperation(), nil
                    case "#microsoft.graph.teamsTab":
                        return NewTeamsTab(), nil
                    case "#microsoft.graph.teamsTemplate":
                        return NewTeamsTemplate(), nil
                    case "#microsoft.graph.teamwork":
                        return NewTeamwork(), nil
                    case "#microsoft.graph.teamworkBot":
                        return NewTeamworkBot(), nil
                    case "#microsoft.graph.teamworkDevice":
                        return NewTeamworkDevice(), nil
                    case "#microsoft.graph.teamworkDeviceActivity":
                        return NewTeamworkDeviceActivity(), nil
                    case "#microsoft.graph.teamworkDeviceConfiguration":
                        return NewTeamworkDeviceConfiguration(), nil
                    case "#microsoft.graph.teamworkDeviceHealth":
                        return NewTeamworkDeviceHealth(), nil
                    case "#microsoft.graph.teamworkDeviceOperation":
                        return NewTeamworkDeviceOperation(), nil
                    case "#microsoft.graph.teamworkHostedContent":
                        return NewTeamworkHostedContent(), nil
                    case "#microsoft.graph.teamworkPeripheral":
                        return NewTeamworkPeripheral(), nil
                    case "#microsoft.graph.teamworkTag":
                        return NewTeamworkTag(), nil
                    case "#microsoft.graph.teamworkTagMember":
                        return NewTeamworkTagMember(), nil
                    case "#microsoft.graph.telecomExpenseManagementPartner":
                        return NewTelecomExpenseManagementPartner(), nil
                    case "#microsoft.graph.tenantAdmin.settings":
                        return i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.NewSettings(), nil
                    case "#microsoft.graph.tenantAdmin.sharepoint":
                        return i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.NewSharepoint(), nil
                    case "#microsoft.graph.tenantAttachRBAC":
                        return NewTenantAttachRBAC(), nil
                    case "#microsoft.graph.tenantRelationship":
                        return NewTenantRelationship(), nil
                    case "#microsoft.graph.tenantSetupInfo":
                        return NewTenantSetupInfo(), nil
                    case "#microsoft.graph.termsAndConditions":
                        return NewTermsAndConditions(), nil
                    case "#microsoft.graph.termsAndConditionsAcceptanceStatus":
                        return NewTermsAndConditionsAcceptanceStatus(), nil
                    case "#microsoft.graph.termsAndConditionsAssignment":
                        return NewTermsAndConditionsAssignment(), nil
                    case "#microsoft.graph.termsAndConditionsGroupAssignment":
                        return NewTermsAndConditionsGroupAssignment(), nil
                    case "#microsoft.graph.termsOfUseContainer":
                        return NewTermsOfUseContainer(), nil
                    case "#microsoft.graph.termStore.group":
                        return i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.NewGroup(), nil
                    case "#microsoft.graph.termStore.relation":
                        return i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.NewRelation(), nil
                    case "#microsoft.graph.termStore.set":
                        return i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.NewSet(), nil
                    case "#microsoft.graph.termStore.store":
                        return i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.NewStore(), nil
                    case "#microsoft.graph.termStore.term":
                        return i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.NewTerm(), nil
                    case "#microsoft.graph.textClassificationRequest":
                        return NewTextClassificationRequest(), nil
                    case "#microsoft.graph.threatAssessmentRequest":
                        return NewThreatAssessmentRequest(), nil
                    case "#microsoft.graph.threatAssessmentResult":
                        return NewThreatAssessmentResult(), nil
                    case "#microsoft.graph.thumbnailSet":
                        return NewThumbnailSet(), nil
                    case "#microsoft.graph.tiIndicator":
                        return NewTiIndicator(), nil
                    case "#microsoft.graph.todo":
                        return NewTodo(), nil
                    case "#microsoft.graph.todoTask":
                        return NewTodoTask(), nil
                    case "#microsoft.graph.todoTaskList":
                        return NewTodoTaskList(), nil
                    case "#microsoft.graph.trending":
                        return NewTrending(), nil
                    case "#microsoft.graph.trustFrameworkKeySet":
                        return NewTrustFrameworkKeySet(), nil
                    case "#microsoft.graph.trustFrameworkPolicy":
                        return NewTrustFrameworkPolicy(), nil
                    case "#microsoft.graph.unifiedRbacResourceAction":
                        return NewUnifiedRbacResourceAction(), nil
                    case "#microsoft.graph.unifiedRbacResourceNamespace":
                        return NewUnifiedRbacResourceNamespace(), nil
                    case "#microsoft.graph.unifiedRbacResourceScope":
                        return NewUnifiedRbacResourceScope(), nil
                    case "#microsoft.graph.unifiedRoleAssignment":
                        return NewUnifiedRoleAssignment(), nil
                    case "#microsoft.graph.unifiedRoleAssignmentMultiple":
                        return NewUnifiedRoleAssignmentMultiple(), nil
                    case "#microsoft.graph.unifiedRoleDefinition":
                        return NewUnifiedRoleDefinition(), nil
                    case "#microsoft.graph.unifiedRoleManagementPolicy":
                        return NewUnifiedRoleManagementPolicy(), nil
                    case "#microsoft.graph.unifiedRoleManagementPolicyAssignment":
                        return NewUnifiedRoleManagementPolicyAssignment(), nil
                    case "#microsoft.graph.unifiedRoleManagementPolicyRule":
                        return NewUnifiedRoleManagementPolicyRule(), nil
                    case "#microsoft.graph.unifiedRoleScheduleBase":
                        return NewUnifiedRoleScheduleBase(), nil
                    case "#microsoft.graph.unifiedRoleScheduleInstanceBase":
                        return NewUnifiedRoleScheduleInstanceBase(), nil
                    case "#microsoft.graph.unitOfMeasure":
                        return NewUnitOfMeasure(), nil
                    case "#microsoft.graph.unsupportedGroupPolicyExtension":
                        return NewUnsupportedGroupPolicyExtension(), nil
                    case "#microsoft.graph.usageRight":
                        return NewUsageRight(), nil
                    case "#microsoft.graph.usedInsight":
                        return NewUsedInsight(), nil
                    case "#microsoft.graph.userActivity":
                        return NewUserActivity(), nil
                    case "#microsoft.graph.userAnalytics":
                        return NewUserAnalytics(), nil
                    case "#microsoft.graph.userAppInstallStatus":
                        return NewUserAppInstallStatus(), nil
                    case "#microsoft.graph.userConfiguration":
                        return NewUserConfiguration(), nil
                    case "#microsoft.graph.userCredentialUsageDetails":
                        return NewUserCredentialUsageDetails(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthApplicationPerformance":
                        return NewUserExperienceAnalyticsAppHealthApplicationPerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersion":
                        return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails":
                        return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId":
                        return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByOSVersion":
                        return NewUserExperienceAnalyticsAppHealthAppPerformanceByOSVersion(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthDeviceModelPerformance":
                        return NewUserExperienceAnalyticsAppHealthDeviceModelPerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthDevicePerformance":
                        return NewUserExperienceAnalyticsAppHealthDevicePerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthDevicePerformanceDetails":
                        return NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails(), nil
                    case "#microsoft.graph.userExperienceAnalyticsAppHealthOSVersionPerformance":
                        return NewUserExperienceAnalyticsAppHealthOSVersionPerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBaseline":
                        return NewUserExperienceAnalyticsBaseline(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBatteryHealthAppImpact":
                        return NewUserExperienceAnalyticsBatteryHealthAppImpact(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBatteryHealthCapacityDetails":
                        return NewUserExperienceAnalyticsBatteryHealthCapacityDetails(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBatteryHealthDeviceAppImpact":
                        return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBatteryHealthDevicePerformance":
                        return NewUserExperienceAnalyticsBatteryHealthDevicePerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory":
                        return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBatteryHealthModelPerformance":
                        return NewUserExperienceAnalyticsBatteryHealthModelPerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBatteryHealthOsPerformance":
                        return NewUserExperienceAnalyticsBatteryHealthOsPerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsBatteryHealthRuntimeDetails":
                        return NewUserExperienceAnalyticsBatteryHealthRuntimeDetails(), nil
                    case "#microsoft.graph.userExperienceAnalyticsCategory":
                        return NewUserExperienceAnalyticsCategory(), nil
                    case "#microsoft.graph.userExperienceAnalyticsDevicePerformance":
                        return NewUserExperienceAnalyticsDevicePerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsDeviceScope":
                        return NewUserExperienceAnalyticsDeviceScope(), nil
                    case "#microsoft.graph.userExperienceAnalyticsDeviceScores":
                        return NewUserExperienceAnalyticsDeviceScores(), nil
                    case "#microsoft.graph.userExperienceAnalyticsDeviceStartupHistory":
                        return NewUserExperienceAnalyticsDeviceStartupHistory(), nil
                    case "#microsoft.graph.userExperienceAnalyticsDeviceStartupProcess":
                        return NewUserExperienceAnalyticsDeviceStartupProcess(), nil
                    case "#microsoft.graph.userExperienceAnalyticsDeviceStartupProcessPerformance":
                        return NewUserExperienceAnalyticsDeviceStartupProcessPerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsDeviceWithoutCloudIdentity":
                        return NewUserExperienceAnalyticsDeviceWithoutCloudIdentity(), nil
                    case "#microsoft.graph.userExperienceAnalyticsImpactingProcess":
                        return NewUserExperienceAnalyticsImpactingProcess(), nil
                    case "#microsoft.graph.userExperienceAnalyticsMetric":
                        return NewUserExperienceAnalyticsMetric(), nil
                    case "#microsoft.graph.userExperienceAnalyticsMetricHistory":
                        return NewUserExperienceAnalyticsMetricHistory(), nil
                    case "#microsoft.graph.userExperienceAnalyticsModelScores":
                        return NewUserExperienceAnalyticsModelScores(), nil
                    case "#microsoft.graph.userExperienceAnalyticsNotAutopilotReadyDevice":
                        return NewUserExperienceAnalyticsNotAutopilotReadyDevice(), nil
                    case "#microsoft.graph.userExperienceAnalyticsOverview":
                        return NewUserExperienceAnalyticsOverview(), nil
                    case "#microsoft.graph.userExperienceAnalyticsRegressionSummary":
                        return NewUserExperienceAnalyticsRegressionSummary(), nil
                    case "#microsoft.graph.userExperienceAnalyticsRemoteConnection":
                        return NewUserExperienceAnalyticsRemoteConnection(), nil
                    case "#microsoft.graph.userExperienceAnalyticsResourcePerformance":
                        return NewUserExperienceAnalyticsResourcePerformance(), nil
                    case "#microsoft.graph.userExperienceAnalyticsScoreHistory":
                        return NewUserExperienceAnalyticsScoreHistory(), nil
                    case "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereDevice":
                        return NewUserExperienceAnalyticsWorkFromAnywhereDevice(), nil
                    case "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric":
                        return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(), nil
                    case "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetric":
                        return NewUserExperienceAnalyticsWorkFromAnywhereMetric(), nil
                    case "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereModelPerformance":
                        return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance(), nil
                    case "#microsoft.graph.userFlowLanguageConfiguration":
                        return NewUserFlowLanguageConfiguration(), nil
                    case "#microsoft.graph.userFlowLanguagePage":
                        return NewUserFlowLanguagePage(), nil
                    case "#microsoft.graph.userInsightsSettings":
                        return NewUserInsightsSettings(), nil
                    case "#microsoft.graph.userInstallStateSummary":
                        return NewUserInstallStateSummary(), nil
                    case "#microsoft.graph.userPFXCertificate":
                        return NewUserPFXCertificate(), nil
                    case "#microsoft.graph.userRegistrationDetails":
                        return NewUserRegistrationDetails(), nil
                    case "#microsoft.graph.userSecurityProfile":
                        return NewUserSecurityProfile(), nil
                    case "#microsoft.graph.userSettings":
                        return NewUserSettings(), nil
                    case "#microsoft.graph.userTeamwork":
                        return NewUserTeamwork(), nil
                    case "#microsoft.graph.vendor":
                        return NewVendor_escaped(), nil
                    case "#microsoft.graph.virtualEndpoint":
                        return NewVirtualEndpoint(), nil
                    case "#microsoft.graph.vppToken":
                        return NewVppToken(), nil
                    case "#microsoft.graph.vulnerableManagedDevice":
                        return NewVulnerableManagedDevice(), nil
                    case "#microsoft.graph.windowsAssignedAccessProfile":
                        return NewWindowsAssignedAccessProfile(), nil
                    case "#microsoft.graph.windowsAutopilotDeploymentProfile":
                        return NewWindowsAutopilotDeploymentProfile(), nil
                    case "#microsoft.graph.windowsAutopilotDeploymentProfileAssignment":
                        return NewWindowsAutopilotDeploymentProfileAssignment(), nil
                    case "#microsoft.graph.windowsAutopilotDeviceIdentity":
                        return NewWindowsAutopilotDeviceIdentity(), nil
                    case "#microsoft.graph.windowsAutopilotSettings":
                        return NewWindowsAutopilotSettings(), nil
                    case "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy":
                        return NewWindowsDefenderApplicationControlSupplementalPolicy(), nil
                    case "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyAssignment":
                        return NewWindowsDefenderApplicationControlSupplementalPolicyAssignment(), nil
                    case "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus":
                        return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus(), nil
                    case "#microsoft.graph.windowsDefenderApplicationControlSupplementalPolicyDeploymentSummary":
                        return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary(), nil
                    case "#microsoft.graph.windowsDeviceMalwareState":
                        return NewWindowsDeviceMalwareState(), nil
                    case "#microsoft.graph.windowsDriverUpdateInventory":
                        return NewWindowsDriverUpdateInventory(), nil
                    case "#microsoft.graph.windowsDriverUpdateProfile":
                        return NewWindowsDriverUpdateProfile(), nil
                    case "#microsoft.graph.windowsDriverUpdateProfileAssignment":
                        return NewWindowsDriverUpdateProfileAssignment(), nil
                    case "#microsoft.graph.windowsFeatureUpdateProfile":
                        return NewWindowsFeatureUpdateProfile(), nil
                    case "#microsoft.graph.windowsFeatureUpdateProfileAssignment":
                        return NewWindowsFeatureUpdateProfileAssignment(), nil
                    case "#microsoft.graph.windowsInformationProtectionAppLearningSummary":
                        return NewWindowsInformationProtectionAppLearningSummary(), nil
                    case "#microsoft.graph.windowsInformationProtectionAppLockerFile":
                        return NewWindowsInformationProtectionAppLockerFile(), nil
                    case "#microsoft.graph.windowsInformationProtectionDeviceRegistration":
                        return NewWindowsInformationProtectionDeviceRegistration(), nil
                    case "#microsoft.graph.windowsInformationProtectionNetworkLearningSummary":
                        return NewWindowsInformationProtectionNetworkLearningSummary(), nil
                    case "#microsoft.graph.windowsInformationProtectionWipeAction":
                        return NewWindowsInformationProtectionWipeAction(), nil
                    case "#microsoft.graph.windowsMalwareInformation":
                        return NewWindowsMalwareInformation(), nil
                    case "#microsoft.graph.windowsManagementApp":
                        return NewWindowsManagementApp(), nil
                    case "#microsoft.graph.windowsManagementAppHealthState":
                        return NewWindowsManagementAppHealthState(), nil
                    case "#microsoft.graph.windowsManagementAppHealthSummary":
                        return NewWindowsManagementAppHealthSummary(), nil
                    case "#microsoft.graph.windowsPrivacyDataAccessControlItem":
                        return NewWindowsPrivacyDataAccessControlItem(), nil
                    case "#microsoft.graph.windowsProtectionState":
                        return NewWindowsProtectionState(), nil
                    case "#microsoft.graph.windowsQualityUpdateProfile":
                        return NewWindowsQualityUpdateProfile(), nil
                    case "#microsoft.graph.windowsQualityUpdateProfileAssignment":
                        return NewWindowsQualityUpdateProfileAssignment(), nil
                    case "#microsoft.graph.windowsUpdateCatalogItem":
                        return NewWindowsUpdateCatalogItem(), nil
                    case "#microsoft.graph.windowsUpdates.catalog":
                        return i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.NewCatalog(), nil
                    case "#microsoft.graph.windowsUpdates.catalogEntry":
                        return i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.NewCatalogEntry(), nil
                    case "#microsoft.graph.windowsUpdates.deployment":
                        return i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.NewDeployment(), nil
                    case "#microsoft.graph.windowsUpdates.deploymentAudience":
                        return i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.NewDeploymentAudience(), nil
                    case "#microsoft.graph.windowsUpdates.resourceConnection":
                        return i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.NewResourceConnection(), nil
                    case "#microsoft.graph.windowsUpdates.updatableAsset":
                        return i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.NewUpdatableAsset(), nil
                    case "#microsoft.graph.windowsUpdates.updates":
                        return i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.NewUpdates(), nil
                    case "#microsoft.graph.windowsUpdates.windows":
                        return i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.NewWindows(), nil
                    case "#microsoft.graph.windowsUpdateState":
                        return NewWindowsUpdateState(), nil
                    case "#microsoft.graph.workbook":
                        return NewWorkbook(), nil
                    case "#microsoft.graph.workbookApplication":
                        return NewWorkbookApplication(), nil
                    case "#microsoft.graph.workbookChart":
                        return NewWorkbookChart(), nil
                    case "#microsoft.graph.workbookChartAreaFormat":
                        return NewWorkbookChartAreaFormat(), nil
                    case "#microsoft.graph.workbookChartAxes":
                        return NewWorkbookChartAxes(), nil
                    case "#microsoft.graph.workbookChartAxis":
                        return NewWorkbookChartAxis(), nil
                    case "#microsoft.graph.workbookChartAxisFormat":
                        return NewWorkbookChartAxisFormat(), nil
                    case "#microsoft.graph.workbookChartAxisTitle":
                        return NewWorkbookChartAxisTitle(), nil
                    case "#microsoft.graph.workbookChartAxisTitleFormat":
                        return NewWorkbookChartAxisTitleFormat(), nil
                    case "#microsoft.graph.workbookChartDataLabelFormat":
                        return NewWorkbookChartDataLabelFormat(), nil
                    case "#microsoft.graph.workbookChartDataLabels":
                        return NewWorkbookChartDataLabels(), nil
                    case "#microsoft.graph.workbookChartFill":
                        return NewWorkbookChartFill(), nil
                    case "#microsoft.graph.workbookChartFont":
                        return NewWorkbookChartFont(), nil
                    case "#microsoft.graph.workbookChartGridlines":
                        return NewWorkbookChartGridlines(), nil
                    case "#microsoft.graph.workbookChartGridlinesFormat":
                        return NewWorkbookChartGridlinesFormat(), nil
                    case "#microsoft.graph.workbookChartLegend":
                        return NewWorkbookChartLegend(), nil
                    case "#microsoft.graph.workbookChartLegendFormat":
                        return NewWorkbookChartLegendFormat(), nil
                    case "#microsoft.graph.workbookChartLineFormat":
                        return NewWorkbookChartLineFormat(), nil
                    case "#microsoft.graph.workbookChartPoint":
                        return NewWorkbookChartPoint(), nil
                    case "#microsoft.graph.workbookChartPointFormat":
                        return NewWorkbookChartPointFormat(), nil
                    case "#microsoft.graph.workbookChartSeries":
                        return NewWorkbookChartSeries(), nil
                    case "#microsoft.graph.workbookChartSeriesFormat":
                        return NewWorkbookChartSeriesFormat(), nil
                    case "#microsoft.graph.workbookChartTitle":
                        return NewWorkbookChartTitle(), nil
                    case "#microsoft.graph.workbookChartTitleFormat":
                        return NewWorkbookChartTitleFormat(), nil
                    case "#microsoft.graph.workbookComment":
                        return NewWorkbookComment(), nil
                    case "#microsoft.graph.workbookCommentReply":
                        return NewWorkbookCommentReply(), nil
                    case "#microsoft.graph.workbookFilter":
                        return NewWorkbookFilter(), nil
                    case "#microsoft.graph.workbookFormatProtection":
                        return NewWorkbookFormatProtection(), nil
                    case "#microsoft.graph.workbookFunctionResult":
                        return NewWorkbookFunctionResult(), nil
                    case "#microsoft.graph.workbookFunctions":
                        return NewWorkbookFunctions(), nil
                    case "#microsoft.graph.workbookNamedItem":
                        return NewWorkbookNamedItem(), nil
                    case "#microsoft.graph.workbookOperation":
                        return NewWorkbookOperation(), nil
                    case "#microsoft.graph.workbookPivotTable":
                        return NewWorkbookPivotTable(), nil
                    case "#microsoft.graph.workbookRange":
                        return NewWorkbookRange(), nil
                    case "#microsoft.graph.workbookRangeBorder":
                        return NewWorkbookRangeBorder(), nil
                    case "#microsoft.graph.workbookRangeFill":
                        return NewWorkbookRangeFill(), nil
                    case "#microsoft.graph.workbookRangeFont":
                        return NewWorkbookRangeFont(), nil
                    case "#microsoft.graph.workbookRangeFormat":
                        return NewWorkbookRangeFormat(), nil
                    case "#microsoft.graph.workbookRangeSort":
                        return NewWorkbookRangeSort(), nil
                    case "#microsoft.graph.workbookRangeView":
                        return NewWorkbookRangeView(), nil
                    case "#microsoft.graph.workbookTable":
                        return NewWorkbookTable(), nil
                    case "#microsoft.graph.workbookTableColumn":
                        return NewWorkbookTableColumn(), nil
                    case "#microsoft.graph.workbookTableRow":
                        return NewWorkbookTableRow(), nil
                    case "#microsoft.graph.workbookTableSort":
                        return NewWorkbookTableSort(), nil
                    case "#microsoft.graph.workbookWorksheet":
                        return NewWorkbookWorksheet(), nil
                    case "#microsoft.graph.workbookWorksheetProtection":
                        return NewWorkbookWorksheetProtection(), nil
                    case "#microsoft.graph.zebraFotaArtifact":
                        return NewZebraFotaArtifact(), nil
                    case "#microsoft.graph.zebraFotaConnector":
                        return NewZebraFotaConnector(), nil
                    case "#microsoft.graph.zebraFotaDeployment":
                        return NewZebraFotaDeployment(), nil
                }
            }
        }
    }
    return NewEntity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Entity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Entity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *Entity) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Serialize serializes information the current object
func (m *Entity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *Entity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. The id property
func (m *Entity) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
