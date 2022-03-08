package exceptionoccurrences

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i42bf19c590c8453c05add5e0d0c0554f2ac1877f283f745258f7a69e194c5772 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/delta"
    i5297b7bfcd041993b9d962a88cb0156bf43da60e03e3fcd29519520babaf32a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/count"
)

// ExceptionOccurrencesRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ExceptionOccurrencesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExceptionOccurrencesRequestBuilderGetOptions options for Get
type ExceptionOccurrencesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ExceptionOccurrencesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExceptionOccurrencesRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type ExceptionOccurrencesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// ExceptionOccurrencesRequestBuilderPostOptions options for Post
type ExceptionOccurrencesRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewExceptionOccurrencesRequestBuilderInternal instantiates a new ExceptionOccurrencesRequestBuilder and sets the default values.
func NewExceptionOccurrencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExceptionOccurrencesRequestBuilder) {
    m := &ExceptionOccurrencesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/events/{event_id}/exceptionOccurrences{?top,skip,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExceptionOccurrencesRequestBuilder instantiates a new ExceptionOccurrencesRequestBuilder and sets the default values.
func NewExceptionOccurrencesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExceptionOccurrencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExceptionOccurrencesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExceptionOccurrencesRequestBuilder) Count()(*i5297b7bfcd041993b9d962a88cb0156bf43da60e03e3fcd29519520babaf32a0.CountRequestBuilder) {
    return i5297b7bfcd041993b9d962a88cb0156bf43da60e03e3fcd29519520babaf32a0.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get exceptionOccurrences from groups
func (m *ExceptionOccurrencesRequestBuilder) CreateGetRequestInformation(options *ExceptionOccurrencesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to exceptionOccurrences for groups
func (m *ExceptionOccurrencesRequestBuilder) CreatePostRequestInformation(options *ExceptionOccurrencesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *ExceptionOccurrencesRequestBuilder) Delta()(*i42bf19c590c8453c05add5e0d0c0554f2ac1877f283f745258f7a69e194c5772.DeltaRequestBuilder) {
    return i42bf19c590c8453c05add5e0d0c0554f2ac1877f283f745258f7a69e194c5772.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
func (m *ExceptionOccurrencesRequestBuilder) Get(options *ExceptionOccurrencesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EventCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EventCollectionResponseable), nil
}
// Post create new navigation property to exceptionOccurrences for groups
func (m *ExceptionOccurrencesRequestBuilder) Post(options *ExceptionOccurrencesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
