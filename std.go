// Generated by generate/std

package interfacer

var stdIfaces = map[string]string{
	"Error() string":                                                                                        "error",
	"ReadByte() (byte, error)":                                                                              "io.ByteReader",
	"ReadByte() (byte, error); UnreadByte() error":                                                          "io.ByteScanner",
	"WriteByte(byte) error":                                                                                 "io.ByteWriter",
	"Close() error":                                                                                         "io.Closer",
	"Close() error; Read([]byte) (int, error)":                                                              "io.ReadCloser",
	"Read([]byte) (int, error); Seek(int64, int) (int64, error)":                                            "io.ReadSeeker",
	"Close() error; Read([]byte) (int, error); Write([]byte) (int, error)":                                  "io.ReadWriteCloser",
	"Read([]byte) (int, error); Seek(int64, int) (int64, error); Write([]byte) (int, error)":                "io.ReadWriteSeeker",
	"Read([]byte) (int, error); Write([]byte) (int, error)":                                                 "io.ReadWriter",
	"Read([]byte) (int, error)":                                                                             "io.Reader",
	"ReadAt([]byte, int64) (int, error)":                                                                    "io.ReaderAt",
	"ReadFrom(io.Reader) (int64, error)":                                                                    "io.ReaderFrom",
	"ReadRune() (rune, int, error)":                                                                         "io.RuneReader",
	"ReadRune() (rune, int, error); UnreadRune() error":                                                     "io.RuneScanner",
	"Seek(int64, int) (int64, error)":                                                                       "io.Seeker",
	"Close() error; Write([]byte) (int, error)":                                                             "io.WriteCloser",
	"Seek(int64, int) (int64, error); Write([]byte) (int, error)":                                           "io.WriteSeeker",
	"Write([]byte) (int, error)":                                                                            "io.Writer",
	"WriteAt([]byte, int64) (int, error)":                                                                   "io.WriterAt",
	"WriteTo(io.Writer) (int64, error)":                                                                     "io.WriterTo",
	"IsDir() bool; ModTime() time.Time; Mode() os.FileMode; Name() string; Size() int64; Sys() interface{}": "os.FileInfo",
	"Signal(); String() string":                                                                             "os.Signal",
	"Format(fmt.State, rune)":                                                                               "fmt.Formatter",
	"GoString() string":                                                                                     "fmt.GoStringer",
	"Read([]byte) (int, error); ReadRune() (rune, int, error); SkipSpace(); Token(bool, func(rune) bool) ([]byte, error); UnreadRune() error; Width() (int, bool)": "fmt.ScanState",
	"Scan(fmt.ScanState, rune) error":                                                          "fmt.Scanner",
	"Flag(int) bool; Precision() (int, bool); Width() (int, bool); Write([]byte) (int, error)": "fmt.State",
	"String() string":                   "fmt.Stringer",
	"Network() string; String() string": "net.Addr",
	"Close() error; LocalAddr() net.Addr; Read([]byte) (int, error); RemoteAddr() net.Addr; SetDeadline(time.Time) error; SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error; Write([]byte) (int, error)": "net.Conn",
	"Error() string; Temporary() bool; Timeout() bool":                                                                                                                                                                        "net.Error",
	"Accept() (net.Conn, error); Addr() net.Addr; Close() error":                                                                                                                                                              "net.Listener",
	"Close() error; LocalAddr() net.Addr; ReadFrom([]byte) (int, net.Addr, error); SetDeadline(time.Time) error; SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error; WriteTo([]byte, net.Addr) (int, error)": "net.PacketConn",
	"Get() interface{}; Set(string) error; String() string":                                                                                                                                                                   "flag.Getter",
	"Set(string) error; String() string":                                                                                                                                                                                      "flag.Value",
	"BlockSize() int; Reset(); Size() int; Sum([]byte) []byte; Write([]byte) (int, error)":                                                                                                                                    "hash.Hash",
	"BlockSize() int; Reset(); Size() int; Sum([]byte) []byte; Sum32() uint32; Write([]byte) (int, error)":                                                                                                                    "hash.Hash32",
	"BlockSize() int; Reset(); Size() int; Sum([]byte) []byte; Sum64() uint64; Write([]byte) (int, error)":                                                                                                                    "hash.Hash64",
	"Len() int; Less(int, int) bool; Swap(int, int)":                                                                                                                                                                          "sort.Interface",
	"Lock(); Unlock()":                                                                                                                                                                                                        "sync.Locker",
	"At(int, int) image/color.Color; Bounds() image.Rectangle; ColorModel() image/color.Model":                                                           "image.Image",
	"At(int, int) image/color.Color; Bounds() image.Rectangle; ColorIndexAt(int, int) uint8; ColorModel() image/color.Model":                             "image.PalettedImage",
	"Decrypt(io.Reader, []byte, crypto.DecrypterOpts) ([]byte, error); Public() crypto.PublicKey":                                                        "crypto.Decrypter",
	"Public() crypto.PublicKey; Sign(io.Reader, []byte, crypto.SignerOpts) ([]byte, error)":                                                              "crypto.Signer",
	"HashFunc() crypto.Hash":                                                                                                                             "crypto.SignerOpts",
	"End() go/token.Pos; Pos() go/token.Pos":                                                                                                             "go/ast.Node",
	"Visit(go/ast.Node) go/ast.Visitor":                                                                                                                  "go/ast.Visitor",
	"Deadline() (time.Time, bool); Done() <-chan struct{}; Err() error; Value(interface{}) interface{}":                                                  "context.Context",
	"Close() error; ReadResponseBody(interface{}) error; ReadResponseHeader(*net/rpc.Response) error; WriteRequest(*net/rpc.Request, interface{}) error": "net/rpc.ClientCodec",
	"Close() error; ReadRequestBody(interface{}) error; ReadRequestHeader(*net/rpc.Request) error; WriteResponse(*net/rpc.Response, interface{}) error":  "net/rpc.ServerCodec",
	"Error() string; RuntimeError()":                                                                                                              "runtime.Error",
	"MarshalBinary() ([]byte, error)":                                                                                                             "encoding.BinaryMarshaler",
	"UnmarshalBinary([]byte) error":                                                                                                               "encoding.BinaryUnmarshaler",
	"MarshalText() ([]byte, error)":                                                                                                               "encoding.TextMarshaler",
	"UnmarshalText([]byte) error":                                                                                                                 "encoding.TextUnmarshaler",
	"Import(string) (*go/types.Package, error)":                                                                                                   "go/types.Importer",
	"Import(string) (*go/types.Package, error); ImportFrom(string, string, go/types.ImportMode) (*go/types.Package, error)":                       "go/types.ImporterFrom",
	"Alignof(go/types.Type) int64; Offsetsof([]*go/types.Var) []int64; Sizeof(go/types.Type) int64":                                               "go/types.Sizes",
	"String() string; Underlying() go/types.Type":                                                                                                 "go/types.Type",
	"CloseNotify() <-chan bool":                                                                                                                   "net/http.CloseNotifier",
	"Cookies(*net/url.URL) []*net/http.Cookie; SetCookies(*net/url.URL, []*net/http.Cookie)":                                                      "net/http.CookieJar",
	"Close() error; Read([]byte) (int, error); Readdir(int) ([]os.FileInfo, error); Seek(int64, int) (int64, error); Stat() (os.FileInfo, error)": "net/http.File",
	"Open(string) (net/http.File, error)":                                                                                                         "net/http.FileSystem",
	"Flush()": "net/http.Flusher",
	"ServeHTTP(net/http.ResponseWriter, *net/http.Request)":                                                                      "net/http.Handler",
	"Hijack() (net.Conn, *bufio.ReadWriter, error)":                                                                              "net/http.Hijacker",
	"Push(string, *net/http.PushOptions) error":                                                                                  "net/http.Pusher",
	"Header() net/http.Header; Write([]byte) (int, error); WriteHeader(int)":                                                     "net/http.ResponseWriter",
	"RoundTrip(*net/http.Request) (*net/http.Response, error)":                                                                   "net/http.RoundTripper",
	"Next([]byte, bool) ([]byte, error); Start(*net/smtp.ServerInfo) (string, []byte, error)":                                    "net/smtp.Auth",
	"Get() *image/png.EncoderBuffer; Put(*image/png.EncoderBuffer)":                                                              "image/png.EncoderBufferPool",
	"Int63() int64; Seed(int64)":                                                                                                 "math/rand.Source",
	"Int63() int64; Seed(int64); Uint64() uint64":                                                                                "math/rand.Source64",
	"Get(string) (*crypto/tls.ClientSessionState, bool); Put(string, *crypto/tls.ClientSessionState)":                            "crypto/tls.ClientSessionCache",
	"Draw(image/draw.Image, image.Rectangle, image.Image, image.Point)":                                                          "image/draw.Drawer",
	"At(int, int) image/color.Color; Bounds() image.Rectangle; ColorModel() image/color.Model; Set(int, int, image/color.Color)": "image/draw.Image",
	"Quantize(image/color.Palette, image.Image) image/color.Palette":                                                             "image/draw.Quantizer",
	"Read([]byte) (int, error); ReadByte() (byte, error)":                                                                        "image/jpeg.Reader",
	"Common() *debug/dwarf.CommonType; Size() int64; String() string":                                                            "debug/dwarf.Type",
	"Raw() []byte":                                                         "debug/macho.Load",
	"RGBA() (uint32, uint32, uint32, uint32)":                              "image/color.Color",
	"Convert(image/color.Color) image/color.Color":                         "image/color.Model",
	"LastInsertId() (int64, error); RowsAffected() (int64, error)":         "database/sql.Result",
	"Scan(interface{}) error":                                              "database/sql.Scanner",
	"GobDecode([]byte) error":                                              "encoding/gob.GobDecoder",
	"GobEncode() ([]byte, error)":                                          "encoding/gob.GobEncoder",
	"MarshalXML(*encoding/xml.Encoder, encoding/xml.StartElement) error":   "encoding/xml.Marshaler",
	"MarshalXMLAttr(encoding/xml.Name) (encoding/xml.Attr, error)":         "encoding/xml.MarshalerAttr",
	"UnmarshalXML(*encoding/xml.Decoder, encoding/xml.StartElement) error": "encoding/xml.Unmarshaler",
	"UnmarshalXMLAttr(encoding/xml.Attr) error":                            "encoding/xml.UnmarshalerAttr",
	"Reset(io.Reader, []byte) error":                                       "compress/zlib.Resetter",
	"NonceSize() int; Open([]byte, []byte, []byte, []byte) ([]byte, error); Overhead() int; Seal([]byte, []byte, []byte, []byte) []byte":                                                                                                                                                                                                                                                               "crypto/cipher.AEAD",
	"BlockSize() int; Decrypt([]byte, []byte); Encrypt([]byte, []byte)":                                                                                                                                                                                                                                                                                                                                "crypto/cipher.Block",
	"BlockSize() int; CryptBlocks([]byte, []byte)":                                                                                                                                                                                                                                                                                                                                                     "crypto/cipher.BlockMode",
	"XORKeyStream([]byte, []byte)":                                                                                                                                                                                                                                                                                                                                                                     "crypto/cipher.Stream",
	"MarshalJSON() ([]byte, error)":                                                                                                                                                                                                                                                                                                                                                                    "encoding/json.Marshaler",
	"UnmarshalJSON([]byte) error":                                                                                                                                                                                                                                                                                                                                                                      "encoding/json.Unmarshaler",
	"Generate(*math/rand.Rand, int) reflect.Value":                                                                                                                                                                                                                                                                                                                                                     "testing/quick.Generator",
	"Len() int; Less(int, int) bool; Pop() interface{}; Push(interface{}); Swap(int, int)":                                                                                                                                                                                                                                                                                                             "container/heap.Interface",
	"Close() error; Read([]byte) (int, error); ReadAt([]byte, int64) (int, error); Seek(int64, int) (int64, error)":                                                                                                                                                                                                                                                                                    "mime/multipart.File",
	"Add(*math/big.Int, *math/big.Int, *math/big.Int, *math/big.Int) (*math/big.Int, *math/big.Int); Double(*math/big.Int, *math/big.Int) (*math/big.Int, *math/big.Int); IsOnCurve(*math/big.Int, *math/big.Int) bool; Params() *crypto/elliptic.CurveParams; ScalarBaseMult([]byte) (*math/big.Int, *math/big.Int); ScalarMult(*math/big.Int, *math/big.Int, []byte) (*math/big.Int, *math/big.Int)": "crypto/elliptic.Curve",
	"PutUint16([]byte, uint16); PutUint32([]byte, uint32); PutUint64([]byte, uint64); String() string; Uint16([]byte) uint16; Uint32([]byte) uint32; Uint64([]byte) uint64":                                                                                                                                                                                                                            "encoding/binary.ByteOrder",
	"Get() []byte; Put([]byte)":                                                                                                                                                  "net/http/httputil.BufferPool",
	"PublicSuffix(string) string; String() string":                                                                                                                               "net/http/cookiejar.PublicSuffixList",
	"ColumnConverter(int) database/sql/driver.ValueConverter":                                                                                                                    "database/sql/driver.ColumnConverter",
	"Begin() (database/sql/driver.Tx, error); Close() error; Prepare(string) (database/sql/driver.Stmt, error)":                                                                  "database/sql/driver.Conn",
	"BeginTx(context.Context, database/sql/driver.TxOptions) (database/sql/driver.Tx, error)":                                                                                    "database/sql/driver.ConnBeginTx",
	"PrepareContext(context.Context, string) (database/sql/driver.Stmt, error)":                                                                                                  "database/sql/driver.ConnPrepareContext",
	"Open(string) (database/sql/driver.Conn, error)":                                                                                                                             "database/sql/driver.Driver",
	"Exec(string, []database/sql/driver.Value) (database/sql/driver.Result, error)":                                                                                              "database/sql/driver.Execer",
	"ExecContext(context.Context, string, []database/sql/driver.NamedValue) (database/sql/driver.Result, error)":                                                                 "database/sql/driver.ExecerContext",
	"Ping(context.Context) error":                                                                                                                                                "database/sql/driver.Pinger",
	"Query(string, []database/sql/driver.Value) (database/sql/driver.Rows, error)":                                                                                               "database/sql/driver.Queryer",
	"QueryContext(context.Context, string, []database/sql/driver.NamedValue) (database/sql/driver.Rows, error)":                                                                  "database/sql/driver.QueryerContext",
	"Close() error; Columns() []string; Next([]database/sql/driver.Value) error":                                                                                                 "database/sql/driver.Rows",
	"Close() error; ColumnTypeDatabaseTypeName(int) string; Columns() []string; Next([]database/sql/driver.Value) error":                                                         "database/sql/driver.RowsColumnTypeDatabaseTypeName",
	"Close() error; ColumnTypeLength(int) (int64, bool); Columns() []string; Next([]database/sql/driver.Value) error":                                                            "database/sql/driver.RowsColumnTypeLength",
	"Close() error; ColumnTypeNullable(int) (bool, bool); Columns() []string; Next([]database/sql/driver.Value) error":                                                           "database/sql/driver.RowsColumnTypeNullable",
	"Close() error; ColumnTypePrecisionScale(int) (int64, int64, bool); Columns() []string; Next([]database/sql/driver.Value) error":                                             "database/sql/driver.RowsColumnTypePrecisionScale",
	"Close() error; ColumnTypeScanType(int) reflect.Type; Columns() []string; Next([]database/sql/driver.Value) error":                                                           "database/sql/driver.RowsColumnTypeScanType",
	"Close() error; Columns() []string; HasNextResultSet() bool; Next([]database/sql/driver.Value) error; NextResultSet() error":                                                 "database/sql/driver.RowsNextResultSet",
	"Close() error; Exec([]database/sql/driver.Value) (database/sql/driver.Result, error); NumInput() int; Query([]database/sql/driver.Value) (database/sql/driver.Rows, error)": "database/sql/driver.Stmt",
	"ExecContext(context.Context, []database/sql/driver.NamedValue) (database/sql/driver.Result, error)":                                                                         "database/sql/driver.StmtExecContext",
	"QueryContext(context.Context, []database/sql/driver.NamedValue) (database/sql/driver.Rows, error)":                                                                          "database/sql/driver.StmtQueryContext",
	"Commit() error; Rollback() error":                                                                                                                                           "database/sql/driver.Tx",
	"ConvertValue(interface{}) (database/sql/driver.Value, error)":                                                                                                               "database/sql/driver.ValueConverter",
	"Value() (database/sql/driver.Value, error)":                                                                                                                                 "database/sql/driver.Valuer",
	"Reset(); Span([]byte, bool) (int, error); Transform([]byte, []byte, bool) (int, int, error)":                                                                                "vendor/golang_org/x/text/transform.SpanningTransformer",
	"Reset(); Transform([]byte, []byte, bool) (int, int, error)":                                                                                                                 "vendor/golang_org/x/text/transform.Transformer",
}

