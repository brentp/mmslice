set -euo pipefail
for itype in uint8 uint32 uint64 int8 int16 int32 int64 float32 float64; do
    rm -rf ${itype}s
    mkdir -p ${itype}s/

    utype=$(echo $itype | sed 's/\(.\)/\U\1/')
    echo $utype
    sed "s/uint16/$itype/g" uint16s/mmslice.go | sed "s/Uint16/$utype/g" > ${itype}s/mmslice.go
    sed "s/uint16/$itype/g" uint16s/mmslice_test.go | sed "s/Uint16/$utype/g" > ${itype}s/mmslice_test.go

    cd ${itype}s && go test
    cd ..
done
