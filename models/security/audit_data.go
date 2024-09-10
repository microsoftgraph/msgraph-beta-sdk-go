package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AuditData struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAuditData instantiates a new AuditData and sets the default values.
func NewAuditData()(*AuditData) {
    m := &AuditData{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuditDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                switch *mappingValue {
                    case "#microsoft.graph.security.aadRiskDetectionAuditRecord":
                        return NewAadRiskDetectionAuditRecord(), nil
                    case "#microsoft.graph.security.aedAuditRecord":
                        return NewAedAuditRecord(), nil
                    case "#microsoft.graph.security.aiAppInteractionAuditRecord":
                        return NewAiAppInteractionAuditRecord(), nil
                    case "#microsoft.graph.security.aipFileDeleted":
                        return NewAipFileDeleted(), nil
                    case "#microsoft.graph.security.aipHeartBeat":
                        return NewAipHeartBeat(), nil
                    case "#microsoft.graph.security.aipProtectionActionLogRequest":
                        return NewAipProtectionActionLogRequest(), nil
                    case "#microsoft.graph.security.aipScannerDiscoverEvent":
                        return NewAipScannerDiscoverEvent(), nil
                    case "#microsoft.graph.security.aipSensitivityLabelActionLogRequest":
                        return NewAipSensitivityLabelActionLogRequest(), nil
                    case "#microsoft.graph.security.airAdminActionInvestigationData":
                        return NewAirAdminActionInvestigationData(), nil
                    case "#microsoft.graph.security.airInvestigationData":
                        return NewAirInvestigationData(), nil
                    case "#microsoft.graph.security.airManualInvestigationData":
                        return NewAirManualInvestigationData(), nil
                    case "#microsoft.graph.security.attackSimAdminAuditRecord":
                        return NewAttackSimAdminAuditRecord(), nil
                    case "#microsoft.graph.security.auditSearchAuditRecord":
                        return NewAuditSearchAuditRecord(), nil
                    case "#microsoft.graph.security.azureActiveDirectoryAccountLogonAuditRecord":
                        return NewAzureActiveDirectoryAccountLogonAuditRecord(), nil
                    case "#microsoft.graph.security.azureActiveDirectoryAuditRecord":
                        return NewAzureActiveDirectoryAuditRecord(), nil
                    case "#microsoft.graph.security.azureActiveDirectoryBaseAuditRecord":
                        return NewAzureActiveDirectoryBaseAuditRecord(), nil
                    case "#microsoft.graph.security.azureActiveDirectoryStsLogonAuditRecord":
                        return NewAzureActiveDirectoryStsLogonAuditRecord(), nil
                    case "#microsoft.graph.security.campaignAuditRecord":
                        return NewCampaignAuditRecord(), nil
                    case "#microsoft.graph.security.caseAuditRecord":
                        return NewCaseAuditRecord(), nil
                    case "#microsoft.graph.security.caseInvestigation":
                        return NewCaseInvestigation(), nil
                    case "#microsoft.graph.security.cdpColdCrawlStatusRecord":
                        return NewCdpColdCrawlStatusRecord(), nil
                    case "#microsoft.graph.security.cdpContentExplorerAggregateRecord":
                        return NewCdpContentExplorerAggregateRecord(), nil
                    case "#microsoft.graph.security.cdpDlpSensitiveAuditRecord":
                        return NewCdpDlpSensitiveAuditRecord(), nil
                    case "#microsoft.graph.security.cdpDlpSensitiveEndpointAuditRecord":
                        return NewCdpDlpSensitiveEndpointAuditRecord(), nil
                    case "#microsoft.graph.security.cdpLogRecord":
                        return NewCdpLogRecord(), nil
                    case "#microsoft.graph.security.cdpOcrBillingRecord":
                        return NewCdpOcrBillingRecord(), nil
                    case "#microsoft.graph.security.cdpResourceScopeChangeEventRecord":
                        return NewCdpResourceScopeChangeEventRecord(), nil
                    case "#microsoft.graph.security.cernerSMSLinkRecord":
                        return NewCernerSMSLinkRecord(), nil
                    case "#microsoft.graph.security.cernerSMSSettingsUpdateRecord":
                        return NewCernerSMSSettingsUpdateRecord(), nil
                    case "#microsoft.graph.security.cernerSMSUnlinkRecord":
                        return NewCernerSMSUnlinkRecord(), nil
                    case "#microsoft.graph.security.complianceConnectorAuditRecord":
                        return NewComplianceConnectorAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDLMExchangeAuditRecord":
                        return NewComplianceDLMExchangeAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDLMSharePointAuditRecord":
                        return NewComplianceDLMSharePointAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpApplicationsAuditRecord":
                        return NewComplianceDlpApplicationsAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpApplicationsClassificationAuditRecord":
                        return NewComplianceDlpApplicationsClassificationAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpBaseAuditRecord":
                        return NewComplianceDlpBaseAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpClassificationBaseAuditRecord":
                        return NewComplianceDlpClassificationBaseAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpClassificationBaseCdpRecord":
                        return NewComplianceDlpClassificationBaseCdpRecord(), nil
                    case "#microsoft.graph.security.complianceDlpEndpointAuditRecord":
                        return NewComplianceDlpEndpointAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpEndpointDiscoveryAuditRecord":
                        return NewComplianceDlpEndpointDiscoveryAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpExchangeAuditRecord":
                        return NewComplianceDlpExchangeAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpExchangeClassificationAuditRecord":
                        return NewComplianceDlpExchangeClassificationAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpExchangeClassificationCdpRecord":
                        return NewComplianceDlpExchangeClassificationCdpRecord(), nil
                    case "#microsoft.graph.security.complianceDlpExchangeDiscoveryAuditRecord":
                        return NewComplianceDlpExchangeDiscoveryAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpSharePointAuditRecord":
                        return NewComplianceDlpSharePointAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpSharePointClassificationAuditRecord":
                        return NewComplianceDlpSharePointClassificationAuditRecord(), nil
                    case "#microsoft.graph.security.complianceDlpSharePointClassificationExtendedAuditRecord":
                        return NewComplianceDlpSharePointClassificationExtendedAuditRecord(), nil
                    case "#microsoft.graph.security.complianceManagerActionRecord":
                        return NewComplianceManagerActionRecord(), nil
                    case "#microsoft.graph.security.complianceSupervisionBaseAuditRecord":
                        return NewComplianceSupervisionBaseAuditRecord(), nil
                    case "#microsoft.graph.security.complianceSupervisionExchangeAuditRecord":
                        return NewComplianceSupervisionExchangeAuditRecord(), nil
                    case "#microsoft.graph.security.consumptionResourceAuditRecord":
                        return NewConsumptionResourceAuditRecord(), nil
                    case "#microsoft.graph.security.copilotInteractionAuditRecord":
                        return NewCopilotInteractionAuditRecord(), nil
                    case "#microsoft.graph.security.coreReportingSettingsAuditRecord":
                        return NewCoreReportingSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.cortanaBriefingAuditRecord":
                        return NewCortanaBriefingAuditRecord(), nil
                    case "#microsoft.graph.security.cpsCommonPolicyAuditRecord":
                        return NewCpsCommonPolicyAuditRecord(), nil
                    case "#microsoft.graph.security.cpsPolicyConfigAuditRecord":
                        return NewCpsPolicyConfigAuditRecord(), nil
                    case "#microsoft.graph.security.crmBaseAuditRecord":
                        return NewCrmBaseAuditRecord(), nil
                    case "#microsoft.graph.security.crmEntityOperationAuditRecord":
                        return NewCrmEntityOperationAuditRecord(), nil
                    case "#microsoft.graph.security.customerKeyServiceEncryptionAuditRecord":
                        return NewCustomerKeyServiceEncryptionAuditRecord(), nil
                    case "#microsoft.graph.security.dataCenterSecurityBaseAuditRecord":
                        return NewDataCenterSecurityBaseAuditRecord(), nil
                    case "#microsoft.graph.security.dataCenterSecurityCmdletAuditRecord":
                        return NewDataCenterSecurityCmdletAuditRecord(), nil
                    case "#microsoft.graph.security.dataGovernanceAuditRecord":
                        return NewDataGovernanceAuditRecord(), nil
                    case "#microsoft.graph.security.dataInsightsRestApiAuditRecord":
                        return NewDataInsightsRestApiAuditRecord(), nil
                    case "#microsoft.graph.security.dataLakeExportOperationAuditRecord":
                        return NewDataLakeExportOperationAuditRecord(), nil
                    case "#microsoft.graph.security.dataShareOperationAuditRecord":
                        return NewDataShareOperationAuditRecord(), nil
                    case "#microsoft.graph.security.defaultAuditData":
                        return NewDefaultAuditData(), nil
                    case "#microsoft.graph.security.defenderSecurityAlertBaseRecord":
                        return NewDefenderSecurityAlertBaseRecord(), nil
                    case "#microsoft.graph.security.deleteCertificateRecord":
                        return NewDeleteCertificateRecord(), nil
                    case "#microsoft.graph.security.disableConsentRecord":
                        return NewDisableConsentRecord(), nil
                    case "#microsoft.graph.security.discoveryAuditRecord":
                        return NewDiscoveryAuditRecord(), nil
                    case "#microsoft.graph.security.dlpEndpointAuditRecord":
                        return NewDlpEndpointAuditRecord(), nil
                    case "#microsoft.graph.security.dlpSensitiveInformationTypeCmdletRecord":
                        return NewDlpSensitiveInformationTypeCmdletRecord(), nil
                    case "#microsoft.graph.security.dlpSensitiveInformationTypeRulePackageCmdletRecord":
                        return NewDlpSensitiveInformationTypeRulePackageCmdletRecord(), nil
                    case "#microsoft.graph.security.downloadCertificateRecord":
                        return NewDownloadCertificateRecord(), nil
                    case "#microsoft.graph.security.dynamics365BusinessCentralAuditRecord":
                        return NewDynamics365BusinessCentralAuditRecord(), nil
                    case "#microsoft.graph.security.enableConsentRecord":
                        return NewEnableConsentRecord(), nil
                    case "#microsoft.graph.security.epicSMSLinkRecord":
                        return NewEpicSMSLinkRecord(), nil
                    case "#microsoft.graph.security.epicSMSSettingsUpdateRecord":
                        return NewEpicSMSSettingsUpdateRecord(), nil
                    case "#microsoft.graph.security.epicSMSUnlinkRecord":
                        return NewEpicSMSUnlinkRecord(), nil
                    case "#microsoft.graph.security.exchangeAdminAuditRecord":
                        return NewExchangeAdminAuditRecord(), nil
                    case "#microsoft.graph.security.exchangeAggregatedMailboxAuditRecord":
                        return NewExchangeAggregatedMailboxAuditRecord(), nil
                    case "#microsoft.graph.security.exchangeAggregatedOperationRecord":
                        return NewExchangeAggregatedOperationRecord(), nil
                    case "#microsoft.graph.security.exchangeMailboxAuditBaseRecord":
                        return NewExchangeMailboxAuditBaseRecord(), nil
                    case "#microsoft.graph.security.exchangeMailboxAuditGroupRecord":
                        return NewExchangeMailboxAuditGroupRecord(), nil
                    case "#microsoft.graph.security.exchangeMailboxAuditRecord":
                        return NewExchangeMailboxAuditRecord(), nil
                    case "#microsoft.graph.security.fhirBaseUrlAddRecord":
                        return NewFhirBaseUrlAddRecord(), nil
                    case "#microsoft.graph.security.fhirBaseUrlApproveRecord":
                        return NewFhirBaseUrlApproveRecord(), nil
                    case "#microsoft.graph.security.fhirBaseUrlDeleteRecord":
                        return NewFhirBaseUrlDeleteRecord(), nil
                    case "#microsoft.graph.security.fhirBaseUrlUpdateRecord":
                        return NewFhirBaseUrlUpdateRecord(), nil
                    case "#microsoft.graph.security.healthcareSignalRecord":
                        return NewHealthcareSignalRecord(), nil
                    case "#microsoft.graph.security.hostedRpaAuditRecord":
                        return NewHostedRpaAuditRecord(), nil
                    case "#microsoft.graph.security.hrSignalAuditRecord":
                        return NewHrSignalAuditRecord(), nil
                    case "#microsoft.graph.security.hygieneEventRecord":
                        return NewHygieneEventRecord(), nil
                    case "#microsoft.graph.security.informationBarrierPolicyApplicationAuditRecord":
                        return NewInformationBarrierPolicyApplicationAuditRecord(), nil
                    case "#microsoft.graph.security.informationWorkerProtectionAuditRecord":
                        return NewInformationWorkerProtectionAuditRecord(), nil
                    case "#microsoft.graph.security.insiderRiskScopedUserInsightsRecord":
                        return NewInsiderRiskScopedUserInsightsRecord(), nil
                    case "#microsoft.graph.security.insiderRiskScopedUsersRecord":
                        return NewInsiderRiskScopedUsersRecord(), nil
                    case "#microsoft.graph.security.irmSecurityAlertRecord":
                        return NewIrmSecurityAlertRecord(), nil
                    case "#microsoft.graph.security.irmUserDefinedDetectionRecord":
                        return NewIrmUserDefinedDetectionRecord(), nil
                    case "#microsoft.graph.security.kaizalaAuditRecord":
                        return NewKaizalaAuditRecord(), nil
                    case "#microsoft.graph.security.labelAnalyticsAggregateAuditRecord":
                        return NewLabelAnalyticsAggregateAuditRecord(), nil
                    case "#microsoft.graph.security.labelContentExplorerAuditRecord":
                        return NewLabelContentExplorerAuditRecord(), nil
                    case "#microsoft.graph.security.largeContentMetadataAuditRecord":
                        return NewLargeContentMetadataAuditRecord(), nil
                    case "#microsoft.graph.security.m365ComplianceConnectorAuditRecord":
                        return NewM365ComplianceConnectorAuditRecord(), nil
                    case "#microsoft.graph.security.m365DAADAuditRecord":
                        return NewM365DAADAuditRecord(), nil
                    case "#microsoft.graph.security.mailSubmissionData":
                        return NewMailSubmissionData(), nil
                    case "#microsoft.graph.security.managedServicesAuditRecord":
                        return NewManagedServicesAuditRecord(), nil
                    case "#microsoft.graph.security.managedTenantsAuditRecord":
                        return NewManagedTenantsAuditRecord(), nil
                    case "#microsoft.graph.security.mapgAlertsAuditRecord":
                        return NewMapgAlertsAuditRecord(), nil
                    case "#microsoft.graph.security.mapgOnboardAuditRecord":
                        return NewMapgOnboardAuditRecord(), nil
                    case "#microsoft.graph.security.mapgPolicyAuditRecord":
                        return NewMapgPolicyAuditRecord(), nil
                    case "#microsoft.graph.security.mcasAlertsAuditRecord":
                        return NewMcasAlertsAuditRecord(), nil
                    case "#microsoft.graph.security.mdaDataSecuritySignalRecord":
                        return NewMdaDataSecuritySignalRecord(), nil
                    case "#microsoft.graph.security.mdatpAuditRecord":
                        return NewMdatpAuditRecord(), nil
                    case "#microsoft.graph.security.mdcEventsRecord":
                        return NewMdcEventsRecord(), nil
                    case "#microsoft.graph.security.mdiAuditRecord":
                        return NewMdiAuditRecord(), nil
                    case "#microsoft.graph.security.meshWorldsAuditRecord":
                        return NewMeshWorldsAuditRecord(), nil
                    case "#microsoft.graph.security.microsoft365BackupBackupItemAuditRecord":
                        return NewMicrosoft365BackupBackupItemAuditRecord(), nil
                    case "#microsoft.graph.security.microsoft365BackupBackupPolicyAuditRecord":
                        return NewMicrosoft365BackupBackupPolicyAuditRecord(), nil
                    case "#microsoft.graph.security.microsoft365BackupRestoreItemAuditRecord":
                        return NewMicrosoft365BackupRestoreItemAuditRecord(), nil
                    case "#microsoft.graph.security.microsoft365BackupRestoreTaskAuditRecord":
                        return NewMicrosoft365BackupRestoreTaskAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftDefenderExpertsBaseAuditRecord":
                        return NewMicrosoftDefenderExpertsBaseAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftDefenderExpertsXDRAuditRecord":
                        return NewMicrosoftDefenderExpertsXDRAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftFlowAuditRecord":
                        return NewMicrosoftFlowAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftFormsAuditRecord":
                        return NewMicrosoftFormsAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftGraphDataConnectConsent":
                        return NewMicrosoftGraphDataConnectConsent(), nil
                    case "#microsoft.graph.security.microsoftGraphDataConnectOperation":
                        return NewMicrosoftGraphDataConnectOperation(), nil
                    case "#microsoft.graph.security.microsoftPurviewDataMapOperationRecord":
                        return NewMicrosoftPurviewDataMapOperationRecord(), nil
                    case "#microsoft.graph.security.microsoftPurviewMetadataPolicyOperationRecord":
                        return NewMicrosoftPurviewMetadataPolicyOperationRecord(), nil
                    case "#microsoft.graph.security.microsoftPurviewPolicyOperationRecord":
                        return NewMicrosoftPurviewPolicyOperationRecord(), nil
                    case "#microsoft.graph.security.microsoftPurviewPrivacyAuditEvent":
                        return NewMicrosoftPurviewPrivacyAuditEvent(), nil
                    case "#microsoft.graph.security.microsoftStreamAuditRecord":
                        return NewMicrosoftStreamAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftTeamsAdminAuditRecord":
                        return NewMicrosoftTeamsAdminAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftTeamsAnalyticsAuditRecord":
                        return NewMicrosoftTeamsAnalyticsAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftTeamsAuditRecord":
                        return NewMicrosoftTeamsAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftTeamsDeviceAuditRecord":
                        return NewMicrosoftTeamsDeviceAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftTeamsRetentionLabelActionAuditRecord":
                        return NewMicrosoftTeamsRetentionLabelActionAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftTeamsSensitivityLabelActionAuditRecord":
                        return NewMicrosoftTeamsSensitivityLabelActionAuditRecord(), nil
                    case "#microsoft.graph.security.microsoftTeamsShiftsAuditRecord":
                        return NewMicrosoftTeamsShiftsAuditRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelExchangeItemAuditRecord":
                        return NewMipAutoLabelExchangeItemAuditRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelItemAuditRecord":
                        return NewMipAutoLabelItemAuditRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelPolicyAuditRecord":
                        return NewMipAutoLabelPolicyAuditRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelProgressFeedbackAuditRecord":
                        return NewMipAutoLabelProgressFeedbackAuditRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelSharePointItemAuditRecord":
                        return NewMipAutoLabelSharePointItemAuditRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelSharePointPolicyLocationAuditRecord":
                        return NewMipAutoLabelSharePointPolicyLocationAuditRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelSimulationSharePointCompletionRecord":
                        return NewMipAutoLabelSimulationSharePointCompletionRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelSimulationSharePointProgressRecord":
                        return NewMipAutoLabelSimulationSharePointProgressRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelSimulationStatisticsRecord":
                        return NewMipAutoLabelSimulationStatisticsRecord(), nil
                    case "#microsoft.graph.security.mipAutoLabelSimulationStatusRecord":
                        return NewMipAutoLabelSimulationStatusRecord(), nil
                    case "#microsoft.graph.security.mipExactDataMatchAuditRecord":
                        return NewMipExactDataMatchAuditRecord(), nil
                    case "#microsoft.graph.security.mipLabelAnalyticsAuditRecord":
                        return NewMipLabelAnalyticsAuditRecord(), nil
                    case "#microsoft.graph.security.mipLabelAuditRecord":
                        return NewMipLabelAuditRecord(), nil
                    case "#microsoft.graph.security.mS365DCustomDetectionAuditRecord":
                        return NewMS365DCustomDetectionAuditRecord(), nil
                    case "#microsoft.graph.security.mS365DIncidentAuditRecord":
                        return NewMS365DIncidentAuditRecord(), nil
                    case "#microsoft.graph.security.mS365DSuppressionRuleAuditRecord":
                        return NewMS365DSuppressionRuleAuditRecord(), nil
                    case "#microsoft.graph.security.msdeGeneralSettingsAuditRecord":
                        return NewMsdeGeneralSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.msdeIndicatorsSettingsAuditRecord":
                        return NewMsdeIndicatorsSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.msdeResponseActionsAuditRecord":
                        return NewMsdeResponseActionsAuditRecord(), nil
                    case "#microsoft.graph.security.msdeRolesSettingsAuditRecord":
                        return NewMsdeRolesSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.msticNationStateNotificationRecord":
                        return NewMsticNationStateNotificationRecord(), nil
                    case "#microsoft.graph.security.multiStageDispositionAuditRecord":
                        return NewMultiStageDispositionAuditRecord(), nil
                    case "#microsoft.graph.security.myAnalyticsSettingsAuditRecord":
                        return NewMyAnalyticsSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.officeNativeAuditRecord":
                        return NewOfficeNativeAuditRecord(), nil
                    case "#microsoft.graph.security.omePortalAuditRecord":
                        return NewOmePortalAuditRecord(), nil
                    case "#microsoft.graph.security.oneDriveAuditRecord":
                        return NewOneDriveAuditRecord(), nil
                    case "#microsoft.graph.security.onPremisesFileShareScannerDlpAuditRecord":
                        return NewOnPremisesFileShareScannerDlpAuditRecord(), nil
                    case "#microsoft.graph.security.onPremisesScannerDlpAuditRecord":
                        return NewOnPremisesScannerDlpAuditRecord(), nil
                    case "#microsoft.graph.security.onPremisesSharePointScannerDlpAuditRecord":
                        return NewOnPremisesSharePointScannerDlpAuditRecord(), nil
                    case "#microsoft.graph.security.owaGetAccessTokenForResourceAuditRecord":
                        return NewOwaGetAccessTokenForResourceAuditRecord(), nil
                    case "#microsoft.graph.security.peopleAdminSettingsAuditRecord":
                        return NewPeopleAdminSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.physicalBadgingSignalAuditRecord":
                        return NewPhysicalBadgingSignalAuditRecord(), nil
                    case "#microsoft.graph.security.plannerCopyPlanAuditRecord":
                        return NewPlannerCopyPlanAuditRecord(), nil
                    case "#microsoft.graph.security.plannerPlanAuditRecord":
                        return NewPlannerPlanAuditRecord(), nil
                    case "#microsoft.graph.security.plannerPlanListAuditRecord":
                        return NewPlannerPlanListAuditRecord(), nil
                    case "#microsoft.graph.security.plannerRosterAuditRecord":
                        return NewPlannerRosterAuditRecord(), nil
                    case "#microsoft.graph.security.plannerRosterSensitivityLabelAuditRecord":
                        return NewPlannerRosterSensitivityLabelAuditRecord(), nil
                    case "#microsoft.graph.security.plannerTaskAuditRecord":
                        return NewPlannerTaskAuditRecord(), nil
                    case "#microsoft.graph.security.plannerTaskListAuditRecord":
                        return NewPlannerTaskListAuditRecord(), nil
                    case "#microsoft.graph.security.plannerTenantSettingsAuditRecord":
                        return NewPlannerTenantSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.powerAppsAuditAppRecord":
                        return NewPowerAppsAuditAppRecord(), nil
                    case "#microsoft.graph.security.powerAppsAuditPlanRecord":
                        return NewPowerAppsAuditPlanRecord(), nil
                    case "#microsoft.graph.security.powerAppsAuditResourceRecord":
                        return NewPowerAppsAuditResourceRecord(), nil
                    case "#microsoft.graph.security.powerBiAuditRecord":
                        return NewPowerBiAuditRecord(), nil
                    case "#microsoft.graph.security.powerBiDlpAuditRecord":
                        return NewPowerBiDlpAuditRecord(), nil
                    case "#microsoft.graph.security.powerPagesSiteAuditRecord":
                        return NewPowerPagesSiteAuditRecord(), nil
                    case "#microsoft.graph.security.powerPlatformAdminDlpAuditRecord":
                        return NewPowerPlatformAdminDlpAuditRecord(), nil
                    case "#microsoft.graph.security.powerPlatformAdminEnvironmentAuditRecord":
                        return NewPowerPlatformAdminEnvironmentAuditRecord(), nil
                    case "#microsoft.graph.security.powerPlatformAdministratorActivityRecord":
                        return NewPowerPlatformAdministratorActivityRecord(), nil
                    case "#microsoft.graph.security.powerPlatformLockboxResourceAccessRequestAuditRecord":
                        return NewPowerPlatformLockboxResourceAccessRequestAuditRecord(), nil
                    case "#microsoft.graph.security.powerPlatformLockboxResourceCommandAuditRecord":
                        return NewPowerPlatformLockboxResourceCommandAuditRecord(), nil
                    case "#microsoft.graph.security.powerPlatformServiceActivityAuditRecord":
                        return NewPowerPlatformServiceActivityAuditRecord(), nil
                    case "#microsoft.graph.security.privacyDataMatchAuditRecord":
                        return NewPrivacyDataMatchAuditRecord(), nil
                    case "#microsoft.graph.security.privacyDataMinimizationRecord":
                        return NewPrivacyDataMinimizationRecord(), nil
                    case "#microsoft.graph.security.privacyDigestEmailRecord":
                        return NewPrivacyDigestEmailRecord(), nil
                    case "#microsoft.graph.security.privacyOpenAccessAuditRecord":
                        return NewPrivacyOpenAccessAuditRecord(), nil
                    case "#microsoft.graph.security.privacyPortalAuditRecord":
                        return NewPrivacyPortalAuditRecord(), nil
                    case "#microsoft.graph.security.privacyRemediationActionRecord":
                        return NewPrivacyRemediationActionRecord(), nil
                    case "#microsoft.graph.security.privacyRemediationRecord":
                        return NewPrivacyRemediationRecord(), nil
                    case "#microsoft.graph.security.privacyTenantAuditHistoryRecord":
                        return NewPrivacyTenantAuditHistoryRecord(), nil
                    case "#microsoft.graph.security.projectAuditRecord":
                        return NewProjectAuditRecord(), nil
                    case "#microsoft.graph.security.projectForTheWebAssignedToMeSettingsAuditRecord":
                        return NewProjectForTheWebAssignedToMeSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.projectForTheWebProjectAuditRecord":
                        return NewProjectForTheWebProjectAuditRecord(), nil
                    case "#microsoft.graph.security.projectForTheWebProjectSettingsAuditRecord":
                        return NewProjectForTheWebProjectSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.projectForTheWebRoadmapAuditRecord":
                        return NewProjectForTheWebRoadmapAuditRecord(), nil
                    case "#microsoft.graph.security.projectForTheWebRoadmapItemAuditRecord":
                        return NewProjectForTheWebRoadmapItemAuditRecord(), nil
                    case "#microsoft.graph.security.projectForTheWebRoadmapSettingsAuditRecord":
                        return NewProjectForTheWebRoadmapSettingsAuditRecord(), nil
                    case "#microsoft.graph.security.projectForTheWebTaskAuditRecord":
                        return NewProjectForTheWebTaskAuditRecord(), nil
                    case "#microsoft.graph.security.publicFolderAuditRecord":
                        return NewPublicFolderAuditRecord(), nil
                    case "#microsoft.graph.security.purviewInsiderRiskAlertsRecord":
                        return NewPurviewInsiderRiskAlertsRecord(), nil
                    case "#microsoft.graph.security.purviewInsiderRiskCasesRecord":
                        return NewPurviewInsiderRiskCasesRecord(), nil
                    case "#microsoft.graph.security.quarantineAuditRecord":
                        return NewQuarantineAuditRecord(), nil
                    case "#microsoft.graph.security.recordsManagementAuditRecord":
                        return NewRecordsManagementAuditRecord(), nil
                    case "#microsoft.graph.security.retentionPolicyAuditRecord":
                        return NewRetentionPolicyAuditRecord(), nil
                    case "#microsoft.graph.security.scoreEvidence":
                        return NewScoreEvidence(), nil
                    case "#microsoft.graph.security.scorePlatformGenericAuditRecord":
                        return NewScorePlatformGenericAuditRecord(), nil
                    case "#microsoft.graph.security.scriptRunAuditRecord":
                        return NewScriptRunAuditRecord(), nil
                    case "#microsoft.graph.security.searchAuditRecord":
                        return NewSearchAuditRecord(), nil
                    case "#microsoft.graph.security.securityComplianceAlertRecord":
                        return NewSecurityComplianceAlertRecord(), nil
                    case "#microsoft.graph.security.securityComplianceCenterEOPCmdletAuditRecord":
                        return NewSecurityComplianceCenterEOPCmdletAuditRecord(), nil
                    case "#microsoft.graph.security.securityComplianceInsightsAuditRecord":
                        return NewSecurityComplianceInsightsAuditRecord(), nil
                    case "#microsoft.graph.security.securityComplianceRBACAuditRecord":
                        return NewSecurityComplianceRBACAuditRecord(), nil
                    case "#microsoft.graph.security.securityComplianceUserChangeAuditRecord":
                        return NewSecurityComplianceUserChangeAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointAppPermissionOperationAuditRecord":
                        return NewSharePointAppPermissionOperationAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointAuditRecord":
                        return NewSharePointAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointCommentOperationAuditRecord":
                        return NewSharePointCommentOperationAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointContentTypeOperationAuditRecord":
                        return NewSharePointContentTypeOperationAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointESignatureAuditRecord":
                        return NewSharePointESignatureAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointFieldOperationAuditRecord":
                        return NewSharePointFieldOperationAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointFileOperationAuditRecord":
                        return NewSharePointFileOperationAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointListOperationAuditRecord":
                        return NewSharePointListOperationAuditRecord(), nil
                    case "#microsoft.graph.security.sharePointSharingOperationAuditRecord":
                        return NewSharePointSharingOperationAuditRecord(), nil
                    case "#microsoft.graph.security.skypeForBusinessBaseAuditRecord":
                        return NewSkypeForBusinessBaseAuditRecord(), nil
                    case "#microsoft.graph.security.skypeForBusinessCmdletsAuditRecord":
                        return NewSkypeForBusinessCmdletsAuditRecord(), nil
                    case "#microsoft.graph.security.skypeForBusinessPSTNUsageAuditRecord":
                        return NewSkypeForBusinessPSTNUsageAuditRecord(), nil
                    case "#microsoft.graph.security.skypeForBusinessUsersBlockedAuditRecord":
                        return NewSkypeForBusinessUsersBlockedAuditRecord(), nil
                    case "#microsoft.graph.security.smsCreatePhoneNumberRecord":
                        return NewSmsCreatePhoneNumberRecord(), nil
                    case "#microsoft.graph.security.smsDeletePhoneNumberRecord":
                        return NewSmsDeletePhoneNumberRecord(), nil
                    case "#microsoft.graph.security.supervisoryReviewDayXInsightsAuditRecord":
                        return NewSupervisoryReviewDayXInsightsAuditRecord(), nil
                    case "#microsoft.graph.security.syntheticProbeAuditRecord":
                        return NewSyntheticProbeAuditRecord(), nil
                    case "#microsoft.graph.security.teamsEasyApprovalsAuditRecord":
                        return NewTeamsEasyApprovalsAuditRecord(), nil
                    case "#microsoft.graph.security.teamsHealthcareAuditRecord":
                        return NewTeamsHealthcareAuditRecord(), nil
                    case "#microsoft.graph.security.teamsUpdatesAuditRecord":
                        return NewTeamsUpdatesAuditRecord(), nil
                    case "#microsoft.graph.security.tenantAllowBlockListAuditRecord":
                        return NewTenantAllowBlockListAuditRecord(), nil
                    case "#microsoft.graph.security.threatFinderAuditRecord":
                        return NewThreatFinderAuditRecord(), nil
                    case "#microsoft.graph.security.threatIntelligenceAtpContentData":
                        return NewThreatIntelligenceAtpContentData(), nil
                    case "#microsoft.graph.security.threatIntelligenceMailData":
                        return NewThreatIntelligenceMailData(), nil
                    case "#microsoft.graph.security.threatIntelligenceUrlClickData":
                        return NewThreatIntelligenceUrlClickData(), nil
                    case "#microsoft.graph.security.todoAuditRecord":
                        return NewTodoAuditRecord(), nil
                    case "#microsoft.graph.security.uamOperationAuditRecord":
                        return NewUamOperationAuditRecord(), nil
                    case "#microsoft.graph.security.unifiedGroupAuditRecord":
                        return NewUnifiedGroupAuditRecord(), nil
                    case "#microsoft.graph.security.unifiedSimulationMatchedItemAuditRecord":
                        return NewUnifiedSimulationMatchedItemAuditRecord(), nil
                    case "#microsoft.graph.security.unifiedSimulationSummaryAuditRecord":
                        return NewUnifiedSimulationSummaryAuditRecord(), nil
                    case "#microsoft.graph.security.uploadCertificateRecord":
                        return NewUploadCertificateRecord(), nil
                    case "#microsoft.graph.security.urbacAssignmentAuditRecord":
                        return NewUrbacAssignmentAuditRecord(), nil
                    case "#microsoft.graph.security.urbacEnableStateAuditRecord":
                        return NewUrbacEnableStateAuditRecord(), nil
                    case "#microsoft.graph.security.urbacRoleAuditRecord":
                        return NewUrbacRoleAuditRecord(), nil
                    case "#microsoft.graph.security.userTrainingAuditRecord":
                        return NewUserTrainingAuditRecord(), nil
                    case "#microsoft.graph.security.vfamBasePolicyAuditRecord":
                        return NewVfamBasePolicyAuditRecord(), nil
                    case "#microsoft.graph.security.vfamCreatePolicyAuditRecord":
                        return NewVfamCreatePolicyAuditRecord(), nil
                    case "#microsoft.graph.security.vfamDeletePolicyAuditRecord":
                        return NewVfamDeletePolicyAuditRecord(), nil
                    case "#microsoft.graph.security.vfamUpdatePolicyAuditRecord":
                        return NewVfamUpdatePolicyAuditRecord(), nil
                    case "#microsoft.graph.security.vivaGoalsAuditRecord":
                        return NewVivaGoalsAuditRecord(), nil
                    case "#microsoft.graph.security.vivaLearningAdminAuditRecord":
                        return NewVivaLearningAdminAuditRecord(), nil
                    case "#microsoft.graph.security.vivaLearningAuditRecord":
                        return NewVivaLearningAuditRecord(), nil
                    case "#microsoft.graph.security.vivaPulseAdminAuditRecord":
                        return NewVivaPulseAdminAuditRecord(), nil
                    case "#microsoft.graph.security.vivaPulseOrganizerAuditRecord":
                        return NewVivaPulseOrganizerAuditRecord(), nil
                    case "#microsoft.graph.security.vivaPulseReportAuditRecord":
                        return NewVivaPulseReportAuditRecord(), nil
                    case "#microsoft.graph.security.vivaPulseResponseAuditRecord":
                        return NewVivaPulseResponseAuditRecord(), nil
                    case "#microsoft.graph.security.wdatpAlertsAuditRecord":
                        return NewWdatpAlertsAuditRecord(), nil
                    case "#microsoft.graph.security.windows365CustomerLockboxAuditRecord":
                        return NewWindows365CustomerLockboxAuditRecord(), nil
                    case "#microsoft.graph.security.workplaceAnalyticsAuditRecord":
                        return NewWorkplaceAnalyticsAuditRecord(), nil
                    case "#microsoft.graph.security.yammerAuditRecord":
                        return NewYammerAuditRecord(), nil
                }
            }
        }
    }
    return NewAuditData(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AuditData) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AuditData) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuditData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AuditData) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuditData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditData) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AuditData) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuditData) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type AuditDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
}
