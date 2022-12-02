package bookingbusinesses

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BookingBusinessesItemUnpublishRequestBuilder provides operations to call the unpublish method.
type BookingBusinessesItemUnpublishRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookingBusinessesItemUnpublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemUnpublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBookingBusinessesItemUnpublishRequestBuilderInternal instantiates a new UnpublishRequestBuilder and sets the default values.
func NewBookingBusinessesItemUnpublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemUnpublishRequestBuilder) {
    m := &BookingBusinessesItemUnpublishRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/bookingBusinesses/{bookingBusiness%2Did}/microsoft.graph.unpublish";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookingBusinessesItemUnpublishRequestBuilder instantiates a new UnpublishRequestBuilder and sets the default values.
func NewBookingBusinessesItemUnpublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesItemUnpublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesItemUnpublishRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation make the scheduling page of this business not available to external customers. Set the **isPublished** property to false, and **publicUrl** property to null.
func (m *BookingBusinessesItemUnpublishRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesItemUnpublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post make the scheduling page of this business not available to external customers. Set the **isPublished** property to false, and **publicUrl** property to null.
func (m *BookingBusinessesItemUnpublishRequestBuilder) Post(ctx context.Context, requestConfiguration *BookingBusinessesItemUnpublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
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
