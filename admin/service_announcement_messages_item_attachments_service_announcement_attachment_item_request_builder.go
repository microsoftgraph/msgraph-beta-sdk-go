package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder provides operations to manage the attachments property of the microsoft.graph.serviceUpdateMessage entity.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters read the properties and relationships of a serviceAnnouncementAttachment object.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters
}
// ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal instantiates a new ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) {
    m := &ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/{serviceUpdateMessage%2Did}/attachments/{serviceAnnouncementAttachment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder instantiates a new ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the admin entity.
// returns a *ServiceAnnouncementMessagesItemAttachmentsItemContentRequestBuilder when successful
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Content()(*ServiceAnnouncementMessagesItemAttachmentsItemContentRequestBuilder) {
    return NewServiceAnnouncementMessagesItemAttachmentsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property attachments for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a serviceAnnouncementAttachment object.
// returns a ServiceAnnouncementAttachmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceannouncementattachment-get?view=graph-rest-1.0
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceAnnouncementAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable), nil
}
// Patch update the navigation property attachments in admin
// returns a ServiceAnnouncementAttachmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceAnnouncementAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable), nil
}
// ToDeleteRequestInformation delete navigation property attachments for admin
// returns a *RequestInformation when successful
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a serviceAnnouncementAttachment object.
// returns a *RequestInformation when successful
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property attachments in admin
// returns a *RequestInformation when successful
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable, requestConfiguration *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder when successful
func (m *ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) WithUrl(rawUrl string)(*ServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) {
    return NewServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
