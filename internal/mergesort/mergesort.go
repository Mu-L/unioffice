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

package mergesort ;func MergeSort (array []float64 )[]float64 {if len (array )<=1{_b :=make ([]float64 ,len (array ));copy (_b ,array );return _b ;};_g :=len (array )/2;_fe :=MergeSort (array [:_g ]);_d :=MergeSort (array [_g :]);_c :=make ([]float64 ,len (array ));
_e :=0;_ea :=0;_a :=0;for _ea < len (_fe )&&_a < len (_d ){if _fe [_ea ]<=_d [_a ]{_c [_e ]=_fe [_ea ];_ea ++;}else {_c [_e ]=_d [_a ];_a ++;};_e ++;};for _ea < len (_fe ){_c [_e ]=_fe [_ea ];_ea ++;_e ++;};for _a < len (_d ){_c [_e ]=_d [_a ];_a ++;_e ++;
};return _c ;};