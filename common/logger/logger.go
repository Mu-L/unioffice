//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package logger ;import (_cd "fmt";_bb "io";_d "os";_b "path/filepath";_cc "runtime";);

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};func (_feb WriterLogger )logToWriter (_fa _bb .Writer ,_cb string ,_ca string ,_cca ...interface{}){_ddb (_fa ,_cb ,_ca ,_cca );};

// Debug logs debug message.
func (_fe WriterLogger )Debug (format string ,args ...interface{}){if _fe .LogLevel >=LogLevelDebug {_bgd :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_fe .logToWriter (_fe .Output ,_bgd ,format ,args ...);};};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// Trace logs trace message.
func (_fd WriterLogger )Trace (format string ,args ...interface{}){if _fd .LogLevel >=LogLevelTrace {_dec :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_fd .logToWriter (_fd .Output ,_dec ,format ,args ...);};};

// Notice logs notice message.
func (_ba ConsoleLogger )Notice (format string ,args ...interface{}){if _ba .LogLevel >=LogLevelNotice {_ee :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_ba .output (_d .Stdout ,_ee ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_abc WriterLogger )IsLogLevel (level LogLevel )bool {return _abc .LogLevel >=level };

// Warning logs warning message.
func (_abb WriterLogger )Warning (format string ,args ...interface{}){if _abb .LogLevel >=LogLevelWarning {_ga :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_abb .logToWriter (_abb .Output ,_ga ,format ,args ...);};};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// Error logs error message.
func (_dd WriterLogger )Error (format string ,args ...interface{}){if _dd .LogLevel >=LogLevelError {_aed :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_dd .logToWriter (_dd .Output ,_aed ,format ,args ...);};};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_ce string ,_f ...interface{});Warning (_bd string ,_a ...interface{});Notice (_g string ,_af ...interface{});Info (_bf string ,_e ...interface{});Debug (_db string ,_cec ...interface{});Trace (_cf string ,_fb ...interface{});
IsLogLevel (_ff LogLevel )bool ;};var Log Logger =DummyLogger {};

// Debug logs debug message.
func (_bbd ConsoleLogger )Debug (format string ,args ...interface{}){if _bbd .LogLevel >=LogLevelDebug {_fbg :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_bbd .output (_d .Stdout ,_fbg ,format ,args ...);};};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_dbf ConsoleLogger )IsLogLevel (level LogLevel )bool {return _dbf .LogLevel >=level };

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Warning logs warning message.
func (_ccg ConsoleLogger )Warning (format string ,args ...interface{}){if _ccg .LogLevel >=LogLevelWarning {_gd :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ccg .output (_d .Stdout ,_gd ,format ,args ...);};};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _bb .Writer )*WriterLogger {logger :=WriterLogger {Output :writer ,LogLevel :logLevel };return &logger ;};

// Info logs info message.
func (_ab ConsoleLogger )Info (format string ,args ...interface{}){if _ab .LogLevel >=LogLevelInfo {_eg :="\u005bI\u004e\u0046\u004f\u005d\u0020";_ab .output (_d .Stdout ,_eg ,format ,args ...);};};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Info logs info message.
func (_gg WriterLogger )Info (format string ,args ...interface{}){if _gg .LogLevel >=LogLevelInfo {_aea :="\u005bI\u004e\u0046\u004f\u005d\u0020";_gg .logToWriter (_gg .Output ,_aea ,format ,args ...);};};

// Trace logs trace message.
func (_ac ConsoleLogger )Trace (format string ,args ...interface{}){if _ac .LogLevel >=LogLevelTrace {_egg :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_ac .output (_d .Stdout ,_egg ,format ,args ...);};};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;
LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// DummyLogger does nothing.
type DummyLogger struct{};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _bb .Writer ;};

// Error logs error message.
func (_ae ConsoleLogger )Error (format string ,args ...interface{}){if _ae .LogLevel >=LogLevelError {_dc :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ae .output (_d .Stdout ,_dc ,format ,args ...);};};func _ddb (_fad _bb .Writer ,_gf string ,_fc string ,_egb ...interface{}){_ ,_afg ,_def ,_ag :=_cc .Caller (3);
if !_ag {_afg ="\u003f\u003f\u003f";_def =0;}else {_afg =_b .Base (_afg );};_aec :=_cd .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_gf ,_afg ,_def )+_fc +"\u000a";_cd .Fprintf (_fad ,_aec ,_egb ...);};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };func (_df ConsoleLogger )output (_ea _bb .Writer ,_da string ,_bg string ,_de ...interface{}){_ddb (_ea ,_da ,_bg ,_de ...);};

// Notice logs notice message.
func (_aeb WriterLogger )Notice (format string ,args ...interface{}){if _aeb .LogLevel >=LogLevelNotice {_acg :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_aeb .logToWriter (_aeb .Output ,_acg ,format ,args ...);};};