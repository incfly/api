# Copyright 2016 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################
#
def zlib_repositories(bind=True):
    BUILD = """
# Copyright 2016 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################
#

licenses(["notice"])

exports_files(["README"])

cc_library(
    name = "zlib",
    srcs = [
        "adler32.c",
        "crc32.c",
        "crc32.h",
        "deflate.c",
        "deflate.h",
        "infback.c",
        "inffast.c",
        "inffast.h",
        "inffixed.h",
        "inflate.c",
        "inflate.h",
        "inftrees.c",
        "inftrees.h",
        "trees.c",
        "trees.h",
        "zconf.h",
        "zutil.c",
        "zutil.h",
    ],
    hdrs = [
        "zlib.h",
    ],
    copts = [
        "-Wno-shift-negative-value",
        "-Wno-unknown-warning-option",
    ],
    defines = [
        "Z_SOLO",
    ],
    visibility = [
        "//visibility:public",
    ],
)
"""

    native.new_git_repository(
        name = "zlib_git",
        build_file_content = BUILD,
        commit = "50893291621658f355bc5b4d450a8d06a563053d",  # v1.2.8
        remote = "https://github.com/madler/zlib.git",
    )

    native.bind(
        name = "zlib",
        actual = "@zlib_git//:zlib"
    )

def nanopb_repositories(bind=True):
    BUILD = """
# Copyright 2016 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################
#

licenses(["notice"])

exports_files(["LICENSE.txt"])

cc_library(
    name = "nanopb",
    srcs = [
        "pb.h",
        "pb_common.c",
        "pb_common.h",
        "pb_decode.c",
        "pb_decode.h",
        "pb_encode.c",
        "pb_encode.h",
    ],
    hdrs = [":includes"],
    visibility = [
        "//visibility:public",
    ],
)

genrule(
    name = "includes",
    srcs = [
        "pb.h",
        "pb_common.h",
        "pb_decode.h",
        "pb_encode.h",
    ],
    outs = [
        "third_party/nanopb/pb.h",
        "third_party/nanopb/pb_common.h",
        "third_party/nanopb/pb_decode.h",
        "third_party/nanopb/pb_encode.h",
    ],
    cmd = "mkdir -p $(@D)/third_party/nanopb && cp $(SRCS) $(@D)/third_party/nanopb",
)
"""

    native.new_git_repository(
        name = "nanopb_git",
        build_file_content = BUILD,
        commit = "f8ac463766281625ad710900479130c7fcb4d63b",
        remote = "https://github.com/nanopb/nanopb.git",
    )

    native.bind(
        name = "nanopb",
        actual = "@nanopb_git//:nanopb",
    )

def grpc_repositories(bind=True):
    zlib_repositories(bind)
    nanopb_repositories(bind)

    native.git_repository(
        name = "grpc_git",
        commit = "d28417c856366df704200f544e72d31056931bce",
        remote = "https://github.com/grpc/grpc.git",
    )

    if bind:
        native.bind(
            name = "gpr",
            actual = "@grpc_git//:gpr",
        )

        native.bind(
            name = "grpc",
            actual = "@grpc_git//:grpc",
        )

        native.bind(
            name = "grpc_cpp_plugin",
            actual = "@grpc_git//:grpc_cpp_plugin",
        )

        native.bind(
            name = "grpc++",
            actual = "@grpc_git//:grpc++",
        )

        native.bind(
            name = "grpc_lib",
            actual = "@grpc_git//:grpc++_reflection",
        )

