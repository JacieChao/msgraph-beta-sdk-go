package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i004be3867596cb38b67288439f8616ebc5f3a4c8649aaf07d480b72367da3970 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/totalrowrange"
    i3369442607bcaf08199f59c55db37e37c05065a5728c7248042881b859ceeb90 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/databodyrange"
    i79a6a3c0b80f2ac094f87b3fe87c31493fd136dfcd5cfd6a7913960efe86ff82 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet"
    i8137a7430ecf6f2bc82484c2017370d5429f9ea963d557344da321942fb27cf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/converttorange"
    i8701821bc3f8584c9d04b7577ebef04893b2c2f1f701ab9dbf8cef0fa75aa9d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/range_escaped"
    i8bc422b58ac52eb24ebefaba090e07fad1c12191d8aac96f2aa3073497974f74 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns"
    i8fb64ec9a3bd0ae69eadf0df70a0d0f2f556cb108870d620c0e7f6b8115e35bf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/reapplyfilters"
    iabd66df5e0ed60aa469a19ba02b1077f32538932106090e70a4e118a14dd6bf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/headerrowrange"
    ic099875b03bb18d4e62b09ea7b9b450bfcd8f4a15c5040cacfb802b26852b36d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/sort"
    ief003cdeca5162b5e32434942f4100a37c5d27e4493efcc3434d8ae1a1e75e75 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/clearfilters"
    iefb4255a7d9d563de3e077b3be82fafefa3007ff27d26cd15affb3f81ca7bdd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/rows"
    i2ff7d61355f35acfab76bec2082848fec1887425cd85200de3f6ddae8cda2847 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/rows/item"
    i86a6a19eac431ddcf501212e02e852ddcc92d1eec422d0b5e4c97df9673e9207 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/columns/item"
)

type WorkbookTableRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorkbookTableRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *WorkbookTableRequestBuilder) ClearFilters()(*ief003cdeca5162b5e32434942f4100a37c5d27e4493efcc3434d8ae1a1e75e75.ClearFiltersRequestBuilder) {
    return ief003cdeca5162b5e32434942f4100a37c5d27e4493efcc3434d8ae1a1e75e75.NewClearFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Columns()(*i8bc422b58ac52eb24ebefaba090e07fad1c12191d8aac96f2aa3073497974f74.ColumnsRequestBuilder) {
    return i8bc422b58ac52eb24ebefaba090e07fad1c12191d8aac96f2aa3073497974f74.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ColumnsById(id string)(*i86a6a19eac431ddcf501212e02e852ddcc92d1eec422d0b5e4c97df9673e9207.WorkbookTableColumnRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableColumn_id"] = id
    }
    return i86a6a19eac431ddcf501212e02e852ddcc92d1eec422d0b5e4c97df9673e9207.NewWorkbookTableColumnRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorkbookTableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    m := &WorkbookTableRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWorkbookTableRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookTableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookTableRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookTableRequestBuilder) ConvertToRange()(*i8137a7430ecf6f2bc82484c2017370d5429f9ea963d557344da321942fb27cf4.ConvertToRangeRequestBuilder) {
    return i8137a7430ecf6f2bc82484c2017370d5429f9ea963d557344da321942fb27cf4.NewConvertToRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookTableRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookTableRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookTableRequestBuilderGetQueryParameters)
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
func (m *WorkbookTableRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookTableRequestBuilder) DataBodyRange()(*i3369442607bcaf08199f59c55db37e37c05065a5728c7248042881b859ceeb90.DataBodyRangeRequestBuilder) {
    return i3369442607bcaf08199f59c55db37e37c05065a5728c7248042881b859ceeb90.NewDataBodyRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookTableRequestBuilder) Get(q func (value *WorkbookTableRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookTable() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable), nil
}
func (m *WorkbookTableRequestBuilder) HeaderRowRange()(*iabd66df5e0ed60aa469a19ba02b1077f32538932106090e70a4e118a14dd6bf5.HeaderRowRangeRequestBuilder) {
    return iabd66df5e0ed60aa469a19ba02b1077f32538932106090e70a4e118a14dd6bf5.NewHeaderRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTable, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookTableRequestBuilder) Range_escaped()(*i8701821bc3f8584c9d04b7577ebef04893b2c2f1f701ab9dbf8cef0fa75aa9d0.RangeRequestBuilder) {
    return i8701821bc3f8584c9d04b7577ebef04893b2c2f1f701ab9dbf8cef0fa75aa9d0.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) ReapplyFilters()(*i8fb64ec9a3bd0ae69eadf0df70a0d0f2f556cb108870d620c0e7f6b8115e35bf.ReapplyFiltersRequestBuilder) {
    return i8fb64ec9a3bd0ae69eadf0df70a0d0f2f556cb108870d620c0e7f6b8115e35bf.NewReapplyFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Rows()(*iefb4255a7d9d563de3e077b3be82fafefa3007ff27d26cd15affb3f81ca7bdd5.RowsRequestBuilder) {
    return iefb4255a7d9d563de3e077b3be82fafefa3007ff27d26cd15affb3f81ca7bdd5.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) RowsById(id string)(*i2ff7d61355f35acfab76bec2082848fec1887425cd85200de3f6ddae8cda2847.WorkbookTableRowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTableRow_id"] = id
    }
    return i2ff7d61355f35acfab76bec2082848fec1887425cd85200de3f6ddae8cda2847.NewWorkbookTableRowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Sort()(*ic099875b03bb18d4e62b09ea7b9b450bfcd8f4a15c5040cacfb802b26852b36d.SortRequestBuilder) {
    return ic099875b03bb18d4e62b09ea7b9b450bfcd8f4a15c5040cacfb802b26852b36d.NewSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) TotalRowRange()(*i004be3867596cb38b67288439f8616ebc5f3a4c8649aaf07d480b72367da3970.TotalRowRangeRequestBuilder) {
    return i004be3867596cb38b67288439f8616ebc5f3a4c8649aaf07d480b72367da3970.NewTotalRowRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookTableRequestBuilder) Worksheet()(*i79a6a3c0b80f2ac094f87b3fe87c31493fd136dfcd5cfd6a7913960efe86ff82.WorksheetRequestBuilder) {
    return i79a6a3c0b80f2ac094f87b3fe87c31493fd136dfcd5cfd6a7913960efe86ff82.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}