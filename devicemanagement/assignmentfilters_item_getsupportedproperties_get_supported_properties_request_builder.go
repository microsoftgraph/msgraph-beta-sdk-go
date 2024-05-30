package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder provides operations to call the getSupportedProperties method.
type AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderGetQueryParameters invoke function getSupportedProperties
type AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderGetQueryParameters struct {
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
// AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderGetQueryParameters
}
// NewAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderInternal instantiates a new AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder and sets the default values.
func NewAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder) {
    m := &AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/assignmentFilters/{deviceAndAppManagementAssignmentFilter%2Did}/getSupportedProperties(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder instantiates a new AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder and sets the default values.
func NewAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getSupportedProperties
// Deprecated: This method is obsolete. Use GetAsGetSupportedPropertiesGetResponse instead.
// returns a AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder) Get(ctx context.Context, requestConfiguration *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderGetRequestConfiguration)(AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesResponseable), nil
}
// GetAsGetSupportedPropertiesGetResponse invoke function getSupportedProperties
// returns a AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder) GetAsGetSupportedPropertiesGetResponse(ctx context.Context, requestConfiguration *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderGetRequestConfiguration)(AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesGetResponseable), nil
}
// ToGetRequestInformation invoke function getSupportedProperties
// returns a *RequestInformation when successful
func (m *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder when successful
func (m *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder) WithUrl(rawUrl string)(*AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder) {
    return NewAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
