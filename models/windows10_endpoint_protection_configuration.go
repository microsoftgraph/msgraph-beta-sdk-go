package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EndpointProtectionConfiguration this topic provides descriptions of the declared methods, properties and relationships exposed by the Windows10EndpointProtectionConfiguration resource.
type Windows10EndpointProtectionConfiguration struct {
    DeviceConfiguration
}
// NewWindows10EndpointProtectionConfiguration instantiates a new Windows10EndpointProtectionConfiguration and sets the default values.
func NewWindows10EndpointProtectionConfiguration()(*Windows10EndpointProtectionConfiguration) {
    m := &Windows10EndpointProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10EndpointProtectionConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10EndpointProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindows10EndpointProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EndpointProtectionConfiguration(), nil
}
// GetApplicationGuardAllowCameraMicrophoneRedirection gets the applicationGuardAllowCameraMicrophoneRedirection property value. Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and microphone.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowCameraMicrophoneRedirection()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardAllowCameraMicrophoneRedirection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardAllowFileSaveOnHost gets the applicationGuardAllowFileSaveOnHost property value. Allow users to download files from Edge in the application guard container and save them on the host file system
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowFileSaveOnHost()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardAllowFileSaveOnHost")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardAllowPersistence gets the applicationGuardAllowPersistence property value. Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPersistence()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardAllowPersistence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardAllowPrintToLocalPrinters gets the applicationGuardAllowPrintToLocalPrinters property value. Allow printing to Local Printers from Container
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToLocalPrinters()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardAllowPrintToLocalPrinters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardAllowPrintToNetworkPrinters gets the applicationGuardAllowPrintToNetworkPrinters property value. Allow printing to Network Printers from Container
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToNetworkPrinters()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardAllowPrintToNetworkPrinters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardAllowPrintToPDF gets the applicationGuardAllowPrintToPDF property value. Allow printing to PDF from Container
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToPDF()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardAllowPrintToPDF")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardAllowPrintToXPS gets the applicationGuardAllowPrintToXPS property value. Allow printing to XPS from Container
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToXPS()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardAllowPrintToXPS")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardAllowVirtualGPU gets the applicationGuardAllowVirtualGPU property value. Allow application guard to use virtual GPU
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowVirtualGPU()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardAllowVirtualGPU")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardBlockClipboardSharing gets the applicationGuardBlockClipboardSharing property value. Possible values for applicationGuardBlockClipboardSharingType
// returns a *ApplicationGuardBlockClipboardSharingType when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockClipboardSharing()(*ApplicationGuardBlockClipboardSharingType) {
    val, err := m.GetBackingStore().Get("applicationGuardBlockClipboardSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApplicationGuardBlockClipboardSharingType)
    }
    return nil
}
// GetApplicationGuardBlockFileTransfer gets the applicationGuardBlockFileTransfer property value. Possible values for applicationGuardBlockFileTransfer
// returns a *ApplicationGuardBlockFileTransferType when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockFileTransfer()(*ApplicationGuardBlockFileTransferType) {
    val, err := m.GetBackingStore().Get("applicationGuardBlockFileTransfer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApplicationGuardBlockFileTransferType)
    }
    return nil
}
// GetApplicationGuardBlockNonEnterpriseContent gets the applicationGuardBlockNonEnterpriseContent property value. Block enterprise sites to load non-enterprise content, such as third party plug-ins
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockNonEnterpriseContent()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardBlockNonEnterpriseContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardCertificateThumbprints gets the applicationGuardCertificateThumbprints property value. Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
// returns a []string when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardCertificateThumbprints()([]string) {
    val, err := m.GetBackingStore().Get("applicationGuardCertificateThumbprints")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetApplicationGuardEnabled gets the applicationGuardEnabled property value. Enable Windows Defender Application Guard
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplicationGuardEnabledOptions gets the applicationGuardEnabledOptions property value. Possible values for ApplicationGuardEnabledOptions
// returns a *ApplicationGuardEnabledOptions when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardEnabledOptions()(*ApplicationGuardEnabledOptions) {
    val, err := m.GetBackingStore().Get("applicationGuardEnabledOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApplicationGuardEnabledOptions)
    }
    return nil
}
// GetApplicationGuardForceAuditing gets the applicationGuardForceAuditing property value. Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardForceAuditing()(*bool) {
    val, err := m.GetBackingStore().Get("applicationGuardForceAuditing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAppLockerApplicationControl gets the appLockerApplicationControl property value. Possible values of AppLocker Application Control Types
// returns a *AppLockerApplicationControlType when successful
func (m *Windows10EndpointProtectionConfiguration) GetAppLockerApplicationControl()(*AppLockerApplicationControlType) {
    val, err := m.GetBackingStore().Get("appLockerApplicationControl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppLockerApplicationControlType)
    }
    return nil
}
// GetBitLockerAllowStandardUserEncryption gets the bitLockerAllowStandardUserEncryption property value. Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerAllowStandardUserEncryption()(*bool) {
    val, err := m.GetBackingStore().Get("bitLockerAllowStandardUserEncryption")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBitLockerDisableWarningForOtherDiskEncryption gets the bitLockerDisableWarningForOtherDiskEncryption property value. Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerDisableWarningForOtherDiskEncryption()(*bool) {
    val, err := m.GetBackingStore().Get("bitLockerDisableWarningForOtherDiskEncryption")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBitLockerEnableStorageCardEncryptionOnMobile gets the bitLockerEnableStorageCardEncryptionOnMobile property value. Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerEnableStorageCardEncryptionOnMobile()(*bool) {
    val, err := m.GetBackingStore().Get("bitLockerEnableStorageCardEncryptionOnMobile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBitLockerEncryptDevice gets the bitLockerEncryptDevice property value. Allows the admin to require encryption to be turned on using BitLocker.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerEncryptDevice()(*bool) {
    val, err := m.GetBackingStore().Get("bitLockerEncryptDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBitLockerFixedDrivePolicy gets the bitLockerFixedDrivePolicy property value. BitLocker Fixed Drive Policy.
// returns a BitLockerFixedDrivePolicyable when successful
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerFixedDrivePolicy()(BitLockerFixedDrivePolicyable) {
    val, err := m.GetBackingStore().Get("bitLockerFixedDrivePolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BitLockerFixedDrivePolicyable)
    }
    return nil
}
// GetBitLockerRecoveryPasswordRotation gets the bitLockerRecoveryPasswordRotation property value. BitLocker recovery password rotation type
// returns a *BitLockerRecoveryPasswordRotationType when successful
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerRecoveryPasswordRotation()(*BitLockerRecoveryPasswordRotationType) {
    val, err := m.GetBackingStore().Get("bitLockerRecoveryPasswordRotation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*BitLockerRecoveryPasswordRotationType)
    }
    return nil
}
// GetBitLockerRemovableDrivePolicy gets the bitLockerRemovableDrivePolicy property value. BitLocker Removable Drive Policy.
// returns a BitLockerRemovableDrivePolicyable when successful
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerRemovableDrivePolicy()(BitLockerRemovableDrivePolicyable) {
    val, err := m.GetBackingStore().Get("bitLockerRemovableDrivePolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BitLockerRemovableDrivePolicyable)
    }
    return nil
}
// GetBitLockerSystemDrivePolicy gets the bitLockerSystemDrivePolicy property value. BitLocker System Drive Policy.
// returns a BitLockerSystemDrivePolicyable when successful
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerSystemDrivePolicy()(BitLockerSystemDrivePolicyable) {
    val, err := m.GetBackingStore().Get("bitLockerSystemDrivePolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BitLockerSystemDrivePolicyable)
    }
    return nil
}
// GetDefenderAdditionalGuardedFolders gets the defenderAdditionalGuardedFolders property value. List of folder paths to be added to the list of protected folders
// returns a []string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdditionalGuardedFolders()([]string) {
    val, err := m.GetBackingStore().Get("defenderAdditionalGuardedFolders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDefenderAdobeReaderLaunchChildProcess gets the defenderAdobeReaderLaunchChildProcess property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdobeReaderLaunchChildProcess()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderAdobeReaderLaunchChildProcess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderAdvancedRansomewareProtectionType gets the defenderAdvancedRansomewareProtectionType property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdvancedRansomewareProtectionType()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderAdvancedRansomewareProtectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderAllowBehaviorMonitoring gets the defenderAllowBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowBehaviorMonitoring()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowBehaviorMonitoring")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowCloudProtection gets the defenderAllowCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowCloudProtection()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowCloudProtection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowEndUserAccess gets the defenderAllowEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowEndUserAccess()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowEndUserAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowIntrusionPreventionSystem gets the defenderAllowIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowIntrusionPreventionSystem()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowIntrusionPreventionSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowOnAccessProtection gets the defenderAllowOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowOnAccessProtection()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowOnAccessProtection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowRealTimeMonitoring gets the defenderAllowRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowRealTimeMonitoring()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowRealTimeMonitoring")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowScanArchiveFiles gets the defenderAllowScanArchiveFiles property value. Allows or disallows scanning of archives.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanArchiveFiles()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowScanArchiveFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowScanDownloads gets the defenderAllowScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanDownloads()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowScanDownloads")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowScanNetworkFiles gets the defenderAllowScanNetworkFiles property value. Allows or disallows a scanning of network files.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanNetworkFiles()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowScanNetworkFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowScanRemovableDrivesDuringFullScan gets the defenderAllowScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanRemovableDrivesDuringFullScan()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowScanRemovableDrivesDuringFullScan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAllowScanScriptsLoadedInInternetExplorer gets the defenderAllowScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAllowScanScriptsLoadedInInternetExplorer()(*bool) {
    val, err := m.GetBackingStore().Get("defenderAllowScanScriptsLoadedInInternetExplorer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderAttackSurfaceReductionExcludedPaths gets the defenderAttackSurfaceReductionExcludedPaths property value. List of exe files and folders to be excluded from attack surface reduction rules
// returns a []string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAttackSurfaceReductionExcludedPaths()([]string) {
    val, err := m.GetBackingStore().Get("defenderAttackSurfaceReductionExcludedPaths")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDefenderBlockEndUserAccess gets the defenderBlockEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderBlockEndUserAccess()(*bool) {
    val, err := m.GetBackingStore().Get("defenderBlockEndUserAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderBlockPersistenceThroughWmiType gets the defenderBlockPersistenceThroughWmiType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderBlockPersistenceThroughWmiType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderBlockPersistenceThroughWmiType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderCheckForSignaturesBeforeRunningScan gets the defenderCheckForSignaturesBeforeRunningScan property value. This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before running a scan.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCheckForSignaturesBeforeRunningScan()(*bool) {
    val, err := m.GetBackingStore().Get("defenderCheckForSignaturesBeforeRunningScan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderCloudBlockLevel gets the defenderCloudBlockLevel property value. Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS' setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
// returns a *DefenderCloudBlockLevelType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCloudBlockLevel()(*DefenderCloudBlockLevelType) {
    val, err := m.GetBackingStore().Get("defenderCloudBlockLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderCloudBlockLevelType)
    }
    return nil
}
// GetDefenderCloudExtendedTimeoutInSeconds gets the defenderCloudExtendedTimeoutInSeconds property value. Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
// returns a *int32 when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderCloudExtendedTimeoutInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("defenderCloudExtendedTimeoutInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDefenderDaysBeforeDeletingQuarantinedMalware gets the defenderDaysBeforeDeletingQuarantinedMalware property value. Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
// returns a *int32 when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDaysBeforeDeletingQuarantinedMalware()(*int32) {
    val, err := m.GetBackingStore().Get("defenderDaysBeforeDeletingQuarantinedMalware")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDefenderDetectedMalwareActions gets the defenderDetectedMalwareActions property value. Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
// returns a DefenderDetectedMalwareActionsable when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDetectedMalwareActions()(DefenderDetectedMalwareActionsable) {
    val, err := m.GetBackingStore().Get("defenderDetectedMalwareActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DefenderDetectedMalwareActionsable)
    }
    return nil
}
// GetDefenderDisableBehaviorMonitoring gets the defenderDisableBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableBehaviorMonitoring()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableBehaviorMonitoring")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableCatchupFullScan gets the defenderDisableCatchupFullScan property value. This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCatchupFullScan()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableCatchupFullScan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableCatchupQuickScan gets the defenderDisableCatchupQuickScan property value. This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCatchupQuickScan()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableCatchupQuickScan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableCloudProtection gets the defenderDisableCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableCloudProtection()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableCloudProtection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableIntrusionPreventionSystem gets the defenderDisableIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableIntrusionPreventionSystem()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableIntrusionPreventionSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableOnAccessProtection gets the defenderDisableOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableOnAccessProtection()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableOnAccessProtection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableRealTimeMonitoring gets the defenderDisableRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableRealTimeMonitoring()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableRealTimeMonitoring")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableScanArchiveFiles gets the defenderDisableScanArchiveFiles property value. Allows or disallows scanning of archives.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanArchiveFiles()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableScanArchiveFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableScanDownloads gets the defenderDisableScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanDownloads()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableScanDownloads")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableScanNetworkFiles gets the defenderDisableScanNetworkFiles property value. Allows or disallows a scanning of network files.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanNetworkFiles()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableScanNetworkFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableScanRemovableDrivesDuringFullScan gets the defenderDisableScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanRemovableDrivesDuringFullScan()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableScanRemovableDrivesDuringFullScan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderDisableScanScriptsLoadedInInternetExplorer gets the defenderDisableScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderDisableScanScriptsLoadedInInternetExplorer()(*bool) {
    val, err := m.GetBackingStore().Get("defenderDisableScanScriptsLoadedInInternetExplorer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderEmailContentExecution gets the defenderEmailContentExecution property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEmailContentExecution()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderEmailContentExecution")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderEmailContentExecutionType gets the defenderEmailContentExecutionType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEmailContentExecutionType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderEmailContentExecutionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderEnableLowCpuPriority gets the defenderEnableLowCpuPriority property value. This policy setting allows you to enable or disable low CPU priority for scheduled scans.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableLowCpuPriority()(*bool) {
    val, err := m.GetBackingStore().Get("defenderEnableLowCpuPriority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderEnableScanIncomingMail gets the defenderEnableScanIncomingMail property value. Allows or disallows scanning of email.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableScanIncomingMail()(*bool) {
    val, err := m.GetBackingStore().Get("defenderEnableScanIncomingMail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderEnableScanMappedNetworkDrivesDuringFullScan gets the defenderEnableScanMappedNetworkDrivesDuringFullScan property value. Allows or disallows a full scan of mapped network drives.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderEnableScanMappedNetworkDrivesDuringFullScan()(*bool) {
    val, err := m.GetBackingStore().Get("defenderEnableScanMappedNetworkDrivesDuringFullScan")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderExploitProtectionXml gets the defenderExploitProtectionXml property value. Xml content containing information regarding exploit protection details.
// returns a []byte when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXml()([]byte) {
    val, err := m.GetBackingStore().Get("defenderExploitProtectionXml")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetDefenderExploitProtectionXmlFileName gets the defenderExploitProtectionXmlFileName property value. Name of the file from which DefenderExploitProtectionXml was obtained.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXmlFileName()(*string) {
    val, err := m.GetBackingStore().Get("defenderExploitProtectionXmlFileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDefenderFileExtensionsToExclude gets the defenderFileExtensionsToExclude property value. File extensions to exclude from scans and real time protection.
// returns a []string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderFileExtensionsToExclude()([]string) {
    val, err := m.GetBackingStore().Get("defenderFileExtensionsToExclude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDefenderFilesAndFoldersToExclude gets the defenderFilesAndFoldersToExclude property value. Files and folder to exclude from scans and real time protection.
// returns a []string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderFilesAndFoldersToExclude()([]string) {
    val, err := m.GetBackingStore().Get("defenderFilesAndFoldersToExclude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDefenderGuardedFoldersAllowedAppPaths gets the defenderGuardedFoldersAllowedAppPaths property value. List of paths to exe that are allowed to access protected folders
// returns a []string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderGuardedFoldersAllowedAppPaths()([]string) {
    val, err := m.GetBackingStore().Get("defenderGuardedFoldersAllowedAppPaths")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDefenderGuardMyFoldersType gets the defenderGuardMyFoldersType property value. Possible values of Folder Protection
// returns a *FolderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderGuardMyFoldersType()(*FolderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderGuardMyFoldersType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FolderProtectionType)
    }
    return nil
}
// GetDefenderNetworkProtectionType gets the defenderNetworkProtectionType property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderNetworkProtectionType()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderNetworkProtectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderOfficeAppsExecutableContentCreationOrLaunch gets the defenderOfficeAppsExecutableContentCreationOrLaunch property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsExecutableContentCreationOrLaunch()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderOfficeAppsExecutableContentCreationOrLaunch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderOfficeAppsExecutableContentCreationOrLaunchType gets the defenderOfficeAppsExecutableContentCreationOrLaunchType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsExecutableContentCreationOrLaunchType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderOfficeAppsExecutableContentCreationOrLaunchType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderOfficeAppsLaunchChildProcess gets the defenderOfficeAppsLaunchChildProcess property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsLaunchChildProcess()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderOfficeAppsLaunchChildProcess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderOfficeAppsLaunchChildProcessType gets the defenderOfficeAppsLaunchChildProcessType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsLaunchChildProcessType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderOfficeAppsLaunchChildProcessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderOfficeAppsOtherProcessInjection gets the defenderOfficeAppsOtherProcessInjection property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsOtherProcessInjection()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderOfficeAppsOtherProcessInjection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderOfficeAppsOtherProcessInjectionType gets the defenderOfficeAppsOtherProcessInjectionType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeAppsOtherProcessInjectionType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderOfficeAppsOtherProcessInjectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderOfficeCommunicationAppsLaunchChildProcess gets the defenderOfficeCommunicationAppsLaunchChildProcess property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeCommunicationAppsLaunchChildProcess()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderOfficeCommunicationAppsLaunchChildProcess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderOfficeMacroCodeAllowWin32Imports gets the defenderOfficeMacroCodeAllowWin32Imports property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeMacroCodeAllowWin32Imports()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderOfficeMacroCodeAllowWin32Imports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderOfficeMacroCodeAllowWin32ImportsType gets the defenderOfficeMacroCodeAllowWin32ImportsType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderOfficeMacroCodeAllowWin32ImportsType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderOfficeMacroCodeAllowWin32ImportsType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderPotentiallyUnwantedAppAction gets the defenderPotentiallyUnwantedAppAction property value. Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs). Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderPotentiallyUnwantedAppAction()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderPotentiallyUnwantedAppAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderPreventCredentialStealingType gets the defenderPreventCredentialStealingType property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderPreventCredentialStealingType()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderPreventCredentialStealingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderProcessCreation gets the defenderProcessCreation property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessCreation()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderProcessCreation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderProcessCreationType gets the defenderProcessCreationType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessCreationType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderProcessCreationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderProcessesToExclude gets the defenderProcessesToExclude property value. Processes to exclude from scans and real time protection.
// returns a []string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderProcessesToExclude()([]string) {
    val, err := m.GetBackingStore().Get("defenderProcessesToExclude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDefenderScanDirection gets the defenderScanDirection property value. Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly, monitorOutgoingFilesOnly.
// returns a *DefenderRealtimeScanDirection when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanDirection()(*DefenderRealtimeScanDirection) {
    val, err := m.GetBackingStore().Get("defenderScanDirection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderRealtimeScanDirection)
    }
    return nil
}
// GetDefenderScanMaxCpuPercentage gets the defenderScanMaxCpuPercentage property value. Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid values 0 to 100
// returns a *int32 when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanMaxCpuPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("defenderScanMaxCpuPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDefenderScanType gets the defenderScanType property value. Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
// returns a *DefenderScanType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScanType()(*DefenderScanType) {
    val, err := m.GetBackingStore().Get("defenderScanType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderScanType)
    }
    return nil
}
// GetDefenderScheduledQuickScanTime gets the defenderScheduledQuickScanTime property value. Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
// returns a *TimeOnly when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledQuickScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("defenderScheduledQuickScanTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetDefenderScheduledScanDay gets the defenderScheduledScanDay property value. Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
// returns a *WeeklySchedule when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledScanDay()(*WeeklySchedule) {
    val, err := m.GetBackingStore().Get("defenderScheduledScanDay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WeeklySchedule)
    }
    return nil
}
// GetDefenderScheduledScanTime gets the defenderScheduledScanTime property value. Selects the time of day that the Windows Defender scan should run.
// returns a *TimeOnly when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScheduledScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("defenderScheduledScanTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetDefenderScriptDownloadedPayloadExecution gets the defenderScriptDownloadedPayloadExecution property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptDownloadedPayloadExecution()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderScriptDownloadedPayloadExecution")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderScriptDownloadedPayloadExecutionType gets the defenderScriptDownloadedPayloadExecutionType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptDownloadedPayloadExecutionType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderScriptDownloadedPayloadExecutionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderScriptObfuscatedMacroCode gets the defenderScriptObfuscatedMacroCode property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptObfuscatedMacroCode()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderScriptObfuscatedMacroCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderScriptObfuscatedMacroCodeType gets the defenderScriptObfuscatedMacroCodeType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderScriptObfuscatedMacroCodeType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderScriptObfuscatedMacroCodeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderSecurityCenterBlockExploitProtectionOverride gets the defenderSecurityCenterBlockExploitProtectionOverride property value. Indicates whether or not to block user from overriding Exploit Protection settings.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterBlockExploitProtectionOverride()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterBlockExploitProtectionOverride")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableAccountUI gets the defenderSecurityCenterDisableAccountUI property value. Used to disable the display of the account protection area.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableAccountUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableAccountUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableAppBrowserUI gets the defenderSecurityCenterDisableAppBrowserUI property value. Used to disable the display of the app and browser protection area.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableAppBrowserUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableAppBrowserUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableClearTpmUI gets the defenderSecurityCenterDisableClearTpmUI property value. Used to disable the display of the Clear TPM button.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableClearTpmUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableClearTpmUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableFamilyUI gets the defenderSecurityCenterDisableFamilyUI property value. Used to disable the display of the family options area.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableFamilyUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableFamilyUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableHardwareUI gets the defenderSecurityCenterDisableHardwareUI property value. Used to disable the display of the hardware protection area.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableHardwareUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableHardwareUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableHealthUI gets the defenderSecurityCenterDisableHealthUI property value. Used to disable the display of the device performance and health area.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableHealthUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableHealthUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableNetworkUI gets the defenderSecurityCenterDisableNetworkUI property value. Used to disable the display of the firewall and network protection area.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableNetworkUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableNetworkUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableNotificationAreaUI gets the defenderSecurityCenterDisableNotificationAreaUI property value. Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot the computer for this setting to take effect.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableNotificationAreaUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableNotificationAreaUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableRansomwareUI gets the defenderSecurityCenterDisableRansomwareUI property value. Used to disable the display of the ransomware protection area.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableRansomwareUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableRansomwareUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableSecureBootUI gets the defenderSecurityCenterDisableSecureBootUI property value. Used to disable the display of the secure boot area under Device security.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableSecureBootUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableSecureBootUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableTroubleshootingUI gets the defenderSecurityCenterDisableTroubleshootingUI property value. Used to disable the display of the security process troubleshooting under Device security.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableTroubleshootingUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableTroubleshootingUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableVirusUI gets the defenderSecurityCenterDisableVirusUI property value. Used to disable the display of the virus and threat protection area.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableVirusUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableVirusUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI gets the defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI property value. Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI()(*bool) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefenderSecurityCenterHelpEmail gets the defenderSecurityCenterHelpEmail property value. The email address that is displayed to users.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpEmail()(*string) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterHelpEmail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDefenderSecurityCenterHelpPhone gets the defenderSecurityCenterHelpPhone property value. The phone number or Skype ID that is displayed to users.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpPhone()(*string) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterHelpPhone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDefenderSecurityCenterHelpURL gets the defenderSecurityCenterHelpURL property value. The help portal URL this is displayed to users.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterHelpURL()(*string) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterHelpURL")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDefenderSecurityCenterITContactDisplay gets the defenderSecurityCenterITContactDisplay property value. Possible values for defenderSecurityCenterITContactDisplay
// returns a *DefenderSecurityCenterITContactDisplayType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterITContactDisplay()(*DefenderSecurityCenterITContactDisplayType) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterITContactDisplay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderSecurityCenterITContactDisplayType)
    }
    return nil
}
// GetDefenderSecurityCenterNotificationsFromApp gets the defenderSecurityCenterNotificationsFromApp property value. Possible values for defenderSecurityCenterNotificationsFromApp
// returns a *DefenderSecurityCenterNotificationsFromAppType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterNotificationsFromApp()(*DefenderSecurityCenterNotificationsFromAppType) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterNotificationsFromApp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderSecurityCenterNotificationsFromAppType)
    }
    return nil
}
// GetDefenderSecurityCenterOrganizationDisplayName gets the defenderSecurityCenterOrganizationDisplayName property value. The company name that is displayed to the users.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterOrganizationDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("defenderSecurityCenterOrganizationDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDefenderSignatureUpdateIntervalInHours gets the defenderSignatureUpdateIntervalInHours property value. Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
// returns a *int32 when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSignatureUpdateIntervalInHours()(*int32) {
    val, err := m.GetBackingStore().Get("defenderSignatureUpdateIntervalInHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDefenderSubmitSamplesConsentType gets the defenderSubmitSamplesConsentType property value. Checks for the user consent level in Windows Defender to send data. Possible values are: sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
// returns a *DefenderSubmitSamplesConsentType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSubmitSamplesConsentType()(*DefenderSubmitSamplesConsentType) {
    val, err := m.GetBackingStore().Get("defenderSubmitSamplesConsentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderSubmitSamplesConsentType)
    }
    return nil
}
// GetDefenderUntrustedExecutable gets the defenderUntrustedExecutable property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedExecutable()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderUntrustedExecutable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderUntrustedExecutableType gets the defenderUntrustedExecutableType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedExecutableType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderUntrustedExecutableType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDefenderUntrustedUSBProcess gets the defenderUntrustedUSBProcess property value. Possible values of Defender PUA Protection
// returns a *DefenderProtectionType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedUSBProcess()(*DefenderProtectionType) {
    val, err := m.GetBackingStore().Get("defenderUntrustedUSBProcess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderProtectionType)
    }
    return nil
}
// GetDefenderUntrustedUSBProcessType gets the defenderUntrustedUSBProcessType property value. Possible values of Defender Attack Surface Reduction Rules
// returns a *DefenderAttackSurfaceType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDefenderUntrustedUSBProcessType()(*DefenderAttackSurfaceType) {
    val, err := m.GetBackingStore().Get("defenderUntrustedUSBProcessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DefenderAttackSurfaceType)
    }
    return nil
}
// GetDeviceGuardEnableSecureBootWithDMA gets the deviceGuardEnableSecureBootWithDMA property value. This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA. Specifies whether Platform Security Level is enabled at next reboot.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardEnableSecureBootWithDMA()(*bool) {
    val, err := m.GetBackingStore().Get("deviceGuardEnableSecureBootWithDMA")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceGuardEnableVirtualizationBasedSecurity gets the deviceGuardEnableVirtualizationBasedSecurity property value. Turns On Virtualization Based Security(VBS).
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardEnableVirtualizationBasedSecurity()(*bool) {
    val, err := m.GetBackingStore().Get("deviceGuardEnableVirtualizationBasedSecurity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceGuardLaunchSystemGuard gets the deviceGuardLaunchSystemGuard property value. Possible values of a property
// returns a *Enablement when successful
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardLaunchSystemGuard()(*Enablement) {
    val, err := m.GetBackingStore().Get("deviceGuardLaunchSystemGuard")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Enablement)
    }
    return nil
}
// GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings gets the deviceGuardLocalSystemAuthorityCredentialGuardSettings property value. Possible values of Credential Guard settings.
// returns a *DeviceGuardLocalSystemAuthorityCredentialGuardType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings()(*DeviceGuardLocalSystemAuthorityCredentialGuardType) {
    val, err := m.GetBackingStore().Get("deviceGuardLocalSystemAuthorityCredentialGuardSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceGuardLocalSystemAuthorityCredentialGuardType)
    }
    return nil
}
// GetDeviceGuardSecureBootWithDMA gets the deviceGuardSecureBootWithDMA property value. Possible values of Secure Boot with DMA
// returns a *SecureBootWithDMAType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDeviceGuardSecureBootWithDMA()(*SecureBootWithDMAType) {
    val, err := m.GetBackingStore().Get("deviceGuardSecureBootWithDMA")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecureBootWithDMAType)
    }
    return nil
}
// GetDmaGuardDeviceEnumerationPolicy gets the dmaGuardDeviceEnumerationPolicy property value. Possible values of the DmaGuardDeviceEnumerationPolicy.
// returns a *DmaGuardDeviceEnumerationPolicyType when successful
func (m *Windows10EndpointProtectionConfiguration) GetDmaGuardDeviceEnumerationPolicy()(*DmaGuardDeviceEnumerationPolicyType) {
    val, err := m.GetBackingStore().Get("dmaGuardDeviceEnumerationPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DmaGuardDeviceEnumerationPolicyType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = v.(WindowsFirewallRuleable)
                }
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
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallBlockStatefulFTP()(*bool) {
    val, err := m.GetBackingStore().Get("firewallBlockStatefulFTP")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallCertificateRevocationListCheckMethod gets the firewallCertificateRevocationListCheckMethod property value. Possible values for firewallCertificateRevocationListCheckMethod
// returns a *FirewallCertificateRevocationListCheckMethodType when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallCertificateRevocationListCheckMethod()(*FirewallCertificateRevocationListCheckMethodType) {
    val, err := m.GetBackingStore().Get("firewallCertificateRevocationListCheckMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FirewallCertificateRevocationListCheckMethodType)
    }
    return nil
}
// GetFirewallIdleTimeoutForSecurityAssociationInSeconds gets the firewallIdleTimeoutForSecurityAssociationInSeconds property value. Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
// returns a *int32 when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIdleTimeoutForSecurityAssociationInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("firewallIdleTimeoutForSecurityAssociationInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFirewallIPSecExemptionsAllowDHCP gets the firewallIPSecExemptionsAllowDHCP property value. Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowDHCP()(*bool) {
    val, err := m.GetBackingStore().Get("firewallIPSecExemptionsAllowDHCP")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallIPSecExemptionsAllowICMP gets the firewallIPSecExemptionsAllowICMP property value. Configures IPSec exemptions to allow ICMP
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowICMP()(*bool) {
    val, err := m.GetBackingStore().Get("firewallIPSecExemptionsAllowICMP")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallIPSecExemptionsAllowNeighborDiscovery gets the firewallIPSecExemptionsAllowNeighborDiscovery property value. Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowNeighborDiscovery()(*bool) {
    val, err := m.GetBackingStore().Get("firewallIPSecExemptionsAllowNeighborDiscovery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallIPSecExemptionsAllowRouterDiscovery gets the firewallIPSecExemptionsAllowRouterDiscovery property value. Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowRouterDiscovery()(*bool) {
    val, err := m.GetBackingStore().Get("firewallIPSecExemptionsAllowRouterDiscovery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallIPSecExemptionsNone gets the firewallIPSecExemptionsNone property value. Configures IPSec exemptions to no exemptions
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsNone()(*bool) {
    val, err := m.GetBackingStore().Get("firewallIPSecExemptionsNone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallMergeKeyingModuleSettings gets the firewallMergeKeyingModuleSettings property value. If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallMergeKeyingModuleSettings()(*bool) {
    val, err := m.GetBackingStore().Get("firewallMergeKeyingModuleSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFirewallPacketQueueingMethod gets the firewallPacketQueueingMethod property value. Possible values for firewallPacketQueueingMethod
// returns a *FirewallPacketQueueingMethodType when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallPacketQueueingMethod()(*FirewallPacketQueueingMethodType) {
    val, err := m.GetBackingStore().Get("firewallPacketQueueingMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FirewallPacketQueueingMethodType)
    }
    return nil
}
// GetFirewallPreSharedKeyEncodingMethod gets the firewallPreSharedKeyEncodingMethod property value. Possible values for firewallPreSharedKeyEncodingMethod
// returns a *FirewallPreSharedKeyEncodingMethodType when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallPreSharedKeyEncodingMethod()(*FirewallPreSharedKeyEncodingMethodType) {
    val, err := m.GetBackingStore().Get("firewallPreSharedKeyEncodingMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FirewallPreSharedKeyEncodingMethodType)
    }
    return nil
}
// GetFirewallProfileDomain gets the firewallProfileDomain property value. Configures the firewall profile settings for domain networks
// returns a WindowsFirewallNetworkProfileable when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfileDomain()(WindowsFirewallNetworkProfileable) {
    val, err := m.GetBackingStore().Get("firewallProfileDomain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsFirewallNetworkProfileable)
    }
    return nil
}
// GetFirewallProfilePrivate gets the firewallProfilePrivate property value. Configures the firewall profile settings for private networks
// returns a WindowsFirewallNetworkProfileable when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfilePrivate()(WindowsFirewallNetworkProfileable) {
    val, err := m.GetBackingStore().Get("firewallProfilePrivate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsFirewallNetworkProfileable)
    }
    return nil
}
// GetFirewallProfilePublic gets the firewallProfilePublic property value. Configures the firewall profile settings for public networks
// returns a WindowsFirewallNetworkProfileable when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfilePublic()(WindowsFirewallNetworkProfileable) {
    val, err := m.GetBackingStore().Get("firewallProfilePublic")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsFirewallNetworkProfileable)
    }
    return nil
}
// GetFirewallRules gets the firewallRules property value. Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
// returns a []WindowsFirewallRuleable when successful
func (m *Windows10EndpointProtectionConfiguration) GetFirewallRules()([]WindowsFirewallRuleable) {
    val, err := m.GetBackingStore().Get("firewallRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsFirewallRuleable)
    }
    return nil
}
// GetLanManagerAuthenticationLevel gets the lanManagerAuthenticationLevel property value. Possible values for LanManagerAuthenticationLevel
// returns a *LanManagerAuthenticationLevel when successful
func (m *Windows10EndpointProtectionConfiguration) GetLanManagerAuthenticationLevel()(*LanManagerAuthenticationLevel) {
    val, err := m.GetBackingStore().Get("lanManagerAuthenticationLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LanManagerAuthenticationLevel)
    }
    return nil
}
// GetLanManagerWorkstationDisableInsecureGuestLogons gets the lanManagerWorkstationDisableInsecureGuestLogons property value. If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure guest logons.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLanManagerWorkstationDisableInsecureGuestLogons()(*bool) {
    val, err := m.GetBackingStore().Get("lanManagerWorkstationDisableInsecureGuestLogons")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsAdministratorAccountName gets the localSecurityOptionsAdministratorAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAdministratorAccountName()(*string) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAdministratorAccountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalSecurityOptionsAdministratorElevationPromptBehavior gets the localSecurityOptionsAdministratorElevationPromptBehavior property value. Possible values for LocalSecurityOptionsAdministratorElevationPromptBehavior
// returns a *LocalSecurityOptionsAdministratorElevationPromptBehaviorType when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAdministratorElevationPromptBehavior()(*LocalSecurityOptionsAdministratorElevationPromptBehaviorType) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAdministratorElevationPromptBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LocalSecurityOptionsAdministratorElevationPromptBehaviorType)
    }
    return nil
}
// GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares gets the localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares property value. This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating the names of domain accounts and network shares.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests gets the localSecurityOptionsAllowPKU2UAuthenticationRequests property value. Block PKU2U authentication requests to this device to use online identities.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAllowPKU2UAuthenticationRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager gets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager property value. Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote calls to the SAM.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager()(*string) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool gets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool property value. UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn gets the localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn property value. This security setting determines whether a computer can be shut down without having to log on to Windows.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsAllowUIAccessApplicationElevation gets the localSecurityOptionsAllowUIAccessApplicationElevation property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUIAccessApplicationElevation()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAllowUIAccessApplicationElevation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations gets the localSecurityOptionsAllowUIAccessApplicationsForSecureLocations property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAllowUIAccessApplicationsForSecureLocations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon gets the localSecurityOptionsAllowUndockWithoutHavingToLogon property value. Prevent a portable computer from being undocked without having to log in.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsAllowUndockWithoutHavingToLogon")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsBlockMicrosoftAccounts gets the localSecurityOptionsBlockMicrosoftAccounts property value. Prevent users from adding new Microsoft accounts to this computer.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockMicrosoftAccounts()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsBlockMicrosoftAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword gets the localSecurityOptionsBlockRemoteLogonWithBlankPassword property value. Enable Local accounts that are not password protected to log on from locations other than the physical device.Default is enabled
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsBlockRemoteLogonWithBlankPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess gets the localSecurityOptionsBlockRemoteOpticalDriveAccess property value. Enabling this settings allows only interactively logged on user to access CD-ROM media.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsBlockRemoteOpticalDriveAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers gets the localSecurityOptionsBlockUsersInstallingPrinterDrivers property value. Restrict installing printer drivers as part of connecting to a shared printer to admins only.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsBlockUsersInstallingPrinterDrivers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsClearVirtualMemoryPageFile gets the localSecurityOptionsClearVirtualMemoryPageFile property value. This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClearVirtualMemoryPageFile()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsClearVirtualMemoryPageFile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways gets the localSecurityOptionsClientDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB client component.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsClientDigitallySignCommunicationsAlways")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers gets the localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers property value. If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords to non-Microsoft SMB servers that do not support password encryption during authentication.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation gets the localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation property value. App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDisableAdministratorAccount gets the localSecurityOptionsDisableAdministratorAccount property value. Determines whether the Local Administrator account is enabled or disabled.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableAdministratorAccount()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDisableAdministratorAccount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees gets the localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees property value. This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDisableGuestAccount gets the localSecurityOptionsDisableGuestAccount property value. Determines if the Guest account is enabled or disabled.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableGuestAccount()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDisableGuestAccount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways gets the localSecurityOptionsDisableServerDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB server component.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDisableServerDigitallySignCommunicationsAlways")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees gets the localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees property value. This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request it.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts gets the localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts property value. This security setting determines what additional permissions will be granted for anonymous connections to the computer.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDoNotRequireCtrlAltDel gets the localSecurityOptionsDoNotRequireCtrlAltDel property value. Require CTRL+ALT+DEL to be pressed before a user can log on.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotRequireCtrlAltDel()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDoNotRequireCtrlAltDel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange gets the localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange property value. This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new password is stored. It’s not stored by default.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser gets the localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser property value. Possible values for LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser
// returns a *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser()(*LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType)
    }
    return nil
}
// GetLocalSecurityOptionsGuestAccountName gets the localSecurityOptionsGuestAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsGuestAccountName()(*string) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsGuestAccountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalSecurityOptionsHideLastSignedInUser gets the localSecurityOptionsHideLastSignedInUser property value. Do not display the username of the last person who signed in on this device.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsHideLastSignedInUser()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsHideLastSignedInUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsHideUsernameAtSignIn gets the localSecurityOptionsHideUsernameAtSignIn property value. Do not display the username of the person signing in to this device after credentials are entered and before the device’s desktop is shown.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsHideUsernameAtSignIn()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsHideUsernameAtSignIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsInformationDisplayedOnLockScreen gets the localSecurityOptionsInformationDisplayedOnLockScreen property value. Possible values for LocalSecurityOptionsInformationDisplayedOnLockScreen
// returns a *LocalSecurityOptionsInformationDisplayedOnLockScreenType when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsInformationDisplayedOnLockScreen()(*LocalSecurityOptionsInformationDisplayedOnLockScreenType) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsInformationDisplayedOnLockScreen")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LocalSecurityOptionsInformationDisplayedOnLockScreenType)
    }
    return nil
}
// GetLocalSecurityOptionsInformationShownOnLockScreen gets the localSecurityOptionsInformationShownOnLockScreen property value. Possible values for LocalSecurityOptionsInformationShownOnLockScreenType
// returns a *LocalSecurityOptionsInformationShownOnLockScreenType when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsInformationShownOnLockScreen()(*LocalSecurityOptionsInformationShownOnLockScreenType) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsInformationShownOnLockScreen")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LocalSecurityOptionsInformationShownOnLockScreenType)
    }
    return nil
}
// GetLocalSecurityOptionsLogOnMessageText gets the localSecurityOptionsLogOnMessageText property value. Set message text for users attempting to log in.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsLogOnMessageText()(*string) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsLogOnMessageText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalSecurityOptionsLogOnMessageTitle gets the localSecurityOptionsLogOnMessageTitle property value. Set message title for users attempting to log in.
// returns a *string when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsLogOnMessageTitle()(*string) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsLogOnMessageTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalSecurityOptionsMachineInactivityLimit gets the localSecurityOptionsMachineInactivityLimit property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
// returns a *int32 when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMachineInactivityLimit()(*int32) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsMachineInactivityLimit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLocalSecurityOptionsMachineInactivityLimitInMinutes gets the localSecurityOptionsMachineInactivityLimitInMinutes property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
// returns a *int32 when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMachineInactivityLimitInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsMachineInactivityLimitInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients gets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients property value. Possible values for LocalSecurityOptionsMinimumSessionSecurity
// returns a *LocalSecurityOptionsMinimumSessionSecurity when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients()(*LocalSecurityOptionsMinimumSessionSecurity) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LocalSecurityOptionsMinimumSessionSecurity)
    }
    return nil
}
// GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers gets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers property value. Possible values for LocalSecurityOptionsMinimumSessionSecurity
// returns a *LocalSecurityOptionsMinimumSessionSecurity when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers()(*LocalSecurityOptionsMinimumSessionSecurity) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LocalSecurityOptionsMinimumSessionSecurity)
    }
    return nil
}
// GetLocalSecurityOptionsOnlyElevateSignedExecutables gets the localSecurityOptionsOnlyElevateSignedExecutables property value. Enforce PKI certification path validation for a given executable file before it is permitted to run.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsOnlyElevateSignedExecutables()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsOnlyElevateSignedExecutables")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares gets the localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares property value. By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that can be accessed anonymously and Shares that can be accessed anonymously
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsSmartCardRemovalBehavior gets the localSecurityOptionsSmartCardRemovalBehavior property value. Possible values for LocalSecurityOptionsSmartCardRemovalBehaviorType
// returns a *LocalSecurityOptionsSmartCardRemovalBehaviorType when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsSmartCardRemovalBehavior()(*LocalSecurityOptionsSmartCardRemovalBehaviorType) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsSmartCardRemovalBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LocalSecurityOptionsSmartCardRemovalBehaviorType)
    }
    return nil
}
// GetLocalSecurityOptionsStandardUserElevationPromptBehavior gets the localSecurityOptionsStandardUserElevationPromptBehavior property value. Possible values for LocalSecurityOptionsStandardUserElevationPromptBehavior
// returns a *LocalSecurityOptionsStandardUserElevationPromptBehaviorType when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsStandardUserElevationPromptBehavior()(*LocalSecurityOptionsStandardUserElevationPromptBehaviorType) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsStandardUserElevationPromptBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LocalSecurityOptionsStandardUserElevationPromptBehaviorType)
    }
    return nil
}
// GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation gets the localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation property value. Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior policy settings for admins and standard users are used.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsUseAdminApprovalMode gets the localSecurityOptionsUseAdminApprovalMode property value. Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin privileges.Default is enabled
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsUseAdminApprovalMode()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsUseAdminApprovalMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators gets the localSecurityOptionsUseAdminApprovalModeForAdministrators property value. Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsUseAdminApprovalModeForAdministrators")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations gets the localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations property value. Virtualize file and registry write failures to per user locations
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations()(*bool) {
    val, err := m.GetBackingStore().Get("localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmartScreenBlockOverrideForFiles gets the smartScreenBlockOverrideForFiles property value. Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetSmartScreenBlockOverrideForFiles()(*bool) {
    val, err := m.GetBackingStore().Get("smartScreenBlockOverrideForFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSmartScreenEnableInShell gets the smartScreenEnableInShell property value. Allows IT Admins to configure SmartScreen for Windows.
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetSmartScreenEnableInShell()(*bool) {
    val, err := m.GetBackingStore().Get("smartScreenEnableInShell")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserRightsAccessCredentialManagerAsTrustedCaller gets the userRightsAccessCredentialManagerAsTrustedCaller property value. This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if this privilege is given to other entities. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsAccessCredentialManagerAsTrustedCaller()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsAccessCredentialManagerAsTrustedCaller")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsActAsPartOfTheOperatingSystem gets the userRightsActAsPartOfTheOperatingSystem property value. This user right allows a process to impersonate any user without authentication. The process can therefore gain access to the same local resources as that user. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsActAsPartOfTheOperatingSystem()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsActAsPartOfTheOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsAllowAccessFromNetwork gets the userRightsAllowAccessFromNetwork property value. This user right determines which users and groups are allowed to connect to the computer over the network. State Allowed is supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsAllowAccessFromNetwork()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsAllowAccessFromNetwork")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsBackupData gets the userRightsBackupData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when backing up files and directories. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsBackupData()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsBackupData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsBlockAccessFromNetwork gets the userRightsBlockAccessFromNetwork property value. This user right determines which users and groups are block from connecting to the computer over the network. State Block is supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsBlockAccessFromNetwork()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsBlockAccessFromNetwork")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsChangeSystemTime gets the userRightsChangeSystemTime property value. This user right determines which users and groups can change the time and date on the internal clock of the computer. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsChangeSystemTime()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsChangeSystemTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsCreateGlobalObjects gets the userRightsCreateGlobalObjects property value. This security setting determines whether users can create global objects that are available to all sessions. Users who can create global objects could affect processes that run under other users' sessions, which could lead to application failure or data corruption. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateGlobalObjects()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsCreateGlobalObjects")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsCreatePageFile gets the userRightsCreatePageFile property value. This user right determines which users and groups can call an internal API to create and change the size of a page file. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreatePageFile()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsCreatePageFile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsCreatePermanentSharedObjects gets the userRightsCreatePermanentSharedObjects property value. This user right determines which accounts can be used by processes to create a directory object using the object manager. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreatePermanentSharedObjects()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsCreatePermanentSharedObjects")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsCreateSymbolicLinks gets the userRightsCreateSymbolicLinks property value. This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateSymbolicLinks()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsCreateSymbolicLinks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsCreateToken gets the userRightsCreateToken property value. This user right determines which users/groups can be used by processes to create a token that can then be used to get access to any local resources when the process uses an internal API to create an access token. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsCreateToken()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsCreateToken")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsDebugPrograms gets the userRightsDebugPrograms property value. This user right determines which users can attach a debugger to any process or to the kernel. Only states NotConfigured and Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDebugPrograms()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsDebugPrograms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsDelegation gets the userRightsDelegation property value. This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDelegation()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsDelegation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsDenyLocalLogOn gets the userRightsDenyLocalLogOn property value. This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsDenyLocalLogOn()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsDenyLocalLogOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsGenerateSecurityAudits gets the userRightsGenerateSecurityAudits property value. This user right determines which accounts can be used by a process to add entries to the security log. The security log is used to trace unauthorized system access.  Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsGenerateSecurityAudits()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsGenerateSecurityAudits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsImpersonateClient gets the userRightsImpersonateClient property value. Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a service that they have created and then impersonating that client, which can elevate the unauthorized user's permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsImpersonateClient()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsImpersonateClient")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsIncreaseSchedulingPriority gets the userRightsIncreaseSchedulingPriority property value. This user right determines which accounts can use a process with Write Property access to another process to increase the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsIncreaseSchedulingPriority()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsIncreaseSchedulingPriority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsLoadUnloadDrivers gets the userRightsLoadUnloadDrivers property value. This user right determines which users can dynamically load and unload device drivers or other code in to kernel mode. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLoadUnloadDrivers()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsLoadUnloadDrivers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsLocalLogOn gets the userRightsLocalLogOn property value. This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLocalLogOn()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsLocalLogOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsLockMemory gets the userRightsLockMemory property value. This user right determines which accounts can use a process to keep data in physical memory, which prevents the system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsLockMemory()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsLockMemory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsManageAuditingAndSecurityLogs gets the userRightsManageAuditingAndSecurityLogs property value. This user right determines which users can specify object access auditing options for individual resources, such as files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsManageAuditingAndSecurityLogs()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsManageAuditingAndSecurityLogs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsManageVolumes gets the userRightsManageVolumes property value. This user right determines which users and groups can run maintenance tasks on a volume, such as remote defragmentation. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsManageVolumes()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsManageVolumes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsModifyFirmwareEnvironment gets the userRightsModifyFirmwareEnvironment property value. This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsModifyFirmwareEnvironment()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsModifyFirmwareEnvironment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsModifyObjectLabels gets the userRightsModifyObjectLabels property value. This user right determines which user accounts can modify the integrity label of objects, such as files, registry keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsModifyObjectLabels()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsModifyObjectLabels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsProfileSingleProcess gets the userRightsProfileSingleProcess property value. This user right determines which users can use performance monitoring tools to monitor the performance of system processes. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsProfileSingleProcess()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsProfileSingleProcess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsRemoteDesktopServicesLogOn gets the userRightsRemoteDesktopServicesLogOn property value. This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client. Only states NotConfigured and Blocked are supported
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRemoteDesktopServicesLogOn()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsRemoteDesktopServicesLogOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsRemoteShutdown gets the userRightsRemoteShutdown property value. This user right determines which users are allowed to shut down a computer from a remote location on the network. Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRemoteShutdown()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsRemoteShutdown")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsRestoreData gets the userRightsRestoreData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when restoring backed up files and directories, and determines which users can set any valid security principal as the owner of an object. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsRestoreData()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsRestoreData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetUserRightsTakeOwnership gets the userRightsTakeOwnership property value. This user right determines which users can take ownership of any securable object in the system, including Active Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and Allowed are supported.
// returns a DeviceManagementUserRightsSettingable when successful
func (m *Windows10EndpointProtectionConfiguration) GetUserRightsTakeOwnership()(DeviceManagementUserRightsSettingable) {
    val, err := m.GetBackingStore().Get("userRightsTakeOwnership")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementUserRightsSettingable)
    }
    return nil
}
// GetWindowsDefenderTamperProtection gets the windowsDefenderTamperProtection property value. Defender TamperProtection setting options
// returns a *WindowsDefenderTamperProtectionOptions when successful
func (m *Windows10EndpointProtectionConfiguration) GetWindowsDefenderTamperProtection()(*WindowsDefenderTamperProtectionOptions) {
    val, err := m.GetBackingStore().Get("windowsDefenderTamperProtection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsDefenderTamperProtectionOptions)
    }
    return nil
}
// GetXboxServicesAccessoryManagementServiceStartupMode gets the xboxServicesAccessoryManagementServiceStartupMode property value. Possible values of xbox service start type
// returns a *ServiceStartType when successful
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesAccessoryManagementServiceStartupMode()(*ServiceStartType) {
    val, err := m.GetBackingStore().Get("xboxServicesAccessoryManagementServiceStartupMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ServiceStartType)
    }
    return nil
}
// GetXboxServicesEnableXboxGameSaveTask gets the xboxServicesEnableXboxGameSaveTask property value. This setting determines whether xbox game save is enabled (1) or disabled (0).
// returns a *bool when successful
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesEnableXboxGameSaveTask()(*bool) {
    val, err := m.GetBackingStore().Get("xboxServicesEnableXboxGameSaveTask")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetXboxServicesLiveAuthManagerServiceStartupMode gets the xboxServicesLiveAuthManagerServiceStartupMode property value. Possible values of xbox service start type
// returns a *ServiceStartType when successful
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveAuthManagerServiceStartupMode()(*ServiceStartType) {
    val, err := m.GetBackingStore().Get("xboxServicesLiveAuthManagerServiceStartupMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ServiceStartType)
    }
    return nil
}
// GetXboxServicesLiveGameSaveServiceStartupMode gets the xboxServicesLiveGameSaveServiceStartupMode property value. Possible values of xbox service start type
// returns a *ServiceStartType when successful
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveGameSaveServiceStartupMode()(*ServiceStartType) {
    val, err := m.GetBackingStore().Get("xboxServicesLiveGameSaveServiceStartupMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ServiceStartType)
    }
    return nil
}
// GetXboxServicesLiveNetworkingServiceStartupMode gets the xboxServicesLiveNetworkingServiceStartupMode property value. Possible values of xbox service start type
// returns a *ServiceStartType when successful
func (m *Windows10EndpointProtectionConfiguration) GetXboxServicesLiveNetworkingServiceStartupMode()(*ServiceStartType) {
    val, err := m.GetBackingStore().Get("xboxServicesLiveNetworkingServiceStartupMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ServiceStartType)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    return nil
}
// SetApplicationGuardAllowCameraMicrophoneRedirection sets the applicationGuardAllowCameraMicrophoneRedirection property value. Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and microphone.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowCameraMicrophoneRedirection(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardAllowCameraMicrophoneRedirection", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardAllowFileSaveOnHost sets the applicationGuardAllowFileSaveOnHost property value. Allow users to download files from Edge in the application guard container and save them on the host file system
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowFileSaveOnHost(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardAllowFileSaveOnHost", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardAllowPersistence sets the applicationGuardAllowPersistence property value. Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPersistence(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardAllowPersistence", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardAllowPrintToLocalPrinters sets the applicationGuardAllowPrintToLocalPrinters property value. Allow printing to Local Printers from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToLocalPrinters(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardAllowPrintToLocalPrinters", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardAllowPrintToNetworkPrinters sets the applicationGuardAllowPrintToNetworkPrinters property value. Allow printing to Network Printers from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToNetworkPrinters(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardAllowPrintToNetworkPrinters", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardAllowPrintToPDF sets the applicationGuardAllowPrintToPDF property value. Allow printing to PDF from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToPDF(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardAllowPrintToPDF", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardAllowPrintToXPS sets the applicationGuardAllowPrintToXPS property value. Allow printing to XPS from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToXPS(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardAllowPrintToXPS", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardAllowVirtualGPU sets the applicationGuardAllowVirtualGPU property value. Allow application guard to use virtual GPU
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowVirtualGPU(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardAllowVirtualGPU", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardBlockClipboardSharing sets the applicationGuardBlockClipboardSharing property value. Possible values for applicationGuardBlockClipboardSharingType
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockClipboardSharing(value *ApplicationGuardBlockClipboardSharingType)() {
    err := m.GetBackingStore().Set("applicationGuardBlockClipboardSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardBlockFileTransfer sets the applicationGuardBlockFileTransfer property value. Possible values for applicationGuardBlockFileTransfer
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockFileTransfer(value *ApplicationGuardBlockFileTransferType)() {
    err := m.GetBackingStore().Set("applicationGuardBlockFileTransfer", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardBlockNonEnterpriseContent sets the applicationGuardBlockNonEnterpriseContent property value. Block enterprise sites to load non-enterprise content, such as third party plug-ins
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockNonEnterpriseContent(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardBlockNonEnterpriseContent", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardCertificateThumbprints sets the applicationGuardCertificateThumbprints property value. Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardCertificateThumbprints(value []string)() {
    err := m.GetBackingStore().Set("applicationGuardCertificateThumbprints", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardEnabled sets the applicationGuardEnabled property value. Enable Windows Defender Application Guard
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardEnabled(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardEnabledOptions sets the applicationGuardEnabledOptions property value. Possible values for ApplicationGuardEnabledOptions
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardEnabledOptions(value *ApplicationGuardEnabledOptions)() {
    err := m.GetBackingStore().Set("applicationGuardEnabledOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationGuardForceAuditing sets the applicationGuardForceAuditing property value. Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardForceAuditing(value *bool)() {
    err := m.GetBackingStore().Set("applicationGuardForceAuditing", value)
    if err != nil {
        panic(err)
    }
}
// SetAppLockerApplicationControl sets the appLockerApplicationControl property value. Possible values of AppLocker Application Control Types
func (m *Windows10EndpointProtectionConfiguration) SetAppLockerApplicationControl(value *AppLockerApplicationControlType)() {
    err := m.GetBackingStore().Set("appLockerApplicationControl", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerAllowStandardUserEncryption sets the bitLockerAllowStandardUserEncryption property value. Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerAllowStandardUserEncryption(value *bool)() {
    err := m.GetBackingStore().Set("bitLockerAllowStandardUserEncryption", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerDisableWarningForOtherDiskEncryption sets the bitLockerDisableWarningForOtherDiskEncryption property value. Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerDisableWarningForOtherDiskEncryption(value *bool)() {
    err := m.GetBackingStore().Set("bitLockerDisableWarningForOtherDiskEncryption", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerEnableStorageCardEncryptionOnMobile sets the bitLockerEnableStorageCardEncryptionOnMobile property value. Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerEnableStorageCardEncryptionOnMobile(value *bool)() {
    err := m.GetBackingStore().Set("bitLockerEnableStorageCardEncryptionOnMobile", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerEncryptDevice sets the bitLockerEncryptDevice property value. Allows the admin to require encryption to be turned on using BitLocker.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerEncryptDevice(value *bool)() {
    err := m.GetBackingStore().Set("bitLockerEncryptDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerFixedDrivePolicy sets the bitLockerFixedDrivePolicy property value. BitLocker Fixed Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerFixedDrivePolicy(value BitLockerFixedDrivePolicyable)() {
    err := m.GetBackingStore().Set("bitLockerFixedDrivePolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerRecoveryPasswordRotation sets the bitLockerRecoveryPasswordRotation property value. BitLocker recovery password rotation type
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerRecoveryPasswordRotation(value *BitLockerRecoveryPasswordRotationType)() {
    err := m.GetBackingStore().Set("bitLockerRecoveryPasswordRotation", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerRemovableDrivePolicy sets the bitLockerRemovableDrivePolicy property value. BitLocker Removable Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerRemovableDrivePolicy(value BitLockerRemovableDrivePolicyable)() {
    err := m.GetBackingStore().Set("bitLockerRemovableDrivePolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetBitLockerSystemDrivePolicy sets the bitLockerSystemDrivePolicy property value. BitLocker System Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerSystemDrivePolicy(value BitLockerSystemDrivePolicyable)() {
    err := m.GetBackingStore().Set("bitLockerSystemDrivePolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAdditionalGuardedFolders sets the defenderAdditionalGuardedFolders property value. List of folder paths to be added to the list of protected folders
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdditionalGuardedFolders(value []string)() {
    err := m.GetBackingStore().Set("defenderAdditionalGuardedFolders", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAdobeReaderLaunchChildProcess sets the defenderAdobeReaderLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdobeReaderLaunchChildProcess(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderAdobeReaderLaunchChildProcess", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAdvancedRansomewareProtectionType sets the defenderAdvancedRansomewareProtectionType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdvancedRansomewareProtectionType(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderAdvancedRansomewareProtectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowBehaviorMonitoring sets the defenderAllowBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowBehaviorMonitoring(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowBehaviorMonitoring", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowCloudProtection sets the defenderAllowCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowCloudProtection(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowCloudProtection", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowEndUserAccess sets the defenderAllowEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowEndUserAccess(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowEndUserAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowIntrusionPreventionSystem sets the defenderAllowIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowIntrusionPreventionSystem(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowIntrusionPreventionSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowOnAccessProtection sets the defenderAllowOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowOnAccessProtection(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowOnAccessProtection", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowRealTimeMonitoring sets the defenderAllowRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowRealTimeMonitoring(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowRealTimeMonitoring", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowScanArchiveFiles sets the defenderAllowScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanArchiveFiles(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowScanArchiveFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowScanDownloads sets the defenderAllowScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanDownloads(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowScanDownloads", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowScanNetworkFiles sets the defenderAllowScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanNetworkFiles(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowScanNetworkFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowScanRemovableDrivesDuringFullScan sets the defenderAllowScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanRemovableDrivesDuringFullScan(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowScanRemovableDrivesDuringFullScan", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAllowScanScriptsLoadedInInternetExplorer sets the defenderAllowScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAllowScanScriptsLoadedInInternetExplorer(value *bool)() {
    err := m.GetBackingStore().Set("defenderAllowScanScriptsLoadedInInternetExplorer", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderAttackSurfaceReductionExcludedPaths sets the defenderAttackSurfaceReductionExcludedPaths property value. List of exe files and folders to be excluded from attack surface reduction rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAttackSurfaceReductionExcludedPaths(value []string)() {
    err := m.GetBackingStore().Set("defenderAttackSurfaceReductionExcludedPaths", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderBlockEndUserAccess sets the defenderBlockEndUserAccess property value. Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will also be suppressed.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderBlockEndUserAccess(value *bool)() {
    err := m.GetBackingStore().Set("defenderBlockEndUserAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderBlockPersistenceThroughWmiType sets the defenderBlockPersistenceThroughWmiType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderBlockPersistenceThroughWmiType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderBlockPersistenceThroughWmiType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderCheckForSignaturesBeforeRunningScan sets the defenderCheckForSignaturesBeforeRunningScan property value. This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before running a scan.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCheckForSignaturesBeforeRunningScan(value *bool)() {
    err := m.GetBackingStore().Set("defenderCheckForSignaturesBeforeRunningScan", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderCloudBlockLevel sets the defenderCloudBlockLevel property value. Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS' setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCloudBlockLevel(value *DefenderCloudBlockLevelType)() {
    err := m.GetBackingStore().Set("defenderCloudBlockLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderCloudExtendedTimeoutInSeconds sets the defenderCloudExtendedTimeoutInSeconds property value. Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
func (m *Windows10EndpointProtectionConfiguration) SetDefenderCloudExtendedTimeoutInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("defenderCloudExtendedTimeoutInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDaysBeforeDeletingQuarantinedMalware sets the defenderDaysBeforeDeletingQuarantinedMalware property value. Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDaysBeforeDeletingQuarantinedMalware(value *int32)() {
    err := m.GetBackingStore().Set("defenderDaysBeforeDeletingQuarantinedMalware", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDetectedMalwareActions sets the defenderDetectedMalwareActions property value. Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDetectedMalwareActions(value DefenderDetectedMalwareActionsable)() {
    err := m.GetBackingStore().Set("defenderDetectedMalwareActions", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableBehaviorMonitoring sets the defenderDisableBehaviorMonitoring property value. Allows or disallows Windows Defender Behavior Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableBehaviorMonitoring(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableBehaviorMonitoring", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableCatchupFullScan sets the defenderDisableCatchupFullScan property value. This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCatchupFullScan(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableCatchupFullScan", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableCatchupQuickScan sets the defenderDisableCatchupQuickScan property value. This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the computer was turned off at the scheduled time.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCatchupQuickScan(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableCatchupQuickScan", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableCloudProtection sets the defenderDisableCloudProtection property value. To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft will analyze that information, learn more about problems affecting you and other customers, and offer improved solutions.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableCloudProtection(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableCloudProtection", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableIntrusionPreventionSystem sets the defenderDisableIntrusionPreventionSystem property value. Allows or disallows Windows Defender Intrusion Prevention functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableIntrusionPreventionSystem(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableIntrusionPreventionSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableOnAccessProtection sets the defenderDisableOnAccessProtection property value. Allows or disallows Windows Defender On Access Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableOnAccessProtection(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableOnAccessProtection", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableRealTimeMonitoring sets the defenderDisableRealTimeMonitoring property value. Allows or disallows Windows Defender Realtime Monitoring functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableRealTimeMonitoring(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableRealTimeMonitoring", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableScanArchiveFiles sets the defenderDisableScanArchiveFiles property value. Allows or disallows scanning of archives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanArchiveFiles(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableScanArchiveFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableScanDownloads sets the defenderDisableScanDownloads property value. Allows or disallows Windows Defender IOAVP Protection functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanDownloads(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableScanDownloads", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableScanNetworkFiles sets the defenderDisableScanNetworkFiles property value. Allows or disallows a scanning of network files.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanNetworkFiles(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableScanNetworkFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableScanRemovableDrivesDuringFullScan sets the defenderDisableScanRemovableDrivesDuringFullScan property value. Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanRemovableDrivesDuringFullScan(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableScanRemovableDrivesDuringFullScan", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderDisableScanScriptsLoadedInInternetExplorer sets the defenderDisableScanScriptsLoadedInInternetExplorer property value. Allows or disallows Windows Defender Script Scanning functionality.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderDisableScanScriptsLoadedInInternetExplorer(value *bool)() {
    err := m.GetBackingStore().Set("defenderDisableScanScriptsLoadedInInternetExplorer", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderEmailContentExecution sets the defenderEmailContentExecution property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEmailContentExecution(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderEmailContentExecution", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderEmailContentExecutionType sets the defenderEmailContentExecutionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEmailContentExecutionType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderEmailContentExecutionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderEnableLowCpuPriority sets the defenderEnableLowCpuPriority property value. This policy setting allows you to enable or disable low CPU priority for scheduled scans.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableLowCpuPriority(value *bool)() {
    err := m.GetBackingStore().Set("defenderEnableLowCpuPriority", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderEnableScanIncomingMail sets the defenderEnableScanIncomingMail property value. Allows or disallows scanning of email.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableScanIncomingMail(value *bool)() {
    err := m.GetBackingStore().Set("defenderEnableScanIncomingMail", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderEnableScanMappedNetworkDrivesDuringFullScan sets the defenderEnableScanMappedNetworkDrivesDuringFullScan property value. Allows or disallows a full scan of mapped network drives.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderEnableScanMappedNetworkDrivesDuringFullScan(value *bool)() {
    err := m.GetBackingStore().Set("defenderEnableScanMappedNetworkDrivesDuringFullScan", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderExploitProtectionXml sets the defenderExploitProtectionXml property value. Xml content containing information regarding exploit protection details.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXml(value []byte)() {
    err := m.GetBackingStore().Set("defenderExploitProtectionXml", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderExploitProtectionXmlFileName sets the defenderExploitProtectionXmlFileName property value. Name of the file from which DefenderExploitProtectionXml was obtained.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXmlFileName(value *string)() {
    err := m.GetBackingStore().Set("defenderExploitProtectionXmlFileName", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderFileExtensionsToExclude sets the defenderFileExtensionsToExclude property value. File extensions to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderFileExtensionsToExclude(value []string)() {
    err := m.GetBackingStore().Set("defenderFileExtensionsToExclude", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderFilesAndFoldersToExclude sets the defenderFilesAndFoldersToExclude property value. Files and folder to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderFilesAndFoldersToExclude(value []string)() {
    err := m.GetBackingStore().Set("defenderFilesAndFoldersToExclude", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderGuardedFoldersAllowedAppPaths sets the defenderGuardedFoldersAllowedAppPaths property value. List of paths to exe that are allowed to access protected folders
func (m *Windows10EndpointProtectionConfiguration) SetDefenderGuardedFoldersAllowedAppPaths(value []string)() {
    err := m.GetBackingStore().Set("defenderGuardedFoldersAllowedAppPaths", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderGuardMyFoldersType sets the defenderGuardMyFoldersType property value. Possible values of Folder Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderGuardMyFoldersType(value *FolderProtectionType)() {
    err := m.GetBackingStore().Set("defenderGuardMyFoldersType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderNetworkProtectionType sets the defenderNetworkProtectionType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderNetworkProtectionType(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderNetworkProtectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeAppsExecutableContentCreationOrLaunch sets the defenderOfficeAppsExecutableContentCreationOrLaunch property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsExecutableContentCreationOrLaunch(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderOfficeAppsExecutableContentCreationOrLaunch", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeAppsExecutableContentCreationOrLaunchType sets the defenderOfficeAppsExecutableContentCreationOrLaunchType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsExecutableContentCreationOrLaunchType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderOfficeAppsExecutableContentCreationOrLaunchType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeAppsLaunchChildProcess sets the defenderOfficeAppsLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsLaunchChildProcess(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderOfficeAppsLaunchChildProcess", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeAppsLaunchChildProcessType sets the defenderOfficeAppsLaunchChildProcessType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsLaunchChildProcessType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderOfficeAppsLaunchChildProcessType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeAppsOtherProcessInjection sets the defenderOfficeAppsOtherProcessInjection property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsOtherProcessInjection(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderOfficeAppsOtherProcessInjection", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeAppsOtherProcessInjectionType sets the defenderOfficeAppsOtherProcessInjectionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeAppsOtherProcessInjectionType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderOfficeAppsOtherProcessInjectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeCommunicationAppsLaunchChildProcess sets the defenderOfficeCommunicationAppsLaunchChildProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeCommunicationAppsLaunchChildProcess(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderOfficeCommunicationAppsLaunchChildProcess", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeMacroCodeAllowWin32Imports sets the defenderOfficeMacroCodeAllowWin32Imports property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeMacroCodeAllowWin32Imports(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderOfficeMacroCodeAllowWin32Imports", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderOfficeMacroCodeAllowWin32ImportsType sets the defenderOfficeMacroCodeAllowWin32ImportsType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderOfficeMacroCodeAllowWin32ImportsType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderOfficeMacroCodeAllowWin32ImportsType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderPotentiallyUnwantedAppAction sets the defenderPotentiallyUnwantedAppAction property value. Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs). Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderPotentiallyUnwantedAppAction(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderPotentiallyUnwantedAppAction", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderPreventCredentialStealingType sets the defenderPreventCredentialStealingType property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderPreventCredentialStealingType(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderPreventCredentialStealingType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderProcessCreation sets the defenderProcessCreation property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessCreation(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderProcessCreation", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderProcessCreationType sets the defenderProcessCreationType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessCreationType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderProcessCreationType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderProcessesToExclude sets the defenderProcessesToExclude property value. Processes to exclude from scans and real time protection.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderProcessesToExclude(value []string)() {
    err := m.GetBackingStore().Set("defenderProcessesToExclude", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScanDirection sets the defenderScanDirection property value. Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly, monitorOutgoingFilesOnly.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanDirection(value *DefenderRealtimeScanDirection)() {
    err := m.GetBackingStore().Set("defenderScanDirection", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScanMaxCpuPercentage sets the defenderScanMaxCpuPercentage property value. Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid values 0 to 100
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanMaxCpuPercentage(value *int32)() {
    err := m.GetBackingStore().Set("defenderScanMaxCpuPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScanType sets the defenderScanType property value. Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScanType(value *DefenderScanType)() {
    err := m.GetBackingStore().Set("defenderScanType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScheduledQuickScanTime sets the defenderScheduledQuickScanTime property value. Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledQuickScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("defenderScheduledQuickScanTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScheduledScanDay sets the defenderScheduledScanDay property value. Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday, monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledScanDay(value *WeeklySchedule)() {
    err := m.GetBackingStore().Set("defenderScheduledScanDay", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScheduledScanTime sets the defenderScheduledScanTime property value. Selects the time of day that the Windows Defender scan should run.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScheduledScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("defenderScheduledScanTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScriptDownloadedPayloadExecution sets the defenderScriptDownloadedPayloadExecution property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptDownloadedPayloadExecution(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderScriptDownloadedPayloadExecution", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScriptDownloadedPayloadExecutionType sets the defenderScriptDownloadedPayloadExecutionType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptDownloadedPayloadExecutionType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderScriptDownloadedPayloadExecutionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScriptObfuscatedMacroCode sets the defenderScriptObfuscatedMacroCode property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptObfuscatedMacroCode(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderScriptObfuscatedMacroCode", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderScriptObfuscatedMacroCodeType sets the defenderScriptObfuscatedMacroCodeType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderScriptObfuscatedMacroCodeType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderScriptObfuscatedMacroCodeType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterBlockExploitProtectionOverride sets the defenderSecurityCenterBlockExploitProtectionOverride property value. Indicates whether or not to block user from overriding Exploit Protection settings.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterBlockExploitProtectionOverride(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterBlockExploitProtectionOverride", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableAccountUI sets the defenderSecurityCenterDisableAccountUI property value. Used to disable the display of the account protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableAccountUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableAccountUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableAppBrowserUI sets the defenderSecurityCenterDisableAppBrowserUI property value. Used to disable the display of the app and browser protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableAppBrowserUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableAppBrowserUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableClearTpmUI sets the defenderSecurityCenterDisableClearTpmUI property value. Used to disable the display of the Clear TPM button.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableClearTpmUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableClearTpmUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableFamilyUI sets the defenderSecurityCenterDisableFamilyUI property value. Used to disable the display of the family options area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableFamilyUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableFamilyUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableHardwareUI sets the defenderSecurityCenterDisableHardwareUI property value. Used to disable the display of the hardware protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableHardwareUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableHardwareUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableHealthUI sets the defenderSecurityCenterDisableHealthUI property value. Used to disable the display of the device performance and health area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableHealthUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableHealthUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableNetworkUI sets the defenderSecurityCenterDisableNetworkUI property value. Used to disable the display of the firewall and network protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableNetworkUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableNetworkUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableNotificationAreaUI sets the defenderSecurityCenterDisableNotificationAreaUI property value. Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot the computer for this setting to take effect.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableNotificationAreaUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableNotificationAreaUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableRansomwareUI sets the defenderSecurityCenterDisableRansomwareUI property value. Used to disable the display of the ransomware protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableRansomwareUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableRansomwareUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableSecureBootUI sets the defenderSecurityCenterDisableSecureBootUI property value. Used to disable the display of the secure boot area under Device security.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableSecureBootUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableSecureBootUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableTroubleshootingUI sets the defenderSecurityCenterDisableTroubleshootingUI property value. Used to disable the display of the security process troubleshooting under Device security.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableTroubleshootingUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableTroubleshootingUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableVirusUI sets the defenderSecurityCenterDisableVirusUI property value. Used to disable the display of the virus and threat protection area.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableVirusUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableVirusUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI sets the defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI property value. Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI(value *bool)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterHelpEmail sets the defenderSecurityCenterHelpEmail property value. The email address that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpEmail(value *string)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterHelpEmail", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterHelpPhone sets the defenderSecurityCenterHelpPhone property value. The phone number or Skype ID that is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpPhone(value *string)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterHelpPhone", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterHelpURL sets the defenderSecurityCenterHelpURL property value. The help portal URL this is displayed to users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterHelpURL(value *string)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterHelpURL", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterITContactDisplay sets the defenderSecurityCenterITContactDisplay property value. Possible values for defenderSecurityCenterITContactDisplay
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterITContactDisplay(value *DefenderSecurityCenterITContactDisplayType)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterITContactDisplay", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterNotificationsFromApp sets the defenderSecurityCenterNotificationsFromApp property value. Possible values for defenderSecurityCenterNotificationsFromApp
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterNotificationsFromApp(value *DefenderSecurityCenterNotificationsFromAppType)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterNotificationsFromApp", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSecurityCenterOrganizationDisplayName sets the defenderSecurityCenterOrganizationDisplayName property value. The company name that is displayed to the users.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterOrganizationDisplayName(value *string)() {
    err := m.GetBackingStore().Set("defenderSecurityCenterOrganizationDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSignatureUpdateIntervalInHours sets the defenderSignatureUpdateIntervalInHours property value. Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSignatureUpdateIntervalInHours(value *int32)() {
    err := m.GetBackingStore().Set("defenderSignatureUpdateIntervalInHours", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderSubmitSamplesConsentType sets the defenderSubmitSamplesConsentType property value. Checks for the user consent level in Windows Defender to send data. Possible values are: sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSubmitSamplesConsentType(value *DefenderSubmitSamplesConsentType)() {
    err := m.GetBackingStore().Set("defenderSubmitSamplesConsentType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderUntrustedExecutable sets the defenderUntrustedExecutable property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedExecutable(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderUntrustedExecutable", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderUntrustedExecutableType sets the defenderUntrustedExecutableType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedExecutableType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderUntrustedExecutableType", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderUntrustedUSBProcess sets the defenderUntrustedUSBProcess property value. Possible values of Defender PUA Protection
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedUSBProcess(value *DefenderProtectionType)() {
    err := m.GetBackingStore().Set("defenderUntrustedUSBProcess", value)
    if err != nil {
        panic(err)
    }
}
// SetDefenderUntrustedUSBProcessType sets the defenderUntrustedUSBProcessType property value. Possible values of Defender Attack Surface Reduction Rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderUntrustedUSBProcessType(value *DefenderAttackSurfaceType)() {
    err := m.GetBackingStore().Set("defenderUntrustedUSBProcessType", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGuardEnableSecureBootWithDMA sets the deviceGuardEnableSecureBootWithDMA property value. This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA. Specifies whether Platform Security Level is enabled at next reboot.
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardEnableSecureBootWithDMA(value *bool)() {
    err := m.GetBackingStore().Set("deviceGuardEnableSecureBootWithDMA", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGuardEnableVirtualizationBasedSecurity sets the deviceGuardEnableVirtualizationBasedSecurity property value. Turns On Virtualization Based Security(VBS).
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardEnableVirtualizationBasedSecurity(value *bool)() {
    err := m.GetBackingStore().Set("deviceGuardEnableVirtualizationBasedSecurity", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGuardLaunchSystemGuard sets the deviceGuardLaunchSystemGuard property value. Possible values of a property
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardLaunchSystemGuard(value *Enablement)() {
    err := m.GetBackingStore().Set("deviceGuardLaunchSystemGuard", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings sets the deviceGuardLocalSystemAuthorityCredentialGuardSettings property value. Possible values of Credential Guard settings.
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings(value *DeviceGuardLocalSystemAuthorityCredentialGuardType)() {
    err := m.GetBackingStore().Set("deviceGuardLocalSystemAuthorityCredentialGuardSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGuardSecureBootWithDMA sets the deviceGuardSecureBootWithDMA property value. Possible values of Secure Boot with DMA
func (m *Windows10EndpointProtectionConfiguration) SetDeviceGuardSecureBootWithDMA(value *SecureBootWithDMAType)() {
    err := m.GetBackingStore().Set("deviceGuardSecureBootWithDMA", value)
    if err != nil {
        panic(err)
    }
}
// SetDmaGuardDeviceEnumerationPolicy sets the dmaGuardDeviceEnumerationPolicy property value. Possible values of the DmaGuardDeviceEnumerationPolicy.
func (m *Windows10EndpointProtectionConfiguration) SetDmaGuardDeviceEnumerationPolicy(value *DmaGuardDeviceEnumerationPolicyType)() {
    err := m.GetBackingStore().Set("dmaGuardDeviceEnumerationPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallBlockStatefulFTP sets the firewallBlockStatefulFTP property value. Blocks stateful FTP connections to the device
func (m *Windows10EndpointProtectionConfiguration) SetFirewallBlockStatefulFTP(value *bool)() {
    err := m.GetBackingStore().Set("firewallBlockStatefulFTP", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallCertificateRevocationListCheckMethod sets the firewallCertificateRevocationListCheckMethod property value. Possible values for firewallCertificateRevocationListCheckMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallCertificateRevocationListCheckMethod(value *FirewallCertificateRevocationListCheckMethodType)() {
    err := m.GetBackingStore().Set("firewallCertificateRevocationListCheckMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallIdleTimeoutForSecurityAssociationInSeconds sets the firewallIdleTimeoutForSecurityAssociationInSeconds property value. Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIdleTimeoutForSecurityAssociationInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("firewallIdleTimeoutForSecurityAssociationInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallIPSecExemptionsAllowDHCP sets the firewallIPSecExemptionsAllowDHCP property value. Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowDHCP(value *bool)() {
    err := m.GetBackingStore().Set("firewallIPSecExemptionsAllowDHCP", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallIPSecExemptionsAllowICMP sets the firewallIPSecExemptionsAllowICMP property value. Configures IPSec exemptions to allow ICMP
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowICMP(value *bool)() {
    err := m.GetBackingStore().Set("firewallIPSecExemptionsAllowICMP", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallIPSecExemptionsAllowNeighborDiscovery sets the firewallIPSecExemptionsAllowNeighborDiscovery property value. Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowNeighborDiscovery(value *bool)() {
    err := m.GetBackingStore().Set("firewallIPSecExemptionsAllowNeighborDiscovery", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallIPSecExemptionsAllowRouterDiscovery sets the firewallIPSecExemptionsAllowRouterDiscovery property value. Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowRouterDiscovery(value *bool)() {
    err := m.GetBackingStore().Set("firewallIPSecExemptionsAllowRouterDiscovery", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallIPSecExemptionsNone sets the firewallIPSecExemptionsNone property value. Configures IPSec exemptions to no exemptions
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsNone(value *bool)() {
    err := m.GetBackingStore().Set("firewallIPSecExemptionsNone", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallMergeKeyingModuleSettings sets the firewallMergeKeyingModuleSettings property value. If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
func (m *Windows10EndpointProtectionConfiguration) SetFirewallMergeKeyingModuleSettings(value *bool)() {
    err := m.GetBackingStore().Set("firewallMergeKeyingModuleSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallPacketQueueingMethod sets the firewallPacketQueueingMethod property value. Possible values for firewallPacketQueueingMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallPacketQueueingMethod(value *FirewallPacketQueueingMethodType)() {
    err := m.GetBackingStore().Set("firewallPacketQueueingMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallPreSharedKeyEncodingMethod sets the firewallPreSharedKeyEncodingMethod property value. Possible values for firewallPreSharedKeyEncodingMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallPreSharedKeyEncodingMethod(value *FirewallPreSharedKeyEncodingMethodType)() {
    err := m.GetBackingStore().Set("firewallPreSharedKeyEncodingMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallProfileDomain sets the firewallProfileDomain property value. Configures the firewall profile settings for domain networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfileDomain(value WindowsFirewallNetworkProfileable)() {
    err := m.GetBackingStore().Set("firewallProfileDomain", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallProfilePrivate sets the firewallProfilePrivate property value. Configures the firewall profile settings for private networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfilePrivate(value WindowsFirewallNetworkProfileable)() {
    err := m.GetBackingStore().Set("firewallProfilePrivate", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallProfilePublic sets the firewallProfilePublic property value. Configures the firewall profile settings for public networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfilePublic(value WindowsFirewallNetworkProfileable)() {
    err := m.GetBackingStore().Set("firewallProfilePublic", value)
    if err != nil {
        panic(err)
    }
}
// SetFirewallRules sets the firewallRules property value. Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
func (m *Windows10EndpointProtectionConfiguration) SetFirewallRules(value []WindowsFirewallRuleable)() {
    err := m.GetBackingStore().Set("firewallRules", value)
    if err != nil {
        panic(err)
    }
}
// SetLanManagerAuthenticationLevel sets the lanManagerAuthenticationLevel property value. Possible values for LanManagerAuthenticationLevel
func (m *Windows10EndpointProtectionConfiguration) SetLanManagerAuthenticationLevel(value *LanManagerAuthenticationLevel)() {
    err := m.GetBackingStore().Set("lanManagerAuthenticationLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetLanManagerWorkstationDisableInsecureGuestLogons sets the lanManagerWorkstationDisableInsecureGuestLogons property value. If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure guest logons.
func (m *Windows10EndpointProtectionConfiguration) SetLanManagerWorkstationDisableInsecureGuestLogons(value *bool)() {
    err := m.GetBackingStore().Set("lanManagerWorkstationDisableInsecureGuestLogons", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAdministratorAccountName sets the localSecurityOptionsAdministratorAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAdministratorAccountName(value *string)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAdministratorAccountName", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAdministratorElevationPromptBehavior sets the localSecurityOptionsAdministratorElevationPromptBehavior property value. Possible values for LocalSecurityOptionsAdministratorElevationPromptBehavior
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAdministratorElevationPromptBehavior(value *LocalSecurityOptionsAdministratorElevationPromptBehaviorType)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAdministratorElevationPromptBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares sets the localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares property value. This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating the names of domain accounts and network shares.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests sets the localSecurityOptionsAllowPKU2UAuthenticationRequests property value. Block PKU2U authentication requests to this device to use online identities.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAllowPKU2UAuthenticationRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager sets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager property value. Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote calls to the SAM.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager(value *string)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool sets the localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool property value. UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn sets the localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn property value. This security setting determines whether a computer can be shut down without having to log on to Windows.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAllowUIAccessApplicationElevation sets the localSecurityOptionsAllowUIAccessApplicationElevation property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUIAccessApplicationElevation(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAllowUIAccessApplicationElevation", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations sets the localSecurityOptionsAllowUIAccessApplicationsForSecureLocations property value. Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAllowUIAccessApplicationsForSecureLocations", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon sets the localSecurityOptionsAllowUndockWithoutHavingToLogon property value. Prevent a portable computer from being undocked without having to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsAllowUndockWithoutHavingToLogon", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsBlockMicrosoftAccounts sets the localSecurityOptionsBlockMicrosoftAccounts property value. Prevent users from adding new Microsoft accounts to this computer.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockMicrosoftAccounts(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsBlockMicrosoftAccounts", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword sets the localSecurityOptionsBlockRemoteLogonWithBlankPassword property value. Enable Local accounts that are not password protected to log on from locations other than the physical device.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsBlockRemoteLogonWithBlankPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess sets the localSecurityOptionsBlockRemoteOpticalDriveAccess property value. Enabling this settings allows only interactively logged on user to access CD-ROM media.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsBlockRemoteOpticalDriveAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers sets the localSecurityOptionsBlockUsersInstallingPrinterDrivers property value. Restrict installing printer drivers as part of connecting to a shared printer to admins only.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsBlockUsersInstallingPrinterDrivers", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsClearVirtualMemoryPageFile sets the localSecurityOptionsClearVirtualMemoryPageFile property value. This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClearVirtualMemoryPageFile(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsClearVirtualMemoryPageFile", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways sets the localSecurityOptionsClientDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB client component.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsClientDigitallySignCommunicationsAlways", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers sets the localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers property value. If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords to non-Microsoft SMB servers that do not support password encryption during authentication.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation sets the localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation property value. App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDisableAdministratorAccount sets the localSecurityOptionsDisableAdministratorAccount property value. Determines whether the Local Administrator account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableAdministratorAccount(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDisableAdministratorAccount", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees sets the localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees property value. This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDisableGuestAccount sets the localSecurityOptionsDisableGuestAccount property value. Determines if the Guest account is enabled or disabled.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableGuestAccount(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDisableGuestAccount", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways sets the localSecurityOptionsDisableServerDigitallySignCommunicationsAlways property value. This security setting determines whether packet signing is required by the SMB server component.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDisableServerDigitallySignCommunicationsAlways", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees sets the localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees property value. This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request it.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts sets the localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts property value. This security setting determines what additional permissions will be granted for anonymous connections to the computer.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDoNotRequireCtrlAltDel sets the localSecurityOptionsDoNotRequireCtrlAltDel property value. Require CTRL+ALT+DEL to be pressed before a user can log on.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotRequireCtrlAltDel(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDoNotRequireCtrlAltDel", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange sets the localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange property value. This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new password is stored. It’s not stored by default.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser sets the localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser property value. Possible values for LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(value *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType)() {
    err := m.GetBackingStore().Set("localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsGuestAccountName sets the localSecurityOptionsGuestAccountName property value. Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsGuestAccountName(value *string)() {
    err := m.GetBackingStore().Set("localSecurityOptionsGuestAccountName", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsHideLastSignedInUser sets the localSecurityOptionsHideLastSignedInUser property value. Do not display the username of the last person who signed in on this device.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsHideLastSignedInUser(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsHideLastSignedInUser", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsHideUsernameAtSignIn sets the localSecurityOptionsHideUsernameAtSignIn property value. Do not display the username of the person signing in to this device after credentials are entered and before the device’s desktop is shown.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsHideUsernameAtSignIn(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsHideUsernameAtSignIn", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsInformationDisplayedOnLockScreen sets the localSecurityOptionsInformationDisplayedOnLockScreen property value. Possible values for LocalSecurityOptionsInformationDisplayedOnLockScreen
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsInformationDisplayedOnLockScreen(value *LocalSecurityOptionsInformationDisplayedOnLockScreenType)() {
    err := m.GetBackingStore().Set("localSecurityOptionsInformationDisplayedOnLockScreen", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsInformationShownOnLockScreen sets the localSecurityOptionsInformationShownOnLockScreen property value. Possible values for LocalSecurityOptionsInformationShownOnLockScreenType
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsInformationShownOnLockScreen(value *LocalSecurityOptionsInformationShownOnLockScreenType)() {
    err := m.GetBackingStore().Set("localSecurityOptionsInformationShownOnLockScreen", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsLogOnMessageText sets the localSecurityOptionsLogOnMessageText property value. Set message text for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsLogOnMessageText(value *string)() {
    err := m.GetBackingStore().Set("localSecurityOptionsLogOnMessageText", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsLogOnMessageTitle sets the localSecurityOptionsLogOnMessageTitle property value. Set message title for users attempting to log in.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsLogOnMessageTitle(value *string)() {
    err := m.GetBackingStore().Set("localSecurityOptionsLogOnMessageTitle", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsMachineInactivityLimit sets the localSecurityOptionsMachineInactivityLimit property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMachineInactivityLimit(value *int32)() {
    err := m.GetBackingStore().Set("localSecurityOptionsMachineInactivityLimit", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsMachineInactivityLimitInMinutes sets the localSecurityOptionsMachineInactivityLimitInMinutes property value. Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid values 0 to 9999
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMachineInactivityLimitInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("localSecurityOptionsMachineInactivityLimitInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients sets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients property value. Possible values for LocalSecurityOptionsMinimumSessionSecurity
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(value *LocalSecurityOptionsMinimumSessionSecurity)() {
    err := m.GetBackingStore().Set("localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers sets the localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers property value. Possible values for LocalSecurityOptionsMinimumSessionSecurity
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(value *LocalSecurityOptionsMinimumSessionSecurity)() {
    err := m.GetBackingStore().Set("localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsOnlyElevateSignedExecutables sets the localSecurityOptionsOnlyElevateSignedExecutables property value. Enforce PKI certification path validation for a given executable file before it is permitted to run.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsOnlyElevateSignedExecutables(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsOnlyElevateSignedExecutables", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares sets the localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares property value. By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that can be accessed anonymously and Shares that can be accessed anonymously
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsSmartCardRemovalBehavior sets the localSecurityOptionsSmartCardRemovalBehavior property value. Possible values for LocalSecurityOptionsSmartCardRemovalBehaviorType
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsSmartCardRemovalBehavior(value *LocalSecurityOptionsSmartCardRemovalBehaviorType)() {
    err := m.GetBackingStore().Set("localSecurityOptionsSmartCardRemovalBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsStandardUserElevationPromptBehavior sets the localSecurityOptionsStandardUserElevationPromptBehavior property value. Possible values for LocalSecurityOptionsStandardUserElevationPromptBehavior
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsStandardUserElevationPromptBehavior(value *LocalSecurityOptionsStandardUserElevationPromptBehaviorType)() {
    err := m.GetBackingStore().Set("localSecurityOptionsStandardUserElevationPromptBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation sets the localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation property value. Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior policy settings for admins and standard users are used.
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsUseAdminApprovalMode sets the localSecurityOptionsUseAdminApprovalMode property value. Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin privileges.Default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsUseAdminApprovalMode(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsUseAdminApprovalMode", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators sets the localSecurityOptionsUseAdminApprovalModeForAdministrators property value. Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsUseAdminApprovalModeForAdministrators", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations sets the localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations property value. Virtualize file and registry write failures to per user locations
func (m *Windows10EndpointProtectionConfiguration) SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations(value *bool)() {
    err := m.GetBackingStore().Set("localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations", value)
    if err != nil {
        panic(err)
    }
}
// SetSmartScreenBlockOverrideForFiles sets the smartScreenBlockOverrideForFiles property value. Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
func (m *Windows10EndpointProtectionConfiguration) SetSmartScreenBlockOverrideForFiles(value *bool)() {
    err := m.GetBackingStore().Set("smartScreenBlockOverrideForFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetSmartScreenEnableInShell sets the smartScreenEnableInShell property value. Allows IT Admins to configure SmartScreen for Windows.
func (m *Windows10EndpointProtectionConfiguration) SetSmartScreenEnableInShell(value *bool)() {
    err := m.GetBackingStore().Set("smartScreenEnableInShell", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsAccessCredentialManagerAsTrustedCaller sets the userRightsAccessCredentialManagerAsTrustedCaller property value. This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if this privilege is given to other entities. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsAccessCredentialManagerAsTrustedCaller(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsAccessCredentialManagerAsTrustedCaller", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsActAsPartOfTheOperatingSystem sets the userRightsActAsPartOfTheOperatingSystem property value. This user right allows a process to impersonate any user without authentication. The process can therefore gain access to the same local resources as that user. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsActAsPartOfTheOperatingSystem(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsActAsPartOfTheOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsAllowAccessFromNetwork sets the userRightsAllowAccessFromNetwork property value. This user right determines which users and groups are allowed to connect to the computer over the network. State Allowed is supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsAllowAccessFromNetwork(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsAllowAccessFromNetwork", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsBackupData sets the userRightsBackupData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when backing up files and directories. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsBackupData(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsBackupData", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsBlockAccessFromNetwork sets the userRightsBlockAccessFromNetwork property value. This user right determines which users and groups are block from connecting to the computer over the network. State Block is supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsBlockAccessFromNetwork(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsBlockAccessFromNetwork", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsChangeSystemTime sets the userRightsChangeSystemTime property value. This user right determines which users and groups can change the time and date on the internal clock of the computer. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsChangeSystemTime(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsChangeSystemTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsCreateGlobalObjects sets the userRightsCreateGlobalObjects property value. This security setting determines whether users can create global objects that are available to all sessions. Users who can create global objects could affect processes that run under other users' sessions, which could lead to application failure or data corruption. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateGlobalObjects(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsCreateGlobalObjects", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsCreatePageFile sets the userRightsCreatePageFile property value. This user right determines which users and groups can call an internal API to create and change the size of a page file. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreatePageFile(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsCreatePageFile", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsCreatePermanentSharedObjects sets the userRightsCreatePermanentSharedObjects property value. This user right determines which accounts can be used by processes to create a directory object using the object manager. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreatePermanentSharedObjects(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsCreatePermanentSharedObjects", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsCreateSymbolicLinks sets the userRightsCreateSymbolicLinks property value. This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateSymbolicLinks(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsCreateSymbolicLinks", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsCreateToken sets the userRightsCreateToken property value. This user right determines which users/groups can be used by processes to create a token that can then be used to get access to any local resources when the process uses an internal API to create an access token. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsCreateToken(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsCreateToken", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsDebugPrograms sets the userRightsDebugPrograms property value. This user right determines which users can attach a debugger to any process or to the kernel. Only states NotConfigured and Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDebugPrograms(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsDebugPrograms", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsDelegation sets the userRightsDelegation property value. This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDelegation(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsDelegation", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsDenyLocalLogOn sets the userRightsDenyLocalLogOn property value. This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsDenyLocalLogOn(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsDenyLocalLogOn", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsGenerateSecurityAudits sets the userRightsGenerateSecurityAudits property value. This user right determines which accounts can be used by a process to add entries to the security log. The security log is used to trace unauthorized system access.  Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsGenerateSecurityAudits(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsGenerateSecurityAudits", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsImpersonateClient sets the userRightsImpersonateClient property value. Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a service that they have created and then impersonating that client, which can elevate the unauthorized user's permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsImpersonateClient(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsImpersonateClient", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsIncreaseSchedulingPriority sets the userRightsIncreaseSchedulingPriority property value. This user right determines which accounts can use a process with Write Property access to another process to increase the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsIncreaseSchedulingPriority(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsIncreaseSchedulingPriority", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsLoadUnloadDrivers sets the userRightsLoadUnloadDrivers property value. This user right determines which users can dynamically load and unload device drivers or other code in to kernel mode. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLoadUnloadDrivers(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsLoadUnloadDrivers", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsLocalLogOn sets the userRightsLocalLogOn property value. This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLocalLogOn(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsLocalLogOn", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsLockMemory sets the userRightsLockMemory property value. This user right determines which accounts can use a process to keep data in physical memory, which prevents the system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsLockMemory(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsLockMemory", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsManageAuditingAndSecurityLogs sets the userRightsManageAuditingAndSecurityLogs property value. This user right determines which users can specify object access auditing options for individual resources, such as files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsManageAuditingAndSecurityLogs(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsManageAuditingAndSecurityLogs", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsManageVolumes sets the userRightsManageVolumes property value. This user right determines which users and groups can run maintenance tasks on a volume, such as remote defragmentation. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsManageVolumes(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsManageVolumes", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsModifyFirmwareEnvironment sets the userRightsModifyFirmwareEnvironment property value. This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsModifyFirmwareEnvironment(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsModifyFirmwareEnvironment", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsModifyObjectLabels sets the userRightsModifyObjectLabels property value. This user right determines which user accounts can modify the integrity label of objects, such as files, registry keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsModifyObjectLabels(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsModifyObjectLabels", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsProfileSingleProcess sets the userRightsProfileSingleProcess property value. This user right determines which users can use performance monitoring tools to monitor the performance of system processes. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsProfileSingleProcess(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsProfileSingleProcess", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsRemoteDesktopServicesLogOn sets the userRightsRemoteDesktopServicesLogOn property value. This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client. Only states NotConfigured and Blocked are supported
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRemoteDesktopServicesLogOn(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsRemoteDesktopServicesLogOn", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsRemoteShutdown sets the userRightsRemoteShutdown property value. This user right determines which users are allowed to shut down a computer from a remote location on the network. Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRemoteShutdown(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsRemoteShutdown", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsRestoreData sets the userRightsRestoreData property value. This user right determines which users can bypass file, directory, registry, and other persistent objects permissions when restoring backed up files and directories, and determines which users can set any valid security principal as the owner of an object. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsRestoreData(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsRestoreData", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRightsTakeOwnership sets the userRightsTakeOwnership property value. This user right determines which users can take ownership of any securable object in the system, including Active Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and Allowed are supported.
func (m *Windows10EndpointProtectionConfiguration) SetUserRightsTakeOwnership(value DeviceManagementUserRightsSettingable)() {
    err := m.GetBackingStore().Set("userRightsTakeOwnership", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsDefenderTamperProtection sets the windowsDefenderTamperProtection property value. Defender TamperProtection setting options
func (m *Windows10EndpointProtectionConfiguration) SetWindowsDefenderTamperProtection(value *WindowsDefenderTamperProtectionOptions)() {
    err := m.GetBackingStore().Set("windowsDefenderTamperProtection", value)
    if err != nil {
        panic(err)
    }
}
// SetXboxServicesAccessoryManagementServiceStartupMode sets the xboxServicesAccessoryManagementServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesAccessoryManagementServiceStartupMode(value *ServiceStartType)() {
    err := m.GetBackingStore().Set("xboxServicesAccessoryManagementServiceStartupMode", value)
    if err != nil {
        panic(err)
    }
}
// SetXboxServicesEnableXboxGameSaveTask sets the xboxServicesEnableXboxGameSaveTask property value. This setting determines whether xbox game save is enabled (1) or disabled (0).
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesEnableXboxGameSaveTask(value *bool)() {
    err := m.GetBackingStore().Set("xboxServicesEnableXboxGameSaveTask", value)
    if err != nil {
        panic(err)
    }
}
// SetXboxServicesLiveAuthManagerServiceStartupMode sets the xboxServicesLiveAuthManagerServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveAuthManagerServiceStartupMode(value *ServiceStartType)() {
    err := m.GetBackingStore().Set("xboxServicesLiveAuthManagerServiceStartupMode", value)
    if err != nil {
        panic(err)
    }
}
// SetXboxServicesLiveGameSaveServiceStartupMode sets the xboxServicesLiveGameSaveServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveGameSaveServiceStartupMode(value *ServiceStartType)() {
    err := m.GetBackingStore().Set("xboxServicesLiveGameSaveServiceStartupMode", value)
    if err != nil {
        panic(err)
    }
}
// SetXboxServicesLiveNetworkingServiceStartupMode sets the xboxServicesLiveNetworkingServiceStartupMode property value. Possible values of xbox service start type
func (m *Windows10EndpointProtectionConfiguration) SetXboxServicesLiveNetworkingServiceStartupMode(value *ServiceStartType)() {
    err := m.GetBackingStore().Set("xboxServicesLiveNetworkingServiceStartupMode", value)
    if err != nil {
        panic(err)
    }
}
type Windows10EndpointProtectionConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationGuardAllowCameraMicrophoneRedirection()(*bool)
    GetApplicationGuardAllowFileSaveOnHost()(*bool)
    GetApplicationGuardAllowPersistence()(*bool)
    GetApplicationGuardAllowPrintToLocalPrinters()(*bool)
    GetApplicationGuardAllowPrintToNetworkPrinters()(*bool)
    GetApplicationGuardAllowPrintToPDF()(*bool)
    GetApplicationGuardAllowPrintToXPS()(*bool)
    GetApplicationGuardAllowVirtualGPU()(*bool)
    GetApplicationGuardBlockClipboardSharing()(*ApplicationGuardBlockClipboardSharingType)
    GetApplicationGuardBlockFileTransfer()(*ApplicationGuardBlockFileTransferType)
    GetApplicationGuardBlockNonEnterpriseContent()(*bool)
    GetApplicationGuardCertificateThumbprints()([]string)
    GetApplicationGuardEnabled()(*bool)
    GetApplicationGuardEnabledOptions()(*ApplicationGuardEnabledOptions)
    GetApplicationGuardForceAuditing()(*bool)
    GetAppLockerApplicationControl()(*AppLockerApplicationControlType)
    GetBitLockerAllowStandardUserEncryption()(*bool)
    GetBitLockerDisableWarningForOtherDiskEncryption()(*bool)
    GetBitLockerEnableStorageCardEncryptionOnMobile()(*bool)
    GetBitLockerEncryptDevice()(*bool)
    GetBitLockerFixedDrivePolicy()(BitLockerFixedDrivePolicyable)
    GetBitLockerRecoveryPasswordRotation()(*BitLockerRecoveryPasswordRotationType)
    GetBitLockerRemovableDrivePolicy()(BitLockerRemovableDrivePolicyable)
    GetBitLockerSystemDrivePolicy()(BitLockerSystemDrivePolicyable)
    GetDefenderAdditionalGuardedFolders()([]string)
    GetDefenderAdobeReaderLaunchChildProcess()(*DefenderProtectionType)
    GetDefenderAdvancedRansomewareProtectionType()(*DefenderProtectionType)
    GetDefenderAllowBehaviorMonitoring()(*bool)
    GetDefenderAllowCloudProtection()(*bool)
    GetDefenderAllowEndUserAccess()(*bool)
    GetDefenderAllowIntrusionPreventionSystem()(*bool)
    GetDefenderAllowOnAccessProtection()(*bool)
    GetDefenderAllowRealTimeMonitoring()(*bool)
    GetDefenderAllowScanArchiveFiles()(*bool)
    GetDefenderAllowScanDownloads()(*bool)
    GetDefenderAllowScanNetworkFiles()(*bool)
    GetDefenderAllowScanRemovableDrivesDuringFullScan()(*bool)
    GetDefenderAllowScanScriptsLoadedInInternetExplorer()(*bool)
    GetDefenderAttackSurfaceReductionExcludedPaths()([]string)
    GetDefenderBlockEndUserAccess()(*bool)
    GetDefenderBlockPersistenceThroughWmiType()(*DefenderAttackSurfaceType)
    GetDefenderCheckForSignaturesBeforeRunningScan()(*bool)
    GetDefenderCloudBlockLevel()(*DefenderCloudBlockLevelType)
    GetDefenderCloudExtendedTimeoutInSeconds()(*int32)
    GetDefenderDaysBeforeDeletingQuarantinedMalware()(*int32)
    GetDefenderDetectedMalwareActions()(DefenderDetectedMalwareActionsable)
    GetDefenderDisableBehaviorMonitoring()(*bool)
    GetDefenderDisableCatchupFullScan()(*bool)
    GetDefenderDisableCatchupQuickScan()(*bool)
    GetDefenderDisableCloudProtection()(*bool)
    GetDefenderDisableIntrusionPreventionSystem()(*bool)
    GetDefenderDisableOnAccessProtection()(*bool)
    GetDefenderDisableRealTimeMonitoring()(*bool)
    GetDefenderDisableScanArchiveFiles()(*bool)
    GetDefenderDisableScanDownloads()(*bool)
    GetDefenderDisableScanNetworkFiles()(*bool)
    GetDefenderDisableScanRemovableDrivesDuringFullScan()(*bool)
    GetDefenderDisableScanScriptsLoadedInInternetExplorer()(*bool)
    GetDefenderEmailContentExecution()(*DefenderProtectionType)
    GetDefenderEmailContentExecutionType()(*DefenderAttackSurfaceType)
    GetDefenderEnableLowCpuPriority()(*bool)
    GetDefenderEnableScanIncomingMail()(*bool)
    GetDefenderEnableScanMappedNetworkDrivesDuringFullScan()(*bool)
    GetDefenderExploitProtectionXml()([]byte)
    GetDefenderExploitProtectionXmlFileName()(*string)
    GetDefenderFileExtensionsToExclude()([]string)
    GetDefenderFilesAndFoldersToExclude()([]string)
    GetDefenderGuardedFoldersAllowedAppPaths()([]string)
    GetDefenderGuardMyFoldersType()(*FolderProtectionType)
    GetDefenderNetworkProtectionType()(*DefenderProtectionType)
    GetDefenderOfficeAppsExecutableContentCreationOrLaunch()(*DefenderProtectionType)
    GetDefenderOfficeAppsExecutableContentCreationOrLaunchType()(*DefenderAttackSurfaceType)
    GetDefenderOfficeAppsLaunchChildProcess()(*DefenderProtectionType)
    GetDefenderOfficeAppsLaunchChildProcessType()(*DefenderAttackSurfaceType)
    GetDefenderOfficeAppsOtherProcessInjection()(*DefenderProtectionType)
    GetDefenderOfficeAppsOtherProcessInjectionType()(*DefenderAttackSurfaceType)
    GetDefenderOfficeCommunicationAppsLaunchChildProcess()(*DefenderProtectionType)
    GetDefenderOfficeMacroCodeAllowWin32Imports()(*DefenderProtectionType)
    GetDefenderOfficeMacroCodeAllowWin32ImportsType()(*DefenderAttackSurfaceType)
    GetDefenderPotentiallyUnwantedAppAction()(*DefenderProtectionType)
    GetDefenderPreventCredentialStealingType()(*DefenderProtectionType)
    GetDefenderProcessCreation()(*DefenderProtectionType)
    GetDefenderProcessCreationType()(*DefenderAttackSurfaceType)
    GetDefenderProcessesToExclude()([]string)
    GetDefenderScanDirection()(*DefenderRealtimeScanDirection)
    GetDefenderScanMaxCpuPercentage()(*int32)
    GetDefenderScanType()(*DefenderScanType)
    GetDefenderScheduledQuickScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetDefenderScheduledScanDay()(*WeeklySchedule)
    GetDefenderScheduledScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetDefenderScriptDownloadedPayloadExecution()(*DefenderProtectionType)
    GetDefenderScriptDownloadedPayloadExecutionType()(*DefenderAttackSurfaceType)
    GetDefenderScriptObfuscatedMacroCode()(*DefenderProtectionType)
    GetDefenderScriptObfuscatedMacroCodeType()(*DefenderAttackSurfaceType)
    GetDefenderSecurityCenterBlockExploitProtectionOverride()(*bool)
    GetDefenderSecurityCenterDisableAccountUI()(*bool)
    GetDefenderSecurityCenterDisableAppBrowserUI()(*bool)
    GetDefenderSecurityCenterDisableClearTpmUI()(*bool)
    GetDefenderSecurityCenterDisableFamilyUI()(*bool)
    GetDefenderSecurityCenterDisableHardwareUI()(*bool)
    GetDefenderSecurityCenterDisableHealthUI()(*bool)
    GetDefenderSecurityCenterDisableNetworkUI()(*bool)
    GetDefenderSecurityCenterDisableNotificationAreaUI()(*bool)
    GetDefenderSecurityCenterDisableRansomwareUI()(*bool)
    GetDefenderSecurityCenterDisableSecureBootUI()(*bool)
    GetDefenderSecurityCenterDisableTroubleshootingUI()(*bool)
    GetDefenderSecurityCenterDisableVirusUI()(*bool)
    GetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI()(*bool)
    GetDefenderSecurityCenterHelpEmail()(*string)
    GetDefenderSecurityCenterHelpPhone()(*string)
    GetDefenderSecurityCenterHelpURL()(*string)
    GetDefenderSecurityCenterITContactDisplay()(*DefenderSecurityCenterITContactDisplayType)
    GetDefenderSecurityCenterNotificationsFromApp()(*DefenderSecurityCenterNotificationsFromAppType)
    GetDefenderSecurityCenterOrganizationDisplayName()(*string)
    GetDefenderSignatureUpdateIntervalInHours()(*int32)
    GetDefenderSubmitSamplesConsentType()(*DefenderSubmitSamplesConsentType)
    GetDefenderUntrustedExecutable()(*DefenderProtectionType)
    GetDefenderUntrustedExecutableType()(*DefenderAttackSurfaceType)
    GetDefenderUntrustedUSBProcess()(*DefenderProtectionType)
    GetDefenderUntrustedUSBProcessType()(*DefenderAttackSurfaceType)
    GetDeviceGuardEnableSecureBootWithDMA()(*bool)
    GetDeviceGuardEnableVirtualizationBasedSecurity()(*bool)
    GetDeviceGuardLaunchSystemGuard()(*Enablement)
    GetDeviceGuardLocalSystemAuthorityCredentialGuardSettings()(*DeviceGuardLocalSystemAuthorityCredentialGuardType)
    GetDeviceGuardSecureBootWithDMA()(*SecureBootWithDMAType)
    GetDmaGuardDeviceEnumerationPolicy()(*DmaGuardDeviceEnumerationPolicyType)
    GetFirewallBlockStatefulFTP()(*bool)
    GetFirewallCertificateRevocationListCheckMethod()(*FirewallCertificateRevocationListCheckMethodType)
    GetFirewallIdleTimeoutForSecurityAssociationInSeconds()(*int32)
    GetFirewallIPSecExemptionsAllowDHCP()(*bool)
    GetFirewallIPSecExemptionsAllowICMP()(*bool)
    GetFirewallIPSecExemptionsAllowNeighborDiscovery()(*bool)
    GetFirewallIPSecExemptionsAllowRouterDiscovery()(*bool)
    GetFirewallIPSecExemptionsNone()(*bool)
    GetFirewallMergeKeyingModuleSettings()(*bool)
    GetFirewallPacketQueueingMethod()(*FirewallPacketQueueingMethodType)
    GetFirewallPreSharedKeyEncodingMethod()(*FirewallPreSharedKeyEncodingMethodType)
    GetFirewallProfileDomain()(WindowsFirewallNetworkProfileable)
    GetFirewallProfilePrivate()(WindowsFirewallNetworkProfileable)
    GetFirewallProfilePublic()(WindowsFirewallNetworkProfileable)
    GetFirewallRules()([]WindowsFirewallRuleable)
    GetLanManagerAuthenticationLevel()(*LanManagerAuthenticationLevel)
    GetLanManagerWorkstationDisableInsecureGuestLogons()(*bool)
    GetLocalSecurityOptionsAdministratorAccountName()(*string)
    GetLocalSecurityOptionsAdministratorElevationPromptBehavior()(*LocalSecurityOptionsAdministratorElevationPromptBehaviorType)
    GetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares()(*bool)
    GetLocalSecurityOptionsAllowPKU2UAuthenticationRequests()(*bool)
    GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager()(*string)
    GetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool()(*bool)
    GetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn()(*bool)
    GetLocalSecurityOptionsAllowUIAccessApplicationElevation()(*bool)
    GetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations()(*bool)
    GetLocalSecurityOptionsAllowUndockWithoutHavingToLogon()(*bool)
    GetLocalSecurityOptionsBlockMicrosoftAccounts()(*bool)
    GetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword()(*bool)
    GetLocalSecurityOptionsBlockRemoteOpticalDriveAccess()(*bool)
    GetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers()(*bool)
    GetLocalSecurityOptionsClearVirtualMemoryPageFile()(*bool)
    GetLocalSecurityOptionsClientDigitallySignCommunicationsAlways()(*bool)
    GetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers()(*bool)
    GetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation()(*bool)
    GetLocalSecurityOptionsDisableAdministratorAccount()(*bool)
    GetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees()(*bool)
    GetLocalSecurityOptionsDisableGuestAccount()(*bool)
    GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways()(*bool)
    GetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees()(*bool)
    GetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts()(*bool)
    GetLocalSecurityOptionsDoNotRequireCtrlAltDel()(*bool)
    GetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange()(*bool)
    GetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser()(*LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType)
    GetLocalSecurityOptionsGuestAccountName()(*string)
    GetLocalSecurityOptionsHideLastSignedInUser()(*bool)
    GetLocalSecurityOptionsHideUsernameAtSignIn()(*bool)
    GetLocalSecurityOptionsInformationDisplayedOnLockScreen()(*LocalSecurityOptionsInformationDisplayedOnLockScreenType)
    GetLocalSecurityOptionsInformationShownOnLockScreen()(*LocalSecurityOptionsInformationShownOnLockScreenType)
    GetLocalSecurityOptionsLogOnMessageText()(*string)
    GetLocalSecurityOptionsLogOnMessageTitle()(*string)
    GetLocalSecurityOptionsMachineInactivityLimit()(*int32)
    GetLocalSecurityOptionsMachineInactivityLimitInMinutes()(*int32)
    GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients()(*LocalSecurityOptionsMinimumSessionSecurity)
    GetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers()(*LocalSecurityOptionsMinimumSessionSecurity)
    GetLocalSecurityOptionsOnlyElevateSignedExecutables()(*bool)
    GetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares()(*bool)
    GetLocalSecurityOptionsSmartCardRemovalBehavior()(*LocalSecurityOptionsSmartCardRemovalBehaviorType)
    GetLocalSecurityOptionsStandardUserElevationPromptBehavior()(*LocalSecurityOptionsStandardUserElevationPromptBehaviorType)
    GetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation()(*bool)
    GetLocalSecurityOptionsUseAdminApprovalMode()(*bool)
    GetLocalSecurityOptionsUseAdminApprovalModeForAdministrators()(*bool)
    GetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations()(*bool)
    GetSmartScreenBlockOverrideForFiles()(*bool)
    GetSmartScreenEnableInShell()(*bool)
    GetUserRightsAccessCredentialManagerAsTrustedCaller()(DeviceManagementUserRightsSettingable)
    GetUserRightsActAsPartOfTheOperatingSystem()(DeviceManagementUserRightsSettingable)
    GetUserRightsAllowAccessFromNetwork()(DeviceManagementUserRightsSettingable)
    GetUserRightsBackupData()(DeviceManagementUserRightsSettingable)
    GetUserRightsBlockAccessFromNetwork()(DeviceManagementUserRightsSettingable)
    GetUserRightsChangeSystemTime()(DeviceManagementUserRightsSettingable)
    GetUserRightsCreateGlobalObjects()(DeviceManagementUserRightsSettingable)
    GetUserRightsCreatePageFile()(DeviceManagementUserRightsSettingable)
    GetUserRightsCreatePermanentSharedObjects()(DeviceManagementUserRightsSettingable)
    GetUserRightsCreateSymbolicLinks()(DeviceManagementUserRightsSettingable)
    GetUserRightsCreateToken()(DeviceManagementUserRightsSettingable)
    GetUserRightsDebugPrograms()(DeviceManagementUserRightsSettingable)
    GetUserRightsDelegation()(DeviceManagementUserRightsSettingable)
    GetUserRightsDenyLocalLogOn()(DeviceManagementUserRightsSettingable)
    GetUserRightsGenerateSecurityAudits()(DeviceManagementUserRightsSettingable)
    GetUserRightsImpersonateClient()(DeviceManagementUserRightsSettingable)
    GetUserRightsIncreaseSchedulingPriority()(DeviceManagementUserRightsSettingable)
    GetUserRightsLoadUnloadDrivers()(DeviceManagementUserRightsSettingable)
    GetUserRightsLocalLogOn()(DeviceManagementUserRightsSettingable)
    GetUserRightsLockMemory()(DeviceManagementUserRightsSettingable)
    GetUserRightsManageAuditingAndSecurityLogs()(DeviceManagementUserRightsSettingable)
    GetUserRightsManageVolumes()(DeviceManagementUserRightsSettingable)
    GetUserRightsModifyFirmwareEnvironment()(DeviceManagementUserRightsSettingable)
    GetUserRightsModifyObjectLabels()(DeviceManagementUserRightsSettingable)
    GetUserRightsProfileSingleProcess()(DeviceManagementUserRightsSettingable)
    GetUserRightsRemoteDesktopServicesLogOn()(DeviceManagementUserRightsSettingable)
    GetUserRightsRemoteShutdown()(DeviceManagementUserRightsSettingable)
    GetUserRightsRestoreData()(DeviceManagementUserRightsSettingable)
    GetUserRightsTakeOwnership()(DeviceManagementUserRightsSettingable)
    GetWindowsDefenderTamperProtection()(*WindowsDefenderTamperProtectionOptions)
    GetXboxServicesAccessoryManagementServiceStartupMode()(*ServiceStartType)
    GetXboxServicesEnableXboxGameSaveTask()(*bool)
    GetXboxServicesLiveAuthManagerServiceStartupMode()(*ServiceStartType)
    GetXboxServicesLiveGameSaveServiceStartupMode()(*ServiceStartType)
    GetXboxServicesLiveNetworkingServiceStartupMode()(*ServiceStartType)
    SetApplicationGuardAllowCameraMicrophoneRedirection(value *bool)()
    SetApplicationGuardAllowFileSaveOnHost(value *bool)()
    SetApplicationGuardAllowPersistence(value *bool)()
    SetApplicationGuardAllowPrintToLocalPrinters(value *bool)()
    SetApplicationGuardAllowPrintToNetworkPrinters(value *bool)()
    SetApplicationGuardAllowPrintToPDF(value *bool)()
    SetApplicationGuardAllowPrintToXPS(value *bool)()
    SetApplicationGuardAllowVirtualGPU(value *bool)()
    SetApplicationGuardBlockClipboardSharing(value *ApplicationGuardBlockClipboardSharingType)()
    SetApplicationGuardBlockFileTransfer(value *ApplicationGuardBlockFileTransferType)()
    SetApplicationGuardBlockNonEnterpriseContent(value *bool)()
    SetApplicationGuardCertificateThumbprints(value []string)()
    SetApplicationGuardEnabled(value *bool)()
    SetApplicationGuardEnabledOptions(value *ApplicationGuardEnabledOptions)()
    SetApplicationGuardForceAuditing(value *bool)()
    SetAppLockerApplicationControl(value *AppLockerApplicationControlType)()
    SetBitLockerAllowStandardUserEncryption(value *bool)()
    SetBitLockerDisableWarningForOtherDiskEncryption(value *bool)()
    SetBitLockerEnableStorageCardEncryptionOnMobile(value *bool)()
    SetBitLockerEncryptDevice(value *bool)()
    SetBitLockerFixedDrivePolicy(value BitLockerFixedDrivePolicyable)()
    SetBitLockerRecoveryPasswordRotation(value *BitLockerRecoveryPasswordRotationType)()
    SetBitLockerRemovableDrivePolicy(value BitLockerRemovableDrivePolicyable)()
    SetBitLockerSystemDrivePolicy(value BitLockerSystemDrivePolicyable)()
    SetDefenderAdditionalGuardedFolders(value []string)()
    SetDefenderAdobeReaderLaunchChildProcess(value *DefenderProtectionType)()
    SetDefenderAdvancedRansomewareProtectionType(value *DefenderProtectionType)()
    SetDefenderAllowBehaviorMonitoring(value *bool)()
    SetDefenderAllowCloudProtection(value *bool)()
    SetDefenderAllowEndUserAccess(value *bool)()
    SetDefenderAllowIntrusionPreventionSystem(value *bool)()
    SetDefenderAllowOnAccessProtection(value *bool)()
    SetDefenderAllowRealTimeMonitoring(value *bool)()
    SetDefenderAllowScanArchiveFiles(value *bool)()
    SetDefenderAllowScanDownloads(value *bool)()
    SetDefenderAllowScanNetworkFiles(value *bool)()
    SetDefenderAllowScanRemovableDrivesDuringFullScan(value *bool)()
    SetDefenderAllowScanScriptsLoadedInInternetExplorer(value *bool)()
    SetDefenderAttackSurfaceReductionExcludedPaths(value []string)()
    SetDefenderBlockEndUserAccess(value *bool)()
    SetDefenderBlockPersistenceThroughWmiType(value *DefenderAttackSurfaceType)()
    SetDefenderCheckForSignaturesBeforeRunningScan(value *bool)()
    SetDefenderCloudBlockLevel(value *DefenderCloudBlockLevelType)()
    SetDefenderCloudExtendedTimeoutInSeconds(value *int32)()
    SetDefenderDaysBeforeDeletingQuarantinedMalware(value *int32)()
    SetDefenderDetectedMalwareActions(value DefenderDetectedMalwareActionsable)()
    SetDefenderDisableBehaviorMonitoring(value *bool)()
    SetDefenderDisableCatchupFullScan(value *bool)()
    SetDefenderDisableCatchupQuickScan(value *bool)()
    SetDefenderDisableCloudProtection(value *bool)()
    SetDefenderDisableIntrusionPreventionSystem(value *bool)()
    SetDefenderDisableOnAccessProtection(value *bool)()
    SetDefenderDisableRealTimeMonitoring(value *bool)()
    SetDefenderDisableScanArchiveFiles(value *bool)()
    SetDefenderDisableScanDownloads(value *bool)()
    SetDefenderDisableScanNetworkFiles(value *bool)()
    SetDefenderDisableScanRemovableDrivesDuringFullScan(value *bool)()
    SetDefenderDisableScanScriptsLoadedInInternetExplorer(value *bool)()
    SetDefenderEmailContentExecution(value *DefenderProtectionType)()
    SetDefenderEmailContentExecutionType(value *DefenderAttackSurfaceType)()
    SetDefenderEnableLowCpuPriority(value *bool)()
    SetDefenderEnableScanIncomingMail(value *bool)()
    SetDefenderEnableScanMappedNetworkDrivesDuringFullScan(value *bool)()
    SetDefenderExploitProtectionXml(value []byte)()
    SetDefenderExploitProtectionXmlFileName(value *string)()
    SetDefenderFileExtensionsToExclude(value []string)()
    SetDefenderFilesAndFoldersToExclude(value []string)()
    SetDefenderGuardedFoldersAllowedAppPaths(value []string)()
    SetDefenderGuardMyFoldersType(value *FolderProtectionType)()
    SetDefenderNetworkProtectionType(value *DefenderProtectionType)()
    SetDefenderOfficeAppsExecutableContentCreationOrLaunch(value *DefenderProtectionType)()
    SetDefenderOfficeAppsExecutableContentCreationOrLaunchType(value *DefenderAttackSurfaceType)()
    SetDefenderOfficeAppsLaunchChildProcess(value *DefenderProtectionType)()
    SetDefenderOfficeAppsLaunchChildProcessType(value *DefenderAttackSurfaceType)()
    SetDefenderOfficeAppsOtherProcessInjection(value *DefenderProtectionType)()
    SetDefenderOfficeAppsOtherProcessInjectionType(value *DefenderAttackSurfaceType)()
    SetDefenderOfficeCommunicationAppsLaunchChildProcess(value *DefenderProtectionType)()
    SetDefenderOfficeMacroCodeAllowWin32Imports(value *DefenderProtectionType)()
    SetDefenderOfficeMacroCodeAllowWin32ImportsType(value *DefenderAttackSurfaceType)()
    SetDefenderPotentiallyUnwantedAppAction(value *DefenderProtectionType)()
    SetDefenderPreventCredentialStealingType(value *DefenderProtectionType)()
    SetDefenderProcessCreation(value *DefenderProtectionType)()
    SetDefenderProcessCreationType(value *DefenderAttackSurfaceType)()
    SetDefenderProcessesToExclude(value []string)()
    SetDefenderScanDirection(value *DefenderRealtimeScanDirection)()
    SetDefenderScanMaxCpuPercentage(value *int32)()
    SetDefenderScanType(value *DefenderScanType)()
    SetDefenderScheduledQuickScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetDefenderScheduledScanDay(value *WeeklySchedule)()
    SetDefenderScheduledScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetDefenderScriptDownloadedPayloadExecution(value *DefenderProtectionType)()
    SetDefenderScriptDownloadedPayloadExecutionType(value *DefenderAttackSurfaceType)()
    SetDefenderScriptObfuscatedMacroCode(value *DefenderProtectionType)()
    SetDefenderScriptObfuscatedMacroCodeType(value *DefenderAttackSurfaceType)()
    SetDefenderSecurityCenterBlockExploitProtectionOverride(value *bool)()
    SetDefenderSecurityCenterDisableAccountUI(value *bool)()
    SetDefenderSecurityCenterDisableAppBrowserUI(value *bool)()
    SetDefenderSecurityCenterDisableClearTpmUI(value *bool)()
    SetDefenderSecurityCenterDisableFamilyUI(value *bool)()
    SetDefenderSecurityCenterDisableHardwareUI(value *bool)()
    SetDefenderSecurityCenterDisableHealthUI(value *bool)()
    SetDefenderSecurityCenterDisableNetworkUI(value *bool)()
    SetDefenderSecurityCenterDisableNotificationAreaUI(value *bool)()
    SetDefenderSecurityCenterDisableRansomwareUI(value *bool)()
    SetDefenderSecurityCenterDisableSecureBootUI(value *bool)()
    SetDefenderSecurityCenterDisableTroubleshootingUI(value *bool)()
    SetDefenderSecurityCenterDisableVirusUI(value *bool)()
    SetDefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI(value *bool)()
    SetDefenderSecurityCenterHelpEmail(value *string)()
    SetDefenderSecurityCenterHelpPhone(value *string)()
    SetDefenderSecurityCenterHelpURL(value *string)()
    SetDefenderSecurityCenterITContactDisplay(value *DefenderSecurityCenterITContactDisplayType)()
    SetDefenderSecurityCenterNotificationsFromApp(value *DefenderSecurityCenterNotificationsFromAppType)()
    SetDefenderSecurityCenterOrganizationDisplayName(value *string)()
    SetDefenderSignatureUpdateIntervalInHours(value *int32)()
    SetDefenderSubmitSamplesConsentType(value *DefenderSubmitSamplesConsentType)()
    SetDefenderUntrustedExecutable(value *DefenderProtectionType)()
    SetDefenderUntrustedExecutableType(value *DefenderAttackSurfaceType)()
    SetDefenderUntrustedUSBProcess(value *DefenderProtectionType)()
    SetDefenderUntrustedUSBProcessType(value *DefenderAttackSurfaceType)()
    SetDeviceGuardEnableSecureBootWithDMA(value *bool)()
    SetDeviceGuardEnableVirtualizationBasedSecurity(value *bool)()
    SetDeviceGuardLaunchSystemGuard(value *Enablement)()
    SetDeviceGuardLocalSystemAuthorityCredentialGuardSettings(value *DeviceGuardLocalSystemAuthorityCredentialGuardType)()
    SetDeviceGuardSecureBootWithDMA(value *SecureBootWithDMAType)()
    SetDmaGuardDeviceEnumerationPolicy(value *DmaGuardDeviceEnumerationPolicyType)()
    SetFirewallBlockStatefulFTP(value *bool)()
    SetFirewallCertificateRevocationListCheckMethod(value *FirewallCertificateRevocationListCheckMethodType)()
    SetFirewallIdleTimeoutForSecurityAssociationInSeconds(value *int32)()
    SetFirewallIPSecExemptionsAllowDHCP(value *bool)()
    SetFirewallIPSecExemptionsAllowICMP(value *bool)()
    SetFirewallIPSecExemptionsAllowNeighborDiscovery(value *bool)()
    SetFirewallIPSecExemptionsAllowRouterDiscovery(value *bool)()
    SetFirewallIPSecExemptionsNone(value *bool)()
    SetFirewallMergeKeyingModuleSettings(value *bool)()
    SetFirewallPacketQueueingMethod(value *FirewallPacketQueueingMethodType)()
    SetFirewallPreSharedKeyEncodingMethod(value *FirewallPreSharedKeyEncodingMethodType)()
    SetFirewallProfileDomain(value WindowsFirewallNetworkProfileable)()
    SetFirewallProfilePrivate(value WindowsFirewallNetworkProfileable)()
    SetFirewallProfilePublic(value WindowsFirewallNetworkProfileable)()
    SetFirewallRules(value []WindowsFirewallRuleable)()
    SetLanManagerAuthenticationLevel(value *LanManagerAuthenticationLevel)()
    SetLanManagerWorkstationDisableInsecureGuestLogons(value *bool)()
    SetLocalSecurityOptionsAdministratorAccountName(value *string)()
    SetLocalSecurityOptionsAdministratorElevationPromptBehavior(value *LocalSecurityOptionsAdministratorElevationPromptBehaviorType)()
    SetLocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares(value *bool)()
    SetLocalSecurityOptionsAllowPKU2UAuthenticationRequests(value *bool)()
    SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager(value *string)()
    SetLocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool(value *bool)()
    SetLocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn(value *bool)()
    SetLocalSecurityOptionsAllowUIAccessApplicationElevation(value *bool)()
    SetLocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations(value *bool)()
    SetLocalSecurityOptionsAllowUndockWithoutHavingToLogon(value *bool)()
    SetLocalSecurityOptionsBlockMicrosoftAccounts(value *bool)()
    SetLocalSecurityOptionsBlockRemoteLogonWithBlankPassword(value *bool)()
    SetLocalSecurityOptionsBlockRemoteOpticalDriveAccess(value *bool)()
    SetLocalSecurityOptionsBlockUsersInstallingPrinterDrivers(value *bool)()
    SetLocalSecurityOptionsClearVirtualMemoryPageFile(value *bool)()
    SetLocalSecurityOptionsClientDigitallySignCommunicationsAlways(value *bool)()
    SetLocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers(value *bool)()
    SetLocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation(value *bool)()
    SetLocalSecurityOptionsDisableAdministratorAccount(value *bool)()
    SetLocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees(value *bool)()
    SetLocalSecurityOptionsDisableGuestAccount(value *bool)()
    SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways(value *bool)()
    SetLocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees(value *bool)()
    SetLocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts(value *bool)()
    SetLocalSecurityOptionsDoNotRequireCtrlAltDel(value *bool)()
    SetLocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange(value *bool)()
    SetLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(value *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType)()
    SetLocalSecurityOptionsGuestAccountName(value *string)()
    SetLocalSecurityOptionsHideLastSignedInUser(value *bool)()
    SetLocalSecurityOptionsHideUsernameAtSignIn(value *bool)()
    SetLocalSecurityOptionsInformationDisplayedOnLockScreen(value *LocalSecurityOptionsInformationDisplayedOnLockScreenType)()
    SetLocalSecurityOptionsInformationShownOnLockScreen(value *LocalSecurityOptionsInformationShownOnLockScreenType)()
    SetLocalSecurityOptionsLogOnMessageText(value *string)()
    SetLocalSecurityOptionsLogOnMessageTitle(value *string)()
    SetLocalSecurityOptionsMachineInactivityLimit(value *int32)()
    SetLocalSecurityOptionsMachineInactivityLimitInMinutes(value *int32)()
    SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients(value *LocalSecurityOptionsMinimumSessionSecurity)()
    SetLocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers(value *LocalSecurityOptionsMinimumSessionSecurity)()
    SetLocalSecurityOptionsOnlyElevateSignedExecutables(value *bool)()
    SetLocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares(value *bool)()
    SetLocalSecurityOptionsSmartCardRemovalBehavior(value *LocalSecurityOptionsSmartCardRemovalBehaviorType)()
    SetLocalSecurityOptionsStandardUserElevationPromptBehavior(value *LocalSecurityOptionsStandardUserElevationPromptBehaviorType)()
    SetLocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation(value *bool)()
    SetLocalSecurityOptionsUseAdminApprovalMode(value *bool)()
    SetLocalSecurityOptionsUseAdminApprovalModeForAdministrators(value *bool)()
    SetLocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations(value *bool)()
    SetSmartScreenBlockOverrideForFiles(value *bool)()
    SetSmartScreenEnableInShell(value *bool)()
    SetUserRightsAccessCredentialManagerAsTrustedCaller(value DeviceManagementUserRightsSettingable)()
    SetUserRightsActAsPartOfTheOperatingSystem(value DeviceManagementUserRightsSettingable)()
    SetUserRightsAllowAccessFromNetwork(value DeviceManagementUserRightsSettingable)()
    SetUserRightsBackupData(value DeviceManagementUserRightsSettingable)()
    SetUserRightsBlockAccessFromNetwork(value DeviceManagementUserRightsSettingable)()
    SetUserRightsChangeSystemTime(value DeviceManagementUserRightsSettingable)()
    SetUserRightsCreateGlobalObjects(value DeviceManagementUserRightsSettingable)()
    SetUserRightsCreatePageFile(value DeviceManagementUserRightsSettingable)()
    SetUserRightsCreatePermanentSharedObjects(value DeviceManagementUserRightsSettingable)()
    SetUserRightsCreateSymbolicLinks(value DeviceManagementUserRightsSettingable)()
    SetUserRightsCreateToken(value DeviceManagementUserRightsSettingable)()
    SetUserRightsDebugPrograms(value DeviceManagementUserRightsSettingable)()
    SetUserRightsDelegation(value DeviceManagementUserRightsSettingable)()
    SetUserRightsDenyLocalLogOn(value DeviceManagementUserRightsSettingable)()
    SetUserRightsGenerateSecurityAudits(value DeviceManagementUserRightsSettingable)()
    SetUserRightsImpersonateClient(value DeviceManagementUserRightsSettingable)()
    SetUserRightsIncreaseSchedulingPriority(value DeviceManagementUserRightsSettingable)()
    SetUserRightsLoadUnloadDrivers(value DeviceManagementUserRightsSettingable)()
    SetUserRightsLocalLogOn(value DeviceManagementUserRightsSettingable)()
    SetUserRightsLockMemory(value DeviceManagementUserRightsSettingable)()
    SetUserRightsManageAuditingAndSecurityLogs(value DeviceManagementUserRightsSettingable)()
    SetUserRightsManageVolumes(value DeviceManagementUserRightsSettingable)()
    SetUserRightsModifyFirmwareEnvironment(value DeviceManagementUserRightsSettingable)()
    SetUserRightsModifyObjectLabels(value DeviceManagementUserRightsSettingable)()
    SetUserRightsProfileSingleProcess(value DeviceManagementUserRightsSettingable)()
    SetUserRightsRemoteDesktopServicesLogOn(value DeviceManagementUserRightsSettingable)()
    SetUserRightsRemoteShutdown(value DeviceManagementUserRightsSettingable)()
    SetUserRightsRestoreData(value DeviceManagementUserRightsSettingable)()
    SetUserRightsTakeOwnership(value DeviceManagementUserRightsSettingable)()
    SetWindowsDefenderTamperProtection(value *WindowsDefenderTamperProtectionOptions)()
    SetXboxServicesAccessoryManagementServiceStartupMode(value *ServiceStartType)()
    SetXboxServicesEnableXboxGameSaveTask(value *bool)()
    SetXboxServicesLiveAuthManagerServiceStartupMode(value *ServiceStartType)()
    SetXboxServicesLiveGameSaveServiceStartupMode(value *ServiceStartType)()
    SetXboxServicesLiveNetworkingServiceStartupMode(value *ServiceStartType)()
}
