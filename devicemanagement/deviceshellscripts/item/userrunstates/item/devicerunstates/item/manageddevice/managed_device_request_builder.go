package manageddevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i04c393d23cba562fcc9574307fc38997f7db0d9898a18e0a1c1e2316041162a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/requestremoteassistance"
    i0983e20d99eb6f73ea094d9dbf0774346f9241737433b97b856bb4e16f807c26 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/playlostmodesound"
    i09f836965832899fa6311020777c4b6298db48ed92ac54590457f23f2ebf6aec "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/retire"
    i1a8779d0a0c7c6ca86c0a6497210c7bf0d739251cfb2ed6e4e0bd792ef6068a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/syncdevice"
    i1f36242d462111330d1f4560717fee153e6ecea56e6ea19b927f143001f051ed "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/updatewindowsdeviceaccount"
    i2370de97e55ebaaa039774252c01a834b7cbf622485b557a93f8858b88f78987 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/windowsdefenderscan"
    i2d814547aec73ec57e913a19addc51fd67f0e07c1ade2836db4224920d2ba7b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/activatedeviceesim"
    i336ad37a7e1d698ad2d7aa608ffd4f17076747c0a75509bfbe30066412ff10b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/remotelock"
    i384b04dbeaf121b322b59ee189ca311192720b30722a0142d559d59b12e4f214 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/logoutsharedappledeviceactiveuser"
    i3f2d3a51dc171e8ba015d086d2f69815144d23159c7c98de3e3424d25368f623 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/deleteuserfromsharedappledevice"
    i452ab6c1f7a7459c799f22086ba2572ff9eb1f4607a9cbb5c74da7bd6acf9dc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/ref"
    i45e0a5b55ae929bbd1cb6135c7d1b84507888096351778cf0d098d8319063249 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/shutdown"
    i5c246acefab1f224dda6ade5d86086c470d82aa25f432ab70ef41f57278cea9e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rotatefilevaultkey"
    i6073ba1153c0afebdcd4608ccea5e7e4d1ff3213972361f1a5fe33603f180b7c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getcloudpcremoteactionresults"
    i6664e22465e1a7f842f7dbd337c85b3ec72d349a0d265d2ec04d4cd9b5eedc89 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getfilevaultkey"
    i67729b0402c3b9a7b854cfe37980a5f4ec64510fc241864ce0fb67d0e793ab8f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/disablelostmode"
    i6bf5bed63eb697a2c063bb2f62300161c7f9d67c5ec0ae7c385edd58eb38b939 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/setdevicename"
    i6cac51c6a5d88fc80593b07dfbe51f119392a6cb61bae1b733a9ff65433d409e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/createdevicelogcollectionrequest"
    i6dd3182fce92abf4b80d02a5058e8787af08fc2ec085394edbec7fe59ed7aab7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/overridecompliancestate"
    i6fef8231babfcbef12ceac4670e4640b1f9fa5908c3ac23cfc5b59942f395990 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/bypassactivationlock"
    i79a6973f4d40aa7a8a7dcbc19ded45b92b4972f9a2a79acd08913532444ec442 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/revokeapplevpplicenses"
    i7b158fccf72d7011089fe0ebdde1bba6013d6389629feb801a04a2289d8311f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/sendcustomnotificationtocompanyportal"
    i7fcf2a2dc546dd9324d7b6b2b82d8c933dbbd65c24f896f1602dcb9594f15d7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/windowsdefenderupdatesignatures"
    ia85e2afd6a2f8100fa7ba4a967610063545b9a81015f9ac320a2db7a8cfd4ca5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/locatedevice"
    iab8439718df12560a25e2edcf55a786fead5110795dc7808a30663503f11afe2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/triggerconfigurationmanageraction"
    iae5f043d307dcb4475baf1d4aebe21e39c5a372abfb29c50b4697c9d9848641e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/recoverpasscode"
    iaf4a01e505fa240a196e8bfc5e8c60c8afad974a2592239f7c74d2f3da69a287 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/reprovisioncloudpc"
    iba49f8e40bde0802817c9f12f495f49cc5d6cdf1846369d99009c4cd02b0335d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/resetpasscode"
    ibfe9201bb2f1ba2cbc9ea02bc0e5d1eabda3657ebb00981cb1c7c0505405045c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rebootnow"
    ic522015ed6052a1a116e1f940fd65c8b287aeb78ea6ded4da47f66542394fa45 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/cleanwindowsdevice"
    ica2174b36d03086aba77732cdf90ecc3468f45216ebd7bb3e226d50d5cd5c599 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/enablelostmode"
    ica29b53b04a9291bf2d8ebf510dd31b2b001bd11d4c0e3d753b6fdcbd3ec4575 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rotatebitlockerkeys"
    idcce05877586645ad78653ca5f3c880e093f991fda5538b2acddacadef588f69 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/wipe"
    iea13686895285205bc0a71513b3d539595cd074a638334066aac1f89b78dac1f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getnoncompliantsettings"
    if153307cb65252c61c9127e2418f320bebced5b858e7c9067ba06930afbc92c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/resizecloudpc"
)

type ManagedDeviceRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ManagedDeviceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*i2d814547aec73ec57e913a19addc51fd67f0e07c1ade2836db4224920d2ba7b9.ActivateDeviceEsimRequestBuilder) {
    return i2d814547aec73ec57e913a19addc51fd67f0e07c1ade2836db4224920d2ba7b9.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*i6fef8231babfcbef12ceac4670e4640b1f9fa5908c3ac23cfc5b59942f395990.BypassActivationLockRequestBuilder) {
    return i6fef8231babfcbef12ceac4670e4640b1f9fa5908c3ac23cfc5b59942f395990.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*ic522015ed6052a1a116e1f940fd65c8b287aeb78ea6ded4da47f66542394fa45.CleanWindowsDeviceRequestBuilder) {
    return ic522015ed6052a1a116e1f940fd65c8b287aeb78ea6ded4da47f66542394fa45.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/deviceShellScripts/{deviceShellScript_id}/userRunStates/{deviceManagementScriptUserState_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewManagedDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*i6cac51c6a5d88fc80593b07dfbe51f119392a6cb61bae1b733a9ff65433d409e.CreateDeviceLogCollectionRequestRequestBuilder) {
    return i6cac51c6a5d88fc80593b07dfbe51f119392a6cb61bae1b733a9ff65433d409e.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CreateGetRequestInformation(q func (value *ManagedDeviceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedDeviceRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*i3f2d3a51dc171e8ba015d086d2f69815144d23159c7c98de3e3424d25368f623.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i3f2d3a51dc171e8ba015d086d2f69815144d23159c7c98de3e3424d25368f623.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i67729b0402c3b9a7b854cfe37980a5f4ec64510fc241864ce0fb67d0e793ab8f.DisableLostModeRequestBuilder) {
    return i67729b0402c3b9a7b854cfe37980a5f4ec64510fc241864ce0fb67d0e793ab8f.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*ica2174b36d03086aba77732cdf90ecc3468f45216ebd7bb3e226d50d5cd5c599.EnableLostModeRequestBuilder) {
    return ica2174b36d03086aba77732cdf90ecc3468f45216ebd7bb3e226d50d5cd5c599.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Get(q func (value *ManagedDeviceRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedDevice() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice), nil
}
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*i6073ba1153c0afebdcd4608ccea5e7e4d1ff3213972361f1a5fe33603f180b7c.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i6073ba1153c0afebdcd4608ccea5e7e4d1ff3213972361f1a5fe33603f180b7c.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*i6664e22465e1a7f842f7dbd337c85b3ec72d349a0d265d2ec04d4cd9b5eedc89.GetFileVaultKeyRequestBuilder) {
    return i6664e22465e1a7f842f7dbd337c85b3ec72d349a0d265d2ec04d4cd9b5eedc89.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*iea13686895285205bc0a71513b3d539595cd074a638334066aac1f89b78dac1f.GetNonCompliantSettingsRequestBuilder) {
    return iea13686895285205bc0a71513b3d539595cd074a638334066aac1f89b78dac1f.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*ia85e2afd6a2f8100fa7ba4a967610063545b9a81015f9ac320a2db7a8cfd4ca5.LocateDeviceRequestBuilder) {
    return ia85e2afd6a2f8100fa7ba4a967610063545b9a81015f9ac320a2db7a8cfd4ca5.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i384b04dbeaf121b322b59ee189ca311192720b30722a0142d559d59b12e4f214.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i384b04dbeaf121b322b59ee189ca311192720b30722a0142d559d59b12e4f214.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*i6dd3182fce92abf4b80d02a5058e8787af08fc2ec085394edbec7fe59ed7aab7.OverrideComplianceStateRequestBuilder) {
    return i6dd3182fce92abf4b80d02a5058e8787af08fc2ec085394edbec7fe59ed7aab7.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*i0983e20d99eb6f73ea094d9dbf0774346f9241737433b97b856bb4e16f807c26.PlayLostModeSoundRequestBuilder) {
    return i0983e20d99eb6f73ea094d9dbf0774346f9241737433b97b856bb4e16f807c26.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*ibfe9201bb2f1ba2cbc9ea02bc0e5d1eabda3657ebb00981cb1c7c0505405045c.RebootNowRequestBuilder) {
    return ibfe9201bb2f1ba2cbc9ea02bc0e5d1eabda3657ebb00981cb1c7c0505405045c.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*iae5f043d307dcb4475baf1d4aebe21e39c5a372abfb29c50b4697c9d9848641e.RecoverPasscodeRequestBuilder) {
    return iae5f043d307dcb4475baf1d4aebe21e39c5a372abfb29c50b4697c9d9848641e.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Ref()(*i452ab6c1f7a7459c799f22086ba2572ff9eb1f4607a9cbb5c74da7bd6acf9dc0.RefRequestBuilder) {
    return i452ab6c1f7a7459c799f22086ba2572ff9eb1f4607a9cbb5c74da7bd6acf9dc0.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*i336ad37a7e1d698ad2d7aa608ffd4f17076747c0a75509bfbe30066412ff10b9.RemoteLockRequestBuilder) {
    return i336ad37a7e1d698ad2d7aa608ffd4f17076747c0a75509bfbe30066412ff10b9.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*iaf4a01e505fa240a196e8bfc5e8c60c8afad974a2592239f7c74d2f3da69a287.ReprovisionCloudPcRequestBuilder) {
    return iaf4a01e505fa240a196e8bfc5e8c60c8afad974a2592239f7c74d2f3da69a287.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*i04c393d23cba562fcc9574307fc38997f7db0d9898a18e0a1c1e2316041162a7.RequestRemoteAssistanceRequestBuilder) {
    return i04c393d23cba562fcc9574307fc38997f7db0d9898a18e0a1c1e2316041162a7.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*iba49f8e40bde0802817c9f12f495f49cc5d6cdf1846369d99009c4cd02b0335d.ResetPasscodeRequestBuilder) {
    return iba49f8e40bde0802817c9f12f495f49cc5d6cdf1846369d99009c4cd02b0335d.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*if153307cb65252c61c9127e2418f320bebced5b858e7c9067ba06930afbc92c8.ResizeCloudPcRequestBuilder) {
    return if153307cb65252c61c9127e2418f320bebced5b858e7c9067ba06930afbc92c8.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*i09f836965832899fa6311020777c4b6298db48ed92ac54590457f23f2ebf6aec.RetireRequestBuilder) {
    return i09f836965832899fa6311020777c4b6298db48ed92ac54590457f23f2ebf6aec.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*i79a6973f4d40aa7a8a7dcbc19ded45b92b4972f9a2a79acd08913532444ec442.RevokeAppleVppLicensesRequestBuilder) {
    return i79a6973f4d40aa7a8a7dcbc19ded45b92b4972f9a2a79acd08913532444ec442.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*ica29b53b04a9291bf2d8ebf510dd31b2b001bd11d4c0e3d753b6fdcbd3ec4575.RotateBitLockerKeysRequestBuilder) {
    return ica29b53b04a9291bf2d8ebf510dd31b2b001bd11d4c0e3d753b6fdcbd3ec4575.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*i5c246acefab1f224dda6ade5d86086c470d82aa25f432ab70ef41f57278cea9e.RotateFileVaultKeyRequestBuilder) {
    return i5c246acefab1f224dda6ade5d86086c470d82aa25f432ab70ef41f57278cea9e.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*i7b158fccf72d7011089fe0ebdde1bba6013d6389629feb801a04a2289d8311f7.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return i7b158fccf72d7011089fe0ebdde1bba6013d6389629feb801a04a2289d8311f7.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*i6bf5bed63eb697a2c063bb2f62300161c7f9d67c5ec0ae7c385edd58eb38b939.SetDeviceNameRequestBuilder) {
    return i6bf5bed63eb697a2c063bb2f62300161c7f9d67c5ec0ae7c385edd58eb38b939.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*i45e0a5b55ae929bbd1cb6135c7d1b84507888096351778cf0d098d8319063249.ShutDownRequestBuilder) {
    return i45e0a5b55ae929bbd1cb6135c7d1b84507888096351778cf0d098d8319063249.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*i1a8779d0a0c7c6ca86c0a6497210c7bf0d739251cfb2ed6e4e0bd792ef6068a0.SyncDeviceRequestBuilder) {
    return i1a8779d0a0c7c6ca86c0a6497210c7bf0d739251cfb2ed6e4e0bd792ef6068a0.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*iab8439718df12560a25e2edcf55a786fead5110795dc7808a30663503f11afe2.TriggerConfigurationManagerActionRequestBuilder) {
    return iab8439718df12560a25e2edcf55a786fead5110795dc7808a30663503f11afe2.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*i1f36242d462111330d1f4560717fee153e6ecea56e6ea19b927f143001f051ed.UpdateWindowsDeviceAccountRequestBuilder) {
    return i1f36242d462111330d1f4560717fee153e6ecea56e6ea19b927f143001f051ed.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*i2370de97e55ebaaa039774252c01a834b7cbf622485b557a93f8858b88f78987.WindowsDefenderScanRequestBuilder) {
    return i2370de97e55ebaaa039774252c01a834b7cbf622485b557a93f8858b88f78987.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i7fcf2a2dc546dd9324d7b6b2b82d8c933dbbd65c24f896f1602dcb9594f15d7d.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i7fcf2a2dc546dd9324d7b6b2b82d8c933dbbd65c24f896f1602dcb9594f15d7d.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*idcce05877586645ad78653ca5f3c880e093f991fda5538b2acddacadef588f69.WipeRequestBuilder) {
    return idcce05877586645ad78653ca5f3c880e093f991fda5538b2acddacadef588f69.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
