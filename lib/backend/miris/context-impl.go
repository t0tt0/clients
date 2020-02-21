package miris

import (
	"fmt"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/serial"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-gonic/gin/binding"
	"github.com/kataras/iris"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type Context struct {
	iris.Context
	vars map[string]interface{}
	meta interface{}
}

func (c Context) IsAborted() bool {
	return c.Context.IsStopped()
}

func (c Context) Abort() {
	c.Context.StopExecution()
}

func (c Context) AbortWithStatus(code int) {
	c.Context.StatusCode(code)
	c.Abort()
}

func (c Context) AbortWithStatusJSON(code int, jsonObj interface{}) {
	c.Context.StatusCode(code)
	procSerializeError(c.Context.JSON(jsonObj))
	c.Abort()
	return
}

func (c Context) AbortWithError(code int, err error) controller.Error {
	c.Context.StatusCode(code)
	procSerializeError(c.Context.JSON(serial.ErrorSerializer{
		Code: -1,
		Err:  err.Error(),
	}))
	c.Abort()
	return IrisError{
		Err:  err,
		Type: controller.ErrorTypeAny,
		Meta: nil,
	}
}

func (c Context) Error(err error) controller.Error {
	return IrisError{
		Err:  err,
		Type: controller.ErrorTypeAny,
		Meta: nil,
	}
}

func (c Context) Set(key string, value interface{}) {
	if c.vars == nil {
		c.vars = make(map[string]interface{})
	}
	c.vars[key] = value
}

func (c Context) Get(key string) (value interface{}, exists bool) {
	value, exists = c.vars[key]
	return
}

func (c Context) MustGet(key string) interface{} {
	return c.vars[key]
}

func (c Context) GetString(key string) (s string) {
	return c.MustGet(key).(string)
}

func (c Context) GetBool(key string) (b bool) {
	return c.MustGet(key).(bool)
}

func (c Context) GetInt(key string) (i int) {
	return c.MustGet(key).(int)
}

func (c Context) GetInt64(key string) (i64 int64) {
	return c.MustGet(key).(int64)
}

func (c Context) GetFloat64(key string) (f64 float64) {
	return c.MustGet(key).(float64)
}

func (c Context) GetTime(key string) (t time.Time) {
	return c.MustGet(key).(time.Time)
}

func (c Context) GetDuration(key string) (d time.Duration) {
	return c.MustGet(key).(time.Duration)
}

func (c Context) GetStringSlice(key string) (ss []string) {
	return c.MustGet(key).([]string)
}

func (c Context) GetStringMap(key string) (sm map[string]interface{}) {
	return c.MustGet(key).(map[string]interface{})
}

func (c Context) GetStringMapString(key string) (sms map[string]string) {
	return c.MustGet(key).(map[string]string)
}

func (c Context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	return c.MustGet(key).(map[string][]string)
}

func (c Context) Param(key string) string {
	return c.Context.Params().Get(key)
}

func (c Context) Query(key string) string {
	return c.Context.URLParam(key)
}

func (c Context) DefaultQuery(key, defaultValue string) string {
	return c.Context.URLParamDefault(key, defaultValue)
}

func (c Context) GetQuery(key string) (value string, exists bool) {
	value, exists = c.Context.URLParams()[key]
	return
}

func (c Context) PostForm(key string) string {
	return c.Context.FormValue(key)
}

func (c Context) DefaultPostForm(key, defaultValue string) string {
	return c.Context.FormValueDefault(key, defaultValue)
}

func (c Context) GetPostForm(key string) (value string, exists bool) {
	var values []string
	values, exists = c.Context.FormValues()[key]
	if exists && len(values) >= 1 {
		return values[0], true
	}
	return
}

func (c Context) PostFormArray(key string) []string {
	return c.FormValues()[key]
}

func (c Context) GetPostFormArray(key string) (values []string, exists bool) {
	values, exists = c.Context.FormValues()[key]
	return
}

func (c Context) FormFile(name string) (header *multipart.FileHeader, err error) {
	_, header, err = c.Context.FormFile(name)
	return
}

func (c Context) Bind(obj interface{}) error {
	return binding.JSON.Bind(c.Context.Request(), obj)
}

func (c Context) BindJSON(obj interface{}) error {
	return binding.JSON.Bind(c.Context.Request(), obj)
}

func (c Context) BindXML(obj interface{}) error {
	return binding.XML.Bind(c.Context.Request(), obj)
}

func (c Context) BindQuery(obj interface{}) error {
	return binding.Query.Bind(c.Context.Request(), obj)
}

func (c Context) BindYAML(obj interface{}) error {
	return binding.YAML.Bind(c.Context.Request(), obj)
}

func (c Context) BindHeader(obj interface{}) error {
	return binding.Header.Bind(c.Context.Request(), obj)
}

func (c Context) ClientIP() string {
	return c.Context.RemoteAddr()
}

func (c Context) ContentType() string {
	return c.Context.GetContentType()
}

func (c Context) Status(code int) {
	c.Context.StatusCode(code)
	return
}

func (c Context) MultipartForm() (*multipart.Form, error) {
	panic("implement me")
}

func (c Context) PostFormMap(key string) map[string]string {
	panic("implement me")
}

func (c Context) GetPostFormMap(key string) (map[string]string, bool) {
	panic("implement me")
}

func (c Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	panic("todo")
}

func (c Context) IsWebsocket() bool {
	panic("implement me")
}

func (c Context) BindUri(obj interface{}) error {
	panic("implement me")
}

func (c Context) ShouldBind(obj interface{}) error {
	return binding.JSON.Bind(c.Context.Request(), obj)
}

func (c Context) ShouldBindJSON(obj interface{}) error {
	return binding.JSON.Bind(c.Context.Request(), obj)
}

func (c Context) ShouldBindXML(obj interface{}) error {
	return binding.XML.Bind(c.Context.Request(), obj)
}

func (c Context) ShouldBindQuery(obj interface{}) error {
	return binding.Query.Bind(c.Context.Request(), obj)
}

func (c Context) ShouldBindYAML(obj interface{}) error {
	return binding.YAML.Bind(c.Context.Request(), obj)
}

func (c Context) ShouldBindHeader(obj interface{}) error {
	return binding.Header.Bind(c.Context.Request(), obj)
}

func (c Context) ShouldBindUri(obj interface{}) error {
	panic("implement me")
}

func (c Context) GetRawData() ([]byte, error) {
	panic("implement me")
}

func (c Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	panic("implement me")
}

func (c Context) Cookie(name string) (string, error) {
	return c.Context.GetCookie(name), nil
}

func (c Context) HTML(code int, name string, obj interface{}) {
	//c.Context.HTML()
}

func (c Context) IndentedJSON(code int, obj interface{}) {
	panic("implement me")
}

func (c Context) SecureJSON(code int, obj interface{}) {
	panic("implement me")
}

func (c Context) JSONP(code int, obj interface{}) {
	panic("implement me")
}

func (c Context) JSON(code int, obj interface{}) {
	c.StatusCode(code)
	procSerializeError(c.Context.JSON(obj))
	return
}

func (c Context) AsciiJSON(code int, obj interface{}) {
	panic("implement me")
}

func (c Context) PureJSON(code int, obj interface{}) {
	panic("implement me")
}

func (c Context) XML(code int, obj interface{}) {
	panic("implement me")
}

func (c Context) YAML(code int, obj interface{}) {
	panic("implement me")
}

func (c Context) ProtoBuf(code int, obj interface{}) {
	panic("implement me")
}

func (c Context) String(code int, format string, values ...interface{}) {
	panic("implement me")
}

func (c Context) Redirect(code int, location string) {
	c.Context.Redirect(location, code)
}

func (c Context) Data(code int, contentType string, data []byte) {
	panic("implement me")
}

func (c Context) DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) {
	panic("implement me")
}

