package list

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i43d48455d4a3f90f42f7efb37baeed2ed81beefe0ecd17a3198b2e7e74cb37e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes"
    i616acf5322905d44724a74277658e2fbb29de886ec41936610c9749a797d2c1e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/subscriptions"
    i815383e92ce310234c9046f660a2267b3f7ce71b9bd91e00791d5fdcb4c6fa80 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/drive"
    ic2ffaaa4c45cb050337e17bffa8e9e570c425d2bdbc32ffa3f4f20174d8584f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/columns"
    ic7ac1062584b45359ea2ca21560d1896ce081d0787c5ad76e8810dd7d12f47a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities"
    ie4c04524750ad89f45c079100aec7678e2c748e6ebaa485ab81bc9fbac539aa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items"
    i19e80f4d1e56025acb195480cb0d830d9d2555bbd4658dd0c0f6d4debb627d6c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/columns/item"
    i6d9fa267ba65ffc69588b5404ba4f8366dc0856106861fc7605e3228178bafa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item"
    i8848c51df74f0989f6c22e145c2ee13b413f255dbd7be95258420ca5d687fd44 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/contenttypes/item"
    ia59fbfe37a44fd2bc2fa238f5b2ad8db05da0c21b1f29efe1fb2808d1a63369e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/subscriptions/item"
    ic6aa6f1d9d4cfd6e95be88447503f291b59998d7a6197a071f773049847458a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/activities/item"
)

type ListRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ListRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ListRequestBuilder) Activities()(*ic7ac1062584b45359ea2ca21560d1896ce081d0787c5ad76e8810dd7d12f47a6.ActivitiesRequestBuilder) {
    return ic7ac1062584b45359ea2ca21560d1896ce081d0787c5ad76e8810dd7d12f47a6.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ActivitiesById(id string)(*ic6aa6f1d9d4cfd6e95be88447503f291b59998d7a6197a071f773049847458a8.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ic6aa6f1d9d4cfd6e95be88447503f291b59998d7a6197a071f773049847458a8.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Columns()(*ic2ffaaa4c45cb050337e17bffa8e9e570c425d2bdbc32ffa3f4f20174d8584f4.ColumnsRequestBuilder) {
    return ic2ffaaa4c45cb050337e17bffa8e9e570c425d2bdbc32ffa3f4f20174d8584f4.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ColumnsById(id string)(*i19e80f4d1e56025acb195480cb0d830d9d2555bbd4658dd0c0f6d4debb627d6c.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i19e80f4d1e56025acb195480cb0d830d9d2555bbd4658dd0c0f6d4debb627d6c.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/drives/{drive_id}/list{?select,expand}";
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
func NewListRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListRequestBuilder) ContentTypes()(*i43d48455d4a3f90f42f7efb37baeed2ed81beefe0ecd17a3198b2e7e74cb37e9.ContentTypesRequestBuilder) {
    return i43d48455d4a3f90f42f7efb37baeed2ed81beefe0ecd17a3198b2e7e74cb37e9.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ContentTypesById(id string)(*i8848c51df74f0989f6c22e145c2ee13b413f255dbd7be95258420ca5d687fd44.ContentTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["contentType_id"] = id
    }
    return i8848c51df74f0989f6c22e145c2ee13b413f255dbd7be95258420ca5d687fd44.NewContentTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListRequestBuilder) CreateGetRequestInformation(q func (value *ListRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ListRequestBuilderGetQueryParameters)
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
func (m *ListRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListRequestBuilder) Drive()(*i815383e92ce310234c9046f660a2267b3f7ce71b9bd91e00791d5fdcb4c6fa80.DriveRequestBuilder) {
    return i815383e92ce310234c9046f660a2267b3f7ce71b9bd91e00791d5fdcb4c6fa80.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) Get(q func (value *ListRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewList() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List), nil
}
func (m *ListRequestBuilder) Items()(*ie4c04524750ad89f45c079100aec7678e2c748e6ebaa485ab81bc9fbac539aa7.ItemsRequestBuilder) {
    return ie4c04524750ad89f45c079100aec7678e2c748e6ebaa485ab81bc9fbac539aa7.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) ItemsById(id string)(*i6d9fa267ba65ffc69588b5404ba4f8366dc0856106861fc7605e3228178bafa8.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["listItem_id"] = id
    }
    return i6d9fa267ba65ffc69588b5404ba4f8366dc0856106861fc7605e3228178bafa8.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.List, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListRequestBuilder) Subscriptions()(*i616acf5322905d44724a74277658e2fbb29de886ec41936610c9749a797d2c1e.SubscriptionsRequestBuilder) {
    return i616acf5322905d44724a74277658e2fbb29de886ec41936610c9749a797d2c1e.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListRequestBuilder) SubscriptionsById(id string)(*ia59fbfe37a44fd2bc2fa238f5b2ad8db05da0c21b1f29efe1fb2808d1a63369e.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ia59fbfe37a44fd2bc2fa238f5b2ad8db05da0c21b1f29efe1fb2808d1a63369e.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
