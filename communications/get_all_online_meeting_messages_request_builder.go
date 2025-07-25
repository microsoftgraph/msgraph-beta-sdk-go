// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetAllOnlineMeetingMessagesRequestBuilder provides operations to call the getAllOnlineMeetingMessages method.
type GetAllOnlineMeetingMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetAllOnlineMeetingMessagesRequestBuilderGetQueryParameters get all Teams question and answer (Q&A) conversation messages in a tenant. This function returns a snapshot of all Q&A activity in JSON format. The export includes:- The original question or discussion text- The user who posted the message- All replies and responders- Vote counts- Moderation status (pending or dismissed)- Private replies- The meeting ID and organizer ID that are used for mapping to meeting metadata.
type GetAllOnlineMeetingMessagesRequestBuilderGetQueryParameters struct {
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
// GetAllOnlineMeetingMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAllOnlineMeetingMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetAllOnlineMeetingMessagesRequestBuilderGetQueryParameters
}
// NewGetAllOnlineMeetingMessagesRequestBuilderInternal instantiates a new GetAllOnlineMeetingMessagesRequestBuilder and sets the default values.
func NewGetAllOnlineMeetingMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAllOnlineMeetingMessagesRequestBuilder) {
    m := &GetAllOnlineMeetingMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/getAllOnlineMeetingMessages(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGetAllOnlineMeetingMessagesRequestBuilder instantiates a new GetAllOnlineMeetingMessagesRequestBuilder and sets the default values.
func NewGetAllOnlineMeetingMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAllOnlineMeetingMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAllOnlineMeetingMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get all Teams question and answer (Q&A) conversation messages in a tenant. This function returns a snapshot of all Q&A activity in JSON format. The export includes:- The original question or discussion text- The user who posted the message- All replies and responders- Vote counts- Moderation status (pending or dismissed)- Private replies- The meeting ID and organizer ID that are used for mapping to meeting metadata.
// Deprecated: This method is obsolete. Use GetAsGetAllOnlineMeetingMessagesGetResponse instead.
// returns a GetAllOnlineMeetingMessagesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudcommunications-getallonlinemeetingmessages?view=graph-rest-beta
func (m *GetAllOnlineMeetingMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *GetAllOnlineMeetingMessagesRequestBuilderGetRequestConfiguration)(GetAllOnlineMeetingMessagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetAllOnlineMeetingMessagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAllOnlineMeetingMessagesResponseable), nil
}
// GetAsGetAllOnlineMeetingMessagesGetResponse get all Teams question and answer (Q&A) conversation messages in a tenant. This function returns a snapshot of all Q&A activity in JSON format. The export includes:- The original question or discussion text- The user who posted the message- All replies and responders- Vote counts- Moderation status (pending or dismissed)- Private replies- The meeting ID and organizer ID that are used for mapping to meeting metadata.
// returns a GetAllOnlineMeetingMessagesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudcommunications-getallonlinemeetingmessages?view=graph-rest-beta
func (m *GetAllOnlineMeetingMessagesRequestBuilder) GetAsGetAllOnlineMeetingMessagesGetResponse(ctx context.Context, requestConfiguration *GetAllOnlineMeetingMessagesRequestBuilderGetRequestConfiguration)(GetAllOnlineMeetingMessagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetAllOnlineMeetingMessagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAllOnlineMeetingMessagesGetResponseable), nil
}
// ToGetRequestInformation get all Teams question and answer (Q&A) conversation messages in a tenant. This function returns a snapshot of all Q&A activity in JSON format. The export includes:- The original question or discussion text- The user who posted the message- All replies and responders- Vote counts- Moderation status (pending or dismissed)- Private replies- The meeting ID and organizer ID that are used for mapping to meeting metadata.
// returns a *RequestInformation when successful
func (m *GetAllOnlineMeetingMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetAllOnlineMeetingMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetAllOnlineMeetingMessagesRequestBuilder when successful
func (m *GetAllOnlineMeetingMessagesRequestBuilder) WithUrl(rawUrl string)(*GetAllOnlineMeetingMessagesRequestBuilder) {
    return NewGetAllOnlineMeetingMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
