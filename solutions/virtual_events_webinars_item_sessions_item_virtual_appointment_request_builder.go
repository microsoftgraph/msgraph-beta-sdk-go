package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder provides operations to manage the virtualAppointment property of the microsoft.graph.onlineMeeting entity.
type VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderGetQueryParameters read the properties and relationships of a virtualAppointment object.
type VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderGetQueryParameters
}
// VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderInternal instantiates a new VirtualAppointmentRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) {
    m := &VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/sessions/{virtualEventSession%2Did}/virtualAppointment{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder instantiates a new VirtualAppointmentRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-delete?view=graph-rest-1.0
func (m *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-get?view=graph-rest-1.0
func (m *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualAppointmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable), nil
}
// Patch create a new virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-put?view=graph-rest-1.0
func (m *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable, requestConfiguration *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualAppointmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable), nil
}
// ToDeleteRequestInformation delete a virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
func (m *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
func (m *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation create a new virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
func (m *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable, requestConfiguration *VirtualEventsWebinarsItemSessionsItemVirtualAppointmentRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
