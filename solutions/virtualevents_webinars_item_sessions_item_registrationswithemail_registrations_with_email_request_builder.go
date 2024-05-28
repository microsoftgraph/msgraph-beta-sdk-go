package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
type VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderGetQueryParameters get registrations from solutions
type VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderGetQueryParameters
}
// NewVirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderInternal instantiates a new VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, email *string)(*VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder) {
    m := &VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/sessions/{virtualEventSession%2Did}/registrations(email='{email}'){?%24expand,%24select}", pathParameters),
    }
    if email != nil {
        m.BaseRequestBuilder.PathParameters["email"] = *email
    }
    return m
}
// NewVirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder instantiates a new VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get registrations from solutions
// returns a VirtualEventRegistrationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationable), nil
}
// ToGetRequestInformation get registrations from solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder when successful
func (m *VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder) {
    return NewVirtualeventsWebinarsItemSessionsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
