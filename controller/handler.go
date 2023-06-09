/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import ctrl "sigs.k8s.io/controller-runtime"

// ReconcileHandler will invoke all the operations to be performed as part of an object reconcile, managing the queue
// based on the operations' results.
func ReconcileHandler(operations []Operation) (ctrl.Result, error) {
	for _, operation := range operations {
		result, err := operation()

		switch {
		case err != nil || result.RequeueRequest:
			return ctrl.Result{RequeueAfter: result.RequeueDelay}, err
		case result.CancelRequest:
			return ctrl.Result{}, nil
		}
	}

	return ctrl.Result{}, nil
}
