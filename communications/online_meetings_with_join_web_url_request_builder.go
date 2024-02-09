package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingsWithJoinWebUrlRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.cloudCommunications entity.
type OnlineMeetingsWithJoinWebUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingsWithJoinWebUrlRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsWithJoinWebUrlRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlineMeetingsWithJoinWebUrlRequestBuilderGetQueryParameters get onlineMeetings from communications
type OnlineMeetingsWithJoinWebUrlRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnlineMeetingsWithJoinWebUrlRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsWithJoinWebUrlRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlineMeetingsWithJoinWebUrlRequestBuilderGetQueryParameters
}
// OnlineMeetingsWithJoinWebUrlRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsWithJoinWebUrlRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlineMeetingsWithJoinWebUrlRequestBuilderInternal instantiates a new OnlineMeetingsWithJoinWebUrlRequestBuilder and sets the default values.
func NewOnlineMeetingsWithJoinWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, joinWebUrl *string)(*OnlineMeetingsWithJoinWebUrlRequestBuilder) {
    m := &OnlineMeetingsWithJoinWebUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings(joinWebUrl='{joinWebUrl}'){?%24expand,%24select}", pathParameters),
    }
    if joinWebUrl != nil {
        m.BaseRequestBuilder.PathParameters["joinWebUrl"] = *joinWebUrl
    }
    return m
}
// NewOnlineMeetingsWithJoinWebUrlRequestBuilder instantiates a new OnlineMeetingsWithJoinWebUrlRequestBuilder and sets the default values.
func NewOnlineMeetingsWithJoinWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsWithJoinWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsWithJoinWebUrlRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property onlineMeetings for communications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlineMeetingsWithJoinWebUrlRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get onlineMeetings from communications
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingsWithJoinWebUrlRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// GetVirtualAppointmentJoinWebUrl provides operations to call the getVirtualAppointmentJoinWebUrl method.
// returns a *OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder when successful
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) GetVirtualAppointmentJoinWebUrl()(*OnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilder) {
    return NewOnlineMeetingsWithJoinWebUrlGetVirtualAppointmentJoinWebUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property onlineMeetings in communications
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *OnlineMeetingsWithJoinWebUrlRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// SendVirtualAppointmentReminderSms provides operations to call the sendVirtualAppointmentReminderSms method.
// returns a *OnlineMeetingsWithJoinWebUrlSendVirtualAppointmentReminderSmsRequestBuilder when successful
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) SendVirtualAppointmentReminderSms()(*OnlineMeetingsWithJoinWebUrlSendVirtualAppointmentReminderSmsRequestBuilder) {
    return NewOnlineMeetingsWithJoinWebUrlSendVirtualAppointmentReminderSmsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendVirtualAppointmentSms provides operations to call the sendVirtualAppointmentSms method.
// returns a *OnlineMeetingsWithJoinWebUrlSendVirtualAppointmentSmsRequestBuilder when successful
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) SendVirtualAppointmentSms()(*OnlineMeetingsWithJoinWebUrlSendVirtualAppointmentSmsRequestBuilder) {
    return NewOnlineMeetingsWithJoinWebUrlSendVirtualAppointmentSmsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property onlineMeetings for communications
// returns a *RequestInformation when successful
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsWithJoinWebUrlRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/communications/onlineMeetings(joinWebUrl='{joinWebUrl}')", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get onlineMeetings from communications
// returns a *RequestInformation when successful
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsWithJoinWebUrlRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property onlineMeetings in communications
// returns a *RequestInformation when successful
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *OnlineMeetingsWithJoinWebUrlRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/communications/onlineMeetings(joinWebUrl='{joinWebUrl}')", m.BaseRequestBuilder.PathParameters)
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
// returns a *OnlineMeetingsWithJoinWebUrlRequestBuilder when successful
func (m *OnlineMeetingsWithJoinWebUrlRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingsWithJoinWebUrlRequestBuilder) {
    return NewOnlineMeetingsWithJoinWebUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
