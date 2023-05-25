package fuzz_gommit_messagequery

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/antham/gommit/gommit"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            case 1:
                var testMatch gommit.MessageQuery
                fuzzConsumer.GenerateStruct(&testMatch)

                gommit.MatchMessageQuery(testMatch)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}