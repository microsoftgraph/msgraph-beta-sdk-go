package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder provides operations to call the getPlatformSupportedProperties method.
type AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetQueryParameters invoke function getPlatformSupportedProperties
type AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetQueryParameters struct {
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
// AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetQueryParameters
}
// NewAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal instantiates a new AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder and sets the default values.
func NewAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, platform *string)(*AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    m := &AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/assignmentFilters/getPlatformSupportedProperties(platform='{platform}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if platform != nil {
        m.BaseRequestBuilder.PathParameters["platform"] = *platform
    }
    return m
}
// NewAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder instantiates a new AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder and sets the default values.
func NewAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getPlatformSupportedProperties
// Deprecated: This method is obsolete. Use GetAsGetPlatformSupportedPropertiesWithPlatformGetResponse instead.
// returns a AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder) Get(ctx context.Context, requestConfiguration *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration)(AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformResponseable), nil
}
// GetAsGetPlatformSupportedPropertiesWithPlatformGetResponse invoke function getPlatformSupportedProperties
// returns a AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder) GetAsGetPlatformSupportedPropertiesWithPlatformGetResponse(ctx context.Context, requestConfiguration *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration)(AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformGetResponseable), nil
}
// ToGetRequestInformation invoke function getPlatformSupportedProperties
// returns a *RequestInformation when successful
func (m *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder when successful
func (m *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder) WithUrl(rawUrl string)(*AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    return NewAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
