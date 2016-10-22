set -euo pipefail
for itype in uint8 uint32 uint64 int8 int16 int32 int64 float32 float64; do
    rm -rf ${itype}mm
    mkdir -p ${itype}mm/

    utype=$(echo $itype | sed 's/\(.\)/\U\1/')
    echo $utype
    sed "s/uint16/$itype/g" uint16mm/mmslice.go | sed "s/Uint16/$utype/g" > ${itype}mm/mmslice.go
    sed "s/uint16/$itype/g" uint16mm/mmslice_test.go | sed "s/Uint16/$utype/g" > ${itype}mm/mmslice_test.go

    cd ${itype}mm && go test
    cd ..
done
