package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder provides operations to call the compare method.
type TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderGetQueryParameters invoke function compare
type TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderGetQueryParameters
}
// NewTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderInternal instantiates a new TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder and sets the default values.
func NewTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, templateId *string)(*TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) {
    m := &TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/{deviceManagementTemplate%2Did1}/compare(templateId='{templateId}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if templateId != nil {
        m.BaseRequestBuilder.PathParameters["templateId"] = *templateId
    }
    return m
}
// NewTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder instantiates a new TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder and sets the default values.
func NewTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function compare
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a TemplatesItemMigratableToItemCompareWithTemplateIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(TemplatesItemMigratableToItemCompareWithTemplateIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesItemMigratableToItemCompareWithTemplateIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesItemMigratableToItemCompareWithTemplateIdResponseable), nil
}
// GetAsCompareWithTemplateIdGetResponse invoke function compare
// returns a TemplatesItemMigratableToItemCompareWithTemplateIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) GetAsCompareWithTemplateIdGetResponse(ctx context.Context, requestConfiguration *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(TemplatesItemMigratableToItemCompareWithTemplateIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesItemMigratableToItemCompareWithTemplateIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesItemMigratableToItemCompareWithTemplateIdGetResponseable), nil
}
// ToGetRequestInformation invoke function compare
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder when successful
func (m *TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) WithUrl(rawUrl string)(*TemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder) {
    return NewTemplatesItemMigratableToItemCompareWithTemplateIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
