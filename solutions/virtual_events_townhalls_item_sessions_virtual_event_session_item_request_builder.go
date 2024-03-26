package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder provides operations to manage the sessions property of the microsoft.graph.virtualEvent entity.
type VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters sessions for the virtual event.
type VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderGetQueryParameters
}
// VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceReports provides operations to manage the attendanceReports property of the microsoft.graph.onlineMeetingBase entity.
// returns a *VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) AttendanceReports()(*VirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemAttendanceReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderInternal instantiates a new VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) {
    m := &VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/sessions/{virtualEventSession%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder instantiates a new VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder and sets the default values.
func NewVirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sessions for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get sessions for the virtual event.
// returns a VirtualEventSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
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
// Patch update the navigation property sessions in solutions
// returns a VirtualEventSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, requestConfiguration *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Presenters provides operations to manage the presenters property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) Presenters()(*VirtualEventsTownhallsItemSessionsItemPresentersRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemPresentersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Registrations provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsTownhallsItemSessionsItemRegistrationsRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) Registrations()(*VirtualEventsTownhallsItemSessionsItemRegistrationsRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RegistrationsWithEmail provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsTownhallsItemSessionsItemRegistrationsWithEmailRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) RegistrationsWithEmail(email *string)(*VirtualEventsTownhallsItemSessionsItemRegistrationsWithEmailRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemRegistrationsWithEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, email)
}
// RegistrationsWithUserId provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualEventsTownhallsItemSessionsItemRegistrationsWithUserIdRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) RegistrationsWithUserId(userId *string)(*VirtualEventsTownhallsItemSessionsItemRegistrationsWithUserIdRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsItemRegistrationsWithUserIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, userId)
}
// ToDeleteRequestInformation delete navigation property sessions for solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/sessions/{virtualEventSession%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation sessions for the virtual event.
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sessions in solutions
// returns a *RequestInformation when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventSessionable, requestConfiguration *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/solutions/virtualEvents/townhalls/{virtualEventTownhall%2Did}/sessions/{virtualEventSession%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder when successful
func (m *VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder) {
    return NewVirtualEventsTownhallsItemSessionsVirtualEventSessionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
