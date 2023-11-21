load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    embed = [":cmd_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "cmd_lib",
    srcs = ["main.go"],
    importpath = "github.com/Iywkm/bazel-go/cmd",
    visibility = ["//visibility:private"],
    deps = ["//uuid"],
)
