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

package rw ;import (_fd "bytes";_fb "encoding/binary";_g "errors";_fg "fmt";_ca "io";_gd "io/ioutil";_c "reflect";);func NewReader (b []byte )(*Reader ,error ){return &Reader {_fd .NewReader (b )},nil };func (_fga *Writer )tryGrowByReslice (_ffd int )(int ,bool ){if _be :=len (_fga ._cc );_ffd <=cap (_fga ._cc )-_be {_fga ._cc =_fga ._cc [:_be +_ffd ];return _be ,true ;};return 0,false ;};func (_ggc *Writer )reset (){_ggc ._cc =_ggc ._cc [:0];_ggc ._cf =0};func (_cbc *Writer )Bytes ()[]byte {return _cbc ._cc };func PopRightUI64 (v uint64 )(bool ,uint64 ){return (v &uint64 (1))==1,v >>1};func (_gdg *Writer )WritePropertyNoAlign (a interface{})error {if _cee :=_fb .Write (_gdg ,_fb .LittleEndian ,a );_cee !=nil {return _cee ;};return nil ;};func (_bff *Writer )AlignLength (alignTo int )error {_ff :=_bff .Len ()%alignTo ;if _ff > 0{_ ,_ed :=_bff .Write (make ([]byte ,alignTo -_ff ));if _ed !=nil {return _ed ;};};return nil ;};func (_aff *Writer )WriteTo (wTo _ca .Writer )(_dba int64 ,_fe error ){if _aa :=_aff .Len ();_aa > 0{_bbc ,_baf :=wTo .Write (_aff ._cc [_aff ._cf :]);if _bbc > _aa {return 0,_g .New ("\u0072\u0077\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f\u003a\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0057\u0072\u0069t\u0065\u0020c\u006f\u0075\u006e\u0074");};_aff ._cf +=_bbc ;_dba =int64 (_bbc );if _baf !=nil {return _dba ,_baf ;};if _bbc !=_aa {return _dba ,_ca .ErrShortWrite ;};};_aff .reset ();return _dba ,nil ;};var _ab =_g .New ("r\u0077.\u0057\u0072\u0069\u0074\u0065\u0072\u003a\u0020t\u006f\u006f\u0020\u006car\u0067\u0065");func (_cb *Writer )align (_cda int )error {return _cb .Skip ((_cda -(_cb .Len ())%_cda )%_cda )};type Writer struct{_cc []byte ;_cf int ;};func (_bc *Writer )WriteStringProperty (s string )error {_bc .align (4);_dbe :=[]byte (s );if _de :=_fb .Write (_bc ,_fb .LittleEndian ,&_dbe );_de !=nil {return _de ;};return nil ;};func (_d *Reader )curPos ()int {return int (_d .Size ())-_d .Len ()};func (_egg *Writer )Cap ()int {return cap (_egg ._cc )};func (_gcg *Writer )WriteProperty (a interface{})error {if _ce :=_gcg .align (int (_c .TypeOf (a ).Size ()));_ce !=nil {return _ce ;};return _gcg .WritePropertyNoAlign (a );};type Reader struct{*_fd .Reader };func (_dfa *Writer )curPos ()int {return int (_dfa .Cap ())-_dfa .Len ()};func (_eg *Writer )WriteByteAt (b byte ,off int )error {if off >=len (_eg ._cc ){return _g .New ("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};_eg ._cc [off ]=b ;return nil ;};func (_bf *Reader )ReadProperty (a interface{})error {_ga :=_c .ValueOf (a );for _ga .Kind ()==_c .Ptr {_ga =_ga .Elem ();};if !_ga .IsValid (){return _fg .Errorf ("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064");};if _df :=_bf .align (int (_ga .Type ().Size ()));_df !=nil {return _df ;};if _eb :=_fb .Read (_bf ,_fb .LittleEndian ,a );_eb !=nil {return _eb ;};return nil ;};func (_b *Reader )skip (_ba int )error {_ ,_dc :=_ca .CopyN (_gd .Discard ,_b ,int64 (_ba ));if _dc !=nil {return _dc ;};return nil ;};func (_dd *Writer )Skip (n int )error {if n ==0{return nil ;};_ ,_cg :=_dd .Write (make ([]byte ,n ));return _cg ;};func (_ag *Reader )ReadPairProperty (p interface{})error {if _ec :=_ag .align (4);_ec !=nil {return _ec ;};_db :=_c .ValueOf (p );for _db .Kind ()==_c .Ptr {_db =_db .Elem ();};if !_db .IsValid (){return _fg .Errorf ("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064");};if _dfe :=_fb .Read (_ag ,_fb .LittleEndian ,p );_dfe !=nil {return _dfe ;};return nil ;};func (_abc *Writer )grow (_bcd int )(int ,error ){_bce :=_abc .Len ();if _bce ==0&&_abc ._cf !=0{_abc .reset ();};if _acg ,_acb :=_abc .tryGrowByReslice (_bcd );_acb {return _acg ,nil ;};if _abc ._cc ==nil &&_bcd <=_fgc {_abc ._cc =make ([]byte ,_bcd ,_fgc );return 0,nil ;};_fbed :=cap (_abc ._cc );if _bcd <=_fbed /2-_bce {copy (_abc ._cc ,_abc ._cc [_abc ._cf :]);}else if _fbed > _ae -_fbed -_bcd {return 0,_ab ;}else {_abb :=_fbe (2*_fbed +_bcd );copy (_abb ,_abc ._cc [_abc ._cf :]);_abc ._cc =_abb ;};_abc ._cf =0;_abc ._cc =_abc ._cc [:_bce +_bcd ];return _bce ,nil ;};func PushLeftUI64 (v uint64 ,flag bool )uint64 {v >>=1;if flag {v |=1<<63;};return v ;};func _fbe (_bfc int )[]byte {defer func (){if recover ()!=nil {panic (_ab );};}();return make ([]byte ,_bfc );};func (_fa *Reader )align (_a int )error {return _fa .skip ((_a -_fa .curPos ()%_a )%_a )};func PopRightUI32 (v uint32 )(bool ,uint32 ){return (v &uint32 (1))==1,v >>1};const _ae =int (^uint (0)>>1);func (_gg *Writer )Len ()int {return len (_gg ._cc )-_gg ._cf };func (_dcc *Reader )ReadStringProperty (n uint32 )(string ,error ){if _cd :=_dcc .align (4);_cd !=nil {return "",_cd ;};_ea :=make ([]byte ,n );if _gc :=_fb .Read (_dcc ,_fb .LittleEndian ,&_ea );_gc !=nil {return "",_gc ;};return string (_ea ),nil ;};func PushLeftUI32 (v uint32 ,flag bool )uint32 {v >>=1;if flag {v |=1<<31;};return v ;};func NewWriter ()*Writer {return &Writer {_cc :[]byte {}}};const _fgc =64;func (_af *Writer )Write (p []byte )(_ef int ,_ac error ){_def ,_bg :=_af .tryGrowByReslice (len (p ));if !_bg {var _gce error ;_def ,_gce =_af .grow (len (p ));if _gce !=nil {return 0,_gce ;};};return copy (_af ._cc [_def :],p ),nil ;};func (_afc *Writer )FillWithByte (fillSize int ,b byte )error {for _fgg :=0;_fgg < fillSize ;_fgg ++{if _cgd :=_afc .WritePropertyNoAlign (b );_cgd !=nil {return _cgd ;};};return nil ;};