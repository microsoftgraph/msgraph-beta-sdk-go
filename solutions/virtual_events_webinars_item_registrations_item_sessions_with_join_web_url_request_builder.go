// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder provides operations to manage the sessions property of the microsoft.graph.virtualEventRegistration entity.
type VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderGetQueryParameters sessions for a registration.
type VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderGetQueryParameters
}
// NewVirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderInternal instantiates a new VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, joinWebUrl *string)(*VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder) {
    m := &VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/registrations/{virtualEventRegistration%2Did}/sessions(joinWebUrl='{joinWebUrl}'){?%24expand,%24select}", pathParameters),
    }
    if joinWebUrl != nil {
        m.BaseRequestBuilder.PathParameters["joinWebUrl"] = *joinWebUrl
    }
    return m
}
// NewVirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder instantiates a new VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get sessions for a registration.
// returns a VirtualEventSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable), nil
}
// ToGetRequestInformation sessions for a registration.
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder when successful
func (m *VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationsItemSessionsWithJoinWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
