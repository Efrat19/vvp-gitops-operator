package appmanagervvp

// import (
// 	"testing"

// 	mocks "efrat19.io/vvp-gitops-operator/mocks/vvp_client"
// 	"github.com/golang/mock/gomock"
// 	user "github.com/golang/mock/sample"
// 	"github.com/golang/mock/sample/imp1"
// appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/appmanager.vvp/v1alpha1"
// appmanagervvpv1alpha1 "efrat19.io/vvp-gitops-operator/apis/appmanager.vvp/v1alpha1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// )

// var _ = Describe("Deployment controller", func() {

// const (
//     timeout  = time.Second * 10
//     duration = time.Second * 10
//     interval = time.Millisecond * 250
// )

// Context("When creating ", func() {
//     It("Should increase CronJob Status.Active count when new Jobs are created", func() {
//         By("By creating a new CronJob")
//         ctx := context.Background()
//         cronJob := &cronjobv1.CronJob{
//             TypeMeta: metav1.TypeMeta{
//                 APIVersion: "batch.tutorial.kubebuilder.io/v1",
//                 Kind:       "CronJob",
//             },
//             ObjectMeta: metav1.ObjectMeta{
//                 Name:      CronjobName,
//                 Namespace: CronjobNamespace,
//             },
//             Spec: cronjobv1.CronJobSpec{
//                 Schedule: "1 * * * *",
//                 JobTemplate: batchv1beta1.JobTemplateSpec{
//                     Spec: batchv1.JobSpec{
//                         // For simplicity, we only fill out the required fields.
//                         Template: v1.PodTemplateSpec{
//                             Spec: v1.PodSpec{
//                                 // For simplicity, we only fill out the required fields.
//                                 Containers: []v1.Container{
//                                     {
//                                         Name:  "test-container",
//                                         Image: "test-image",
//                                     },
//                                 },
//                                 RestartPolicy: v1.RestartPolicyOnFailure,
//                             },
//                         },
//                     },
//                 },
//             },
//         }
//         Expect(k8sClient.Create(ctx, cronJob)).Should(Succeed())

// func TestCreateResource(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	reconciler := DeploymentReconciler{
// 		vvpClient: mocks.NewMockVvpClient(ctrl),
// 	}
// 	exampleDeployment := &v1alpha1.Deployment{
// 		TypeMeta: metav1.TypeMeta{
// 			APIVersion: "appmanager.vvp.efrat19.io/v1alpha1",
// 			Kind:       "Deployment",
// 		},
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      "foo",
// 			Namespace: "default",
// 		},
// 		Spec: v1alpha1.DeploymentSpec{
// 			Metadata: appmanager_apis.DeploymentMetadata{
// 				Annotations:    nil,
// 				CreatedAt:       metav1.Now(),
// 				DisplayName:     "test",
// 				Id:              "test" ,
// 				Labels:          nil,
// 				ModifiedAt:      metav1.Now(),
// 				Name:            "test"       ,
// 				Namespace:       "default"    ,
// 				ResourceVersion: 1,
// 			},
// 			Spec:     appmanager_apis.DeploymentSpec{
// 				DeploymentTargetId:           "test"
// 				DeploymentTargetName:         "test"
// 				MaxJobCreationAttempts:       1
// 				MaxSavepointCreationAttempts: 1
// 				RestoreStrategy:              nil
// 				SessionClusterName:           "test"
// 				State:                        "test"
// 				Template:                     nil
// 				UpgradeStrategy:              nil

// 			},
// 			Status:   appmanager_apis.DeploymentStatus{
// 				Running nil,
// 				State   "RUNNING"
// 			},
// 		},
// 	}

// }

// func TestGrabPointer(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockIndex := NewMockIndex(ctrl)
// 	mockIndex.EXPECT().Ptr(gomock.Any()).SetArg(0, 7) // set first argument to 7

// 	i := user.GrabPointer(mockIndex)
// 	if i != 7 {
// 		t.Errorf("Expected 7, got %d", i)
// 	}
// }

// func TestEmbeddedInterface(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockEmbed := NewMockEmbed(ctrl)
// 	mockEmbed.EXPECT().RegularMethod()
// 	mockEmbed.EXPECT().EmbeddedMethod()
// 	mockEmbed.EXPECT().ForeignEmbeddedMethod()

// 	mockEmbed.RegularMethod()
// 	mockEmbed.EmbeddedMethod()
// 	var emb imp1.ForeignEmbedded = mockEmbed // also does interface check
// 	emb.ForeignEmbeddedMethod()
// }

// func TestExpectTrueNil(t *testing.T) {
// 	// Make sure that passing "nil" to EXPECT (thus as a nil interface value),
// 	// will correctly match a nil concrete type.
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockIndex := NewMockIndex(ctrl)
// 	mockIndex.EXPECT().Ptr(nil) // this nil is a nil interface{}
// 	mockIndex.Ptr(nil)          // this nil is a nil *int
// }

// func TestDoAndReturnSignature(t *testing.T) {
// 	t.Run("wrong number of return args", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		defer ctrl.Finish()

// 		mockIndex := NewMockIndex(ctrl)

// 		mockIndex.EXPECT().Slice(gomock.Any(), gomock.Any()).DoAndReturn(
// 			func(_ []int, _ []byte) {},
// 		)

// 		defer func() {
// 			if r := recover(); r == nil {
// 				t.Error("expected panic")
// 			}
// 		}()

// 		mockIndex.Slice([]int{0}, []byte("meow"))
// 	})

// 	t.Run("wrong type of return arg", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		defer ctrl.Finish()

// 		mockIndex := NewMockIndex(ctrl)

// 		mockIndex.EXPECT().Slice(gomock.Any(), gomock.Any()).DoAndReturn(
// 			func(_ []int, _ []byte) bool {
// 				return true
// 			})

// 		mockIndex.Slice([]int{0}, []byte("meow"))
// 	})
// }

// var _ = Describe("Deployment Controller", func() {
// 	var reconciler DeploymentReconciler

// 	BeforeEach(func() {
// 		reconciler = DeploymentReconciler{
// 			vvpClient: mocks.NewMockVvpClient(),
// 		}

// 	})
// })
