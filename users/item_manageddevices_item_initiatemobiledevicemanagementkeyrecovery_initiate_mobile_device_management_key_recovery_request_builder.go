package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
type ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal instantiates a new ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder and sets the default values.
func NewItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    m := &ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/initiateMobileDeviceManagementKeyRecovery", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder instantiates a new ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder and sets the default values.
func NewItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform MDM key recovery and TPM attestation
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder when successful
func (m *ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    return NewItemManageddevicesItemInitiatemobiledevicemanagementkeyrecoveryInitiateMobileDeviceManagementKeyRecoveryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
