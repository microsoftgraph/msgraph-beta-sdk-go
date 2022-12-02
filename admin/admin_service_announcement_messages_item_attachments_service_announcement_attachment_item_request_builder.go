package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder provides operations to manage the attachments property of the microsoft.graph.serviceUpdateMessage entity.
type AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters a collection of serviceAnnouncementAttachments.
type AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetQueryParameters
}
// AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal instantiates a new ServiceAnnouncementAttachmentItemRequestBuilder and sets the default values.
func NewAdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) {
    m := &AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages/{serviceUpdateMessage%2Did}/attachments/{serviceAnnouncementAttachment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder instantiates a new ServiceAnnouncementAttachmentItemRequestBuilder and sets the default values.
func NewAdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the admin entity.
func (m *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Content()(*AdminServiceAnnouncementMessagesItemAttachmentsItemContentRequestBuilder) {
    return NewAdminServiceAnnouncementMessagesItemAttachmentsItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property attachments for admin
func (m *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of serviceAnnouncementAttachments.
func (m *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property attachments in admin
func (m *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property attachments for admin
func (m *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of serviceAnnouncementAttachments.
func (m *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceAnnouncementAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable), nil
}
// Patch update the navigation property attachments in admin
func (m *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsServiceAnnouncementAttachmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceAnnouncementAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementAttachmentable), nil
}
