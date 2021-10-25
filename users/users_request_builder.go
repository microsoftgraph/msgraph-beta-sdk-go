package users

import (
    i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/validateproperties"
    i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getmanagedappblockedusers"
    i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getuserownedobjects"
    i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/delta"
    ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getbyids"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type UsersRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type UsersRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escaped []string;
    Skip *int32;
    Top *int32;
}
func NewUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UsersRequestBuilder) {
    m := &UsersRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewUsersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UsersRequestBuilder) CreateGetRequestInformation(q func (value *UsersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(UsersRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *UsersRequestBuilder) CreatePostRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.User, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *UsersRequestBuilder) Delta()(*i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be.DeltaRequestBuilder) {
    return i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UsersRequestBuilder) Get(q func (value *UsersRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*UsersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUsersResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*UsersResponse), nil
}
func (m *UsersRequestBuilder) GetByIds()(*ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef.GetByIdsRequestBuilder) {
    return ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UsersRequestBuilder) GetManagedAppBlockedUsers()(*i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275.GetManagedAppBlockedUsersRequestBuilder) {
    return i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275.NewGetManagedAppBlockedUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UsersRequestBuilder) GetUserOwnedObjects()(*i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70.GetUserOwnedObjectsRequestBuilder) {
    return i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UsersRequestBuilder) Post(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.User, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.User, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUser() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.User), nil
}
func (m *UsersRequestBuilder) ValidateProperties()(*i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c.ValidatePropertiesRequestBuilder) {
    return i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
