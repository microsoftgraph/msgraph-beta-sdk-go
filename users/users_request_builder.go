package users

import (
    i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/validateproperties"
    i2c83b78d88f4ae5d43180bd3ceb5b2b18f835891b4d285c93889e5b4e7dadfa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/validatepassword"
    i4955c8ead1dfb331d7b6c5cb84c83e17fe654822a84f8a3c4f9681b987a8f982 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/count"
    i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getmanagedappblockedusers"
    i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getuserownedobjects"
    i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/delta"
    ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getbyids"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// UsersRequestBuilder provides operations to manage the collection of user entities.
type UsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UsersRequestBuilderGetOptions options for Get
type UsersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UsersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UsersRequestBuilderGetQueryParameters get entities from users
type UsersRequestBuilderGetQueryParameters struct {
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
// UsersRequestBuilderPostOptions options for Post
type UsersRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Userable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUsersRequestBuilderInternal instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UsersRequestBuilder) {
    m := &UsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersRequestBuilder instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UsersRequestBuilder) Count()(*i4955c8ead1dfb331d7b6c5cb84c83e17fe654822a84f8a3c4f9681b987a8f982.CountRequestBuilder) {
    return i4955c8ead1dfb331d7b6c5cb84c83e17fe654822a84f8a3c4f9681b987a8f982.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from users
func (m *UsersRequestBuilder) CreateGetRequestInformation(options *UsersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to users
func (m *UsersRequestBuilder) CreatePostRequestInformation(options *UsersRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UsersRequestBuilder) Delta()(*i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be.DeltaRequestBuilder) {
    return i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from users
func (m *UsersRequestBuilder) Get(options *UsersRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserCollectionResponseable), nil
}
func (m *UsersRequestBuilder) GetByIds()(*ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef.GetByIdsRequestBuilder) {
    return ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppBlockedUsers provides operations to call the getManagedAppBlockedUsers method.
func (m *UsersRequestBuilder) GetManagedAppBlockedUsers()(*i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275.GetManagedAppBlockedUsersRequestBuilder) {
    return i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275.NewGetManagedAppBlockedUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UsersRequestBuilder) GetUserOwnedObjects()(*i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70.GetUserOwnedObjectsRequestBuilder) {
    return i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to users
func (m *UsersRequestBuilder) Post(options *UsersRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Userable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Userable), nil
}
func (m *UsersRequestBuilder) ValidatePassword()(*i2c83b78d88f4ae5d43180bd3ceb5b2b18f835891b4d285c93889e5b4e7dadfa7.ValidatePasswordRequestBuilder) {
    return i2c83b78d88f4ae5d43180bd3ceb5b2b18f835891b4d285c93889e5b4e7dadfa7.NewValidatePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UsersRequestBuilder) ValidateProperties()(*i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c.ValidatePropertiesRequestBuilder) {
    return i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
