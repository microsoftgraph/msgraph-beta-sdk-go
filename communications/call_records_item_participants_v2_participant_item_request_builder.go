package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/callrecords"
)

// CallRecordsItemParticipants_v2ParticipantItemRequestBuilder provides operations to manage the participants_v2 property of the microsoft.graph.callRecords.callRecord entity.
type CallRecordsItemParticipants_v2ParticipantItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallRecordsItemParticipants_v2ParticipantItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsItemParticipants_v2ParticipantItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallRecordsItemParticipants_v2ParticipantItemRequestBuilderGetQueryParameters list of distinct participants in the call.
type CallRecordsItemParticipants_v2ParticipantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CallRecordsItemParticipants_v2ParticipantItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsItemParticipants_v2ParticipantItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallRecordsItemParticipants_v2ParticipantItemRequestBuilderGetQueryParameters
}
// CallRecordsItemParticipants_v2ParticipantItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallRecordsItemParticipants_v2ParticipantItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallRecordsItemParticipants_v2ParticipantItemRequestBuilderInternal instantiates a new CallRecordsItemParticipants_v2ParticipantItemRequestBuilder and sets the default values.
func NewCallRecordsItemParticipants_v2ParticipantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) {
    m := &CallRecordsItemParticipants_v2ParticipantItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/{callRecord%2Did}/participants_v2/{participant%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCallRecordsItemParticipants_v2ParticipantItemRequestBuilder instantiates a new CallRecordsItemParticipants_v2ParticipantItemRequestBuilder and sets the default values.
func NewCallRecordsItemParticipants_v2ParticipantItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallRecordsItemParticipants_v2ParticipantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property participants_v2 for communications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CallRecordsItemParticipants_v2ParticipantItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of distinct participants in the call.
// returns a Participantable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CallRecordsItemParticipants_v2ParticipantItemRequestBuilderGetRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Participantable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property participants_v2 in communications
// returns a Participantable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) Patch(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Participantable, requestConfiguration *CallRecordsItemParticipants_v2ParticipantItemRequestBuilderPatchRequestConfiguration)(iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Participantable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property participants_v2 for communications
// returns a *RequestInformation when successful
func (m *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CallRecordsItemParticipants_v2ParticipantItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/communications/callRecords/{callRecord%2Did}/participants_v2/{participant%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of distinct participants in the call.
// returns a *RequestInformation when successful
func (m *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallRecordsItemParticipants_v2ParticipantItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property participants_v2 in communications
// returns a *RequestInformation when successful
func (m *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iabe42a55de44a0960e4cc683a105812061defb936fe89e1bc4ab83c390c3839f.Participantable, requestConfiguration *CallRecordsItemParticipants_v2ParticipantItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/communications/callRecords/{callRecord%2Did}/participants_v2/{participant%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder when successful
func (m *CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) WithUrl(rawUrl string)(*CallRecordsItemParticipants_v2ParticipantItemRequestBuilder) {
    return NewCallRecordsItemParticipants_v2ParticipantItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
