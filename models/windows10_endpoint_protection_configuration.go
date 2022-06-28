package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EndpointProtectionConfiguration 
type Windows10EndpointProtectionConfiguration struct {
    DeviceConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and microphone.
    applicationGuardAllowCameraMicrophoneRedirection *bool
    // Allow users to download files from Edge in the application guard container and save them on the host file system
    applicationGuardAllowFileSaveOnHost *bool
    // Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
    applicationGuardAllowPersistence *bool
    // Allow printing to Local Printers from Container
    applicationGuardAllowPrintToLocalPrinters *bool
    // Allow printing to Network Printers from Container
    applicationGuardAllowPrintToNetworkPrinters *bool
    // Allow printing to PDF from Container
    applicationGuardAllowPrintToPDF *bool
    // Allow printing to XPS from Container
    applicationGuardAllowPrintToXPS *bool
    // Allow application guard to use virtual GPU
    applicationGuardAllowVirtualGPU *bool
    // Block clipboard to share data from Host to Container, or from Container to Host, or both ways, or neither ways. Possible values are: notConfigured, blockBoth, blockHostToContainer, blockContainerToHost, blockNone.
    applicationGuardBlockClipboardSharing *ApplicationGuardBlockClipboardSharingType
    // Block clipboard to transfer image file, text file or neither of them. Possible values are: notConfigured, blockImageAndTextFile, blockImageFile, blockNone, blockTextFile.
    applicationGuardBlockFileTransfer *ApplicationGuardBlockFileTransferType
    // Block enterprise sites to load non-enterprise content, such as third party plug-ins
    applicationGuardBlockNonEnterpriseContent *bool
    // Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
    applicationGuardCertificateThumbprints []string
    // Enable Windows Defender Application Guard
    applicationGuardEnabled *bool
    // Enable Windows Defender Application Guard for newer Windows builds. Possible values are: notConfigured, enabledForEdge, enabledForOffice, enabledForEdgeAndOffice.
    applicationGuardEnabledOptions *ApplicationGuardEnabledOptions
    // Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
    applicationGuardForceAuditing *bool
    // Enables the Admin to choose what types of app to allow on devices. Possible values are: notConfigured, enforceComponentsAndStoreApps, auditComponentsAndStoreApps, enforceComponentsStoreAppsAndSmartlocker, auditComponentsStoreAppsAndSmartlocker.
    appLockerApplicationControl *AppLockerApplicationControlType
    // Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
    bitLockerAllowStandardUserEncryption *bool
    // Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
    bitLockerDisableWarningForOtherDiskEncryption *bool
    // Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
    bitLockerEnableStorageCardEncryptionOnMobile *bool
    // Allows the admin to require encryption to be turned on using BitLocker.
    bitLockerEncryptDevice *bool
    // BitLocker Fixed Drive Policy.
    bitLockerFixedDrivePolicy BitLockerFixedDrivePolicyable
    // This setting initiates a client-driven recovery password rotation after an OS drive recovery (either by using bootmgr or WinRE). Possible values are: notConfigured, disabled, enabledForAzureAd, enabledForAzureAdAndHybrid.
    bitLockerRecoveryPasswordRotation *BitLockerRecoveryPasswordRotationType
    // BitLocker Removable Drive Policy.
    bitLockerRemovableDrivePolicy BitLockerRemovableDrivePolicyable
    // BitLocker System Drive Policy.
    bitLockerSystemDrivePolicy BitLockerSystemDrivePolicyable
    // List of folder paths to be added to the list of protected folders
    defenderAdditionalGuardedFolders []string
    // Value indicating the behavior of Adobe Reader from creating child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderAdobeReaderLaunchChildProcess *DefenderProtectionType
    // Value indicating use of advanced protection against ransomeware. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderAdvancedRansomewareProtectionType *DefenderProtectionType
    // Allows or disallows Windows Defender Behavior Monitoring functionality.
    defenderAllowBehaviorMonitoring *bool
    // To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
    defenderAllowCloudProtection *bool
    // Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
    defenderAllowEndUserAccess *bool
    // Allows or disallows Windows Defender Intrusion Prevention functionality.
    defenderAllowIntrusionPreventionSystem *bool
    // Allows or disallows Windows Defender On Access Protection functionality.
    defenderAllowOnAccessProtection *bool
    // Allows or disallows Windows Defender Realtime Monitoring functionality.
    defenderAllowRealTimeMonitoring *bool
    // Allows or disallows scanning of archives.
    defenderAllowScanArchiveFiles *bool
    // Allows or disallows Windows Defender IOAVP Protection functionality.
    defenderAllowScanDownloads *bool
    // Allows or disallows a scanning of network files.
    defenderAllowScanNetworkFiles *bool
    // Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
    defenderAllowScanRemovableDrivesDuringFullScan *bool
    // Allows or disallows Windows Defender Script Scanning functionality.
    defenderAllowScanScriptsLoadedInInternetExplorer *bool
    // List of exe files and folders to be excluded from attack surface reduction rules
    defenderAttackSurfaceReductionExcludedPaths []string
    // Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
    defenderBlockEndUserAccess *bool
    // Value indicating the behavior of Block persistence through WMI event subscription. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderBlockPersistenceThroughWmiType *DefenderAttackSurfaceType
    // This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before running a scan.
    defenderCheckForSignaturesBeforeRunningScan *bool
    // Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS' setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
    defenderCloudBlockLevel *DefenderCloudBlockLevelType
    // Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
    defenderCloudExtendedTimeoutInSeconds *int32
    // Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
    defenderDaysBeforeDeletingQuarantinedMalware *int32
    // Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
    defenderDetectedMalwareActions DefenderDetectedMalwareActionsable
    // Allows or disallows Windows Defender Behavior Monitoring functionality.
    defenderDisableBehaviorMonitoring *bool
    // This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
    defenderDisableCatchupFullScan *bool
    // This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
    defenderDisableCatchupQuickScan *bool
    // To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
    defenderDisableCloudProtection *bool
    // Allows or disallows Windows Defender Intrusion Prevention functionality.
    defenderDisableIntrusionPreventionSystem *bool
    // Allows or disallows Windows Defender On Access Protection functionality.
    defenderDisableOnAccessProtection *bool
    // Allows or disallows Windows Defender Realtime Monitoring functionality.
    defenderDisableRealTimeMonitoring *bool
    // Allows or disallows scanning of archives.
    defenderDisableScanArchiveFiles *bool
    // Allows or disallows Windows Defender IOAVP Protection functionality.
    defenderDisableScanDownloads *bool
    // Allows or disallows a scanning of network files.
    defenderDisableScanNetworkFiles *bool
    // Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
    defenderDisableScanRemovableDrivesDuringFullScan *bool
    // Allows or disallows Windows Defender Script Scanning functionality.
    defenderDisableScanScriptsLoadedInInternetExplorer *bool
    // Value indicating if execution of executable content (exe, dll, ps, js, vbs, etc) should be dropped from email (webmail/mail-client). Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderEmailContentExecution *DefenderProtectionType
    // Value indicating if execution of executable content (exe, dll, ps, js, vbs, etc) should be dropped from email (webmail/mail-client). Possible values are: userDefined, block, auditMode, warn, disable.
    defenderEmailContentExecutionType *DefenderAttackSurfaceType
    // This policy setting allows you to enable or disable low CPU priority for scheduled scans.
    defenderEnableLowCpuPriority *bool
    // Allows or disallows scanning of email.
    defenderEnableScanIncomingMail *bool
    // Allows or disallows a full scan of mapped network drives.
    defenderEnableScanMappedNetworkDrivesDuringFullScan *bool
    // Xml content containing information regarding exploit protection details.
    defenderExploitProtectionXml []byte
    // Name of the file from which DefenderExploitProtectionXml was obtained.
    defenderExploitProtectionXmlFileName *string
    // File extensions to exclude from scans and real time protection.
    defenderFileExtensionsToExclude []string
    // Files and folder to exclude from scans and real time protection.
    defenderFilesAndFoldersToExclude []string
    // List of paths to exe that are allowed to access protected folders
    defenderGuardedFoldersAllowedAppPaths []string
    // Value indicating the behavior of protected folders. Possible values are: userDefined, enable, auditMode, blockDiskModification, auditDiskModification.
    defenderGuardMyFoldersType *FolderProtectionType
    // Value indicating the behavior of NetworkProtection. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderNetworkProtectionType *DefenderProtectionType
    // Value indicating the behavior of Office applications/macros creating or launching executable content. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderOfficeAppsExecutableContentCreationOrLaunch *DefenderProtectionType
    // Value indicating the behavior of Office applications/macros creating or launching executable content. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderOfficeAppsExecutableContentCreationOrLaunchType *DefenderAttackSurfaceType
    // Value indicating the behavior of Office application launching child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderOfficeAppsLaunchChildProcess *DefenderProtectionType
    // Value indicating the behavior of Office application launching child processes. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderOfficeAppsLaunchChildProcessType *DefenderAttackSurfaceType
    // Value indicating the behavior of  Office applications injecting into other processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderOfficeAppsOtherProcessInjection *DefenderProtectionType
    // Value indicating the behavior of Office applications injecting into other processes. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderOfficeAppsOtherProcessInjectionType *DefenderAttackSurfaceType
    // Value indicating the behavior of Office communication applications, including Microsoft Outlook, from creating child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderOfficeCommunicationAppsLaunchChildProcess *DefenderProtectionType
    // Value indicating the behavior of Win32 imports from Macro code in Office. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderOfficeMacroCodeAllowWin32Imports *DefenderProtectionType
    // Value indicating the behavior of Win32 imports from Macro code in Office. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderOfficeMacroCodeAllowWin32ImportsType *DefenderAttackSurfaceType
    // Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs). Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderPotentiallyUnwantedAppAction *DefenderProtectionType
    // Value indicating if credential stealing from the Windows local security authority subsystem is permitted. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderPreventCredentialStealingType *DefenderProtectionType
    // Value indicating response to process creations originating from PSExec and WMI commands. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderProcessCreation *DefenderProtectionType
    // Value indicating response to process creations originating from PSExec and WMI commands. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderProcessCreationType *DefenderAttackSurfaceType
    // Processes to exclude from scans and real time protection.
    defenderProcessesToExclude []string
    // Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly, monitorOutgoingFilesOnly.
    defenderScanDirection *DefenderRealtimeScanDirection
    // Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid values 0 to 100
    defenderScanMaxCpuPercentage *int32
    // Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
    defenderScanType *DefenderScanType
    // Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
    defenderScheduledQuickScanTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
    defenderScheduledScanDay *WeeklySchedule
    // Selects the time of day that the Windows Defender scan should run.
    defenderScheduledScanTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Value indicating the behavior of js/vbs executing payload downloaded from Internet. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderScriptDownloadedPayloadExecution *DefenderProtectionType
    // Value indicating the behavior of js/vbs executing payload downloaded from Internet. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderScriptDownloadedPayloadExecutionType *DefenderAttackSurfaceType
    // Value indicating the behavior of obfuscated js/vbs/ps/macro code. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderScriptObfuscatedMacroCode *DefenderProtectionType
    // Value indicating the behavior of obfuscated js/vbs/ps/macro code. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderScriptObfuscatedMacroCodeType *DefenderAttackSurfaceType
    // Indicates whether or not to block user from overriding Exploit Protection settings.
    defenderSecurityCenterBlockExploitProtectionOverride *bool
    // Used to disable the display of the account protection area.
    defenderSecurityCenterDisableAccountUI *bool
    // Used to disable the display of the app and browser protection area.
    defenderSecurityCenterDisableAppBrowserUI *bool
    // Used to disable the display of the Clear TPM button.
    defenderSecurityCenterDisableClearTpmUI *bool
    // Used to disable the display of the family options area.
    defenderSecurityCenterDisableFamilyUI *bool
    // Used to disable the display of the hardware protection area.
    defenderSecurityCenterDisableHardwareUI *bool
    // Used to disable the display of the device performance and health area.
    defenderSecurityCenterDisableHealthUI *bool
    // Used to disable the display of the firewall and network protection area.
    defenderSecurityCenterDisableNetworkUI *bool
    // Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot the computer for this setting to take effect.
    defenderSecurityCenterDisableNotificationAreaUI *bool
    // Used to disable the display of the ransomware protection area.
    defenderSecurityCenterDisableRansomwareUI *bool
    // Used to disable the display of the secure boot area under Device security.
    defenderSecurityCenterDisableSecureBootUI *bool
    // Used to disable the display of the security process troubleshooting under Device security.
    defenderSecurityCenterDisableTroubleshootingUI *bool
    // Used to disable the display of the virus and threat protection area.
    defenderSecurityCenterDisableVirusUI *bool
    // Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
    defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI *bool
    // The email address that is displayed to users.
    defenderSecurityCenterHelpEmail *string
    // The phone number or Skype ID that is displayed to users.
    defenderSecurityCenterHelpPhone *string
    // The help portal URL this is displayed to users.
    defenderSecurityCenterHelpURL *string
    // Configure where to display IT contact information to end users. Possible values are: notConfigured, displayInAppAndInNotifications, displayOnlyInApp, displayOnlyInNotifications.
    defenderSecurityCenterITContactDisplay *DefenderSecurityCenterITContactDisplayType
    // Notifications to show from the displayed areas of app. Possible values are: notConfigured, blockNoncriticalNotifications, blockAllNotifications.
    defenderSecurityCenterNotificationsFromApp *DefenderSecurityCenterNotificationsFromAppType
    // The company name that is displayed to the users.
    defenderSecurityCenterOrganizationDisplayName *string
    // Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
    defenderSignatureUpdateIntervalInHours *int32
    // Checks for the user consent level in Windows Defender to send data. Possible values are: sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
    defenderSubmitSamplesConsentType *DefenderSubmitSamplesConsentType
    // Value indicating response to executables that don't meet a prevalence, age, or trusted list criteria. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderUntrustedExecutable *DefenderProtectionType
    // Value indicating response to executables that don't meet a prevalence, age, or trusted list criteria. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderUntrustedExecutableType *DefenderAttackSurfaceType
    // Value indicating response to untrusted and unsigned processes that run from USB. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
    defenderUntrustedUSBProcess *DefenderProtectionType
    // Value indicating response to untrusted and unsigned processes that run from USB. Possible values are: userDefined, block, auditMode, warn, disable.
    defenderUntrustedUSBProcessType *DefenderAttackSurfaceType
    // This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA. Specifies whether Platform Security Level is enabled at next reboot.
    deviceGuardEnableSecureBootWithDMA *bool
    // Turns On Virtualization Based Security(VBS).
    deviceGuardEnableVirtualizationBasedSecurity *bool
    // Allows the IT admin to configure the launch of System Guard. Possible values are: notConfigured, enabled, disabled.
    deviceGuardLaunchSystemGuard *Enablement
    // Turn on Credential Guard when Platform Security Level with Secure Boot and Virtualization Based Security are both enabled. Possible values are: notConfigured, enableWithUEFILock, enableWithoutUEFILock, disable.
    deviceGuardLocalSystemAuthorityCredentialGuardSettings *DeviceGuardLocalSystemAuthorityCredentialGuardType
    // Specifies whether Platform Security Level is enabled at next reboot. Possible values are: notConfigured, withoutDMA, withDMA.
    deviceGuardSecureBootWithDMA *SecureBootWithDMAType
    // This policy is intended to provide additional security against external DMA capable devices. It allows for more control over the enumeration of external DMA capable devices incompatible with DMA Remapping/device memory isolation and sandboxing. This policy only takes effect when Kernel DMA Protection is supported and enabled by the system firmware. Kernel DMA Protection is a platform feature that cannot be controlled via policy or by end user. It has to be supported by the system at the time of manufacturing. To check if the system supports Kernel DMA Protection, please check the Kernel DMA Protection field in the Summary page of MSINFO32.exe. Possible values are: deviceDefault, blockAll, allowAll.
    dmaGuardDeviceEnumerationPolicy *DmaGuardDeviceEnumerationPolicyType
    // Blocks stateful FTP connections to the device
    firewallBlockStatefulFTP *bool
    // Specify how the certificate revocation list is to be enforced. Possible values are: deviceDefault, none, attempt, require.
    firewallCertificateRevocationListCheckMethod *FirewallCertificateRevocationListCheckMethodType
    // Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
    firewallIdleTimeoutForSecurityAssociationInSeconds *int32
    // Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
    firewallIPSecExemptionsAllowDHCP *bool
    // Configures IPSec exemptions to allow ICMP
    firewallIPSecExemptionsAllowICMP *bool
    // Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
    firewallIPSecExemptionsAllowNeighborDiscovery *bool
    // Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
    firewallIPSecExemptionsAllowRouterDiscovery *bool
    // Configures IPSec exemptions to no exemptions
    firewallIPSecExemptionsNone *bool
    // If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
    firewallMergeKeyingModuleSettings *bool
    // Configures how packet queueing should be applied in the tunnel gateway scenario. Possible values are: deviceDefault, disabled, queueInbound, queueOutbound, queueBoth.
    firewallPacketQueueingMethod *FirewallPacketQueueingMethodType
    // Select the preshared key encoding to be used. Possible values are: deviceDefault, none, utF8.
    firewallPreSharedKeyEncodingMethod *FirewallPreSharedKeyEncodingMethodType
    // Configures the firewall profile settings for domain networks
    firewallProfileDomain WindowsFirewallNetworkProfileable
    // Configures the firewall profile settings for private networks
    firewallProfilePrivate WindowsFirewallNetworkProfileable
    // Configures the firewall profile settings for public networks
    firewallProfilePublic WindowsFirewallNetworkProfileable
    // Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
    firewallRules []WindowsFirewallRuleable
    // This security setting determines which challenge/response authentication protocol is used for network logons. Possible values are: lmAndNltm, lmNtlmAndNtlmV2, lmAndNtlmOnly, lmAndNtlmV2, lmNtlmV2AndNotLm, lmNtlmV2AndNotLmOrNtm.
    lanManagerAuthenticationLevel *LanManagerAuthenticationLevel
    // If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure guest logons.
    lanManagerWorkstationDisableInsecureGuestLogons *bool
    // Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
    localSecurityOptionsAdministratorAccountName *string
    // Define the behavior of the elevation prompt for admins in Admin Approval Mode. Possible values are: notConfigured, elevateWithoutPrompting, promptForCredentialsOnTheSecureDesktop, promptForConsentOnTheSecureDesktop, promptForCredentials, promptForConsent, promptForConsentForNonWindowsBinaries.
    localSecurityOptionsAdministratorElevationPromptBehavior *LocalSecurityOptionsAdministratorElevationPromptBehaviorType
    // This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating the names of domain accounts and network shares.
    localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares *bool
    // Block PKU2U authentication requests to this device to use online identities.
    localSecurityOptionsAllowPKU2UAuthenticationRequests *bool
    // Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote calls to the SAM.
    localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager *string
    // UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
    localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool *bool
    // This security setting determines whether a computer can be shut down without having to log on to Windows.
    localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn *bool
    // Allow UIAccess apps to prompt for elevation without using the secure desktop.
    localSecurityOptionsAllowUIAccessApplicationElevation *bool
    // Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
    localSecurityOptionsAllowUIAccessApplicationsForSecureLocations *bool
    // Prevent a portable computer from being undocked without having to log in.
    localSecurityOptionsAllowUndockWithoutHavingToLogon *bool
    // Prevent users from adding new Microsoft accounts to this computer.
    localSecurityOptionsBlockMicrosoftAccounts *bool
    // Enable Local accounts that are not password protected to log on from locations other than the physical device.Default is enabled
    localSecurityOptionsBlockRemoteLogonWithBlankPassword *bool
    // Enabling this settings allows only interactively logged on user to access CD-ROM media.
    localSecurityOptionsBlockRemoteOpticalDriveAccess *bool
    // Restrict installing printer drivers as part of connecting to a shared printer to admins only.
    localSecurityOptionsBlockUsersInstallingPrinterDrivers *bool
    // This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
    localSecurityOptionsClearVirtualMemoryPageFile *bool
    // This security setting determines whether packet signing is required by the SMB client component.
    localSecurityOptionsClientDigitallySignCommunicationsAlways *bool
    // If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords to non-Microsoft SMB servers that do not support password encryption during authentication.
    localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers *bool
    // App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
    localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation *bool
    // Determines whether the Local Administrator account is enabled or disabled.
    localSecurityOptionsDisableAdministratorAccount *bool
    // This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
    localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees *bool
    // Determines if the Guest account is enabled or disabled.
    localSecurityOptionsDisableGuestAccount *bool
    // This security setting determines whether packet signing is required by the SMB server component.
    localSecurityOptionsDisableServerDigitallySignCommunicationsAlways *bool
    // This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request it.
    localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees *bool
    // This security setting determines what additional permissions will be granted for anonymous connections to the computer.
    localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts *bool
    // Require CTRL+ALT+DEL to be pressed before a user can log on.
    localSecurityOptionsDoNotRequireCtrlAltDel *bool
    // This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new password is stored. It’s not stored by default.
    localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange *bool
    // Define who is allowed to format and eject removable NTFS media. Possible values are: notConfigured, administrators, administratorsAndPowerUsers, administratorsAndInteractiveUsers.
    localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType
    // Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
    localSecurityOptionsGuestAccountName *string
    // Do not display the username of the last person who signed in on this device.
    localSecurityOptionsHideLastSignedInUser *bool
    // Do not display the username of the person signing in to this device after credentials are entered and before the device’s desktop is shown.
    localSecurityOptionsHideUsernameAtSignIn *bool
    // Configure the user information that is displayed when the session is locked. If not configured, user display name, domain and username are shown. Possible values are: notConfigured, administrators, administratorsAndPowerUsers, administratorsAndInteractiveUsers.
    localSecurityOptionsInformationDisplayedOnLockScreen *LocalSecurityOptionsInformationDisplayedOnLockScreenType
    // Configure the user information that is displayed when the session is locked. If not configured, user display name, domain and username are shown. Possible values are: notConfigured, userDisplayNameDomainUser, userDisplayNameOnly, doNotDisplayUser.
    localSecurityOptionsInformationShownOnLockScreen *LocalSecurityOptionsInformationShownOnLockScreenType
    // Set message text for users attempting to log in.
    localSecurityOptionsLogOnMessageText *string
    // Set message title for users attempting to log in.
    localSecurityOptionsLogOnMessageTitle *string
    // Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
    localSecurityOptionsMachineInactivityLimit *int32
    // Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
    localSecurityOptionsMachineInactivityLimitInMinutes *int32
    // This security setting allows a client to require the negotiation of 128-bit encryption and/or NTLMv2 session security. Possible values are: none, requireNtmlV2SessionSecurity, require128BitEncryption, ntlmV2And128BitEncryption.
    localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients *LocalSecurityOptionsMinimumSessionSecurity
    // This security setting allows a server to require the negotiation of 128-bit encryption and/or NTLMv2 session security. Possible values are: none, requireNtmlV2SessionSecurity, require128BitEncryption, ntlmV2And128BitEncryption.
    localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers *LocalSecurityOptionsMinimumSessionSecurity
    // Enforce PKI certification path validation for a given executable file before it is permitted to run.
    localSecurityOptionsOnlyElevateSignedExecutables *bool
    // By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that can be accessed anonymously and Shares that can be accessed anonymously
    localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares *bool
    // This security setting determines what happens when the smart card for a logged-on user is removed from the smart card reader. Possible values are: noAction, lockWorkstation, forceLogoff, disconnectRemoteDesktopSession.
    localSecurityOptionsSmartCardRemovalBehavior *LocalSecurityOptionsSmartCardRemovalBehaviorType
    // Define the behavior of the elevation prompt for standard users. Possible values are: notConfigured, automaticallyDenyElevationRequests, promptForCredentialsOnTheSecureDesktop, promptForCredentials.
    localSecurityOptionsStandardUserElevationPromptBehavior *LocalSecurityOptionsStandardUserElevationPromptBehaviorType
    // Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior policy settings for admins and standard users are used.
    localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation *bool
    // Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin privileges.Default is enabled
    localSecurityOptionsUseAdminApprovalMode *bool
    // Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
    localSecurityOptionsUseAdminApprovalModeForAdministrators *bool
    // Virtualize file and registry write failures to per user locations
    localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations *bool
    // Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
    smartScreenBlockOverrideForFiles *bool
    // Allows IT Admins to configure SmartScreen for Windows.
    smartScreenEnableInShell *bool
    // This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if this privilege is given to other entities. Only states NotConfigured and Allowed are supported
    userRightsAccessCredentialManagerAsTrustedCaller DeviceManagementUserRightsSettingable
    // This user right allows a process to impersonate any user without authentication. The process can therefore gain access to the same local resources as that user. Only states NotConfigured and Allowed are supported
    userRightsActAsPartOfTheOperatingSystem DeviceManagementUserRightsSettingable
    // This user right determines which users and groups are allowed to connect to the computer over the network. State Allowed is supported.
    userRightsAllowAccessFromNetwork DeviceManagementUserRightsSettingable
    // This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when backing up files and directories. Only states NotConfigured and Allowed are supported
    userRightsBackupData DeviceManagementUserRightsSettingable
    // This user right determines which users and groups are block from connecting to the computer over the network. State Block is supported.
    userRightsBlockAccessFromNetwork DeviceManagementUserRightsSettingable
    // This user right determines which users and groups can change the time and date on the internal clock of the computer. Only states NotConfigured and Allowed are supported
    userRightsChangeSystemTime DeviceManagementUserRightsSettingable
    // This security setting determines whether users can create global objects that are available to all sessions. Users who can create global objects could affect processes that run under other users' sessions, which could lead to application failure or data corruption. Only states NotConfigured and Allowed are supported
    userRightsCreateGlobalObjects DeviceManagementUserRightsSettingable
    // This user right determines which users and groups can call an internal API to create and change the size of a page file. Only states NotConfigured and Allowed are supported
    userRightsCreatePageFile DeviceManagementUserRightsSettingable
    // This user right determines which accounts can be used by processes to create a directory object using the object manager. Only states NotConfigured and Allowed are supported
    userRightsCreatePermanentSharedObjects DeviceManagementUserRightsSettingable
    // This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only states NotConfigured and Allowed are supported
    userRightsCreateSymbolicLinks DeviceManagementUserRightsSettingable
    // This user right determines which users/groups can be used by processes to create a token that can then be used to get access to any local resources when the process uses an internal API to create an access token. Only states NotConfigured and Allowed are supported
    userRightsCreateToken DeviceManagementUserRightsSettingable
    // This user right determines which users can attach a debugger to any process or to the kernel. Only states NotConfigured and Allowed are supported
    userRightsDebugPrograms DeviceManagementUserRightsSettingable
    // This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only states NotConfigured and Allowed are supported.
    userRightsDelegation DeviceManagementUserRightsSettingable
    // This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
    userRightsDenyLocalLogOn DeviceManagementUserRightsSettingable
    // This user right determines which accounts can be used by a process to add entries to the security log. The security log is used to trace unauthorized system access.  Only states NotConfigured and Allowed are supported.
    userRightsGenerateSecurityAudits DeviceManagementUserRightsSettingable
    // Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a service that they have created and then impersonating that client, which can elevate the unauthorized user's permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
    userRightsImpersonateClient DeviceManagementUserRightsSettingable
    // This user right determines which accounts can use a process with Write Property access to another process to increase the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
    userRightsIncreaseSchedulingPriority DeviceManagementUserRightsSettingable
    // This user right determines which users can dynamically load and unload device drivers or other code in to kernel mode. Only states NotConfigured and Allowed are supported.
    userRightsLoadUnloadDrivers DeviceManagementUserRightsSettingable
    // This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
    userRightsLocalLogOn DeviceManagementUserRightsSettingable
    // This user right determines which accounts can use a process to keep data in physical memory, which prevents the system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
    userRightsLockMemory DeviceManagementUserRightsSettingable
    // This user right determines which users can specify object access auditing options for individual resources, such as files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
    userRightsManageAuditingAndSecurityLogs DeviceManagementUserRightsSettingable
    // This user right determines which users and groups can run maintenance tasks on a volume, such as remote defragmentation. Only states NotConfigured and Allowed are supported.
    userRightsManageVolumes DeviceManagementUserRightsSettingable
    // This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are supported.
    userRightsModifyFirmwareEnvironment DeviceManagementUserRightsSettingable
    // This user right determines which user accounts can modify the integrity label of objects, such as files, registry keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
    userRightsModifyObjectLabels DeviceManagementUserRightsSettingable
    // This user right determines which users can use performance monitoring tools to monitor the performance of system processes. Only states NotConfigured and Allowed are supported.
    userRightsProfileSingleProcess DeviceManagementUserRightsSettingable
    // This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client. Only states NotConfigured and Blocked are supported
    userRightsRemoteDesktopServicesLogOn DeviceManagementUserRightsSettingable
    // This user right determines which users are allowed to shut down a computer from a remote location on the network. Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
    userRightsRemoteShutdown DeviceManagementUserRightsSettingable
    // This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when restoring backed up files and directories, and determines which users can set any valid security principal as the owner of an object. Only states NotConfigured and Allowed are supported.
    userRightsRestoreData DeviceManagementUserRightsSettingable
    // This user right determines which users can take ownership of any securable object in the system, including Active Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and Allowed are supported.
    userRightsTakeOwnership DeviceManagementUserRightsSettingable
    // Configure windows defender TamperProtection settings. Possible values are: notConfigured, enable, disable.
    windowsDefenderTamperProtection *WindowsDefenderTamperProtectionOptions
    // This setting determines whether the Accessory management service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
    xboxServicesAccessoryManagementServiceStartupMode *ServiceStartType
    // This setting determines whether xbox game save is enabled (1) or disabled (0).
    xboxServicesEnableXboxGameSaveTask *bool
    // This setting determines whether Live Auth Manager service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
    xboxServicesLiveAuthManagerServiceStartupMode *ServiceStartType
    // This setting determines whether Live Game save service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
    xboxServicesLiveGameSaveServiceStartupMode *ServiceStartType
    // This setting determines whether Networking service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
    xboxServicesLiveNetworkingServiceStartupMode *ServiceStartType
}
// NewWindows10EndpointProtectionConfiguration instantiates a new Windows10EndpointProtectionConfiguration and sets the default values.
func NewWindows10EndpointProtectionConfiguration()(*Windows10EndpointProtectionConfiguration) {
    m := &Windows10EndpointProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindows10EndpointProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EndpointProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EndpointProtectionConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10EndpointProtectionConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationGuardAllowCameraMicrophoneRedirection gets the applicationGuardAllowCameraMicrophoneRedirection property value. Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and microphone.
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowCameraMicrophoneRedirection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardAllowCameraMicrophoneRedirection
    }
}
// GetApplicationGuardAllowFileSaveOnHost gets the applicationGuardAllowFileSaveOnHost property value. Allow users to download files from Edge in the application guard container and save them on the host file system
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowFileSaveOnHost()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardAllowFileSaveOnHost
    }
}
// GetApplicationGuardAllowPersistence gets the applicationGuardAllowPersistence property value. Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPersistence()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardAllowPersistence
    }
}
// GetApplicationGuardAllowPrintToLocalPrinters gets the applicationGuardAllowPrintToLocalPrinters property value. Allow printing to Local Printers from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToLocalPrinters()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardAllowPrintToLocalPrinters
    }
}
// GetApplicationGuardAllowPrintToNetworkPrinters gets the applicationGuardAllowPrintToNetworkPrinters property value. Allow printing to Network Printers from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToNetworkPrinters()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardAllowPrintToNetworkPrinters
    }
}
// GetApplicationGuardAllowPrintToPDF gets the applicationGuardAllowPrintToPDF property value. Allow printing to PDF from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToPDF()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardAllowPrintToPDF
    }
}
// GetApplicationGuardAllowPrintToXPS gets the applicationGuardAllowPrintToXPS property value. Allow printing to XPS from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToXPS()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardAllowPrintToXPS
    }
}
// GetApplicationGuardAllowVirtualGPU gets the applicationGuardAllowVirtualGPU property value. Allow application guard to use virtual GPU
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowVirtualGPU()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardAllowVirtualGPU
    }
}
// GetApplicationGuardBlockClipboardSharing gets the applicationGuardBlockClipboardSharing property value. Block clipboard to share data from Host to Container, or from Container to Host, or both ways, or neither ways. Possible values are: notConfigured, blockBoth, blockHostToContainer, blockContainerToHost, blockNone.
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockClipboardSharing()(*ApplicationGuardBlockClipboardSharingType) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardBlockClipboardSharing
    }
}
// GetApplicationGuardBlockFileTransfer gets the applicationGuardBlockFileTransfer property value. Block clipboard to transfer image file, text file or neither of them. Possible values are: notConfigured, blockImageAndTextFile, blockImageFile, blockNone, blockTextFile.
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockFileTransfer()(*ApplicationGuardBlockFileTransferType) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardBlockFileTransfer
    }
}
// GetApplicationGuardBlockNonEnterpriseContent gets the applicationGuardBlockNonEnterpriseContent property value. Block enterprise sites to load non-enterprise content, such as third party plug-ins
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockNonEnterpriseContent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardBlockNonEnterpriseContent
    }
}
// GetApplicationGuardCertificateThumbprints gets the applicationGuardCertificateThumbprints property value. Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardCertificateThumbprints()([]string) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardCertificateThumbprints
    }
}
// GetApplicationGuardEnabled gets the applicationGuardEnabled property value. Enable Windows Defender Application Guard
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardEnabled
    }
}
// GetApplicationGuardEnabledOptions gets the applicationGuardEnabledOptions property value. Enable Windows Defender Application Guard for newer Windows builds. Possible values are: notConfigured, enabledForEdge, enabledForOffice, enabledForEdgeAndOffice.
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardEnabledOptions()(*ApplicationGuardEnabledOptions) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardEnabledOptions
    }
}
// GetApplicationGuardForceAuditing gets the applicationGuardForceAuditing property value. Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardForceAuditing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applicationGuardForceAuditing
    }
}
// GetAppLockerApplicationControl gets the appLockerApplicationControl property value. Enables the Admin to choose what types of app to allow on devices. Possible values are: notConfigured, enforceComponentsAndStoreApps, auditComponentsAndStoreApps, enforceComponentsStoreAppsAndSmartlocker, auditComponentsStoreAppsAndSmartlocker.
func (m *Windows10EndpointProtectionConfiguration) GetAppLockerApplicationControl()(*AppLockerApplicationControlType) {
    if m == nil {
        return nil
    } else {
        return m.appLockerApplicationControl
    }
}
// GetBitLockerAllowStandardUserEncryption gets the bitLockerAllowStandardUserEncryption property value. Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerAllowStandardUserEncryption()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerAllowStandardUserEncryption
    }
}
// GetBitLockerDisableWarningForOtherDiskEncryption gets the bitLockerDisableWarningForOtherDiskEncryption property value. Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerDisableWarningForOtherDiskEncryption()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerDisableWarningForOtherDiskEncryption
    }
}
// GetBitLockerEnableStorageCardEncryptionOnMobile gets the bitLockerEnableStorageCardEncryptionOnMobile property value. Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerEnableStorageCardEncryptionOnMobile()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerEnableStorageCardEncryptionOnMobile
    }
}
// GetBitLockerEncryptDevice gets the bitLockerEncryptDevice property value. Allows the admin to require encryption to be turned on using BitLocker.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerEncryptDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerEncryptDevice
    }
}
// GetBitLockerFixedDrivePolicy gets the bitLockerFixedDrivePolicy property value. BitLocker Fixed Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerFixedDrivePolicy()(BitLockerFixedDrivePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerFixedDrivePolicy
    }
}
// GetBitLockerRecoveryPasswordRotation gets the bitLockerRecoveryPasswordRotation property value. This setting initiates a client-driven recovery password rotation after an OS drive recovery (either by using bootmgr or WinRE). Possible values are: notConfigured, disabled, enabledForAzureAd, enabledForAzureAdAndHybrid.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerRecoveryPasswordRotation()(*BitLockerRecoveryPasswordRotationType) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerRecoveryPasswordRotation
    }
}
// GetBitLockerRemovableDrivePolicy gets the bitLockerRemovableDrivePolicy property value. BitLocker Removable Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerRemovableDrivePolicy()(BitLockerRemovableDrivePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerRemovableDrivePolicy
    }
}
// GetBitLockerSystemDrivePolicy gets the bitLockerSystemDrivePolicy property value. BitLocker System Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerSystemDrivePolicy()(BitLockerSystemDrivePolicyable) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerSystemDrivePolicy
    }
}
// GetDefenderAdditionalGuardedFolders gets the defenderAdditionalGuardedFolders property value. List of folder paths to be added to the list of protected folders
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdditionalGuardedFolders()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defenderAdditionalGuardedFolders
    }
}
// GetDefenderAdobeReaderLaunchChildProcess gets the defenderAdobeReaderLaunchChildProcess property value. Value indicating the behavior of Adobe Reader from creating child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdobeReaderLaunchChildProcess()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderAdobeReaderLaunchChildProcess
    }
}
// GetDefenderAdvancedRansomewareProtectionType gets the defenderAdvancedRansomewareProtectionType property value. Value indicating use of advanced protection against ransomeware. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdvancedRansomewareProtectionType()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderAdvancedRansomewareProtectionType
    }
}
// GetDefenderAllowBehaviorMonitoring gets the defenderAllowBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowBehaviorMonitoring()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowBehaviorMonitoring
    }
}
// GetDefenderAllowCloudProtection gets the defenderAllowCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowCloudProtection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowCloudProtection
    }
}
// GetDefenderAllowEndUserAccess gets the defenderAllowEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowEndUserAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowEndUserAccess
    }
}
// GetDefenderAllowIntrusionPreventionSystem gets the defenderAllowIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowIntrusionPreventionSystem()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowIntrusionPreventionSystem
    }
}
// GetDefenderAllowOnAccessProtection gets the defenderAllowOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowOnAccessProtection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowOnAccessProtection
    }
}
// GetDefenderAllowRealTimeMonitoring gets the defenderAllowRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowRealTimeMonitoring()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowRealTimeMonitoring
    }
}
// GetDefenderAllowScanArchiveFiles gets the defenderAllowScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanArchiveFiles()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowScanArchiveFiles
    }
}
// GetDefenderAllowScanDownloads gets the defenderAllowScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanDownloads()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowScanDownloads
    }
}
// GetDefenderAllowScanNetworkFiles gets the defenderAllowScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanNetworkFiles()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowScanNetworkFiles
    }
}
// GetDefenderAllowScanRemovableDrivesDuringFullScan gets the defenderAllowScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanRemovableDrivesDuringFullScan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowScanRemovableDrivesDuringFullScan
    }
}
// GetDefenderAllowScanScriptsLoadedInInternetExplorer gets the defenderAllowScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanScriptsLoadedInInternetExplorer()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderAllowScanScriptsLoadedInInternetExplorer
    }
}
// GetDefenderAttackSurfaceReductionExcludedPaths gets the defenderAttackSurfaceReductionExcludedPaths property value. List of exe files and folders to be excluded from attack surface reduction rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAttackSurfaceReductionExcludedPaths()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defenderAttackSurfaceReductionExcludedPaths
    }
}
// GetDefenderBlockEndUserAccess gets the defenderBlockEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderBlockEndUserAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderBlockEndUserAccess
    }
}
// GetDefenderBlockPersistenceThroughWmiType gets the defenderBlockPersistenceThroughWmiType property value. Value indicating the behavior of Block persistence through WMI event subscription. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderBlockPersistenceThroughWmiType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderBlockPersistenceThroughWmiType
    }
}
// GetDefenderCheckForSignaturesBeforeRunningScan gets the defenderCheckForSignaturesBeforeRunningScan property value. This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before running a scan.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCheckForSignaturesBeforeRunningScan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderCheckForSignaturesBeforeRunningScan
    }
}
// GetDefenderCloudBlockLevel gets the defenderCloudBlockLevel property value. Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS' setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCloudBlockLevel()(*DefenderCloudBlockLevelType) {
    if m == nil {
        return nil
    } else {
        return m.defenderCloudBlockLevel
    }
}
// GetDefenderCloudExtendedTimeoutInSeconds gets the defenderCloudExtendedTimeoutInSeconds property value. Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCloudExtendedTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.defenderCloudExtendedTimeoutInSeconds
    }
}
// GetDefenderDaysBeforeDeletingQuarantinedMalware gets the defenderDaysBeforeDeletingQuarantinedMalware property value. Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDaysBeforeDeletingQuarantinedMalware()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.defenderDaysBeforeDeletingQuarantinedMalware
    }
}
// GetDefenderDetectedMalwareActions gets the defenderDetectedMalwareActions property value. Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDetectedMalwareActions()(DefenderDetectedMalwareActionsable) {
    if m == nil {
        return nil
    } else {
        return m.defenderDetectedMalwareActions
    }
}
// GetDefenderDisableBehaviorMonitoring gets the defenderDisableBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableBehaviorMonitoring()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableBehaviorMonitoring
    }
}
// GetDefenderDisableCatchupFullScan gets the defenderDisableCatchupFullScan property value. This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCatchupFullScan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableCatchupFullScan
    }
}
// GetDefenderDisableCatchupQuickScan gets the defenderDisableCatchupQuickScan property value. This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCatchupQuickScan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableCatchupQuickScan
    }
}
// GetDefenderDisableCloudProtection gets the defenderDisableCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCloudProtection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableCloudProtection
    }
}
// GetDefenderDisableIntrusionPreventionSystem gets the defenderDisableIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableIntrusionPreventionSystem()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableIntrusionPreventionSystem
    }
}
// GetDefenderDisableOnAccessProtection gets the defenderDisableOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableOnAccessProtection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableOnAccessProtection
    }
}
// GetDefenderDisableRealTimeMonitoring gets the defenderDisableRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableRealTimeMonitoring()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableRealTimeMonitoring
    }
}
// GetDefenderDisableScanArchiveFiles gets the defenderDisableScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanArchiveFiles()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableScanArchiveFiles
    }
}
// GetDefenderDisableScanDownloads gets the defenderDisableScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanDownloads()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableScanDownloads
    }
}
// GetDefenderDisableScanNetworkFiles gets the defenderDisableScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanNetworkFiles()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableScanNetworkFiles
    }
}
// GetDefenderDisableScanRemovableDrivesDuringFullScan gets the defenderDisableScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanRemovableDrivesDuringFullScan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableScanRemovableDrivesDuringFullScan
    }
}
// GetDefenderDisableScanScriptsLoadedInInternetExplorer gets the defenderDisableScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanScriptsLoadedInInternetExplorer()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderDisableScanScriptsLoadedInInternetExplorer
    }
}
// GetDefenderEmailContentExecution gets the defenderEmailContentExecution property value. Value indicating if execution of executable content (exe, dll, ps, js, vbs, etc) should be dropped from email (webmail/mail-client). Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEmailContentExecution()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderEmailContentExecution
    }
}
// GetDefenderEmailContentExecutionType gets the defenderEmailContentExecutionType property value. Value indicating if execution of executable content (exe, dll, ps, js, vbs, etc) should be dropped from email (webmail/mail-client). Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEmailContentExecutionType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderEmailContentExecutionType
    }
}
// GetDefenderEnableLowCpuPriority gets the defenderEnableLowCpuPriority property value. This policy setting allows you to enable or disable low CPU priority for scheduled scans.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableLowCpuPriority()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderEnableLowCpuPriority
    }
}
// GetDefenderEnableScanIncomingMail gets the defenderEnableScanIncomingMail property value. Allows or disallows scanning of email.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableScanIncomingMail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderEnableScanIncomingMail
    }
}
// GetDefenderEnableScanMappedNetworkDrivesDuringFullScan gets the defenderEnableScanMappedNetworkDrivesDuringFullScan property value. Allows or disallows a full scan of mapped network drives.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableScanMappedNetworkDrivesDuringFullScan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderEnableScanMappedNetworkDrivesDuringFullScan
    }
}
// GetDefenderExploitProtectionXml gets the defenderExploitProtectionXml property value. Xml content containing information regarding exploit protection details.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.defenderExploitProtectionXml
    }
}
// GetDefenderExploitProtectionXmlFileName gets the defenderExploitProtectionXmlFileName property value. Name of the file from which DefenderExploitProtectionXml was obtained.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXmlFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defenderExploitProtectionXmlFileName
    }
}
// GetDefenderFileExtensionsToExclude gets the defenderFileExtensionsToExclude property value. File extensions to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderFileExtensionsToExclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defenderFileExtensionsToExclude
    }
}
// GetDefenderFilesAndFoldersToExclude gets the defenderFilesAndFoldersToExclude property value. Files and folder to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderFilesAndFoldersToExclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defenderFilesAndFoldersToExclude
    }
}
// GetDefenderGuardedFoldersAllowedAppPaths gets the defenderGuardedFoldersAllowedAppPaths property value. List of paths to exe that are allowed to access protected folders
func (m *Windows10EndpointProtectionConfiguration) GetDefenderGuardedFoldersAllowedAppPaths()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defenderGuardedFoldersAllowedAppPaths
    }
}
// GetDefenderGuardMyFoldersType gets the defenderGuardMyFoldersType property value. Value indicating the behavior of protected folders. Possible values are: userDefined, enable, auditMode, blockDiskModification, auditDiskModification.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderGuardMyFoldersType()(*FolderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderGuardMyFoldersType
    }
}
// GetDefenderNetworkProtectionType gets the defenderNetworkProtectionType property value. Value indicating the behavior of NetworkProtection. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderNetworkProtectionType()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderNetworkProtectionType
    }
}
// GetDefenderOfficeAppsExecutableContentCreationOrLaunch gets the defenderOfficeAppsExecutableContentCreationOrLaunch property value. Value indicating the behavior of Office applications/macros creating or launching executable content. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsExecutableContentCreationOrLaunch()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeAppsExecutableContentCreationOrLaunch
    }
}
// GetDefenderOfficeAppsExecutableContentCreationOrLaunchType gets the defenderOfficeAppsExecutableContentCreationOrLaunchType property value. Value indicating the behavior of Office applications/macros creating or launching executable content. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsExecutableContentCreationOrLaunchType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeAppsExecutableContentCreationOrLaunchType
    }
}
// GetDefenderOfficeAppsLaunchChildProcess gets the defenderOfficeAppsLaunchChildProcess property value. Value indicating the behavior of Office application launching child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsLaunchChildProcess()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeAppsLaunchChildProcess
    }
}
// GetDefenderOfficeAppsLaunchChildProcessType gets the defenderOfficeAppsLaunchChildProcessType property value. Value indicating the behavior of Office application launching child processes. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsLaunchChildProcessType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeAppsLaunchChildProcessType
    }
}
// GetDefenderOfficeAppsOtherProcessInjection gets the defenderOfficeAppsOtherProcessInjection property value. Value indicating the behavior of  Office applications injecting into other processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsOtherProcessInjection()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeAppsOtherProcessInjection
    }
}
// GetDefenderOfficeAppsOtherProcessInjectionType gets the defenderOfficeAppsOtherProcessInjectionType property value. Value indicating the behavior of Office applications injecting into other processes. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsOtherProcessInjectionType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeAppsOtherProcessInjectionType
    }
}
// GetDefenderOfficeCommunicationAppsLaunchChildProcess gets the defenderOfficeCommunicationAppsLaunchChildProcess property value. Value indicating the behavior of Office communication applications, including Microsoft Outlook, from creating child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeCommunicationAppsLaunchChildProcess()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeCommunicationAppsLaunchChildProcess
    }
}
// GetDefenderOfficeMacroCodeAllowWin32Imports gets the defenderOfficeMacroCodeAllowWin32Imports property value. Value indicating the behavior of Win32 imports from Macro code in Office. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeMacroCodeAllowWin32Imports()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeMacroCodeAllowWin32Imports
    }
}
// GetDefenderOfficeMacroCodeAllowWin32ImportsType gets the defenderOfficeMacroCodeAllowWin32ImportsType property value. Value indicating the behavior of Win32 imports from Macro code in Office. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeMacroCodeAllowWin32ImportsType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderOfficeMacroCodeAllowWin32ImportsType
    }
}
// GetDefenderPotentiallyUnwantedAppAction gets the defenderPotentiallyUnwantedAppAction property value. Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs). Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderPotentiallyUnwantedAppAction()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderPotentiallyUnwantedAppAction
    }
}
// GetDefenderPreventCredentialStealingType gets the defenderPreventCredentialStealingType property value. Value indicating if credential stealing from the Windows local security authority subsystem is permitted. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderPreventCredentialStealingType()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderPreventCredentialStealingType
    }
}
// GetDefenderProcessCreation gets the defenderProcessCreation property value. Value indicating response to process creations originating from PSExec and WMI commands. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessCreation()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderProcessCreation
    }
}
// GetDefenderProcessCreationType gets the defenderProcessCreationType property value. Value indicating response to process creations originating from PSExec and WMI commands. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessCreationType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderProcessCreationType
    }
}
// GetDefenderProcessesToExclude gets the defenderProcessesToExclude property value. Processes to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessesToExclude()([]string) {
    if m == nil {
        return nil
    } else {
        return m.defenderProcessesToExclude
    }
}
// GetDefenderScanDirection gets the defenderScanDirection property value. Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly, monitorOutgoingFilesOnly.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanDirection()(*DefenderRealtimeScanDirection) {
    if m == nil {
        return nil
    } else {
        return m.defenderScanDirection
    }
}
// GetDefenderScanMaxCpuPercentage gets the defenderScanMaxCpuPercentage property value. Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid values 0 to 100
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanMaxCpuPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.defenderScanMaxCpuPercentage
    }
}
// GetDefenderScanType gets the defenderScanType property value. Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanType()(*DefenderScanType) {
    if m == nil {
        return nil
    } else {
        return m.defenderScanType
    }
}
// GetDefenderScheduledQuickScanTime gets the defenderScheduledQuickScanTime property value. Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledQuickScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.defenderScheduledQuickScanTime
    }
}
// GetDefenderScheduledScanDay gets the defenderScheduledScanDay property value. Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledScanDay()(*WeeklySchedule) {
    if m == nil {
        return nil
    } else {
        return m.defenderScheduledScanDay
    }
}
// GetDefenderScheduledScanTime gets the defenderScheduledScanTime property value. Selects the time of day that the Windows Defender scan should run.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.defenderScheduledScanTime
    }
}
// GetDefenderScriptDownloadedPayloadExecution gets the defenderScriptDownloadedPayloadExecution property value. Value indicating the behavior of js/vbs executing payload downloaded from Internet. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptDownloadedPayloadExecution()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderScriptDownloadedPayloadExecution
    }
}
// GetDefenderScriptDownloadedPayloadExecutionType gets the defenderScriptDownloadedPayloadExecutionType property value. Value indicating the behavior of js/vbs executing payload downloaded from Internet. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptDownloadedPayloadExecutionType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderScriptDownloadedPayloadExecutionType
    }
}
// GetDefenderScriptObfuscatedMacroCode gets the defenderScriptObfuscatedMacroCode property value. Value indicating the behavior of obfuscated js/vbs/ps/macro code. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptObfuscatedMacroCode()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderScriptObfuscatedMacroCode
    }
}
// GetDefenderScriptObfuscatedMacroCodeType gets the defenderScriptObfuscatedMacroCodeType property value. Value indicating the behavior of obfuscated js/vbs/ps/macro code. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptObfuscatedMacroCodeType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderScriptObfuscatedMacroCodeType
    }
}
// GetDefenderSecurityCenterBlockExploitProtectionOverride gets the defenderSecurityCenterBlockExploitProtectionOverride property value. Indicates whether or not to block user from overriding Exploit Protection settings.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterBlockExploitProtectionOverride()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterBlockExploitProtectionOverride
    }
}
// GetDefenderSecurityCenterDisableAccountUI gets the defenderSecurityCenterDisableAccountUI property value. Used to disable the display of the account protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableAccountUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableAccountUI
    }
}
// GetDefenderSecurityCenterDisableAppBrowserUI gets the defenderSecurityCenterDisableAppBrowserUI property value. Used to disable the display of the app and browser protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableAppBrowserUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableAppBrowserUI
    }
}
// GetDefenderSecurityCenterDisableClearTpmUI gets the defenderSecurityCenterDisableClearTpmUI property value. Used to disable the display of the Clear TPM button.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableClearTpmUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableClearTpmUI
    }
}
// GetDefenderSecurityCenterDisableFamilyUI gets the defenderSecurityCenterDisableFamilyUI property value. Used to disable the display of the family options area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableFamilyUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableFamilyUI
    }
}
// GetDefenderSecurityCenterDisableHardwareUI gets the defenderSecurityCenterDisableHardwareUI property value. Used to disable the display of the hardware protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableHardwareUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableHardwareUI
    }
}
// GetDefenderSecurityCenterDisableHealthUI gets the defenderSecurityCenterDisableHealthUI property value. Used to disable the display of the device performance and health area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableHealthUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableHealthUI
    }
}
// GetDefenderSecurityCenterDisableNetworkUI gets the defenderSecurityCenterDisableNetworkUI property value. Used to disable the display of the firewall and network protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableNetworkUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableNetworkUI
    }
}
// GetDefenderSecurityCenterDisableNotificationAreaUI gets the defenderSecurityCenterDisableNotificationAreaUI property value. Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot the computer for this setting to take effect.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableNotificationAreaUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableNotificationAreaUI
    }
}
// GetDefenderSecurityCenterDisableRansomwareUI gets the defenderSecurityCenterDisableRansomwareUI property value. Used to disable the display of the ransomware protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableRansomwareUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableRansomwareUI
    }
}
// GetDefenderSecurityCenterDisableSecureBootUI gets the defenderSecurityCenterDisableSecureBootUI property value. Used to disable the display of the secure boot area under Device security.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableSecureBootUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableSecureBootUI
    }
}
// GetDefenderSecurityCenterDisableTroubleshootingUI gets the defenderSecurityCenterDisableTroubleshootingUI property value. Used to disable the display of the security process troubleshooting under Device security.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableTroubleshootingUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableTroubleshootingUI
    }
}
// GetDefenderSecurityCenterDisableVirusUI gets the defenderSecurityCenterDisableVirusUI property value. Used to disable the display of the virus and threat protection area.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableVirusUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableVirusUI
    }
}
// GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI gets the defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI property value. Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI
    }
}
// GetDefenderSecurityCenterHelpEmail gets the defenderSecurityCenterHelpEmail property value. The email address that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterHelpEmail
    }
}
// GetDefenderSecurityCenterHelpPhone gets the defenderSecurityCenterHelpPhone property value. The phone number or Skype ID that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpPhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterHelpPhone
    }
}
// GetDefenderSecurityCenterHelpURL gets the defenderSecurityCenterHelpURL property value. The help portal URL this is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpURL()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterHelpURL
    }
}
// GetDefenderSecurityCenterITContactDisplay gets the defenderSecurityCenterITContactDisplay property value. Configure where to display IT contact information to end users. Possible values are: notConfigured, displayInAppAndInNotifications, displayOnlyInApp, displayOnlyInNotifications.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterITContactDisplay()(*DefenderSecurityCenterITContactDisplayType) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterITContactDisplay
    }
}
// GetDefenderSecurityCenterNotificationsFromApp gets the defenderSecurityCenterNotificationsFromApp property value. Notifications to show from the displayed areas of app. Possible values are: notConfigured, blockNoncriticalNotifications, blockAllNotifications.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterNotificationsFromApp()(*DefenderSecurityCenterNotificationsFromAppType) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterNotificationsFromApp
    }
}
// GetDefenderSecurityCenterOrganizationDisplayName gets the defenderSecurityCenterOrganizationDisplayName property value. The company name that is displayed to the users.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterOrganizationDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defenderSecurityCenterOrganizationDisplayName
    }
}
// GetDefenderSignatureUpdateIntervalInHours gets the defenderSignatureUpdateIntervalInHours property value. Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSignatureUpdateIntervalInHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.defenderSignatureUpdateIntervalInHours
    }
}
// GetDefenderSubmitSamplesConsentType gets the defenderSubmitSamplesConsentType property value. Checks for the user consent level in Windows Defender to send data. Possible values are: sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSubmitSamplesConsentType()(*DefenderSubmitSamplesConsentType) {
    if m == nil {
        return nil
    } else {
        return m.defenderSubmitSamplesConsentType
    }
}
// GetDefenderUntrustedExecutable gets the defenderUntrustedExecutable property value. Value indicating response to executables that don't meet a prevalence, age, or trusted list criteria. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedExecutable()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderUntrustedExecutable
    }
}
// GetDefenderUntrustedExecutableType gets the defenderUntrustedExecutableType property value. Value indicating response to executables that don't meet a prevalence, age, or trusted list criteria. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedExecutableType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderUntrustedExecutableType
    }
}
// GetDefenderUntrustedUSBProcess gets the defenderUntrustedUSBProcess property value. Value indicating response to untrusted and unsigned processes that run from USB. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedUSBProcess()(*DefenderProtectionType) {
    if m == nil {
        return nil
    } else {
        return m.defenderUntrustedUSBProcess
    }
}
// GetDefenderUntrustedUSBProcessType gets the defenderUntrustedUSBProcessType property value. Value indicating response to untrusted and unsigned processes that run from USB. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedUSBProcessType()(*DefenderAttackSurfaceType) {
    if m == nil {
        return nil
    } else {
        return m.defenderUntrustedUSBProcessType
    }
}
// GetDeviceGuardEnableSecureBootWithDMA gets the deviceGuardEnableSecureBootWithDMA property value. This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA. Specifies whether Platform Security Level is enabled at next reboot.
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardEnableSecureBootWithDMA()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardEnableSecureBootWithDMA
    }
}
// GetDeviceGuardEnableVirtualizationBasedSecurity gets the deviceGuardEnableVirtualizationBasedSecurity property value. Turns On Virtualization Based Security(VBS).
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardEnableVirtualizationBasedSecurity()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardEnableVirtualizationBasedSecurity
    }
}
// GetDeviceGuardLaunchSystemGuard gets the deviceGuardLaunchSystemGuard property value. Allows the IT admin to configure the launch of System Guard. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardLaunchSystemGuard()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardLaunchSystemGuard
    }
}
// GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings gets the deviceGuardLocalSystemAuthorityCredentialGuardSettings property value. Turn on Credential Guard when Platform Security Level with Secure Boot and Virtualization Based Security are both enabled. Possible values are: notConfigured, enableWithUEFILock, enableWithoutUEFILock, disable.
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings()(*DeviceGuardLocalSystemAuthorityCredentialGuardType) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardLocalSystemAuthorityCredentialGuardSettings
    }
}
// GetDeviceGuardSecureBootWithDMA gets the deviceGuardSecureBootWithDMA property value. Specifies whether Platform Security Level is enabled at next reboot. Possible values are: notConfigured, withoutDMA, withDMA.
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardSecureBootWithDMA()(*SecureBootWithDMAType) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardSecureBootWithDMA
    }
}
// GetDmaGuardDeviceEnumerationPolicy gets the dmaGuardDeviceEnumerationPolicy property value. This policy is intended to provide additional security against external DMA capable devices. It allows for more control over the enumeration of external DMA capable devices incompatible with DMA Remapping/device memory isolation and sandboxing. This policy only takes effect when Kernel DMA Protection is supported and enabled by the system firmware. Kernel DMA Protection is a platform feature that cannot be controlled via policy or by end user. It has to be supported by the system at the time of manufacturing. To check if the system supports Kernel DMA Protection, please check the Kernel DMA Protection field in the Summary page of MSINFO32.exe. Possible values are: deviceDefault, blockAll, allowAll.
func (m *Windows10EndpointProtectionConfiguration) GetDmaGuardDeviceEnumerationPolicy()(*DmaGuardDeviceEnumerationPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.dmaGuardDeviceEnumerationPolicy
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10EndpointProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["applicationGuardAllowCameraMicrophoneRedirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowCameraMicrophoneRedirection(val)
        }
        return nil
    }
    res["applicationGuardAllowFileSaveOnHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowFileSaveOnHost(val)
        }
        return nil
    }
    res["applicationGuardAllowPersistence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPersistence(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToLocalPrinters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToLocalPrinters(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToNetworkPrinters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToNetworkPrinters(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToPDF"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToPDF(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToXPS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToXPS(val)
        }
        return nil
    }
    res["applicationGuardAllowVirtualGPU"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowVirtualGPU(val)
        }
        return nil
    }
    res["applicationGuardBlockClipboardSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationGuardBlockClipboardSharingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockClipboardSharing(val.(*ApplicationGuardBlockClipboardSharingType))
        }
        return nil
    }
    res["applicationGuardBlockFileTransfer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationGuardBlockFileTransferType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockFileTransfer(val.(*ApplicationGuardBlockFileTransferType))
        }
        return nil
    }
    res["applicationGuardBlockNonEnterpriseContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockNonEnterpriseContent(val)
        }
        return nil
    }
    res["applicationGuardCertificateThumbprints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetApplicationGuardCertificateThumbprints(res)
        }
        return nil
    }
    res["applicationGuardEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardEnabled(val)
        }
        return nil
    }
    res["applicationGuardEnabledOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationGuardEnabledOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardEnabledOptions(val.(*ApplicationGuardEnabledOptions))
        }
        return nil
    }
    res["applicationGuardForceAuditing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardForceAuditing(val)
        }
        return nil
    }
    res["appLockerApplicationControl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLockerApplicationControlType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppLockerApplicationControl(val.(*AppLockerApplicationControlType))
        }
        return nil
    }
    res["bitLockerAllowStandardUserEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerAllowStandardUserEncryption(val)
        }
        return nil
    }
    res["bitLockerDisableWarningForOtherDiskEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerDisableWarningForOtherDiskEncryption(val)
        }
        return nil
    }
    res["bitLockerEnableStorageCardEncryptionOnMobile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerEnableStorageCardEncryptionOnMobile(val)
        }
        return nil
    }
    res["bitLockerEncryptDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerEncryptDevice(val)
        }
        return nil
    }
    res["bitLockerFixedDrivePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerFixedDrivePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerFixedDrivePolicy(val.(BitLockerFixedDrivePolicyable))
        }
        return nil
    }
    res["bitLockerRecoveryPasswordRotation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBitLockerRecoveryPasswordRotationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerRecoveryPasswordRotation(val.(*BitLockerRecoveryPasswordRotationType))
        }
        return nil
    }
    res["bitLockerRemovableDrivePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerRemovableDrivePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerRemovableDrivePolicy(val.(BitLockerRemovableDrivePolicyable))
        }
        return nil
    }
    res["bitLockerSystemDrivePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerSystemDrivePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerSystemDrivePolicy(val.(BitLockerSystemDrivePolicyable))
        }
        return nil
    }
    res["defenderAdditionalGuardedFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderAdditionalGuardedFolders(res)
        }
        return nil
    }
    res["defenderAdobeReaderLaunchChildProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAdobeReaderLaunchChildProcess(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderAdvancedRansomewareProtectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAdvancedRansomewareProtectionType(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderAllowBehaviorMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowBehaviorMonitoring(val)
        }
        return nil
    }
    res["defenderAllowCloudProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowCloudProtection(val)
        }
        return nil
    }
    res["defenderAllowEndUserAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowEndUserAccess(val)
        }
        return nil
    }
    res["defenderAllowIntrusionPreventionSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowIntrusionPreventionSystem(val)
        }
        return nil
    }
    res["defenderAllowOnAccessProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowOnAccessProtection(val)
        }
        return nil
    }
    res["defenderAllowRealTimeMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowRealTimeMonitoring(val)
        }
        return nil
    }
    res["defenderAllowScanArchiveFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanArchiveFiles(val)
        }
        return nil
    }
    res["defenderAllowScanDownloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanDownloads(val)
        }
        return nil
    }
    res["defenderAllowScanNetworkFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanNetworkFiles(val)
        }
        return nil
    }
    res["defenderAllowScanRemovableDrivesDuringFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanRemovableDrivesDuringFullScan(val)
        }
        return nil
    }
    res["defenderAllowScanScriptsLoadedInInternetExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAllowScanScriptsLoadedInInternetExplorer(val)
        }
        return nil
    }
    res["defenderAttackSurfaceReductionExcludedPaths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderAttackSurfaceReductionExcludedPaths(res)
        }
        return nil
    }
    res["defenderBlockEndUserAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderBlockEndUserAccess(val)
        }
        return nil
    }
    res["defenderBlockPersistenceThroughWmiType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderBlockPersistenceThroughWmiType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderCheckForSignaturesBeforeRunningScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderCheckForSignaturesBeforeRunningScan(val)
        }
        return nil
    }
    res["defenderCloudBlockLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderCloudBlockLevelType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderCloudBlockLevel(val.(*DefenderCloudBlockLevelType))
        }
        return nil
    }
    res["defenderCloudExtendedTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderCloudExtendedTimeoutInSeconds(val)
        }
        return nil
    }
    res["defenderDaysBeforeDeletingQuarantinedMalware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDaysBeforeDeletingQuarantinedMalware(val)
        }
        return nil
    }
    res["defenderDetectedMalwareActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDefenderDetectedMalwareActionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDetectedMalwareActions(val.(DefenderDetectedMalwareActionsable))
        }
        return nil
    }
    res["defenderDisableBehaviorMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableBehaviorMonitoring(val)
        }
        return nil
    }
    res["defenderDisableCatchupFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableCatchupFullScan(val)
        }
        return nil
    }
    res["defenderDisableCatchupQuickScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableCatchupQuickScan(val)
        }
        return nil
    }
    res["defenderDisableCloudProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableCloudProtection(val)
        }
        return nil
    }
    res["defenderDisableIntrusionPreventionSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableIntrusionPreventionSystem(val)
        }
        return nil
    }
    res["defenderDisableOnAccessProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableOnAccessProtection(val)
        }
        return nil
    }
    res["defenderDisableRealTimeMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableRealTimeMonitoring(val)
        }
        return nil
    }
    res["defenderDisableScanArchiveFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanArchiveFiles(val)
        }
        return nil
    }
    res["defenderDisableScanDownloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanDownloads(val)
        }
        return nil
    }
    res["defenderDisableScanNetworkFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanNetworkFiles(val)
        }
        return nil
    }
    res["defenderDisableScanRemovableDrivesDuringFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanRemovableDrivesDuringFullScan(val)
        }
        return nil
    }
    res["defenderDisableScanScriptsLoadedInInternetExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDisableScanScriptsLoadedInInternetExplorer(val)
        }
        return nil
    }
    res["defenderEmailContentExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEmailContentExecution(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderEmailContentExecutionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEmailContentExecutionType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderEnableLowCpuPriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEnableLowCpuPriority(val)
        }
        return nil
    }
    res["defenderEnableScanIncomingMail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEnableScanIncomingMail(val)
        }
        return nil
    }
    res["defenderEnableScanMappedNetworkDrivesDuringFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderEnableScanMappedNetworkDrivesDuringFullScan(val)
        }
        return nil
    }
    res["defenderExploitProtectionXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderExploitProtectionXml(val)
        }
        return nil
    }
    res["defenderExploitProtectionXmlFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderExploitProtectionXmlFileName(val)
        }
        return nil
    }
    res["defenderFileExtensionsToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderFileExtensionsToExclude(res)
        }
        return nil
    }
    res["defenderFilesAndFoldersToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderFilesAndFoldersToExclude(res)
        }
        return nil
    }
    res["defenderGuardedFoldersAllowedAppPaths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderGuardedFoldersAllowedAppPaths(res)
        }
        return nil
    }
    res["defenderGuardMyFoldersType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFolderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderGuardMyFoldersType(val.(*FolderProtectionType))
        }
        return nil
    }
    res["defenderNetworkProtectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderNetworkProtectionType(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeAppsExecutableContentCreationOrLaunch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsExecutableContentCreationOrLaunch(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeAppsExecutableContentCreationOrLaunchType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsExecutableContentCreationOrLaunchType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderOfficeAppsLaunchChildProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsLaunchChildProcess(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeAppsLaunchChildProcessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsLaunchChildProcessType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderOfficeAppsOtherProcessInjection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsOtherProcessInjection(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeAppsOtherProcessInjectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeAppsOtherProcessInjectionType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderOfficeCommunicationAppsLaunchChildProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeCommunicationAppsLaunchChildProcess(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeMacroCodeAllowWin32Imports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeMacroCodeAllowWin32Imports(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderOfficeMacroCodeAllowWin32ImportsType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderOfficeMacroCodeAllowWin32ImportsType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderPotentiallyUnwantedAppAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderPotentiallyUnwantedAppAction(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderPreventCredentialStealingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderPreventCredentialStealingType(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderProcessCreation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderProcessCreation(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderProcessCreationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderProcessCreationType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderProcessesToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderProcessesToExclude(res)
        }
        return nil
    }
    res["defenderScanDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderRealtimeScanDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanDirection(val.(*DefenderRealtimeScanDirection))
        }
        return nil
    }
    res["defenderScanMaxCpuPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanMaxCpuPercentage(val)
        }
        return nil
    }
    res["defenderScanType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderScanType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanType(val.(*DefenderScanType))
        }
        return nil
    }
    res["defenderScheduledQuickScanTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScheduledQuickScanTime(val)
        }
        return nil
    }
    res["defenderScheduledScanDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeeklySchedule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScheduledScanDay(val.(*WeeklySchedule))
        }
        return nil
    }
    res["defenderScheduledScanTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScheduledScanTime(val)
        }
        return nil
    }
    res["defenderScriptDownloadedPayloadExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScriptDownloadedPayloadExecution(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderScriptDownloadedPayloadExecutionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScriptDownloadedPayloadExecutionType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderScriptObfuscatedMacroCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScriptObfuscatedMacroCode(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderScriptObfuscatedMacroCodeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScriptObfuscatedMacroCodeType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderSecurityCenterBlockExploitProtectionOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterBlockExploitProtectionOverride(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableAccountUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableAccountUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableAppBrowserUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableAppBrowserUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableClearTpmUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableClearTpmUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableFamilyUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableFamilyUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableHardwareUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableHardwareUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableHealthUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableHealthUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableNetworkUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableNetworkUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableNotificationAreaUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableNotificationAreaUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableRansomwareUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableRansomwareUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableSecureBootUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableSecureBootUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableTroubleshootingUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableTroubleshootingUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableVirusUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableVirusUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI(val)
        }
        return nil
    }
    res["defenderSecurityCenterHelpEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterHelpEmail(val)
        }
        return nil
    }
    res["defenderSecurityCenterHelpPhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterHelpPhone(val)
        }
        return nil
    }
    res["defenderSecurityCenterHelpURL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterHelpURL(val)
        }
        return nil
    }
    res["defenderSecurityCenterITContactDisplay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderSecurityCenterITContactDisplayType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterITContactDisplay(val.(*DefenderSecurityCenterITContactDisplayType))
        }
        return nil
    }
    res["defenderSecurityCenterNotificationsFromApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderSecurityCenterNotificationsFromAppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterNotificationsFromApp(val.(*DefenderSecurityCenterNotificationsFromAppType))
        }
        return nil
    }
    res["defenderSecurityCenterOrganizationDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterOrganizationDisplayName(val)
        }
        return nil
    }
    res["defenderSignatureUpdateIntervalInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSignatureUpdateIntervalInHours(val)
        }
        return nil
    }
    res["defenderSubmitSamplesConsentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderSubmitSamplesConsentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSubmitSamplesConsentType(val.(*DefenderSubmitSamplesConsentType))
        }
        return nil
    }
    res["defenderUntrustedExecutable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderUntrustedExecutable(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderUntrustedExecutableType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderUntrustedExecutableType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["defenderUntrustedUSBProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderProtectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderUntrustedUSBProcess(val.(*DefenderProtectionType))
        }
        return nil
    }
    res["defenderUntrustedUSBProcessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAttackSurfaceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderUntrustedUSBProcessType(val.(*DefenderAttackSurfaceType))
        }
        return nil
    }
    res["deviceGuardEnableSecureBootWithDMA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardEnableSecureBootWithDMA(val)
        }
        return nil
    }
    res["deviceGuardEnableVirtualizationBasedSecurity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardEnableVirtualizationBasedSecurity(val)
        }
        return nil
    }
    res["deviceGuardLaunchSystemGuard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardLaunchSystemGuard(val.(*Enablement))
        }
        return nil
    }
    res["deviceGuardLocalSystemAuthorityCredentialGuardSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardLocalSystemAuthorityCredentialGuardType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings(val.(*DeviceGuardLocalSystemAuthorityCredentialGuardType))
        }
        return nil
    }
    res["deviceGuardSecureBootWithDMA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecureBootWithDMAType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardSecureBootWithDMA(val.(*SecureBootWithDMAType))
        }
        return nil
    }
    res["dmaGuardDeviceEnumerationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDmaGuardDeviceEnumerationPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDmaGuardDeviceEnumerationPolicy(val.(*DmaGuardDeviceEnumerationPolicyType))
        }
        return nil
    }
    res["firewallBlockStatefulFTP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallBlockStatefulFTP(val)
        }
        return nil
    }
    res["firewallCertificateRevocationListCheckMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallCertificateRevocationListCheckMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallCertificateRevocationListCheckMethod(val.(*FirewallCertificateRevocationListCheckMethodType))
        }
        return nil
    }
    res["firewallIdleTimeoutForSecurityAssociationInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIdleTimeoutForSecurityAssociationInSeconds(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowDHCP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowDHCP(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowICMP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowICMP(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowNeighborDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowNeighborDiscovery(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowRouterDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowRouterDiscovery(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsNone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsNone(val)
        }
        return nil
    }
    res["firewallMergeKeyingModuleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallMergeKeyingModuleSettings(val)
        }
        return nil
    }
    res["firewallPacketQueueingMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallPacketQueueingMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallPacketQueueingMethod(val.(*FirewallPacketQueueingMethodType))
        }
        return nil
    }
    res["firewallPreSharedKeyEncodingMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallPreSharedKeyEncodingMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallPreSharedKeyEncodingMethod(val.(*FirewallPreSharedKeyEncodingMethodType))
        }
        return nil
    }
    res["firewallProfileDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfileDomain(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["firewallProfilePrivate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfilePrivate(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["firewallProfilePublic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfilePublic(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["firewallRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsFirewallRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsFirewallRuleable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsFirewallRuleable)
            }
            m.SetFirewallRules(res)
        }
        return nil
    }
    res["lanManagerAuthenticationLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanManagerAuthenticationLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanManagerAuthenticationLevel(val.(*LanManagerAuthenticationLevel))
        }
        return nil
    }
    res["lanManagerWorkstationDisableInsecureGuestLogons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanManagerWorkstationDisableInsecureGuestLogons(val)
        }
        return nil
    }
    res["localSecurityOptionsAdministratorAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAdministratorAccountName(val)
        }
        return nil
    }
    res["localSecurityOptionsAdministratorElevationPromptBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsAdministratorElevationPromptBehaviorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAdministratorElevationPromptBehavior(val.(*LocalSecurityOptionsAdministratorElevationPromptBehaviorType))
        }
        return nil
    }
    res["localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowPKU2UAuthenticationRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowUIAccessApplicationElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowUIAccessApplicationElevation(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowUIAccessApplicationsForSecureLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations(val)
        }
        return nil
    }
    res["localSecurityOptionsAllowUndockWithoutHavingToLogon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon(val)
        }
        return nil
    }
    res["localSecurityOptionsBlockMicrosoftAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsBlockMicrosoftAccounts(val)
        }
        return nil
    }
    res["localSecurityOptionsBlockRemoteLogonWithBlankPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword(val)
        }
        return nil
    }
    res["localSecurityOptionsBlockRemoteOpticalDriveAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess(val)
        }
        return nil
    }
    res["localSecurityOptionsBlockUsersInstallingPrinterDrivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers(val)
        }
        return nil
    }
    res["localSecurityOptionsClearVirtualMemoryPageFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsClearVirtualMemoryPageFile(val)
        }
        return nil
    }
    res["localSecurityOptionsClientDigitallySignCommunicationsAlways"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways(val)
        }
        return nil
    }
    res["localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers(val)
        }
        return nil
    }
    res["localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableAdministratorAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableAdministratorAccount(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableGuestAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableGuestAccount(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableServerDigitallySignCommunicationsAlways"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways(val)
        }
        return nil
    }
    res["localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees(val)
        }
        return nil
    }
    res["localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts(val)
        }
        return nil
    }
    res["localSecurityOptionsDoNotRequireCtrlAltDel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDoNotRequireCtrlAltDel(val)
        }
        return nil
    }
    res["localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange(val)
        }
        return nil
    }
    res["localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(val.(*LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType))
        }
        return nil
    }
    res["localSecurityOptionsGuestAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsGuestAccountName(val)
        }
        return nil
    }
    res["localSecurityOptionsHideLastSignedInUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsHideLastSignedInUser(val)
        }
        return nil
    }
    res["localSecurityOptionsHideUsernameAtSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsHideUsernameAtSignIn(val)
        }
        return nil
    }
    res["localSecurityOptionsInformationDisplayedOnLockScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsInformationDisplayedOnLockScreenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsInformationDisplayedOnLockScreen(val.(*LocalSecurityOptionsInformationDisplayedOnLockScreenType))
        }
        return nil
    }
    res["localSecurityOptionsInformationShownOnLockScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsInformationShownOnLockScreenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsInformationShownOnLockScreen(val.(*LocalSecurityOptionsInformationShownOnLockScreenType))
        }
        return nil
    }
    res["localSecurityOptionsLogOnMessageText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsLogOnMessageText(val)
        }
        return nil
    }
    res["localSecurityOptionsLogOnMessageTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsLogOnMessageTitle(val)
        }
        return nil
    }
    res["localSecurityOptionsMachineInactivityLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsMachineInactivityLimit(val)
        }
        return nil
    }
    res["localSecurityOptionsMachineInactivityLimitInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsMachineInactivityLimitInMinutes(val)
        }
        return nil
    }
    res["localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsMinimumSessionSecurity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(val.(*LocalSecurityOptionsMinimumSessionSecurity))
        }
        return nil
    }
    res["localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsMinimumSessionSecurity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(val.(*LocalSecurityOptionsMinimumSessionSecurity))
        }
        return nil
    }
    res["localSecurityOptionsOnlyElevateSignedExecutables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsOnlyElevateSignedExecutables(val)
        }
        return nil
    }
    res["localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares(val)
        }
        return nil
    }
    res["localSecurityOptionsSmartCardRemovalBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsSmartCardRemovalBehaviorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsSmartCardRemovalBehavior(val.(*LocalSecurityOptionsSmartCardRemovalBehaviorType))
        }
        return nil
    }
    res["localSecurityOptionsStandardUserElevationPromptBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocalSecurityOptionsStandardUserElevationPromptBehaviorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsStandardUserElevationPromptBehavior(val.(*LocalSecurityOptionsStandardUserElevationPromptBehaviorType))
        }
        return nil
    }
    res["localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation(val)
        }
        return nil
    }
    res["localSecurityOptionsUseAdminApprovalMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsUseAdminApprovalMode(val)
        }
        return nil
    }
    res["localSecurityOptionsUseAdminApprovalModeForAdministrators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators(val)
        }
        return nil
    }
    res["localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations(val)
        }
        return nil
    }
    res["smartScreenBlockOverrideForFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenBlockOverrideForFiles(val)
        }
        return nil
    }
    res["smartScreenEnableInShell"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenEnableInShell(val)
        }
        return nil
    }
    res["userRightsAccessCredentialManagerAsTrustedCaller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsAccessCredentialManagerAsTrustedCaller(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsActAsPartOfTheOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsActAsPartOfTheOperatingSystem(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsAllowAccessFromNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsAllowAccessFromNetwork(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsBackupData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsBackupData(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsBlockAccessFromNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsBlockAccessFromNetwork(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsChangeSystemTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsChangeSystemTime(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreateGlobalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreateGlobalObjects(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreatePageFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreatePageFile(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreatePermanentSharedObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreatePermanentSharedObjects(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreateSymbolicLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreateSymbolicLinks(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsCreateToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsCreateToken(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsDebugPrograms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsDebugPrograms(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsDelegation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsDelegation(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsDenyLocalLogOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsDenyLocalLogOn(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsGenerateSecurityAudits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsGenerateSecurityAudits(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsImpersonateClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsImpersonateClient(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsIncreaseSchedulingPriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsIncreaseSchedulingPriority(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsLoadUnloadDrivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsLoadUnloadDrivers(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsLocalLogOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsLocalLogOn(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsLockMemory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsLockMemory(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsManageAuditingAndSecurityLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsManageAuditingAndSecurityLogs(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsManageVolumes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsManageVolumes(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsModifyFirmwareEnvironment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsModifyFirmwareEnvironment(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsModifyObjectLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsModifyObjectLabels(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsProfileSingleProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsProfileSingleProcess(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsRemoteDesktopServicesLogOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsRemoteDesktopServicesLogOn(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsRemoteShutdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsRemoteShutdown(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsRestoreData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsRestoreData(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["userRightsTakeOwnership"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementUserRightsSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRightsTakeOwnership(val.(DeviceManagementUserRightsSettingable))
        }
        return nil
    }
    res["windowsDefenderTamperProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderTamperProtectionOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsDefenderTamperProtection(val.(*WindowsDefenderTamperProtectionOptions))
        }
        return nil
    }
    res["xboxServicesAccessoryManagementServiceStartupMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceStartType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesAccessoryManagementServiceStartupMode(val.(*ServiceStartType))
        }
        return nil
    }
    res["xboxServicesEnableXboxGameSaveTask"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesEnableXboxGameSaveTask(val)
        }
        return nil
    }
    res["xboxServicesLiveAuthManagerServiceStartupMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceStartType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesLiveAuthManagerServiceStartupMode(val.(*ServiceStartType))
        }
        return nil
    }
    res["xboxServicesLiveGameSaveServiceStartupMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceStartType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesLiveGameSaveServiceStartupMode(val.(*ServiceStartType))
        }
        return nil
    }
    res["xboxServicesLiveNetworkingServiceStartupMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceStartType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetXboxServicesLiveNetworkingServiceStartupMode(val.(*ServiceStartType))
        }
        return nil
    }
    return res
}
// GetFirewallBlockStatefulFTP gets the firewallBlockStatefulFTP property value. Blocks stateful FTP connections to the device
func (m *Windows10EndpointProtectionConfiguration) GetFirewallBlockStatefulFTP()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallBlockStatefulFTP
    }
}
// GetFirewallCertificateRevocationListCheckMethod gets the firewallCertificateRevocationListCheckMethod property value. Specify how the certificate revocation list is to be enforced. Possible values are: deviceDefault, none, attempt, require.
func (m *Windows10EndpointProtectionConfiguration) GetFirewallCertificateRevocationListCheckMethod()(*FirewallCertificateRevocationListCheckMethodType) {
    if m == nil {
        return nil
    } else {
        return m.firewallCertificateRevocationListCheckMethod
    }
}
// GetFirewallIdleTimeoutForSecurityAssociationInSeconds gets the firewallIdleTimeoutForSecurityAssociationInSeconds property value. Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIdleTimeoutForSecurityAssociationInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.firewallIdleTimeoutForSecurityAssociationInSeconds
    }
}
// GetFirewallIPSecExemptionsAllowDHCP gets the firewallIPSecExemptionsAllowDHCP property value. Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowDHCP()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallIPSecExemptionsAllowDHCP
    }
}
// GetFirewallIPSecExemptionsAllowICMP gets the firewallIPSecExemptionsAllowICMP property value. Configures IPSec exemptions to allow ICMP
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowICMP()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallIPSecExemptionsAllowICMP
    }
}
// GetFirewallIPSecExemptionsAllowNeighborDiscovery gets the firewallIPSecExemptionsAllowNeighborDiscovery property value. Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowNeighborDiscovery()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallIPSecExemptionsAllowNeighborDiscovery
    }
}
// GetFirewallIPSecExemptionsAllowRouterDiscovery gets the firewallIPSecExemptionsAllowRouterDiscovery property value. Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowRouterDiscovery()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallIPSecExemptionsAllowRouterDiscovery
    }
}
// GetFirewallIPSecExemptionsNone gets the firewallIPSecExemptionsNone property value. Configures IPSec exemptions to no exemptions
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsNone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallIPSecExemptionsNone
    }
}
// GetFirewallMergeKeyingModuleSettings gets the firewallMergeKeyingModuleSettings property value. If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
func (m *Windows10EndpointProtectionConfiguration) GetFirewallMergeKeyingModuleSettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.firewallMergeKeyingModuleSettings
    }
}
// GetFirewallPacketQueueingMethod gets the firewallPacketQueueingMethod property value. Configures how packet queueing should be applied in the tunnel gateway scenario. Possible values are: deviceDefault, disabled, queueInbound, queueOutbound, queueBoth.
func (m *Windows10EndpointProtectionConfiguration) GetFirewallPacketQueueingMethod()(*FirewallPacketQueueingMethodType) {
    if m == nil {
        return nil
    } else {
        return m.firewallPacketQueueingMethod
    }
}
// GetFirewallPreSharedKeyEncodingMethod gets the firewallPreSharedKeyEncodingMethod property value. Select the preshared key encoding to be used. Possible values are: deviceDefault, none, utF8.
func (m *Windows10EndpointProtectionConfiguration) GetFirewallPreSharedKeyEncodingMethod()(*FirewallPreSharedKeyEncodingMethodType) {
    if m == nil {
        return nil
    } else {
        return m.firewallPreSharedKeyEncodingMethod
    }
}
// GetFirewallProfileDomain gets the firewallProfileDomain property value. Configures the firewall profile settings for domain networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfileDomain()(WindowsFirewallNetworkProfileable) {
    if m == nil {
        return nil
    } else {
        return m.firewallProfileDomain
    }
}
// GetFirewallProfilePrivate gets the firewallProfilePrivate property value. Configures the firewall profile settings for private networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfilePrivate()(WindowsFirewallNetworkProfileable) {
    if m == nil {
        return nil
    } else {
        return m.firewallProfilePrivate
    }
}
// GetFirewallProfilePublic gets the firewallProfilePublic property value. Configures the firewall profile settings for public networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfilePublic()(WindowsFirewallNetworkProfileable) {
    if m == nil {
        return nil
    } else {
        return m.firewallProfilePublic
    }
}
// GetFirewallRules gets the firewallRules property value. Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
func (m *Windows10EndpointProtectionConfiguration) GetFirewallRules()([]WindowsFirewallRuleable) {
    if m == nil {
        return nil
    } else {
        return m.firewallRules
    }
}
// GetLanManagerAuthenticationLevel gets the lanManagerAuthenticationLevel property value. This security setting determines which challenge/response authentication protocol is used for network logons. Possible values are: lmAndNltm, lmNtlmAndNtlmV2, lmAndNtlmOnly, lmAndNtlmV2, lmNtlmV2AndNotLm, lmNtlmV2AndNotLmOrNtm.
func (m *Windows10EndpointProtectionConfiguration) GetLanManagerAuthenticationLevel()(*LanManagerAuthenticationLevel) {
    if m == nil {
        return nil
    } else {
        return m.lanManagerAuthenticationLevel
    }
}
// GetLanManagerWorkstationDisableInsecureGuestLogons gets the lanManagerWorkstationDisableInsecureGuestLogons property value. If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure guest logons.
func (m *Windows10EndpointProtectionConfiguration) GetLanManagerWorkstationDisableInsecureGuestLogons()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.lanManagerWorkstationDisableInsecureGuestLogons
    }
}
// GetLocalSecurityOptionsAdministratorAccountName gets the localSecurityOptionsAdministratorAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAdministratorAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAdministratorAccountName
    }
}
// GetLocalSecurityOptionsAdministratorElevationPromptBehavior gets the localSecurityOptionsAdministratorElevationPromptBehavior property value. Define the behavior of the elevation prompt for admins in Admin Approval Mode. Possible values are: notConfigured, elevateWithoutPrompting, promptForCredentialsOnTheSecureDesktop, promptForConsentOnTheSecureDesktop, promptForCredentials, promptForConsent, promptForConsentForNonWindowsBinaries.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAdministratorElevationPromptBehavior()(*LocalSecurityOptionsAdministratorElevationPromptBehaviorType) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAdministratorElevationPromptBehavior
    }
}
// GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares gets the localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares property value. This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating the names of domain accounts and network shares.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares
    }
}
// GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests gets the localSecurityOptionsAllowPKU2UAuthenticationRequests property value. Block PKU2U authentication requests to this device to use online identities.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAllowPKU2UAuthenticationRequests
    }
}
// GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager gets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager property value. Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote calls to the SAM.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager
    }
}
// GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool gets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool property value. UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool
    }
}
// GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn gets the localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn property value. This security setting determines whether a computer can be shut down without having to log on to Windows.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn
    }
}
// GetLocalSecurityOptionsAllowUIAccessApplicationElevation gets the localSecurityOptionsAllowUIAccessApplicationElevation property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUIAccessApplicationElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAllowUIAccessApplicationElevation
    }
}
// GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations gets the localSecurityOptionsAllowUIAccessApplicationsForSecureLocations property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAllowUIAccessApplicationsForSecureLocations
    }
}
// GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon gets the localSecurityOptionsAllowUndockWithoutHavingToLogon property value. Prevent a portable computer from being undocked without having to log in.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsAllowUndockWithoutHavingToLogon
    }
}
// GetLocalSecurityOptionsBlockMicrosoftAccounts gets the localSecurityOptionsBlockMicrosoftAccounts property value. Prevent users from adding new Microsoft accounts to this computer.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockMicrosoftAccounts()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsBlockMicrosoftAccounts
    }
}
// GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword gets the localSecurityOptionsBlockRemoteLogonWithBlankPassword property value. Enable Local accounts that are not password protected to log on from locations other than the physical device.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsBlockRemoteLogonWithBlankPassword
    }
}
// GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess gets the localSecurityOptionsBlockRemoteOpticalDriveAccess property value. Enabling this settings allows only interactively logged on user to access CD-ROM media.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsBlockRemoteOpticalDriveAccess
    }
}
// GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers gets the localSecurityOptionsBlockUsersInstallingPrinterDrivers property value. Restrict installing printer drivers as part of connecting to a shared printer to admins only.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsBlockUsersInstallingPrinterDrivers
    }
}
// GetLocalSecurityOptionsClearVirtualMemoryPageFile gets the localSecurityOptionsClearVirtualMemoryPageFile property value. This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClearVirtualMemoryPageFile()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsClearVirtualMemoryPageFile
    }
}
// GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways gets the localSecurityOptionsClientDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB client component.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsClientDigitallySignCommunicationsAlways
    }
}
// GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers gets the localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers property value. If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords to non-Microsoft SMB servers that do not support password encryption during authentication.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers
    }
}
// GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation gets the localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation property value. App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation
    }
}
// GetLocalSecurityOptionsDisableAdministratorAccount gets the localSecurityOptionsDisableAdministratorAccount property value. Determines whether the Local Administrator account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableAdministratorAccount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDisableAdministratorAccount
    }
}
// GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees gets the localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees property value. This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees
    }
}
// GetLocalSecurityOptionsDisableGuestAccount gets the localSecurityOptionsDisableGuestAccount property value. Determines if the Guest account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableGuestAccount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDisableGuestAccount
    }
}
// GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways gets the localSecurityOptionsDisableServerDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB server component.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDisableServerDigitallySignCommunicationsAlways
    }
}
// GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees gets the localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees property value. This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request it.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees
    }
}
// GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts gets the localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts property value. This security setting determines what additional permissions will be granted for anonymous connections to the computer.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts
    }
}
// GetLocalSecurityOptionsDoNotRequireCtrlAltDel gets the localSecurityOptionsDoNotRequireCtrlAltDel property value. Require CTRL+ALT+DEL to be pressed before a user can log on.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotRequireCtrlAltDel()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDoNotRequireCtrlAltDel
    }
}
// GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange gets the localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange property value. This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new password is stored. It’s not stored by default.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange
    }
}
// GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser gets the localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser property value. Define who is allowed to format and eject removable NTFS media. Possible values are: notConfigured, administrators, administratorsAndPowerUsers, administratorsAndInteractiveUsers.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser()(*LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser
    }
}
// GetLocalSecurityOptionsGuestAccountName gets the localSecurityOptionsGuestAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsGuestAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsGuestAccountName
    }
}
// GetLocalSecurityOptionsHideLastSignedInUser gets the localSecurityOptionsHideLastSignedInUser property value. Do not display the username of the last person who signed in on this device.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsHideLastSignedInUser()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsHideLastSignedInUser
    }
}
// GetLocalSecurityOptionsHideUsernameAtSignIn gets the localSecurityOptionsHideUsernameAtSignIn property value. Do not display the username of the person signing in to this device after credentials are entered and before the device’s desktop is shown.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsHideUsernameAtSignIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsHideUsernameAtSignIn
    }
}
// GetLocalSecurityOptionsInformationDisplayedOnLockScreen gets the localSecurityOptionsInformationDisplayedOnLockScreen property value. Configure the user information that is displayed when the session is locked. If not configured, user display name, domain and username are shown. Possible values are: notConfigured, administrators, administratorsAndPowerUsers, administratorsAndInteractiveUsers.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsInformationDisplayedOnLockScreen()(*LocalSecurityOptionsInformationDisplayedOnLockScreenType) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsInformationDisplayedOnLockScreen
    }
}
// GetLocalSecurityOptionsInformationShownOnLockScreen gets the localSecurityOptionsInformationShownOnLockScreen property value. Configure the user information that is displayed when the session is locked. If not configured, user display name, domain and username are shown. Possible values are: notConfigured, userDisplayNameDomainUser, userDisplayNameOnly, doNotDisplayUser.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsInformationShownOnLockScreen()(*LocalSecurityOptionsInformationShownOnLockScreenType) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsInformationShownOnLockScreen
    }
}
// GetLocalSecurityOptionsLogOnMessageText gets the localSecurityOptionsLogOnMessageText property value. Set message text for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsLogOnMessageText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsLogOnMessageText
    }
}
// GetLocalSecurityOptionsLogOnMessageTitle gets the localSecurityOptionsLogOnMessageTitle property value. Set message title for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsLogOnMessageTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsLogOnMessageTitle
    }
}
// GetLocalSecurityOptionsMachineInactivityLimit gets the localSecurityOptionsMachineInactivityLimit property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMachineInactivityLimit()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsMachineInactivityLimit
    }
}
// GetLocalSecurityOptionsMachineInactivityLimitInMinutes gets the localSecurityOptionsMachineInactivityLimitInMinutes property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMachineInactivityLimitInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsMachineInactivityLimitInMinutes
    }
}
// GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients gets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients property value. This security setting allows a client to require the negotiation of 128-bit encryption and/or NTLMv2 session security. Possible values are: none, requireNtmlV2SessionSecurity, require128BitEncryption, ntlmV2And128BitEncryption.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients()(*LocalSecurityOptionsMinimumSessionSecurity) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients
    }
}
// GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers gets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers property value. This security setting allows a server to require the negotiation of 128-bit encryption and/or NTLMv2 session security. Possible values are: none, requireNtmlV2SessionSecurity, require128BitEncryption, ntlmV2And128BitEncryption.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers()(*LocalSecurityOptionsMinimumSessionSecurity) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers
    }
}
// GetLocalSecurityOptionsOnlyElevateSignedExecutables gets the localSecurityOptionsOnlyElevateSignedExecutables property value. Enforce PKI certification path validation for a given executable file before it is permitted to run.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsOnlyElevateSignedExecutables()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsOnlyElevateSignedExecutables
    }
}
// GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares gets the localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares property value. By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that can be accessed anonymously and Shares that can be accessed anonymously
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares
    }
}
// GetLocalSecurityOptionsSmartCardRemovalBehavior gets the localSecurityOptionsSmartCardRemovalBehavior property value. This security setting determines what happens when the smart card for a logged-on user is removed from the smart card reader. Possible values are: noAction, lockWorkstation, forceLogoff, disconnectRemoteDesktopSession.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsSmartCardRemovalBehavior()(*LocalSecurityOptionsSmartCardRemovalBehaviorType) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsSmartCardRemovalBehavior
    }
}
// GetLocalSecurityOptionsStandardUserElevationPromptBehavior gets the localSecurityOptionsStandardUserElevationPromptBehavior property value. Define the behavior of the elevation prompt for standard users. Possible values are: notConfigured, automaticallyDenyElevationRequests, promptForCredentialsOnTheSecureDesktop, promptForCredentials.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsStandardUserElevationPromptBehavior()(*LocalSecurityOptionsStandardUserElevationPromptBehaviorType) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsStandardUserElevationPromptBehavior
    }
}
// GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation gets the localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation property value. Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior policy settings for admins and standard users are used.
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation
    }
}
// GetLocalSecurityOptionsUseAdminApprovalMode gets the localSecurityOptionsUseAdminApprovalMode property value. Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin privileges.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsUseAdminApprovalMode()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsUseAdminApprovalMode
    }
}
// GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators gets the localSecurityOptionsUseAdminApprovalModeForAdministrators property value. Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsUseAdminApprovalModeForAdministrators
    }
}
// GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations gets the localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations property value. Virtualize file and registry write failures to per user locations
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations
    }
}
// GetSmartScreenBlockOverrideForFiles gets the smartScreenBlockOverrideForFiles property value. Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
func (m *Windows10EndpointProtectionConfiguration) GetSmartScreenBlockOverrideForFiles()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.smartScreenBlockOverrideForFiles
    }
}
// GetSmartScreenEnableInShell gets the smartScreenEnableInShell property value. Allows IT Admins to configure SmartScreen for Windows.
func (m *Windows10EndpointProtectionConfiguration) GetSmartScreenEnableInShell()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.smartScreenEnableInShell
    }
}
// GetUserRightsAccessCredentialManagerAsTrustedCaller gets the userRightsAccessCredentialManagerAsTrustedCaller property value. This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if this privilege is given to other entities. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsAccessCredentialManagerAsTrustedCaller()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsAccessCredentialManagerAsTrustedCaller
    }
}
// GetUserRightsActAsPartOfTheOperatingSystem gets the userRightsActAsPartOfTheOperatingSystem property value. This user right allows a process to impersonate any user without authentication. The process can therefore gain access to the same local resources as that user. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsActAsPartOfTheOperatingSystem()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsActAsPartOfTheOperatingSystem
    }
}
// GetUserRightsAllowAccessFromNetwork gets the userRightsAllowAccessFromNetwork property value. This user right determines which users and groups are allowed to connect to the computer over the network. State Allowed is supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsAllowAccessFromNetwork()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsAllowAccessFromNetwork
    }
}
// GetUserRightsBackupData gets the userRightsBackupData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when backing up files and directories. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsBackupData()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsBackupData
    }
}
// GetUserRightsBlockAccessFromNetwork gets the userRightsBlockAccessFromNetwork property value. This user right determines which users and groups are block from connecting to the computer over the network. State Block is supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsBlockAccessFromNetwork()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsBlockAccessFromNetwork
    }
}
// GetUserRightsChangeSystemTime gets the userRightsChangeSystemTime property value. This user right determines which users and groups can change the time and date on the internal clock of the computer. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsChangeSystemTime()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsChangeSystemTime
    }
}
// GetUserRightsCreateGlobalObjects gets the userRightsCreateGlobalObjects property value. This security setting determines whether users can create global objects that are available to all sessions. Users who can create global objects could affect processes that run under other users' sessions, which could lead to application failure or data corruption. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateGlobalObjects()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsCreateGlobalObjects
    }
}
// GetUserRightsCreatePageFile gets the userRightsCreatePageFile property value. This user right determines which users and groups can call an internal API to create and change the size of a page file. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreatePageFile()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsCreatePageFile
    }
}
// GetUserRightsCreatePermanentSharedObjects gets the userRightsCreatePermanentSharedObjects property value. This user right determines which accounts can be used by processes to create a directory object using the object manager. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreatePermanentSharedObjects()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsCreatePermanentSharedObjects
    }
}
// GetUserRightsCreateSymbolicLinks gets the userRightsCreateSymbolicLinks property value. This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateSymbolicLinks()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsCreateSymbolicLinks
    }
}
// GetUserRightsCreateToken gets the userRightsCreateToken property value. This user right determines which users/groups can be used by processes to create a token that can then be used to get access to any local resources when the process uses an internal API to create an access token. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateToken()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsCreateToken
    }
}
// GetUserRightsDebugPrograms gets the userRightsDebugPrograms property value. This user right determines which users can attach a debugger to any process or to the kernel. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDebugPrograms()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsDebugPrograms
    }
}
// GetUserRightsDelegation gets the userRightsDelegation property value. This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDelegation()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsDelegation
    }
}
// GetUserRightsDenyLocalLogOn gets the userRightsDenyLocalLogOn property value. This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDenyLocalLogOn()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsDenyLocalLogOn
    }
}
// GetUserRightsGenerateSecurityAudits gets the userRightsGenerateSecurityAudits property value. This user right determines which accounts can be used by a process to add entries to the security log. The security log is used to trace unauthorized system access.  Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsGenerateSecurityAudits()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsGenerateSecurityAudits
    }
}
// GetUserRightsImpersonateClient gets the userRightsImpersonateClient property value. Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a service that they have created and then impersonating that client, which can elevate the unauthorized user's permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsImpersonateClient()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsImpersonateClient
    }
}
// GetUserRightsIncreaseSchedulingPriority gets the userRightsIncreaseSchedulingPriority property value. This user right determines which accounts can use a process with Write Property access to another process to increase the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsIncreaseSchedulingPriority()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsIncreaseSchedulingPriority
    }
}
// GetUserRightsLoadUnloadDrivers gets the userRightsLoadUnloadDrivers property value. This user right determines which users can dynamically load and unload device drivers or other code in to kernel mode. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLoadUnloadDrivers()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsLoadUnloadDrivers
    }
}
// GetUserRightsLocalLogOn gets the userRightsLocalLogOn property value. This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLocalLogOn()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsLocalLogOn
    }
}
// GetUserRightsLockMemory gets the userRightsLockMemory property value. This user right determines which accounts can use a process to keep data in physical memory, which prevents the system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLockMemory()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsLockMemory
    }
}
// GetUserRightsManageAuditingAndSecurityLogs gets the userRightsManageAuditingAndSecurityLogs property value. This user right determines which users can specify object access auditing options for individual resources, such as files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsManageAuditingAndSecurityLogs()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsManageAuditingAndSecurityLogs
    }
}
// GetUserRightsManageVolumes gets the userRightsManageVolumes property value. This user right determines which users and groups can run maintenance tasks on a volume, such as remote defragmentation. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsManageVolumes()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsManageVolumes
    }
}
// GetUserRightsModifyFirmwareEnvironment gets the userRightsModifyFirmwareEnvironment property value. This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsModifyFirmwareEnvironment()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsModifyFirmwareEnvironment
    }
}
// GetUserRightsModifyObjectLabels gets the userRightsModifyObjectLabels property value. This user right determines which user accounts can modify the integrity label of objects, such as files, registry keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsModifyObjectLabels()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsModifyObjectLabels
    }
}
// GetUserRightsProfileSingleProcess gets the userRightsProfileSingleProcess property value. This user right determines which users can use performance monitoring tools to monitor the performance of system processes. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsProfileSingleProcess()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsProfileSingleProcess
    }
}
// GetUserRightsRemoteDesktopServicesLogOn gets the userRightsRemoteDesktopServicesLogOn property value. This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client. Only states NotConfigured and Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRemoteDesktopServicesLogOn()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsRemoteDesktopServicesLogOn
    }
}
// GetUserRightsRemoteShutdown gets the userRightsRemoteShutdown property value. This user right determines which users are allowed to shut down a computer from a remote location on the network. Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRemoteShutdown()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsRemoteShutdown
    }
}
// GetUserRightsRestoreData gets the userRightsRestoreData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when restoring backed up files and directories, and determines which users can set any valid security principal as the owner of an object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRestoreData()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsRestoreData
    }
}
// GetUserRightsTakeOwnership gets the userRightsTakeOwnership property value. This user right determines which users can take ownership of any securable object in the system, including Active Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsTakeOwnership()(DeviceManagementUserRightsSettingable) {
    if m == nil {
        return nil
    } else {
        return m.userRightsTakeOwnership
    }
}
// GetWindowsDefenderTamperProtection gets the windowsDefenderTamperProtection property value. Configure windows defender TamperProtection settings. Possible values are: notConfigured, enable, disable.
func (m *Windows10EndpointProtectionConfiguration) GetWindowsDefenderTamperProtection()(*WindowsDefenderTamperProtectionOptions) {
    if m == nil {
        return nil
    } else {
        return m.windowsDefenderTamperProtection
    }
}
// GetXboxServicesAccessoryManagementServiceStartupMode gets the xboxServicesAccessoryManagementServiceStartupMode property value. This setting determines whether the Accessory management service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesAccessoryManagementServiceStartupMode()(*ServiceStartType) {
    if m == nil {
        return nil
    } else {
        return m.xboxServicesAccessoryManagementServiceStartupMode
    }
}
// GetXboxServicesEnableXboxGameSaveTask gets the xboxServicesEnableXboxGameSaveTask property value. This setting determines whether xbox game save is enabled (1) or disabled (0).
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesEnableXboxGameSaveTask()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.xboxServicesEnableXboxGameSaveTask
    }
}
// GetXboxServicesLiveAuthManagerServiceStartupMode gets the xboxServicesLiveAuthManagerServiceStartupMode property value. This setting determines whether Live Auth Manager service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveAuthManagerServiceStartupMode()(*ServiceStartType) {
    if m == nil {
        return nil
    } else {
        return m.xboxServicesLiveAuthManagerServiceStartupMode
    }
}
// GetXboxServicesLiveGameSaveServiceStartupMode gets the xboxServicesLiveGameSaveServiceStartupMode property value. This setting determines whether Live Game save service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveGameSaveServiceStartupMode()(*ServiceStartType) {
    if m == nil {
        return nil
    } else {
        return m.xboxServicesLiveGameSaveServiceStartupMode
    }
}
// GetXboxServicesLiveNetworkingServiceStartupMode gets the xboxServicesLiveNetworkingServiceStartupMode property value. This setting determines whether Networking service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveNetworkingServiceStartupMode()(*ServiceStartType) {
    if m == nil {
        return nil
    } else {
        return m.xboxServicesLiveNetworkingServiceStartupMode
    }
}
// Serialize serializes information the current object
func (m *Windows10EndpointProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowCameraMicrophoneRedirection", m.GetApplicationGuardAllowCameraMicrophoneRedirection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowFileSaveOnHost", m.GetApplicationGuardAllowFileSaveOnHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPersistence", m.GetApplicationGuardAllowPersistence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToLocalPrinters", m.GetApplicationGuardAllowPrintToLocalPrinters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToNetworkPrinters", m.GetApplicationGuardAllowPrintToNetworkPrinters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToPDF", m.GetApplicationGuardAllowPrintToPDF())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToXPS", m.GetApplicationGuardAllowPrintToXPS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowVirtualGPU", m.GetApplicationGuardAllowVirtualGPU())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardBlockClipboardSharing() != nil {
        cast := (*m.GetApplicationGuardBlockClipboardSharing()).String()
        err = writer.WriteStringValue("applicationGuardBlockClipboardSharing", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardBlockFileTransfer() != nil {
        cast := (*m.GetApplicationGuardBlockFileTransfer()).String()
        err = writer.WriteStringValue("applicationGuardBlockFileTransfer", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardBlockNonEnterpriseContent", m.GetApplicationGuardBlockNonEnterpriseContent())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardCertificateThumbprints() != nil {
        err = writer.WriteCollectionOfStringValues("applicationGuardCertificateThumbprints", m.GetApplicationGuardCertificateThumbprints())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardEnabled", m.GetApplicationGuardEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardEnabledOptions() != nil {
        cast := (*m.GetApplicationGuardEnabledOptions()).String()
        err = writer.WriteStringValue("applicationGuardEnabledOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardForceAuditing", m.GetApplicationGuardForceAuditing())
        if err != nil {
            return err
        }
    }
    if m.GetAppLockerApplicationControl() != nil {
        cast := (*m.GetAppLockerApplicationControl()).String()
        err = writer.WriteStringValue("appLockerApplicationControl", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerAllowStandardUserEncryption", m.GetBitLockerAllowStandardUserEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerDisableWarningForOtherDiskEncryption", m.GetBitLockerDisableWarningForOtherDiskEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerEnableStorageCardEncryptionOnMobile", m.GetBitLockerEnableStorageCardEncryptionOnMobile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerEncryptDevice", m.GetBitLockerEncryptDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bitLockerFixedDrivePolicy", m.GetBitLockerFixedDrivePolicy())
        if err != nil {
            return err
        }
    }
    if m.GetBitLockerRecoveryPasswordRotation() != nil {
        cast := (*m.GetBitLockerRecoveryPasswordRotation()).String()
        err = writer.WriteStringValue("bitLockerRecoveryPasswordRotation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bitLockerRemovableDrivePolicy", m.GetBitLockerRemovableDrivePolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bitLockerSystemDrivePolicy", m.GetBitLockerSystemDrivePolicy())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAdditionalGuardedFolders() != nil {
        err = writer.WriteCollectionOfStringValues("defenderAdditionalGuardedFolders", m.GetDefenderAdditionalGuardedFolders())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAdobeReaderLaunchChildProcess() != nil {
        cast := (*m.GetDefenderAdobeReaderLaunchChildProcess()).String()
        err = writer.WriteStringValue("defenderAdobeReaderLaunchChildProcess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAdvancedRansomewareProtectionType() != nil {
        cast := (*m.GetDefenderAdvancedRansomewareProtectionType()).String()
        err = writer.WriteStringValue("defenderAdvancedRansomewareProtectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowBehaviorMonitoring", m.GetDefenderAllowBehaviorMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowCloudProtection", m.GetDefenderAllowCloudProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowEndUserAccess", m.GetDefenderAllowEndUserAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowIntrusionPreventionSystem", m.GetDefenderAllowIntrusionPreventionSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowOnAccessProtection", m.GetDefenderAllowOnAccessProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowRealTimeMonitoring", m.GetDefenderAllowRealTimeMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanArchiveFiles", m.GetDefenderAllowScanArchiveFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanDownloads", m.GetDefenderAllowScanDownloads())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanNetworkFiles", m.GetDefenderAllowScanNetworkFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanRemovableDrivesDuringFullScan", m.GetDefenderAllowScanRemovableDrivesDuringFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderAllowScanScriptsLoadedInInternetExplorer", m.GetDefenderAllowScanScriptsLoadedInInternetExplorer())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAttackSurfaceReductionExcludedPaths() != nil {
        err = writer.WriteCollectionOfStringValues("defenderAttackSurfaceReductionExcludedPaths", m.GetDefenderAttackSurfaceReductionExcludedPaths())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderBlockEndUserAccess", m.GetDefenderBlockEndUserAccess())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderBlockPersistenceThroughWmiType() != nil {
        cast := (*m.GetDefenderBlockPersistenceThroughWmiType()).String()
        err = writer.WriteStringValue("defenderBlockPersistenceThroughWmiType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderCheckForSignaturesBeforeRunningScan", m.GetDefenderCheckForSignaturesBeforeRunningScan())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderCloudBlockLevel() != nil {
        cast := (*m.GetDefenderCloudBlockLevel()).String()
        err = writer.WriteStringValue("defenderCloudBlockLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderCloudExtendedTimeoutInSeconds", m.GetDefenderCloudExtendedTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderDaysBeforeDeletingQuarantinedMalware", m.GetDefenderDaysBeforeDeletingQuarantinedMalware())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defenderDetectedMalwareActions", m.GetDefenderDetectedMalwareActions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableBehaviorMonitoring", m.GetDefenderDisableBehaviorMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableCatchupFullScan", m.GetDefenderDisableCatchupFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableCatchupQuickScan", m.GetDefenderDisableCatchupQuickScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableCloudProtection", m.GetDefenderDisableCloudProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableIntrusionPreventionSystem", m.GetDefenderDisableIntrusionPreventionSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableOnAccessProtection", m.GetDefenderDisableOnAccessProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableRealTimeMonitoring", m.GetDefenderDisableRealTimeMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanArchiveFiles", m.GetDefenderDisableScanArchiveFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanDownloads", m.GetDefenderDisableScanDownloads())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanNetworkFiles", m.GetDefenderDisableScanNetworkFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanRemovableDrivesDuringFullScan", m.GetDefenderDisableScanRemovableDrivesDuringFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderDisableScanScriptsLoadedInInternetExplorer", m.GetDefenderDisableScanScriptsLoadedInInternetExplorer())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderEmailContentExecution() != nil {
        cast := (*m.GetDefenderEmailContentExecution()).String()
        err = writer.WriteStringValue("defenderEmailContentExecution", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderEmailContentExecutionType() != nil {
        cast := (*m.GetDefenderEmailContentExecutionType()).String()
        err = writer.WriteStringValue("defenderEmailContentExecutionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderEnableLowCpuPriority", m.GetDefenderEnableLowCpuPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderEnableScanIncomingMail", m.GetDefenderEnableScanIncomingMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderEnableScanMappedNetworkDrivesDuringFullScan", m.GetDefenderEnableScanMappedNetworkDrivesDuringFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("defenderExploitProtectionXml", m.GetDefenderExploitProtectionXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderExploitProtectionXmlFileName", m.GetDefenderExploitProtectionXmlFileName())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderFileExtensionsToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderFileExtensionsToExclude", m.GetDefenderFileExtensionsToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderFilesAndFoldersToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderFilesAndFoldersToExclude", m.GetDefenderFilesAndFoldersToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderGuardedFoldersAllowedAppPaths() != nil {
        err = writer.WriteCollectionOfStringValues("defenderGuardedFoldersAllowedAppPaths", m.GetDefenderGuardedFoldersAllowedAppPaths())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderGuardMyFoldersType() != nil {
        cast := (*m.GetDefenderGuardMyFoldersType()).String()
        err = writer.WriteStringValue("defenderGuardMyFoldersType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderNetworkProtectionType() != nil {
        cast := (*m.GetDefenderNetworkProtectionType()).String()
        err = writer.WriteStringValue("defenderNetworkProtectionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsExecutableContentCreationOrLaunch() != nil {
        cast := (*m.GetDefenderOfficeAppsExecutableContentCreationOrLaunch()).String()
        err = writer.WriteStringValue("defenderOfficeAppsExecutableContentCreationOrLaunch", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsExecutableContentCreationOrLaunchType() != nil {
        cast := (*m.GetDefenderOfficeAppsExecutableContentCreationOrLaunchType()).String()
        err = writer.WriteStringValue("defenderOfficeAppsExecutableContentCreationOrLaunchType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsLaunchChildProcess() != nil {
        cast := (*m.GetDefenderOfficeAppsLaunchChildProcess()).String()
        err = writer.WriteStringValue("defenderOfficeAppsLaunchChildProcess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsLaunchChildProcessType() != nil {
        cast := (*m.GetDefenderOfficeAppsLaunchChildProcessType()).String()
        err = writer.WriteStringValue("defenderOfficeAppsLaunchChildProcessType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsOtherProcessInjection() != nil {
        cast := (*m.GetDefenderOfficeAppsOtherProcessInjection()).String()
        err = writer.WriteStringValue("defenderOfficeAppsOtherProcessInjection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeAppsOtherProcessInjectionType() != nil {
        cast := (*m.GetDefenderOfficeAppsOtherProcessInjectionType()).String()
        err = writer.WriteStringValue("defenderOfficeAppsOtherProcessInjectionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeCommunicationAppsLaunchChildProcess() != nil {
        cast := (*m.GetDefenderOfficeCommunicationAppsLaunchChildProcess()).String()
        err = writer.WriteStringValue("defenderOfficeCommunicationAppsLaunchChildProcess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeMacroCodeAllowWin32Imports() != nil {
        cast := (*m.GetDefenderOfficeMacroCodeAllowWin32Imports()).String()
        err = writer.WriteStringValue("defenderOfficeMacroCodeAllowWin32Imports", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderOfficeMacroCodeAllowWin32ImportsType() != nil {
        cast := (*m.GetDefenderOfficeMacroCodeAllowWin32ImportsType()).String()
        err = writer.WriteStringValue("defenderOfficeMacroCodeAllowWin32ImportsType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderPotentiallyUnwantedAppAction() != nil {
        cast := (*m.GetDefenderPotentiallyUnwantedAppAction()).String()
        err = writer.WriteStringValue("defenderPotentiallyUnwantedAppAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderPreventCredentialStealingType() != nil {
        cast := (*m.GetDefenderPreventCredentialStealingType()).String()
        err = writer.WriteStringValue("defenderPreventCredentialStealingType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderProcessCreation() != nil {
        cast := (*m.GetDefenderProcessCreation()).String()
        err = writer.WriteStringValue("defenderProcessCreation", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderProcessCreationType() != nil {
        cast := (*m.GetDefenderProcessCreationType()).String()
        err = writer.WriteStringValue("defenderProcessCreationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderProcessesToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderProcessesToExclude", m.GetDefenderProcessesToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScanDirection() != nil {
        cast := (*m.GetDefenderScanDirection()).String()
        err = writer.WriteStringValue("defenderScanDirection", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderScanMaxCpuPercentage", m.GetDefenderScanMaxCpuPercentage())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScanType() != nil {
        cast := (*m.GetDefenderScanType()).String()
        err = writer.WriteStringValue("defenderScanType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("defenderScheduledQuickScanTime", m.GetDefenderScheduledQuickScanTime())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScheduledScanDay() != nil {
        cast := (*m.GetDefenderScheduledScanDay()).String()
        err = writer.WriteStringValue("defenderScheduledScanDay", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("defenderScheduledScanTime", m.GetDefenderScheduledScanTime())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScriptDownloadedPayloadExecution() != nil {
        cast := (*m.GetDefenderScriptDownloadedPayloadExecution()).String()
        err = writer.WriteStringValue("defenderScriptDownloadedPayloadExecution", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScriptDownloadedPayloadExecutionType() != nil {
        cast := (*m.GetDefenderScriptDownloadedPayloadExecutionType()).String()
        err = writer.WriteStringValue("defenderScriptDownloadedPayloadExecutionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScriptObfuscatedMacroCode() != nil {
        cast := (*m.GetDefenderScriptObfuscatedMacroCode()).String()
        err = writer.WriteStringValue("defenderScriptObfuscatedMacroCode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScriptObfuscatedMacroCodeType() != nil {
        cast := (*m.GetDefenderScriptObfuscatedMacroCodeType()).String()
        err = writer.WriteStringValue("defenderScriptObfuscatedMacroCodeType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterBlockExploitProtectionOverride", m.GetDefenderSecurityCenterBlockExploitProtectionOverride())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableAccountUI", m.GetDefenderSecurityCenterDisableAccountUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableAppBrowserUI", m.GetDefenderSecurityCenterDisableAppBrowserUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableClearTpmUI", m.GetDefenderSecurityCenterDisableClearTpmUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableFamilyUI", m.GetDefenderSecurityCenterDisableFamilyUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableHardwareUI", m.GetDefenderSecurityCenterDisableHardwareUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableHealthUI", m.GetDefenderSecurityCenterDisableHealthUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableNetworkUI", m.GetDefenderSecurityCenterDisableNetworkUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableNotificationAreaUI", m.GetDefenderSecurityCenterDisableNotificationAreaUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableRansomwareUI", m.GetDefenderSecurityCenterDisableRansomwareUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableSecureBootUI", m.GetDefenderSecurityCenterDisableSecureBootUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableTroubleshootingUI", m.GetDefenderSecurityCenterDisableTroubleshootingUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableVirusUI", m.GetDefenderSecurityCenterDisableVirusUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI", m.GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderSecurityCenterHelpEmail", m.GetDefenderSecurityCenterHelpEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderSecurityCenterHelpPhone", m.GetDefenderSecurityCenterHelpPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderSecurityCenterHelpURL", m.GetDefenderSecurityCenterHelpURL())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderSecurityCenterITContactDisplay() != nil {
        cast := (*m.GetDefenderSecurityCenterITContactDisplay()).String()
        err = writer.WriteStringValue("defenderSecurityCenterITContactDisplay", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderSecurityCenterNotificationsFromApp() != nil {
        cast := (*m.GetDefenderSecurityCenterNotificationsFromApp()).String()
        err = writer.WriteStringValue("defenderSecurityCenterNotificationsFromApp", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderSecurityCenterOrganizationDisplayName", m.GetDefenderSecurityCenterOrganizationDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderSignatureUpdateIntervalInHours", m.GetDefenderSignatureUpdateIntervalInHours())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderSubmitSamplesConsentType() != nil {
        cast := (*m.GetDefenderSubmitSamplesConsentType()).String()
        err = writer.WriteStringValue("defenderSubmitSamplesConsentType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderUntrustedExecutable() != nil {
        cast := (*m.GetDefenderUntrustedExecutable()).String()
        err = writer.WriteStringValue("defenderUntrustedExecutable", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderUntrustedExecutableType() != nil {
        cast := (*m.GetDefenderUntrustedExecutableType()).String()
        err = writer.WriteStringValue("defenderUntrustedExecutableType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderUntrustedUSBProcess() != nil {
        cast := (*m.GetDefenderUntrustedUSBProcess()).String()
        err = writer.WriteStringValue("defenderUntrustedUSBProcess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderUntrustedUSBProcessType() != nil {
        cast := (*m.GetDefenderUntrustedUSBProcessType()).String()
        err = writer.WriteStringValue("defenderUntrustedUSBProcessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceGuardEnableSecureBootWithDMA", m.GetDeviceGuardEnableSecureBootWithDMA())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceGuardEnableVirtualizationBasedSecurity", m.GetDeviceGuardEnableVirtualizationBasedSecurity())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardLaunchSystemGuard() != nil {
        cast := (*m.GetDeviceGuardLaunchSystemGuard()).String()
        err = writer.WriteStringValue("deviceGuardLaunchSystemGuard", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings() != nil {
        cast := (*m.GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings()).String()
        err = writer.WriteStringValue("deviceGuardLocalSystemAuthorityCredentialGuardSettings", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardSecureBootWithDMA() != nil {
        cast := (*m.GetDeviceGuardSecureBootWithDMA()).String()
        err = writer.WriteStringValue("deviceGuardSecureBootWithDMA", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDmaGuardDeviceEnumerationPolicy() != nil {
        cast := (*m.GetDmaGuardDeviceEnumerationPolicy()).String()
        err = writer.WriteStringValue("dmaGuardDeviceEnumerationPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallBlockStatefulFTP", m.GetFirewallBlockStatefulFTP())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallCertificateRevocationListCheckMethod() != nil {
        cast := (*m.GetFirewallCertificateRevocationListCheckMethod()).String()
        err = writer.WriteStringValue("firewallCertificateRevocationListCheckMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("firewallIdleTimeoutForSecurityAssociationInSeconds", m.GetFirewallIdleTimeoutForSecurityAssociationInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowDHCP", m.GetFirewallIPSecExemptionsAllowDHCP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowICMP", m.GetFirewallIPSecExemptionsAllowICMP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowNeighborDiscovery", m.GetFirewallIPSecExemptionsAllowNeighborDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowRouterDiscovery", m.GetFirewallIPSecExemptionsAllowRouterDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsNone", m.GetFirewallIPSecExemptionsNone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallMergeKeyingModuleSettings", m.GetFirewallMergeKeyingModuleSettings())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallPacketQueueingMethod() != nil {
        cast := (*m.GetFirewallPacketQueueingMethod()).String()
        err = writer.WriteStringValue("firewallPacketQueueingMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFirewallPreSharedKeyEncodingMethod() != nil {
        cast := (*m.GetFirewallPreSharedKeyEncodingMethod()).String()
        err = writer.WriteStringValue("firewallPreSharedKeyEncodingMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfileDomain", m.GetFirewallProfileDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfilePrivate", m.GetFirewallProfilePrivate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfilePublic", m.GetFirewallProfilePublic())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFirewallRules()))
        for i, v := range m.GetFirewallRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("firewallRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLanManagerAuthenticationLevel() != nil {
        cast := (*m.GetLanManagerAuthenticationLevel()).String()
        err = writer.WriteStringValue("lanManagerAuthenticationLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lanManagerWorkstationDisableInsecureGuestLogons", m.GetLanManagerWorkstationDisableInsecureGuestLogons())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsAdministratorAccountName", m.GetLocalSecurityOptionsAdministratorAccountName())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsAdministratorElevationPromptBehavior() != nil {
        cast := (*m.GetLocalSecurityOptionsAdministratorElevationPromptBehavior()).String()
        err = writer.WriteStringValue("localSecurityOptionsAdministratorElevationPromptBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares", m.GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowPKU2UAuthenticationRequests", m.GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager", m.GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool", m.GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn", m.GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowUIAccessApplicationElevation", m.GetLocalSecurityOptionsAllowUIAccessApplicationElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowUIAccessApplicationsForSecureLocations", m.GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsAllowUndockWithoutHavingToLogon", m.GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsBlockMicrosoftAccounts", m.GetLocalSecurityOptionsBlockMicrosoftAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsBlockRemoteLogonWithBlankPassword", m.GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsBlockRemoteOpticalDriveAccess", m.GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsBlockUsersInstallingPrinterDrivers", m.GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsClearVirtualMemoryPageFile", m.GetLocalSecurityOptionsClearVirtualMemoryPageFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsClientDigitallySignCommunicationsAlways", m.GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers", m.GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation", m.GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableAdministratorAccount", m.GetLocalSecurityOptionsDisableAdministratorAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees", m.GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableGuestAccount", m.GetLocalSecurityOptionsDisableGuestAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableServerDigitallySignCommunicationsAlways", m.GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees", m.GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts", m.GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDoNotRequireCtrlAltDel", m.GetLocalSecurityOptionsDoNotRequireCtrlAltDel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange", m.GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser() != nil {
        cast := (*m.GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser()).String()
        err = writer.WriteStringValue("localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsGuestAccountName", m.GetLocalSecurityOptionsGuestAccountName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsHideLastSignedInUser", m.GetLocalSecurityOptionsHideLastSignedInUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsHideUsernameAtSignIn", m.GetLocalSecurityOptionsHideUsernameAtSignIn())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsInformationDisplayedOnLockScreen() != nil {
        cast := (*m.GetLocalSecurityOptionsInformationDisplayedOnLockScreen()).String()
        err = writer.WriteStringValue("localSecurityOptionsInformationDisplayedOnLockScreen", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsInformationShownOnLockScreen() != nil {
        cast := (*m.GetLocalSecurityOptionsInformationShownOnLockScreen()).String()
        err = writer.WriteStringValue("localSecurityOptionsInformationShownOnLockScreen", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsLogOnMessageText", m.GetLocalSecurityOptionsLogOnMessageText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localSecurityOptionsLogOnMessageTitle", m.GetLocalSecurityOptionsLogOnMessageTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("localSecurityOptionsMachineInactivityLimit", m.GetLocalSecurityOptionsMachineInactivityLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("localSecurityOptionsMachineInactivityLimitInMinutes", m.GetLocalSecurityOptionsMachineInactivityLimitInMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients() != nil {
        cast := (*m.GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients()).String()
        err = writer.WriteStringValue("localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers() != nil {
        cast := (*m.GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers()).String()
        err = writer.WriteStringValue("localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsOnlyElevateSignedExecutables", m.GetLocalSecurityOptionsOnlyElevateSignedExecutables())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares", m.GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares())
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsSmartCardRemovalBehavior() != nil {
        cast := (*m.GetLocalSecurityOptionsSmartCardRemovalBehavior()).String()
        err = writer.WriteStringValue("localSecurityOptionsSmartCardRemovalBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalSecurityOptionsStandardUserElevationPromptBehavior() != nil {
        cast := (*m.GetLocalSecurityOptionsStandardUserElevationPromptBehavior()).String()
        err = writer.WriteStringValue("localSecurityOptionsStandardUserElevationPromptBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation", m.GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsUseAdminApprovalMode", m.GetLocalSecurityOptionsUseAdminApprovalMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsUseAdminApprovalModeForAdministrators", m.GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations", m.GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenBlockOverrideForFiles", m.GetSmartScreenBlockOverrideForFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenEnableInShell", m.GetSmartScreenEnableInShell())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsAccessCredentialManagerAsTrustedCaller", m.GetUserRightsAccessCredentialManagerAsTrustedCaller())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsActAsPartOfTheOperatingSystem", m.GetUserRightsActAsPartOfTheOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsAllowAccessFromNetwork", m.GetUserRightsAllowAccessFromNetwork())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsBackupData", m.GetUserRightsBackupData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsBlockAccessFromNetwork", m.GetUserRightsBlockAccessFromNetwork())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsChangeSystemTime", m.GetUserRightsChangeSystemTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreateGlobalObjects", m.GetUserRightsCreateGlobalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreatePageFile", m.GetUserRightsCreatePageFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreatePermanentSharedObjects", m.GetUserRightsCreatePermanentSharedObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreateSymbolicLinks", m.GetUserRightsCreateSymbolicLinks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsCreateToken", m.GetUserRightsCreateToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsDebugPrograms", m.GetUserRightsDebugPrograms())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsDelegation", m.GetUserRightsDelegation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsDenyLocalLogOn", m.GetUserRightsDenyLocalLogOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsGenerateSecurityAudits", m.GetUserRightsGenerateSecurityAudits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsImpersonateClient", m.GetUserRightsImpersonateClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsIncreaseSchedulingPriority", m.GetUserRightsIncreaseSchedulingPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsLoadUnloadDrivers", m.GetUserRightsLoadUnloadDrivers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsLocalLogOn", m.GetUserRightsLocalLogOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsLockMemory", m.GetUserRightsLockMemory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsManageAuditingAndSecurityLogs", m.GetUserRightsManageAuditingAndSecurityLogs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsManageVolumes", m.GetUserRightsManageVolumes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsModifyFirmwareEnvironment", m.GetUserRightsModifyFirmwareEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsModifyObjectLabels", m.GetUserRightsModifyObjectLabels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsProfileSingleProcess", m.GetUserRightsProfileSingleProcess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsRemoteDesktopServicesLogOn", m.GetUserRightsRemoteDesktopServicesLogOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsRemoteShutdown", m.GetUserRightsRemoteShutdown())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsRestoreData", m.GetUserRightsRestoreData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userRightsTakeOwnership", m.GetUserRightsTakeOwnership())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsDefenderTamperProtection() != nil {
        cast := (*m.GetWindowsDefenderTamperProtection()).String()
        err = writer.WriteStringValue("windowsDefenderTamperProtection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetXboxServicesAccessoryManagementServiceStartupMode() != nil {
        cast := (*m.GetXboxServicesAccessoryManagementServiceStartupMode()).String()
        err = writer.WriteStringValue("xboxServicesAccessoryManagementServiceStartupMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("xboxServicesEnableXboxGameSaveTask", m.GetXboxServicesEnableXboxGameSaveTask())
        if err != nil {
            return err
        }
    }
    if m.GetXboxServicesLiveAuthManagerServiceStartupMode() != nil {
        cast := (*m.GetXboxServicesLiveAuthManagerServiceStartupMode()).String()
        err = writer.WriteStringValue("xboxServicesLiveAuthManagerServiceStartupMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetXboxServicesLiveGameSaveServiceStartupMode() != nil {
        cast := (*m.GetXboxServicesLiveGameSaveServiceStartupMode()).String()
        err = writer.WriteStringValue("xboxServicesLiveGameSaveServiceStartupMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetXboxServicesLiveNetworkingServiceStartupMode() != nil {
        cast := (*m.GetXboxServicesLiveNetworkingServiceStartupMode()).String()
        err = writer.WriteStringValue("xboxServicesLiveNetworkingServiceStartupMode", &cast)
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
func (m *Windows10EndpointProtectionConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationGuardAllowCameraMicrophoneRedirection sets the applicationGuardAllowCameraMicrophoneRedirection property value. Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and microphone.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowCameraMicrophoneRedirection(value *bool)() {
    if m != nil {
        m.applicationGuardAllowCameraMicrophoneRedirection = value
    }
}
// SetApplicationGuardAllowFileSaveOnHost sets the applicationGuardAllowFileSaveOnHost property value. Allow users to download files from Edge in the application guard container and save them on the host file system
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowFileSaveOnHost(value *bool)() {
    if m != nil {
        m.applicationGuardAllowFileSaveOnHost = value
    }
}
// SetApplicationGuardAllowPersistence sets the applicationGuardAllowPersistence property value. Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPersistence(value *bool)() {
    if m != nil {
        m.applicationGuardAllowPersistence = value
    }
}
// SetApplicationGuardAllowPrintToLocalPrinters sets the applicationGuardAllowPrintToLocalPrinters property value. Allow printing to Local Printers from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToLocalPrinters(value *bool)() {
    if m != nil {
        m.applicationGuardAllowPrintToLocalPrinters = value
    }
}
// SetApplicationGuardAllowPrintToNetworkPrinters sets the applicationGuardAllowPrintToNetworkPrinters property value. Allow printing to Network Printers from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToNetworkPrinters(value *bool)() {
    if m != nil {
        m.applicationGuardAllowPrintToNetworkPrinters = value
    }
}
// SetApplicationGuardAllowPrintToPDF sets the applicationGuardAllowPrintToPDF property value. Allow printing to PDF from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToPDF(value *bool)() {
    if m != nil {
        m.applicationGuardAllowPrintToPDF = value
    }
}
// SetApplicationGuardAllowPrintToXPS sets the applicationGuardAllowPrintToXPS property value. Allow printing to XPS from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToXPS(value *bool)() {
    if m != nil {
        m.applicationGuardAllowPrintToXPS = value
    }
}
// SetApplicationGuardAllowVirtualGPU sets the applicationGuardAllowVirtualGPU property value. Allow application guard to use virtual GPU
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowVirtualGPU(value *bool)() {
    if m != nil {
        m.applicationGuardAllowVirtualGPU = value
    }
}
// SetApplicationGuardBlockClipboardSharing sets the applicationGuardBlockClipboardSharing property value. Block clipboard to share data from Host to Container, or from Container to Host, or both ways, or neither ways. Possible values are: notConfigured, blockBoth, blockHostToContainer, blockContainerToHost, blockNone.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockClipboardSharing(value *ApplicationGuardBlockClipboardSharingType)() {
    if m != nil {
        m.applicationGuardBlockClipboardSharing = value
    }
}
// SetApplicationGuardBlockFileTransfer sets the applicationGuardBlockFileTransfer property value. Block clipboard to transfer image file, text file or neither of them. Possible values are: notConfigured, blockImageAndTextFile, blockImageFile, blockNone, blockTextFile.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockFileTransfer(value *ApplicationGuardBlockFileTransferType)() {
    if m != nil {
        m.applicationGuardBlockFileTransfer = value
    }
}
// SetApplicationGuardBlockNonEnterpriseContent sets the applicationGuardBlockNonEnterpriseContent property value. Block enterprise sites to load non-enterprise content, such as third party plug-ins
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockNonEnterpriseContent(value *bool)() {
    if m != nil {
        m.applicationGuardBlockNonEnterpriseContent = value
    }
}
// SetApplicationGuardCertificateThumbprints sets the applicationGuardCertificateThumbprints property value. Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardCertificateThumbprints(value []string)() {
    if m != nil {
        m.applicationGuardCertificateThumbprints = value
    }
}
// SetApplicationGuardEnabled sets the applicationGuardEnabled property value. Enable Windows Defender Application Guard
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardEnabled(value *bool)() {
    if m != nil {
        m.applicationGuardEnabled = value
    }
}
// SetApplicationGuardEnabledOptions sets the applicationGuardEnabledOptions property value. Enable Windows Defender Application Guard for newer Windows builds. Possible values are: notConfigured, enabledForEdge, enabledForOffice, enabledForEdgeAndOffice.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardEnabledOptions(value *ApplicationGuardEnabledOptions)() {
    if m != nil {
        m.applicationGuardEnabledOptions = value
    }
}
// SetApplicationGuardForceAuditing sets the applicationGuardForceAuditing property value. Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardForceAuditing(value *bool)() {
    if m != nil {
        m.applicationGuardForceAuditing = value
    }
}
// SetAppLockerApplicationControl sets the appLockerApplicationControl property value. Enables the Admin to choose what types of app to allow on devices. Possible values are: notConfigured, enforceComponentsAndStoreApps, auditComponentsAndStoreApps, enforceComponentsStoreAppsAndSmartlocker, auditComponentsStoreAppsAndSmartlocker.
func (m *Windows10EndpointProtectionConfiguration) SetAppLockerApplicationControl(value *AppLockerApplicationControlType)() {
    if m != nil {
        m.appLockerApplicationControl = value
    }
}
// SetBitLockerAllowStandardUserEncryption sets the bitLockerAllowStandardUserEncryption property value. Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerAllowStandardUserEncryption(value *bool)() {
    if m != nil {
        m.bitLockerAllowStandardUserEncryption = value
    }
}
// SetBitLockerDisableWarningForOtherDiskEncryption sets the bitLockerDisableWarningForOtherDiskEncryption property value. Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerDisableWarningForOtherDiskEncryption(value *bool)() {
    if m != nil {
        m.bitLockerDisableWarningForOtherDiskEncryption = value
    }
}
// SetBitLockerEnableStorageCardEncryptionOnMobile sets the bitLockerEnableStorageCardEncryptionOnMobile property value. Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerEnableStorageCardEncryptionOnMobile(value *bool)() {
    if m != nil {
        m.bitLockerEnableStorageCardEncryptionOnMobile = value
    }
}
// SetBitLockerEncryptDevice sets the bitLockerEncryptDevice property value. Allows the admin to require encryption to be turned on using BitLocker.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerEncryptDevice(value *bool)() {
    if m != nil {
        m.bitLockerEncryptDevice = value
    }
}
// SetBitLockerFixedDrivePolicy sets the bitLockerFixedDrivePolicy property value. BitLocker Fixed Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerFixedDrivePolicy(value BitLockerFixedDrivePolicyable)() {
    if m != nil {
        m.bitLockerFixedDrivePolicy = value
    }
}
// SetBitLockerRecoveryPasswordRotation sets the bitLockerRecoveryPasswordRotation property value. This setting initiates a client-driven recovery password rotation after an OS drive recovery (either by using bootmgr or WinRE). Possible values are: notConfigured, disabled, enabledForAzureAd, enabledForAzureAdAndHybrid.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerRecoveryPasswordRotation(value *BitLockerRecoveryPasswordRotationType)() {
    if m != nil {
        m.bitLockerRecoveryPasswordRotation = value
    }
}
// SetBitLockerRemovableDrivePolicy sets the bitLockerRemovableDrivePolicy property value. BitLocker Removable Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerRemovableDrivePolicy(value BitLockerRemovableDrivePolicyable)() {
    if m != nil {
        m.bitLockerRemovableDrivePolicy = value
    }
}
// SetBitLockerSystemDrivePolicy sets the bitLockerSystemDrivePolicy property value. BitLocker System Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerSystemDrivePolicy(value BitLockerSystemDrivePolicyable)() {
    if m != nil {
        m.bitLockerSystemDrivePolicy = value
    }
}
// SetDefenderAdditionalGuardedFolders sets the defenderAdditionalGuardedFolders property value. List of folder paths to be added to the list of protected folders
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdditionalGuardedFolders(value []string)() {
    if m != nil {
        m.defenderAdditionalGuardedFolders = value
    }
}
// SetDefenderAdobeReaderLaunchChildProcess sets the defenderAdobeReaderLaunchChildProcess property value. Value indicating the behavior of Adobe Reader from creating child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdobeReaderLaunchChildProcess(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderAdobeReaderLaunchChildProcess = value
    }
}
// SetDefenderAdvancedRansomewareProtectionType sets the defenderAdvancedRansomewareProtectionType property value. Value indicating use of advanced protection against ransomeware. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdvancedRansomewareProtectionType(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderAdvancedRansomewareProtectionType = value
    }
}
// SetDefenderAllowBehaviorMonitoring sets the defenderAllowBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowBehaviorMonitoring(value *bool)() {
    if m != nil {
        m.defenderAllowBehaviorMonitoring = value
    }
}
// SetDefenderAllowCloudProtection sets the defenderAllowCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowCloudProtection(value *bool)() {
    if m != nil {
        m.defenderAllowCloudProtection = value
    }
}
// SetDefenderAllowEndUserAccess sets the defenderAllowEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowEndUserAccess(value *bool)() {
    if m != nil {
        m.defenderAllowEndUserAccess = value
    }
}
// SetDefenderAllowIntrusionPreventionSystem sets the defenderAllowIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowIntrusionPreventionSystem(value *bool)() {
    if m != nil {
        m.defenderAllowIntrusionPreventionSystem = value
    }
}
// SetDefenderAllowOnAccessProtection sets the defenderAllowOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowOnAccessProtection(value *bool)() {
    if m != nil {
        m.defenderAllowOnAccessProtection = value
    }
}
// SetDefenderAllowRealTimeMonitoring sets the defenderAllowRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowRealTimeMonitoring(value *bool)() {
    if m != nil {
        m.defenderAllowRealTimeMonitoring = value
    }
}
// SetDefenderAllowScanArchiveFiles sets the defenderAllowScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanArchiveFiles(value *bool)() {
    if m != nil {
        m.defenderAllowScanArchiveFiles = value
    }
}
// SetDefenderAllowScanDownloads sets the defenderAllowScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanDownloads(value *bool)() {
    if m != nil {
        m.defenderAllowScanDownloads = value
    }
}
// SetDefenderAllowScanNetworkFiles sets the defenderAllowScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanNetworkFiles(value *bool)() {
    if m != nil {
        m.defenderAllowScanNetworkFiles = value
    }
}
// SetDefenderAllowScanRemovableDrivesDuringFullScan sets the defenderAllowScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanRemovableDrivesDuringFullScan(value *bool)() {
    if m != nil {
        m.defenderAllowScanRemovableDrivesDuringFullScan = value
    }
}
// SetDefenderAllowScanScriptsLoadedInInternetExplorer sets the defenderAllowScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanScriptsLoadedInInternetExplorer(value *bool)() {
    if m != nil {
        m.defenderAllowScanScriptsLoadedInInternetExplorer = value
    }
}
// SetDefenderAttackSurfaceReductionExcludedPaths sets the defenderAttackSurfaceReductionExcludedPaths property value. List of exe files and folders to be excluded from attack surface reduction rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAttackSurfaceReductionExcludedPaths(value []string)() {
    if m != nil {
        m.defenderAttackSurfaceReductionExcludedPaths = value
    }
}
// SetDefenderBlockEndUserAccess sets the defenderBlockEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderBlockEndUserAccess(value *bool)() {
    if m != nil {
        m.defenderBlockEndUserAccess = value
    }
}
// SetDefenderBlockPersistenceThroughWmiType sets the defenderBlockPersistenceThroughWmiType property value. Value indicating the behavior of Block persistence through WMI event subscription. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderBlockPersistenceThroughWmiType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderBlockPersistenceThroughWmiType = value
    }
}
// SetDefenderCheckForSignaturesBeforeRunningScan sets the defenderCheckForSignaturesBeforeRunningScan property value. This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before running a scan.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCheckForSignaturesBeforeRunningScan(value *bool)() {
    if m != nil {
        m.defenderCheckForSignaturesBeforeRunningScan = value
    }
}
// SetDefenderCloudBlockLevel sets the defenderCloudBlockLevel property value. Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS' setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCloudBlockLevel(value *DefenderCloudBlockLevelType)() {
    if m != nil {
        m.defenderCloudBlockLevel = value
    }
}
// SetDefenderCloudExtendedTimeoutInSeconds sets the defenderCloudExtendedTimeoutInSeconds property value. Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCloudExtendedTimeoutInSeconds(value *int32)() {
    if m != nil {
        m.defenderCloudExtendedTimeoutInSeconds = value
    }
}
// SetDefenderDaysBeforeDeletingQuarantinedMalware sets the defenderDaysBeforeDeletingQuarantinedMalware property value. Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDaysBeforeDeletingQuarantinedMalware(value *int32)() {
    if m != nil {
        m.defenderDaysBeforeDeletingQuarantinedMalware = value
    }
}
// SetDefenderDetectedMalwareActions sets the defenderDetectedMalwareActions property value. Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDetectedMalwareActions(value DefenderDetectedMalwareActionsable)() {
    if m != nil {
        m.defenderDetectedMalwareActions = value
    }
}
// SetDefenderDisableBehaviorMonitoring sets the defenderDisableBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableBehaviorMonitoring(value *bool)() {
    if m != nil {
        m.defenderDisableBehaviorMonitoring = value
    }
}
// SetDefenderDisableCatchupFullScan sets the defenderDisableCatchupFullScan property value. This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCatchupFullScan(value *bool)() {
    if m != nil {
        m.defenderDisableCatchupFullScan = value
    }
}
// SetDefenderDisableCatchupQuickScan sets the defenderDisableCatchupQuickScan property value. This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCatchupQuickScan(value *bool)() {
    if m != nil {
        m.defenderDisableCatchupQuickScan = value
    }
}
// SetDefenderDisableCloudProtection sets the defenderDisableCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCloudProtection(value *bool)() {
    if m != nil {
        m.defenderDisableCloudProtection = value
    }
}
// SetDefenderDisableIntrusionPreventionSystem sets the defenderDisableIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableIntrusionPreventionSystem(value *bool)() {
    if m != nil {
        m.defenderDisableIntrusionPreventionSystem = value
    }
}
// SetDefenderDisableOnAccessProtection sets the defenderDisableOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableOnAccessProtection(value *bool)() {
    if m != nil {
        m.defenderDisableOnAccessProtection = value
    }
}
// SetDefenderDisableRealTimeMonitoring sets the defenderDisableRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableRealTimeMonitoring(value *bool)() {
    if m != nil {
        m.defenderDisableRealTimeMonitoring = value
    }
}
// SetDefenderDisableScanArchiveFiles sets the defenderDisableScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanArchiveFiles(value *bool)() {
    if m != nil {
        m.defenderDisableScanArchiveFiles = value
    }
}
// SetDefenderDisableScanDownloads sets the defenderDisableScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanDownloads(value *bool)() {
    if m != nil {
        m.defenderDisableScanDownloads = value
    }
}
// SetDefenderDisableScanNetworkFiles sets the defenderDisableScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanNetworkFiles(value *bool)() {
    if m != nil {
        m.defenderDisableScanNetworkFiles = value
    }
}
// SetDefenderDisableScanRemovableDrivesDuringFullScan sets the defenderDisableScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanRemovableDrivesDuringFullScan(value *bool)() {
    if m != nil {
        m.defenderDisableScanRemovableDrivesDuringFullScan = value
    }
}
// SetDefenderDisableScanScriptsLoadedInInternetExplorer sets the defenderDisableScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanScriptsLoadedInInternetExplorer(value *bool)() {
    if m != nil {
        m.defenderDisableScanScriptsLoadedInInternetExplorer = value
    }
}
// SetDefenderEmailContentExecution sets the defenderEmailContentExecution property value. Value indicating if execution of executable content (exe, dll, ps, js, vbs, etc) should be dropped from email (webmail/mail-client). Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEmailContentExecution(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderEmailContentExecution = value
    }
}
// SetDefenderEmailContentExecutionType sets the defenderEmailContentExecutionType property value. Value indicating if execution of executable content (exe, dll, ps, js, vbs, etc) should be dropped from email (webmail/mail-client). Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEmailContentExecutionType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderEmailContentExecutionType = value
    }
}
// SetDefenderEnableLowCpuPriority sets the defenderEnableLowCpuPriority property value. This policy setting allows you to enable or disable low CPU priority for scheduled scans.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableLowCpuPriority(value *bool)() {
    if m != nil {
        m.defenderEnableLowCpuPriority = value
    }
}
// SetDefenderEnableScanIncomingMail sets the defenderEnableScanIncomingMail property value. Allows or disallows scanning of email.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableScanIncomingMail(value *bool)() {
    if m != nil {
        m.defenderEnableScanIncomingMail = value
    }
}
// SetDefenderEnableScanMappedNetworkDrivesDuringFullScan sets the defenderEnableScanMappedNetworkDrivesDuringFullScan property value. Allows or disallows a full scan of mapped network drives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableScanMappedNetworkDrivesDuringFullScan(value *bool)() {
    if m != nil {
        m.defenderEnableScanMappedNetworkDrivesDuringFullScan = value
    }
}
// SetDefenderExploitProtectionXml sets the defenderExploitProtectionXml property value. Xml content containing information regarding exploit protection details.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXml(value []byte)() {
    if m != nil {
        m.defenderExploitProtectionXml = value
    }
}
// SetDefenderExploitProtectionXmlFileName sets the defenderExploitProtectionXmlFileName property value. Name of the file from which DefenderExploitProtectionXml was obtained.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXmlFileName(value *string)() {
    if m != nil {
        m.defenderExploitProtectionXmlFileName = value
    }
}
// SetDefenderFileExtensionsToExclude sets the defenderFileExtensionsToExclude property value. File extensions to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderFileExtensionsToExclude(value []string)() {
    if m != nil {
        m.defenderFileExtensionsToExclude = value
    }
}
// SetDefenderFilesAndFoldersToExclude sets the defenderFilesAndFoldersToExclude property value. Files and folder to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderFilesAndFoldersToExclude(value []string)() {
    if m != nil {
        m.defenderFilesAndFoldersToExclude = value
    }
}
// SetDefenderGuardedFoldersAllowedAppPaths sets the defenderGuardedFoldersAllowedAppPaths property value. List of paths to exe that are allowed to access protected folders
func (m *Windows10EndpointProtectionConfiguration) SetDefenderGuardedFoldersAllowedAppPaths(value []string)() {
    if m != nil {
        m.defenderGuardedFoldersAllowedAppPaths = value
    }
}
// SetDefenderGuardMyFoldersType sets the defenderGuardMyFoldersType property value. Value indicating the behavior of protected folders. Possible values are: userDefined, enable, auditMode, blockDiskModification, auditDiskModification.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderGuardMyFoldersType(value *FolderProtectionType)() {
    if m != nil {
        m.defenderGuardMyFoldersType = value
    }
}
// SetDefenderNetworkProtectionType sets the defenderNetworkProtectionType property value. Value indicating the behavior of NetworkProtection. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderNetworkProtectionType(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderNetworkProtectionType = value
    }
}
// SetDefenderOfficeAppsExecutableContentCreationOrLaunch sets the defenderOfficeAppsExecutableContentCreationOrLaunch property value. Value indicating the behavior of Office applications/macros creating or launching executable content. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsExecutableContentCreationOrLaunch(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderOfficeAppsExecutableContentCreationOrLaunch = value
    }
}
// SetDefenderOfficeAppsExecutableContentCreationOrLaunchType sets the defenderOfficeAppsExecutableContentCreationOrLaunchType property value. Value indicating the behavior of Office applications/macros creating or launching executable content. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsExecutableContentCreationOrLaunchType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderOfficeAppsExecutableContentCreationOrLaunchType = value
    }
}
// SetDefenderOfficeAppsLaunchChildProcess sets the defenderOfficeAppsLaunchChildProcess property value. Value indicating the behavior of Office application launching child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsLaunchChildProcess(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderOfficeAppsLaunchChildProcess = value
    }
}
// SetDefenderOfficeAppsLaunchChildProcessType sets the defenderOfficeAppsLaunchChildProcessType property value. Value indicating the behavior of Office application launching child processes. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsLaunchChildProcessType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderOfficeAppsLaunchChildProcessType = value
    }
}
// SetDefenderOfficeAppsOtherProcessInjection sets the defenderOfficeAppsOtherProcessInjection property value. Value indicating the behavior of  Office applications injecting into other processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsOtherProcessInjection(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderOfficeAppsOtherProcessInjection = value
    }
}
// SetDefenderOfficeAppsOtherProcessInjectionType sets the defenderOfficeAppsOtherProcessInjectionType property value. Value indicating the behavior of Office applications injecting into other processes. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsOtherProcessInjectionType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderOfficeAppsOtherProcessInjectionType = value
    }
}
// SetDefenderOfficeCommunicationAppsLaunchChildProcess sets the defenderOfficeCommunicationAppsLaunchChildProcess property value. Value indicating the behavior of Office communication applications, including Microsoft Outlook, from creating child processes. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeCommunicationAppsLaunchChildProcess(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderOfficeCommunicationAppsLaunchChildProcess = value
    }
}
// SetDefenderOfficeMacroCodeAllowWin32Imports sets the defenderOfficeMacroCodeAllowWin32Imports property value. Value indicating the behavior of Win32 imports from Macro code in Office. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeMacroCodeAllowWin32Imports(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderOfficeMacroCodeAllowWin32Imports = value
    }
}
// SetDefenderOfficeMacroCodeAllowWin32ImportsType sets the defenderOfficeMacroCodeAllowWin32ImportsType property value. Value indicating the behavior of Win32 imports from Macro code in Office. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeMacroCodeAllowWin32ImportsType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderOfficeMacroCodeAllowWin32ImportsType = value
    }
}
// SetDefenderPotentiallyUnwantedAppAction sets the defenderPotentiallyUnwantedAppAction property value. Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs). Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderPotentiallyUnwantedAppAction(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderPotentiallyUnwantedAppAction = value
    }
}
// SetDefenderPreventCredentialStealingType sets the defenderPreventCredentialStealingType property value. Value indicating if credential stealing from the Windows local security authority subsystem is permitted. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderPreventCredentialStealingType(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderPreventCredentialStealingType = value
    }
}
// SetDefenderProcessCreation sets the defenderProcessCreation property value. Value indicating response to process creations originating from PSExec and WMI commands. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessCreation(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderProcessCreation = value
    }
}
// SetDefenderProcessCreationType sets the defenderProcessCreationType property value. Value indicating response to process creations originating from PSExec and WMI commands. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessCreationType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderProcessCreationType = value
    }
}
// SetDefenderProcessesToExclude sets the defenderProcessesToExclude property value. Processes to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessesToExclude(value []string)() {
    if m != nil {
        m.defenderProcessesToExclude = value
    }
}
// SetDefenderScanDirection sets the defenderScanDirection property value. Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly, monitorOutgoingFilesOnly.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanDirection(value *DefenderRealtimeScanDirection)() {
    if m != nil {
        m.defenderScanDirection = value
    }
}
// SetDefenderScanMaxCpuPercentage sets the defenderScanMaxCpuPercentage property value. Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid values 0 to 100
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanMaxCpuPercentage(value *int32)() {
    if m != nil {
        m.defenderScanMaxCpuPercentage = value
    }
}
// SetDefenderScanType sets the defenderScanType property value. Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanType(value *DefenderScanType)() {
    if m != nil {
        m.defenderScanType = value
    }
}
// SetDefenderScheduledQuickScanTime sets the defenderScheduledQuickScanTime property value. Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledQuickScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.defenderScheduledQuickScanTime = value
    }
}
// SetDefenderScheduledScanDay sets the defenderScheduledScanDay property value. Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledScanDay(value *WeeklySchedule)() {
    if m != nil {
        m.defenderScheduledScanDay = value
    }
}
// SetDefenderScheduledScanTime sets the defenderScheduledScanTime property value. Selects the time of day that the Windows Defender scan should run.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.defenderScheduledScanTime = value
    }
}
// SetDefenderScriptDownloadedPayloadExecution sets the defenderScriptDownloadedPayloadExecution property value. Value indicating the behavior of js/vbs executing payload downloaded from Internet. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptDownloadedPayloadExecution(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderScriptDownloadedPayloadExecution = value
    }
}
// SetDefenderScriptDownloadedPayloadExecutionType sets the defenderScriptDownloadedPayloadExecutionType property value. Value indicating the behavior of js/vbs executing payload downloaded from Internet. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptDownloadedPayloadExecutionType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderScriptDownloadedPayloadExecutionType = value
    }
}
// SetDefenderScriptObfuscatedMacroCode sets the defenderScriptObfuscatedMacroCode property value. Value indicating the behavior of obfuscated js/vbs/ps/macro code. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptObfuscatedMacroCode(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderScriptObfuscatedMacroCode = value
    }
}
// SetDefenderScriptObfuscatedMacroCodeType sets the defenderScriptObfuscatedMacroCodeType property value. Value indicating the behavior of obfuscated js/vbs/ps/macro code. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptObfuscatedMacroCodeType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderScriptObfuscatedMacroCodeType = value
    }
}
// SetDefenderSecurityCenterBlockExploitProtectionOverride sets the defenderSecurityCenterBlockExploitProtectionOverride property value. Indicates whether or not to block user from overriding Exploit Protection settings.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterBlockExploitProtectionOverride(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterBlockExploitProtectionOverride = value
    }
}
// SetDefenderSecurityCenterDisableAccountUI sets the defenderSecurityCenterDisableAccountUI property value. Used to disable the display of the account protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableAccountUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableAccountUI = value
    }
}
// SetDefenderSecurityCenterDisableAppBrowserUI sets the defenderSecurityCenterDisableAppBrowserUI property value. Used to disable the display of the app and browser protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableAppBrowserUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableAppBrowserUI = value
    }
}
// SetDefenderSecurityCenterDisableClearTpmUI sets the defenderSecurityCenterDisableClearTpmUI property value. Used to disable the display of the Clear TPM button.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableClearTpmUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableClearTpmUI = value
    }
}
// SetDefenderSecurityCenterDisableFamilyUI sets the defenderSecurityCenterDisableFamilyUI property value. Used to disable the display of the family options area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableFamilyUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableFamilyUI = value
    }
}
// SetDefenderSecurityCenterDisableHardwareUI sets the defenderSecurityCenterDisableHardwareUI property value. Used to disable the display of the hardware protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableHardwareUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableHardwareUI = value
    }
}
// SetDefenderSecurityCenterDisableHealthUI sets the defenderSecurityCenterDisableHealthUI property value. Used to disable the display of the device performance and health area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableHealthUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableHealthUI = value
    }
}
// SetDefenderSecurityCenterDisableNetworkUI sets the defenderSecurityCenterDisableNetworkUI property value. Used to disable the display of the firewall and network protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableNetworkUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableNetworkUI = value
    }
}
// SetDefenderSecurityCenterDisableNotificationAreaUI sets the defenderSecurityCenterDisableNotificationAreaUI property value. Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot the computer for this setting to take effect.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableNotificationAreaUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableNotificationAreaUI = value
    }
}
// SetDefenderSecurityCenterDisableRansomwareUI sets the defenderSecurityCenterDisableRansomwareUI property value. Used to disable the display of the ransomware protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableRansomwareUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableRansomwareUI = value
    }
}
// SetDefenderSecurityCenterDisableSecureBootUI sets the defenderSecurityCenterDisableSecureBootUI property value. Used to disable the display of the secure boot area under Device security.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableSecureBootUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableSecureBootUI = value
    }
}
// SetDefenderSecurityCenterDisableTroubleshootingUI sets the defenderSecurityCenterDisableTroubleshootingUI property value. Used to disable the display of the security process troubleshooting under Device security.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableTroubleshootingUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableTroubleshootingUI = value
    }
}
// SetDefenderSecurityCenterDisableVirusUI sets the defenderSecurityCenterDisableVirusUI property value. Used to disable the display of the virus and threat protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableVirusUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableVirusUI = value
    }
}
// SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI sets the defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI property value. Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI(value *bool)() {
    if m != nil {
        m.defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI = value
    }
}
// SetDefenderSecurityCenterHelpEmail sets the defenderSecurityCenterHelpEmail property value. The email address that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpEmail(value *string)() {
    if m != nil {
        m.defenderSecurityCenterHelpEmail = value
    }
}
// SetDefenderSecurityCenterHelpPhone sets the defenderSecurityCenterHelpPhone property value. The phone number or Skype ID that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpPhone(value *string)() {
    if m != nil {
        m.defenderSecurityCenterHelpPhone = value
    }
}
// SetDefenderSecurityCenterHelpURL sets the defenderSecurityCenterHelpURL property value. The help portal URL this is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpURL(value *string)() {
    if m != nil {
        m.defenderSecurityCenterHelpURL = value
    }
}
// SetDefenderSecurityCenterITContactDisplay sets the defenderSecurityCenterITContactDisplay property value. Configure where to display IT contact information to end users. Possible values are: notConfigured, displayInAppAndInNotifications, displayOnlyInApp, displayOnlyInNotifications.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterITContactDisplay(value *DefenderSecurityCenterITContactDisplayType)() {
    if m != nil {
        m.defenderSecurityCenterITContactDisplay = value
    }
}
// SetDefenderSecurityCenterNotificationsFromApp sets the defenderSecurityCenterNotificationsFromApp property value. Notifications to show from the displayed areas of app. Possible values are: notConfigured, blockNoncriticalNotifications, blockAllNotifications.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterNotificationsFromApp(value *DefenderSecurityCenterNotificationsFromAppType)() {
    if m != nil {
        m.defenderSecurityCenterNotificationsFromApp = value
    }
}
// SetDefenderSecurityCenterOrganizationDisplayName sets the defenderSecurityCenterOrganizationDisplayName property value. The company name that is displayed to the users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterOrganizationDisplayName(value *string)() {
    if m != nil {
        m.defenderSecurityCenterOrganizationDisplayName = value
    }
}
// SetDefenderSignatureUpdateIntervalInHours sets the defenderSignatureUpdateIntervalInHours property value. Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSignatureUpdateIntervalInHours(value *int32)() {
    if m != nil {
        m.defenderSignatureUpdateIntervalInHours = value
    }
}
// SetDefenderSubmitSamplesConsentType sets the defenderSubmitSamplesConsentType property value. Checks for the user consent level in Windows Defender to send data. Possible values are: sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSubmitSamplesConsentType(value *DefenderSubmitSamplesConsentType)() {
    if m != nil {
        m.defenderSubmitSamplesConsentType = value
    }
}
// SetDefenderUntrustedExecutable sets the defenderUntrustedExecutable property value. Value indicating response to executables that don't meet a prevalence, age, or trusted list criteria. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedExecutable(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderUntrustedExecutable = value
    }
}
// SetDefenderUntrustedExecutableType sets the defenderUntrustedExecutableType property value. Value indicating response to executables that don't meet a prevalence, age, or trusted list criteria. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedExecutableType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderUntrustedExecutableType = value
    }
}
// SetDefenderUntrustedUSBProcess sets the defenderUntrustedUSBProcess property value. Value indicating response to untrusted and unsigned processes that run from USB. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedUSBProcess(value *DefenderProtectionType)() {
    if m != nil {
        m.defenderUntrustedUSBProcess = value
    }
}
// SetDefenderUntrustedUSBProcessType sets the defenderUntrustedUSBProcessType property value. Value indicating response to untrusted and unsigned processes that run from USB. Possible values are: userDefined, block, auditMode, warn, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedUSBProcessType(value *DefenderAttackSurfaceType)() {
    if m != nil {
        m.defenderUntrustedUSBProcessType = value
    }
}
// SetDeviceGuardEnableSecureBootWithDMA sets the deviceGuardEnableSecureBootWithDMA property value. This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA. Specifies whether Platform Security Level is enabled at next reboot.
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardEnableSecureBootWithDMA(value *bool)() {
    if m != nil {
        m.deviceGuardEnableSecureBootWithDMA = value
    }
}
// SetDeviceGuardEnableVirtualizationBasedSecurity sets the deviceGuardEnableVirtualizationBasedSecurity property value. Turns On Virtualization Based Security(VBS).
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardEnableVirtualizationBasedSecurity(value *bool)() {
    if m != nil {
        m.deviceGuardEnableVirtualizationBasedSecurity = value
    }
}
// SetDeviceGuardLaunchSystemGuard sets the deviceGuardLaunchSystemGuard property value. Allows the IT admin to configure the launch of System Guard. Possible values are: notConfigured, enabled, disabled.
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardLaunchSystemGuard(value *Enablement)() {
    if m != nil {
        m.deviceGuardLaunchSystemGuard = value
    }
}
// SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings sets the deviceGuardLocalSystemAuthorityCredentialGuardSettings property value. Turn on Credential Guard when Platform Security Level with Secure Boot and Virtualization Based Security are both enabled. Possible values are: notConfigured, enableWithUEFILock, enableWithoutUEFILock, disable.
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings(value *DeviceGuardLocalSystemAuthorityCredentialGuardType)() {
    if m != nil {
        m.deviceGuardLocalSystemAuthorityCredentialGuardSettings = value
    }
}
// SetDeviceGuardSecureBootWithDMA sets the deviceGuardSecureBootWithDMA property value. Specifies whether Platform Security Level is enabled at next reboot. Possible values are: notConfigured, withoutDMA, withDMA.
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardSecureBootWithDMA(value *SecureBootWithDMAType)() {
    if m != nil {
        m.deviceGuardSecureBootWithDMA = value
    }
}
// SetDmaGuardDeviceEnumerationPolicy sets the dmaGuardDeviceEnumerationPolicy property value. This policy is intended to provide additional security against external DMA capable devices. It allows for more control over the enumeration of external DMA capable devices incompatible with DMA Remapping/device memory isolation and sandboxing. This policy only takes effect when Kernel DMA Protection is supported and enabled by the system firmware. Kernel DMA Protection is a platform feature that cannot be controlled via policy or by end user. It has to be supported by the system at the time of manufacturing. To check if the system supports Kernel DMA Protection, please check the Kernel DMA Protection field in the Summary page of MSINFO32.exe. Possible values are: deviceDefault, blockAll, allowAll.
func (m *Windows10EndpointProtectionConfiguration) SetDmaGuardDeviceEnumerationPolicy(value *DmaGuardDeviceEnumerationPolicyType)() {
    if m != nil {
        m.dmaGuardDeviceEnumerationPolicy = value
    }
}
// SetFirewallBlockStatefulFTP sets the firewallBlockStatefulFTP property value. Blocks stateful FTP connections to the device
func (m *Windows10EndpointProtectionConfiguration) SetFirewallBlockStatefulFTP(value *bool)() {
    if m != nil {
        m.firewallBlockStatefulFTP = value
    }
}
// SetFirewallCertificateRevocationListCheckMethod sets the firewallCertificateRevocationListCheckMethod property value. Specify how the certificate revocation list is to be enforced. Possible values are: deviceDefault, none, attempt, require.
func (m *Windows10EndpointProtectionConfiguration) SetFirewallCertificateRevocationListCheckMethod(value *FirewallCertificateRevocationListCheckMethodType)() {
    if m != nil {
        m.firewallCertificateRevocationListCheckMethod = value
    }
}
// SetFirewallIdleTimeoutForSecurityAssociationInSeconds sets the firewallIdleTimeoutForSecurityAssociationInSeconds property value. Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIdleTimeoutForSecurityAssociationInSeconds(value *int32)() {
    if m != nil {
        m.firewallIdleTimeoutForSecurityAssociationInSeconds = value
    }
}
// SetFirewallIPSecExemptionsAllowDHCP sets the firewallIPSecExemptionsAllowDHCP property value. Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowDHCP(value *bool)() {
    if m != nil {
        m.firewallIPSecExemptionsAllowDHCP = value
    }
}
// SetFirewallIPSecExemptionsAllowICMP sets the firewallIPSecExemptionsAllowICMP property value. Configures IPSec exemptions to allow ICMP
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowICMP(value *bool)() {
    if m != nil {
        m.firewallIPSecExemptionsAllowICMP = value
    }
}
// SetFirewallIPSecExemptionsAllowNeighborDiscovery sets the firewallIPSecExemptionsAllowNeighborDiscovery property value. Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowNeighborDiscovery(value *bool)() {
    if m != nil {
        m.firewallIPSecExemptionsAllowNeighborDiscovery = value
    }
}
// SetFirewallIPSecExemptionsAllowRouterDiscovery sets the firewallIPSecExemptionsAllowRouterDiscovery property value. Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowRouterDiscovery(value *bool)() {
    if m != nil {
        m.firewallIPSecExemptionsAllowRouterDiscovery = value
    }
}
// SetFirewallIPSecExemptionsNone sets the firewallIPSecExemptionsNone property value. Configures IPSec exemptions to no exemptions
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsNone(value *bool)() {
    if m != nil {
        m.firewallIPSecExemptionsNone = value
    }
}
// SetFirewallMergeKeyingModuleSettings sets the firewallMergeKeyingModuleSettings property value. If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
func (m *Windows10EndpointProtectionConfiguration) SetFirewallMergeKeyingModuleSettings(value *bool)() {
    if m != nil {
        m.firewallMergeKeyingModuleSettings = value
    }
}
// SetFirewallPacketQueueingMethod sets the firewallPacketQueueingMethod property value. Configures how packet queueing should be applied in the tunnel gateway scenario. Possible values are: deviceDefault, disabled, queueInbound, queueOutbound, queueBoth.
func (m *Windows10EndpointProtectionConfiguration) SetFirewallPacketQueueingMethod(value *FirewallPacketQueueingMethodType)() {
    if m != nil {
        m.firewallPacketQueueingMethod = value
    }
}
// SetFirewallPreSharedKeyEncodingMethod sets the firewallPreSharedKeyEncodingMethod property value. Select the preshared key encoding to be used. Possible values are: deviceDefault, none, utF8.
func (m *Windows10EndpointProtectionConfiguration) SetFirewallPreSharedKeyEncodingMethod(value *FirewallPreSharedKeyEncodingMethodType)() {
    if m != nil {
        m.firewallPreSharedKeyEncodingMethod = value
    }
}
// SetFirewallProfileDomain sets the firewallProfileDomain property value. Configures the firewall profile settings for domain networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfileDomain(value WindowsFirewallNetworkProfileable)() {
    if m != nil {
        m.firewallProfileDomain = value
    }
}
// SetFirewallProfilePrivate sets the firewallProfilePrivate property value. Configures the firewall profile settings for private networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfilePrivate(value WindowsFirewallNetworkProfileable)() {
    if m != nil {
        m.firewallProfilePrivate = value
    }
}
// SetFirewallProfilePublic sets the firewallProfilePublic property value. Configures the firewall profile settings for public networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfilePublic(value WindowsFirewallNetworkProfileable)() {
    if m != nil {
        m.firewallProfilePublic = value
    }
}
// SetFirewallRules sets the firewallRules property value. Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
func (m *Windows10EndpointProtectionConfiguration) SetFirewallRules(value []WindowsFirewallRuleable)() {
    if m != nil {
        m.firewallRules = value
    }
}
// SetLanManagerAuthenticationLevel sets the lanManagerAuthenticationLevel property value. This security setting determines which challenge/response authentication protocol is used for network logons. Possible values are: lmAndNltm, lmNtlmAndNtlmV2, lmAndNtlmOnly, lmAndNtlmV2, lmNtlmV2AndNotLm, lmNtlmV2AndNotLmOrNtm.
func (m *Windows10EndpointProtectionConfiguration) SetLanManagerAuthenticationLevel(value *LanManagerAuthenticationLevel)() {
    if m != nil {
        m.lanManagerAuthenticationLevel = value
    }
}
// SetLanManagerWorkstationDisableInsecureGuestLogons sets the lanManagerWorkstationDisableInsecureGuestLogons property value. If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure guest logons.
func (m *Windows10EndpointProtectionConfiguration) SetLanManagerWorkstationDisableInsecureGuestLogons(value *bool)() {
    if m != nil {
        m.lanManagerWorkstationDisableInsecureGuestLogons = value
    }
}
// SetLocalSecurityOptionsAdministratorAccountName sets the localSecurityOptionsAdministratorAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAdministratorAccountName(value *string)() {
    if m != nil {
        m.localSecurityOptionsAdministratorAccountName = value
    }
}
// SetLocalSecurityOptionsAdministratorElevationPromptBehavior sets the localSecurityOptionsAdministratorElevationPromptBehavior property value. Define the behavior of the elevation prompt for admins in Admin Approval Mode. Possible values are: notConfigured, elevateWithoutPrompting, promptForCredentialsOnTheSecureDesktop, promptForConsentOnTheSecureDesktop, promptForCredentials, promptForConsent, promptForConsentForNonWindowsBinaries.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAdministratorElevationPromptBehavior(value *LocalSecurityOptionsAdministratorElevationPromptBehaviorType)() {
    if m != nil {
        m.localSecurityOptionsAdministratorElevationPromptBehavior = value
    }
}
// SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares sets the localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares property value. This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating the names of domain accounts and network shares.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares(value *bool)() {
    if m != nil {
        m.localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares = value
    }
}
// SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests sets the localSecurityOptionsAllowPKU2UAuthenticationRequests property value. Block PKU2U authentication requests to this device to use online identities.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests(value *bool)() {
    if m != nil {
        m.localSecurityOptionsAllowPKU2UAuthenticationRequests = value
    }
}
// SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager sets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager property value. Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote calls to the SAM.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager(value *string)() {
    if m != nil {
        m.localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager = value
    }
}
// SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool sets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool property value. UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool(value *bool)() {
    if m != nil {
        m.localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool = value
    }
}
// SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn sets the localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn property value. This security setting determines whether a computer can be shut down without having to log on to Windows.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn(value *bool)() {
    if m != nil {
        m.localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn = value
    }
}
// SetLocalSecurityOptionsAllowUIAccessApplicationElevation sets the localSecurityOptionsAllowUIAccessApplicationElevation property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUIAccessApplicationElevation(value *bool)() {
    if m != nil {
        m.localSecurityOptionsAllowUIAccessApplicationElevation = value
    }
}
// SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations sets the localSecurityOptionsAllowUIAccessApplicationsForSecureLocations property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations(value *bool)() {
    if m != nil {
        m.localSecurityOptionsAllowUIAccessApplicationsForSecureLocations = value
    }
}
// SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon sets the localSecurityOptionsAllowUndockWithoutHavingToLogon property value. Prevent a portable computer from being undocked without having to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon(value *bool)() {
    if m != nil {
        m.localSecurityOptionsAllowUndockWithoutHavingToLogon = value
    }
}
// SetLocalSecurityOptionsBlockMicrosoftAccounts sets the localSecurityOptionsBlockMicrosoftAccounts property value. Prevent users from adding new Microsoft accounts to this computer.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockMicrosoftAccounts(value *bool)() {
    if m != nil {
        m.localSecurityOptionsBlockMicrosoftAccounts = value
    }
}
// SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword sets the localSecurityOptionsBlockRemoteLogonWithBlankPassword property value. Enable Local accounts that are not password protected to log on from locations other than the physical device.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword(value *bool)() {
    if m != nil {
        m.localSecurityOptionsBlockRemoteLogonWithBlankPassword = value
    }
}
// SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess sets the localSecurityOptionsBlockRemoteOpticalDriveAccess property value. Enabling this settings allows only interactively logged on user to access CD-ROM media.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess(value *bool)() {
    if m != nil {
        m.localSecurityOptionsBlockRemoteOpticalDriveAccess = value
    }
}
// SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers sets the localSecurityOptionsBlockUsersInstallingPrinterDrivers property value. Restrict installing printer drivers as part of connecting to a shared printer to admins only.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers(value *bool)() {
    if m != nil {
        m.localSecurityOptionsBlockUsersInstallingPrinterDrivers = value
    }
}
// SetLocalSecurityOptionsClearVirtualMemoryPageFile sets the localSecurityOptionsClearVirtualMemoryPageFile property value. This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClearVirtualMemoryPageFile(value *bool)() {
    if m != nil {
        m.localSecurityOptionsClearVirtualMemoryPageFile = value
    }
}
// SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways sets the localSecurityOptionsClientDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB client component.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways(value *bool)() {
    if m != nil {
        m.localSecurityOptionsClientDigitallySignCommunicationsAlways = value
    }
}
// SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers sets the localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers property value. If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords to non-Microsoft SMB servers that do not support password encryption during authentication.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers(value *bool)() {
    if m != nil {
        m.localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers = value
    }
}
// SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation sets the localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation property value. App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation = value
    }
}
// SetLocalSecurityOptionsDisableAdministratorAccount sets the localSecurityOptionsDisableAdministratorAccount property value. Determines whether the Local Administrator account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableAdministratorAccount(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDisableAdministratorAccount = value
    }
}
// SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees sets the localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees property value. This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees = value
    }
}
// SetLocalSecurityOptionsDisableGuestAccount sets the localSecurityOptionsDisableGuestAccount property value. Determines if the Guest account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableGuestAccount(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDisableGuestAccount = value
    }
}
// SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways sets the localSecurityOptionsDisableServerDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB server component.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDisableServerDigitallySignCommunicationsAlways = value
    }
}
// SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees sets the localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees property value. This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request it.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees = value
    }
}
// SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts sets the localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts property value. This security setting determines what additional permissions will be granted for anonymous connections to the computer.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts = value
    }
}
// SetLocalSecurityOptionsDoNotRequireCtrlAltDel sets the localSecurityOptionsDoNotRequireCtrlAltDel property value. Require CTRL+ALT+DEL to be pressed before a user can log on.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotRequireCtrlAltDel(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDoNotRequireCtrlAltDel = value
    }
}
// SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange sets the localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange property value. This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new password is stored. It’s not stored by default.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange(value *bool)() {
    if m != nil {
        m.localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange = value
    }
}
// SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser sets the localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser property value. Define who is allowed to format and eject removable NTFS media. Possible values are: notConfigured, administrators, administratorsAndPowerUsers, administratorsAndInteractiveUsers.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(value *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType)() {
    if m != nil {
        m.localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = value
    }
}
// SetLocalSecurityOptionsGuestAccountName sets the localSecurityOptionsGuestAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsGuestAccountName(value *string)() {
    if m != nil {
        m.localSecurityOptionsGuestAccountName = value
    }
}
// SetLocalSecurityOptionsHideLastSignedInUser sets the localSecurityOptionsHideLastSignedInUser property value. Do not display the username of the last person who signed in on this device.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsHideLastSignedInUser(value *bool)() {
    if m != nil {
        m.localSecurityOptionsHideLastSignedInUser = value
    }
}
// SetLocalSecurityOptionsHideUsernameAtSignIn sets the localSecurityOptionsHideUsernameAtSignIn property value. Do not display the username of the person signing in to this device after credentials are entered and before the device’s desktop is shown.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsHideUsernameAtSignIn(value *bool)() {
    if m != nil {
        m.localSecurityOptionsHideUsernameAtSignIn = value
    }
}
// SetLocalSecurityOptionsInformationDisplayedOnLockScreen sets the localSecurityOptionsInformationDisplayedOnLockScreen property value. Configure the user information that is displayed when the session is locked. If not configured, user display name, domain and username are shown. Possible values are: notConfigured, administrators, administratorsAndPowerUsers, administratorsAndInteractiveUsers.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsInformationDisplayedOnLockScreen(value *LocalSecurityOptionsInformationDisplayedOnLockScreenType)() {
    if m != nil {
        m.localSecurityOptionsInformationDisplayedOnLockScreen = value
    }
}
// SetLocalSecurityOptionsInformationShownOnLockScreen sets the localSecurityOptionsInformationShownOnLockScreen property value. Configure the user information that is displayed when the session is locked. If not configured, user display name, domain and username are shown. Possible values are: notConfigured, userDisplayNameDomainUser, userDisplayNameOnly, doNotDisplayUser.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsInformationShownOnLockScreen(value *LocalSecurityOptionsInformationShownOnLockScreenType)() {
    if m != nil {
        m.localSecurityOptionsInformationShownOnLockScreen = value
    }
}
// SetLocalSecurityOptionsLogOnMessageText sets the localSecurityOptionsLogOnMessageText property value. Set message text for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsLogOnMessageText(value *string)() {
    if m != nil {
        m.localSecurityOptionsLogOnMessageText = value
    }
}
// SetLocalSecurityOptionsLogOnMessageTitle sets the localSecurityOptionsLogOnMessageTitle property value. Set message title for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsLogOnMessageTitle(value *string)() {
    if m != nil {
        m.localSecurityOptionsLogOnMessageTitle = value
    }
}
// SetLocalSecurityOptionsMachineInactivityLimit sets the localSecurityOptionsMachineInactivityLimit property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMachineInactivityLimit(value *int32)() {
    if m != nil {
        m.localSecurityOptionsMachineInactivityLimit = value
    }
}
// SetLocalSecurityOptionsMachineInactivityLimitInMinutes sets the localSecurityOptionsMachineInactivityLimitInMinutes property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMachineInactivityLimitInMinutes(value *int32)() {
    if m != nil {
        m.localSecurityOptionsMachineInactivityLimitInMinutes = value
    }
}
// SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients sets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients property value. This security setting allows a client to require the negotiation of 128-bit encryption and/or NTLMv2 session security. Possible values are: none, requireNtmlV2SessionSecurity, require128BitEncryption, ntlmV2And128BitEncryption.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(value *LocalSecurityOptionsMinimumSessionSecurity)() {
    if m != nil {
        m.localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients = value
    }
}
// SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers sets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers property value. This security setting allows a server to require the negotiation of 128-bit encryption and/or NTLMv2 session security. Possible values are: none, requireNtmlV2SessionSecurity, require128BitEncryption, ntlmV2And128BitEncryption.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(value *LocalSecurityOptionsMinimumSessionSecurity)() {
    if m != nil {
        m.localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers = value
    }
}
// SetLocalSecurityOptionsOnlyElevateSignedExecutables sets the localSecurityOptionsOnlyElevateSignedExecutables property value. Enforce PKI certification path validation for a given executable file before it is permitted to run.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsOnlyElevateSignedExecutables(value *bool)() {
    if m != nil {
        m.localSecurityOptionsOnlyElevateSignedExecutables = value
    }
}
// SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares sets the localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares property value. By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that can be accessed anonymously and Shares that can be accessed anonymously
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares(value *bool)() {
    if m != nil {
        m.localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares = value
    }
}
// SetLocalSecurityOptionsSmartCardRemovalBehavior sets the localSecurityOptionsSmartCardRemovalBehavior property value. This security setting determines what happens when the smart card for a logged-on user is removed from the smart card reader. Possible values are: noAction, lockWorkstation, forceLogoff, disconnectRemoteDesktopSession.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsSmartCardRemovalBehavior(value *LocalSecurityOptionsSmartCardRemovalBehaviorType)() {
    if m != nil {
        m.localSecurityOptionsSmartCardRemovalBehavior = value
    }
}
// SetLocalSecurityOptionsStandardUserElevationPromptBehavior sets the localSecurityOptionsStandardUserElevationPromptBehavior property value. Define the behavior of the elevation prompt for standard users. Possible values are: notConfigured, automaticallyDenyElevationRequests, promptForCredentialsOnTheSecureDesktop, promptForCredentials.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsStandardUserElevationPromptBehavior(value *LocalSecurityOptionsStandardUserElevationPromptBehaviorType)() {
    if m != nil {
        m.localSecurityOptionsStandardUserElevationPromptBehavior = value
    }
}
// SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation sets the localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation property value. Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior policy settings for admins and standard users are used.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation(value *bool)() {
    if m != nil {
        m.localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation = value
    }
}
// SetLocalSecurityOptionsUseAdminApprovalMode sets the localSecurityOptionsUseAdminApprovalMode property value. Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin privileges.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsUseAdminApprovalMode(value *bool)() {
    if m != nil {
        m.localSecurityOptionsUseAdminApprovalMode = value
    }
}
// SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators sets the localSecurityOptionsUseAdminApprovalModeForAdministrators property value. Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators(value *bool)() {
    if m != nil {
        m.localSecurityOptionsUseAdminApprovalModeForAdministrators = value
    }
}
// SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations sets the localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations property value. Virtualize file and registry write failures to per user locations
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations(value *bool)() {
    if m != nil {
        m.localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations = value
    }
}
// SetSmartScreenBlockOverrideForFiles sets the smartScreenBlockOverrideForFiles property value. Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
func (m *Windows10EndpointProtectionConfiguration) SetSmartScreenBlockOverrideForFiles(value *bool)() {
    if m != nil {
        m.smartScreenBlockOverrideForFiles = value
    }
}
// SetSmartScreenEnableInShell sets the smartScreenEnableInShell property value. Allows IT Admins to configure SmartScreen for Windows.
func (m *Windows10EndpointProtectionConfiguration) SetSmartScreenEnableInShell(value *bool)() {
    if m != nil {
        m.smartScreenEnableInShell = value
    }
}
// SetUserRightsAccessCredentialManagerAsTrustedCaller sets the userRightsAccessCredentialManagerAsTrustedCaller property value. This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if this privilege is given to other entities. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsAccessCredentialManagerAsTrustedCaller(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsAccessCredentialManagerAsTrustedCaller = value
    }
}
// SetUserRightsActAsPartOfTheOperatingSystem sets the userRightsActAsPartOfTheOperatingSystem property value. This user right allows a process to impersonate any user without authentication. The process can therefore gain access to the same local resources as that user. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsActAsPartOfTheOperatingSystem(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsActAsPartOfTheOperatingSystem = value
    }
}
// SetUserRightsAllowAccessFromNetwork sets the userRightsAllowAccessFromNetwork property value. This user right determines which users and groups are allowed to connect to the computer over the network. State Allowed is supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsAllowAccessFromNetwork(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsAllowAccessFromNetwork = value
    }
}
// SetUserRightsBackupData sets the userRightsBackupData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when backing up files and directories. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsBackupData(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsBackupData = value
    }
}
// SetUserRightsBlockAccessFromNetwork sets the userRightsBlockAccessFromNetwork property value. This user right determines which users and groups are block from connecting to the computer over the network. State Block is supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsBlockAccessFromNetwork(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsBlockAccessFromNetwork = value
    }
}
// SetUserRightsChangeSystemTime sets the userRightsChangeSystemTime property value. This user right determines which users and groups can change the time and date on the internal clock of the computer. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsChangeSystemTime(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsChangeSystemTime = value
    }
}
// SetUserRightsCreateGlobalObjects sets the userRightsCreateGlobalObjects property value. This security setting determines whether users can create global objects that are available to all sessions. Users who can create global objects could affect processes that run under other users' sessions, which could lead to application failure or data corruption. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateGlobalObjects(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsCreateGlobalObjects = value
    }
}
// SetUserRightsCreatePageFile sets the userRightsCreatePageFile property value. This user right determines which users and groups can call an internal API to create and change the size of a page file. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreatePageFile(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsCreatePageFile = value
    }
}
// SetUserRightsCreatePermanentSharedObjects sets the userRightsCreatePermanentSharedObjects property value. This user right determines which accounts can be used by processes to create a directory object using the object manager. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreatePermanentSharedObjects(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsCreatePermanentSharedObjects = value
    }
}
// SetUserRightsCreateSymbolicLinks sets the userRightsCreateSymbolicLinks property value. This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateSymbolicLinks(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsCreateSymbolicLinks = value
    }
}
// SetUserRightsCreateToken sets the userRightsCreateToken property value. This user right determines which users/groups can be used by processes to create a token that can then be used to get access to any local resources when the process uses an internal API to create an access token. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateToken(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsCreateToken = value
    }
}
// SetUserRightsDebugPrograms sets the userRightsDebugPrograms property value. This user right determines which users can attach a debugger to any process or to the kernel. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDebugPrograms(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsDebugPrograms = value
    }
}
// SetUserRightsDelegation sets the userRightsDelegation property value. This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDelegation(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsDelegation = value
    }
}
// SetUserRightsDenyLocalLogOn sets the userRightsDenyLocalLogOn property value. This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDenyLocalLogOn(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsDenyLocalLogOn = value
    }
}
// SetUserRightsGenerateSecurityAudits sets the userRightsGenerateSecurityAudits property value. This user right determines which accounts can be used by a process to add entries to the security log. The security log is used to trace unauthorized system access.  Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsGenerateSecurityAudits(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsGenerateSecurityAudits = value
    }
}
// SetUserRightsImpersonateClient sets the userRightsImpersonateClient property value. Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a service that they have created and then impersonating that client, which can elevate the unauthorized user's permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsImpersonateClient(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsImpersonateClient = value
    }
}
// SetUserRightsIncreaseSchedulingPriority sets the userRightsIncreaseSchedulingPriority property value. This user right determines which accounts can use a process with Write Property access to another process to increase the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsIncreaseSchedulingPriority(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsIncreaseSchedulingPriority = value
    }
}
// SetUserRightsLoadUnloadDrivers sets the userRightsLoadUnloadDrivers property value. This user right determines which users can dynamically load and unload device drivers or other code in to kernel mode. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLoadUnloadDrivers(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsLoadUnloadDrivers = value
    }
}
// SetUserRightsLocalLogOn sets the userRightsLocalLogOn property value. This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLocalLogOn(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsLocalLogOn = value
    }
}
// SetUserRightsLockMemory sets the userRightsLockMemory property value. This user right determines which accounts can use a process to keep data in physical memory, which prevents the system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLockMemory(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsLockMemory = value
    }
}
// SetUserRightsManageAuditingAndSecurityLogs sets the userRightsManageAuditingAndSecurityLogs property value. This user right determines which users can specify object access auditing options for individual resources, such as files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsManageAuditingAndSecurityLogs(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsManageAuditingAndSecurityLogs = value
    }
}
// SetUserRightsManageVolumes sets the userRightsManageVolumes property value. This user right determines which users and groups can run maintenance tasks on a volume, such as remote defragmentation. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsManageVolumes(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsManageVolumes = value
    }
}
// SetUserRightsModifyFirmwareEnvironment sets the userRightsModifyFirmwareEnvironment property value. This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsModifyFirmwareEnvironment(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsModifyFirmwareEnvironment = value
    }
}
// SetUserRightsModifyObjectLabels sets the userRightsModifyObjectLabels property value. This user right determines which user accounts can modify the integrity label of objects, such as files, registry keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsModifyObjectLabels(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsModifyObjectLabels = value
    }
}
// SetUserRightsProfileSingleProcess sets the userRightsProfileSingleProcess property value. This user right determines which users can use performance monitoring tools to monitor the performance of system processes. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsProfileSingleProcess(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsProfileSingleProcess = value
    }
}
// SetUserRightsRemoteDesktopServicesLogOn sets the userRightsRemoteDesktopServicesLogOn property value. This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client. Only states NotConfigured and Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRemoteDesktopServicesLogOn(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsRemoteDesktopServicesLogOn = value
    }
}
// SetUserRightsRemoteShutdown sets the userRightsRemoteShutdown property value. This user right determines which users are allowed to shut down a computer from a remote location on the network. Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRemoteShutdown(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsRemoteShutdown = value
    }
}
// SetUserRightsRestoreData sets the userRightsRestoreData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when restoring backed up files and directories, and determines which users can set any valid security principal as the owner of an object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRestoreData(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsRestoreData = value
    }
}
// SetUserRightsTakeOwnership sets the userRightsTakeOwnership property value. This user right determines which users can take ownership of any securable object in the system, including Active Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsTakeOwnership(value DeviceManagementUserRightsSettingable)() {
    if m != nil {
        m.userRightsTakeOwnership = value
    }
}
// SetWindowsDefenderTamperProtection sets the windowsDefenderTamperProtection property value. Configure windows defender TamperProtection settings. Possible values are: notConfigured, enable, disable.
func (m *Windows10EndpointProtectionConfiguration) SetWindowsDefenderTamperProtection(value *WindowsDefenderTamperProtectionOptions)() {
    if m != nil {
        m.windowsDefenderTamperProtection = value
    }
}
// SetXboxServicesAccessoryManagementServiceStartupMode sets the xboxServicesAccessoryManagementServiceStartupMode property value. This setting determines whether the Accessory management service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesAccessoryManagementServiceStartupMode(value *ServiceStartType)() {
    if m != nil {
        m.xboxServicesAccessoryManagementServiceStartupMode = value
    }
}
// SetXboxServicesEnableXboxGameSaveTask sets the xboxServicesEnableXboxGameSaveTask property value. This setting determines whether xbox game save is enabled (1) or disabled (0).
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesEnableXboxGameSaveTask(value *bool)() {
    if m != nil {
        m.xboxServicesEnableXboxGameSaveTask = value
    }
}
// SetXboxServicesLiveAuthManagerServiceStartupMode sets the xboxServicesLiveAuthManagerServiceStartupMode property value. This setting determines whether Live Auth Manager service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveAuthManagerServiceStartupMode(value *ServiceStartType)() {
    if m != nil {
        m.xboxServicesLiveAuthManagerServiceStartupMode = value
    }
}
// SetXboxServicesLiveGameSaveServiceStartupMode sets the xboxServicesLiveGameSaveServiceStartupMode property value. This setting determines whether Live Game save service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveGameSaveServiceStartupMode(value *ServiceStartType)() {
    if m != nil {
        m.xboxServicesLiveGameSaveServiceStartupMode = value
    }
}
// SetXboxServicesLiveNetworkingServiceStartupMode sets the xboxServicesLiveNetworkingServiceStartupMode property value. This setting determines whether Networking service's start type is Automatic(2), Manual(3), Disabled(4). Default: Manual. Possible values are: manual, automatic, disabled.
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveNetworkingServiceStartupMode(value *ServiceStartType)() {
    if m != nil {
        m.xboxServicesLiveNetworkingServiceStartupMode = value
    }
}
