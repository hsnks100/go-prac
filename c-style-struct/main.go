
package main

/*
#include <string.h>
#include <stdint.h>
typedef struct tagA {
    int32_t a;
    int32_t b;
}A;

A N={12,22};
*/
import "C"

import (
    "fmt"
    "encoding/binary"
    "bytes"
)

type A struct {
    Ea int32
    Eb int16
    Ec int16
    Ed [3]uint8
}

func main() {

    cbyte := make([]byte, 11)
    cbyte[0] = 2
    cbyte[1] = 0
    cbyte[2] = 0
    cbyte[3] = 0
    cbyte[4] = 5
    cbyte[5] = 0
    cbyte[6] = 7
    cbyte[7] = 0
    cbyte[8] = 97
    cbyte[9] = 98
    cbyte[10] = 0

    myStruct := A{}
    // myStruct.Ed = make([]byte, 3)
    _ = binary.Read(bytes.NewBuffer(cbyte[:]), binary.LittleEndian, &myStruct)
    fmt.Println(myStruct.Ea)
    fmt.Println(myStruct.Eb)
    fmt.Println(myStruct.Ec)
    fmt.Println(myStruct.Ed)


    buf := new(bytes.Buffer)
    _ = binary.Write(buf, binary.LittleEndian, &myStruct)
    fmt.Println(buf.Bytes())


    // t := C.A(cbyte)

    // s := &C.N // var s *C.struct_tagA = &C.N
    //
    // t := A{a: int64(s.a), b: int64(s.b)}
    // length := 0
    // for i, v := range s.c {
    //     t.c[i] = byte(v)
    //     if v == 0 {
    //         length = i
    //         break
    //     }
    // }
    //
    // fmt.Println("len(s.c):", len(s.c)) // 1024
    // str := string(t.c[0:length])       
    // fmt.Printf("len:%d %q \n", len(str), str) // len:4 "test" 
    //
    // s.a *= 10
    // fmt.Println(s.a) // 120
}
