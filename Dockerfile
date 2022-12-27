FROM scratch
COPY bazel-diff-as-a-service /
ENTRYPOINT ["/bazel-diff-as-a-service"]
