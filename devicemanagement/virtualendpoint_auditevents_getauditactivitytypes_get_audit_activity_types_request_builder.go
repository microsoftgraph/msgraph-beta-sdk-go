package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder provides operations to call the getAuditActivityTypes method.
type VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderGetQueryParameters get audit activity types by tenant ID.
type VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderGetQueryParameters struct {
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
// VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderGetQueryParameters
}
// NewVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderInternal instantiates a new VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder and sets the default values.
func NewVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder) {
    m := &VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/auditEvents/getAuditActivityTypes(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder instantiates a new VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder and sets the default values.
func NewVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get audit activity types by tenant ID.
// Deprecated: This method is obsolete. Use GetAsGetAuditActivityTypesGetResponse instead.
// returns a VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcauditevent-getauditactivitytypes?view=graph-rest-beta
func (m *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderGetRequestConfiguration)(VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesResponseable), nil
}
// GetAsGetAuditActivityTypesGetResponse get audit activity types by tenant ID.
// returns a VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcauditevent-getauditactivitytypes?view=graph-rest-beta
func (m *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder) GetAsGetAuditActivityTypesGetResponse(ctx context.Context, requestConfiguration *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderGetRequestConfiguration)(VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesGetResponseable), nil
}
// ToGetRequestInformation get audit activity types by tenant ID.
// returns a *RequestInformation when successful
func (m *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder when successful
func (m *VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder) {
    return NewVirtualendpointAuditeventsGetauditactivitytypesGetAuditActivityTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