def googleapis_repositories(protobuf_repo="@protobuf_git//", bind=True):
    BUILD = """
# Copyright 2016 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################
#

licenses(["notice"])

load("{}:protobuf.bzl", "cc_proto_library")

exports_files(glob(["google/**"]))

cc_proto_library(
    name = "servicecontrol",
    srcs = [
        "google/api/servicecontrol/v1/check_error.proto",
        "google/api/servicecontrol/v1/distribution.proto",
        "google/api/servicecontrol/v1/log_entry.proto",
        "google/api/servicecontrol/v1/metric_value.proto",
        "google/api/servicecontrol/v1/operation.proto",
        "google/api/servicecontrol/v1/service_controller.proto",
        "google/logging/type/http_request.proto",
        "google/logging/type/log_severity.proto",
        "google/rpc/error_details.proto",
        "google/rpc/status.proto",
        "google/type/money.proto",
    ],
    include = ".",
    visibility = ["//visibility:public"],
    deps = [
        ":service_config",
    ],
    protoc = "//external:protoc",
    default_runtime = "//external:protobuf",
)

cc_proto_library(
    name = "service_config",
    srcs = [
        "google/api/annotations.proto",
        "google/api/auth.proto",
        "google/api/backend.proto",
        "google/api/billing.proto",
        "google/api/consumer.proto",
        "google/api/context.proto",
        "google/api/control.proto",
        "google/api/documentation.proto",
        "google/api/endpoint.proto",
        "google/api/http.proto",
        "google/api/label.proto",
        "google/api/log.proto",
        "google/api/logging.proto",
        "google/api/metric.proto",
        "google/api/monitored_resource.proto",
        "google/api/monitoring.proto",
        "google/api/service.proto",
        "google/api/system_parameter.proto",
        "google/api/usage.proto",
    ],
    include = ".",
    visibility = ["//visibility:public"],
    deps = [
        "//external:cc_wkt_protos",
    ],
    protoc = "//external:protoc",
    default_runtime = "//external:protobuf",
)

cc_proto_library(
    name = "cloud_trace",
    srcs = [
        "google/devtools/cloudtrace/v1/trace.proto",
    ],
    include = ".",
    default_runtime = "//external:protobuf",
    protoc = "//external:protoc",
    visibility = ["//visibility:public"],
    deps = [
        ":service_config",
        "//external:cc_wkt_protos",
    ],
)
""".format(protobuf_repo)


    native.new_git_repository(
        name = "googleapis_git",
        commit = "db1d4547dc56a798915e0eb2c795585385922165",
        remote = "https://github.com/googleapis/googleapis.git",
        build_file_content = BUILD,
    )

    if bind:
        native.bind(
            name = "servicecontrol",
            actual = "@googleapis_git//:servicecontrol",
        )

        native.bind(
            name = "servicecontrol_genproto",
            actual = "@googleapis_git//:servicecontrol_genproto",
        )

        native.bind(
            name = "service_config",
            actual = "@googleapis_git//:service_config",
        )

        native.bind(
            name = "cloud_trace",
            actual = "@googleapis_git//:cloud_trace",
        )

def servicecontrol_client_repositories(bind=True):
    googleapis_repositories(bind=bind)

    native.git_repository(
        name = "servicecontrol_client_git",
        commit = "d739d755365c6a13d0b4164506fd593f53932f5d",
        remote = "https://github.com/cloudendpoints/service-control-client-cxx.git",
    )

    if bind:
        native.bind(
            name = "servicecontrol_client",
            actual = "@servicecontrol_client_git//:service_control_client_lib",
        )

def mixerapi_repositories(protobuf_repo="@protobuf_git//", bind=True):
    BUILD = """
# Copyright 2016 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################
#
licenses(["notice"])

load("{}:protobuf.bzl", "cc_proto_library")

cc_proto_library(
    name = "mixer_api_cc_proto",
    srcs = glob(
        ["mixer/api/v1/*.proto"],
    ),
    default_runtime = "//external:protobuf",
    protoc = "//external:protoc",
    visibility = ["//visibility:public"],
    deps = [
        "//external:cc_wkt_protos",
        "//external:servicecontrol",
    ],
)
""".format(protobuf_repo)

    native.new_git_repository(
        name = "mixerapi_git",
        commit = "fc5a396185edc72d06d1937f30a8148a37d4fc1b",
        remote = "https://github.com/istio/api.git",
        build_file_content = BUILD,
    )
    if bind:
        native.bind(
            name = "mixer_api_cc_proto",
            actual = "@mixerapi_git//:mixer_api_cc_proto",
        )