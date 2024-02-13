package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemResourcesItemRoleAssignmentsExportRequestBuilder provides operations to call the export method.
type ItemResourcesItemRoleAssignmentsExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemResourcesItemRoleAssignmentsExportRequestBuilderGetQueryParameters invoke function export
type ItemResourcesItemRoleAssignmentsExportRequestBuilderGetQueryParameters struct {
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
// ItemResourcesItemRoleAssignmentsExportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesItemRoleAssignmentsExportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemResourcesItemRoleAssignmentsExportRequestBuilderGetQueryParameters
}
// NewItemResourcesItemRoleAssignmentsExportRequestBuilderInternal instantiates a new ItemResourcesItemRoleAssignmentsExportRequestBuilder and sets the default values.
func NewItemResourcesItemRoleAssignmentsExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleAssignmentsExportRequestBuilder) {
    m := &ItemResourcesItemRoleAssignmentsExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}/roleAssignments/export(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemResourcesItemRoleAssignmentsExportRequestBuilder instantiates a new ItemResourcesItemRoleAssignmentsExportRequestBuilder and sets the default values.
func NewItemResourcesItemRoleAssignmentsExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleAssignmentsExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResourcesItemRoleAssignmentsExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function export
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ItemResourcesItemRoleAssignmentsExportResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesItemRoleAssignmentsExportRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemResourcesItemRoleAssignmentsExportRequestBuilderGetRequestConfiguration)(ItemResourcesItemRoleAssignmentsExportResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemResourcesItemRoleAssignmentsExportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemResourcesItemRoleAssignmentsExportResponseable), nil
}
// GetAsExportGetResponse invoke function export
// returns a ItemResourcesItemRoleAssignmentsExportGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesItemRoleAssignmentsExportRequestBuilder) GetAsExportGetResponse(ctx context.Context, requestConfiguration *ItemResourcesItemRoleAssignmentsExportRequestBuilderGetRequestConfiguration)(ItemResourcesItemRoleAssignmentsExportGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemResourcesItemRoleAssignmentsExportGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemResourcesItemRoleAssignmentsExportGetResponseable), nil
}
// ToGetRequestInformation invoke function export
// returns a *RequestInformation when successful
func (m *ItemResourcesItemRoleAssignmentsExportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemResourcesItemRoleAssignmentsExportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemResourcesItemRoleAssignmentsExportRequestBuilder when successful
func (m *ItemResourcesItemRoleAssignmentsExportRequestBuilder) WithUrl(rawUrl string)(*ItemResourcesItemRoleAssignmentsExportRequestBuilder) {
    return NewItemResourcesItemRoleAssignmentsExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
