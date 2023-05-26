package fuzz_gommit

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
            case 0:
                var testMatch gommit.Matching
                fuzzConsumer.GenerateStruct(&testMatch)

                gommit.IsZeroMatching(&testMatch)
                return 0

            case 1:
                var testMatch gommit.CommitQuery
                fuzzConsumer.GenerateStruct(&testMatch)

                gommit.MatchCommitQuery(testMatch)
                return 0

            case 2:
                var testMatch gommit.RangeQuery
                fuzzConsumer.GenerateStruct(&testMatch)

                gommit.MatchRangeQuery(testMatch)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}