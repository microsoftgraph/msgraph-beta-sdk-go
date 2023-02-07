package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder provides operations to call the updateSoftware method.
type DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicesItemMicrosoftGraphUpdateSoftwareRequestBuilderInternal instantiates a new MicrosoftGraphUpdateSoftwareRequestBuilder and sets the default values.
func NewDevicesItemMicrosoftGraphUpdateSoftwareRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder) {
    m := &DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/devices/{teamworkDevice%2Did}/microsoft.graph.updateSoftware";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder instantiates a new MicrosoftGraphUpdateSoftwareRequestBuilder and sets the default values.
func NewDevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicesItemMicrosoftGraphUpdateSoftwareRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the software for a Microsoft Teams-enabled device. This API triggers a long-running operation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/teamworkdevice-updatesoftware?view=graph-rest-1.0
func (m *DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder) Post(ctx context.Context, body DevicesItemMicrosoftGraphUpdateSoftwareUpdateSoftwarePostRequestBodyable, requestConfiguration *DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation update the software for a Microsoft Teams-enabled device. This API triggers a long-running operation.
func (m *DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilder) ToPostRequestInformation(ctx context.Context, body DevicesItemMicrosoftGraphUpdateSoftwareUpdateSoftwarePostRequestBodyable, requestConfiguration *DevicesItemMicrosoftGraphUpdateSoftwareRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
