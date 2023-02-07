package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder provides operations to call the getExpiringVppTokenCount method.
type DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal instantiates a new MicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewDepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, expiringBeforeDateTime *string)(*DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    m := &DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/microsoft.graph.getExpiringVppTokenCount(expiringBeforeDateTime='{expiringBeforeDateTime}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if expiringBeforeDateTime != nil {
        urlTplParams["expiringBeforeDateTime"] = *expiringBeforeDateTime
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder instantiates a new MicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewDepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getExpiringVppTokenCount
func (m *DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable), nil
}
// ToGetRequestInformation invoke function getExpiringVppTokenCount
func (m *DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsMicrosoftGraphGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
