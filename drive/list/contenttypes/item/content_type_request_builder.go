package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/publish"
    i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/copytodefaultcontentlocation"
    i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/associatewithhubsites"
    i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnlinks"
    i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/unpublish"
    i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/base"
    i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/basetypes"
    ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/ispublished"
    icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columns"
    ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnpositions"
    ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columns/item"
    if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/contenttypes/item/columnlinks/item"
)

type ContentTypeRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ContentTypeRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a.AssociateWithHubSitesRequestBuilder) {
    return i48e086ed64fcd59d03281d5ffa804a5c862c38e82c9234f5eceee38d4e99a30a.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641.BaseRequestBuilder) {
    return i7bed78d5c7b182bb6b35c0cbd54072e0caa68f8aaee7b2c72ed1530d26e66641.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d.BaseTypesRequestBuilder) {
    return i929fc0841e147d731d330fffac2b603f59d2e7ee09ba0bb76a82c060f484478d.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d.ColumnLinksRequestBuilder) {
    return i5591abfce713804f2fb74b208efc6db0ee225ff736b8e9ac134955aad253bf7d.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return if1bf402e4bfc313e6d11fabbebde44dffbe81a5aa7a12e657549b48143aefbb0.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90.ColumnPositionsRequestBuilder) {
    return ieabef4a072848f73eb7851ae8297bf4213dbcc6b84e915fef10453e9c720ea90.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc.ColumnsRequestBuilder) {
    return icc9cf5371dacb763ed32f4832e68daf65622d2d8a427a8b01fd79665f253c4cc.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ib1338a26b88df4781df2ab470814626fd7c5b61709da7294d63f28bea3263a37.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/drive/list/contentTypes/{contentType_id}{?select,expand}";
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
func NewContentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1.CopyToDefaultContentLocationRequestBuilder) {
    return i2e24155317ddebf9f80799a7e7567793da2868054b1116675108e3b4674a0fc1.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(q func (value *ContentTypeRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ContentTypeRequestBuilderGetQueryParameters)
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
func (m *ContentTypeRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *ContentTypeRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Get(q func (value *ContentTypeRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, error) {
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
func (m *ContentTypeRequestBuilder) IsPublished()(*ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc.IsPublishedRequestBuilder) {
    return ib19674cd06300cabca5bf2c5ea72a1bbc22118fa8aabb80e98be0a419c35e6cc.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Publish()(*i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3.PublishRequestBuilder) {
    return i072c7323df45db1f1ee748560abd200a43db5ebff2606b32cdad943a8247a3e3.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb.UnpublishRequestBuilder) {
    return i69a0f0826d29adb7fdeff42634c0cfeaa9dfca0c790eb1d7972dc8275fdcd7bb.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
