package groups

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i3ac608d31c4029613f5f3c3d50a2d13a52484d199656da3d51731cf4e049bfb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/groups/count"
    i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/termstore"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// GroupsRequestBuilder provides operations to manage the groups property of the microsoft.graph.termStore.store entity.
type GroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupsRequestBuilderGetOptions options for Get
type GroupsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupsRequestBuilderGetQueryParameters collection of all groups available in the term store.
type GroupsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// GroupsRequestBuilderPostOptions options for Post
type GroupsRequestBuilderPostOptions struct {
    // 
    Body i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.Groupable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGroupsRequestBuilderInternal instantiates a new GroupsRequestBuilder and sets the default values.
func NewGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupsRequestBuilder) {
    m := &GroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/termStore/groups{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsRequestBuilder instantiates a new GroupsRequestBuilder and sets the default values.
func NewGroupsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GroupsRequestBuilder) Count()(*i3ac608d31c4029613f5f3c3d50a2d13a52484d199656da3d51731cf4e049bfb9.CountRequestBuilder) {
    return i3ac608d31c4029613f5f3c3d50a2d13a52484d199656da3d51731cf4e049bfb9.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation collection of all groups available in the term store.
func (m *GroupsRequestBuilder) CreateGetRequestInformation(options *GroupsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to groups for termStore
func (m *GroupsRequestBuilder) CreatePostRequestInformation(options *GroupsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get collection of all groups available in the term store.
func (m *GroupsRequestBuilder) Get(options *GroupsRequestBuilderGetOptions)(i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.CreateGroupCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.GroupCollectionResponseable), nil
}
// Post create new navigation property to groups for termStore
func (m *GroupsRequestBuilder) Post(options *GroupsRequestBuilderPostOptions)(i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.Groupable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.CreateGroupFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.Groupable), nil
}
