package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder provides operations to manage the media for the organization entity.
type ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderInternal instantiates a new SquareLogoDarkRequestBuilder and sets the default values.
func NewItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) {
    m := &ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/branding/localizations/{organizationalBrandingLocalization%2Did}/squareLogoDark", pathParameters),
    }
    return m
}
// NewItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder instantiates a new SquareLogoDarkRequestBuilder and sets the default values.
func NewItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a square dark version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/organizationalbranding-list-localizations?view=graph-rest-1.0
func (m *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put a square dark version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) Put(ctx context.Context, body []byte, contentType *string, requestConfiguration *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, contentType, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation a square dark version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "image/bmp, image/jpg, image/jpeg, image/gif, image/vnd.microsoft.icon, image/png, image/tiff, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation a square dark version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG not larger than 240 x 240 pixels and not more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, contentType *string, requestConfiguration *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, *contentType)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) WithUrl(rawUrl string)(*ItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder) {
    return NewItemBrandingLocalizationsItemSquareLogoDarkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
