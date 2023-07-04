package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder provides operations to manage the virtualAppointment property of the microsoft.graph.onlineMeeting entity.
type VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderGetQueryParameters read the properties and relationships of a virtualAppointment object.
type VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderGetQueryParameters
}
// VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderInternal instantiates a new VirtualAppointmentRequestBuilder and sets the default values.
func NewVirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder) {
    m := &VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/sessions/{virtualEventSession%2Did}/virtualAppointment{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder instantiates a new VirtualAppointmentRequestBuilder and sets the default values.
func NewVirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/virtualappointment-delete?view=graph-rest-1.0
func (m *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderDeleteRequestConfiguration)(error) {
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
// [Find more info here]: https://docs.microsoft.com/graph/api/virtualappointment-get?view=graph-rest-1.0
func (m *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable, error) {
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
// Patch update the properties of a virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/virtualappointment-update?view=graph-rest-1.0
func (m *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable, requestConfiguration *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable, error) {
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
func (m *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a virtualAppointment object.
// Deprecated: The Virtual appointment resource is deprecated and will stop returning data on May 31, 2023. Existing apps that use this feature should be updated to the new getVirtualAppointmentJoinWebUrl API. as of 2023-04/VirtualAppointment on 2023-04-04 and will be removed 2023-05-31
func (m *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualAppointmentable, requestConfiguration *VirtualEventsEventsItemSessionsItemVirtualAppointmentRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
