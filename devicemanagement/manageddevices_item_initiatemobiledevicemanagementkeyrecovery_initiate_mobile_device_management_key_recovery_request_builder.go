package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
type ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal instantiates a new ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder and sets the default values.
func NewManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    m := &ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/initiateMobileDeviceManagementKeyRecovery", pathParameters),
    }
    return m
}
// NewManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder instantiates a new ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder and sets the default values.
func NewManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform MDM key recovery and TPM attestation
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) Post(ctx context.Context, requestConfiguration *ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation perform MDM key recovery and TPM attestation
// returns a *RequestInformation when successful
func (m *ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder when successful
func (m *ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
