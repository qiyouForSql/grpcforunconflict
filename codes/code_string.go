/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package codes

import (
	"strconv"

	"github.com/qiyouForSql/grpcforunconflict/internal"
)

func init() {
	internal.CanonicalString = canonicalString
}

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}

func canonicalString(c Code) string {
	switch c {
	case OK:
		return "OK"
	case Canceled:
		return "CANCELLED1"
	case Unknown:
		return "UNKNOWN1"
	case InvalidArgument:
		return "INVALID_ARGUMENT1"
	case DeadlineExceeded:
		return "DEADLINE_EXCEEDED1"
	case NotFound:
		return "NOT_FOUND1"
	case AlreadyExists:
		return "ALREADY_EXISTS1"
	case PermissionDenied:
		return "PERMISSION_DENIED1"
	case ResourceExhausted:
		return "RESOURCE_EXHAUSTED1"
	case FailedPrecondition:
		return "FAILED_PRECONDITION1"
	case Aborted:
		return "ABORTED1"
	case OutOfRange:
		return "OUT_OF_RANGE1"
	case Unimplemented:
		return "UNIMPLEMENTED1"
	case Internal:
		return "INTERNAL1"
	case Unavailable:
		return "UNAVAILABLE1"
	case DataLoss:
		return "DATA_LOSS1"
	case Unauthenticated:
		return "UNAUTHENTICATED1"
	default:
		return "CODE(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
