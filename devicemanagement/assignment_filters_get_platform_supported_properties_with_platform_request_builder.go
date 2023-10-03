package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder provides operations to call the getPlatformSupportedProperties method.
type AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetQueryParameters invoke function getPlatformSupportedProperties
type AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetQueryParameters struct {
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
// AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetQueryParameters
}
// NewAssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal instantiates a new GetPlatformSupportedPropertiesWithPlatformRequestBuilder and sets the default values.
func NewAssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, platform *string)(*AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    m := &AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/assignmentFilters/getPlatformSupportedProperties(platform='{platform}'){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    if platform != nil {
        m.BaseRequestBuilder.PathParameters["platform"] = *platform
    }
    return m
}
// NewAssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder instantiates a new GetPlatformSupportedPropertiesWithPlatformRequestBuilder and sets the default values.
func NewAssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getPlatformSupportedProperties
// Deprecated: This method is obsolete. Use GetAsGetPlatformSupportedPropertiesWithPlatformGetResponse instead.
func (m *AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder) Get(ctx context.Context, requestConfiguration *AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration)(AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAssignmentFiltersGetPlatformSupportedPropertiesWithPlatformResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformResponseable), nil
}
// GetAsGetPlatformSupportedPropertiesWithPlatformGetResponse invoke function getPlatformSupportedProperties
func (m *AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder) GetAsGetPlatformSupportedPropertiesWithPlatformGetResponse(ctx context.Context, requestConfiguration *AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration)(AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAssignmentFiltersGetPlatformSupportedPropertiesWithPlatformGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformGetResponseable), nil
}
// ToGetRequestInformation invoke function getPlatformSupportedProperties
func (m *AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder) WithUrl(rawUrl string)(*AssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    return NewAssignmentFiltersGetPlatformSupportedPropertiesWithPlatformRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
