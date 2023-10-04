package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentFiltersItemGetSupportedPropertiesRequestBuilder provides operations to call the getSupportedProperties method.
type AssignmentFiltersItemGetSupportedPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentFiltersItemGetSupportedPropertiesRequestBuilderGetQueryParameters invoke function getSupportedProperties
type AssignmentFiltersItemGetSupportedPropertiesRequestBuilderGetQueryParameters struct {
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
// AssignmentFiltersItemGetSupportedPropertiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentFiltersItemGetSupportedPropertiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AssignmentFiltersItemGetSupportedPropertiesRequestBuilderGetQueryParameters
}
// NewAssignmentFiltersItemGetSupportedPropertiesRequestBuilderInternal instantiates a new GetSupportedPropertiesRequestBuilder and sets the default values.
func NewAssignmentFiltersItemGetSupportedPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentFiltersItemGetSupportedPropertiesRequestBuilder) {
    m := &AssignmentFiltersItemGetSupportedPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/assignmentFilters/{deviceAndAppManagementAssignmentFilter%2Did}/getSupportedProperties(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewAssignmentFiltersItemGetSupportedPropertiesRequestBuilder instantiates a new GetSupportedPropertiesRequestBuilder and sets the default values.
func NewAssignmentFiltersItemGetSupportedPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentFiltersItemGetSupportedPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentFiltersItemGetSupportedPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getSupportedProperties
// Deprecated: This method is obsolete. Use GetAsGetSupportedPropertiesGetResponse instead.
func (m *AssignmentFiltersItemGetSupportedPropertiesRequestBuilder) Get(ctx context.Context, requestConfiguration *AssignmentFiltersItemGetSupportedPropertiesRequestBuilderGetRequestConfiguration)(AssignmentFiltersItemGetSupportedPropertiesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAssignmentFiltersItemGetSupportedPropertiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignmentFiltersItemGetSupportedPropertiesResponseable), nil
}
// GetAsGetSupportedPropertiesGetResponse invoke function getSupportedProperties
func (m *AssignmentFiltersItemGetSupportedPropertiesRequestBuilder) GetAsGetSupportedPropertiesGetResponse(ctx context.Context, requestConfiguration *AssignmentFiltersItemGetSupportedPropertiesRequestBuilderGetRequestConfiguration)(AssignmentFiltersItemGetSupportedPropertiesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAssignmentFiltersItemGetSupportedPropertiesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignmentFiltersItemGetSupportedPropertiesGetResponseable), nil
}
// ToGetRequestInformation invoke function getSupportedProperties
func (m *AssignmentFiltersItemGetSupportedPropertiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AssignmentFiltersItemGetSupportedPropertiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AssignmentFiltersItemGetSupportedPropertiesRequestBuilder) WithUrl(rawUrl string)(*AssignmentFiltersItemGetSupportedPropertiesRequestBuilder) {
    return NewAssignmentFiltersItemGetSupportedPropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
