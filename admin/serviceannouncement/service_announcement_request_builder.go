package serviceannouncement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i41ae9b5ca5aa3a84169898e2b617c6ea1d88b829fb2659c23e21b770a541a9d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/healthoverviews"
    i6bb2abfee987b65b1b3392b8a360998369927f072881a1c395ccb59e5b16a948 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/issues"
    iba795784a084f257c4f80d391dcf5228c541b81940b1988d008a703891f942fa "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages"
    ia4194987de629e901159377d6071685c8aff20bacfa55b2cdb6c27925b728fa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/issues/item"
    ie72cf01cf6ce5564b072dca815b13dbf9c0f248c55417fb37202d44f8b47ac04 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/item"
    ifeaaa36be4c91c2207be47a8feb667e2c552fe82abae810a59e3babd3be24a98 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/healthoverviews/item"
)

// ServiceAnnouncementRequestBuilder provides operations to manage the serviceAnnouncement property of the microsoft.graph.admin entity.
type ServiceAnnouncementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServiceAnnouncementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceAnnouncementRequestBuilderGetQueryParameters a container for service communications resources. Read-only.
type ServiceAnnouncementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceAnnouncementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceAnnouncementRequestBuilderGetQueryParameters
}
// ServiceAnnouncementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAnnouncementRequestBuilderInternal instantiates a new ServiceAnnouncementRequestBuilder and sets the default values.
func NewServiceAnnouncementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    m := &ServiceAnnouncementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServiceAnnouncementRequestBuilder instantiates a new ServiceAnnouncementRequestBuilder and sets the default values.
func NewServiceAnnouncementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property serviceAnnouncement for admin
func (m *ServiceAnnouncementRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceAnnouncementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a container for service communications resources. Read-only.
func (m *ServiceAnnouncementRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ServiceAnnouncementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property serviceAnnouncement in admin
func (m *ServiceAnnouncementRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementable, requestConfiguration *ServiceAnnouncementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property serviceAnnouncement for admin
func (m *ServiceAnnouncementRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceAnnouncementRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a container for service communications resources. Read-only.
func (m *ServiceAnnouncementRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceAnnouncementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceAnnouncementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementable), nil
}
// HealthOverviews provides operations to manage the healthOverviews property of the microsoft.graph.serviceAnnouncement entity.
func (m *ServiceAnnouncementRequestBuilder) HealthOverviews()(*i41ae9b5ca5aa3a84169898e2b617c6ea1d88b829fb2659c23e21b770a541a9d4.HealthOverviewsRequestBuilder) {
    return i41ae9b5ca5aa3a84169898e2b617c6ea1d88b829fb2659c23e21b770a541a9d4.NewHealthOverviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HealthOverviewsById provides operations to manage the healthOverviews property of the microsoft.graph.serviceAnnouncement entity.
func (m *ServiceAnnouncementRequestBuilder) HealthOverviewsById(id string)(*ifeaaa36be4c91c2207be47a8feb667e2c552fe82abae810a59e3babd3be24a98.ServiceHealthItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealth%2Did"] = id
    }
    return ifeaaa36be4c91c2207be47a8feb667e2c552fe82abae810a59e3babd3be24a98.NewServiceHealthItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Issues provides operations to manage the issues property of the microsoft.graph.serviceAnnouncement entity.
func (m *ServiceAnnouncementRequestBuilder) Issues()(*i6bb2abfee987b65b1b3392b8a360998369927f072881a1c395ccb59e5b16a948.IssuesRequestBuilder) {
    return i6bb2abfee987b65b1b3392b8a360998369927f072881a1c395ccb59e5b16a948.NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IssuesById provides operations to manage the issues property of the microsoft.graph.serviceAnnouncement entity.
func (m *ServiceAnnouncementRequestBuilder) IssuesById(id string)(*ia4194987de629e901159377d6071685c8aff20bacfa55b2cdb6c27925b728fa7.ServiceHealthIssueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealthIssue%2Did"] = id
    }
    return ia4194987de629e901159377d6071685c8aff20bacfa55b2cdb6c27925b728fa7.NewServiceHealthIssueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.serviceAnnouncement entity.
func (m *ServiceAnnouncementRequestBuilder) Messages()(*iba795784a084f257c4f80d391dcf5228c541b81940b1988d008a703891f942fa.MessagesRequestBuilder) {
    return iba795784a084f257c4f80d391dcf5228c541b81940b1988d008a703891f942fa.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.serviceAnnouncement entity.
func (m *ServiceAnnouncementRequestBuilder) MessagesById(id string)(*ie72cf01cf6ce5564b072dca815b13dbf9c0f248c55417fb37202d44f8b47ac04.ServiceUpdateMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceUpdateMessage%2Did"] = id
    }
    return ie72cf01cf6ce5564b072dca815b13dbf9c0f248c55417fb37202d44f8b47ac04.NewServiceUpdateMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property serviceAnnouncement in admin
func (m *ServiceAnnouncementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementable, requestConfiguration *ServiceAnnouncementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceAnnouncementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceAnnouncementable), nil
}
