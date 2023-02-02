package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilderInternal instantiates a new HasPayloadLinksRequestBuilder and sets the default values.
func NewMobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder) {
    m := &MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/microsoft.graph.hasPayloadLinks";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder instantiates a new HasPayloadLinksRequestBuilder and sets the default values.
func NewMobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
func (m *MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder) Post(ctx context.Context, body MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBodyable, requestConfiguration *MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
func (m *MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBodyable, requestConfiguration *MobileAppsMicrosoftGraphHasPayloadLinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
