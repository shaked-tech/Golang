# Copyright 2021 The Kubernetes Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# origin: https://github.com/kubernetes/ingress-nginx/blob/main/images/custom-error-pages/rootfs/Dockerfile

FROM golang:1.17-alpine as builder
RUN apk add git

WORKDIR /go/src/k8s.io/ingress-nginx/images/custom-error-pages

COPY . .

RUN go get . && \
    CGO_ENABLED=0 go build -a -installsuffix cgo \
	-ldflags "-s -w" \
	-o nginx-errors .

# Use distroless as minimal base image to package the binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

ENV DefaultFormatVar="application/json"

COPY src/page-error-403.html /www/403.html
COPY src/page-error-404.html /www/404.html
COPY src/page-error-500.html /www/500.html
COPY src/page-error-503.html /www/503.html
COPY src/page-error-504.html /www/504.html
COPY src/page-error-511.html /www/511.html

COPY src/page-error-403.json /www/403.json
COPY src/page-error-404.json /www/404.json
COPY src/page-error-500.json /www/500.json
COPY src/page-error-503.json /www/503.json
COPY src/page-error-504.json /www/504.json
COPY src/page-error-511.json /www/511.json

COPY --from=builder /go/src/k8s.io/ingress-nginx/images/custom-error-pages/nginx-errors /
USER nonroot:nonroot

CMD ["/nginx-errors"]
