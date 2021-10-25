package base

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i24a5b99ac039073ae8cdc79b9ea77ee0e7d9f2b84d2a33eb98e6739d0cdb618a "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/base/associatewithhubsites"
    i2567d0898f6efd1b1e6856db090ddcdf6c2c5cc9a0358c22f3d850ec6cc1278f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/base/publish"
    i43fc81b930520b7b32f5b1b7b3a81599e21eb0685f51e6fa6d1b7ba637182b2a "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/base/ref"
    i6ae6fc694747e0f41f4cb7a9aedf68da550d37f98d0155a8cf3f9e2e0c383a9c "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/base/unpublish"
    iad65d8bf28c0c3276b9c319f45b4cf3cf6c5610cb1346e981ee7bbf6bca4feb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/base/ispublished"
    ifbb2743a51310b7bd231986655861e65ac5f35e98c7106065d409579c2a532b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/base/copytodefaultcontentlocation"
)

type BaseRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type BaseRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *BaseRequestBuilder) AssociateWithHubSites()(*i24a5b99ac039073ae8cdc79b9ea77ee0e7d9f2b84d2a33eb98e6739d0cdb618a.AssociateWithHubSitesRequestBuilder) {
    return i24a5b99ac039073ae8cdc79b9ea77ee0e7d9f2b84d2a33eb98e6739d0cdb618a.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewBaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    m := &BaseRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/contentTypes/{contentType_id}/base{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewBaseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *BaseRequestBuilder) CopyToDefaultContentLocation()(*ifbb2743a51310b7bd231986655861e65ac5f35e98c7106065d409579c2a532b6.CopyToDefaultContentLocationRequestBuilder) {
    return ifbb2743a51310b7bd231986655861e65ac5f35e98c7106065d409579c2a532b6.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) CreateGetRequestInformation(q func (value *BaseRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(BaseRequestBuilderGetQueryParameters)
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
func (m *BaseRequestBuilder) Get(q func (value *BaseRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContentType() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType), nil
}
func (m *BaseRequestBuilder) IsPublished()(*iad65d8bf28c0c3276b9c319f45b4cf3cf6c5610cb1346e981ee7bbf6bca4feb4.IsPublishedRequestBuilder) {
    return iad65d8bf28c0c3276b9c319f45b4cf3cf6c5610cb1346e981ee7bbf6bca4feb4.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Publish()(*i2567d0898f6efd1b1e6856db090ddcdf6c2c5cc9a0358c22f3d850ec6cc1278f.PublishRequestBuilder) {
    return i2567d0898f6efd1b1e6856db090ddcdf6c2c5cc9a0358c22f3d850ec6cc1278f.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Ref()(*i43fc81b930520b7b32f5b1b7b3a81599e21eb0685f51e6fa6d1b7ba637182b2a.RefRequestBuilder) {
    return i43fc81b930520b7b32f5b1b7b3a81599e21eb0685f51e6fa6d1b7ba637182b2a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseRequestBuilder) Unpublish()(*i6ae6fc694747e0f41f4cb7a9aedf68da550d37f98d0155a8cf3f9e2e0c383a9c.UnpublishRequestBuilder) {
    return i6ae6fc694747e0f41f4cb7a9aedf68da550d37f98d0155a8cf3f9e2e0c383a9c.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
