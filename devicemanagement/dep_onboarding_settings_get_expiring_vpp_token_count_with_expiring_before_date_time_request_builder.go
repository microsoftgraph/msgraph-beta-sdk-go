package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder provides operations to call the getExpiringVppTokenCount method.
type DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal instantiates a new GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewDepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, expiringBeforeDateTime *string)(*DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    m := &DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/getExpiringVppTokenCount(expiringBeforeDateTime='{expiringBeforeDateTime}')", pathParameters),
    }
    if expiringBeforeDateTime != nil {
        m.BaseRequestBuilder.PathParameters["expiringBeforeDateTime"] = *expiringBeforeDateTime
    }
    return m
}
// NewDepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder instantiates a new GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewDepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getExpiringVppTokenCount
// Deprecated: This method is obsolete. Use GetAsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponse instead.
func (m *DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable), nil
}
// GetAsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponse invoke function getExpiringVppTokenCount
func (m *DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) GetAsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponse(ctx context.Context, requestConfiguration *DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponseable), nil
}
// ToGetRequestInformation invoke function getExpiringVppTokenCount
func (m *DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) WithUrl(rawUrl string)(*DepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    return NewDepOnboardingSettingsGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
