FROM golang:1.18 as builder

COPY . /gommit
WORKDIR /gommit/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_gommit && go-fuzz-build -libfuzzer -o fuzz_gommit.a && \
    clang -fsanitize=fuzzer fuzz_gommit.a -o fuzz_gommit.libfuzzer

RUN cd fuzz_gommit_messagequery && go-fuzz-build -libfuzzer -o fuzz_gommit_messagequery.a && \
    clang -fsanitize=fuzzer fuzz_gommit_messagequery.a -o fuzz_gommit_messagequery.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /gommit/mayhem/fuzz_gommit/fuzz_gommit.libfuzzer /
COPY --from=builder /gommit/mayhem/fuzz_gommit_messagequery/fuzz_gommit_messagequery.libfuzzer /