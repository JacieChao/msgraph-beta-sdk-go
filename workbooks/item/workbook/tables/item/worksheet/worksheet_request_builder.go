package worksheet

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0a95af900300faaed3357b3f622cbbf00989f037fa708defe188d9a16fc51418 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/usedrange"
    i2c38aa9fcf0ef15c13b582943a8ecdb0ca24c5c4953d2374c4c673501da58f63 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/cellwithrowwithcolumn"
    i3ddcdce4491c9d04c88133619f912a234e6c29f49882574f3899a635b9c598c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/pivottables"
    i986fff784f42ffbf5b0746413f92c6f4ccf48b191a61b742e953699e8d3413f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/protection"
    ia1f2ce8b385e4c6a5361e5411d36c56faf61f0a71bb670b41377ca5a5b1b10ae "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/range_escaped"
    ia276ceb85fd09a2117364dbbb767b09b3c5479169f5e4f85a8f8f9bab3a80f6b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables"
    ic89adfad75cf5592a2cf67fbb4dd1075ab6ca605f469fa0591123cbb68a446d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/rangewithaddress"
    idc2bc5c4d1bbb00a5a2c6a8e8b705aa87b682acec7daec80757f8192ef1b758d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/usedrangewithvaluesonly"
    ide4ba23235f326614c8805813356139727b8263a14f3a4d923bb40121a7ba365 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts"
    iee6c370a85cb6bb0fefbc8dc3939ea618f9c1c5957db5c49369a7a35b0d095a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/names"
    i53303b45d7d648df95cc91c6507cc59b5af02e940b75e6a563d137849b1f393d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/names/item"
    ia1cc623eddd9df9f17c0b9701a2a3539b2f5442517a879409721736eb8a27048 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item"
    iec90b216a18e128f226f8a1291a9f75dd1d3bed00c0efd71469331f485a8cb0a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/tables/item"
    if823d55338eadd2201e55a09008562794fbe17cc448d5a324b81519ccff4008a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/pivottables/item"
)

type WorksheetRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorksheetRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *WorksheetRequestBuilder) CellWithRowWithColumn(row *int32, column *int32)(*i2c38aa9fcf0ef15c13b582943a8ecdb0ca24c5c4953d2374c4c673501da58f63.CellWithRowWithColumnRequestBuilder) {
    return i2c38aa9fcf0ef15c13b582943a8ecdb0ca24c5c4953d2374c4c673501da58f63.NewCellWithRowWithColumnRequestBuilderInternal(m.pathParameters, m.requestAdapter, row, column);
}
func (m *WorksheetRequestBuilder) Charts()(*ide4ba23235f326614c8805813356139727b8263a14f3a4d923bb40121a7ba365.ChartsRequestBuilder) {
    return ide4ba23235f326614c8805813356139727b8263a14f3a4d923bb40121a7ba365.NewChartsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) ChartsById(id string)(*ia1cc623eddd9df9f17c0b9701a2a3539b2f5442517a879409721736eb8a27048.WorkbookChartRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChart_id"] = id
    }
    return ia1cc623eddd9df9f17c0b9701a2a3539b2f5442517a879409721736eb8a27048.NewWorkbookChartRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    m := &WorksheetRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWorksheetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorksheetRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorksheetRequestBuilder) CreateGetRequestInformation(q func (value *WorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorksheetRequestBuilderGetQueryParameters)
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
func (m *WorksheetRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorksheetRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorksheetRequestBuilder) Get(q func (value *WorksheetRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookWorksheet() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet), nil
}
func (m *WorksheetRequestBuilder) Names()(*iee6c370a85cb6bb0fefbc8dc3939ea618f9c1c5957db5c49369a7a35b0d095a7.NamesRequestBuilder) {
    return iee6c370a85cb6bb0fefbc8dc3939ea618f9c1c5957db5c49369a7a35b0d095a7.NewNamesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) NamesById(id string)(*i53303b45d7d648df95cc91c6507cc59b5af02e940b75e6a563d137849b1f393d.WorkbookNamedItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookNamedItem_id"] = id
    }
    return i53303b45d7d648df95cc91c6507cc59b5af02e940b75e6a563d137849b1f393d.NewWorkbookNamedItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookWorksheet, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorksheetRequestBuilder) PivotTables()(*i3ddcdce4491c9d04c88133619f912a234e6c29f49882574f3899a635b9c598c2.PivotTablesRequestBuilder) {
    return i3ddcdce4491c9d04c88133619f912a234e6c29f49882574f3899a635b9c598c2.NewPivotTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) PivotTablesById(id string)(*if823d55338eadd2201e55a09008562794fbe17cc448d5a324b81519ccff4008a.WorkbookPivotTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookPivotTable_id"] = id
    }
    return if823d55338eadd2201e55a09008562794fbe17cc448d5a324b81519ccff4008a.NewWorkbookPivotTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Protection()(*i986fff784f42ffbf5b0746413f92c6f4ccf48b191a61b742e953699e8d3413f0.ProtectionRequestBuilder) {
    return i986fff784f42ffbf5b0746413f92c6f4ccf48b191a61b742e953699e8d3413f0.NewProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) Range_escaped()(*ia1f2ce8b385e4c6a5361e5411d36c56faf61f0a71bb670b41377ca5a5b1b10ae.RangeRequestBuilder) {
    return ia1f2ce8b385e4c6a5361e5411d36c56faf61f0a71bb670b41377ca5a5b1b10ae.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) RangeWithAddress(address *string)(*ic89adfad75cf5592a2cf67fbb4dd1075ab6ca605f469fa0591123cbb68a446d9.RangeWithAddressRequestBuilder) {
    return ic89adfad75cf5592a2cf67fbb4dd1075ab6ca605f469fa0591123cbb68a446d9.NewRangeWithAddressRequestBuilderInternal(m.pathParameters, m.requestAdapter, address);
}
func (m *WorksheetRequestBuilder) Tables()(*ia276ceb85fd09a2117364dbbb767b09b3c5479169f5e4f85a8f8f9bab3a80f6b.TablesRequestBuilder) {
    return ia276ceb85fd09a2117364dbbb767b09b3c5479169f5e4f85a8f8f9bab3a80f6b.NewTablesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) TablesById(id string)(*iec90b216a18e128f226f8a1291a9f75dd1d3bed00c0efd71469331f485a8cb0a.WorkbookTableRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookTable_id1"] = id
    }
    return iec90b216a18e128f226f8a1291a9f75dd1d3bed00c0efd71469331f485a8cb0a.NewWorkbookTableRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) UsedRange()(*i0a95af900300faaed3357b3f622cbbf00989f037fa708defe188d9a16fc51418.UsedRangeRequestBuilder) {
    return i0a95af900300faaed3357b3f622cbbf00989f037fa708defe188d9a16fc51418.NewUsedRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorksheetRequestBuilder) UsedRangeWithValuesOnly(valuesOnly *bool)(*idc2bc5c4d1bbb00a5a2c6a8e8b705aa87b682acec7daec80757f8192ef1b758d.UsedRangeWithValuesOnlyRequestBuilder) {
    return idc2bc5c4d1bbb00a5a2c6a8e8b705aa87b682acec7daec80757f8192ef1b758d.NewUsedRangeWithValuesOnlyRequestBuilderInternal(m.pathParameters, m.requestAdapter, valuesOnly);
}