package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BookingBusinessesItemCustomQuestionsRequestBuilder provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
type BookingBusinessesItemCustomQuestionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BookingBusinessesItemCustomQuestionsRequestBuilderGetQueryParameters get the bookingCustomQuestion resources associated with a bookingBusiness.
type BookingBusinessesItemCustomQuestionsRequestBuilderGetQueryParameters struct {
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
// BookingBusinessesItemCustomQuestionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemCustomQuestionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingBusinessesItemCustomQuestionsRequestBuilderGetQueryParameters
}
// BookingBusinessesItemCustomQuestionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemCustomQuestionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByBookingCustomQuestionId provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
// returns a *BookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilder when successful
func (m *BookingBusinessesItemCustomQuestionsRequestBuilder) ByBookingCustomQuestionId(bookingCustomQuestionId string)(*BookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if bookingCustomQuestionId != "" {
        urlTplParams["bookingCustomQuestion%2Did"] = bookingCustomQuestionId
    }
    return NewBookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBookingBusinessesItemCustomQuestionsRequestBuilderInternal instantiates a new BookingBusinessesItemCustomQuestionsRequestBuilder and sets the default values.
func NewBookingBusinessesItemCustomQuestionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemCustomQuestionsRequestBuilder) {
    m := &BookingBusinessesItemCustomQuestionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/customQuestions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBookingBusinessesItemCustomQuestionsRequestBuilder instantiates a new BookingBusinessesItemCustomQuestionsRequestBuilder and sets the default values.
func NewBookingBusinessesItemCustomQuestionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemCustomQuestionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesItemCustomQuestionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BookingBusinessesItemCustomQuestionsCountRequestBuilder when successful
func (m *BookingBusinessesItemCustomQuestionsRequestBuilder) Count()(*BookingBusinessesItemCustomQuestionsCountRequestBuilder) {
    return NewBookingBusinessesItemCustomQuestionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the bookingCustomQuestion resources associated with a bookingBusiness.
// returns a BookingCustomQuestionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-list-customquestions?view=graph-rest-1.0
func (m *BookingBusinessesItemCustomQuestionsRequestBuilder) Get(ctx context.Context, requestConfiguration *BookingBusinessesItemCustomQuestionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingCustomQuestionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBookingCustomQuestionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingCustomQuestionCollectionResponseable), nil
}
// Post create a new bookingCustomQuestion object.
// returns a BookingCustomQuestionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bookingbusiness-post-customquestions?view=graph-rest-1.0
func (m *BookingBusinessesItemCustomQuestionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingCustomQuestionable, requestConfiguration *BookingBusinessesItemCustomQuestionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingCustomQuestionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBookingCustomQuestionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingCustomQuestionable), nil
}
// ToGetRequestInformation get the bookingCustomQuestion resources associated with a bookingBusiness.
// returns a *RequestInformation when successful
func (m *BookingBusinessesItemCustomQuestionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesItemCustomQuestionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new bookingCustomQuestion object.
// returns a *RequestInformation when successful
func (m *BookingBusinessesItemCustomQuestionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BookingCustomQuestionable, requestConfiguration *BookingBusinessesItemCustomQuestionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/customQuestions", m.BaseRequestBuilder.PathParameters)
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
// returns a *BookingBusinessesItemCustomQuestionsRequestBuilder when successful
func (m *BookingBusinessesItemCustomQuestionsRequestBuilder) WithUrl(rawUrl string)(*BookingBusinessesItemCustomQuestionsRequestBuilder) {
    return NewBookingBusinessesItemCustomQuestionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