var stdFuncs = map[string]bool{
	"(fmt.State, rune)":                                                                               true,
	"(fmt.ScanState, rune) error":                                                                     true,
	"(time.Time) error":                                                                               true,
	"([]byte, net.Addr) (int, error)":                                                                 true,
	"(string, reflect.Value) bool":                                                                    true,
	"(go/ast.Node) go/ast.Visitor":                                                                    true,
	"(*go/types.Package) string":                                                                      true,
	"(go/types.Type) int64":                                                                           true,
	"(*net/url.URL) []*net/http.Cookie":                                                               true,
	"(*net/url.URL, []*net/http.Cookie)":                                                              true,
	"(net/http.ResponseWriter, *net/http.Request)":                                                    true,
	"(*net/http.Request) (*net/http.Response, error)":                                                 true,
	"(go/token.Position, string)":                                                                     true,
	"(image/draw.Image, image.Rectangle, image.Image, image.Point)":                                   true,
	"(image/color.Palette, image.Image) image/color.Palette":                                          true,
	"(*encoding/xml.Encoder, encoding/xml.StartElement) error":                                        true,
	"(*encoding/xml.Decoder, encoding/xml.StartElement) error":                                        true,
	"(string, os.FileInfo, error) error":                                                              true,
	"(*math/rand.Rand, int) reflect.Value":                                                            true,
	"(*math/big.Int, *math/big.Int, *math/big.Int, *math/big.Int) (*math/big.Int, *math/big.Int)":     true,
	"(*math/big.Int, *math/big.Int) (*math/big.Int, *math/big.Int)":                                   true,
	"(*math/big.Int, *math/big.Int) bool":                                                             true,
	"(*math/big.Int, *math/big.Int, []byte) (*math/big.Int, *math/big.Int)":                           true,
	"(context.Context, database/sql/driver.TxOptions) (database/sql/driver.Tx, error)":                true,
	"(context.Context, string) (database/sql/driver.Stmt, error)":                                     true,
	"(context.Context, string, []database/sql/driver.NamedValue) (database/sql/driver.Result, error)": true,
	"(context.Context) error":                                                                         true,
	"(context.Context, string, []database/sql/driver.NamedValue) (database/sql/driver.Rows, error)":   true,
	"(context.Context, []database/sql/driver.NamedValue) (database/sql/driver.Result, error)":         true,
	"(context.Context, []database/sql/driver.NamedValue) (database/sql/driver.Rows, error)":           true,
}