func (c Context) File(filepath string) {
	panic("implement me")
}

func (c Context) FileAttachment(filepath, filename string) {
	panic("implement me")
}

func (c Context) SSEvent(name string, message interface{}) {
	panic("implement me")
}

func (c Context) Stream(step func(w io.Writer) bool) bool {
	panic("implement me")
}

func (c Context) SetAccepted(formats ...string) {
	panic("implement me")
}

func (c Context) GetMeta() interface{} {
	return c.meta
}

func (c Context) SetMeta(meta interface{}) {
	c.meta = meta
	return
}

func (c Context) HandlerNames() []string {
	panic("implement me")
}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (c Context) Done() <-chan struct{} {
	panic("implement me")
}

func (c Context) Err() error {
	panic("implement me")
}

func (c Context) Value(key interface{}) interface{} {
	panic("implement me")
}

func (c Context) Copy() controller.MContext {
	return NewRaw(c.Request())
}

func NewRaw(request *http.Request) controller.MContext {
	panic("implement me")
}

func (c Context) FullPath() string {
	panic("implement me")
}

func (c Context) QueryArray(key string) []string {
	fmt.Println(c.Context.URLParam(key))
	panic("implement me")
	return nil
}

func (c Context) GetQueryArray(key string) ([]string, bool) {
	panic("implement me")
}

func (c Context) QueryMap(key string) map[string]string {
	panic("implement me")
}

func (c Context) GetQueryMap(key string) (map[string]string, bool) {
	panic("implement me")
}
