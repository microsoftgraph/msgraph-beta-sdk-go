package messages

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1aa791957161149aab4db21f948fcd8c8275c9ea7c5c4177f0e7d036e3a941c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/count"
    i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/unfavorite"
    i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/favorite"
    i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/markunread"
    ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/markread"
    icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/archive"
    icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/messages/unarchive"
)

// MessagesRequestBuilder provides operations to manage the messages property of the microsoft.graph.serviceAnnouncement entity.
type MessagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessagesRequestBuilderGetQueryParameters retrieve the serviceUpdateMessage resources from the **messages** navigation property. This operation retrieves all service update messages that exist for the tenant.
type MessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MessagesRequestBuilderGetQueryParameters
}
// MessagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Archive the archive property
func (m *MessagesRequestBuilder) Archive()(*icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208.ArchiveRequestBuilder) {
    return icbd43ca604dd53bf065ea6257222903649f6b57753bfe458b714f37a75dca208.NewArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMessagesRequestBuilderInternal instantiates a new MessagesRequestBuilder and sets the default values.
func NewMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessagesRequestBuilder) {
    m := &MessagesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessagesRequestBuilder instantiates a new MessagesRequestBuilder and sets the default values.
func NewMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *MessagesRequestBuilder) Count()(*i1aa791957161149aab4db21f948fcd8c8275c9ea7c5c4177f0e7d036e3a941c6.CountRequestBuilder) {
    return i1aa791957161149aab4db21f948fcd8c8275c9ea7c5c4177f0e7d036e3a941c6.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation retrieve the serviceUpdateMessage resources from the **messages** navigation property. This operation retrieves all service update messages that exist for the tenant.
func (m *MessagesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve the serviceUpdateMessage resources from the **messages** navigation property. This operation retrieves all service update messages that exist for the tenant.
func (m *MessagesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to messages for admin
func (m *MessagesRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceUpdateMessageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to messages for admin
func (m *MessagesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceUpdateMessageable, requestConfiguration *MessagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Favorite the favorite property
func (m *MessagesRequestBuilder) Favorite()(*i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb.FavoriteRequestBuilder) {
    return i5f80f1e7991cbd0c6f9ff2d8c45c4858d507b7bba3f68978bb3fe6dc2af201eb.NewFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the serviceUpdateMessage resources from the **messages** navigation property. This operation retrieves all service update messages that exist for the tenant.
func (m *MessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *MessagesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceUpdateMessageCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceUpdateMessageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceUpdateMessageCollectionResponseable), nil
}
// MarkRead the markRead property
func (m *MessagesRequestBuilder) MarkRead()(*ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e.MarkReadRequestBuilder) {
    return ib2da2e30bf0b9948932c34a14cde77af0959c19fa87f122e064fb88a21ea8e4e.NewMarkReadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MarkUnread the markUnread property
func (m *MessagesRequestBuilder) MarkUnread()(*i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83.MarkUnreadRequestBuilder) {
    return i7bd438807fb58b51e66ef1d0e2bdc47e74daeff3fe1a24ff3fe519435f47cf83.NewMarkUnreadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to messages for admin
func (m *MessagesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceUpdateMessageable, requestConfiguration *MessagesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceUpdateMessageable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceUpdateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceUpdateMessageable), nil
}
// Unarchive the unarchive property
func (m *MessagesRequestBuilder) Unarchive()(*icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c.UnarchiveRequestBuilder) {
    return icc1be3ce3319b01a9ddfc76049beefcc1ccda50b3e6821435cd16b349c738c4c.NewUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unfavorite the unfavorite property
func (m *MessagesRequestBuilder) Unfavorite()(*i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112.UnfavoriteRequestBuilder) {
    return i205a7b20019ec5ab2d4fd32c699b4ba886d1261919e1a47f2ecae1252cd1a112.NewUnfavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
