package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder provides operations to call the overrideComplianceState method.
type ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderInternal instantiates a new MicrosoftGraphOverrideComplianceStateRequestBuilder and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder) {
    m := &ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/microsoft.graph.overrideComplianceState";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder instantiates a new MicrosoftGraphOverrideComplianceStateRequestBuilder and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action overrideComplianceState
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder) Post(ctx context.Context, body ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBodyable, requestConfiguration *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action overrideComplianceState
func (m *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateOverrideComplianceStatePostRequestBodyable, requestConfiguration *ComanagedDevicesItemMicrosoftGraphOverrideComplianceStateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
