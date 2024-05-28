package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemResourcesItemRoleassignmentsExportRequestBuilder provides operations to call the export method.
type ItemResourcesItemRoleassignmentsExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemResourcesItemRoleassignmentsExportRequestBuilderGetQueryParameters invoke function export
type ItemResourcesItemRoleassignmentsExportRequestBuilderGetQueryParameters struct {
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
// ItemResourcesItemRoleassignmentsExportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesItemRoleassignmentsExportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemResourcesItemRoleassignmentsExportRequestBuilderGetQueryParameters
}
// NewItemResourcesItemRoleassignmentsExportRequestBuilderInternal instantiates a new ItemResourcesItemRoleassignmentsExportRequestBuilder and sets the default values.
func NewItemResourcesItemRoleassignmentsExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleassignmentsExportRequestBuilder) {
    m := &ItemResourcesItemRoleassignmentsExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}/roleAssignments/export(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemResourcesItemRoleassignmentsExportRequestBuilder instantiates a new ItemResourcesItemRoleassignmentsExportRequestBuilder and sets the default values.
func NewItemResourcesItemRoleassignmentsExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleassignmentsExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResourcesItemRoleassignmentsExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function export
// Deprecated: This method is obsolete. Use GetAsExportGetResponse instead.
// returns a ItemResourcesItemRoleassignmentsExportResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesItemRoleassignmentsExportRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemResourcesItemRoleassignmentsExportRequestBuilderGetRequestConfiguration)(ItemResourcesItemRoleassignmentsExportResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemResourcesItemRoleassignmentsExportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemResourcesItemRoleassignmentsExportResponseable), nil
}
// GetAsExportGetResponse invoke function export
// returns a ItemResourcesItemRoleassignmentsExportGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesItemRoleassignmentsExportRequestBuilder) GetAsExportGetResponse(ctx context.Context, requestConfiguration *ItemResourcesItemRoleassignmentsExportRequestBuilderGetRequestConfiguration)(ItemResourcesItemRoleassignmentsExportGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemResourcesItemRoleassignmentsExportGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemResourcesItemRoleassignmentsExportGetResponseable), nil
}
// ToGetRequestInformation invoke function export
// returns a *RequestInformation when successful
func (m *ItemResourcesItemRoleassignmentsExportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesItemRoleassignmentsExportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemResourcesItemRoleassignmentsExportRequestBuilder when successful
func (m *ItemResourcesItemRoleassignmentsExportRequestBuilder) WithUrl(rawUrl string)(*ItemResourcesItemRoleassignmentsExportRequestBuilder) {
    return NewItemResourcesItemRoleassignmentsExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
