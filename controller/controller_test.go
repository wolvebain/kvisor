package controller

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"

	json "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/castai/sec-agent/version"
)

func TestController(t *testing.T) {
	ctx := context.Background()
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	t.Run("handle events", func(t *testing.T) {
		r := require.New(t)
		testNs := &corev1.Namespace{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name: "kube-system",
				ManagedFields: []metav1.ManagedFieldsEntry{
					{
						Manager:   "mng",
						Operation: "op",
					},
				},
			},
		}
		testDs := &appsv1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kube-proxy",
				Namespace: "kube-system",
				ManagedFields: []metav1.ManagedFieldsEntry{
					{
						Manager:   "mng",
						Operation: "op",
					},
				},
			},
		}
		clientset := fake.NewSimpleClientset(testNs, testDs)
		informersFactory := informers.NewSharedInformerFactory(clientset, 0)
		testSubs := []ObjectSubscriber{
			newTestSubscriber(log.WithField("sub", "sub1")),
			newTestSubscriber(log.WithField("sub", "sub2")),
		}
		ctrl := New(log, informersFactory, testSubs, version.Version{MinorInt: 22})

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		errc := make(chan error)
		go func() {
			if err := ctrl.Run(ctx); err != nil {
				errc <- err
			}
		}()

		for i := 0; i < 2; i++ {
			sub := testSubs[i].(*testSubscriber)
			select {
			case err := <-errc:
				t.Fatal(err)
			case <-time.After(1 * time.Millisecond):
				sub.assertNoManagedFields(r)
				sub.assertObjectMeta(r)
			}
		}
	})

	t.Run("skip events for unchanged daemon sets", func(t *testing.T) {
		r := require.New(t)
		testDs := &appsv1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kube-proxy",
				Namespace: "kube-system",
				ManagedFields: []metav1.ManagedFieldsEntry{
					{
						Manager:   "mng",
						Operation: "op",
					},
				},
			},
		}
		clientset := fake.NewSimpleClientset(testDs)
		informersFactory := informers.NewSharedInformerFactory(clientset, 0)
		testSubs := []ObjectSubscriber{
			newTestSubscriber(log.WithField("sub", "sub1")),
		}
		ctrl := New(log, informersFactory, testSubs, version.Version{MinorInt: 22})

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		errc := make(chan error)
		go func() {
			if err := ctrl.Run(ctx); err != nil {
				errc <- err
			}
		}()

		time.Sleep(100 * time.Millisecond)
		testDs.SetResourceVersion("new")
		_, err := clientset.AppsV1().DaemonSets(testDs.Namespace).Update(ctx, testDs, metav1.UpdateOptions{})
		r.NoError(err)
		time.Sleep(100 * time.Millisecond)
		err = clientset.AppsV1().DaemonSets(testDs.Namespace).Delete(ctx, testDs.Name, metav1.DeleteOptions{})
		r.NoError(err)

		sub := testSubs[0].(*testSubscriber)
		select {
		case err := <-errc:
			t.Fatal(err)
		case <-time.After(1 * time.Millisecond):
			sub.assertNoUpdates(r)
		}
	})
}

