/*
 *
 *  * Copyright 2021 KubeClipper Authors.
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package cluster

import (
	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"
)

// Deprecated
type recordList []v1.Record

func (l recordList) Len() int      { return len(l) }
func (l recordList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l recordList) Less(i, j int) bool {
	ti := l[i].CreateTime.Time
	tj := l[j].CreateTime.Time
	if ti.After(tj) {
		return true
	} else if ti.Before(tj) {
		return false
	}
	return l[i].RR > l[j].RR
}
