package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type RemoteAction int

const (
    // User initiates an unknown action.
    UNKNOWN_REMOTEACTION RemoteAction = iota
    // User initiates an action to factory reset a device. 
    FACTORYRESET_REMOTEACTION
    // User initiates an action to remove company data from a device. 
    REMOVECOMPANYDATA_REMOTEACTION
    // User initiates an action to remove the passcode of an iOS device, or reset the passcode of Android / Windows device. 
    RESETPASSCODE_REMOTEACTION
    // User initiates an action to remote lock a device.
    REMOTELOCK_REMOTEACTION
    // User initiates an action to enable lost mode on a supervised iOS device.
    ENABLELOSTMODE_REMOTEACTION
    // User initiates an action to disable lost mode on a supervised iOS device.
    DISABLELOSTMODE_REMOTEACTION
    // User initiates an action to locate a supervised iOS device.
    LOCATEDEVICE_REMOTEACTION
    // User initiates an action to reboot a Windows device.
    REBOOTNOW_REMOTEACTION
    // User initiates an action to reset the pin for passport for work on windows phone device.
    RECOVERPASSCODE_REMOTEACTION
    // User initiates an action to clean up windows device.
    CLEANWINDOWSDEVICE_REMOTEACTION
    // User initiates an action to log out current user on shared apple device.
    LOGOUTSHAREDAPPLEDEVICEACTIVEUSER_REMOTEACTION
    // User initiates an action to run quick scan on device.
    QUICKSCAN_REMOTEACTION
    // User initiates an action to run full scan on device.
    FULLSCAN_REMOTEACTION
    // User initiates an action to update malware signatures on device.
    WINDOWSDEFENDERUPDATESIGNATURES_REMOTEACTION
    // User initiates an action remote wipe device with keeping enrollment data.
    FACTORYRESETKEEPENROLLMENTDATA_REMOTEACTION
    // User initiates an action to update account on device.
    UPDATEDEVICEACCOUNT_REMOTEACTION
    // User initiates an action to automatice redeploy the device
    AUTOMATICREDEPLOYMENT_REMOTEACTION
    // User initiates an action to shut down the device.
    SHUTDOWN_REMOTEACTION
    // User initiates an action to Rotate BitLockerKeys on the device.
    ROTATEBITLOCKERKEYS_REMOTEACTION
    // User initiates an action to Rotate FileVaultKey on mac.
    ROTATEFILEVAULTKEY_REMOTEACTION
    // User initiates an action to Get FileVaultKey on mac.
    GETFILEVAULTKEY_REMOTEACTION
    // User initiates an action to Set Device Name on the device.
    SETDEVICENAME_REMOTEACTION
    // User initiates an action to Activate eSIM on the device.
    ACTIVATEDEVICEESIM_REMOTEACTION
)

func (i RemoteAction) String() string {
    return []string{"unknown", "factoryReset", "removeCompanyData", "resetPasscode", "remoteLock", "enableLostMode", "disableLostMode", "locateDevice", "rebootNow", "recoverPasscode", "cleanWindowsDevice", "logoutSharedAppleDeviceActiveUser", "quickScan", "fullScan", "windowsDefenderUpdateSignatures", "factoryResetKeepEnrollmentData", "updateDeviceAccount", "automaticRedeployment", "shutDown", "rotateBitLockerKeys", "rotateFileVaultKey", "getFileVaultKey", "setDeviceName", "activateDeviceEsim"}[i]
}
func ParseRemoteAction(v string) (interface{}, error) {
    result := UNKNOWN_REMOTEACTION
    switch v {
        case "unknown":
            result = UNKNOWN_REMOTEACTION
        case "factoryReset":
            result = FACTORYRESET_REMOTEACTION
        case "removeCompanyData":
            result = REMOVECOMPANYDATA_REMOTEACTION
        case "resetPasscode":
            result = RESETPASSCODE_REMOTEACTION
        case "remoteLock":
            result = REMOTELOCK_REMOTEACTION
        case "enableLostMode":
            result = ENABLELOSTMODE_REMOTEACTION
        case "disableLostMode":
            result = DISABLELOSTMODE_REMOTEACTION
        case "locateDevice":
            result = LOCATEDEVICE_REMOTEACTION
        case "rebootNow":
            result = REBOOTNOW_REMOTEACTION
        case "recoverPasscode":
            result = RECOVERPASSCODE_REMOTEACTION
        case "cleanWindowsDevice":
            result = CLEANWINDOWSDEVICE_REMOTEACTION
        case "logoutSharedAppleDeviceActiveUser":
            result = LOGOUTSHAREDAPPLEDEVICEACTIVEUSER_REMOTEACTION
        case "quickScan":
            result = QUICKSCAN_REMOTEACTION
        case "fullScan":
            result = FULLSCAN_REMOTEACTION
        case "windowsDefenderUpdateSignatures":
            result = WINDOWSDEFENDERUPDATESIGNATURES_REMOTEACTION
        case "factoryResetKeepEnrollmentData":
            result = FACTORYRESETKEEPENROLLMENTDATA_REMOTEACTION
        case "updateDeviceAccount":
            result = UPDATEDEVICEACCOUNT_REMOTEACTION
        case "automaticRedeployment":
            result = AUTOMATICREDEPLOYMENT_REMOTEACTION
        case "shutDown":
            result = SHUTDOWN_REMOTEACTION
        case "rotateBitLockerKeys":
            result = ROTATEBITLOCKERKEYS_REMOTEACTION
        case "rotateFileVaultKey":
            result = ROTATEFILEVAULTKEY_REMOTEACTION
        case "getFileVaultKey":
            result = GETFILEVAULTKEY_REMOTEACTION
        case "setDeviceName":
            result = SETDEVICENAME_REMOTEACTION
        case "activateDeviceEsim":
            result = ACTIVATEDEVICEESIM_REMOTEACTION
        default:
            return 0, errors.New("Unknown RemoteAction value: " + v)
    }
    return &result, nil
}
func SerializeRemoteAction(values []RemoteAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
