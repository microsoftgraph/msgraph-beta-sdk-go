package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder provides operations to call the initiateMobileDeviceManagementKeyRecovery method.
type ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal instantiates a new MicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    m := &ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/microsoft.graph.initiateMobileDeviceManagementKeyRecovery";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder instantiates a new MicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform MDM key recovery and TPM attestation
func (m *ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) Post(ctx context.Context, requestConfiguration *ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation perform MDM key recovery and TPM attestation
func (m *ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesItemMicrosoftGraphInitiateMobileDeviceManagementKeyRecoveryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
