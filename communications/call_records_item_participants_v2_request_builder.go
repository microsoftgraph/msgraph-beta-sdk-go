package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
)

// CallRecordsItemParticipants_v2RequestBuilder provides operations to manage the participants_v2 property of the microsoft.graph.callRecords.callRecord entity.
type CallRecordsItemParticipants_v2RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallRecordsItemParticipants_v2RequestBuilderGetQueryParameters list of distinct participants in the call.
type CallRecordsItemParticipants_v2RequestBuilderGetQueryParameters struct {
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
// CallRecordsItemParticipants_v2RequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsItemParticipants_v2RequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallRecordsItemParticipants_v2RequestBuilderGetQueryParameters
}
// CallRecordsItemParticipants_v2RequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsItemParticipants_v2RequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByParticipantId provides operations to manage the participants_v2 property of the microsoft.graph.callRecords.callRecord entity.
// returns a *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder when successful
func (m *CallRecordsItemParticipants_v2RequestBuilder) ByParticipantId(participantId string)(*CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if participantId != "" {
        urlTplParams["participant%2Did"] = participantId
    }
    return NewCallRecordsItemParticipants_v2ParticipantItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCallRecordsItemParticipants_v2RequestBuilderInternal instantiates a new CallRecordsItemParticipants_v2RequestBuilder and sets the default values.
func NewCallRecordsItemParticipants_v2RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordsItemParticipants_v2RequestBuilder) {
    m := &CallRecordsItemParticipants_v2RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/{callRecord%2Did}/participants_v2{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCallRecordsItemParticipants_v2RequestBuilder instantiates a new CallRecordsItemParticipants_v2RequestBuilder and sets the default values.
func NewCallRecordsItemParticipants_v2RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordsItemParticipants_v2RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallRecordsItemParticipants_v2RequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CallRecordsItemParticipants_v2CountRequestBuilder when successful
func (m *CallRecordsItemParticipants_v2RequestBuilder) Count()(*CallRecordsItemParticipants_v2CountRequestBuilder) {
    return NewCallRecordsItemParticipants_v2CountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of distinct participants in the call.
// returns a ParticipantCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallRecordsItemParticipants_v2RequestBuilder) Get(ctx context.Context, requestConfiguration *CallRecordsItemParticipants_v2RequestBuilderGetRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.ParticipantCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateParticipantCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.ParticipantCollectionResponseable), nil
}
// Post create new navigation property to participants_v2 for communications
// returns a Participantable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallRecordsItemParticipants_v2RequestBuilder) Post(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Participantable, requestConfiguration *CallRecordsItemParticipants_v2RequestBuilderPostRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Participantable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.CreateParticipantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Participantable), nil
}
// ToGetRequestInformation list of distinct participants in the call.
// returns a *RequestInformation when successful
func (m *CallRecordsItemParticipants_v2RequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallRecordsItemParticipants_v2RequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to participants_v2 for communications
// returns a *RequestInformation when successful
func (m *CallRecordsItemParticipants_v2RequestBuilder) ToPostRequestInformation(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Participantable, requestConfiguration *CallRecordsItemParticipants_v2RequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CallRecordsItemParticipants_v2RequestBuilder when successful
func (m *CallRecordsItemParticipants_v2RequestBuilder) WithUrl(rawUrl string)(*CallRecordsItemParticipants_v2RequestBuilder) {
    return NewCallRecordsItemParticipants_v2RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
