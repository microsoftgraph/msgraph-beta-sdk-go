package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder provides operations to call the endRemoteHelpSession method.
type ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderInternal instantiates a new MicrosoftGraphEndRemoteHelpSessionRequestBuilder and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder) {
    m := &ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/microsoft.graph.endRemoteHelpSession";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder instantiates a new MicrosoftGraphEndRemoteHelpSessionRequestBuilder and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remote help - End ACS session, Pubsub session and delete Remote help session
func (m *ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder) Post(ctx context.Context, body ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionEndRemoteHelpSessionPostRequestBodyable, requestConfiguration *ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation remote help - End ACS session, Pubsub session and delete Remote help session
func (m *ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionEndRemoteHelpSessionPostRequestBodyable, requestConfiguration *ComanagedDevicesItemMicrosoftGraphEndRemoteHelpSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