func TestObjectHash(t *testing.T) {
	r := require.New(t)
	js := `{
    "apiVersion": "apps/v1",
    "kind": "DaemonSet",
    "metadata": {
        "annotations": {
            "components.gke.io/layer": "addon",
            "deprecated.daemonset.template.generation": "5"
        },
        "creationTimestamp": "2022-04-19T16:20:05Z",
        "generation": 5,
        "labels": {
            "addonmanager.kubernetes.io/mode": "Reconcile",
            "k8s-app": "cilium"
        },
        "name": "anetd",
        "namespace": "kube-system",
        "resourceVersion": "148483317",
        "uid": "c0ed439e-8063-42ad-b523-427e7af5c927"
    },
    "spec": {
        "revisionHistoryLimit": 10,
        "selector": {
            "matchLabels": {
                "k8s-app": "cilium"
            }
        },
        "template": {
            "metadata": {
                "annotations": {
                    "components.gke.io/component-name": "advanceddatapath",
                    "components.gke.io/component-version": "2.5.21",
                    "prometheus.io/port": "9990",
                    "prometheus.io/scrape": "true"
                },
                "creationTimestamp": null,
                "labels": {
                    "k8s-app": "cilium"
                }
            },
            "spec": {
                "nodeSelector": {
                    "kubernetes.io/os": "linux"
                },
                "priorityClassName": "system-node-critical",
                "restartPolicy": "Always2",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "serviceAccount": "cilium",
                "serviceAccountName": "cilium",
                "terminationGracePeriodSeconds": 1,
                "tolerations": [
                    {
                        "operator": "Exists"
                    },
                    {
                        "key": "components.gke.io/gke-managed-components",
                        "operator": "Exists"
                    }
                ]
            }
        },
        "updateStrategy": {
            "rollingUpdate": {
                "maxSurge": 0,
                "maxUnavailable": 2
            },
            "type": "RollingUpdate"
        }
    },
    "status": {
        "currentNumberScheduled": 6,
        "desiredNumberScheduled": 6,
        "numberAvailable": 6,
        "numberMisscheduled": 0,
        "numberReady": 6,
        "observedGeneration": 5,
        "updatedNumberScheduled": 6
    }
}
`
	var ds1 appsv1.DaemonSet
	r.NoError(json.Unmarshal([]byte(js), &ds1))
	var ds2 appsv1.DaemonSet
	r.NoError(json.Unmarshal([]byte(js), &ds2))

	h1, err := ObjectHash(&ds1)
	r.NoError(err)
	h2, err := ObjectHash(&ds2)
	r.NoError(err)
	r.NotEmpty(h1)
	r.Equal(h1, h2)
}

func newTestSubscriber(log logrus.FieldLogger) *testSubscriber {
	return &testSubscriber{
		log:         log,
		addedObjs:   make(map[string]Object),
		updatedObjs: make(map[string]Object),
		deletedObjs: make(map[string]Object),
	}
}

type testSubscriber struct {
	log         logrus.FieldLogger
	mu          sync.Mutex
	addedObjs   map[string]Object
	updatedObjs map[string]Object
	deletedObjs map[string]Object
}

func (t *testSubscriber) assertNoManagedFields(r *require.Assertions) {
	r.Eventually(func() bool {
		t.mu.Lock()
		defer t.mu.Unlock()

		if len(t.addedObjs) != 2 {
			return false
		}
		obj := t.addedObjs["kube-system"]
		return len(obj.GetManagedFields()) == 0
	}, 3*time.Second, 1*time.Millisecond)
}

func (t *testSubscriber) assertObjectMeta(r *require.Assertions) {
	t.mu.Lock()
	defer t.mu.Unlock()

	obj := t.addedObjs["kube-proxy"]
	r.Equal("DaemonSet", obj.(*appsv1.DaemonSet).Kind)
}

func (t *testSubscriber) assertNoUpdates(r *require.Assertions) {
	r.Eventually(func() bool {
		t.mu.Lock()
		defer t.mu.Unlock()

		// We need to wait for delete in order to check if update happened or not.
		return len(t.deletedObjs) > 0 && len(t.updatedObjs) == 0
	}, 3*time.Second, 1*time.Millisecond)
}

func (t *testSubscriber) OnAdd(obj Object) {
	t.log.Debug("add")
	t.mu.Lock()
	defer t.mu.Unlock()
	t.addedObjs[obj.GetName()] = obj
}

func (t *testSubscriber) OnUpdate(obj Object) {
	t.log.Debug("update")
	t.mu.Lock()
	defer t.mu.Unlock()
	t.updatedObjs[obj.GetName()] = obj
}

func (t *testSubscriber) OnDelete(obj Object) {
	t.log.Debug("delete")
	t.mu.Lock()
	defer t.mu.Unlock()
	t.deletedObjs[obj.GetName()] = obj
}

func (t *testSubscriber) Run(ctx context.Context) error {
	t.log.Debug("run start")
	defer t.log.Debug("run done")

	for {
		select {
		case <-time.After(1 * time.Millisecond):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (t *testSubscriber) RequiredInformers() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf(&corev1.Namespace{}),
		reflect.TypeOf(&appsv1.DaemonSet{}),
	}
}
